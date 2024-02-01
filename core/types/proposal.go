package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ProposalInfo struct {
	Id           *big.Int           `json:"id"                      gencodec:"required"`
	Creator      common.Address     `json:"creator"                 gencodec:"required"`
	Title        string             `json:"title"                   gencodec:"required"`
	PayAmount    *big.Int           `json:"payAmount"               gencodec:"required"`
	PayTimes     *big.Int           `json:"payTimes"                gencodec:"required"`
	StartPayTime *big.Int           `json:"startPayTime"            gencodec:"required"`
	EndPayTime   *big.Int           `json:"endPayTime"              gencodec:"required"`
	Description  string             `json:"description"             gencodec:"required"`
	State        *big.Int           `json:"state"                   gencodec:"required"`
	CreateHeight *big.Int           `json:"createHeight"            gencodec:"required"`
	UpdateHeight *big.Int           `json:"UpdateHeight"            gencodec:"required"`
}

type ProposalVoteInfo struct {
	Voter        common.Address     `json:"voter"                   gencodec:"required"`
	VoteResult   *big.Int           `json:"voteResult"              gencodec:"required"`
}