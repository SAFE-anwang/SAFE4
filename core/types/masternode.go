package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync/atomic"
	"time"
)

type MasterNodeMemberInfo struct {
	LockID *big.Int
	Addr   common.Address
	Amount *big.Int
	Height *big.Int
}

type MasterNodeIncentivePlan struct {
	Creator *big.Int
	Partner *big.Int
	Voter   *big.Int
}

type MasterNodeInfo struct {
	Id            *big.Int
	Addr          common.Address
	Creator       common.Address
	Amount        *big.Int
	Enode         string
	Ip            string
	Pubkey        string
	Description   string
	State         *big.Int
	Founders      []MasterNodeMemberInfo
	IncentivePlan MasterNodeIncentivePlan
	CreateHeight  *big.Int
	UpdateHeight  *big.Int
}

// MasterNodePing is an masternode ping.
type MasterNodePing struct {
	version int             `json:"version"        gencodec:"required"`
	signTime time.Time      `json:"signTime"       gencodec:"required"`
	sign []byte             `json:"sign"           gencodec:"required"`
	blockhash common.Hash   `json:"blockhash"      gencodec:"required"`

	// caches
	hash atomic.Value
	size atomic.Value
}

const MnpVersion = 1001

func NewMasterNodePing(masterNodeInfo *MasterNodeInfo) *MasterNodePing {
	mnp := &MasterNodePing{}
	mnp.version = MnpVersion
	mnp.signTime = time.Now()
	mnp.sign = nil
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
