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

func DepositAccount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, to common.Address, lockDay *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, value, systemcontracts.AccountManagerContractAddr, "deposit", getValues(to, lockDay))
}

func BatchDeposit4One(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, to common.Address, times *big.Int, spaceDay *big.Int, startDay *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, value, systemcontracts.AccountManagerContractAddr, "batchDeposit4One", getValues(to, times, spaceDay, startDay))
}

func BatchDeposit4Multi(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, addrs []common.Address, times *big.Int, spaceDay *big.Int, startDay *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, value, systemcontracts.AccountManagerContractAddr, "batchDeposit4Multi", getValues(addrs, times, spaceDay, startDay))
}

func WithdrawAccount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.AccountManagerContractAddr, "withdraw", nil)
}

func WithdrawAccountByID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, ids []*big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.AccountManagerContractAddr, "withdrawByID", getValues(ids))
}

func TransferAccount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, to common.Address, amount *hexutil.Big, lockDay *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.AccountManagerContractAddr, "transfer", getValues(to, amount, lockDay))
}

func AddAccountLockDay(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, day *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.AccountManagerContractAddr, "addLockDay", getValues(id, day))
}

func GetAccountTotalAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, blockChainAPI, "getTotalAmount", addr, blockNrOrHash)
}

func GetAccountAvailableAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, blockChainAPI, "getAvailableAmount", addr, blockNrOrHash)
}

func GetAccountLockedAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, blockChainAPI, "getLockedAmount", addr, blockNrOrHash)
}

func GetAccountUsedAmount(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return getAccountAmountInfo(ctx, blockChainAPI, "getUsedAmount", addr, blockNrOrHash)
}

func getAccountAmountInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, method string, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	outs, err := QueryContract4MultiReturn(ctx, blockChainAPI, blockNrOrHash, systemcontracts.AccountManagerContractAddr, method, getValues(addr))
	if err != nil {
		return nil, err
	}
	ret := new(types.AccountAmountInfo)
	ret.Amount = outs[0].(*big.Int)
	ret.Num = outs[1].(*big.Int)
	return ret, nil
}

func GetAccountTotalIDs(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return getAccountIDs(ctx, blockChainAPI, "getTotalIDs", addr, start, count, blockNrOrHash)
}

func GetAccountAvailableIDs(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return getAccountIDs(ctx, blockChainAPI, "getAvailableIDs", addr, start, count, blockNrOrHash)
}

func GetAccountLockedIDs(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return getAccountIDs(ctx, blockChainAPI, "getLockedIDs", addr, start, count, blockNrOrHash)
}

func GetAccountUsedIDs(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	return getAccountIDs(ctx, blockChainAPI, "getUsedIDs", addr, start, count, blockNrOrHash)
}

func getAccountIDs(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, method string, addr common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]*big.Int, error) {
	ret := new([]*big.Int)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.AccountManagerContractAddr, method, getValues(addr, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetAccountRecord0(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecord, error) {
	ret := new(types.AccountRecord)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.AccountManagerContractAddr, "getRecord0", getValues(addr), &ret)
	return ret, err
}

func GetAccountRecordByID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecord, error) {
	ret := new(types.AccountRecord)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.AccountManagerContractAddr, "getRecordByID", getValues(id), &ret)
	return ret, err
}

func GetAccountRecordUseInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecordUseInfo, error) {
	ret := new(types.AccountRecordUseInfo)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.AccountManagerContractAddr, "getRecordUseInfo", getValues(id), &ret)
	return ret, err
}
