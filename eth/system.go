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

func NewPublicSystemAPI(e *Ethereum, blockChainAPI *ethapi.PublicBlockChainAPI) *PublicSystemAPI {
	return &PublicSystemAPI{e, blockChainAPI}
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
