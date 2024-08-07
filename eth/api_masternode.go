package eth

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
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

	info, err := api.GetInfo(ctx, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
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

func (api *PublicMasterNodeAPI) Register(ctx context.Context, from common.Address, amount *hexutil.Big, isUnion bool, addr common.Address, lockDay *big.Int, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int) (common.Hash, error) {
	return contract_api.RegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, isUnion, addr, lockDay, enode, description, creatorIncentive, partnerIncentive)
}

func (api *PublicMasterNodeAPI) AppendRegister(ctx context.Context, from common.Address, amount *hexutil.Big, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	return contract_api.AppendRegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, addr, lockDay)
}

func (api *PublicMasterNodeAPI) TurnRegister(ctx context.Context, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	return contract_api.TurnRegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, lockID)
}

func (api *PublicMasterNodeAPI) ChangeAddress(ctx context.Context, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	return contract_api.ChangeMasterNodeAddress(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, newAddr)
}

func (api *PublicMasterNodeAPI) ChangeEnode(ctx context.Context, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	return contract_api.ChangeMasterNodeEnode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, enode)
}

func (api *PublicMasterNodeAPI) ChangeDescription(ctx context.Context, from common.Address, addr common.Address, description string) (common.Hash, error) {
	return contract_api.ChangeMasterNodeDescription(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, description)
}

func (api *PublicMasterNodeAPI) ChangeIsOfficial(ctx context.Context, from common.Address, addr common.Address, flag bool) (common.Hash, error) {
	return contract_api.ChangeMasterNodeIsOfficial(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, flag)
}

func (api *PublicMasterNodeAPI) GetInfo(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.MasterNodeInfo, error) {
	return contract_api.GetMasterNodeInfo(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) GetInfoByID(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.MasterNodeInfo, error) {
	return contract_api.GetMasterNodeInfoByID(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) GetNext(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (common.Address, error) {
	return contract_api.GetNextMasterNode(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) GetNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetMasterNodeNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) GetAll(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetAllMasterNodes(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) GetOfficials(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetOfficialMasterNodes(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) Exist(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistMasterNode(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) ExistID(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistMasterNodeID(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) ExistEnode(ctx context.Context, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistMasterNodeEnode(ctx, api.blockChainAPI, enode, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) ExistLockID(ctx context.Context, addr common.Address, lockID *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistMasterNodeLockID(ctx, api.blockChainAPI, addr, lockID, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) IsValid(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.IsValidMasterNode(ctx, api.blockChainAPI, addr, blockNrOrHash)
}