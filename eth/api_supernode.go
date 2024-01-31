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
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"
)

type PublicSuperNodeAPI struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
	enode              string
}

func NewPublicSuperNodeAPI(e *Ethereum) *PublicSuperNodeAPI {
	return &PublicSuperNodeAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI(), ""}
}

func (api *PublicSuperNodeAPI) Start(ctx context.Context, addr common.Address) (bool, error) {
	if len(api.enode) == 0 {
		temp := api.e.p2pServer.NodeInfo().Enode
		arr := strings.Split(temp, "?")
		if len(arr) == 0 {
			return false, errors.New("start supernode failed, invalid local enode")
		}
		api.enode = arr[0]
	}

	info, err := api.GetInfo(ctx, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return false, err
	}
	if api.enode != info.Enode {
		return false, errors.New("start failed, incompatible supernode enode, local: [" + api.enode + "], state: [" + info.Enode + "]")
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

func (api *PublicSuperNodeAPI) Register(ctx context.Context, from common.Address, amount *big.Int, isUnion bool, addr common.Address, lockDay *big.Int, name string, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	return contract_api.RegisterSuperNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, isUnion, addr, lockDay, name, enode, description, creatorIncentive, partnerIncentive, voterIncentive)
}

func (api *PublicSuperNodeAPI) AppendRegister(ctx context.Context, from common.Address, amount *big.Int, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	return contract_api.AppendRegisterSuperNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, addr, lockDay)
}

func (api *PublicSuperNodeAPI) TurnRegister(ctx context.Context, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	return contract_api.TurnRegisterSuperNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, lockID)
}

func (api *PublicSuperNodeAPI) ChangeAddress(ctx context.Context, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	return contract_api.ChangeSuperNodeAddress(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, newAddr)
}

func (api *PublicSuperNodeAPI) ChangeName(ctx context.Context, from common.Address, addr common.Address, name string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeName(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, name)
}

func (api *PublicSuperNodeAPI) ChangeEnode(ctx context.Context, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeEnode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, enode)
}

func (api *PublicSuperNodeAPI) ChangeDescription(ctx context.Context, from common.Address, addr common.Address, description string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeDescription(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, description)
}

func (api *PublicSuperNodeAPI) ChangeIsOfficial(ctx context.Context, from common.Address, addr common.Address, flag bool) (common.Hash, error) {
	return contract_api.ChangeSuperNodeIsOfficial(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, flag)
}

func (api *PublicSuperNodeAPI) GetInfo(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.SuperNodeInfo, error) {
	return contract_api.GetSuperNodeInfo(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetInfoByID(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SuperNodeInfo, error) {
	return contract_api.GetSuperNodeInfoByID(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetAll(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetAllSuperNodes(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetTops(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetTopSuperNodes(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetOfficials(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetOfficialSuperNodes(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetSuperNodeNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) Exist(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistSuperNode(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) ExistID(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistSuperNodeID(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) ExistName(ctx context.Context, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistSuperNodeName(ctx, api.blockChainAPI, name, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) ExistEnode(ctx context.Context, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistSuperNodeEnode(ctx, api.blockChainAPI, enode, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) ExistLockID(ctx context.Context, addr common.Address, lockID *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistSuperNodeLockID(ctx, api.blockChainAPI, addr, lockID, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) IsValid(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.IsValidSuperNode(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) IsFormal(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.IsFormalSuperNode(ctx, api.blockChainAPI, addr, blockNrOrHash)
}
