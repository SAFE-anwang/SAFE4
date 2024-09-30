package types

import (
    "github.com/ethereum/go-ethereum/common"
    "math/big"
)

type AvailableSafe3Info struct {
    Safe3Addr           string              `json:"safe3Addr"       gencodec:"required"`
    Amount              *big.Int            `json:"amount"          gencodec:"required"`
    Safe4Addr           common.Address      `json:"safe4Addr"       gencodec:"required"`
    RedeemHeight        *big.Int            `json:"redeemHeight"    gencodec:"required"`
}

type LockedSafe3Info struct {
    Safe3Addr           string              `json:"safe3Addr"       gencodec:"required"`
    Amount              *big.Int            `json:"amount"          gencodec:"required"`
    RemainLockHeight    *big.Int            `json:"remainLockHeight" gencodec:"required"`
    LockDay             *big.Int            `json:"lockDay"         gencodec:"required"`
    IsMN                bool                `json:"isMN"            gencodec:"required"`
    Safe4Addr           common.Address      `json:"safe4Addr"       gencodec:"required"`
    RedeemHeight        *big.Int            `json:"redeemHeight"    gencodec:"required"`
}

type SpecialSafe3Info struct {
    Safe3Addr           string              `json:"safe3Addr"       gencodec:"required"`
    Amount              *big.Int            `json:"amount"          gencodec:"required"`
    ApplyHeight         *big.Int            `json:"applyHeight"     gencodec:"required"`
    Voters              []common.Address    `json:"voters"          gencodec:"required"`
    VoteResults         []*big.Int          `json:"voteResults"     gencodec:"required"`
    Safe4Addr           common.Address      `json:"safe4Addr"       gencodec:"required"`
    RedeemHeight        *big.Int            `json:"redeemHeight"    gencodec:"required"`
}