package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"math/big"
)

type PublicSysPropertyAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicSysPropertyAPI(e *Ethereum) *PublicSysPropertyAPI {
	return &PublicSysPropertyAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicSysPropertyAPI) Add(ctx context.Context, from common.Address, name string, value *big.Int, description string) (common.Hash, error) {
	return contract_api.AddProperty(ctx, api.blockChainAPI, api.transactionPoolAPI, from, name, value, description)
}

func (api *PublicSysPropertyAPI) ApplyUpdate(ctx context.Context, from common.Address, name string, value *big.Int, reason string) (common.Hash, error) {
	return contract_api.ApplyUpdateProperty(ctx, api.blockChainAPI, api.transactionPoolAPI, from, name, value, reason)
}

func (api *PublicSysPropertyAPI) Vote4Update(ctx context.Context, from common.Address, name string, voteResult *big.Int) (common.Hash, error) {
	return contract_api.Vote4UpdateProperty(ctx, api.blockChainAPI, api.transactionPoolAPI, from, name, voteResult)
}

func (api *PublicSysPropertyAPI) GetInfo(ctx context.Context, name string, blocknumber *big.Int) (*types.PropertyInfo, error) {
	return contract_api.GetPropertyInfo(ctx, api.blockChainAPI, name, blocknumber)
}

func (api *PublicSysPropertyAPI) GetUnconfirmedInfo(ctx context.Context, name string, blocknumber *big.Int) (*types.UnconfirmedPropertyInfo, error) {
	return contract_api.GetUnconfirmedPropertyInfo(ctx, api.blockChainAPI, name, blocknumber)
}

func (api *PublicSysPropertyAPI) GetValue(ctx context.Context, name string, blocknumber *big.Int) (*big.Int, error) {
	return contract_api.GetPropertyValue(ctx, api.blockChainAPI, name, blocknumber)
}

func (api *PublicSysPropertyAPI) GetAll(ctx context.Context, blocknumber *big.Int) ([]types.PropertyInfo, error) {
	return contract_api.GetAllProperties(ctx, api.blockChainAPI, blocknumber)
}

func (api *PublicSysPropertyAPI) GetAllUnconfirmed(ctx context.Context, name string, blocknumber *big.Int) ([]types.UnconfirmedPropertyInfo, error) {
	return contract_api.GetAllUnconfirmedProperties(ctx, api.blockChainAPI, blocknumber)
}

func (api *PublicSysPropertyAPI) Exist(ctx context.Context, name string, blocknumber *big.Int) (bool, error) {
	return contract_api.ExistProperty(ctx, api.blockChainAPI, name, blocknumber)
}

func (api *PublicSysPropertyAPI) ExistUnconfirmed(ctx context.Context, name string, blocknumber *big.Int) (bool, error) {
	return contract_api.ExistUnconfirmedProperty(ctx, api.blockChainAPI, name, blocknumber)
}