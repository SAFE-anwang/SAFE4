// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MasterNode

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IMasterNodeIncentivePlan is an auto generated low-level Go binding around an user-defined struct.
type IMasterNodeIncentivePlan struct {
	Creator *big.Int
	Partner *big.Int
	Voter   *big.Int
}

// IMasterNodeMasterNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type IMasterNodeMasterNodeInfo struct {
	Id            *big.Int
	Addr          common.Address
	Creator       common.Address
	Amount        *big.Int
	Ip            string
	Pubkey        string
	Description   string
	State         *big.Int
	Founders      []IMasterNodeMemberInfo
	IncentivePlan IMasterNodeIncentivePlan
	CreateHeight  *big.Int
	UpdateHeight  *big.Int
}

// IMasterNodeMemberInfo is an auto generated low-level Go binding around an user-defined struct.
type IMasterNodeMemberInfo struct {
	LockID [20]byte
	Addr   common.Address
	Amount *big.Int
	Height *big.Int
}

// MasterNodeMetaData contains all meta data concerning the MasterNode contract.
var MasterNodeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"}],\"name\":\"MNAppendRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_reocrdID\",\"type\":\"bytes20\"}],\"name\":\"MNRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"}],\"name\":\"appendRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAddr\",\"type\":\"address\"}],\"name\":\"changeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"changeDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"}],\"name\":\"changeIP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"changePubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"exist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"existID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"}],\"name\":\"existIP\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_lokcID\",\"type\":\"bytes20\"}],\"name\":\"existLockID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"existPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes20\",\"name\":\"lockID\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structIMasterNode.MemberInfo[]\",\"name\":\"founders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"partner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voter\",\"type\":\"uint256\"}],\"internalType\":\"structIMasterNode.IncentivePlan\",\"name\":\"incentivePlan\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIMasterNode.MasterNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNext\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isUnion\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_creatorIncentive\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_partnerIncentive\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MasterNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterNodeMetaData.ABI instead.
var MasterNodeABI = MasterNodeMetaData.ABI

// MasterNode is an auto generated Go binding around an Ethereum contract.
type MasterNode struct {
	MasterNodeCaller     // Read-only binding to the contract
	MasterNodeTransactor // Write-only binding to the contract
	MasterNodeFilterer   // Log filterer for contract events
}

// MasterNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterNodeSession struct {
	Contract     *MasterNode       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterNodeCallerSession struct {
	Contract *MasterNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MasterNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterNodeTransactorSession struct {
	Contract     *MasterNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MasterNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterNodeRaw struct {
	Contract *MasterNode // Generic contract binding to access the raw methods on
}

// MasterNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterNodeCallerRaw struct {
	Contract *MasterNodeCaller // Generic read-only contract binding to access the raw methods on
}

// MasterNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterNodeTransactorRaw struct {
	Contract *MasterNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterNode creates a new instance of MasterNode, bound to a specific deployed contract.
func NewMasterNode(address common.Address, backend bind.ContractBackend) (*MasterNode, error) {
	contract, err := bindMasterNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterNode{MasterNodeCaller: MasterNodeCaller{contract: contract}, MasterNodeTransactor: MasterNodeTransactor{contract: contract}, MasterNodeFilterer: MasterNodeFilterer{contract: contract}}, nil
}

// NewMasterNodeCaller creates a new read-only instance of MasterNode, bound to a specific deployed contract.
func NewMasterNodeCaller(address common.Address, caller bind.ContractCaller) (*MasterNodeCaller, error) {
	contract, err := bindMasterNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterNodeCaller{contract: contract}, nil
}

// NewMasterNodeTransactor creates a new write-only instance of MasterNode, bound to a specific deployed contract.
func NewMasterNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterNodeTransactor, error) {
	contract, err := bindMasterNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterNodeTransactor{contract: contract}, nil
}

// NewMasterNodeFilterer creates a new log filterer instance of MasterNode, bound to a specific deployed contract.
func NewMasterNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterNodeFilterer, error) {
	contract, err := bindMasterNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterNodeFilterer{contract: contract}, nil
}

// bindMasterNode binds a generic wrapper to an already deployed contract.
func bindMasterNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasterNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterNode *MasterNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterNode.Contract.MasterNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterNode *MasterNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterNode.Contract.MasterNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterNode *MasterNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterNode.Contract.MasterNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterNode *MasterNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterNode *MasterNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterNode *MasterNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterNode.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_MasterNode *MasterNodeCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_MasterNode *MasterNodeSession) GetInitializeData() ([]byte, error) {
	return _MasterNode.Contract.GetInitializeData(&_MasterNode.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_MasterNode *MasterNodeCallerSession) GetInitializeData() ([]byte, error) {
	return _MasterNode.Contract.GetInitializeData(&_MasterNode.CallOpts)
}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _addr) view returns(bool)
func (_MasterNode *MasterNodeCaller) Exist(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "exist", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _addr) view returns(bool)
func (_MasterNode *MasterNodeSession) Exist(_addr common.Address) (bool, error) {
	return _MasterNode.Contract.Exist(&_MasterNode.CallOpts, _addr)
}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _addr) view returns(bool)
func (_MasterNode *MasterNodeCallerSession) Exist(_addr common.Address) (bool, error) {
	return _MasterNode.Contract.Exist(&_MasterNode.CallOpts, _addr)
}

// ExistID is a free data retrieval call binding the contract method 0x829c82e1.
//
// Solidity: function existID(uint256 _id) view returns(bool)
func (_MasterNode *MasterNodeCaller) ExistID(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "existID", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistID is a free data retrieval call binding the contract method 0x829c82e1.
//
// Solidity: function existID(uint256 _id) view returns(bool)
func (_MasterNode *MasterNodeSession) ExistID(_id *big.Int) (bool, error) {
	return _MasterNode.Contract.ExistID(&_MasterNode.CallOpts, _id)
}

// ExistID is a free data retrieval call binding the contract method 0x829c82e1.
//
// Solidity: function existID(uint256 _id) view returns(bool)
func (_MasterNode *MasterNodeCallerSession) ExistID(_id *big.Int) (bool, error) {
	return _MasterNode.Contract.ExistID(&_MasterNode.CallOpts, _id)
}

// ExistIP is a free data retrieval call binding the contract method 0xde1b1428.
//
// Solidity: function existIP(string _ip) view returns(bool)
func (_MasterNode *MasterNodeCaller) ExistIP(opts *bind.CallOpts, _ip string) (bool, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "existIP", _ip)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistIP is a free data retrieval call binding the contract method 0xde1b1428.
//
// Solidity: function existIP(string _ip) view returns(bool)
func (_MasterNode *MasterNodeSession) ExistIP(_ip string) (bool, error) {
	return _MasterNode.Contract.ExistIP(&_MasterNode.CallOpts, _ip)
}

// ExistIP is a free data retrieval call binding the contract method 0xde1b1428.
//
// Solidity: function existIP(string _ip) view returns(bool)
func (_MasterNode *MasterNodeCallerSession) ExistIP(_ip string) (bool, error) {
	return _MasterNode.Contract.ExistIP(&_MasterNode.CallOpts, _ip)
}

// ExistLockID is a free data retrieval call binding the contract method 0x598bdb3d.
//
// Solidity: function existLockID(address _addr, bytes20 _lokcID) view returns(bool)
func (_MasterNode *MasterNodeCaller) ExistLockID(opts *bind.CallOpts, _addr common.Address, _lokcID [20]byte) (bool, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "existLockID", _addr, _lokcID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistLockID is a free data retrieval call binding the contract method 0x598bdb3d.
//
// Solidity: function existLockID(address _addr, bytes20 _lokcID) view returns(bool)
func (_MasterNode *MasterNodeSession) ExistLockID(_addr common.Address, _lokcID [20]byte) (bool, error) {
	return _MasterNode.Contract.ExistLockID(&_MasterNode.CallOpts, _addr, _lokcID)
}

// ExistLockID is a free data retrieval call binding the contract method 0x598bdb3d.
//
// Solidity: function existLockID(address _addr, bytes20 _lokcID) view returns(bool)
func (_MasterNode *MasterNodeCallerSession) ExistLockID(_addr common.Address, _lokcID [20]byte) (bool, error) {
	return _MasterNode.Contract.ExistLockID(&_MasterNode.CallOpts, _addr, _lokcID)
}

// ExistPubkey is a free data retrieval call binding the contract method 0x4afe2bfb.
//
// Solidity: function existPubkey(string _pubkey) view returns(bool)
func (_MasterNode *MasterNodeCaller) ExistPubkey(opts *bind.CallOpts, _pubkey string) (bool, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "existPubkey", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistPubkey is a free data retrieval call binding the contract method 0x4afe2bfb.
//
// Solidity: function existPubkey(string _pubkey) view returns(bool)
func (_MasterNode *MasterNodeSession) ExistPubkey(_pubkey string) (bool, error) {
	return _MasterNode.Contract.ExistPubkey(&_MasterNode.CallOpts, _pubkey)
}

// ExistPubkey is a free data retrieval call binding the contract method 0x4afe2bfb.
//
// Solidity: function existPubkey(string _pubkey) view returns(bool)
func (_MasterNode *MasterNodeCallerSession) ExistPubkey(_pubkey string) (bool, error) {
	return _MasterNode.Contract.ExistPubkey(&_MasterNode.CallOpts, _pubkey)
}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _addr) view returns((uint256,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),uint256,uint256))
func (_MasterNode *MasterNodeCaller) GetInfo(opts *bind.CallOpts, _addr common.Address) (IMasterNodeMasterNodeInfo, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "getInfo", _addr)

	if err != nil {
		return *new(IMasterNodeMasterNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IMasterNodeMasterNodeInfo)).(*IMasterNodeMasterNodeInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _addr) view returns((uint256,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),uint256,uint256))
func (_MasterNode *MasterNodeSession) GetInfo(_addr common.Address) (IMasterNodeMasterNodeInfo, error) {
	return _MasterNode.Contract.GetInfo(&_MasterNode.CallOpts, _addr)
}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _addr) view returns((uint256,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),uint256,uint256))
func (_MasterNode *MasterNodeCallerSession) GetInfo(_addr common.Address) (IMasterNodeMasterNodeInfo, error) {
	return _MasterNode.Contract.GetInfo(&_MasterNode.CallOpts, _addr)
}

// GetNext is a free data retrieval call binding the contract method 0xf3638f78.
//
// Solidity: function getNext() view returns(address)
func (_MasterNode *MasterNodeCaller) GetNext(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "getNext")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNext is a free data retrieval call binding the contract method 0xf3638f78.
//
// Solidity: function getNext() view returns(address)
func (_MasterNode *MasterNodeSession) GetNext() (common.Address, error) {
	return _MasterNode.Contract.GetNext(&_MasterNode.CallOpts)
}

// GetNext is a free data retrieval call binding the contract method 0xf3638f78.
//
// Solidity: function getNext() view returns(address)
func (_MasterNode *MasterNodeCallerSession) GetNext() (common.Address, error) {
	return _MasterNode.Contract.GetNext(&_MasterNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterNode *MasterNodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterNode.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterNode *MasterNodeSession) Owner() (common.Address, error) {
	return _MasterNode.Contract.Owner(&_MasterNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterNode *MasterNodeCallerSession) Owner() (common.Address, error) {
	return _MasterNode.Contract.Owner(&_MasterNode.CallOpts)
}

// AppendRegister is a paid mutator transaction binding the contract method 0x978a11d1.
//
// Solidity: function appendRegister(address _addr, uint256 _lockDay) payable returns()
func (_MasterNode *MasterNodeTransactor) AppendRegister(opts *bind.TransactOpts, _addr common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "appendRegister", _addr, _lockDay)
}

// AppendRegister is a paid mutator transaction binding the contract method 0x978a11d1.
//
// Solidity: function appendRegister(address _addr, uint256 _lockDay) payable returns()
func (_MasterNode *MasterNodeSession) AppendRegister(_addr common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _MasterNode.Contract.AppendRegister(&_MasterNode.TransactOpts, _addr, _lockDay)
}

// AppendRegister is a paid mutator transaction binding the contract method 0x978a11d1.
//
// Solidity: function appendRegister(address _addr, uint256 _lockDay) payable returns()
func (_MasterNode *MasterNodeTransactorSession) AppendRegister(_addr common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _MasterNode.Contract.AppendRegister(&_MasterNode.TransactOpts, _addr, _lockDay)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address _addr, address _newAddr) returns()
func (_MasterNode *MasterNodeTransactor) ChangeAddress(opts *bind.TransactOpts, _addr common.Address, _newAddr common.Address) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "changeAddress", _addr, _newAddr)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address _addr, address _newAddr) returns()
func (_MasterNode *MasterNodeSession) ChangeAddress(_addr common.Address, _newAddr common.Address) (*types.Transaction, error) {
	return _MasterNode.Contract.ChangeAddress(&_MasterNode.TransactOpts, _addr, _newAddr)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address _addr, address _newAddr) returns()
func (_MasterNode *MasterNodeTransactorSession) ChangeAddress(_addr common.Address, _newAddr common.Address) (*types.Transaction, error) {
	return _MasterNode.Contract.ChangeAddress(&_MasterNode.TransactOpts, _addr, _newAddr)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0x1ed6f423.
//
// Solidity: function changeDescription(address _addr, string _description) returns()
func (_MasterNode *MasterNodeTransactor) ChangeDescription(opts *bind.TransactOpts, _addr common.Address, _description string) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "changeDescription", _addr, _description)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0x1ed6f423.
//
// Solidity: function changeDescription(address _addr, string _description) returns()
func (_MasterNode *MasterNodeSession) ChangeDescription(_addr common.Address, _description string) (*types.Transaction, error) {
	return _MasterNode.Contract.ChangeDescription(&_MasterNode.TransactOpts, _addr, _description)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0x1ed6f423.
//
// Solidity: function changeDescription(address _addr, string _description) returns()
func (_MasterNode *MasterNodeTransactorSession) ChangeDescription(_addr common.Address, _description string) (*types.Transaction, error) {
	return _MasterNode.Contract.ChangeDescription(&_MasterNode.TransactOpts, _addr, _description)
}

// ChangeIP is a paid mutator transaction binding the contract method 0x46d3b3be.
//
// Solidity: function changeIP(address _addr, string _ip) returns()
func (_MasterNode *MasterNodeTransactor) ChangeIP(opts *bind.TransactOpts, _addr common.Address, _ip string) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "changeIP", _addr, _ip)
}

// ChangeIP is a paid mutator transaction binding the contract method 0x46d3b3be.
//
// Solidity: function changeIP(address _addr, string _ip) returns()
func (_MasterNode *MasterNodeSession) ChangeIP(_addr common.Address, _ip string) (*types.Transaction, error) {
	return _MasterNode.Contract.ChangeIP(&_MasterNode.TransactOpts, _addr, _ip)
}

// ChangeIP is a paid mutator transaction binding the contract method 0x46d3b3be.
//
// Solidity: function changeIP(address _addr, string _ip) returns()
func (_MasterNode *MasterNodeTransactorSession) ChangeIP(_addr common.Address, _ip string) (*types.Transaction, error) {
	return _MasterNode.Contract.ChangeIP(&_MasterNode.TransactOpts, _addr, _ip)
}

// ChangePubkey is a paid mutator transaction binding the contract method 0xf333e362.
//
// Solidity: function changePubkey(address _addr, string _pubkey) returns()
func (_MasterNode *MasterNodeTransactor) ChangePubkey(opts *bind.TransactOpts, _addr common.Address, _pubkey string) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "changePubkey", _addr, _pubkey)
}

// ChangePubkey is a paid mutator transaction binding the contract method 0xf333e362.
//
// Solidity: function changePubkey(address _addr, string _pubkey) returns()
func (_MasterNode *MasterNodeSession) ChangePubkey(_addr common.Address, _pubkey string) (*types.Transaction, error) {
	return _MasterNode.Contract.ChangePubkey(&_MasterNode.TransactOpts, _addr, _pubkey)
}

// ChangePubkey is a paid mutator transaction binding the contract method 0xf333e362.
//
// Solidity: function changePubkey(address _addr, string _pubkey) returns()
func (_MasterNode *MasterNodeTransactorSession) ChangePubkey(_addr common.Address, _pubkey string) (*types.Transaction, error) {
	return _MasterNode.Contract.ChangePubkey(&_MasterNode.TransactOpts, _addr, _pubkey)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MasterNode *MasterNodeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MasterNode *MasterNodeSession) Initialize() (*types.Transaction, error) {
	return _MasterNode.Contract.Initialize(&_MasterNode.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MasterNode *MasterNodeTransactorSession) Initialize() (*types.Transaction, error) {
	return _MasterNode.Contract.Initialize(&_MasterNode.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xcabad5b0.
//
// Solidity: function register(address _addr, bool _isUnion, uint256 _lockDay, string _ip, string _pubkey, string _description, uint256 _creatorIncentive, uint256 _partnerIncentive) payable returns()
func (_MasterNode *MasterNodeTransactor) Register(opts *bind.TransactOpts, _addr common.Address, _isUnion bool, _lockDay *big.Int, _ip string, _pubkey string, _description string, _creatorIncentive *big.Int, _partnerIncentive *big.Int) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "register", _addr, _isUnion, _lockDay, _ip, _pubkey, _description, _creatorIncentive, _partnerIncentive)
}

// Register is a paid mutator transaction binding the contract method 0xcabad5b0.
//
// Solidity: function register(address _addr, bool _isUnion, uint256 _lockDay, string _ip, string _pubkey, string _description, uint256 _creatorIncentive, uint256 _partnerIncentive) payable returns()
func (_MasterNode *MasterNodeSession) Register(_addr common.Address, _isUnion bool, _lockDay *big.Int, _ip string, _pubkey string, _description string, _creatorIncentive *big.Int, _partnerIncentive *big.Int) (*types.Transaction, error) {
	return _MasterNode.Contract.Register(&_MasterNode.TransactOpts, _addr, _isUnion, _lockDay, _ip, _pubkey, _description, _creatorIncentive, _partnerIncentive)
}

// Register is a paid mutator transaction binding the contract method 0xcabad5b0.
//
// Solidity: function register(address _addr, bool _isUnion, uint256 _lockDay, string _ip, string _pubkey, string _description, uint256 _creatorIncentive, uint256 _partnerIncentive) payable returns()
func (_MasterNode *MasterNodeTransactorSession) Register(_addr common.Address, _isUnion bool, _lockDay *big.Int, _ip string, _pubkey string, _description string, _creatorIncentive *big.Int, _partnerIncentive *big.Int) (*types.Transaction, error) {
	return _MasterNode.Contract.Register(&_MasterNode.TransactOpts, _addr, _isUnion, _lockDay, _ip, _pubkey, _description, _creatorIncentive, _partnerIncentive)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterNode *MasterNodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterNode *MasterNodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterNode.Contract.RenounceOwnership(&_MasterNode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterNode *MasterNodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterNode.Contract.RenounceOwnership(&_MasterNode.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x6353586b.
//
// Solidity: function reward(address _addr) payable returns()
func (_MasterNode *MasterNodeTransactor) Reward(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "reward", _addr)
}

// Reward is a paid mutator transaction binding the contract method 0x6353586b.
//
// Solidity: function reward(address _addr) payable returns()
func (_MasterNode *MasterNodeSession) Reward(_addr common.Address) (*types.Transaction, error) {
	return _MasterNode.Contract.Reward(&_MasterNode.TransactOpts, _addr)
}

// Reward is a paid mutator transaction binding the contract method 0x6353586b.
//
// Solidity: function reward(address _addr) payable returns()
func (_MasterNode *MasterNodeTransactorSession) Reward(_addr common.Address) (*types.Transaction, error) {
	return _MasterNode.Contract.Reward(&_MasterNode.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterNode *MasterNodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MasterNode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterNode *MasterNodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterNode.Contract.TransferOwnership(&_MasterNode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterNode *MasterNodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterNode.Contract.TransferOwnership(&_MasterNode.TransactOpts, newOwner)
}

// MasterNodeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MasterNode contract.
type MasterNodeInitializedIterator struct {
	Event *MasterNodeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MasterNodeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterNodeInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MasterNodeInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MasterNodeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterNodeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterNodeInitialized represents a Initialized event raised by the MasterNode contract.
type MasterNodeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MasterNode *MasterNodeFilterer) FilterInitialized(opts *bind.FilterOpts) (*MasterNodeInitializedIterator, error) {

	logs, sub, err := _MasterNode.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MasterNodeInitializedIterator{contract: _MasterNode.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MasterNode *MasterNodeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MasterNodeInitialized) (event.Subscription, error) {

	logs, sub, err := _MasterNode.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterNodeInitialized)
				if err := _MasterNode.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MasterNode *MasterNodeFilterer) ParseInitialized(log types.Log) (*MasterNodeInitialized, error) {
	event := new(MasterNodeInitialized)
	if err := _MasterNode.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterNodeMNAppendRegisterIterator is returned from FilterMNAppendRegister and is used to iterate over the raw logs and unpacked data for MNAppendRegister events raised by the MasterNode contract.
type MasterNodeMNAppendRegisterIterator struct {
	Event *MasterNodeMNAppendRegister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MasterNodeMNAppendRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterNodeMNAppendRegister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MasterNodeMNAppendRegister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MasterNodeMNAppendRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterNodeMNAppendRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterNodeMNAppendRegister represents a MNAppendRegister event raised by the MasterNode contract.
type MasterNodeMNAppendRegister struct {
	Addr     common.Address
	Operator common.Address
	Amount   *big.Int
	LockDay  *big.Int
	RecordID [20]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMNAppendRegister is a free log retrieval operation binding the contract event 0x58f8614f0d53293d34dde92e63a3061d0dc19f5439f35dd6c0583ad9f8ce4e1c.
//
// Solidity: event MNAppendRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _recordID)
func (_MasterNode *MasterNodeFilterer) FilterMNAppendRegister(opts *bind.FilterOpts) (*MasterNodeMNAppendRegisterIterator, error) {

	logs, sub, err := _MasterNode.contract.FilterLogs(opts, "MNAppendRegister")
	if err != nil {
		return nil, err
	}
	return &MasterNodeMNAppendRegisterIterator{contract: _MasterNode.contract, event: "MNAppendRegister", logs: logs, sub: sub}, nil
}

// WatchMNAppendRegister is a free log subscription operation binding the contract event 0x58f8614f0d53293d34dde92e63a3061d0dc19f5439f35dd6c0583ad9f8ce4e1c.
//
// Solidity: event MNAppendRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _recordID)
func (_MasterNode *MasterNodeFilterer) WatchMNAppendRegister(opts *bind.WatchOpts, sink chan<- *MasterNodeMNAppendRegister) (event.Subscription, error) {

	logs, sub, err := _MasterNode.contract.WatchLogs(opts, "MNAppendRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterNodeMNAppendRegister)
				if err := _MasterNode.contract.UnpackLog(event, "MNAppendRegister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMNAppendRegister is a log parse operation binding the contract event 0x58f8614f0d53293d34dde92e63a3061d0dc19f5439f35dd6c0583ad9f8ce4e1c.
//
// Solidity: event MNAppendRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _recordID)
func (_MasterNode *MasterNodeFilterer) ParseMNAppendRegister(log types.Log) (*MasterNodeMNAppendRegister, error) {
	event := new(MasterNodeMNAppendRegister)
	if err := _MasterNode.contract.UnpackLog(event, "MNAppendRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterNodeMNRegisterIterator is returned from FilterMNRegister and is used to iterate over the raw logs and unpacked data for MNRegister events raised by the MasterNode contract.
type MasterNodeMNRegisterIterator struct {
	Event *MasterNodeMNRegister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MasterNodeMNRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterNodeMNRegister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MasterNodeMNRegister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MasterNodeMNRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterNodeMNRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterNodeMNRegister represents a MNRegister event raised by the MasterNode contract.
type MasterNodeMNRegister struct {
	Addr     common.Address
	Operator common.Address
	Amount   *big.Int
	LockDay  *big.Int
	ReocrdID [20]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMNRegister is a free log retrieval operation binding the contract event 0x21e4571bf1bcae0cf50dd00294f22f2f8fb6df2f4a5ce07c8207d5bb19f9d272.
//
// Solidity: event MNRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_MasterNode *MasterNodeFilterer) FilterMNRegister(opts *bind.FilterOpts) (*MasterNodeMNRegisterIterator, error) {

	logs, sub, err := _MasterNode.contract.FilterLogs(opts, "MNRegister")
	if err != nil {
		return nil, err
	}
	return &MasterNodeMNRegisterIterator{contract: _MasterNode.contract, event: "MNRegister", logs: logs, sub: sub}, nil
}

// WatchMNRegister is a free log subscription operation binding the contract event 0x21e4571bf1bcae0cf50dd00294f22f2f8fb6df2f4a5ce07c8207d5bb19f9d272.
//
// Solidity: event MNRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_MasterNode *MasterNodeFilterer) WatchMNRegister(opts *bind.WatchOpts, sink chan<- *MasterNodeMNRegister) (event.Subscription, error) {

	logs, sub, err := _MasterNode.contract.WatchLogs(opts, "MNRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterNodeMNRegister)
				if err := _MasterNode.contract.UnpackLog(event, "MNRegister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMNRegister is a log parse operation binding the contract event 0x21e4571bf1bcae0cf50dd00294f22f2f8fb6df2f4a5ce07c8207d5bb19f9d272.
//
// Solidity: event MNRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_MasterNode *MasterNodeFilterer) ParseMNRegister(log types.Log) (*MasterNodeMNRegister, error) {
	event := new(MasterNodeMNRegister)
	if err := _MasterNode.contract.UnpackLog(event, "MNRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterNodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MasterNode contract.
type MasterNodeOwnershipTransferredIterator struct {
	Event *MasterNodeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MasterNodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterNodeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MasterNodeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MasterNodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterNodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterNodeOwnershipTransferred represents a OwnershipTransferred event raised by the MasterNode contract.
type MasterNodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterNode *MasterNodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MasterNodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterNode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MasterNodeOwnershipTransferredIterator{contract: _MasterNode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterNode *MasterNodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MasterNodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterNode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterNodeOwnershipTransferred)
				if err := _MasterNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterNode *MasterNodeFilterer) ParseOwnershipTransferred(log types.Log) (*MasterNodeOwnershipTransferred, error) {
	event := new(MasterNodeOwnershipTransferred)
	if err := _MasterNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
