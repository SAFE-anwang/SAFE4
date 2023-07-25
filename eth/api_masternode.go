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

type PublicMasterNodeAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicMasterNodeAPI(e *Ethereum) *PublicMasterNodeAPI {
	return &PublicMasterNodeAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PublicMasterNodeAPI) Start(ctx context.Context, addr common.Address) (bool, error) {
	info, err := api.GetInfo(ctx, addr)
	if err != nil {
		return false, err
	}
	curBlock := api.e.blockchain.CurrentBlock()
	ping, err := types.NewNodePing(info.Id, types.MasterNodeType, curBlock.Hash(), curBlock.Number(), api.e.p2pServer.Config.PrivateKey)
	if err != nil {
		return false, err
	}
	api.e.eventMux.Post(core.NodePingEvent{Ping: ping})
	return true, nil
}

func (api *PublicMasterNodeAPI) Stop(ctx context.Context, addr common.Address) (bool, error) {
	log.Info("Stop masternode", "address", addr)
	return true, nil
}

func (api *PublicMasterNodeAPI) Restart(ctx context.Context, addr common.Address) (bool, error) {
	log.Info("Restart masternode", "address", addr)
	return true, nil
}

func (api *PublicMasterNodeAPI) GetInfo(ctx context.Context, addr common.Address) (*types.MasterNodeInfo, error) {
	return systemcontracts.GetMasterNodeInfo(ctx, api.blockChainAPI, addr)
}

func (api *PublicMasterNodeAPI) GetInfoByID(ctx context.Context, id *big.Int) (*types.MasterNodeInfo, error) {
	return systemcontracts.GetMasterNodeInfoByID(ctx, api.blockChainAPI, id)
}

func (api *PublicMasterNodeAPI) GetNext(ctx context.Context) (*common.Address, error) {
	return systemcontracts.GetNextMasterNode(ctx, api.blockChainAPI)
}

func (api *PublicMasterNodeAPI) GetAll(ctx context.Context) ([]types.MasterNodeInfo, error) {
	return systemcontracts.GetAllMasterNode(ctx, api.blockChainAPI)
}

func (api *PublicMasterNodeAPI) GetNum(ctx context.Context) (*big.Int, error) {
	return systemcontracts.GetMasterNodeNum(ctx, api.blockChainAPI)
}

func (api *PublicMasterNodeAPI) Register(ctx context.Context, from common.Address, amount *big.Int, isUnion bool, mnAddr common.Address, lockDay *big.Int, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int) (common.Hash, error) {
	return systemcontracts.RegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, isUnion, mnAddr, lockDay, enode, description, creatorIncentive, partnerIncentive)
}

func (api *PublicMasterNodeAPI) AppendRegister(ctx context.Context, from common.Address, amount *big.Int, mnAddr common.Address, lockDay *big.Int) (common.Hash, error) {
	return systemcontracts.AppendRegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, mnAddr, lockDay)
}