package contract_api

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"
)

func GetPropertyValue(ctx context.Context, api *ethapi.PublicBlockChainAPI, key string) (*big.Int, error) {
	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(systemcontracts.PropertyABI))
	if err != nil {
		return nil, err
	}

	method := "getValue"
	data, err := vABI.Pack(method, key)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.PropertyContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	value := new(big.Int)
	if err := vABI.UnpackIntoInterface(&value, method, result); err != nil {
		return nil, err
	}
	return value, nil
}