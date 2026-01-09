package contract_api

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
)

func RegisterMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, isUnion bool, addr common.Address, lockDay *big.Int, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, value, systemcontracts.MasterNodeLogicContractAddr, "register", getValues(isUnion, addr, lockDay, enode, description, creatorIncentive, partnerIncentive))
}

func AppendRegisterMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, value, systemcontracts.MasterNodeLogicContractAddr, "appendRegister", getValues(addr, lockDay))
}

func TurnRegisterMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.MasterNodeLogicContractAddr, "turnRegister", getValues(addr, lockID))
}

func ChangeMasterNodeAddress(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.MasterNodeLogicContractAddr, "changeAddress", getValues(addr, newAddr))
}

func ChangeMasterNodeEnode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.MasterNodeLogicContractAddr, "changeEnode", getValues(addr, enode))
}

func ChangeMasterNodeEnodeByID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, enode string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.MasterNodeLogicContractAddr, "changeEnodeByID", getValues(id, enode))
}

func ChangeMasterNodeDescription(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, description string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.MasterNodeLogicContractAddr, "changeDescription", getValues(addr, description))
}

func ChangeMasterNodeDescriptionByID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, id *big.Int, description string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.MasterNodeLogicContractAddr, "changeDescriptionByID", getValues(id, description))
}

func GetMasterNodeInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.MasterNodeInfo, error) {
	ret := new(types.MasterNodeInfo)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getInfo", getValues(addr), &ret)
	return ret, err
}

func GetMasterNodeInfoByID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.MasterNodeInfo, error) {
	ret := new(types.MasterNodeInfo)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getInfoByID", getValues(id), &ret)
	return ret, err
}

func GetNextMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (common.Address, error) {
	ret := new(common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getNext", nil, ret); err != nil {
		return common.Address{}, err
	}
	return *ret, nil
}

func GetMasterNodeNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getNum", nil, &ret)
	return ret, err
}

func GetAllMasterNodes(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getAll", getValues(start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetMasterNodeNum4Creator(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, creator common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getAddrNum4Creator", getValues(creator), &ret)
	return ret, err
}

func GetMasterNodes4Creator(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, creator common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getAddrs4Creator", getValues(creator, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetMasterNodeNum4Partner(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, partner common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getAddrNum4Partner", getValues(partner), &ret)
	return ret, err
}

func GetMasterNodes4Partner(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, partner common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getAddrs4Partner", getValues(partner, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetOfficialMasterNodes(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "getOfficials", nil, &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func ExistMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "exist", getValues(addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistMasterNodeID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "existID", getValues(id), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistMasterNodeEnode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "existEnode", getValues(CompressEnode(enode)), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistMasterNodeLockID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, lockID *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "existLockID", getValues(addr, lockID), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func IsValidMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "isValid", getValues(addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func IsUnionMasterNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "isUnion", getValues(addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistNodeAddress(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "existNodeAddress", getValues(addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistNodeEnode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.MasterNodeStorageContractAddr, "existNodeEnode", getValues(CompressEnode(enode)), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}
