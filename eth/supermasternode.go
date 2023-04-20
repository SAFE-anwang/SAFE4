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
	"math/big"
	"strings"
)

type PublicSuperMasterNodeAPI struct {
	e *Ethereum
	blockChainAPI *ethapi.PublicBlockChainAPI
}

func NewPublicSuperMasterNodeAPI(e *Ethereum) *PublicSuperMasterNodeAPI {
	return &PublicSuperMasterNodeAPI{e, ethapi.NewPublicBlockChainAPI(e.APIBackend)}
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
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperMasterNodeABI))
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
		To: &systemcontracts.SuperMasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new(types.SuperMasterNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return info, nil
}

func (api *PublicSuperMasterNodeAPI) GetTop(ctx context.Context) ([]types.SuperMasterNodeInfo, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperMasterNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getTop"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperMasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	info := new([]types.SuperMasterNodeInfo)
	if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
		return nil, err
	}
	return *info, nil
}

func (api *PublicSuperMasterNodeAPI) GetNum(ctx context.Context) (*big.Int, error) {
	vABI, err := abi.JSON(strings.NewReader(systemcontracts.SuperMasterNodeABI))
	if err != nil {
		return nil, err
	}

	method := "getNum"
	data, err := vABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	args := ethapi.TransactionArgs{
		To: &systemcontracts.SuperMasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.blockChainAPI.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	count := new(big.Int)
	if err := vABI.UnpackIntoInterface(&count, method, result); err != nil {
		return nil, err
	}
	return count, nil
}