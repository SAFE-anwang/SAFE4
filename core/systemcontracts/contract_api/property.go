package contract_api

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
)

func AddProperty(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, name string, value *big.Int, description string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.PropertyContractAddr, "add", getValues(name, value, description))
}

func ApplyUpdateProperty(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, name string, value *big.Int, reason string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.PropertyContractAddr, "applyUpdate", getValues(name, value, reason))
}

func Vote4UpdateProperty(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, name string, voteResult *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.PropertyContractAddr, "vote4Update", getValues(name, voteResult))
}

func GetPropertyInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (*types.PropertyInfo, error) {
	ret := new(types.PropertyInfo)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "getInfo", getValues(name), &ret)
	return ret, err
}

func GetUnconfirmedPropertyInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (*types.UnconfirmedPropertyInfo, error) {
	ret := new(types.UnconfirmedPropertyInfo)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "getUnconfirmedInfo", getValues(name), &ret)
	return ret, err
}

func GetPropertyValue(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "getValue", getValues(name), &ret)
	return ret, err
}

func GetPropertyNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "getNum", nil, &ret)
	return ret, err
}

func GetAllProperties(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	ret := new([]string)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "getAll", getValues(start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetUnconfirmedPropertyNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "getUnconfirmedNum", nil, &ret)
	return ret, err
}

func GetAllUnconfirmedProperties(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	ret := new([]string)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "getAllUnconfirmed", getValues(start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func ExistProperty(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "exist", getValues(name), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistUnconfirmedProperty(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.PropertyContractAddr, "existUnconfirmed", getValues(name), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}
