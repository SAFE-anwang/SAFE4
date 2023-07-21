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

func GetAllSuperNodeState(ctx context.Context, api *ethapi.PublicBlockChainAPI) ([]big.Int, []uint8, error) {
	if api == nil {
		return nil, nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(SuperNodeStateABI))
	if err != nil {
		return nil, nil, err
	}

	method := "getAllState"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &SuperNodeStateContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	out := make(map[string]interface{})
	if err := vABI.UnpackIntoMap(out, method, result); err != nil {
		return nil, nil, err
	}
	if _, ok := out["ids"]; !ok {
		return nil, nil, nil
	}
	return out["ids"].([]big.Int), out["states"].([]uint8), nil
}

func GetSuperNodeStateEntries(ctx context.Context, api *ethapi.PublicBlockChainAPI, addr common.Address) ([]types.StateEntry, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api");
	}

	vABI, err := abi.JSON(strings.NewReader(SuperNodeStateABI))
	if err != nil {
		return nil, err
	}

	method := "getEntries"
	data, err := vABI.Pack(method, addr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &SuperNodeStateContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	var (
		ret0 = new([]types.StateEntry)
	)
	out := ret0
	if err := vABI.UnpackIntoInterface(out, method, result); err != nil {
		return nil, err
	}

	entries := make([]types.StateEntry, len(*ret0))
	for i, sn := range *ret0 {
		entries[i] = sn
	}
	return entries, nil
}

func UploadSuperNodeStates(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, ids []int64, states []uint8) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(SuperNodeStateABI))
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
		To:       &SuperNodeStateContractAddr,
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