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

func RedeemAvailableSafe3(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
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

func RedeemLockedSafe3(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes, enode string) (common.Hash, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return common.Hash{}, err
    }

    method := "redeemLocked"
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

func ApplyRedeemSpecialSafe3(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, pubkey hexutil.Bytes, sig hexutil.Bytes) (common.Hash, error) {
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

func Vote4SpecialSafe3(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, safe3Addr string, voteResult *big.Int) (common.Hash, error) {
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

func GetAvailableSafe3(ctx context.Context, api *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.AvailableSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAvailable"
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

func GetLockedSafe3(ctx context.Context, api *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) ([]types.LockedSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getLocked"
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

    infos := new([]types.LockedSafe3Info)
    if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
        return nil, err
    }
    return *infos, nil
}

func GetSpecialSafe3(ctx context.Context, api *ethapi.PublicBlockChainAPI, safe3Addr string, blockNrOrHash rpc.BlockNumberOrHash) (*types.SpecialSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getSpecial"
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

func GetAllAvailableSafe3(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]types.AvailableSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAllAvailable"
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

    infos := new([]types.AvailableSafe3Info)
    if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
        return nil, err
    }
    return *infos, nil
}

func GetAllLockedSafe3(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]types.LockedSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAllLocked"
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

    infos := new([]types.LockedSafe3Info)
    if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
        return nil, err
    }
    return *infos, nil
}

func GetAllSpecialSafe3(ctx context.Context, api *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]types.SpecialSafe3Info, error) {
    vABI, err := abi.JSON(strings.NewReader(systemcontracts.Safe3ABI))
    if err != nil {
        return nil, err
    }

    method := "getAllSpecial"
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

    infos := new([]types.SpecialSafe3Info)
    if err := vABI.UnpackIntoInterface(infos, method, result); err != nil {
        return nil, err
    }
    return *infos, nil
}
