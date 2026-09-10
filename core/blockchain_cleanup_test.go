package core

import (
	"bytes"
	"math/big"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
)

type cleanupCompactRange struct {
	start, limit []byte
}

type cleanupTestDB struct {
	ethdb.Database
	ranges        []cleanupCompactRange
	beforeCompact func()
}

func (db *cleanupTestDB) Compact(start, limit []byte) error {
	if db.beforeCompact != nil {
		db.beforeCompact()
	}
	db.ranges = append(db.ranges, cleanupCompactRange{common.CopyBytes(start), common.CopyBytes(limit)})
	return nil
}

func TestSidechainCompactPending(t *testing.T) {
	for _, stop := range []bool{false, true} {
		name := "drain"
		if stop {
			name = "stop"
		}
		t.Run(name, func(t *testing.T) {
			bc, db := newCleanupTestChain(t)
			entered, release := make(chan struct{}), make(chan struct{})
			var first, unblock sync.Once
			defer unblock.Do(func() { close(release) })
			db.beforeCompact = func() {
				first.Do(func() {
					close(entered)
					<-release
				})
			}
			bc.compactRangeAsync(1, sidechainBatch)
			select {
			case <-entered:
			case <-time.After(5 * time.Second):
				t.Fatal("compaction did not start")
			}
			bc.compactRangeAsync(sidechainBatch+1, 2*sidechainBatch)
			bc.compactRangeAsync(2*sidechainBatch+1, 3*sidechainBatch)
			// Empty scans must not enqueue new ranges or be needed to drain pending work.
			bc.maybeCleanupSidechains()
			if stop {
				close(bc.quit)
			}
			unblock.Do(func() { close(release) })
			bc.wg.Wait()
			if atomic.LoadInt32(&bc.compacting) != 0 {
				t.Fatal("compaction still marked running")
			}
			if stop {
				if len(db.ranges) != 1 {
					t.Fatalf("compacted after shutdown: %d ranges", len(db.ranges))
				}
				return
			}
			if len(db.ranges) != 6 {
				t.Fatalf("pending compaction not drained: want 6 ranges, got %d", len(db.ranges))
			}
			for i, prefix := range []byte{'h', 'b', 'r'} {
				wantStart := append([]byte{prefix}, encodeBlockNumber(sidechainBatch+1)...)
				wantLimit := append([]byte{prefix}, encodeBlockNumber(3*sidechainBatch+1)...)
				if r := db.ranges[i+3]; !bytes.Equal(r.start, wantStart) || !bytes.Equal(r.limit, wantLimit) {
					t.Fatalf("incorrect merged pending range: %x-%x", r.start, r.limit)
				}
			}
			if bc.compactFrom != 0 || bc.compactTo != 0 {
				t.Fatal("pending range remains after drain")
			}
			bc.compactRangeAsync(3*sidechainBatch+1, 4*sidechainBatch)
			bc.wg.Wait()
			if len(db.ranges) != 9 {
				t.Fatal("compaction did not restart after becoming idle")
			}
		})
	}
}

func newCleanupTestChain(t *testing.T) (*BlockChain, *cleanupTestDB) {
	t.Helper()
	db := &cleanupTestDB{Database: rawdb.NewMemoryDatabase()}
	bc := &BlockChain{
		db:          db,
		cacheConfig: &CacheConfig{TrieDirtyLimit: 256},
		stateCache:  state.NewDatabase(db),
		quit:        make(chan struct{}),
	}
	bc.currentBlock.Store(types.NewBlockWithHeader(&types.Header{Number: big.NewInt(cleanupLag + sidechainBatch)}))
	t.Cleanup(func() {
		bc.wg.Wait()
		db.Close()
	})
	return bc, db
}

func TestSidechainCleanupWithoutDeletes(t *testing.T) {
	bc, db := newCleanupTestChain(t)
	for height := uint64(1); height <= sidechainBatch; height++ {
		block := types.NewBlockWithHeader(&types.Header{Number: new(big.Int).SetUint64(height), Root: types.EmptyRootHash})
		rawdb.WriteBlock(db, block)
		rawdb.WriteCanonicalHash(db, block.Hash(), height)
	}
	bc.maybeCleanupSidechains()
	bc.wg.Wait()
	if len(db.ranges) != 0 {
		t.Fatalf("compacted without deleting blocks: %d ranges", len(db.ranges))
	}
	if progress := loadCleanupState(db); progress.LastSidechainCleanHeight != sidechainBatch {
		t.Fatalf("cleanup progress not persisted: %+v", progress)
	}
	for height := uint64(1); height <= sidechainBatch; height++ {
		hash := rawdb.ReadCanonicalHash(db, height)
		if rawdb.ReadBlock(db, hash, height) == nil {
			t.Fatalf("canonical block %d was deleted", height)
		}
	}
}

// A historical sidechain header can name a root owned by a newer canonical
// block. Deleting that header must not consume the canonical block's reference.
func TestSidechainCleanupPreservesReferencedState(t *testing.T) {
	bc, db := newCleanupTestChain(t)
	statedb, err := state.New(types.EmptyRootHash, bc.stateCache, nil)
	if err != nil {
		t.Fatal(err)
	}
	addr := common.Address{1}
	statedb.AddBalance(addr, big.NewInt(42))
	root, err := statedb.Commit(false)
	if err != nil {
		t.Fatal(err)
	}
	triedb := bc.stateCache.TrieDB()
	triedb.Reference(root, common.Hash{})
	if len(rawdb.ReadTrieNode(db, root)) != 0 {
		t.Fatal("test root must only exist in the dirty cache")
	}
	head := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(cleanupLag + sidechainBatch), Root: root})
	bc.currentBlock.Store(head)
	rawdb.WriteBlock(db, head)
	rawdb.WriteCanonicalHash(db, head.Hash(), head.NumberU64())
	canonical := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(1), Root: types.EmptyRootHash})
	side := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(1), Root: root})
	rawdb.WriteBlock(db, canonical)
	rawdb.WriteCanonicalHash(db, canonical.Hash(), 1)
	rawdb.WriteBlock(db, side)
	rawdb.WriteTd(db, side.Hash(), 1, big.NewInt(1))
	rawdb.WriteReceipts(db, side.Hash(), 1, types.Receipts{})
	bc.maybeCleanupSidechains()
	bc.wg.Wait()
	if rawdb.ReadHeader(db, side.Hash(), 1) != nil || rawdb.ReadBody(db, side.Hash(), 1) != nil ||
		rawdb.ReadTd(db, side.Hash(), 1) != nil || rawdb.ReadHeaderNumber(db, side.Hash()) != nil || rawdb.HasReceipts(db, side.Hash(), 1) {
		t.Fatal("sidechain block data remains")
	}
	for _, block := range []*types.Block{canonical, head} {
		if rawdb.ReadCanonicalHash(db, block.NumberU64()) != block.Hash() || rawdb.ReadBlock(db, block.Hash(), block.NumberU64()) == nil {
			t.Fatalf("canonical block %d was changed", block.NumberU64())
		}
	}
	reopened, err := state.New(root, bc.stateCache, nil)
	if err != nil {
		t.Fatalf("canonical state lost: %v", err)
	}
	if balance := reopened.GetBalance(addr); balance.Cmp(big.NewInt(42)) != 0 || reopened.Error() != nil {
		t.Fatalf("canonical state unreadable: balance %v, error %v", balance, reopened.Error())
	}
	if len(db.ranges) != 3 {
		t.Fatalf("want three compaction ranges after deletion, got %d", len(db.ranges))
	}
	// The original owner can still release the root exactly once.
	triedb.Dereference(root)
	if _, err := triedb.Node(root); err == nil {
		t.Fatal("root remains after its owner released it")
	}
}

func TestSidechainCompactRangeBoundary(t *testing.T) {
	bc, db := newCleanupTestChain(t)
	bc.compactRangeAsync(1, sidechainBatch)
	bc.wg.Wait()
	if len(db.ranges) != 3 {
		t.Fatalf("want three compaction ranges, got %d", len(db.ranges))
	}
	for i, prefix := range []byte{'h', 'b', 'r'} {
		r := db.ranges[i]
		key := func(height uint64, hash byte) []byte {
			return append(append([]byte{prefix}, encodeBlockNumber(height)...), bytes.Repeat([]byte{hash}, common.HashLength)...)
		}
		for _, height := range []uint64{1, sidechainBatch} {
			for _, hash := range []byte{0, 0xff} {
				k := key(height, hash)
				if bytes.Compare(k, r.start) < 0 || bytes.Compare(k, r.limit) >= 0 {
					t.Fatalf("range %q excludes block %d", prefix, height)
				}
			}
		}
		if bytes.Compare(key(0, 0xff), r.start) >= 0 || bytes.Compare(key(sidechainBatch+1, 0), r.limit) < 0 {
			t.Fatalf("range %q includes neighboring heights", prefix)
		}
	}
}
