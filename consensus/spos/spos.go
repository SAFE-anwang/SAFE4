// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package spos implements the SAFE-proof-of-stake consensus engine.
package spos

import (
	"bytes"
	"context"
	//"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
)

const (
	inmemorySignatures = 4096     // Number of recent block signatures to keep in memory
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	checkpointInterval = 1024     // Number of blocks after which to save the snapshot to the database
	validatorBytesLength = common.AddressLength

	wiggleTime = 5000 * time.Millisecond // Random delay (per signer) to allow concurrent signers
	sposAllowedFutureBlockTimeSeconds = int64(60)   // Max seconds from current time allowed for blocks, before they're considered future blocks

	//superNodeSPosCount = 7           //Total number of bookkeepers
	superNodeSPosCount = 3
	pushForwardHeight  = 14	          //Push forward the block height
	//chtAddress         = "0x043807066705c6EF9EB3D28D5D230b4d87EC4832" //Contract address
)

// Spos SAFE-proof-of-stake protocol constants.
var (
	epochLength = uint64(2000) // Default number of blocks after which to checkpoint and reset the pending votes

	extraVanity = 32                     // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = crypto.SignatureLength // Fixed number of extra-data suffix bytes reserved for signer seal

	BlockReward = big.NewInt(1e+18)

	subsidyHalvingInterval = big.NewInt(1051200)  //Number of blocks per year
	nextDecrementHeight =  big.NewInt(200)      //Half the height the next time

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.

	diffInTurn = big.NewInt(2) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(1) // Block difficulty for out-of-turn signatures
)

//var SposTxLock    sync.RWMutex
//var MinerRewardTx *types.Transaction
var ReceiptsLock  sync.RWMutex
var Receipts []*types.Receipt

func SetReceipts(receipts []*types.Receipt){
	ReceiptsLock.Lock()
	defer ReceiptsLock.Unlock()
	Receipts = make([]*types.Receipt, len(receipts))
	copy(Receipts, receipts)
}

func GetReceipts() []*types.Receipt{
	var receipts []*types.Receipt

	ReceiptsLock.Lock()
	receipts = make([]*types.Receipt, len(Receipts))
	copy(receipts, Receipts)
	ReceiptsLock.Unlock()
	return receipts
}

var TxsLock  sync.RWMutex
var Txs []*types.Transaction
func SetTxs(txs []*types.Transaction){
	TxsLock.Lock()
	defer TxsLock.Unlock()
	Txs = make([]*types.Transaction, len(txs))
	copy(Txs, txs)
}

func GetTxs() []*types.Transaction{
	var txs []*types.Transaction

	TxsLock.Lock()
	txs = make([]*types.Transaction, len(Receipts))
	copy(txs, Txs)
	TxsLock.Unlock()
	return txs
}

var CompleteBlockLock sync.RWMutex
var CompleteBlockFlag bool

func SetCompleteBlockFlag(flag bool) {
	CompleteBlockLock.Lock()
	defer CompleteBlockLock.Unlock()
	CompleteBlockFlag = flag
}

func GetCompleteBlockFlag() bool {
	var flag bool
	CompleteBlockLock.Lock()
	defer CompleteBlockLock.Unlock()
	flag = CompleteBlockFlag
	return flag
}


var TxPool *core.TxPool
func SetTxPool(txpool *core.TxPool) {
	TxPool = txpool
}

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte signature suffix missing")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block neither 1 or 2.
	errInvalidDifficulty = errors.New("invalid difficulty")

	// errInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	errInvalidTimestamp = errors.New("invalid timestamp")

	// errInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errInvalidVotingChain = errors.New("invalid voting chain")

	// errUnauthorizedSigner is returned if a header is signed by a non-authorized entity.
	errUnauthorizedSigner = errors.New("unauthorized signer")

	// errCoinBaseMisMatch is returned if a header's coinbase do not match with signature
	errCoinBaseMisMatch = errors.New("coinbase do not match with signature")

	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// invalid list of signers (i.e. non divisible by 20 bytes).
	errInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")

	// errMismatchingCheckpointSigners is returned if a checkpoint block contains a
	// list of signers different than the one the local node calculated.
	errMismatchingCheckpointSigners = errors.New("mismatching signer list on checkpoint block")

	// errUnauthorizedValidator is returned if a header is signed by a non-authorized entity.
	errUnauthorizedValidator = errors.New("unauthorized validator")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	//errRecentlySigned = errors.New("recently signed")

	// errWrongDifficulty is returned if the difficulty of a block doesn't match the
	// turn of the signer.
	errWrongDifficulty = errors.New("wrong difficulty")

	// errBlockHashInconsistent is returned if an authorization list is attempted to
	// insert an inconsistent block.
	errBlockHashInconsistent = errors.New("the block hash is inconsistent")
)

// SignerFn hashes and signs the data to be signed by a backing account.
type SignerFn func(signer accounts.Account, mimeType string, message []byte) ([]byte, error)
type SignerTxFn func(signer accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

// Spos is the SAFE-proof-of-stake consensus engine proposed to support the
// Ethereum testnet following the Ropsten attacks.
type Spos struct {
	chainConfig *params.ChainConfig
	config *params.SposConfig   // Consensus engine configuration parameters
	db     ethdb.Database       // Database to store and retrieve snapshot checkpoints

	recents     *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures  *lru.ARCCache // Signatures of recent blocks to speed up mining

	proposals map[common.Address]bool // Current list of proposals we are pushing

	signer common.Address // Ethereum address of the signing key
	signFn SignerFn       // Signer function to authorize hashes with
	signTxFn SignerTxFn
	lock   sync.RWMutex   // Protects the signer and proposals fields

	// The fields below are for testing only
	fakeDiff bool // Skip difficulty verifications

	//etherbaseprivatekey *ecdsa.PrivateKey

	blockChainAPI *ethapi.PublicBlockChainAPI
}

// New creates a Spos SAFE-proof-of-stack consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.ChainConfig, db ethdb.Database, blockChainAPI *ethapi.PublicBlockChainAPI) *Spos {
	// Set any missing consensus parameters to their defaults
	conf := config.Spos
	if conf != nil && conf.Epoch == 0 {
		conf.Epoch = epochLength
	}
	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	return &Spos{
		chainConfig: config,
		config:     conf,
		db:         db,
		recents:    recents,
		signatures: signatures,
		proposals:  make(map[common.Address]bool),
		blockChainAPI: blockChainAPI,
	}
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (s *Spos) Author(header *types.Header) (common.Address, error) {
	return ecrecover(header, s.signatures)
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (s *Spos) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	return s.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (s *Spos) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := s.verifyHeader(chain, header, headers[:i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (s *Spos) verifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix() + sposAllowedFutureBlockTimeSeconds){
		return consensus.ErrFutureBlock
	}

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}

	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}

	// Checkpoint blocks need to enforce zero beneficiary
	isEpoch := (number % s.config.Epoch) == 0

	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	signersBytes := len(header.Extra) - extraVanity - extraSeal
	if !isEpoch && signersBytes != 0 {
		return errExtraSigners
	}

	if isEpoch && signersBytes%validatorBytesLength != 0 {
		return errInvalidCheckpointSigners
	}

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in Spos
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil  {
			return errInvalidDifficulty
		}
	}

	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return s.verifyCascadingFields(chain, header, parents)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (s *Spos) verifyCascadingFields(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to its parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}

	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time + s.config.Period > header.Time {
		return errInvalidTimestamp
	}

	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := s.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	// If the block is a checkpoint block, verify the signer list
	if number%s.config.Epoch == 0 {
		signers := make([]byte, len(snap.Signers)*common.AddressLength)
		for i, signer := range snap.signers() {
			copy(signers[i*common.AddressLength:], signer[:])
		}
		extraSuffix := len(header.Extra) - extraSeal
		if !bytes.Equal(header.Extra[extraVanity:extraSuffix], signers) {
			return errMismatchingCheckpointSigners
		}
	}

	// Verify that the gas limit is <= 2^63-1
	capacity := uint64(0x7fffffffffffffff)
	if header.GasLimit > capacity {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, capacity)
	}

	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}

	// All basic checks passed, verify the seal and return
	return s.verifySeal(chain, header, parents)
}

// snapshot retrieves the authorization snapshot at a given point in time.
func (s *Spos) snapshot(chain consensus.ChainHeaderReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot
	)
	for snap == nil {
		// If an in-memory snapshot was found, use that
		if s, ok := s.recents.Get(hash); ok {
			snap = s.(*Snapshot)
			break
		}

		// If an on-disk checkpoint snapshot can be found, use that
		if number % checkpointInterval == 0 {
			if s, err := loadSnapshot(s.config, s.signatures, s.db, hash); err == nil {
				log.Trace("Loaded snapshot from disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}

		// If we're at the genesis, snapshot the initial state.
		if number == 0 || number % s.config.Epoch == 0{
			checkpoint := chain.GetHeaderByNumber(number)
			if checkpoint != nil {
				hash := checkpoint.Hash()

				var signers []common.Address
				if number == 0 {
					tempStartNewLoopTime := checkpoint.Time
					tempSignermap := make(map[common.Address]struct{})
					for _, signer := range params.SafeSposOfficialSuperNodeConfig.Signers {
						tempSignermap[signer] = struct{}{}
					}

					resultSuperNode := make([]common.Address, 0)
					resultSuperNode = sortSupernode(tempSignermap, checkpoint, tempStartNewLoopTime)
					signers = make([]common.Address, 0)

					icountsupernode := len(resultSuperNode)
					if icountsupernode > superNodeSPosCount {
						icountsupernode = superNodeSPosCount
					}

					for i := 0; i < icountsupernode; i++{
						signers =  append(signers, resultSuperNode[i])
						log.Info("singers info", "number", number, "signers[i]", signers[i])
					}
				}else{
					signers = make([]common.Address, (len(checkpoint.Extra)-extraVanity-extraSeal)/common.AddressLength)
					for i := 0; i < len(signers); i++ {
						copy(signers[i][:], checkpoint.Extra[extraVanity+i*common.AddressLength:])
					}
				}

				snap = newSnapshot(s.config, s.signatures, number, hash, signers)
				if err := snap.store(s.db); err != nil {
					return nil, err
				}
				log.Info("Stored checkpoint snapshot to disk", "number", number, "hash", hash)
				break
			}
		}else{
			checkpoint := chain.GetHeaderByNumber(number)
			if checkpoint != nil {
				hash := checkpoint.Hash()

				var startNewLoopTime uint64
				signersmap := make(map[common.Address]struct{})
				var signers []common.Address

				if number > pushForwardHeight {
					forwardHeight := number - pushForwardHeight
					forwardblock := chain.GetHeaderByNumber(forwardHeight)
					startNewLoopTime = forwardblock.Time
				}else{
					forwardblock := chain.GetHeaderByNumber(0)
					startNewLoopTime = forwardblock.Time
				}

				if number < params.SafeSposOfficialSuperNodeConfig.StartCommonSuperHeight {
					for _, signer := range params.SafeSposOfficialSuperNodeConfig.Signers {
						signersmap[signer] = struct{}{}
					}
				}else { //TODO Call the contract to get the super node list
					superMasterNodeInfos, err := systemcontracts.GetTopSuperMasterNode(s.blockChainAPI)
					if err != nil {
						log.Error("Failed to GetTopSMN", "error", err)
						return nil, err
					}

					for i := range superMasterNodeInfos {
						log.Info("Super MasterNode Addr Info", "superMasterNodeInfos[i].Addr", superMasterNodeInfos[i].Addr)
						signersmap[superMasterNodeInfos[i].Addr] = struct{}{}
					}
				}

				resultSuperNode := make([]common.Address, 0)
				resultSuperNode = sortSupernode(signersmap, checkpoint, startNewLoopTime)

				signers = make([]common.Address, 0)

				icountsupernode := len(resultSuperNode)
				if icountsupernode > superNodeSPosCount {
					icountsupernode = superNodeSPosCount
				}

				for i := 0; i < icountsupernode; i++{
					signers = append(signers, resultSuperNode[i])
					log.Info("singers info", "number", number, "i", i, "signers[i]", signers[i])
				}

				snap = newSnapshot(s.config, s.signatures, number, hash, signers)
				if err := snap.store(s.db); err != nil {
					return nil, err
				}
				log.Info("Stored checkpoint snapshot to disk", "number", number, "hash", hash)
				break
			}
		}

		// No snapshot for this header, gather the header and move backward
		var header *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			header = parents[len(parents)-1]
			if header.Hash() != hash || header.Number.Uint64() != number {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			header = chain.GetHeader(hash, number)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, header)
		number, hash = number-1, header.ParentHash
	}
	// check if snapshot is nil
	if snap == nil {
		return nil, fmt.Errorf("unknown error while retrieving snapshot at block number %v", number)
	}
	// Previous snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}

	snap, err := snap.apply(headers)
	if err != nil {
		return nil, err
	}
	s.recents.Add(snap.Hash, snap)

	// If we've generated a new checkpoint snapshot, save to disk
	if snap.Number % checkpointInterval == 0 && len(headers) > 0 {
		if err = snap.store(s.db); err != nil {
			return nil, err
		}
		log.Trace("Stored voting snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}
	return snap, err
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (s *Spos) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (s *Spos) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := s.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, s.signatures)
	if err != nil {
		return err
	}

	if signer != header.Coinbase {
		return errCoinBaseMisMatch
	}

	if _, ok := snap.Signers[signer]; !ok {
		return errUnauthorizedValidator
	}

	/*
	for seen, recent := range snap.Recents {
		if recent == signer {
			// Signer is among recents, only fail if the current block doesn't shift it out
			if limit := uint64(len(snap.Signers)/2 + 1); seen > number-limit {
				return errRecentlySigned
			}
		}
	}*/

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	if !s.fakeDiff {
		inturn := snap.inturn(header.Number.Uint64(),signer)
		if inturn && header.Difficulty.Cmp(diffInTurn) != 0 {
			return errWrongDifficulty
		}
		if !inturn && header.Difficulty.Cmp(diffNoTurn) != 0 {
			return errWrongDifficulty
		}
	}

	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (s *Spos) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	// If the block isn't a checkpoint, cast a random vote (good enough for now)
	header.Coinbase = s.signer
	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()
	// Assemble the voting snapshot to check which votes make sense
	snap, err := s.snapshot(chain, number - 1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	// Set the correct difficulty
	header.Difficulty = calcDifficulty(snap, s.signer)

	// Ensure the extra data has all its components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	if number % s.config.Epoch == 0 {
		for _, signer := range snap.signers() {
			header.Extra = append(header.Extra, signer[:]...)
		}
	}
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Time = parent.Time + s.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	return nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set
func (s *Spos) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	accumulateRewards(state, header)

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set, and returns the final block.
func (s *Spos) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Finalize block
	s.Finalize(chain, header, state, txs, uncles)
	number := header.Number.Uint64()

	completeBlockFlag := GetCompleteBlockFlag()

	//clearExpiredBlockRewardData(number)

	//Whether block rewards have been allocated
	//distributeRewardFlag := getDistributeRewardFlag(number)

	//if number >= params.SafeSposOfficialSuperNodeConfig.StartCommonSuperHeight && !distributeRewardFlag && completeBlockFlag {
	if number >= params.SafeSposOfficialSuperNodeConfig.StartCommonSuperHeight && completeBlockFlag {
		cx := chainContext{Chain: chain, Spos: s}
		_, err := s.distributeReward(header, cx, state, &txs, &receipts, &header.GasUsed)
		if err != nil {
			return nil, err
		}
		//setDistributeRewardFlag(number, true)
	}

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil)), nil
}


//Whether the current block height has been awarded
var distributeRewardLock   sync.RWMutex
var distributeRewardMap    map[uint64] bool

func getDistributeRewardFlag(number uint64) bool{
	distributeRewardLock.Lock()
	defer distributeRewardLock.Unlock()
	_, ok := distributeRewardMap[number]
	return ok
}

func setDistributeRewardFlag(number uint64, flag bool) {
	distributeRewardLock.Lock()
	defer distributeRewardLock.Unlock()
	if distributeRewardMap == nil {
		distributeRewardMap = make(map[uint64] bool, 10)
	}
	distributeRewardMap[number] = flag
}

func clearExpiredBlockRewardData(number uint64) {
	distributeRewardLock.Lock()
	defer distributeRewardLock.Unlock()
	for k,_ := range distributeRewardMap{
		if k < number {
			delete(distributeRewardMap, k)
		}
	}
}

func (s *Spos) DistributeIncoming(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) error {
	number := header.Number.Uint64()
	clearExpiredBlockRewardData(number)
	distributeRewardFlag := getDistributeRewardFlag(number)
	if number >= params.SafeSposOfficialSuperNodeConfig.StartCommonSuperHeight && !distributeRewardFlag {
		cx := chainContext{Chain: chain, Spos: s}
		_, err := s.distributeReward(header, cx, state, &txs, &receipts, &header.GasUsed)
		if err != nil {
			return err
		}
		setDistributeRewardFlag(number, true)
	}
	return nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
//func (s *Spos) Authorize(signer common.Address, signFn SignerFn, ebpk *ecdsa.PrivateKey) {
func (s *Spos) Authorize(signer common.Address, signFn SignerFn, signTxFn SignerTxFn) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.signer = signer
	s.signFn = signFn
	s.signTxFn = signTxFn
	//s.etherbaseprivatekey = ebpk
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (s *Spos) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	if s.config.Period == 0 {
		log.Info("Sealing paused")
		return nil
	}

	// Don't hold the signer fields for the entire sealing procedure
	s.lock.RLock()
	signer, signFn := s.signer, s.signFn
	s.lock.RUnlock()

	// Bail out if we're unauthorized to sign a block
	snap, err := s.snapshot(chain, number - 1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	if _, authorized := snap.Signers[signer]; !authorized {
		return errUnauthorizedSigner
	}

	/*
	// If we're amongst the recent signers, wait for the next block
	for seen, recent := range snap.Recents {
		if recent == signer {
			// Signer is among recents, only wait if the current block doesn't shift it out
			if limit := uint64(len(snap.Signers)/2 + 1); number < limit || seen > number-limit {
				log.Info("Signed recently, must wait for others")
				return nil
			}
		}
	}*/

	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Until(time.Unix(int64(header.Time), 0)) // nolint: gosimple
	if header.Difficulty.Cmp(diffNoTurn) == 0 {
		// It's not our turn explicitly to sign, delay it a bit
		wiggle := time.Duration(len(snap.Signers)/2+1) * wiggleTime
		delay += time.Duration(rand.Int63n(int64(wiggle)))

		log.Info("Sealing block with", "number", number, "delay", delay, "headerDifficulty", header.Difficulty, "signer", signer.Hex())
	}

	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeSpos, SposRLP(header))
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	return nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have:
// * DIFF_NOTURN(2) if BLOCK_NUMBER % SIGNER_COUNT != SIGNER_INDEX
// * DIFF_INTURN(1) if BLOCK_NUMBER % SIGNER_COUNT == SIGNER_INDEX
func (s *Spos) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	snap, err := s.snapshot(chain, parent.Number.Uint64(), parent.Hash(), nil)
	if err != nil {
		return nil
	}
	return calcDifficulty(snap,  s.signer)
}

func (s *Spos) calcDifficulty(snap *Snapshot, signer common.Address)*big.Int {
	if snap.inturn(snap.Number + 1,signer) {
		return new(big.Int).Set(diffInTurn)
	}
	return new(big.Int).Set(diffNoTurn)
}

// SealHash returns the hash of a block prior to it being sealed.
func (s *Spos) SealHash(header *types.Header) common.Hash {
	return SealHash(header)
}

/*
func (s * Spos) EhterbasePrivatekey() (ehterbasePrivatekey *ecdsa.PrivateKey, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	ebpk := s.etherbaseprivatekey
	if ebpk != nil {
		return ebpk, nil
	}
	return nil, fmt.Errorf("etherbase privatekey missing: %v", err)
}*/

// Close implements consensus.Engine. It's a noop for spos as there are no background threads.
func (s *Spos) Close() error {
	return nil
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (s *Spos) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return []rpc.API{{
		Namespace: "spos",
		Version:   "1.0",
		Service:   &API{chain: chain, spos: s},
		Public:    false,
	}}
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.(crypto.KeccakState).Read(hash[:])
	return hash
}

// SposRLP returns the rlp bytes which needs to be signed for the proof-of-stack
// sealing. The RLP to sign consists of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func SposRLP(header *types.Header) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header)
	return b.Bytes()
}

func encodeSigHeader(w io.Writer, header *types.Header) {
	enc := []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-crypto.SignatureLength], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	}
	if header.BaseFee != nil {
		enc = append(enc, header.BaseFee)
	}
	if err := rlp.Encode(w, enc); err != nil {
		panic("can't encode: " + err.Error())
	}
}

func accumulateRewards(state *state.StateDB, header *types.Header) {
	// Accumulate the rewards for the miner
	number := header.Number.Uint64()
	totalReward := getBlockSubsidy(number, false)
	state.AddBalance(header.Coinbase, totalReward)
}

func sortKey (mp map[string]common.Address) map[string]common.Address{
	var newMap = make([]string, 0)
	for k,_ := range mp {
		newMap = append(newMap, k)
	}

	sort.Strings(newMap)

	newScoreMasternode := make(map[string]common.Address,len(mp))
	for _, v:= range newMap {
		newScoreMasternode[v] = mp[v]
	}

	return newScoreMasternode
}

func sortSupernode(Signers map[common.Address]struct{}, header *types.Header, scoreTime uint64) []common.Address {
	scoreSupernode := make(map[string]common.Address,len(Signers))
	afterscoreSupernode := make(map[string]common.Address,len(Signers))

	for signer,_ := range Signers {
		hasher := sha3.NewLegacyKeccak256()
		enc := []interface{}{
			signer.Hash(),
			scoreTime,
		}

		if err := rlp.Encode(hasher, enc); err != nil {
			panic("can't encode: " + err.Error())
		}

		hash := common.Hash{}
		hasher.(crypto.KeccakState).Read(hash[:])
		scoreSupernode[hash.String()] = signer
	}

	afterscoreSupernode = sortKey(scoreSupernode)

	resultSuperMasterNode := make([]common.Address, 0)
	for _,address := range afterscoreSupernode {
		resultSuperMasterNode = append(resultSuperMasterNode, address)
	}
	now_hi := uint64(header.Time) << 32
	for i := 0; i < len(resultSuperMasterNode); i++ {
		k := now_hi + uint64(i) * 2685821657736338717
		k ^= (k >> 12)
		k ^= (k << 25)
		k ^= (k >> 27)
		k *= 2685821657736338717

		jmax := len(resultSuperMasterNode) - i
		j := uint64(i) + k % uint64(jmax)
		resultSuperMasterNode[i], resultSuperMasterNode[j] = resultSuperMasterNode[j],resultSuperMasterNode[i]
	}

	return resultSuperMasterNode
}

func getBlockSubsidy(nBlockNum uint64, fSuperblockPartOnly bool) *big.Int {
	subsidy := BlockReward.Uint64()

	// yearly decline of production by ~7.1% per year, projected ~18M coins max by year 2050+.
	for  i := nextDecrementHeight.Uint64(); i <= nBlockNum; i += subsidyHalvingInterval.Uint64(){
		subsidy -= subsidy / 14
	}

    superblockPart := subsidy / 10

    if fSuperblockPartOnly {
    	return new(big.Int).SetUint64(superblockPart)
	}else{
		return new(big.Int).SetUint64(subsidy - superblockPart)
	}
}

func getMasternodePayment(blockReward *big.Int) *big.Int {
	//start at 20%
	masternodePayment := blockReward.Uint64() / 5

	//The SAFE 3 height is greater than 935600, and the revenue of the master node is only about 50%
	masternodePayment += blockReward.Uint64() / 20
	masternodePayment += blockReward.Uint64() / 20
	masternodePayment += blockReward.Uint64() / 20
	masternodePayment += blockReward.Uint64() / 40
	masternodePayment += blockReward.Uint64() / 40
	masternodePayment += blockReward.Uint64() / 40
	masternodePayment += blockReward.Uint64() / 40
	masternodePayment += blockReward.Uint64() / 40
	masternodePayment += blockReward.Uint64() / 40

	return new(big.Int).SetUint64(masternodePayment)
}

func (s *Spos) distributeReward(header *types.Header, cx core.ChainContext, state *state.StateDB, txs *[]*types.Transaction, receipts *[]*types.Receipt, usedGas *uint64) (*types.Transaction, error) {
	number := header.Number.Uint64()
	totalReward := getBlockSubsidy(number, false)
	masterNodePayment := getMasternodePayment(totalReward)
	superNodeReward := new(big.Int).Sub(totalReward, masterNodePayment)
	mnAddr, err := systemcontracts.GetNextMasterNode(s.blockChainAPI)
	if err != nil {
		return nil, err
	}
	return s.Reward(header.Coinbase, superNodeReward, *mnAddr, masterNodePayment, header, cx, state, txs, receipts, usedGas)
}

func (s *Spos) Reward(smnAddr common.Address, smnCount *big.Int, mnAddr common.Address, mnCount *big.Int, header *types.Header, cx core.ChainContext, state *state.StateDB, txs *[]*types.Transaction, receipts *[]*types.Receipt, usedGas *uint64) (*types.Transaction, error){
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SystemRewardABI))
	if err != nil {
		return nil, err
	}

	method := "reward"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := vABI.Pack(method, smnAddr, smnCount, mnAddr, mnCount)
	if err != nil {
		return nil, err
	}

	value := new(big.Int)
	value.Add(smnCount, mnCount)
	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = systemcontracts.GetPropertyValue(s.blockChainAPI, "gas_price")
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}
	args := ethapi.TransactionArgs{
		From:     &smnAddr,
		To:       &systemcontracts.SystemRewardContractAddr,
		Data:     &msgData,
		Value:    (*hexutil.Big)(value),
		GasPrice: (*hexutil.Big)(gasPrice),
	}
	//blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(header.Number.Uint64() - 1))
	gas, err := s.blockChainAPI.EstimateGas(ctx, args, nil)
	args.Gas = &gas

	nonce := state.GetNonce(*args.From)

	rawTx :=types.NewTransaction(nonce, *args.To, value, uint64(*args.Gas), args.GasPrice.ToInt(), msgData)
	tx, err := s.signTxFn(accounts.Account{Address: smnAddr}, rawTx, s.chainConfig.ChainID)

	if err != nil {
		return nil, err
	}
	state.Prepare(tx.Hash(), len(*txs))
	gasUsed, err := applyMessage(args, header, cx, state, s.chainConfig)
	if err != nil {
		return nil, err
	}
	*txs = append(*txs, tx)
	root := state.IntermediateRoot(s.chainConfig.IsEIP158(header.Number)).Bytes()
	*usedGas += gasUsed
	receipt := types.NewReceipt(root, false, *usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = gasUsed

	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = state.GetLogs(tx.Hash(), header.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	receipt.BlockHash = header.Hash()
	receipt.BlockNumber = header.Number
	receipt.TransactionIndex = uint(state.TxIndex())
	*receipts = append(*receipts, receipt)
	state.SetNonce(*args.From, nonce+1)

//	TxPool.AddLocal(tx)
	SetReceipts(*receipts)
	SetTxs(*txs)

	return tx, err
}

// chain context
type chainContext struct {
	Chain  consensus.ChainHeaderReader
	Spos consensus.Engine
}

func (c chainContext) Engine() consensus.Engine {
	return c.Spos
}

func (c chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return c.Chain.GetHeader(hash, number)
}

func applyMessage(
	args ethapi.TransactionArgs,
	header *types.Header,
	cx core.ChainContext,
	state *state.StateDB,
	chainConfig *params.ChainConfig,
) (uint64, error) {
	// Create a new context to be used in the EVM environment
	context := core.NewEVMBlockContext(header, cx, nil)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, vm.TxContext{Origin: *args.From, GasPrice: big.NewInt(0)}, state, chainConfig, vm.Config{})
	// Apply the transaction to the current state (included in the env)
	ret, returnGas, err := vmenv.Call(vm.AccountRef(*args.From), *args.To, *args.Data, uint64(*args.Gas), args.Value.ToInt())
	if err != nil {
		log.Error("apply message failed", "msg", string(ret), "err", err)
		return 0, err
	}
	return uint64(*args.Gas) - returnGas, err
}
