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

func GetSuperMasterNodeInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address) (*types.SuperMasterNodeInfo, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(SuperMasterNodeABI))
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
		To: &SuperMasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new(types.SuperMasterNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}
func GetAllSuperMasterNode(ctx context.Context, api *ethapi.PublicBlockChainAPI) ([]types.SuperMasterNodeInfo, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api");
	}

	vABI, err := abi.JSON(strings.NewReader(SuperMasterNodeABI))
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
		To: &SuperMasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	var (
		ret0 = new([]types.SuperMasterNodeInfo)
	)
	out := ret0
	if err := vABI.UnpackIntoInterface(out, method, result); err != nil {
		return nil, err
	}

	smnList := make([]types.SuperMasterNodeInfo, len(*ret0))
	for i, smn := range *ret0 {
		smnList[i] = smn
	}
	return smnList, nil
}

func GetTopSuperMasterNode(ctx context.Context, api *ethapi.PublicBlockChainAPI) ([]types.SuperMasterNodeInfo, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api");
	}

	vABI, err := abi.JSON(strings.NewReader(SuperMasterNodeABI))
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
		To: &SuperMasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	var (
		ret0 = new([]types.SuperMasterNodeInfo)
	)
	out := ret0
	if err := vABI.UnpackIntoInterface(out, method, result); err != nil {
		return nil, err
	}

	smnList := make([]types.SuperMasterNodeInfo, len(*ret0))
	for i, smn := range *ret0 {
		smnList[i] = smn
	}
	return smnList, nil
}

func GetSuperMasterNodeNum(ctx context.Context, api *ethapi.PublicBlockChainAPI) (*big.Int, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(SuperMasterNodeABI))
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
		To: &SuperMasterNodeContractAddr,
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

func RegisterSuperMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, isUnion bool, smnAddr common.Address, lockDay *big.Int, name string, enode string, pubkey string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	if blockChainAPI == nil || transactionPoolAPI == nil {
		return common.Hash{}, errors.New("invalid blockchain/transactionpool api")
	}

	vABI, err := abi.JSON(strings.NewReader(SuperMasterNodeABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "register"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := vABI.Pack(method, isUnion, smnAddr, lockDay, name, enode, pubkey, description, creatorIncentive, partnerIncentive, voterIncentive)
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

func AppendRegisterSuperMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, amount *big.Int, mnAddr common.Address, lockDay *big.Int) (common.Hash, error) {
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
