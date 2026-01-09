package eth

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts/contract_api"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type PrivateMasterNodeAPI struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPrivateMasterNodeAPI(e *Ethereum) *PrivateMasterNodeAPI {
	return &PrivateMasterNodeAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PrivateMasterNodeAPI) Start(ctx context.Context, addr common.Address) (bool, error) {
	progress := api.e.APIBackend.SyncProgress()
	if progress.CurrentBlock < progress.HighestBlock {
		return false, errors.New("syncing now")
	}

	info, err := contract_api.GetMasterNodeInfo(ctx, api.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return false, err
	}
	enode := contract_api.CompressEnode(api.e.p2pServer.NodeInfo().Enode)
	if !contract_api.CompareEnode(enode, info.Enode) {
		return false, errors.New("start failed, incompatible masternode enode, local: [" + enode + "], state: [" + info.Enode + "]")
	}

	curBlock := api.e.blockchain.CurrentBlock()
	ping, err := types.NewNodePing(info.Id, types.MasterNodeType, curBlock.Hash(), curBlock.Number(), api.e.p2pServer.Config.PrivateKey)
	if err != nil {
		return false, err
	}
	api.e.handler.BroadcastNodePing(ping)
	return true, nil
}

func (api *PrivateMasterNodeAPI) Register(ctx context.Context, from common.Address, value *hexutil.Big, isUnion bool, addr common.Address, lockDay *big.Int, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int) (common.Hash, error) {
	return contract_api.RegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, value, isUnion, addr, lockDay, enode, description, creatorIncentive, partnerIncentive)
}

func (api *PrivateMasterNodeAPI) AppendRegister(ctx context.Context, from common.Address, value *hexutil.Big, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	return contract_api.AppendRegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, value, addr, lockDay)
}

func (api *PrivateMasterNodeAPI) TurnRegister(ctx context.Context, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	return contract_api.TurnRegisterMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, lockID)
}

func (api *PrivateMasterNodeAPI) ChangeAddress(ctx context.Context, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	return contract_api.ChangeMasterNodeAddress(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, newAddr)
}

func (api *PrivateMasterNodeAPI) ChangeEnode(ctx context.Context, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	return contract_api.ChangeMasterNodeEnode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, enode)
}

func (api *PrivateMasterNodeAPI) ChangeEnodeByID(ctx context.Context, from common.Address, id *big.Int, enode string) (common.Hash, error) {
	return contract_api.ChangeMasterNodeEnodeByID(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, enode)
}

func (api *PrivateMasterNodeAPI) ChangeDescription(ctx context.Context, from common.Address, addr common.Address, description string) (common.Hash, error) {
	return contract_api.ChangeMasterNodeDescription(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, description)
}

func (api *PrivateMasterNodeAPI) ChangeDescriptionByID(ctx context.Context, from common.Address, id *big.Int, description string) (common.Hash, error) {
	return contract_api.ChangeMasterNodeDescriptionByID(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, description)
}

type PublicMasterNodeAPI struct {
	e             *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
}

func NewPublicMasterNodeAPI(e *Ethereum) *PublicMasterNodeAPI {
	return &PublicMasterNodeAPI{e, e.GetPublicBlockChainAPI()}
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

func (api *PublicMasterNodeAPI) GetAddrNum4Creator(ctx context.Context, creator common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetMasterNodeNum4Creator(ctx, api.blockChainAPI, creator, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) GetAddrs4Creator(ctx context.Context, creator common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetMasterNodes4Creator(ctx, api.blockChainAPI, creator, start, count, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) GetAddrNum4Partner(ctx context.Context, partner common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetMasterNodeNum4Partner(ctx, api.blockChainAPI, partner, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) GetAddrs4Partner(ctx context.Context, partner common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetMasterNodes4Partner(ctx, api.blockChainAPI, partner, start, count, blockNrOrHash)
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

func (api *PublicMasterNodeAPI) IsUnion(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.IsUnionMasterNode(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) ExistNodeAddress(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistNodeAddress(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicMasterNodeAPI) ExistNodeEnode(ctx context.Context, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistNodeEnode(ctx, api.blockChainAPI, enode, blockNrOrHash)
}
