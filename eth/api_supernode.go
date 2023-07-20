package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type PublicSuperNodeAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicSuperNodeAPI(e *Ethereum, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI) *PublicSuperNodeAPI {
	return &PublicSuperNodeAPI{e, blockChainAPI, transactionPoolAPI}
}

func (api *PublicSuperNodeAPI) Start(ctx context.Context, addr common.Address) (bool, error) {
	info, err := api.GetInfo(ctx, addr)
	if err != nil {
		return false, err
	}
	curBlock := api.e.blockchain.CurrentBlock()
	ping, err := types.NewNodePing(info.Id, types.SuperNodeType, curBlock.Hash(), curBlock.Number(), api.e.p2pServer.Config.PrivateKey)
	if err != nil {
		return false, err
	}
	api.e.eventMux.Post(core.NodePingEvent{Ping: ping})
	return true, nil
}

func (api *PublicSuperNodeAPI) Stop(ctx context.Context, addr common.Address) (bool, error) {
	log.Info("Stop supernode", "address", addr)
	return true, nil
}

func (api *PublicSuperNodeAPI) Restart(ctx context.Context, addr common.Address) (bool, error) {
	log.Info("Restart supernode", "address", addr)
	return true, nil
}

func (api *PublicSuperNodeAPI) GetInfo(ctx context.Context, addr common.Address) (*types.SuperNodeInfo, error) {
	return systemcontracts.GetSuperNodeInfo(ctx, api.blockChainAPI, addr)
}

func (api *PublicSuperNodeAPI) GetAll(ctx context.Context) ([]types.SuperNodeInfo, error) {
	return systemcontracts.GetAllSuperNode(ctx, api.blockChainAPI)
}

func (api *PublicSuperNodeAPI) GetTop(ctx context.Context) ([]types.SuperNodeInfo, error) {
	return systemcontracts.GetTopSuperNode(ctx, api.blockChainAPI)
}

func (api *PublicSuperNodeAPI) GetNum(ctx context.Context) (*big.Int, error) {
	return systemcontracts.GetSuperNodeNum(ctx, api.blockChainAPI)
}

func (api *PublicSuperNodeAPI) RegisterSuperNode(ctx context.Context, from common.Address, amount *big.Int, isUnion bool, snAddr common.Address, lockDay *big.Int, name string, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	return systemcontracts.RegisterSuperNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, isUnion, snAddr, lockDay, name, enode, description, creatorIncentive, partnerIncentive, voterIncentive)
}

func (api *PublicSuperNodeAPI) AppendRegisterSuperNode(ctx context.Context, from common.Address, amount *big.Int, snAddr common.Address, lockDay *big.Int) (common.Hash, error) {
	return systemcontracts.AppendRegisterSuperNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, snAddr, lockDay)
}