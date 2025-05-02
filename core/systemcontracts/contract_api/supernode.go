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

func RegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, isUnion bool, addr common.Address, lockDay *big.Int, name string, enode string, description string, creatorIncentive *big.Int, partnerIncentive *big.Int, voterIncentive *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, value, systemcontracts.SuperNodeLogicContractAddr, "register", getValues(isUnion, addr, lockDay, name, enode, description, creatorIncentive, partnerIncentive, voterIncentive))
}

func AppendRegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, value *hexutil.Big, addr common.Address, lockDay *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, value, systemcontracts.SuperNodeLogicContractAddr, "appendRegister", getValues(addr, lockDay))
}

func TurnRegisterSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, lockID *big.Int) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SuperNodeLogicContractAddr, "turnRegister", getValues(addr, lockID))
}

func ChangeSuperNodeAddress(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, newAddr common.Address) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SuperNodeLogicContractAddr, "changeAddress", getValues(addr, newAddr))
}

func ChangeSuperNodeName(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, name string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SuperNodeLogicContractAddr, "changeName", getValues(addr, name))
}

func ChangeSuperNodeEnode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, enode string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SuperNodeLogicContractAddr, "changeEnode", getValues(addr, enode))
}

func ChangeSuperNodeDescription(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, description string) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SuperNodeLogicContractAddr, "changeDescription", getValues(addr, description))
}

func ChangeSuperNodeIsOfficial(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, transactionPoolAPI *ethapi.PublicTransactionPoolAPI, from common.Address, addr common.Address, flag bool) (common.Hash, error) {
	return CallContract(ctx, blockChainAPI, transactionPoolAPI, from, nil, systemcontracts.SuperNodeLogicContractAddr, "changeIsOfficial", getValues(addr, flag))
}

func GetSuperNodeInfo(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*types.SuperNodeInfo, error) {
	ret := new(types.SuperNodeInfo)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getInfo", getValues(addr), &ret)
	return ret, err
}

func GetSuperNodeInfoByID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (*types.SuperNodeInfo, error) {
	ret := new(types.SuperNodeInfo)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getInfoByID", getValues(id), &ret)
	return ret, err
}

func GetSuperNodeNum(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getNum", nil, &ret)
	return ret, err
}

func GetAllSuperNodes(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getAll", getValues(start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetSuperNodeNum4Creator(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, creator common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getAddrs4Creator", getValues(creator), &ret)
	return ret, err
}

func GetSuperNodes4Creator(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, creator common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getAddrs4Creator", getValues(creator, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetSuperNodeNum4Partner(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, partner common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	ret := new(big.Int)
	err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getAddrNum4Partner", getValues(partner), &ret)
	return ret, err
}

func GetSuperNodes4Partner(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, partner common.Address, start *big.Int, count *big.Int, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getAddrs4Partner", getValues(partner, start, count), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetTopSuperNodes(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getTops", nil, &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetTopSuperNodes4Creator(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, creator common.Address, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getTops4Creator", getValues(creator), &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func GetOfficialSuperNodes(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, blockNrOrHash rpc.BlockNumberOrHash) ([]common.Address, error) {
	ret := new([]common.Address)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "getOfficials", nil, &ret); err != nil {
		return nil, err
	}
	return *ret, nil
}

func ExistSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "exist", getValues(addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistSuperNodeID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, id *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "existID", getValues(id), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistSuperNodeName(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, name string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "existName", getValues(name), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistSuperNodeEnode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, enode string, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "existEnode", getValues(CompressEnode(enode)), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistSuperNodeLockID(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, lockID *big.Int, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "existLockID", getValues(addr, lockID), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func ExistSuperNodeFounder(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, founder common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "existFounder", getValues(founder), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func IsValidSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "isValid", getValues(addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}

func IsFormalSuperNode(ctx context.Context, blockChainAPI *ethapi.PublicBlockChainAPI, addr common.Address, blockNrOrHash rpc.BlockNumberOrHash) (bool, error) {
	ret := new(bool)
	if err := QueryContract(ctx, blockChainAPI, blockNrOrHash, systemcontracts.SuperNodeStorageContractAddr, "isFormal", getValues(addr), &ret); err != nil {
		return false, err
	}
	return *ret, nil
}
