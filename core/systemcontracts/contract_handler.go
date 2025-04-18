package systemcontracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"strings"
)

func IsSystemRewardTx(tx *types.Transaction) bool {
	if tx.To() == nil || len(tx.Data()) == 0 {
		return false
	}
	if *tx.To() != SystemRewardContractAddr {
		return false
	}

	vABI, err := abi.JSON(strings.NewReader(SystemRewardABI))
	if err != nil {
		log.Info("IsSystemRewardTx: parse SystemRewardABI failed", "error", err)
		return false
	}
	method, err := vABI.MethodById(tx.Data())
	if err != nil {
		return false
	}
	return method.Name == "reward"
}

func IsSystemRewardMessage(to *common.Address, data []byte) bool {
	if to == nil || len(data) == 0 {
		return false
	}
	if *to != SystemRewardContractAddr {
		return false
	}

	vABI, err := abi.JSON(strings.NewReader(SystemRewardABI))
	if err != nil {
		log.Info("IsSystemRewardMessage: parse SystemRewardABI failed", "error", err)
		return false
	}
	method, err := vABI.MethodById(data)
	if err != nil {
		return false
	}
	return method.Name == "reward"
}

func IsNodeStateMessage(to *common.Address, data []byte) bool {
	if to == nil || len(data) == 0 {
		return false
	}
	if *to != MasterNodeStateContractAddr && *to != SuperNodeStateContractAddr {
		return false
	}

	vABI, err := abi.JSON(strings.NewReader(MasterNodeStateABI))
	if err != nil {
		log.Info("IsNodeStateMessage: parse NodeStateABI failed", "error", err)
		return false
	}
	method, err := vABI.MethodById(data)
	if err != nil {
		return false
	}
	return method.Name == "upload"
}
