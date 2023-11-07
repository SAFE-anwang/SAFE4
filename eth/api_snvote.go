package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
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

func (api *PublicSNVoteAPI) GetSuperNodes4Voter(ctx context.Context, voterAddr common.Address) (*types.SNVoteRetInfo, error) {
	return contract_api.GetSuperNodes4Voter(ctx, api.blockChainAPI, voterAddr)
}

func (api *PublicSNVoteAPI) GetRecordIDs4Voter(ctx context.Context, voterAddr common.Address) ([]big.Int, error) {
	return contract_api.GetRecordIDs4Voter(ctx, api.blockChainAPI, voterAddr)
}

func (api *PublicSNVoteAPI) GetVoters4SN(ctx context.Context, snAddr common.Address) (*types.SNVoteRetInfo, error) {
	return contract_api.GetVoters4SN(ctx, api.blockChainAPI, snAddr)
}

func (api *PublicSNVoteAPI) GetVoteNum4SN(ctx context.Context, snAddr common.Address) (*big.Int, error) {
	return contract_api.GetVoteNum4SN(ctx, api.blockChainAPI, snAddr)
}

func (api *PublicSNVoteAPI) GetProxies4Voter(ctx context.Context, voterAddr common.Address) (*types.SNVoteRetInfo, error) {
	return contract_api.GetProxies4Voter(ctx, api.blockChainAPI, voterAddr)
}

func (api *PublicSNVoteAPI) GetProxiedRecordIDs4Voter(ctx context.Context, voterAddr common.Address) ([]big.Int, error) {
	return contract_api.GetProxiedRecordIDs4Voter(ctx, api.blockChainAPI, voterAddr)
}

func (api *PublicSNVoteAPI) GetVoters4Proxy(ctx context.Context, proxyAddr common.Address) (*types.SNVoteRetInfo, error) {
	return contract_api.GetVoters4Proxy(ctx, api.blockChainAPI, proxyAddr)
}

func (api *PublicSNVoteAPI) GetVoteNum4Proxy(ctx context.Context, proxyAddr common.Address) (*big.Int, error) {
	return contract_api.GetVoteNum4Proxy(ctx, api.blockChainAPI, proxyAddr)
}
