package types

import (
    "github.com/ethereum/go-ethereum/common"
    "math/big"
)

type SNVoteRetInfo struct {
    Addrs       []common.Address    `json:"addrs"           gencodec:"required"`
    VoteNums    []*big.Int          `json:"voteNums"        gencodec:"required"`
}

type VoteRecord struct {
    VoterAddr   common.Address      `json:"voterAddr"       gencodec:"required"`
    DstAddr     common.Address      `json:"dstAddr"         gencodec:"required"`
    Amount      *big.Int            `json:"amount"          gencodec:"required"`
    Num         *big.Int            `json:"num"             gencodec:"required"`
    Height      *big.Int            `json:"height"          gencodec:"required"`
}