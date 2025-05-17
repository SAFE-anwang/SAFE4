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

type PrivateNodeStateAPI struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPrivateNodeStateAPI(e *Ethereum) *PrivateNodeStateAPI {
	return &PrivateNodeStateAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PrivateNodeStateAPI) UploadMnStates(ctx context.Context, from common.Address, ids []*big.Int, states []*big.Int) (common.Hash, error) {
	return contract_api.UploadMasterNodeStates(ctx, api.blockChainAPI, api.transactionPoolAPI, from, ids, states)
}

func (api *PrivateNodeStateAPI) UploadSnStates(ctx context.Context, from common.Address, ids []*big.Int, states []*big.Int) (common.Hash, error) {
	return contract_api.UploadSuperNodeStates(ctx, api.blockChainAPI, api.transactionPoolAPI, from, ids, states)
}

type PublicNodeStateAPI struct {
	e             *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
}

func NewPublicNodeStateAPI(e *Ethereum) *PublicNodeStateAPI {
	return &PublicNodeStateAPI{e, e.GetPublicBlockChainAPI()}
}

func (api *PublicNodeStateAPI) GetMnTempStates(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.StateEntry, error) {
	return contract_api.GetMasterNodeUploadEntries(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicNodeStateAPI) GetSnTempStates(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.StateEntry, error) {
	return contract_api.GetSuperNodeUploadEntries(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicNodeStateAPI) GetMnTempStateByAddr(ctx context.Context, id *big.Int, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetMasterNodeUploadState(ctx, api.blockChainAPI, id, addr, blockNrOrHash)
}

func (api *PublicNodeStateAPI) GetSnTempStateByAddr(ctx context.Context, id *big.Int, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetSuperNodeUploadState(ctx, api.blockChainAPI, id, addr, blockNrOrHash)
}
