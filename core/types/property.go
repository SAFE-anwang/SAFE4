package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type PropertyInfo struct {
	Name         string         `json:"name"            gencodec:"required"`
	Value        *big.Int       `json:"value"           gencodec:"required"`
	Description  string         `json:"description"     gencodec:"required"`
	CreateHeight *big.Int       `json:"createHeight,omitempty"`
	UpdateHeight *big.Int       `json:"updateHeight,omitempty"`
}

type UnconfirmedPropertyInfo struct {
	Name        string              `json:"name"          gencodec:"required"`
	Value       *big.Int            `json:"value"         gencodec:"required"`
	Applicant   common.Address      `json:"applicant"     gencodec:"required"`
	Voters      []common.Address    `json:"voters,omitempty"`
	VoteResults []*big.Int          `json:"voteResults,omitempty"`
	Reason      string              `json:"reason,omitempty"`
	ApplyHeight *big.Int            `json:"applyHeight,omitempty"`
}