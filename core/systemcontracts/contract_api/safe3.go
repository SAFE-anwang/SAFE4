package contract_api

import (
    "context"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/core/systemcontracts"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/internal/ethapi"
    "github.com/ethereum/go-ethereum/rpc"
    "math/big"
    "strings"
)

func RedeemAvailable(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return common.Hash{}, err
    }

    method := "redeemAvailable"
    data, err := vABI.Pack(method, pubkey, sig)
    if err != nil {
        return common.Hash{}, err
    }

    msgData := (hexutil.Bytes)(data)
    value := big.NewInt(1000000000000000000) // 1 SAFE
    args := ethapi.TransactionArgs{
        From:     &from,
        To:       &systemcontracts.Safe3ContractAddr,
        Data:     &msgData,
        Value:    (*hexutil.Big)(value),
        GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
    }
    gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
    if err != nil {
        return common.Hash{}, err
    }
    args.Gas = &gas
    return transactionPoolAPI.SendTransaction(ctx, args)
}

func RedeemLocked(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return common.Hash{}, err
    }

    method := "redeemLocked"
    data, err := vABI.Pack(method, pubkey, sig)
    if err != nil {
        return common.Hash{}, err
    }

    msgData := (hexutil.Bytes)(data)
    value := big.NewInt(1000000000000000000) // 1 SAFE
    args := ethapi.TransactionArgs{
        From:     &from,
        To:       &systemcontracts.Safe3ContractAddr,
        Data:     &msgData,
        Value:    (*hexutil.Big)(value),
        GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
    }
    gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
    if err != nil {
        return common.Hash{}, err
    }
    args.Gas = &gas
    return transactionPoolAPI.SendTransaction(ctx, args)
}

func RedeemMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes, enode string) (common.Hash, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return common.Hash{}, err
    }

    method := "redeemMasterNode"
    data, err := vABI.Pack(method, pubkey, sig, enode)
    if err != nil {
        return common.Hash{}, err
    }

    msgData := (hexutil.Bytes)(data)
    value := big.NewInt(1000000000000000000) // 1 SAFE
    args := ethapi.TransactionArgs{
        From:     &from,
        To:       &systemcontracts.Safe3ContractAddr,
        Data:     &msgData,
        Value:    (*hexutil.Big)(value),
        GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
    }
    gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
    if err != nil {
        return common.Hash{}, err
    }
    args.Gas = &gas
    return transactionPoolAPI.SendTransaction(ctx, args)
}

func ApplyRedeemSpecial(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return common.Hash{}, err
    }

    method := "applyRedeemSpecial"
    data, err := vABI.Pack(method, pubkey, sig)
    if err != nil {
        return common.Hash{}, err
    }

    msgData := (hexutil.Bytes)(data)
    value := big.NewInt(1000000000000000000) // 1 SAFE
    args := ethapi.TransactionArgs{
        From:     &from,
        To:       &systemcontracts.Safe3ContractAddr,
        Data:     &msgData,
        Value:    (*hexutil.Big)(value),
        GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
    }
    gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
    if err != nil {
        return common.Hash{}, err
    }
    args.Gas = &gas
    return transactionPoolAPI.SendTransaction(ctx, args)
}

func Vote4Special(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, safe3Addr string, voteResult *big.Int) (common.Hash, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return common.Hash{}, err
    }

    method := "applyRedeemSpecial"
    data, err := vABI.Pack(method, safe3Addr, voteResult)
    if err != nil {
        return common.Hash{}, err
    }

    msgData := (hexutil.Bytes)(data)
    value := big.NewInt(1000000000000000000) // 1 SAFE
    args := ethapi.TransactionArgs{
        From:     &from,
        To:       &systemcontracts.Safe3ContractAddr,
        Data:     &msgData,
        Value:    (*hexutil.Big)(value),
        GasPrice: (*hexutil.Big)(GetCurrentGasPrice(ctx, blockChainAPI)),
    }
    gas, err := blockChainAPI.EstimateGas(ctx, args, nil)
    if err != nil {
        return common.Hash{}, err
    }
    args.Gas = &gas
    return transactionPoolAPI.SendTransaction(ctx, args)
}

func GetAllAvailableNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAllAvailableNum"
    data, err := vABI.Pack(method)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    info := new(big.Int)
    if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
        return nil, err
    }
    return info, nil
}

func GetAvailableInfos(ctx context.Context, api *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.AvailableSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAvailableInfos"
    data, err := vABI.Pack(method, start, count)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    infos := new([]types.AvailableSafe3Info)
    if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
        return nil, err
    }
    return *infos, nil
}

func GetAvailableInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.AvailableSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAvailableInfo"
    data, err := vABI.Pack(method, safe3Addr)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    info := new(types.AvailableSafe3Info)
    if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
        return nil, err
    }
    return info, nil
}

func GetAllLockedNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAllLockedNum"
    data, err := vABI.Pack(method)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    info := new(big.Int)
    if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
        return nil, err
    }
    return info, nil
}

func GetLockedAddrNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getLockedAddrNum"
    data, err := vABI.Pack(method)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    info := new(big.Int)
    if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
        return nil, err
    }
    return info, nil
}

func GetLockedAddrs(ctx context.Context, api *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]string, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getLockedAddrs"
    data, err := vABI.Pack(method, start, count)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    infos := new([]string)
    if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
        return nil, err
    }
    return *infos, nil
}

func GetLockedNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getLockedNum"
    data, err := vABI.Pack(method, safe3Addr)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    info := new(big.Int)
    if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
        return nil, err
    }
    return info, nil
}

func GetLockedInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, safe3Addr string, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.LockedSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getLockedInfo"
    data, err := vABI.Pack(method, safe3Addr, start, count)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    infos := new([]types.LockedSafe3Info)
    if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
        return nil, err
    }
    return *infos, nil
}

func GetAllSpecialNum(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAllSpecialNum"
    data, err := vABI.Pack(method)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    info := new(big.Int)
    if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
        return nil, err
    }
    return info, nil
}

func GetSpecialInfos(ctx context.Context, api *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]types.SpecialSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getSpecialInfos"
    data, err := vABI.Pack(method, start, count)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    infos := new([]types.SpecialSafe3Info)
    if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
        return nil, err
    }
    return *infos, nil
}

func GetSpecialInfo(ctx context.Context, api *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.SpecialSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getSpecialInfo"
    data, err := vABI.Pack(method, safe3Addr)
    if err != nil {
        return nil, err
    }

    msgData := (hexutil.Bytes)(data)
    args := ethapi.TransactionArgs{
        To: &systemcontracts.Safe3ContractAddr,
        Data: &msgData,
    }
    result, err := api.Call(ctx, args, blockNrOrHash, nil)

    if err != nil {
        return nil, err
    }

    info := new(types.SpecialSafe3Info)
    if err := vABI.UnpackIntoInterface(&info, method, result); err != nil {
        return nil, err
    }
    return info, nil
}
