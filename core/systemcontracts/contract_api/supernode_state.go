package contract_api

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"
)

func UploadSuperNodeStates(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, ids []*big.Int, states []*big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStateABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "upload"
	data, err := vABI.Pack(method, ids, states)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeStateContractAddr,
		Data:     &msgData,
		GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	gas = gas * 3 / 2
	args.Gas = &gas
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func GetSuperNodeUploadEntries(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.StateEntry, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStateABI))
	if err != nil {
		return nil, err
	}

	method := "get"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeStateContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	entries := new([]types.StateEntry)
	if err := vABI.UnpackIntoInterface(entries, method, result); err != nil {
		return nil, err
	}
	return *entries, nil
}