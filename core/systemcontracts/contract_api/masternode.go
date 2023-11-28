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

func RegisterMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, isUnion bool, addr common.Address, lockDay *big.Int, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "register"
	data, err := vABI.Pack(method, isUnion, addr, lockDay, enode, description, creatorIncentive, partnerIncentive)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.MasterNodeLogicContractAddr,
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

func AppendRegisterMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeLogicABI))
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
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.MasterNodeLogicContractAddr,
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

func TurnRegisterMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeLogicABI))
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
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.MasterNodeLogicContractAddr,
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

func ChangeMasterNodeAddress(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeLogicABI))
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
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.MasterNodeLogicContractAddr,
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

func ChangeMasterNodeEnode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeLogicABI))
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
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.MasterNodeLogicContractAddr,
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

func ChangeMasterNodeDescription(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, description string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeLogicABI))
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
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.MasterNodeLogicContractAddr,
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

func ChangeMasterNodeIsOfficial(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, flag bool) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeIsOfficial"
	data, err := vABI.Pack(method, addr, flag)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	gasPrice := big.NewInt(params.GWei)
	gasPrice, err = GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		gasPrice = big.NewInt(params.GWei / 100)
	}

	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.MasterNodeLogicContractAddr,
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

func GetMasterNodeInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.MasterNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	info := new(types.MasterNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetMasterNodeInfoByID(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.MasterNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	info := new(types.MasterNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetNextMasterNode(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (common.Address, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
	if err != nil {
		return common.Address{}, err
	}

	method := "getNext"
	data, err := vABI.Pack(method)
	if err != nil {
		return common.Address{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.MasterNodeStorageContractAddr,
		Data: &msgData,
	}

	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return common.Address{}, err
	}

	addr := new(common.Address)
	if err := vABI.UnpackIntoInterface(&addr, method, result); err != nil {
		return common.Address{}, err
	}
	return *addr, nil
}

func GetAllMasterNodes(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]types.MasterNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	infos := new([]types.MasterNodeInfo)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetOfficialMasterNodes(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]types.MasterNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	infos := new([]types.MasterNodeInfo)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetMasterNodeNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	count := new(big.Int)
	if err := vABI.UnpackIntoInterface(&count, method, result); err != nil {
		return nil, err
	}
	return count, nil
}

func ExistMasterNode(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
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

func ExistMasterNodeID(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
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

func ExistMasterNodeEnode(ctx context.Context, api *ethapi.PublicBlockChainAPI, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
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

func ExistMasterNodeLockID(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, lockID *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeStorageABI))
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
		To: &systemcontracts.MasterNodeStorageContractAddr,
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