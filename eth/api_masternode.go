package eth

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"
)

type PublicMasterNodeAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
	enode string
}

func NewPublicMasterNodeAPI(e *Ethereum) *PublicMasterNodeAPI {
	return &PublicMasterNodeAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI(), ""}
}

func (api *PublicMasterNodeAPI) Start(ctx context.Context, addr common.Address) (bool, error) {
	if len(api.enode) == 0 {
		temp := api.e.p2pServer.NodeInfo().Enode
		arr := strings.Split(temp, "?")
		if len(arr) == 0 {
			return false, errors.New("start masternode failed, invalid local enode")
		}
		api.enode = arr[0]
	}

	info, err := api.GetInfo(ctx, addr)
	if err != nil {
		return false, err
	}
	if api.enode != info.Enode {
		return false, errors.New("start failed, incompatible masternode enode, local: [" + api.enode + "], state: [" + info.Enode + "]")
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
	return contract_api.GetMasterNodeInfo(ctx, api.blockChainAPI, addr)
}

func (api *PublicMasterNodeAPI) GetInfoByID(ctx context.Context, id *big.Int) (*types.MasterNodeInfo, error) {
	return contract_api.GetMasterNodeInfoByID(ctx, api.blockChainAPI, id)
}

func (api *PublicMasterNodeAPI) GetNext(ctx context.Context) (*common.Address, error) {
	return contract_api.GetNextMasterNode(ctx, api.blockChainAPI)
}

func (api *PublicMasterNodeAPI) GetAll(ctx context.Context) ([]types.MasterNodeInfo, error) {
	return contract_api.GetAllMasterNode(ctx, api.blockChainAPI)
}

func (api *PublicMasterNodeAPI) GetNum(ctx context.Context) (*big.Int, error) {
	return contract_api.GetMasterNodeNum(ctx, api.blockChainAPI)
}

func (api *PublicMasterNodeAPI) Register(ctx context.Context, from common.Address, amount *big.Int, isUnion bool, mnAddr common.Address, lockDay *big.Int, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int) (common.Hash, error) {
	return contract_api.RegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, isUnion, mnAddr, lockDay, enode, description, creatorIncentive, partnerIncentive)
}

func (api *PublicMasterNodeAPI) AppendRegister(ctx context.Context, from common.Address, amount *big.Int, mnAddr common.Address, lockDay *big.Int) (common.Hash, error) {
	return contract_api.AppendRegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, mnAddr, lockDay)
}