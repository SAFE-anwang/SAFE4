package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync/atomic"
	"time"
)

type SuperMasterNodeMemberInfo struct {
	LockID [20]byte
	Addr   common.Address
	Amount *big.Int
	Height *big.Int
}

type SuperMasterNodeIncentivePlan struct {
	Creator *big.Int
	Partner *big.Int
	Voter   *big.Int
}

type SuperMasterNodeInfo struct {
	Id               *big.Int
	Name             string
	Addr             common.Address
	Creator          common.Address
	Amount           *big.Int
	Ip               string
	Pubkey           string
	Description      string
	State            *big.Int
	Founders         []SuperMasterNodeMemberInfo
	IncentivePlan    SuperMasterNodeIncentivePlan
	Voters           []SuperMasterNodeMemberInfo
	TotalVoteNum     *big.Int
	TotalVoterAmount *big.Int
	CreateHeight     *big.Int
	UpdateHeight     *big.Int
}

// SuperMasterNodePing is a supermasternode ping.
type SuperMasterNodePing struct {
	version int             `json:"version"        gencodec:"required"`
	signTime time.Time      `json:"signTime"       gencodec:"required"`
	sign []byte             `json:"sign"           gencodec:"required"`
	blockhash common.Hash   `json:"blockhash"      gencodec:"required"`

	// caches
	hash atomic.Value
	size atomic.Value
}

const SmnpVersion = 1001

func NewSuperMasterNodePing(superMasterNodeInfo *SuperMasterNodeInfo) *SuperMasterNodePing {
	smnp := &SuperMasterNodePing{}
	smnp.version = SmnpVersion
	smnp.signTime = time.Now()
	smnp.sign = nil
	return smnp
}

// Hash returns the transaction hash.
func (smnp *SuperMasterNodePing) Hash() common.Hash {
	if hash := smnp.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	h := rlpHash(smnp)
	smnp.hash.Store(h)
	return h
}
