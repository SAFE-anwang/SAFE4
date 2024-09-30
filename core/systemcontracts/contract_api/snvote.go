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

func VoteOrApproval(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, isVote bool, dstAddr common.Address, recordIDs []*big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SNVoteContractAddr, "voteOrApproval", getValues(isVote, dstAddr, recordIDs))
}

func VoteOrApprovalWithAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, isVote bool, dstAddr common.Address) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, value, systemcontracts.SNVoteContractAddr, "voteOrApprovalWithAmount", getValues(isVote, dstAddr))
}

func RemoveVoteOrApproval(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, recordIDs []*big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SNVoteContractAddr, "removeVoteOrApproval", getValues(recordIDs))
}

func ProxyVote(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, snAddr common.Address) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SNVoteContractAddr, "proxyVote", getValues(snAddr))
}

func GetAmount4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getAmount4Voter", getValues(voterAddr), &ret)
	return ret, err
}

func GetVoteNum4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getVoteNum4Voter", getValues(voterAddr), &ret)
	return ret, err
}

func GetSNNum4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getSNNum4Voter", getValues(voterAddr), &ret)
	return ret, err
}

func GetSNs4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	outs, err := QueryContract4MultiReturn(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getSNs4Voter", getValues(voterAddr, start, count))
	if err != nil {
		return nil, err
	}
	ret := new(types.SNVoteRetInfo)
	ret.Addrs = outs[0].([]common.Address)
	ret.VoteNums = outs[1].([]*big.Int)
	return ret, nil
}

func GetProxyNum4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getProxyNum4Voter", getValues(voterAddr), &ret)
	return ret, err
}

func GetProxies4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	outs, err := QueryContract4MultiReturn(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getProxies4Voter", getValues(voterAddr, start, count))
	if err != nil {
		return nil, err
	}
	ret := new(types.SNVoteRetInfo)
	ret.Addrs = outs[0].([]common.Address)
	ret.VoteNums = outs[1].([]*big.Int)
	return ret, nil
}

func GetVotedIDNum4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getVotedIDNum4Voter", getValues(voterAddr), &ret)
	return ret, err
}

func GetVotedIDs4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	ret := new([]*big.Int)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getVotedIDs4Voter", getValues(voterAddr, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetProxiedIDNum4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getProxiedIDNum4Voter", getValues(voterAddr), &ret)
	return ret, err
}

func GetProxiedIDs4Voter(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, voterAddr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	ret := new([]*big.Int)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getProxiedIDs4Voter", getValues(voterAddr, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetTotalAmount4SNOrProxy(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getTotalAmount", getValues(addr), &ret)
	return ret, err
}

func GetTotalVoteNum4SNOrProxy(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getTotalVoteNum", getValues(addr), &ret)
	return ret, err
}

func GetVoterNum4SNOrProxy(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getVoterNum", getValues(addr), &ret)
	return ret, err
}

func GetVoters4SNOrProxy(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SNVoteRetInfo, error) {
	outs, err := QueryContract4MultiReturn(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getVoters", getValues(addr, start, count))
	if err != nil {
		return nil, err
	}
	ret := new(types.SNVoteRetInfo)
	ret.Addrs = outs[0].([]common.Address)
	ret.VoteNums = outs[1].([]*big.Int)
	return ret, nil
}

func GetIDNum4SNOrProxy(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getIDNum", getValues(addr), &ret)
	return ret, err
}

func GetIDs4SNOrProxy(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	ret := new([]*big.Int)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getIDs", getValues(addr, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetAllVoteAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getAllAmount", nil, &ret)
	return ret, err
}

func GetAllVoteNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getAllVoteNum", nil, &ret)
	return ret, err
}

func GetAllProxiedAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getAllProxiedAmount", nil, &ret)
	return ret, err
}

func GetAllProxiedVoteNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SNVoteContractAddr, "getAllProxiedVoteNum", nil, &ret)
	return ret, err
}
