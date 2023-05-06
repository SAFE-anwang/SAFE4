package systemcontracts

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"strings"
)

func GetNextMasterNode(api *ethapi.PublicBlockChainAPI) (*common.Address, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if api == nil {
		return nil, errors.New("invalid blockchain api")
	}

	vABI, err := abi.JSON(strings.NewReader(MasterNodeABI))
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
		To: &MasterNodeContractAddr,
		Data: &msgData,
	}
	result, err := api.Call(ctx, args, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	addr := new(common.Address)
	if err := vABI.UnpackIntoInterface(&addr, method, result); err != nil {
		return nil, err
	}
	return addr, nil
}