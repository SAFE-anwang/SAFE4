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

type SuperNodeStateInfo struct {
	State  	uint8           `json:"state"     gencodec:"required"`
	Height  *big.Int        `json:"height"    gencodec:"required"`
}

type SuperVoteInfo struct {
	Voters       []SuperNodeMemberInfo   `json:"voters"      gencodec:"required"`
	TotalAmount  *big.Int                `json:"totalAmount" gencodec:"required"`
	TotalNum     *big.Int                `json:"totalNum"    gencodec:"required"`
	Height       *big.Int                `json:"height"      gencodec:"required"`
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
	IsOfficial          bool                    `json:"isOfficial"    gencodec:"required"`
	StateInfo           SuperNodeStateInfo      `json:"stateInfo"     gencodec:"required"`
	Founders            []SuperNodeMemberInfo   `json:"founders"      gencodec:"required"`
	IncentivePlan       SuperNodeIncentivePlan  `json:"incentivePlan" gencodec:"required"`
	VoteInfo            SuperVoteInfo           `json:"voteInfo"      gencodec:"required"`
	LastRewardHeight  *big.Int                  `json:"lastRewardHeight" gencodec:"required"`
	CreateHeight  *big.Int                      `json:"createHeight"  gencodec:"required"`
	UpdateHeight  *big.Int                      `json:"updateHeight"  gencodec:"required"`
}

// SuperNodePing is a supernode ping.
type SuperNodePing struct {
	Version int             `json:"version"        gencodec:"required"`
	SignTime time.Time      `json:"signTime"       gencodec:"required"`
	Sign []byte             `json:"sign"           gencodec:"required"`
	BlockHash common.Hash   `json:"blockhash"      gencodec:"required"`

	// caches
	hash atomic.Value
	size atomic.Value
}

const SnpVersion = 1001

func NewSuperNodePing(superNodeInfo *SuperNodeInfo) *SuperNodePing {
	snp := &SuperNodePing{}
	snp.Version = SnpVersion
	snp.SignTime = time.Now()
	snp.Sign = nil
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
