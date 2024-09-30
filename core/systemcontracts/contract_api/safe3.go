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

func BatchRedeemAvailable(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkeys []hexutil.Bytes, sigs []hexutil.Bytes, targetAddr common.Address) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.Safe3ContractAddr, "batchRedeemAvailable", getValues(pubkeys, sigs, targetAddr))
}

func BatchRedeemLocked(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkeys []hexutil.Bytes, sigs []hexutil.Bytes, targetAddr common.Address) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.Safe3ContractAddr, "batchRedeemLocked", getValues(pubkeys, sigs, targetAddr))
}

func BatchRedeemMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkeys []hexutil.Bytes, sigs []hexutil.Bytes, enodes []string, targetAddr common.Address) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.Safe3ContractAddr, "batchRedeemMasterNode", getValues(pubkeys, sigs, enodes, targetAddr))
}

func ApplyRedeemSpecial(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkey []byte, sig []byte) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.Safe3ContractAddr, "applyRedeemSpecial", getValues(pubkey, sig))
}

func Vote4Special(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, safe3Addr string, voteResult *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.Safe3ContractAddr, "vote4Special", getValues(safe3Addr, voteResult))
}

func GetAllAvailableNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getAllAvailableNum", nil, &ret)
	return ret, err
}

func GetAvailableInfos(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.AvailableSafe3Info, error) {
	ret := new([]types.AvailableSafe3Info)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getAvailableInfos", getValues(start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetAvailableInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.AvailableSafe3Info, error) {
	ret := new(types.AvailableSafe3Info)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getAvailableInfo", getValues(safe3Addr), &ret)
	return ret, err
}

func GetAllLockedNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getAllLockedNum", nil, &ret)
	return ret, err
}

func GetLockedAddrNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getLockedAddrNum", nil, &ret)
	return ret, err
}

func GetLockedAddrs(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
	ret := new([]string)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getLockedAddrs", getValues(start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetLockedNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getLockedNum", getValues(safe3Addr), &ret)
	return ret, err
}

func GetLockedInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, safe3Addr string, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.LockedSafe3Info, error) {
	ret := new([]types.LockedSafe3Info)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getLockedInfo", getValues(safe3Addr, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetAllSpecialNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getAllSpecialNum", nil, &ret)
	return ret, err
}

func GetSpecialInfos(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.SpecialSafe3Info, error) {
	ret := new([]types.SpecialSafe3Info)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getSpecialInfos", getValues(start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetSpecialInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.SpecialSafe3Info, error) {
	ret := new(types.SpecialSafe3Info)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "getSpecialInfo", getValues(safe3Addr), &ret)
	return ret, err
}

func ExistAvailableNeedToRedeem(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "existAvailableNeedToRedeem", getValues(safe3Addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistLockedNeedToRedeem(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "existLockedNeedToRedeem", getValues(safe3Addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistMasterNodeNeedToRedeem(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.Safe3ContractAddr, "existMasterNodeNeedToRedeem", getValues(safe3Addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}
