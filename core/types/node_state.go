package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type StateEntry struct {
	Addr   common.Address   `json:"addr"      gencodec:"required"`
	state  uint8            `json:"state"     gencodec:"required"`
}