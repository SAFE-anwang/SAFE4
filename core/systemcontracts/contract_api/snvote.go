package contract_api

import (
	"context"
	"fmt"
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

func VoteOrApproval(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, isVote bool, dstAddr common.Address, recordIDs []*big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "voteOrApproval"
	data, err := vABI.Pack(method, isVote, dstAddr, recordIDs)
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
		To:       &systemcontracts.SNVoteContractAddr,
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

func RemoveVoteOrApproval(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, recordIDs []*big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "removeVoteOrApproval"
	data, err := vABI.Pack(method, recordIDs)
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
		To:       &systemcontracts.SNVoteContractAddr,
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

func ProxyVote(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, snAddr common.Address) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "proxyVote"
	data, err := vABI.Pack(method, snAddr)
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
		To:       &systemcontracts.SNVoteContractAddr,
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

func GetSuperNodes4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blocknumber *big.Int) (*types.SNVoteRetInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getSuperNodes4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}

	if !blocknumber.IsInt64() {
		return nil, fmt.Errorf("big.Int is out of int64 range")
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

	info := new(types.SNVoteRetInfo)
	info.Addrs = unpacked[0].([]common.Address)
	info.VoteNums = unpacked[1].([]*big.Int)
	return info, nil
}

func GetRecordIDs4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blocknumber *big.Int) ([]big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getRecordIDs4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}

	if !blocknumber.IsInt64() {
		return nil, fmt.Errorf("big.Int is out of int64 range")
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	recordIDs := new([]big.Int)
	if err := vABI.UnpackIntoInterface(&recordIDs, method, result); err != nil {
		return nil, err
	}
	return *recordIDs, nil
}

func GetVoters4SN(ctx context.Context, api *ethapi.PublicBlockChainAPI, snAddr common.Address, blocknumber *big.Int) (*types.SNVoteRetInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVoters4SN"
	data, err := vABI.Pack(method, snAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}

	if !blocknumber.IsInt64() {
		return nil, fmt.Errorf("big.Int is out of int64 range")
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

	info := new(types.SNVoteRetInfo)
	info.Addrs = unpacked[0].([]common.Address)
	info.VoteNums = unpacked[1].([]*big.Int)
	return info, nil
}

func GetVoteNum4SN(ctx context.Context, api *ethapi.PublicBlockChainAPI, snAddr common.Address, blocknumber *big.Int) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVoteNum4SN"
	data, err := vABI.Pack(method, snAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}

	if !blocknumber.IsInt64() {
		return nil, fmt.Errorf("big.Int is out of int64 range")
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	num := new(big.Int)
	if err := vABI.UnpackIntoInterface(&num, method, result); err != nil {
		return nil, err
	}
	return num, nil
}

func GetProxies4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blocknumber *big.Int) (*types.SNVoteRetInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getProxies4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}

	if !blocknumber.IsInt64() {
		return nil, fmt.Errorf("big.Int is out of int64 range")
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

	info := new(types.SNVoteRetInfo)
	info.Addrs = unpacked[0].([]common.Address)
	info.VoteNums = unpacked[1].([]*big.Int)
	return info, nil
}

func GetProxiedRecordIDs4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blocknumber *big.Int) ([]big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getProxiedRecordIDs4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}

	if !blocknumber.IsInt64() {
		return nil, fmt.Errorf("big.Int is out of int64 range")
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	recordIDs := new([]big.Int)
	if err := vABI.UnpackIntoInterface(&recordIDs, method, result); err != nil {
		return nil, err
	}
	return *recordIDs, nil
}

func GetVoters4Proxy(ctx context.Context, api *ethapi.PublicBlockChainAPI, proxyAddr common.Address, blocknumber *big.Int) (*types.SNVoteRetInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVoters4Proxy"
	data, err := vABI.Pack(method, proxyAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}

	if !blocknumber.IsInt64() {
		return nil, fmt.Errorf("big.Int is out of int64 range")
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

	info := new(types.SNVoteRetInfo)
	info.Addrs = unpacked[0].([]common.Address)
	info.VoteNums = unpacked[1].([]*big.Int)
	return info, nil
}

func GetVoteNum4Proxy(ctx context.Context, api *ethapi.PublicBlockChainAPI, proxyAddr common.Address, blocknumber *big.Int) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVoteNum4Proxy"
	data, err := vABI.Pack(method, proxyAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}

	if !blocknumber.IsInt64() {
		return nil, fmt.Errorf("big.Int is out of int64 range")
	}

	//result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blocknumber.Int64())), nil)
	if err != nil {
		return nil, err
	}

	num := new(big.Int)
	if err := vABI.UnpackIntoInterface(&num, method, result); err != nil {
		return nil, err
	}
	return num, nil
}