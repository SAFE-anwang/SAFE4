package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync/atomic"
	"time"
)

type SuperNodeMemberInfo struct {
	LockID *big.Int
	Addr   common.Address
	Amount *big.Int
	Height *big.Int
}

type SuperNodeIncentivePlan struct {
	Creator *big.Int
	Partner *big.Int
	Voter   *big.Int
}

type SuperNodeInfo struct {
	Id               *big.Int
	Name             string
	Addr             common.Address
	Creator          common.Address
	Amount           *big.Int
	Enode            string
	Ip               string
	Description      string
	State            *big.Int
	Founders         []SuperNodeMemberInfo
	IncentivePlan    SuperNodeIncentivePlan
	Voters           []SuperNodeMemberInfo
	TotalVoteNum     *big.Int
	TotalVoterAmount *big.Int
	CreateHeight     *big.Int
	UpdateHeight     *big.Int
}

// SuperNodePing is a supernode ping.
type SuperNodePing struct {
	version int             `json:"version"        gencodec:"required"`
	signTime time.Time      `json:"signTime"       gencodec:"required"`
	sign []byte             `json:"sign"           gencodec:"required"`
	blockhash common.Hash   `json:"blockhash"      gencodec:"required"`

	// caches
	hash atomic.Value
	size atomic.Value
}

const SnpVersion = 1001

func NewSuperNodePing(superNodeInfo *SuperNodeInfo) *SuperNodePing {
	snp := &SuperNodePing{}
	snp.version = SnpVersion
	snp.signTime = time.Now()
	snp.sign = nil
	return snp
}

// Hash returns the transaction hash.
func (snp *SuperNodePing) Hash() common.Hash {
	if hash := snp.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	h := rlpHash(snp)
	snp.hash.Store(h)
	return h
}
