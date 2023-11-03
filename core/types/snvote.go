package types

import (
    "github.com/ethereum/go-ethereum/common"
    "math/big"
)

type SuperNodes4VoterInfo struct {
    SuperNodes  []common.Address    `json:"superNodes"      gencodec:"required"`
    VoteNums    []big.Int           `json:"voteNums"        gencodec:"required"`
}

type Voters4SNInfo struct {
    Voters      []common.Address    `json:"voters"          gencodec:"required"`
    VoteNums    []big.Int           `json:"voteNums"        gencodec:"required"`
}

type Proxies4VoterInfo struct {
    Proxies     []common.Address    `json:"proxies"         gencodec:"required"`
    VoteNums    []big.Int           `json:"voteNums"        gencodec:"required"`
}

type Voters4ProxyInfo struct {
    Voters      []common.Address    `json:"voters"          gencodec:"required"`
    VoteNums    []big.Int           `json:"voteNums"        gencodec:"required"`
}