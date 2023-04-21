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
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
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
	checkpointInterval = 1024     // Number of blocks after which to save the snapshot to the database
	validatorBytesLength = common.AddressLength

	wiggleTime = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers

	//superNodeSPosCount = 7           //Total number of bookkeepers
	superNodeSPosCount = 2
	pushForwardHeight  = 14	          //Push forward the block height
	//chtAddress         = "0x043807066705c6EF9EB3D28D5D230b4d87EC4832" //Contract address

	// genesis contracts
	AccountManagerContract ="0x0000000000000000000000000000000000001012"
	MasterNodeContract = "0x0000000000000000000000000000000000001022"
	NodeStateContract = "0x0000000000000000000000000000000000001052"
	SuperNodeStateContract = "0x0000000000000000000000000000000000001062"
	PropertyContract = "0x0000000000000000000000000000000000001002"
	ProposalContract = "0x0000000000000000000000000000000000001072"
	SMNVoteContract = "0x0000000000000000000000000000000000001042"
	SuperMasterNodeContract = "0x0000000000000000000000000000000000001032"
	SystemRewardContract ="0x0000000000000000000000000000000000001082"
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

var SposLock   sync.RWMutex
var Signerlist []common.Address
var StartNewLoopTime uint64
var PushForwardTime uint64
//var SposTxLock    sync.RWMutex
//var MinerRewardTx *types.Transaction
var ReceiptsLock  sync.RWMutex
var Receipts []*types.Receipt

//Whether the current block height has been awarded
var distributeRewardLock   sync.RWMutex
var distributeRewardMap    map[uint64] bool



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

	//errIndexError is returned if Index error
	errIndexError = errors.New("Index error")

	//errAllowTime The maximum allowed time is exceeded
	errAllowTime = errors.New("invalid time")

	// errCoinBaseMisMatch is returned if a header's coinbase do not match with signature
	errCoinBaseMisMatch = errors.New("coinbase do not match with signature")

	//errBillinglist The billing list is empty
	errBillinglist = errors.New("invalid Billing list")

	//errSignatory The signatory is incorrect
	errSignatory = errors.New("invalid signatory")

	//errKeyStore The KeyStore is empty
	errKeyStore = errors.New("invalid KeyStore")
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

	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

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
	signatures, _ := lru.NewARC(inmemorySignatures)

	return &Spos{
		chainConfig: config,
		config:     conf,
		db:         db,
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
	if header.Time > uint64(time.Now().Unix()){
		return consensus.ErrFutureBlock
	}

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
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
		if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
			return errInvalidDifficulty
		}
	}
	// Verify that the gas limit is <= 2^63-1
	if header.GasLimit > params.MaxGasLimit {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, params.MaxGasLimit)
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
	if header.Time > uint64(time.Now().Unix()) {
		return errInvalidTimestamp
	}

	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}

	if !chain.Config().IsLondon(header.Number) {
		// Verify BaseFee not present before EIP-1559 fork.
		if header.BaseFee != nil {
			return fmt.Errorf("invalid baseFee before fork: have %d, want <nil>", header.BaseFee)
		}
		if err := misc.VerifyGaslimit(parent.GasLimit, header.GasLimit); err != nil {
			return err
		}
	} else if err := misc.VerifyEip1559Header(chain.Config(), parent, header); err != nil {
		// Verify the header's EIP-1559 attributes.
		return err
	}

	// All basic checks passed, verify the seal and return
	return s.verifySeal(header, parents)
}

// snapshot retrieves the authorization snapshot at a given point in time.
func (s *Spos) snapshot(chain consensus.ChainHeaderReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot
	)
	for snap == nil {
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

				var signers   []common.Address

				if number == 0 {
					signers = params.SafeSposOfficialSuperNodeConfig.Signers
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
	// Previous snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}

	snap, err := snap.apply(headers)
	if err != nil {
		return nil, err
	}

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
func (s *Spos) verifySeal( header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, s.signatures)
	if err != nil {
		return err
	}

	if signer != header.Coinbase {
		return errCoinBaseMisMatch
	}

	if header.Time > uint64(time.Now().Unix()) {
		return errAllowTime
	}

	SposLock.RLock()
	defer SposLock.RUnlock()
	nmnSize := len(Signerlist)
	if nmnSize == 0 {
		return errBillinglist
	}

	ninterval := (header.Time - StartNewLoopTime - PushForwardTime) / s.config.Period - 1
	nindex := ninterval % uint64(nmnSize)
	if nindex < 0 {
		return errIndexError
	}

	blocksigner := Signerlist[nindex]
	if blocksigner != signer {
		return errSignatory
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
	s.lock.RLock()

	//Select 7 bookkeepers
	if number - 1 == 0 || number % superNodeSPosCount == 0 {
		if number > pushForwardHeight {
			forwardHeight :=  number - pushForwardHeight
			forwardblock := chain.GetHeaderByNumber(forwardHeight)

			SposLock.Lock()
			PushForwardTime = pushForwardHeight * s.config.Period
			StartNewLoopTime = forwardblock.Time
			SposLock.Unlock()
		}else{
			parentblock := chain.GetHeaderByHash(header.ParentHash)

			SposLock.Lock()
			PushForwardTime = 0
			StartNewLoopTime = parentblock.Time
			SposLock.Unlock()
		}

		if number < params.SafeSposOfficialSuperNodeConfig.StartCommonSuperHeight {
			snap.Signers = make(map[common.Address]struct{})

			for _, signer := range params.SafeSposOfficialSuperNodeConfig.Signers {
				snap.Signers[signer] = struct{}{}
			}
		}else { //TODO Call the contract to get the super node list
			superMasterNodeInfos, err := s.getTopSuperMasterNode(number - 1)
			if err != nil {
				log.Error("Failed to GetTopSMN", "error", err)
				return err
			}

			snap.Signers = make(map[common.Address]struct{})
			for i := range superMasterNodeInfos {
				log.Info("Super MasterNode Addr Info", "superMasterNodeInfos[i].Addr", superMasterNodeInfos[i].Addr)
				snap.Signers[superMasterNodeInfos[i].Addr] = struct{}{}
			}
		}

		resultSuperNode := []common.Address{}
		resultSuperNode = sortSupernode(snap, header)

		SposLock.Lock()
		if len(Signerlist) > 0 {
			Signerlist = []common.Address{}
		}
		SposLock.Unlock()

		for i := 0; i < superNodeSPosCount; i++{
			SposLock.Lock()
			Signerlist = append(Signerlist, resultSuperNode[i])
			SposLock.Unlock()
		}
	}

	// Copy signer protected by mutex to avoid race condition
	signer := s.signer
	s.lock.RUnlock()

	// Set the correct difficulty
	header.Difficulty = calcSposDifficulty(snap, signer)

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

	nCurTime := uint64(time.Now().Unix())
	var nTimeInterval uint64 = 0
	nSposTargetSpacing := s.config.Period

	SposLock.RLock()
	defer SposLock.RUnlock()
	nTimeInterval = nCurTime - PushForwardTime - StartNewLoopTime
	if nTimeInterval < 0 {
		return errIndexError
	}

	nTargetSpacingInterval := nTimeInterval / nSposTargetSpacing
	header.Time = StartNewLoopTime + PushForwardTime + (nTargetSpacingInterval + 1) * nSposTargetSpacing

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
	clearExpiredBlockRewardData(number)

	//Whether block rewards have been allocated
	distributeRewardFlag := getDistributeRewardFlag(number)

	var rewardTx *types.Transaction
	if number >= params.SafeSposOfficialSuperNodeConfig.StartCommonSuperHeight && !distributeRewardFlag {
		distributeRewardTx, err := s.distributeReward(header, state)
		if err != nil {
			return nil, err
		}
		rewardTx = distributeRewardTx
		setDistributeRewardFlag(number, true)
	}

	//The reward distribution transaction in mining is the first transaction in this block
	//rewardTx := getMinerRewardTx()
	if rewardTx != nil {
		/*tempTransaction := new(types.Transaction)
		txs = append(txs, tempTransaction)
		copy(txs[1:], txs[0:])
		txs[0] = rewardTx
		 */

		state.Prepare(rewardTx.Hash(), len(txs))
		txs = append(txs, rewardTx)

		usedGas := rewardTx.Gas()
		receipt := types.NewReceipt(header.Root.Bytes(), false, usedGas)
		receipt.TxHash = rewardTx.Hash()
		receipt.GasUsed = usedGas

		header.GasUsed += usedGas

		// Set the receipt logs and create a bloom for filtering
		nonce := state.GetNonce(header.Coinbase)
		receipt.Logs = state.GetLogs(rewardTx.Hash(), header.Hash())
		receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
		receipt.BlockHash = header.Hash()
		receipt.BlockNumber = header.Number
		receipt.TransactionIndex = uint(state.TxIndex())
		state.SetNonce(header.Coinbase, nonce + 1)

		receipts = append(receipts, receipt)
		/*tempreceipts := new(types.Receipt)
		receipts = append(receipts, tempreceipts)
		copy(receipts[1:], receipts[0:])
		receipts[0] = receipt
		 */
	}

	SetReceipts(receipts)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil)), nil
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

	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	if header.Difficulty.Cmp(diffNoTurn) == 0 {
		// It's not our turn explicitly to sign, delay it a bit
		wiggle := time.Duration(len(snap.Signers)/2+1) * wiggleTime
		delay += time.Duration(rand.Int63n(int64(wiggle)))

		log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
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
	s.lock.RLock()
	signer := s.signer
	s.lock.RUnlock()
	return calcSposDifficulty(snap,  signer)
}

func calcSposDifficulty(snap *Snapshot, signer common.Address)*big.Int {
	if snap.inturnblock(signer) {
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

func sortSupernode(snap *Snapshot, header *types.Header) []common.Address {
	scoreSupernode := make(map[string]common.Address,len(snap.Signers))
	afterscoreSupernode := make(map[string]common.Address,len(snap.Signers))

	SposLock.RLock()
	tempPushForwardBlockTime := StartNewLoopTime
	SposLock.RUnlock()

	for signer,_ := range snap.Signers {
		hasher := sha3.NewLegacyKeccak256()
		enc := []interface{}{
			signer.Hash(),
			tempPushForwardBlockTime,
		}

		if err := rlp.Encode(hasher, enc); err != nil {
			panic("can't encode: " + err.Error())
		}

		hash := common.Hash{}
		hasher.(crypto.KeccakState).Read(hash[:])
		scoreSupernode[hash.String()] = signer
	}

	afterscoreSupernode = sortKey(scoreSupernode)

	resultSuperMasterNode := []common.Address{}
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

/*
func setMinerRewardTx(tx *types.Transaction) {
	SposTxLock.Lock()
	defer SposTxLock.Unlock()
	if MinerRewardTx != nil {
		MinerRewardTx = nil
	}
	MinerRewardTx = tx
}

func getMinerRewardTx() *types.Transaction{
	var minerRewardTx *types.Transaction
	SposTxLock.Lock()
	minerRewardTx = MinerRewardTx
	SposTxLock.Unlock()
	return minerRewardTx
}*/

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

func (s *Spos) distributeReward(header *types.Header, state *state.StateDB) (*types.Transaction, error) {
	number := header.Number.Uint64()
	totalReward := getBlockSubsidy(number, false)
	masterNodePayment := getMasternodePayment(totalReward)
	superNodeReward := new(big.Int).Sub(totalReward, masterNodePayment)
	log.Info("Block reward info", "reward", totalReward,"superNodeReward", superNodeReward,
		     "masterNodePayment", masterNodePayment)

	mnAddr, err := s.getNextMasterNode(number - 1)
	if err != nil {
		return nil, err
	}
	return s.Reward(number - 1, state, header.Coinbase, superNodeReward, *mnAddr, masterNodePayment)
}
func (s *Spos) getNextMasterNode(height uint64) (*common.Address, error) {
	log.Info("lemengbin", "height", height, "latest", rpc.LatestBlockNumber)
	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(height))
	method := "getNext"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeABI))
	if err != nil {
		return nil, nil
	}

	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}
	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.MasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := s.blockChainAPI.Call(ctx, args, blockNrOrHash, nil)
	if err != nil {
		return nil, err
	}

	addr := new(common.Address)
	if err := vABI.UnpackIntoInterface(&addr, method, result); err != nil {
		return nil, err
	}
	return addr, nil
}

func (s *Spos) getTopSuperMasterNode(height uint64) ([]types.SuperMasterNodeInfo, error) {
	log.Info("lemengbin", "height", height, "latest", rpc.LatestBlockNumber)
	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(height))
	method := "getTop"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperMasterNodeABI))
	if err != nil {
		return nil, err
	}

	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}
	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperMasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := s.blockChainAPI.Call(ctx, args, blockNrOrHash, nil)
	if err != nil {
		return nil, err
	}

	var (
		ret0 = new([]types.SuperMasterNodeInfo)
	)
	out := ret0
	if err := vABI.UnpackIntoInterface(out, method, result); err != nil {
		return nil, err
	}

	addrs :=make([]types.SuperMasterNodeInfo, len(*ret0))
	for i, addr := range *ret0 {
		addrs[i] = addr
	}
	return addrs, nil
}

func (s *Spos) Reward(height uint64, state *state.StateDB, smnAddr common.Address, smnCount *big.Int, mnAddr common.Address, mnCount *big.Int) (*types.Transaction, error){
	log.Info("lemengbin", "height", height, "latest", rpc.LatestBlockNumber)
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
	args := ethapi.TransactionArgs{
		From:  &smnAddr,
		To:    &systemcontracts.SystemRewardContractAddr,
		Data:  &msgData,
		Value: (*hexutil.Big)(value),
	}
	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(height))
	gas, err := s.blockChainAPI.EstimateGas(ctx, args, &blockNrOrHash)
	args.Gas = &gas
	nonce := state.GetNonce(smnAddr)

	rawTx :=types.NewTransaction(nonce, *args.To, value, uint64(*args.Gas), args.GasPrice.ToInt(), msgData)
	tx, err := s.signTxFn(accounts.Account{Address: smnAddr}, rawTx, s.chainConfig.ChainID)
	if err != nil {
		return nil, err
	}
	return tx, err
}

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
