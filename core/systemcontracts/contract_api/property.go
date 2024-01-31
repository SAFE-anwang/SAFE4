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

func AddProperty(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, name string, value *big.Int, description string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "add"
	data, err := vABI.Pack(method, name, value, description)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.PropertyContractAddr,
		Data:     &msgData,
		GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	args.Gas = &gas
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func ApplyUpdateProperty(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, name string, value *big.Int, reason string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "applyUpdate"
	data, err := vABI.Pack(method, name, value, reason)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.PropertyContractAddr,
		Data:     &msgData,
		GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	args.Gas = &gas
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func Vote4UpdateProperty(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, name string, voteResult *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "vote4Update"
	data, err := vABI.Pack(method, name, voteResult)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.PropertyContractAddr,
		Data:     &msgData,
		GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	args.Gas = &gas
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func GetPropertyInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (*types.PropertyInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getInfo"
	data, err := vABI.Pack(method, name)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	info := new(types.PropertyInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetUnconfirmedPropertyInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (*types.UnconfirmedPropertyInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getUnconfirmedInfo"
	data, err := vABI.Pack(method, name)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	info := new(types.UnconfirmedPropertyInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetPropertyValue(ctx context.Context, api *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getValue"
	data, err := vABI.Pack(method, name)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	value := new(big.Int)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return nil, err
	}
	return value, nil
}

func GetPropertyNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getNum"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	value := new(big.Int)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return nil, err
	}
	return value, nil
}

func GetAllProperties(ctx context.Context, api *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getAll"
	data, err := vABI.Pack(method, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	infos := new([]string)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetUnconfirmedPropertyNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getUnconfirmedNum"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	value := new(big.Int)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return nil, err
	}
	return value, nil
}

func GetAllUnconfirmedProperties(ctx context.Context, api *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getAllUnconfirmed"
	data, err := vABI.Pack(method, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	infos := new([]string)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func ExistProperty(ctx context.Context, api *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return false, err
	}

	method := "exist"
	data, err := vABI.Pack(method, name)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return false, err
	}

	value := new(bool)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return false, err
	}
	return *value, nil
}

func ExistUnconfirmedProperty(ctx context.Context, api *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return false, err
	}

	method := "existUnconfirmed"
	data, err := vABI.Pack(method, name)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return false, err
	}

	value := new(bool)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return false, err
	}
	return *value, nil
}