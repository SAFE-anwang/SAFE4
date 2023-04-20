package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ProposalInfo struct {
	Id           *big.Int
	Creator      common.Address
	Title        string
	PayAmount    *big.Int
	PayTimes     *big.Int
	StartPayTime *big.Int
	EndPayTime   *big.Int
	Description  string
	Detail       string
	Voters       []common.Address
	VoteResults  []*big.Int
	State        *big.Int
	CreateHeight *big.Int
	UpdateHeight *big.Int
}