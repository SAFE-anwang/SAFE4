package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ProposalInfo struct {
	Id           *big.Int           `json:"id"              gencodec:"required"`
	Creator      common.Address     `json:"creator"         gencodec:"required"`
	Title        string             `json:"title"           gencodec:"required"`
	PayAmount    *big.Int           `json:"payAmount"       gencodec:"required"`
	PayTimes     *big.Int           `json:"payTimes"        gencodec:"required"`
	StartPayTime *big.Int           `json:"startPayTime"    gencodec:"required"`
	EndPayTime   *big.Int           `json:"endPayTime"      gencodec:"required"`
	Description  string             `json:"description"     gencodec:"required"`
	Detail       string             `json:"detail"          gencodec:"required"`
	Voters       []common.Address   `json:"voters,omitempty"`
	VoteResults  []*big.Int         `json:"voteResults,omitempty"`
	State        *big.Int           `json:"state"           gencodec:"required"`
	CreateHeight *big.Int           `json:"createHeight,omitempty"`
	UpdateHeight *big.Int           `json:"UpdateHeight,omitempty"`
}