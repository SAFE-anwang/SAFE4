package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type PublicSafe3API struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicSafe3API(e *Ethereum) *PublicSafe3API {
	return &PublicSafe3API{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicSafe3API) RedeemAvailable(ctx context.Context, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
	return contract_api.RedeemAvailableSafe3(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, sig)
}

func (api *PublicSafe3API) RedeemLocked(ctx context.Context, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes, enode string) (common.Hash, error) {
	return contract_api.RedeemLockedSafe3(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, sig, enode)
}

func (api *PublicSafe3API) ApplyRedeemSpecial(ctx context.Context, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
	return contract_api.ApplyRedeemSpecialSafe3(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, sig)
}

func (api *PublicSafe3API) Vote4Special(ctx context.Context, from common.Address, safe3Addr string, voteResult *big.Int) (common.Hash, error) {
	return contract_api.Vote4SpecialSafe3(ctx, api.blockChainAPI, api.transactionPoolAPI, from, safe3Addr, voteResult)
}

func (api *PublicSafe3API) GetAvailable(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.AvailableSafe3Info, error) {
	return contract_api.GetAvailableSafe3(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) GetLocked(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) ([]types.LockedSafe3Info, error) {
	return contract_api.GetLockedSafe3(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) GetSpecial(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.SpecialSafe3Info, error) {
	return contract_api.GetSpecialSafe3(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) GetAllAvailable(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]types.AvailableSafe3Info, error) {
	return contract_api.GetAllAvailableSafe3(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetAllLocked(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]types.LockedSafe3Info, error) {
	return contract_api.GetAllLockedSafe3(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetAllSpecial(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]types.SpecialSafe3Info, error) {
	return contract_api.GetAllSpecialSafe3(ctx, api.blockChainAPI, blockNrOrHash)
}
