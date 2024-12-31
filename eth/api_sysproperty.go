package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type PrivateSysPropertyAPI struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPrivateSysPropertyAPI(e *Ethereum) *PrivateSysPropertyAPI {
	return &PrivateSysPropertyAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PrivateSysPropertyAPI) Add(ctx context.Context, from common.Address, name string, value *big.Int, description string) (common.Hash, error) {
	return contract_api.AddProperty(ctx, api.blockChainAPI, api.transactionPoolAPI, from, name, value, description)
}

func (api *PrivateSysPropertyAPI) ApplyUpdate(ctx context.Context, from common.Address, name string, value *big.Int, reason string) (common.Hash, error) {
	return contract_api.ApplyUpdateProperty(ctx, api.blockChainAPI, api.transactionPoolAPI, from, name, value, reason)
}

func (api *PrivateSysPropertyAPI) Vote4Update(ctx context.Context, from common.Address, name string, voteResult *big.Int) (common.Hash, error) {
	return contract_api.Vote4UpdateProperty(ctx, api.blockChainAPI, api.transactionPoolAPI, from, name, voteResult)
}

type PublicSysPropertyAPI struct {
	e             *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
}

func NewPublicSysPropertyAPI(e *Ethereum) *PublicSysPropertyAPI {
	return &PublicSysPropertyAPI{e, e.GetPublicBlockChainAPI()}
}

func (api *PublicSysPropertyAPI) GetInfo(ctx context.Context, name string, blockNrOrHash rpc.BlockNumberOrHash) (*types.PropertyInfo, error) {
	return contract_api.GetPropertyInfo(ctx, api.blockChainAPI, name, blockNrOrHash)
}

func (api *PublicSysPropertyAPI) GetUnconfirmedInfo(ctx context.Context, name string, blockNrOrHash rpc.BlockNumberOrHash) (*types.UnconfirmedPropertyInfo, error) {
	return contract_api.GetUnconfirmedPropertyInfo(ctx, api.blockChainAPI, name, blockNrOrHash)
}

func (api *PublicSysPropertyAPI) GetValue(ctx context.Context, name string, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetPropertyValue(ctx, api.blockChainAPI, name, blockNrOrHash)
}

func (api *PublicSysPropertyAPI) GetNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetPropertyNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSysPropertyAPI) GetAll(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	return contract_api.GetAllProperties(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSysPropertyAPI) GetUnconfirmedNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetUnconfirmedPropertyNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSysPropertyAPI) GetAllUnconfirmed(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	return contract_api.GetAllUnconfirmedProperties(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSysPropertyAPI) Exist(ctx context.Context, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistProperty(ctx, api.blockChainAPI, name, blockNrOrHash)
}

func (api *PublicSysPropertyAPI) ExistUnconfirmed(ctx context.Context, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistUnconfirmedProperty(ctx, api.blockChainAPI, name, blockNrOrHash)
}
