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

func DepositAccount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *hexutil.Big, to common.Address, lockDay *big.Int) (common.Hash, error) {
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
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
		Data:     &msgData,
		Value:    amount,
		GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
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
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
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
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
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

func TransferAccount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, to common.Address, amount *hexutil.Big, lockDay *big.Int) (common.Hash, error) {
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
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
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
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.AccountManagerContractAddr,
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

func GetAccountTotalAmount(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, api, "getTotalAmount", addr, blockNrOrHash)
}

func GetAccountAvailableAmount(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, api, "getAvailableAmount", addr, blockNrOrHash)
}

func GetAccountLockedAmount(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, api, "getLockedAmount", addr, blockNrOrHash)
}

func GetAccountUsedAmount(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, api, "getUsedAmount", addr, blockNrOrHash)
}

func getAccountAmountInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, method string, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
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

	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	unpacked, err := vABI.Unpack(method, result)
	if err != nil {
		return nil, err
	}

	info := new(types.AccountAmountInfo)
	info.Amount = unpacked[0].(*big.Int)
	info.Num = unpacked[1].(*big.Int)
	return info, nil
}

func GetAccountTotalIDs(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return getAccountIDs(ctx, api, "getTotalIDs", addr, start, count, blockNrOrHash)
}

func GetAccountAvailableIDs(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return getAccountIDs(ctx, api, "getAvailableIDs", addr, start, count, blockNrOrHash)
}

func GetAccountLockedIDs(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return getAccountIDs(ctx, api, "getLockedIDs", addr, start, count, blockNrOrHash)
}

func GetAccountUsedIDs(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return getAccountIDs(ctx, api, "getUsedIDs", addr, start, count, blockNrOrHash)
}

func getAccountIDs(ctx context.Context, api *ethapi.PublicBlockChainAPI, method string, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.AccountManagerABI))
	if err != nil {
		return nil, err
	}

	data, err := vABI.Pack(method, addr, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.AccountManagerContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	records := new([]*big.Int)
	if err := vABI.UnpackIntoInterface(records, method, result); err != nil {
		return nil, err
	}
	return *records, nil
}

func GetAccountRecord0(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecord, error) {
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
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	record := new(types.AccountRecord)
	if err := vABI.UnpackIntoInterface(&record, method, result); err != nil {
		return nil, err
	}
	return record, nil
}

func GetAccountRecordByID(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecord, error) {
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
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	record := new(types.AccountRecord)
	if err := vABI.UnpackIntoInterface(&record, method, result); err != nil {
		return nil, err
	}
	return record, nil
}

func GetAccountRecordUseInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecordUseInfo, error) {
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
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	info := new(types.AccountRecordUseInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}