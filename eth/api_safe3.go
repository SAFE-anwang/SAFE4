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
	return contract_api.RedeemAvailable(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, sig)
}

func (api *PublicSafe3API) RedeemLocked(ctx context.Context, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
	return contract_api.RedeemLocked(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, sig)
}

func (api *PublicSafe3API) RedeemMasterNode(ctx context.Context, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes, enode string) (common.Hash, error) {
	return contract_api.RedeemMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, sig, enode)
}

func (api *PublicSafe3API) ApplyRedeemSpecial(ctx context.Context, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
	return contract_api.ApplyRedeemSpecial(ctx, api.blockChainAPI, api.transactionPoolAPI, from, pubkey, sig)
}

func (api *PublicSafe3API) Vote4Special(ctx context.Context, from common.Address, safe3Addr string, voteResult *big.Int) (common.Hash, error) {
	return contract_api.Vote4Special(ctx, api.blockChainAPI, api.transactionPoolAPI, from, safe3Addr, voteResult)
}

func (api *PublicSafe3API) GetAllAvailableNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetAllAvailableNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetAvailableInfos(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.AvailableSafe3Info, error) {
	return contract_api.GetAvailableInfos(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSafe3API) GetAvailableInfo(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.AvailableSafe3Info, error) {
	return contract_api.GetAvailableInfo(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) GetAllLockedNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetAllLockedNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetLockedAddrNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetLockedAddrNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetLockedAddrs(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	return contract_api.GetLockedAddrs(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSafe3API) GetLockedNum(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetLockedNum(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}

func (api *PublicSafe3API) GetLockedInfo(ctx context.Context, safe3Addr string, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.LockedSafe3Info, error) {
	return contract_api.GetLockedInfo(ctx, api.blockChainAPI, safe3Addr, start, count, blockNrOrHash)
}

func (api *PublicSafe3API) GetAllSpecialNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetAllSpecialNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSafe3API) GetSpecialInfos(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.SpecialSafe3Info, error) {
	return contract_api.GetSpecialInfos(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSafe3API) GetSpecialInfo(ctx context.Context, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.SpecialSafe3Info, error) {
	return contract_api.GetSpecialInfo(ctx, api.blockChainAPI, safe3Addr, blockNrOrHash)
}
