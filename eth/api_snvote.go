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

type PublicSNVoteAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicSNVoteAPI(e *Ethereum) *PublicSNVoteAPI {
	return &PublicSNVoteAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicSNVoteAPI) VoteOrApproval(ctx context.Context, from common.Address, isVote bool, dstAddr common.Address, recordIDs []*big.Int) (common.Hash, error) {
	return contract_api.VoteOrApproval(ctx, api.blockChainAPI, api.transactionPoolAPI, from, isVote, dstAddr, recordIDs)
}

func (api *PublicSNVoteAPI) RemoveVoteOrApproval(ctx context.Context, from common.Address, recordIDs []*big.Int) (common.Hash, error) {
	return contract_api.RemoveVoteOrApproval(ctx, api.blockChainAPI, api.transactionPoolAPI, from, recordIDs)
}

func (api *PublicSNVoteAPI) ProxyVote(ctx context.Context, from common.Address, snAddr common.Address) (common.Hash, error) {
	return contract_api.ProxyVote(ctx, api.blockChainAPI, api.transactionPoolAPI, from, snAddr)
}

func (api *PublicSNVoteAPI) GetAmount4Voter(ctx context.Context, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetAmount4Voter(ctx, api.blockChainAPI, voterAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetVoteNum4Voter(ctx context.Context, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetVoteNum4Voter(ctx, api.blockChainAPI, voterAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetSNNum4Voter(ctx context.Context, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetSNNum4Voter(ctx, api.blockChainAPI, voterAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetSNs4Voter(ctx context.Context, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	return contract_api.GetSNs4Voter(ctx, api.blockChainAPI, voterAddr, start, count, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetProxyNum4Voter(ctx context.Context, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetProxyNum4Voter(ctx, api.blockChainAPI, voterAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetProxies4Voter(ctx context.Context, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	return contract_api.GetProxies4Voter(ctx, api.blockChainAPI, voterAddr, start, count, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetVotedIDNum4Voter(ctx context.Context, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetVotedIDNum4Voter(ctx, api.blockChainAPI, voterAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetVotedIDs4Voter(ctx context.Context, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return contract_api.GetVotedIDs4Voter(ctx, api.blockChainAPI, voterAddr, start, count, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetProxiedIDNum4Voter(ctx context.Context, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetProxiedIDNum4Voter(ctx, api.blockChainAPI, voterAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetProxiedIDs4Voter(ctx context.Context, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return contract_api.GetProxiedIDs4Voter(ctx, api.blockChainAPI, voterAddr, start, count, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetTotalAmount(ctx context.Context, dstAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetTotalAmount4SNOrProxy(ctx, api.blockChainAPI, dstAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetTotalVoteNum(ctx context.Context, dstAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetTotalVoteNum4SNOrProxy(ctx, api.blockChainAPI, dstAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetVoterNum(ctx context.Context, dstAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetVoterNum4SNOrProxy(ctx, api.blockChainAPI, dstAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetVoters(ctx context.Context, dstAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	return contract_api.GetVoters4SNOrProxy(ctx, api.blockChainAPI, dstAddr, start, count, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetIDNum(ctx context.Context, dstAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetIDNum4SNOrProxy(ctx, api.blockChainAPI, dstAddr, blockNrOrHash)
}

func (api *PublicSNVoteAPI) GetIDs(ctx context.Context, dstAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return contract_api.GetIDs4SNOrProxy(ctx, api.blockChainAPI, dstAddr, start, count, blockNrOrHash)
}
