package contract_api

import (
	"context"
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

func DepositAccount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, to common.Address, lockDay *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "deposit"
	data, err := vABI.Pack(method, to, lockDay)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", new(big.Int).SetInt64(int64(rpc.LatestBlockNumber)))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
		Data:     &msgData,
		Value:    (*hexutil.Big)(amount),
		GasPrice: (*hexutil.Big)(gasPrice),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	args.Gas = &gas
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func WithdrawAccount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "withdraw"
	data, err := vABI.Pack(method)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", new(big.Int).SetInt64(int64(rpc.LatestBlockNumber)))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
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

func WithdrawAccountByID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, ids []*big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "withdrawByID"
	data, err := vABI.Pack(method, ids)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", new(big.Int).SetInt64(int64(rpc.LatestBlockNumber)))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
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

func TransferAccount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, to common.Address, amount *big.Int, lockDay *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "transfer"
	data, err := vABI.Pack(method, to, amount, lockDay)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", new(big.Int).SetInt64(int64(rpc.LatestBlockNumber)))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
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

func AddAccountLockDay(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, day *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "addLockDay"
	data, err := vABI.Pack(method, id, day)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", new(big.Int).SetInt64(int64(rpc.LatestBlockNumber)))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
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

func GetAccountTotalAmount(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blocknumber *big.Int) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, api, "getTotalAmount", addr, blocknumber)
}

func GetAccountAvailableAmount(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blocknumber *big.Int) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, api, "getAvailableAmount", addr, blocknumber)
}

func GetAccountLockedAmount(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blocknumber *big.Int) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, api, "getLockedAmount", addr, blocknumber)
}

func GetAccountUsedAmount(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blocknumber *big.Int) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, api, "getUsedAmount", addr, blocknumber)
}

func getAccountAmountInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, method string, addr common.Address, blocknumber *big.Int) (*types.AccountAmountInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return nil, err
	}

	data, err := vABI.Pack(method, addr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.AccountManagerContractAddr,
		Data: &msgData,
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	unpacked, err := vABI.Unpack(method, result)
	if err != nil {
		return nil, err
	}

	info := new(types.AccountAmountInfo)
	info.Amount = unpacked[0].(*big.Int)
	info.IDs = unpacked[1].([]*big.Int)
	return info, nil
}

func GetAccountRecords(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blocknumber *big.Int) ([]types.AccountRecord, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return nil, err
	}

	method := "getRecords"
	data, err := vABI.Pack(method, addr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.AccountManagerContractAddr,
		Data: &msgData,
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	records := new([]types.AccountRecord)
	if err := vABI.UnpackIntoInterface(records, method, result); err != nil {
		return nil, err
	}
	return *records, nil
}

func GetAccountRecord0(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blocknumber *big.Int) (*types.AccountRecord, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return nil, err
	}

	method := "getRecord0"
	data, err := vABI.Pack(method, addr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.AccountManagerContractAddr,
		Data: &msgData,
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	record := new(types.AccountRecord)
	if err := vABI.UnpackIntoInterface(&record, method, result); err != nil {
		return nil, err
	}
	return record, nil
}

func GetAccountRecordByID(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blocknumber *big.Int) (*types.AccountRecord, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return nil, err
	}

	method := "getRecordByID"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.AccountManagerContractAddr,
		Data: &msgData,
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	record := new(types.AccountRecord)
	if err := vABI.UnpackIntoInterface(&record, method, result); err != nil {
		return nil, err
	}
	return record, nil
}

func GetAccountRecordUseInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blocknumber *big.Int) (*types.AccountRecordUseInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return nil, err
	}

	method := "getRecordUseInfo"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.AccountManagerContractAddr,
		Data: &msgData,
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	info := new(types.AccountRecordUseInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}