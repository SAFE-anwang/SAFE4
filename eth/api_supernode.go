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

type PrivateSuperNodeAPI struct {
	e                  *Ethereum
	blockChainAPI      *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPrivateSuperNodeAPI(e *Ethereum) *PrivateSuperNodeAPI {
	return &PrivateSuperNodeAPI{e, e.GetPublicBlockChainAPI(), e.GetPublicTransactionPoolAPI()}
}

func (api *PrivateSuperNodeAPI) Start(ctx context.Context, addr common.Address) (bool, error) {
	progress := api.e.APIBackend.SyncProgress()
	if progress.CurrentBlock < progress.HighestBlock {
		return false, errors.New("syncing now")
	}

	info, err := contract_api.GetSuperNodeInfo(ctx, api.blockChainAPI, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return false, err
	}
	enode := contract_api.CompressEnode(api.e.p2pServer.NodeInfo().Enode)
	if !contract_api.CompareEnode(enode, info.Enode) {
		return false, errors.New("start failed, incompatible supernode enode, local: [" + enode + "], state: [" + info.Enode + "]")
	}

	curBlock := api.e.blockchain.CurrentBlock()
	ping, err := types.NewNodePing(info.Id, types.SuperNodeType, curBlock.Hash(), curBlock.Number(), api.e.p2pServer.Config.PrivateKey)
	if err != nil {
		return false, err
	}
	api.e.handler.BroadcastNodePing(ping)
	return true, nil
}

func (api *PrivateSuperNodeAPI) Register(ctx context.Context, from common.Address, value *hexutil.Big, isUnion bool, addr common.Address, lockDay *big.Int, name string, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	return contract_api.RegisterSuperNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, value, isUnion, addr, lockDay, name, enode, description, creatorIncentive, partnerIncentive, voterIncentive)
}

func (api *PrivateSuperNodeAPI) AppendRegister(ctx context.Context, from common.Address, value *hexutil.Big, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	return contract_api.AppendRegisterSuperNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, value, addr, lockDay)
}

func (api *PrivateSuperNodeAPI) TurnRegister(ctx context.Context, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	return contract_api.TurnRegisterSuperNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, lockID)
}

func (api *PrivateSuperNodeAPI) ChangeAddress(ctx context.Context, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	return contract_api.ChangeSuperNodeAddress(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, newAddr)
}

func (api *PrivateSuperNodeAPI) ChangeName(ctx context.Context, from common.Address, addr common.Address, name string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeName(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, name)
}

func (api *PrivateSuperNodeAPI) ChangeNameByID(ctx context.Context, from common.Address, id *big.Int, name string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeNameByID(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, name)
}

func (api *PrivateSuperNodeAPI) ChangeEnode(ctx context.Context, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeEnode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, enode)
}

func (api *PrivateSuperNodeAPI) ChangeEnodeByID(ctx context.Context, from common.Address, id *big.Int, enode string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeEnodeByID(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, enode)
}

func (api *PrivateSuperNodeAPI) ChangeDescription(ctx context.Context, from common.Address, addr common.Address, description string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeDescription(ctx, api.blockChainAPI, api.transactionPoolAPI, from, addr, description)
}

func (api *PrivateSuperNodeAPI) ChangeDescriptionByID(ctx context.Context, from common.Address, id *big.Int, description string) (common.Hash, error) {
	return contract_api.ChangeSuperNodeDescriptionByID(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, description)
}

func (api *PrivateSuperNodeAPI) ChangeIncentivePlan(ctx context.Context, from common.Address, id *big.Int, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	return contract_api.ChangeSuperNodeIncentivePlan(ctx, api.blockChainAPI, api.transactionPoolAPI, from, id, creatorIncentive, partnerIncentive, voterIncentive)
}

type PublicSuperNodeAPI struct {
	e             *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
}

func NewPublicSuperNodeAPI(e *Ethereum) *PublicSuperNodeAPI {
	return &PublicSuperNodeAPI{e, e.GetPublicBlockChainAPI()}
}

func (api *PublicSuperNodeAPI) GetInfo(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.SuperNodeInfo, error) {
	return contract_api.GetSuperNodeInfo(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetInfoByID(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SuperNodeInfo, error) {
	return contract_api.GetSuperNodeInfoByID(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetDisableHeight(ctx context.Context, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetSuperNodeDisableHeight(ctx, api.blockChainAPI, id, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetNum(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetSuperNodeNum(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetAll(ctx context.Context, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetAllSuperNodes(ctx, api.blockChainAPI, start, count, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetAddrNum4Creator(ctx context.Context, creator common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetSuperNodeNum4Creator(ctx, api.blockChainAPI, creator, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetAddrs4Creator(ctx context.Context, creator common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetSuperNodes4Creator(ctx, api.blockChainAPI, creator, start, count, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetAddrNum4Partner(ctx context.Context, partner common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return contract_api.GetSuperNodeNum4Partner(ctx, api.blockChainAPI, partner, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetAddrs4Partner(ctx context.Context, partner common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetSuperNodes4Partner(ctx, api.blockChainAPI, partner, start, count, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetTops(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetTopSuperNodes(ctx, api.blockChainAPI, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetTops4Creator(ctx context.Context, creator common.Address, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetTopSuperNodes4Creator(ctx, api.blockChainAPI, creator, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) GetOfficials(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	return contract_api.GetOfficialSuperNodes(ctx, api.blockChainAPI, blockNrOrHash)
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

func (api *PublicSuperNodeAPI) ExistFounder(ctx context.Context, founder common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistSuperNodeFounder(ctx, api.blockChainAPI, founder, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) IsValid(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.IsValidSuperNode(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) IsFormal(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.IsFormalSuperNode(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) ExistNodeAddress(ctx context.Context, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistNodeAddress(ctx, api.blockChainAPI, addr, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) ExistNodeEnode(ctx context.Context, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistNodeEnode(ctx, api.blockChainAPI, enode, blockNrOrHash)
}

func (api *PublicSuperNodeAPI) ExistNodeFounder(ctx context.Context, founder common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	return contract_api.ExistNodeFounder(ctx, api.blockChainAPI, founder, blockNrOrHash)
}
