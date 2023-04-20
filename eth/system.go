package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"
)

type PublicSystemAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
}

func NewPublicSystemAPI(e *Ethereum) *PublicSystemAPI {
	return &PublicSystemAPI{e, ethapi.NewPublicBlockChainAPI(e.APIBackend)}
}

func (api *PublicSystemAPI) GetProperty(ctx context.Context, name string) (*types.PropertyInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getInfo"
	data, err := vABI.Pack(method, name)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
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

func (api *PublicSystemAPI) GetPropertyValue(ctx context.Context, name string) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getValue"
	data, err := vABI.Pack(method, name)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	value := new(big.Int)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return nil, err
	}
	return value, nil
}
/*
func (api *PublicSystemAPI) Reward(ctx context.Context, smnAddr common.Address, smnCount *big.Int, mnAddr common.Address, mnCount *big.Int) (common.Hash, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SystemRewardABI))
	if err != nil {
		return common.Hash{}, err
	}

	method := "reward"
	data, err := vABI.Pack(method, smnAddr, smnCount, mnAddr, mnCount)
	if err != nil {
		return common.Hash{}, err
	}

	value := new(big.Int)
	value.Add(smnCount, mnCount)
	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		From: &smnAddr,
		To: &systemcontracts.SystemRewardContractAddr,
		Data: &msgData,
	}
	blockNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
	gas, err := api.blockChainAPI.EstimateGas(ctx, args, &blockNrOrHash)
	args.Gas = &gas

	hash, err := ethapi.NewPrivateAccountAPI(api.e.APIBackend, new(ethapi.AddrLocker)).SendTransaction(ctx, args, "123")
	if err != nil {
		return common.Hash{}, err
	}
	return hash, err
}
*/
