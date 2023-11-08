package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type AccountRecord struct {
	Id           *big.Int           `json:"id"              gencodec:"required"`
	Addr         common.Address     `json:"addr"            gencodec:"required"`
	Amount       *big.Int           `json:"amount"          gencodec:"required"`
	LockDay      *big.Int           `json:"lockDay"         gencodec:"required"`
	StartHeight  *big.Int           `json:"startHeight"     gencodec:"required"`
	UnlockHeight *big.Int           `json:"unlockHeight"    gencodec:"required"`
}

type AccountRecordUseInfo struct {
	FrozenAddr      common.Address  `json:"frozenAddr"      gencodec:"required"`
	FreezeHeight    *big.Int        `json:"freezeHeight"    gencodec:"required"`
	UnfreezeHeight  *big.Int        `json:"unfreezeHeight"  gencodec:"required"`
	VotedAddr       common.Address  `json:"votedAddr"       gencodec:"required"`
	VoteHeight      *big.Int        `json:"voteHeight"      gencodec:"required"`
	ReleaseHeight   *big.Int        `json:"releaseHeight"   gencodec:"required"`
}

type AccountAmountInfo struct {
	Amount      *big.Int            `json:"amount"      gencodec:"required"`
	IDs         []*big.Int          `json:"ids"         gencodec:"required"`
}