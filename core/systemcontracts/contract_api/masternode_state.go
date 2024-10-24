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

func UploadMasterNodeStates(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, ids []*big.Int, states []*big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.MasterNodeStateContractAddr, "upload", getValues(ids, states))
}

func GetMasterNodeUploadEntries(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.StateEntry, error) {
	ret := new([]types.StateEntry)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStateContractAddr, "get", getValues(id), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}
