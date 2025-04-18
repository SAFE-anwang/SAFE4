package contract_api

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	eth_params "github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

func getValues(args ...interface{}) []reflect.Value {
	var values []reflect.Value
	for _, arg := range args {
		values = append(values, reflect.ValueOf(arg))
	}
	return values
}

func getABI(contractAddr common.Address) (abi.ABI, error) {
	contractABI := ""
	switch contractAddr {
	case systemcontracts.PropertyContractAddr:
		contractABI = systemcontracts.PropertyABI
	case systemcontracts.AccountManagerContractAddr:
		contractABI = systemcontracts.AccountManagerABI
	case systemcontracts.MasterNodeStorageContractAddr:
		contractABI = systemcontracts.MasterNodeStorageABI
	case systemcontracts.MasterNodeLogicContractAddr:
		contractABI = systemcontracts.MasterNodeLogicABI
	case systemcontracts.SuperNodeStorageContractAddr:
		contractABI = systemcontracts.SuperNodeStorageABI
	case systemcontracts.SuperNodeLogicContractAddr:
		contractABI = systemcontracts.SuperNodeLogicABI
	case systemcontracts.SNVoteContractAddr:
		contractABI = systemcontracts.SNVoteABI
	case systemcontracts.MasterNodeStateContractAddr:
		contractABI = systemcontracts.MasterNodeStateABI
	case systemcontracts.SuperNodeStateContractAddr:
		contractABI = systemcontracts.SuperNodeStateABI
	case systemcontracts.ProposalContractAddr:
		contractABI = systemcontracts.ProposalABI
	case systemcontracts.Safe3ContractAddr:
		contractABI = systemcontracts.Safe3ABI
	default:
		return abi.ABI{}, fmt.Errorf("unknown contract")
	}
	return abi.JSON(strings.NewReader(contractABI))
}

func CallContract(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, contractAddr common.Address, method string, inputs []reflect.Value) (common.Hash, error) {
	vABI, err := getABI(contractAddr)
	if err != nil {
		return common.Hash{}, err
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf(method))
	if inputs != nil {
		params = append(params, inputs...)
	}

	f := reflect.ValueOf(vABI.Pack)
	result := f.Call(params)
	if result[1].Interface() != nil {
		err = result[1].Interface().(error)
		return common.Hash{}, err
	}
	data := result[0].Interface().([]byte)

	msgData := (hexutil.Bytes)(data)

	var args ethapi.TransactionArgs
	if contractAddr != systemcontracts.MasterNodeStateContractAddr && contractAddr != systemcontracts.SuperNodeStateContractAddr {
		args = ethapi.TransactionArgs{
			From:     &from,
			To:       &contractAddr,
			Data:     &msgData,
			GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
		}
		if value != nil {
			args.Value = value
		}
		gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
		if err != nil {
			return common.Hash{}, err
		}
		gas = gas * 6 / 5
		args.Gas = &gas
	} else {
		gas := eth_params.MaxSystemRewardTxGas
		args = ethapi.TransactionArgs{
			From:     &from,
			To:       &contractAddr,
			Data:     &msgData,
			Gas:      (*hexutil.Uint64)(&gas),
			GasPrice: (*hexutil.Big)(common.Big0),
		}
	}
	return transactionPoolAPI.SendTransaction(ctx, args)
}

func QueryContract(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash, contractAddr common.Address, method string, inputs []reflect.Value, output interface{}) error {
	vABI, err := getABI(contractAddr)
	if err != nil {
		return err
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf(method))
	if inputs != nil {
		params = append(params, inputs...)
	}

	f := reflect.ValueOf(vABI.Pack)
	result := f.Call(params)
	if result[1].Interface() != nil {
		err = result[1].Interface().(error)
		return err
	}
	data := result[0].Interface().([]byte)
	args := ethapi.TransactionArgs{
		To:   &contractAddr,
		Data: (*hexutil.Bytes)(&data),
	}
	buf, err := blockChainAPI.Call(ctx, args, blockNrOrHash, nil)
	if err != nil {
		return err
	}
	return vABI.UnpackIntoInterface(output, method, buf)
}

func QueryContract4MultiReturn(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash, contractAddr common.Address, method string, inputs []reflect.Value) ([]interface{}, error) {
	vABI, err := getABI(contractAddr)
	if err != nil {
		return nil, err
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf(method))
	params = append(params, inputs...)

	f := reflect.ValueOf(vABI.Pack)
	result := f.Call(params)
	data := result[0].Interface().([]byte)
	if result[1].Interface() != nil {
		err = result[1].Interface().(error)
		return nil, err
	}

	args := ethapi.TransactionArgs{
		To:   &contractAddr,
		Data: (*hexutil.Bytes)(&data),
	}
	buf, err := blockChainAPI.Call(ctx, args, blockNrOrHash, nil)
	if err != nil {
		return nil, err
	}
	return vABI.Unpack(method, buf)
}

func GetCurrentGasPrice(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI) *big.Int {
	gasPrice, err := GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
	if err != nil || gasPrice.Int64() == 0 {
		gasPrice = big.NewInt(eth_params.GWei / 100)
	}
	return gasPrice
}

func GetLatestGasPrice(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI) *big.Int {
	gasPrice, err := GetPropertyValue(ctx, blockChainAPI, "gas_price", rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil || gasPrice.Int64() == 0 {
		gasPrice = big.NewInt(eth_params.GWei / 100)
	}
	return gasPrice
}

func CompressEnode(enode string) string {
	arr := strings.Split(enode, "?")
	if len(arr) == 0 {
		return enode
	}
	return arr[0]
}

func CompareEnode(e1, e2 string) bool {
	return CompressEnode(e1) == CompressEnode(e2)
}
