package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type StateEntry struct {
	Caller common.Address   `json:"caller"    gencodec:"required"`
	State  *big.Int         `json:"state"     gencodec:"required"`
}