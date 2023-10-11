package contract_api

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"
)

func GetAllSuperNodeState(ctx context.Context, api *ethapi.PublicBlockChainAPI) ([]types.StateInfo, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStateABI))
	if err != nil {
		return nil, err
	}

	method := "getAllState"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeStateContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	infos := new([]types.StateInfo)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetSuperNodeStateEntries(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int) ([]types.StateEntry, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api");
	}

	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStateABI))
	if err != nil {
		return nil, err
	}

	method := "getEntries"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeStateContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	entries := new([]types.StateEntry)
	if err := vABI.UnpackIntoInterface(entries, method, result); err != nil {
		return nil, err
	}
	return *entries, nil
}

func UploadSuperNodeStates(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, ids []*big.Int, states []uint8) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStateABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "uploadState"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := vABI.Pack(method, ids, states)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price")
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeStateContractAddr,
		Data:     &msgData,
		GasPrice: (*hexutil.Big)(gasPrice),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	args.Gas = &gas

	return transactionPoolAPI.SendTransaction(ctx, args)
}