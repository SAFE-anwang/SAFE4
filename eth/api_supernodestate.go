package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"math/big"
)

type PublicSuperNodeStateAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicSuperNodeStateAPI(e *Ethereum) *PublicSuperNodeStateAPI {
	return &PublicSuperNodeStateAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicSuperNodeStateAPI) GetAll(ctx context.Context) ([]types.StateInfo, error) {
	return contract_api.GetAllSuperNodeState(ctx, api.blockChainAPI)
}

func (api *PublicSuperNodeStateAPI) GetEntries(ctx context.Context, id *big.Int) ([]types.StateEntry, error) {
	return contract_api.GetSuperNodeStateEntries(ctx, api.blockChainAPI, id)
}

func (api *PublicSuperNodeStateAPI) GetEntriesByAddr(ctx context.Context, addr common.Address) ([]types.StateEntry, error) {
	info, err := contract_api.GetSuperNodeInfo(ctx, api.blockChainAPI, addr)
	if err != nil {
		return nil, err
	}
	return contract_api.GetSuperNodeStateEntries(ctx, api.blockChainAPI, info.Id)
}

func (api *PublicSuperNodeStateAPI) GetEntriesByID(ctx context.Context, id *big.Int) ([]types.StateEntry, error) {
	return contract_api.GetSuperNodeStateEntries(ctx, api.blockChainAPI, id)
}

func (api *PublicSuperNodeStateAPI) Upload(ctx context.Context, from common.Address, ids []*big.Int, states []*big.Int) (common.Hash, error) {
	return contract_api.UploadSuperNodeStates(ctx, api.blockChainAPI, api.transactionPoolAPI, from, ids, states)
}