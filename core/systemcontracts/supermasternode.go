package systemcontracts

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"strings"
)

func GetTopSuperMasterNode(api *ethapi.PublicBlockChainAPI) ([]types.SuperMasterNodeInfo, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if api == nil {
		return nil, errors.New("invalid blockchain api");
	}

	vABI, err := abi.JSON(strings.NewReader(SuperMasterNodeABI))
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
		To: &SuperMasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	var (
		ret0 = new([]types.SuperMasterNodeInfo)
	)
	out := ret0
	if err := vABI.UnpackIntoInterface(out, method, result); err != nil {
		return nil, err
	}

	smnList := make([]types.SuperMasterNodeInfo, len(*ret0))
	for i, addr := range *ret0 {
		smnList[i] = addr
	}
	return smnList, nil
}
