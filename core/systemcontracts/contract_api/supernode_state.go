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

func UploadSuperNodeStates(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, ids []*big.Int, states []*big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SuperNodeStateContractAddr, "upload", getValues(ids, states))
}

func GetSuperNodeUploadEntries(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.StateEntry, error) {
	ret := new([]types.StateEntry)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStateContractAddr, "get", getValues(id), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetSuperNodeUploadState(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStateContractAddr, "getByAddr", getValues(id, addr), &ret)
	return ret, err
}
