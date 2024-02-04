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

func CreateProposal(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, title string, payAmount *hexutil.Big, payTimes *big.Int, startPayTime *big.Int, endPayTime *big.Int, description string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "create"
	data, err := vABI.Pack(method, title, payAmount.ToInt(), payTimes, startPayTime, endPayTime, description)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	value := big.NewInt(1000000000000000000) // 1 SAFE
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.ProposalContractAddr,
		Data:     &msgData,
		Value:    (*hexutil.Big)(value),
		GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
	}
	gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
	if err != nil {
		return common.Hash{}, err
	}
	args.Gas = &gas
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func Vote4Proposal(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, voteResult *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "vote"
	data, err := vABI.Pack(method, id, voteResult)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.ProposalContractAddr,
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

func ChangeProposalTitle(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, title string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeTitle"
	data, err := vABI.Pack(method, id, title)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.ProposalContractAddr,
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

func ChangeProposalPayAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, payAmount *hexutil.Big) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changePayAmount"
	data, err := vABI.Pack(method, id, payAmount.ToInt())
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.ProposalContractAddr,
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

func ChangeProposalPayTimes(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, payTimes *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changePayTimes"
	data, err := vABI.Pack(method, id, payTimes)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.ProposalContractAddr,
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

func ChangeProposalStartPayTime(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, startPayTime *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeStartPayTime"
	data, err := vABI.Pack(method, id, startPayTime)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.ProposalContractAddr,
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

func ChangeProposalEndPayTime(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, endPayTime *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changeEndPayTime"
	data, err := vABI.Pack(method, id, endPayTime)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.ProposalContractAddr,
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

func ChangeProposalDescription(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, description string) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "changePayDescription"
	data, err := vABI.Pack(method, id, description)
	if err != nil {
		return common.Hash{}, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From:     &from,
		To:       &systemcontracts.ProposalContractAddr,
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

func GetProposalBalance(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return nil, err
	}

	method := "getBalance"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.ProposalContractAddr,
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

func GetProposalInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.ProposalInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return nil, err
	}

	method := "getInfo"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.ProposalContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new(types.ProposalInfo)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return nil, err
	}
	return ret, nil
}

func GetProposalVoterNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return nil, err
	}

	method := "getVoterNum"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.ProposalContractAddr,
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

func GetProposalVoteInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.ProposalVoteInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return nil, err
	}

	method := "getVoteInfo"
	data, err := vABI.Pack(method, id, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.ProposalContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	ret := new([]types.ProposalVoteInfo)
	if err := vABI.UnpackIntoInterface(ret, method, result); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetProposalNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
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
		To: &systemcontracts.ProposalContractAddr,
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

func GetAllProposals(ctx context.Context, api *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
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
		To: &systemcontracts.ProposalContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return nil, err
	}

	infos := new([]*big.Int)
	if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
		return nil, err
	}
	return *infos, nil
}

func GetMineProposalNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, from common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return nil, err
	}

	method := "getMineNum"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From: &from,
		To: &systemcontracts.ProposalContractAddr,
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

func GetMineProposals(ctx context.Context, api *ethapi.PublicBlockChainAPI, from common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return nil, err
	}

	method := "getMines"
	data, err := vABI.Pack(method, start, count)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From: &from,
		To: &systemcontracts.ProposalContractAddr,
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

func ExistProposal(ctx context.Context, api *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return false, err
	}

	method := "exist"
	data, err := vABI.Pack(method, id)
	if err != nil {
		return false, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.ProposalContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, blockNrOrHash, nil)

	if err != nil {
		return false, err
	}

	ret := new(bool)
	if err := vABI.UnpackIntoInterface(&ret, method, result); err != nil {
		return false, err
	}
	return *ret, nil
}