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
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SNVoteContractAddr,
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
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SNVoteContractAddr,
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
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.SNVoteContractAddr,
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

func GetAmount4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getAmount4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetVoteNum4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVoteNum4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetSNNum4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getSNNum4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetSNs4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getSNs4Voter"
	data, err := vABI.Pack(method, voterAddr, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
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

	ret := new(types.SNVoteRetInfo)
	ret.Addrs = unpacked[0].([]common.Address)
	ret.VoteNums = unpacked[1].([]*big.Int)
	return ret, nil
}

func GetProxyNum4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getProxyNum4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetProxies4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getProxies4Voter"
	data, err := vABI.Pack(method, voterAddr, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
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

	ret := new(types.SNVoteRetInfo)
	ret.Addrs = unpacked[0].([]common.Address)
	ret.VoteNums = unpacked[1].([]*big.Int)
	return ret, nil
}

func GetVotedIDNum4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVotedIDNum4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetVotedIDs4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVotedIDs4Voter"
	data, err := vABI.Pack(method, voterAddr, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new([]*big.Int)
	if err := vABI.UnpackIntoInterface(ret, method, result); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetProxiedIDNum4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getProxiedIDNum4Voter"
	data, err := vABI.Pack(method, voterAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetProxiedIDs4Voter(ctx context.Context, api *ethapi.PublicBlockChainAPI, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getProxiedIDs4Voter"
	data, err := vABI.Pack(method, voterAddr, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new([]*big.Int)
	if err := vABI.UnpackIntoInterface(ret, method, result); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetTotalAmount4SNOrProxy(ctx context.Context, api *ethapi.PublicBlockChainAPI, dstAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getTotalAmount"
	data, err := vABI.Pack(method, dstAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetTotalVoteNum4SNOrProxy(ctx context.Context, api *ethapi.PublicBlockChainAPI, dstAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getTotalVoteNum"
	data, err := vABI.Pack(method, dstAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetVoterNum4SNOrProxy(ctx context.Context, api *ethapi.PublicBlockChainAPI, dstAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVoterNum"
	data, err := vABI.Pack(method, dstAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetVoters4SNOrProxy(ctx context.Context, api *ethapi.PublicBlockChainAPI, dstAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getVoters"
	data, err := vABI.Pack(method, dstAddr, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
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

	ret := new(types.SNVoteRetInfo)
	ret.Addrs = unpacked[0].([]common.Address)
	ret.VoteNums = unpacked[1].([]*big.Int)
	return ret, nil
}

func GetIDNum4SNOrProxy(ctx context.Context, api *ethapi.PublicBlockChainAPI, dstAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getIDNum"
	data, err := vABI.Pack(method, dstAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(big.Int)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetIDs4SNOrProxy(ctx context.Context, api *ethapi.PublicBlockChainAPI, dstAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SNVoteABI))
	if err != nil {
		return nil, err
	}

	method := "getIDs"
	data, err := vABI.Pack(method, dstAddr, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SNVoteContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new([]*big.Int)
	if err := vABI.UnpackIntoInterface(ret, method, result); err != nil {
		return nil, err
	}
	return *ret, nil
}
