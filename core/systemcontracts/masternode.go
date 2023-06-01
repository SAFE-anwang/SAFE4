package systemcontracts

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"
)

func GetMasterNodeInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address) (*types.MasterNodeInfo, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(MasterNodeABI))
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
		To: &MasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new(types.MasterNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func GetNextMasterNode(ctx context.Context, api *ethapi.PublicBlockChainAPI) (*common.Address, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(MasterNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getNext"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &MasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	addr := new(common.Address)
	if err := vABI.UnpackIntoInterface(&addr, method, result); err != nil {
		return nil, err
	}
	return addr, nil
}

func GetAllMasterNode(ctx context.Context, api *ethapi.PublicBlockChainAPI) ([]types.MasterNodeInfo, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(MasterNodeABI))
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
		To: &MasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	var (
		ret0 = new([]types.MasterNodeInfo)
	)
	out := ret0
	if err := vABI.UnpackIntoInterface(out, method, result); err != nil {
		return nil, err
	}

	mnList := make([]types.MasterNodeInfo, len(*ret0))
	for i, mn := range *ret0 {
		mnList[i] = mn
	}
	return mnList, nil
}

func GetMasterNodeNum(ctx context.Context, api *ethapi.PublicBlockChainAPI) (*big.Int, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(MasterNodeABI))
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
		To: &MasterNodeContractAddr,
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

func RegisterMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, isUnion bool, mnAddr common.Address, lockDay *big.Int, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(MasterNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "register"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := vABI.Pack(method, isUnion, mnAddr, lockDay, enode, description, creatorIncentive, partnerIncentive)
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
		To:       &MasterNodeContractAddr,
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

func AppendRegisterMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, mnAddr common.Address, lockDay *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(MasterNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "appendRegister"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := vABI.Pack(method, mnAddr, lockDay)
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
		To:       &MasterNodeContractAddr,
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