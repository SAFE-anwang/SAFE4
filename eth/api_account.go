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

type PublicAccountAPI struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicAccountAPI(e *Ethereum) *PublicAccountAPI {
	return &PublicAccountAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicAccountAPI) Deposit(ctx context.Context, from common.Address, amount *big.Int, to common.Address, lockDay *big.Int) (common.Hash, error) {
	return contract_api.DepositAccount(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, to, lockDay)
}

func (api *PublicAccountAPI) Withdraw(ctx context.Context, from common.Address) (common.Hash, error) {
	return contract_api.WithdrawAccount(ctx, api.blockChainAPI, api.transactionPoolAPI, from)
}

func (api *PublicAccountAPI) WithdrawByID(ctx context.Context, from common.Address, ids []*big.Int) (common.Hash, error) {
	return contract_api.WithdrawAccountByID(ctx, api.blockChainAPI, api.transactionPoolAPI, from, ids)
}

func (api *PublicAccountAPI) Transfer(ctx context.Context, from common.Address, to common.Address, amount *big.Int, lockDay *big.Int) (common.Hash, error) {
	return contract_api.TransferAccount(ctx, api.blockChainAPI, api.transactionPoolAPI, from, to, amount, lockDay)
}

func (api *PublicAccountAPI) AddLockDay(ctx context.Context, from common.Address, id *big.Int, day *big.Int) (common.Hash, error) {
	return contract_api.AddAccountLockDay(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, day)
}

func (api *PublicAccountAPI) GetTotalAmount(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return contract_api.GetAccountTotalAmount(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicAccountAPI) GetAvailableAmount(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return contract_api.GetAccountAvailableAmount(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicAccountAPI) GetLockedAmount(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return contract_api.GetAccountLockedAmount(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicAccountAPI) GetUsedAmount(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountAmountInfo, error) {
	return contract_api.GetAccountUsedAmount(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicAccountAPI) GetRecords(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) ([]types.AccountRecord, error) {
	return contract_api.GetAccountRecords(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicAccountAPI) GetRecord0(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecord, error) {
	return contract_api.GetAccountRecord0(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicAccountAPI) GetRecordByID(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecord, error) {
	return contract_api.GetAccountRecordByID(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicAccountAPI) GetRecordUseInfo(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.AccountRecordUseInfo, error) {
	return contract_api.GetAccountRecordUseInfo(ctx, api.blockChainAPI, id, blockNrOrHash)
}
