package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"strings"
)

type PublicMasterNodeAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
}

func NewPublicMasterNodeAPI(e *Ethereum) *PublicMasterNodeAPI {
	return &PublicMasterNodeAPI{e, ethapi.NewPublicBlockChainAPI(e.APIBackend)}
}

func (api *PublicMasterNodeAPI) Start(ctx context.Context, addr common.Address) (bool, error) {
	info, err := api.GetInfo(ctx, addr)
	if err != nil {
		return false, err
	}
	api.e.eventMux.Post(core.NewMasterNodeEvent{MasterNodeInfo: info, Operator: core.NODE_START})
	return true, nil
}

func (api *PublicMasterNodeAPI) Stop(ctx context.Context,addr common.Address) (bool, error) {
	log.Info("Stop masternode", "address", addr)
	return true, nil
}

func (api *PublicMasterNodeAPI) Restart(ctx context.Context, addr common.Address) (bool, error) {
	log.Info("Restart masternode", "address", addr)
	return true, nil
}

func (api *PublicMasterNodeAPI) GetInfo(ctx context.Context, addr common.Address) (*types.MasterNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getInfo"
	data, err := vABI.Pack(method, addr)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.MasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new(types.MasterNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func (api *PublicMasterNodeAPI) GetNext(ctx context.Context) (*common.Address, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.MasterNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getNext"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.MasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new(common.Address)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}