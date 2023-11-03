package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type PropertyInfo struct {
	Name         string             `json:"name"                    gencodec:"required"`
	Value        *big.Int           `json:"value"                   gencodec:"required"`
	Description  string             `json:"description"             gencodec:"required"`
	CreateHeight *big.Int           `json:"createHeight"            gencodec:"required"`
	UpdateHeight *big.Int           `json:"updateHeight"            gencodec:"required"`
}

type UnconfirmedPropertyInfo struct {
	Name        string              `json:"name"                    gencodec:"required"`
	Value       *big.Int            `json:"value"                   gencodec:"required"`
	Applicant   common.Address      `json:"applicant"               gencodec:"required"`
	Voters      []common.Address    `json:"voters"                  gencodec:"required"`
	VoteResults []*big.Int          `json:"voteResults"             gencodec:"required"`
	Reason      string              `json:"reason"                  gencodec:"required"`
	ApplyHeight *big.Int            `json:"applyHeight"             gencodec:"required"`
}