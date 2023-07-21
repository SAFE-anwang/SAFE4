package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
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

func (api *PublicSuperNodeStateAPI) GetAll(ctx context.Context) ([]big.Int, []uint8, error) {
	return systemcontracts.GetAllSuperNodeState(ctx, api.blockChainAPI)
}

func (api *PublicSuperNodeStateAPI) GetEntries(ctx context.Context, addr common.Address) ([]types.StateEntry, error) {
	return systemcontracts.GetSuperNodeStateEntries(ctx, api.blockChainAPI, addr)
}

func (api *PublicSuperNodeStateAPI) Upload(ctx context.Context, from common.Address, ids []int64, states []uint8) (common.Hash, error) {
	return systemcontracts.UploadSuperNodeStates(ctx, api.blockChainAPI, api.transactionPoolAPI, from, ids, states)
}