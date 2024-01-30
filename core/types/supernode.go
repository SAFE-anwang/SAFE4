package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SuperNodeMemberInfo struct {
	LockID *big.Int         `json:"lockID"    gencodec:"required"`
	Addr   common.Address   `json:"addr"      gencodec:"required"`
	Amount *big.Int         `json:"amount"    gencodec:"required"`
	Height *big.Int         `json:"height"    gencodec:"required"`
}

type SuperNodeIncentivePlan struct {
	Creator *big.Int        `json:"creator"   gencodec:"required"`
	Partner *big.Int        `json:"partner"   gencodec:"required"`
	Voter   *big.Int        `json:"voter"     gencodec:"required"`
}

type SuperNodeInfo struct {
	Id                  *big.Int                `json:"id"            gencodec:"required"`
	Name                string                  `json:"name"          gencodec:"required"`
	Addr                common.Address          `json:"addr"          gencodec:"required"`
	Creator             common.Address          `json:"creator"       gencodec:"required"`
	Enode               string                  `json:"enode"         gencodec:"required"`
	Description         string                  `json:"description"   gencodec:"required"`
	IsOfficial          bool                    `json:"isOfficial"    gencodec:"required"`
	State               *big.Int                `json:"state"         gencodec:"required"`
	Founders            []SuperNodeMemberInfo   `json:"founders"      gencodec:"required"`
	IncentivePlan       SuperNodeIncentivePlan  `json:"incentivePlan" gencodec:"required"`
	LastRewardHeight    *big.Int                `json:"lastRewardHeight" gencodec:"required"`
	CreateHeight        *big.Int                `json:"createHeight"  gencodec:"required"`
	UpdateHeight        *big.Int                `json:"updateHeight"  gencodec:"required"`
}
