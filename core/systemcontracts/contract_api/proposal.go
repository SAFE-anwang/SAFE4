package contract_api

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetProposalBalance(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getBalance", nil, &ret)
	return ret, err
}

func GetProposalImmatureBalance(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getImmatureBalance", nil, &ret)
	return ret, err
}

func CreateProposal(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, title string, payAmount *hexutil.Big, payTimes *big.Int, startPayTime *big.Int, endPayTime *big.Int, description string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, (*hexutil.Big)(big.NewInt(1000000000000000000)), systemcontracts.ProposalContractAddr, "create", getValues(title, payAmount, payTimes, startPayTime, endPayTime, description))
}

func Vote4Proposal(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, voteResult *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.ProposalContractAddr, "vote", getValues(id, voteResult))
}

func ChangeProposalTitle(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, title string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.ProposalContractAddr, "changeTitle", getValues(id, title))
}

func ChangeProposalPayAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, payAmount *hexutil.Big) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.ProposalContractAddr, "changePayAmount", getValues(id, payAmount))
}

func ChangeProposalPayTimes(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, payTimes *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.ProposalContractAddr, "changePayTimes", getValues(id, payTimes))
}

func ChangeProposalStartPayTime(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, startPayTime *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.ProposalContractAddr, "changeStartPayTime", getValues(id, startPayTime))
}

func ChangeProposalEndPayTime(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, endPayTime *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.ProposalContractAddr, "changeEndPayTime", getValues(id, endPayTime))
}

func ChangeProposalDescription(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, description string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.ProposalContractAddr, "changeDescription", getValues(id, description))
}

func GetProposalInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.ProposalInfo, error) {
	ret := new(types.ProposalInfo)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getInfo", getValues(id), &ret)
	return ret, err
}

func GetProposalRewardIDs(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	ret := new([]*big.Int)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getRewardIDs", getValues(id), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetProposalVoterNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getVoterNum", getValues(id), &ret)
	return ret, err
}

func GetProposalVoteInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.ProposalVoteInfo, error) {
	ret := new([]types.ProposalVoteInfo)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getVoteInfo", getValues(id, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetProposalNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getNum", nil, &ret)
	return ret, err
}

func GetAllProposals(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	ret := new([]*big.Int)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getAll", getValues(start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetMineProposalNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, creator common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getMineNum", getValues(creator), &ret)
	return ret, err
}

func GetMineProposals(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, creator common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	ret := new([]*big.Int)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "getMines", getValues(creator, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func ExistProposal(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.ProposalContractAddr, "exist", getValues(id), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}
