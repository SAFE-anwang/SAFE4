package types

import (
    "github.com/ethereum/go-ethereum/common"
    "math/big"
)

type SNVoteRetInfo struct {
    Addrs       []common.Address    `json:"addrs"           gencodec:"required"`
    VoteNums    []*big.Int           `json:"voteNums"        gencodec:"required"`
}