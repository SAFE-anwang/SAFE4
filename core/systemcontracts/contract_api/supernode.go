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

func RegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, isUnion bool, addr common.Address, lockDay *big.Int, name string, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "register"
	data, err := vABI.Pack(method, isUnion, addr, lockDay, name, enode, description, creatorIncentive, partnerIncentive, voterIncentive)
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
		To:       &systemcontracts.SuperNodeContractAddr,
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

func AppendRegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "appendRegister"
	data, err := vABI.Pack(method, addr, lockDay)
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
		To:       &systemcontracts.SuperNodeContractAddr,
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

func TurnRegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "turnRegister"
	data, err := vABI.Pack(method, addr, lockID)
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
		To:       &systemcontracts.SuperNodeContractAddr,
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

func ChangeSuperNodeAddress(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeAddress"
	data, err := vABI.Pack(method, addr, newAddr)
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
		To:       &systemcontracts.SuperNodeContractAddr,
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

func ChangeSuperNodeEnode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeEnode"
	data, err := vABI.Pack(method, addr, enode)
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
		To:       &systemcontracts.SuperNodeContractAddr,
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

func ChangeSuperNodeName(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, name string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeEnode"
	data, err := vABI.Pack(method, addr, name)
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
		To:       &systemcontracts.SuperNodeContractAddr,
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

func ChangeSuperNodeDescription(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, description string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeDescription"
	data, err := vABI.Pack(method, addr, description)
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
		To:       &systemcontracts.SuperNodeContractAddr,
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

func ChangeSuperNodeOfficial(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, flag bool) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeOfficial"
	data, err := vABI.Pack(method, addr, flag)
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
		To:       &systemcontracts.SuperNodeContractAddr,
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

func GetSuperNodeInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address) (*types.SuperNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getInfo"
	data, err := vABI.Pack(method, addr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new(types.SuperNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetSuperNodeInfoByID(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int) (*types.SuperNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getInfoByID"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new(types.SuperNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetAllSuperNodes(ctx context.Context, api *ethapi.PublicBlockChainAPI) ([]types.SuperNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getAll"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	infos := new([]types.SuperNodeInfo)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetTopSuperNodes(ctx context.Context, api *ethapi.PublicBlockChainAPI) ([]types.SuperNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getTop"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	infos := new([]types.SuperNodeInfo)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetOfficialSuperNodes(ctx context.Context, api *ethapi.PublicBlockChainAPI) ([]types.SuperNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getOfficials"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	infos := new([]types.SuperNodeInfo)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetSuperNodeNum(ctx context.Context, api *ethapi.PublicBlockChainAPI) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
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
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	count := new(big.Int)
	if err := vABI.UnpackIntoInterface(&count, method, result); err != nil {
		return nil, err
	}
	return count, nil
}

func ExistSuperNode(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return false, err
	}

	method := "exist"
	data, err := vABI.Pack(method, addr)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return false, err
	}

	value := new(bool)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return false, err
	}
	return *value, nil
}

func ExistSuperNodeID(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return false, err
	}

	method := "existID"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return false, err
	}

	value := new(bool)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return false, err
	}
	return *value, nil
}

func ExistSuperNodeName(ctx context.Context, api *ethapi.PublicBlockChainAPI, name string) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return false, err
	}

	method := "existName"
	data, err := vABI.Pack(method, name)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return false, err
	}

	value := new(bool)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return false, err
	}
	return *value, nil
}

func ExistSuperNodeEnode(ctx context.Context, api *ethapi.PublicBlockChainAPI, enode string) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return false, err
	}

	method := "existEnode"
	data, err := vABI.Pack(method, enode)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return false, err
	}

	value := new(bool)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return false, err
	}
	return *value, nil
}

func ExistSuperNodeLockID(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, lockID *big.Int) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeABI))
	if err != nil {
		return false, err
	}

	method := "existLockID"
	data, err := vABI.Pack(method, addr, lockID)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return false, err
	}

	value := new(bool)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return false, err
	}
	return *value, nil
}