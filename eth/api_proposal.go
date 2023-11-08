package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"math/big"
)

type PublicProposalAPI struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicProposalAPI(e *Ethereum) *PublicProposalAPI {
	return &PublicProposalAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicProposalAPI) Create(ctx context.Context, from common.Address, title string, payAmount *big.Int, payTimes *big.Int, startPayTime *big.Int, endPayTime *big.Int, description string) (common.Hash, error) {
	return contract_api.CreateProposal(ctx, api.blockChainAPI, api.transactionPoolAPI, from, title, payAmount, payTimes, startPayTime, endPayTime, description)
}

func (api *PublicProposalAPI) Vote(ctx context.Context, from common.Address, id *big.Int, voteResult *big.Int) (common.Hash, error) {
	return contract_api.Vote4Proposal(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, voteResult)
}

func (api *PublicProposalAPI) ChangeTitle(ctx context.Context, from common.Address, id *big.Int, title string) (common.Hash, error) {
	return contract_api.ChangeProposalTitle(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, title)
}

func (api *PublicProposalAPI) ChangePayAmount(ctx context.Context, from common.Address, id *big.Int, payAmount *big.Int) (common.Hash, error) {
	return contract_api.ChangeProposalPayAmount(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, payAmount)
}

func (api *PublicProposalAPI) ChangePayTimes(ctx context.Context, from common.Address, id *big.Int, payTimes *big.Int) (common.Hash, error) {
	return contract_api.ChangeProposalPayTimes(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, payTimes)
}

func (api *PublicProposalAPI) ChangeStartPayTime(ctx context.Context, from common.Address, id *big.Int, startPayTime *big.Int) (common.Hash, error) {
	return contract_api.ChangeProposalStartPayTime(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, startPayTime)
}

func (api *PublicProposalAPI) ChangeEndPayTime(ctx context.Context, from common.Address, id *big.Int, endPayTime *big.Int) (common.Hash, error) {
	return contract_api.ChangeProposalEndPayTime(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, endPayTime)
}

func (api *PublicProposalAPI) ChangeDescription(ctx context.Context, from common.Address, id *big.Int, description string) (common.Hash, error) {
	return contract_api.ChangeProposalDescription(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, description)
}

func (api *PublicProposalAPI) GetInfo(ctx context.Context, id *big.Int) (*types.ProposalInfo, error) {
	return contract_api.GetProposalInfo(ctx, api.blockChainAPI, id)
}

func (api *PublicProposalAPI) GetAll(ctx context.Context) ([]types.ProposalInfo, error) {
	return contract_api.GetAllProposals(ctx, api.blockChainAPI)
}

func (api *PublicProposalAPI) GetMines(ctx context.Context, from common.Address) ([]types.ProposalInfo, error) {
	return contract_api.GetMineProposals(ctx, api.blockChainAPI, from)
}

func (api *PublicProposalAPI) Exist(ctx context.Context, id *big.Int) (bool, error) {
	return contract_api.ExistProposal(ctx, api.blockChainAPI, id)
}