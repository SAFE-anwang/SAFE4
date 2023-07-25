package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"math/big"
)

type PublicMasterNodeStateAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicMasterNodeStateAPI(e *Ethereum) *PublicMasterNodeStateAPI {
	return &PublicMasterNodeStateAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicMasterNodeStateAPI) GetAll(ctx context.Context) ([]types.StateInfo, error) {
	return systemcontracts.GetAllMasterNodeState(ctx, api.blockChainAPI)
}

func (api *PublicMasterNodeStateAPI) GetEntriesByAddr(ctx context.Context, addr common.Address) ([]types.StateEntry, error) {
	info, err := systemcontracts.GetMasterNodeInfo(ctx, api.blockChainAPI, addr)
	if err != nil {
		return nil, err
	}
	return systemcontracts.GetMasterNodeStateEntries(ctx, api.blockChainAPI, info.Id)
}

func (api *PublicMasterNodeStateAPI) GetEntriesByID(ctx context.Context, id *big.Int) ([]types.StateEntry, error) {
	return systemcontracts.GetMasterNodeStateEntries(ctx, api.blockChainAPI, id)
}

func (api *PublicMasterNodeStateAPI) Upload(ctx context.Context, from common.Address, ids []int64, states []uint8) (common.Hash, error) {
	return systemcontracts.UploadMasterNodeStates(ctx, api.blockChainAPI, api.transactionPoolAPI, from, ids, states)
}