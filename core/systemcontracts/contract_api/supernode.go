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

func RegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, isUnion bool, addr common.Address, lockDay *big.Int, name string, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "register"
	data, err := vABI.Pack(method, isUnion, addr, lockDay, name, enode, description, creatorIncentive, partnerIncentive, voterIncentive)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeLogicContractAddr,
		Data:     &msgData,
		Value:    (*hexutil.Big)(amount),
		GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	args.Gas = &gas
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func AppendRegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "appendRegister"
	data, err := vABI.Pack(method, addr, lockDay)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeLogicContractAddr,
		Data:     &msgData,
		Value:    (*hexutil.Big)(amount),
		GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	args.Gas = &gas
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func TurnRegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "turnRegister"
	data, err := vABI.Pack(method, addr, lockID)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeLogicContractAddr,
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

func ChangeSuperNodeAddress(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeAddress"
	data, err := vABI.Pack(method, addr, newAddr)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeLogicContractAddr,
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

func ChangeSuperNodeName(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, name string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeName"
	data, err := vABI.Pack(method, addr, name)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeLogicContractAddr,
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

func ChangeSuperNodeEnode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeEnode"
	data, err := vABI.Pack(method, addr, enode)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeLogicContractAddr,
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

func ChangeSuperNodeDescription(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, description string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeDescription"
	data, err := vABI.Pack(method, addr, description)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeLogicContractAddr,
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

func ChangeSuperNodeIsOfficial(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, flag bool) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeLogicABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeIsOfficial"
	data, err := vABI.Pack(method, addr, flag)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SuperNodeLogicContractAddr,
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

func GetSuperNodeInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.SuperNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	info := new(types.SuperNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetSuperNodeInfoByID(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SuperNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	info := new(types.SuperNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetAllSuperNodes(ctx context.Context, api *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	infos := new([]common.Address)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetTopSuperNodes(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
	if err != nil {
		return nil, err
	}

	method := "getTops"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	addrs := new([]common.Address)
	if err := vABI.UnpackIntoInterface(addrs, method, result); err != nil {
		return nil, err
	}
	return *addrs, nil
}

func GetOfficialSuperNodes(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	infos := new([]common.Address)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetSuperNodeNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
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

func ExistSuperNode(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
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

func ExistSuperNodeID(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
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

func ExistSuperNodeName(ctx context.Context, api *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
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

func ExistSuperNodeEnode(ctx context.Context, api *ethapi.PublicBlockChainAPI, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
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

func ExistSuperNodeLockID(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, lockID *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
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
		To: &systemcontracts.SuperNodeStorageContractAddr,
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

func IsValidSuperNode(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
	if err != nil {
		return false, err
	}

	method := "isValid"
	data, err := vABI.Pack(method, addr)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeStorageContractAddr,
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

func IsFormalSuperNode(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperNodeStorageABI))
	if err != nil {
		return false, err
	}

	method := "isFormal"
	data, err := vABI.Pack(method, addr)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperNodeStorageContractAddr,
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