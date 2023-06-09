package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"math/big"
	"sync/atomic"
	"time"
)

type MasterNodeMemberInfo struct {
	LockID *big.Int         `json:"lockID"    gencodec:"required"`
	Addr   common.Address   `json:"addr"      gencodec:"required"`
	Amount *big.Int         `json:"amount"    gencodec:"required"`
	Height *big.Int         `json:"height"    gencodec:"required"`
}

type MasterNodeIncentivePlan struct {
	Creator *big.Int        `json:"creator"   gencodec:"required"`
	Partner *big.Int        `json:"partner"   gencodec:"required"`
	Voter   *big.Int        `json:"voter"     gencodec:"required"`
}

type MasterNodeInfo struct {
	Id            *big.Int                  `json:"id"            gencodec:"required"`
	Addr          common.Address            `json:"addr"          gencodec:"required"`
	Creator       common.Address            `json:"creator"       gencodec:"required"`
	Amount        *big.Int                  `json:"amount"        gencodec:"required"`
	Enode         string                    `json:"enode"         gencodec:"required"`
	Ip            string                    `json:"ip"            gencodec:"required"`
	Description   string                    `json:"description,omitempty"`
	State               *big.Int            `json:"state"         gencodec:"required"`
	Founders      []MasterNodeMemberInfo    `json:"founders"      gencodec:"required"`
	IncentivePlan MasterNodeIncentivePlan   `json:"incentivePlan" gencodec:"required"`
	CreateHeight  *big.Int                  `json:"createHeight,omitempty"`
	UpdateHeight  *big.Int                  `json:"updateHeight,omitempty"`
}

// MasterNodePing is an masternode ping.
type MasterNodePing struct {
	version *big.Int        `json:"version"        gencodec:"required"`
	signTime time.Time      `json:"signTime"       gencodec:"required"`
	sign []byte             `json:"sign"           gencodec:"required"`
	blockHash common.Hash   `json:"blockHash"      gencodec:"required"`

	// caches
	hash atomic.Value
	size atomic.Value
}

const MnpVersion = 1001

func NewMasterNodePing(masterNodeInfo *MasterNodeInfo, blockHash common.Hash) *MasterNodePing {
	mnp := &MasterNodePing{}
	mnp.version = big.NewInt(MnpVersion)
	mnp.signTime = time.Now()
	mnp.sign = nil
	mnp.blockHash = blockHash
	return mnp
}

// Hash returns the transaction hash.
func (mnp *MasterNodePing) Hash() common.Hash {
	if hash := mnp.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	h := rlpHash(mnp)
	mnp.hash.Store(h)
	return h
}

func (mnp *MasterNodePing) Size() common.StorageSize {
	if size := mnp.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, mnp)
	mnp.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

type extMasterNodePing struct {
	Version *big.Int
	SignTime time.Time
	Sign []byte
	BlockHash common.Hash
}

// DecodeRLP decodes the Ethereum
func (mnp *MasterNodePing) DecodeRLP(s *rlp.Stream) error {
	var emnp extMasterNodePing
	_, size, _ := s.Kind()
	if err := s.Decode(&emnp); err != nil {
		return err
	}
	mnp.version, mnp.signTime, mnp.sign, mnp.blockHash = emnp.Version, emnp.SignTime, emnp.Sign, emnp.BlockHash
	mnp.size.Store(common.StorageSize(rlp.ListSize(size)))
	return nil
}

// EncodeRLP serializes b into the Ethereum RLP block format.
func (mnp *MasterNodePing) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, extMasterNodePing{
		Version: mnp.version,
		SignTime: mnp.signTime,
		Sign: mnp.sign,
		BlockHash: mnp.blockHash,
	})
}
