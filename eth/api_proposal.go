package eth

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

type PublicProposalAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicProposalAPI(e *Ethereum) *PublicProposalAPI {
	return &PublicProposalAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicProposalAPI) Create(ctx context.Context, creatorAddr common.Address, title string, payAmount *big.Int, payTimes *big.Int, startPayTime *big.Int, endPayTime *big.Int, description string, detail string) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return nil, err
	}

	method := "create"
	data, err := vABI.Pack(method, creatorAddr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From: &creatorAddr,
		To: &systemcontracts.ProposalContractAddr,
		Data: &msgData,
	}
	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
	gas, err := ethapi.DoEstimateGas(ctx, api.e.APIBackend, args, blockNrOrHash, api.e.APIBackend.RPCGasCap())
	if err != nil {
		return nil, err
	}
	args.Gas = &gas
	result, err := api.blockChainAPI.Call(ctx, args, blockNrOrHash, nil)
	if err != nil {
		return nil, err
	}

	id := new(big.Int)
	if err := vABI.UnpackIntoInterface(&id, method, result); err != nil {
		return nil, err
	}
	return id, nil
}

func (api *PublicProposalAPI) Vote(ctx context.Context, addr common.Address, id *big.Int, voteResult *big.Int) error {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return err
	}

	method := "vote"
	data, err := vABI.Pack(method, id, voteResult)
	if err != nil {
		return err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From: &addr,
		To: &systemcontracts.ProposalContractAddr,
		Data: &msgData,
	}
	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(api.blockChainAPI.BlockNumber()))
	gas, err := ethapi.DoEstimateGas(ctx, api.e.APIBackend, args, blockNrOrHash, api.e.APIBackend.RPCGasCap())
	if err != nil {
		return err
	}
	args.Gas = &gas

	result, err := api.blockChainAPI.Call(ctx, args, blockNrOrHash, nil)
	if err != nil {
		return err
	}

	info := new([]common.Address)
	return vABI.UnpackIntoInterface(&info, method, result)
}

func (api *PublicProposalAPI) GetInfo(ctx context.Context, id *big.Int) (*types.PropertyInfo, error) {
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

	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new(types.PropertyInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func (api *PublicProposalAPI) GetAll(ctx context.Context) ([]types.PropertyInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
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
		To: &systemcontracts.ProposalContractAddr,
		Data: &msgData,
	}
	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new([]types.PropertyInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return *info, nil
}

func (api *PublicProposalAPI) GetMine(ctx context.Context, addr common.Address) ([]types.PropertyInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.ProposalABI))
	if err != nil {
		return nil, err
	}

	method := "getMine"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From: &addr,
		To: &systemcontracts.ProposalContractAddr,
		Data: &msgData,
	}
	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new([]types.PropertyInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return *info, nil
}
