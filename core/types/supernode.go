package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync/atomic"
	"time"
)

type SuperNodeMemberInfo struct {
	LockID *big.Int         `json:"lockID"    gencodec:"required"`
	Addr   common.Address   `json:"addr"      gencodec:"required"`
	Amount *big.Int         `json:"amount"    gencodec:"required"`
	Height *big.Int         `json:"height"    gencodec:"required"`
}

type SuperNodeIncentivePlan struct {
	Creator *big.Int        `json:"creator"   gencodec:"required"`
	Partner *big.Int        `json:"partner"   gencodec:"required"`
	Voter   *big.Int        `json:"voter"     gencodec:"required"`
}

type SuperNodeInfo struct {
	Id                  *big.Int                `json:"id"            gencodec:"required"`
	Name                string                  `json:"name"          gencodec:"required"`
	Addr                common.Address          `json:"addr"          gencodec:"required"`
	Creator             common.Address          `json:"creator"       gencodec:"required"`
	Amount              *big.Int                `json:"amount"        gencodec:"required"`
	Enode               string                  `json:"enode"         gencodec:"required"`
	Ip                  string                  `json:"ip"            gencodec:"required"`
	Description         string                  `json:"description"   gencodec:"required"`
	State               *big.Int                `json:"state"         gencodec:"required"`
	Founders            []SuperNodeMemberInfo   `json:"founders"      gencodec:"required"`
	IncentivePlan       SuperNodeIncentivePlan  `json:"incentivePlan" gencodec:"required"`
	Voters              []SuperNodeMemberInfo   `json:"voters,omitempty"`
	TotalVoteNum        *big.Int                `json:"totalVoteNum,omitempty"`
	TotalVoterAmount    *big.Int                `json:"totalVoterAmount,omitempty"`
	CreateHeight        *big.Int                `json:"createHeight,omitempty"`
	UpdateHeight        *big.Int                `json:"updateHeight,omitempty"`
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
