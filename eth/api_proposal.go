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

type PrivateProposalAPI struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPrivateProposalAPI(e *Ethereum) *PrivateProposalAPI {
	return &PrivateProposalAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PrivateProposalAPI) Create(ctx context.Context, from common.Address, title string, payAmount *hexutil.Big, payTimes *big.Int, startPayTime *big.Int, endPayTime *big.Int, description string) (common.Hash, error) {
	return contract_api.CreateProposal(ctx, api.blockChainAPI, api.transactionPoolAPI, from, title, payAmount, payTimes, startPayTime, endPayTime, description)
}

func (api *PrivateProposalAPI) Vote(ctx context.Context, from common.Address, id *big.Int, voteResult *big.Int) (common.Hash, error) {
	return contract_api.Vote4Proposal(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, voteResult)
}

func (api *PrivateProposalAPI) ChangeTitle(ctx context.Context, from common.Address, id *big.Int, title string) (common.Hash, error) {
	return contract_api.ChangeProposalTitle(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, title)
}

func (api *PrivateProposalAPI) ChangePayAmount(ctx context.Context, from common.Address, id *big.Int, payAmount *hexutil.Big) (common.Hash, error) {
	return contract_api.ChangeProposalPayAmount(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, payAmount)
}

func (api *PrivateProposalAPI) ChangePayTimes(ctx context.Context, from common.Address, id *big.Int, payTimes *big.Int) (common.Hash, error) {
	return contract_api.ChangeProposalPayTimes(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, payTimes)
}

func (api *PrivateProposalAPI) ChangeStartPayTime(ctx context.Context, from common.Address, id *big.Int, startPayTime *big.Int) (common.Hash, error) {
	return contract_api.ChangeProposalStartPayTime(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, startPayTime)
}

func (api *PrivateProposalAPI) ChangeEndPayTime(ctx context.Context, from common.Address, id *big.Int, endPayTime *big.Int) (common.Hash, error) {
	return contract_api.ChangeProposalEndPayTime(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, endPayTime)
}

func (api *PrivateProposalAPI) ChangeDescription(ctx context.Context, from common.Address, id *big.Int, description string) (common.Hash, error) {
	return contract_api.ChangeProposalDescription(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, description)
}

type PublicProposalAPI struct {
	e             *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
}

func NewPublicProposalAPI(e *Ethereum) *PublicProposalAPI {
	return &PublicProposalAPI{e, e.GetPublicBlockChainAPI()}
}

func (api *PublicProposalAPI) GetBalance(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetProposalBalance(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicProposalAPI) GetImmatureBalance(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetProposalImmatureBalance(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicProposalAPI) GetInfo(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.ProposalInfo, error) {
	return contract_api.GetProposalInfo(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicProposalAPI) GetRewardIDs(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return contract_api.GetProposalRewardIDs(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicProposalAPI) GetVoterNum(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetProposalVoterNum(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicProposalAPI) GetVoteInfo(ctx context.Context, id *big.Int, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.ProposalVoteInfo, error) {
	return contract_api.GetProposalVoteInfo(ctx, api.blockChainAPI, id, start, count, blockNrOrHash)
}

func (api *PublicProposalAPI) GetNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetProposalNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicProposalAPI) GetAll(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return contract_api.GetAllProposals(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicProposalAPI) GetMineNum(ctx context.Context, from common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetMineProposalNum(ctx, api.blockChainAPI, from, blockNrOrHash)
}

func (api *PublicProposalAPI) GetMines(ctx context.Context, from common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return contract_api.GetMineProposals(ctx, api.blockChainAPI, from, start, count, blockNrOrHash)
}

func (api *PublicProposalAPI) Exist(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistProposal(ctx, api.blockChainAPI, id, blockNrOrHash)
}
