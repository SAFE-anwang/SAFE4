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

type PublicSuperMasterNodeAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
	transactionPoolAPI *ethapi.PublicTransactionPoolAPI
}

func NewPublicSuperMasterNodeAPI(e *Ethereum, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI) *PublicSuperMasterNodeAPI {
	return &PublicSuperMasterNodeAPI{e, blockChainAPI, transactionPoolAPI}
}

func (api *PublicSuperMasterNodeAPI) Start(ctx context.Context, addr common.Address) (bool, error) {
	info, err := api.GetInfo(ctx, addr)
	if err != nil {
		log.Info(err.Error())
		return false, nil
		api.blockChainAPI.BlockNumber()
	}
	api.e.eventMux.Post(core.NewSuperMasterNodeEvent{SuperMasterNodeInfo: info})

	return true, nil
}

func (api *PublicSuperMasterNodeAPI) Stop(ctx context.Context, addr common.Address) (bool, error) {
	log.Info("Stop supermasternode", "address", addr)
	return true, nil
}

func (api *PublicSuperMasterNodeAPI) Restart(ctx context.Context, addr common.Address) (bool, error) {
	log.Info("Restart supermasternode", "address", addr)
	return true, nil
}

func (api *PublicSuperMasterNodeAPI) GetInfo(ctx context.Context, addr common.Address) (*types.SuperMasterNodeInfo, error) {
	return systemcontracts.GetSuperMasterNodeInfo(ctx, api.blockChainAPI, addr)
}

func (api *PublicSuperMasterNodeAPI) GetAll(ctx context.Context) ([]types.SuperMasterNodeInfo, error) {
	return systemcontracts.GetAllSuperMasterNode(ctx, api.blockChainAPI)
}

func (api *PublicSuperMasterNodeAPI) GetTop(ctx context.Context) ([]types.SuperMasterNodeInfo, error) {
	return systemcontracts.GetTopSuperMasterNode(ctx, api.blockChainAPI)
}

func (api *PublicSuperMasterNodeAPI) GetNum(ctx context.Context) (*big.Int, error) {
	return systemcontracts.GetSuperMasterNodeNum(ctx, api.blockChainAPI)
}

func (api *PublicSuperMasterNodeAPI) RegisterSuperMasterNode(ctx context.Context, from common.Address, amount *big.Int, isUnion bool, smnAddr common.Address, lockDay *big.Int, name string, enode string, pubkey string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	return systemcontracts.RegisterSuperMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, isUnion, smnAddr, lockDay, name, enode, pubkey, description, creatorIncentive, partnerIncentive, voterIncentive)
}

func (api *PublicSuperMasterNodeAPI) AppendRegisterSuperMasterNode(ctx context.Context, from common.Address, amount *big.Int, smnAddr common.Address, lockDay *big.Int) (common.Hash, error) {
	return systemcontracts.AppendRegisterSuperMasterNode(ctx, api.blockChainAPI, api.transactionPoolAPI, from, amount, smnAddr, lockDay)
}