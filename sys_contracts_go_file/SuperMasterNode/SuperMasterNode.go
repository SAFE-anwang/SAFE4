// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SuperMasterNode

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

// ISuperMasterNodeIncentivePlan is an auto generated low-level Go binding around an user-defined struct.
type ISuperMasterNodeIncentivePlan struct {
	Creator *big.Int
	Partner *big.Int
	Voter   *big.Int
}

// ISuperMasterNodeMemberInfo is an auto generated low-level Go binding around an user-defined struct.
type ISuperMasterNodeMemberInfo struct {
	LockID [20]byte
	Addr   common.Address
	Amount *big.Int
	Height *big.Int
}

// ISuperMasterNodeSuperMasterNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type ISuperMasterNodeSuperMasterNodeInfo struct {
	Id               *big.Int
	Name             string
	Addr             common.Address
	Creator          common.Address
	Amount           *big.Int
	Ip               string
	Pubkey           string
	Description      string
	State            *big.Int
	Founders         []ISuperMasterNodeMemberInfo
	IncentivePlan    ISuperMasterNodeIncentivePlan
	Voters           []ISuperMasterNodeMemberInfo
	TotalVoteNum     *big.Int
	TotalVoterAmount *big.Int
	CreateHeight     *big.Int
	UpdateHeight     *big.Int
}

// SuperMasterNodeMetaData contains all meta data concerning the SuperMasterNode contract.
var SuperMasterNodeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"}],\"name\":\"SMNAppendRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_reocrdID\",\"type\":\"bytes20\"}],\"name\":\"SMNRegister\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"}],\"name\":\"appendRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAddr\",\"type\":\"address\"}],\"name\":\"changeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_newDescription\",\"type\":\"string\"}],\"name\":\"changeDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_newIP\",\"type\":\"string\"}],\"name\":\"changeIP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_newPubkey\",\"type\":\"string\"}],\"name\":\"changePubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"exist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"existID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"}],\"name\":\"existIP\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_lockID\",\"type\":\"bytes20\"}],\"name\":\"existLockID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"existPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes20\",\"name\":\"lockID\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structISuperMasterNode.MemberInfo[]\",\"name\":\"founders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"partner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voter\",\"type\":\"uint256\"}],\"internalType\":\"structISuperMasterNode.IncentivePlan\",\"name\":\"incentivePlan\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes20\",\"name\":\"lockID\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structISuperMasterNode.MemberInfo[]\",\"name\":\"voters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalVoteNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVoterAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structISuperMasterNode.SuperMasterNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes20\",\"name\":\"lockID\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structISuperMasterNode.MemberInfo[]\",\"name\":\"founders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"partner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voter\",\"type\":\"uint256\"}],\"internalType\":\"structISuperMasterNode.IncentivePlan\",\"name\":\"incentivePlan\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes20\",\"name\":\"lockID\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structISuperMasterNode.MemberInfo[]\",\"name\":\"voters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalVoteNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVoterAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structISuperMasterNode.SuperMasterNodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isUnion\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_creatorIncentive\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_partnerIncentive\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_voterIncentive\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SuperMasterNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use SuperMasterNodeMetaData.ABI instead.
var SuperMasterNodeABI = SuperMasterNodeMetaData.ABI

// SuperMasterNode is an auto generated Go binding around an Ethereum contract.
type SuperMasterNode struct {
	SuperMasterNodeCaller     // Read-only binding to the contract
	SuperMasterNodeTransactor // Write-only binding to the contract
	SuperMasterNodeFilterer   // Log filterer for contract events
}

// SuperMasterNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuperMasterNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperMasterNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuperMasterNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperMasterNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuperMasterNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperMasterNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuperMasterNodeSession struct {
	Contract     *SuperMasterNode  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuperMasterNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuperMasterNodeCallerSession struct {
	Contract *SuperMasterNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SuperMasterNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuperMasterNodeTransactorSession struct {
	Contract     *SuperMasterNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SuperMasterNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuperMasterNodeRaw struct {
	Contract *SuperMasterNode // Generic contract binding to access the raw methods on
}

// SuperMasterNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuperMasterNodeCallerRaw struct {
	Contract *SuperMasterNodeCaller // Generic read-only contract binding to access the raw methods on
}

// SuperMasterNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuperMasterNodeTransactorRaw struct {
	Contract *SuperMasterNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuperMasterNode creates a new instance of SuperMasterNode, bound to a specific deployed contract.
func NewSuperMasterNode(address common.Address, backend bind.ContractBackend) (*SuperMasterNode, error) {
	contract, err := bindSuperMasterNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SuperMasterNode{SuperMasterNodeCaller: SuperMasterNodeCaller{contract: contract}, SuperMasterNodeTransactor: SuperMasterNodeTransactor{contract: contract}, SuperMasterNodeFilterer: SuperMasterNodeFilterer{contract: contract}}, nil
}

// NewSuperMasterNodeCaller creates a new read-only instance of SuperMasterNode, bound to a specific deployed contract.
func NewSuperMasterNodeCaller(address common.Address, caller bind.ContractCaller) (*SuperMasterNodeCaller, error) {
	contract, err := bindSuperMasterNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuperMasterNodeCaller{contract: contract}, nil
}

// NewSuperMasterNodeTransactor creates a new write-only instance of SuperMasterNode, bound to a specific deployed contract.
func NewSuperMasterNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*SuperMasterNodeTransactor, error) {
	contract, err := bindSuperMasterNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuperMasterNodeTransactor{contract: contract}, nil
}

// NewSuperMasterNodeFilterer creates a new log filterer instance of SuperMasterNode, bound to a specific deployed contract.
func NewSuperMasterNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*SuperMasterNodeFilterer, error) {
	contract, err := bindSuperMasterNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuperMasterNodeFilterer{contract: contract}, nil
}

// bindSuperMasterNode binds a generic wrapper to an already deployed contract.
func bindSuperMasterNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SuperMasterNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperMasterNode *SuperMasterNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SuperMasterNode.Contract.SuperMasterNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperMasterNode *SuperMasterNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.SuperMasterNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperMasterNode *SuperMasterNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.SuperMasterNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperMasterNode *SuperMasterNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SuperMasterNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperMasterNode *SuperMasterNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperMasterNode *SuperMasterNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SuperMasterNode *SuperMasterNodeCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SuperMasterNode *SuperMasterNodeSession) GetInitializeData() ([]byte, error) {
	return _SuperMasterNode.Contract.GetInitializeData(&_SuperMasterNode.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SuperMasterNode *SuperMasterNodeCallerSession) GetInitializeData() ([]byte, error) {
	return _SuperMasterNode.Contract.GetInitializeData(&_SuperMasterNode.CallOpts)
}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _addr) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCaller) Exist(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "exist", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _addr) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeSession) Exist(_addr common.Address) (bool, error) {
	return _SuperMasterNode.Contract.Exist(&_SuperMasterNode.CallOpts, _addr)
}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _addr) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCallerSession) Exist(_addr common.Address) (bool, error) {
	return _SuperMasterNode.Contract.Exist(&_SuperMasterNode.CallOpts, _addr)
}

// ExistID is a free data retrieval call binding the contract method 0x829c82e1.
//
// Solidity: function existID(uint256 _id) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCaller) ExistID(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "existID", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistID is a free data retrieval call binding the contract method 0x829c82e1.
//
// Solidity: function existID(uint256 _id) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeSession) ExistID(_id *big.Int) (bool, error) {
	return _SuperMasterNode.Contract.ExistID(&_SuperMasterNode.CallOpts, _id)
}

// ExistID is a free data retrieval call binding the contract method 0x829c82e1.
//
// Solidity: function existID(uint256 _id) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCallerSession) ExistID(_id *big.Int) (bool, error) {
	return _SuperMasterNode.Contract.ExistID(&_SuperMasterNode.CallOpts, _id)
}

// ExistIP is a free data retrieval call binding the contract method 0xde1b1428.
//
// Solidity: function existIP(string _ip) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCaller) ExistIP(opts *bind.CallOpts, _ip string) (bool, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "existIP", _ip)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistIP is a free data retrieval call binding the contract method 0xde1b1428.
//
// Solidity: function existIP(string _ip) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeSession) ExistIP(_ip string) (bool, error) {
	return _SuperMasterNode.Contract.ExistIP(&_SuperMasterNode.CallOpts, _ip)
}

// ExistIP is a free data retrieval call binding the contract method 0xde1b1428.
//
// Solidity: function existIP(string _ip) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCallerSession) ExistIP(_ip string) (bool, error) {
	return _SuperMasterNode.Contract.ExistIP(&_SuperMasterNode.CallOpts, _ip)
}

// ExistLockID is a free data retrieval call binding the contract method 0x598bdb3d.
//
// Solidity: function existLockID(address _addr, bytes20 _lockID) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCaller) ExistLockID(opts *bind.CallOpts, _addr common.Address, _lockID [20]byte) (bool, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "existLockID", _addr, _lockID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistLockID is a free data retrieval call binding the contract method 0x598bdb3d.
//
// Solidity: function existLockID(address _addr, bytes20 _lockID) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeSession) ExistLockID(_addr common.Address, _lockID [20]byte) (bool, error) {
	return _SuperMasterNode.Contract.ExistLockID(&_SuperMasterNode.CallOpts, _addr, _lockID)
}

// ExistLockID is a free data retrieval call binding the contract method 0x598bdb3d.
//
// Solidity: function existLockID(address _addr, bytes20 _lockID) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCallerSession) ExistLockID(_addr common.Address, _lockID [20]byte) (bool, error) {
	return _SuperMasterNode.Contract.ExistLockID(&_SuperMasterNode.CallOpts, _addr, _lockID)
}

// ExistPubkey is a free data retrieval call binding the contract method 0x4afe2bfb.
//
// Solidity: function existPubkey(string _pubkey) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCaller) ExistPubkey(opts *bind.CallOpts, _pubkey string) (bool, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "existPubkey", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistPubkey is a free data retrieval call binding the contract method 0x4afe2bfb.
//
// Solidity: function existPubkey(string _pubkey) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeSession) ExistPubkey(_pubkey string) (bool, error) {
	return _SuperMasterNode.Contract.ExistPubkey(&_SuperMasterNode.CallOpts, _pubkey)
}

// ExistPubkey is a free data retrieval call binding the contract method 0x4afe2bfb.
//
// Solidity: function existPubkey(string _pubkey) view returns(bool)
func (_SuperMasterNode *SuperMasterNodeCallerSession) ExistPubkey(_pubkey string) (bool, error) {
	return _SuperMasterNode.Contract.ExistPubkey(&_SuperMasterNode.CallOpts, _pubkey)
}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _addr) view returns((uint256,string,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),(bytes20,address,uint256,uint256)[],uint256,uint256,uint256,uint256))
func (_SuperMasterNode *SuperMasterNodeCaller) GetInfo(opts *bind.CallOpts, _addr common.Address) (ISuperMasterNodeSuperMasterNodeInfo, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "getInfo", _addr)

	if err != nil {
		return *new(ISuperMasterNodeSuperMasterNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISuperMasterNodeSuperMasterNodeInfo)).(*ISuperMasterNodeSuperMasterNodeInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _addr) view returns((uint256,string,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),(bytes20,address,uint256,uint256)[],uint256,uint256,uint256,uint256))
func (_SuperMasterNode *SuperMasterNodeSession) GetInfo(_addr common.Address) (ISuperMasterNodeSuperMasterNodeInfo, error) {
	return _SuperMasterNode.Contract.GetInfo(&_SuperMasterNode.CallOpts, _addr)
}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _addr) view returns((uint256,string,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),(bytes20,address,uint256,uint256)[],uint256,uint256,uint256,uint256))
func (_SuperMasterNode *SuperMasterNodeCallerSession) GetInfo(_addr common.Address) (ISuperMasterNodeSuperMasterNodeInfo, error) {
	return _SuperMasterNode.Contract.GetInfo(&_SuperMasterNode.CallOpts, _addr)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_SuperMasterNode *SuperMasterNodeCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "getNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_SuperMasterNode *SuperMasterNodeSession) GetNum() (*big.Int, error) {
	return _SuperMasterNode.Contract.GetNum(&_SuperMasterNode.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_SuperMasterNode *SuperMasterNodeCallerSession) GetNum() (*big.Int, error) {
	return _SuperMasterNode.Contract.GetNum(&_SuperMasterNode.CallOpts)
}

// GetTop is a free data retrieval call binding the contract method 0x5c2b1119.
//
// Solidity: function getTop() view returns((uint256,string,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),(bytes20,address,uint256,uint256)[],uint256,uint256,uint256,uint256)[])
func (_SuperMasterNode *SuperMasterNodeCaller) GetTop(opts *bind.CallOpts) ([]ISuperMasterNodeSuperMasterNodeInfo, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "getTop")

	if err != nil {
		return *new([]ISuperMasterNodeSuperMasterNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISuperMasterNodeSuperMasterNodeInfo)).(*[]ISuperMasterNodeSuperMasterNodeInfo)

	return out0, err

}

// GetTop is a free data retrieval call binding the contract method 0x5c2b1119.
//
// Solidity: function getTop() view returns((uint256,string,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),(bytes20,address,uint256,uint256)[],uint256,uint256,uint256,uint256)[])
func (_SuperMasterNode *SuperMasterNodeSession) GetTop() ([]ISuperMasterNodeSuperMasterNodeInfo, error) {
	return _SuperMasterNode.Contract.GetTop(&_SuperMasterNode.CallOpts)
}

// GetTop is a free data retrieval call binding the contract method 0x5c2b1119.
//
// Solidity: function getTop() view returns((uint256,string,address,address,uint256,string,string,string,uint256,(bytes20,address,uint256,uint256)[],(uint256,uint256,uint256),(bytes20,address,uint256,uint256)[],uint256,uint256,uint256,uint256)[])
func (_SuperMasterNode *SuperMasterNodeCallerSession) GetTop() ([]ISuperMasterNodeSuperMasterNodeInfo, error) {
	return _SuperMasterNode.Contract.GetTop(&_SuperMasterNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SuperMasterNode *SuperMasterNodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SuperMasterNode.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SuperMasterNode *SuperMasterNodeSession) Owner() (common.Address, error) {
	return _SuperMasterNode.Contract.Owner(&_SuperMasterNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SuperMasterNode *SuperMasterNodeCallerSession) Owner() (common.Address, error) {
	return _SuperMasterNode.Contract.Owner(&_SuperMasterNode.CallOpts)
}

// AppendRegister is a paid mutator transaction binding the contract method 0x978a11d1.
//
// Solidity: function appendRegister(address _addr, uint256 _lockDay) payable returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) AppendRegister(opts *bind.TransactOpts, _addr common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "appendRegister", _addr, _lockDay)
}

// AppendRegister is a paid mutator transaction binding the contract method 0x978a11d1.
//
// Solidity: function appendRegister(address _addr, uint256 _lockDay) payable returns()
func (_SuperMasterNode *SuperMasterNodeSession) AppendRegister(_addr common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.AppendRegister(&_SuperMasterNode.TransactOpts, _addr, _lockDay)
}

// AppendRegister is a paid mutator transaction binding the contract method 0x978a11d1.
//
// Solidity: function appendRegister(address _addr, uint256 _lockDay) payable returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) AppendRegister(_addr common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.AppendRegister(&_SuperMasterNode.TransactOpts, _addr, _lockDay)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address _addr, address _newAddr) returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) ChangeAddress(opts *bind.TransactOpts, _addr common.Address, _newAddr common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "changeAddress", _addr, _newAddr)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address _addr, address _newAddr) returns()
func (_SuperMasterNode *SuperMasterNodeSession) ChangeAddress(_addr common.Address, _newAddr common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.ChangeAddress(&_SuperMasterNode.TransactOpts, _addr, _newAddr)
}

// ChangeAddress is a paid mutator transaction binding the contract method 0xefe08a7d.
//
// Solidity: function changeAddress(address _addr, address _newAddr) returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) ChangeAddress(_addr common.Address, _newAddr common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.ChangeAddress(&_SuperMasterNode.TransactOpts, _addr, _newAddr)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0x1ed6f423.
//
// Solidity: function changeDescription(address _addr, string _newDescription) returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) ChangeDescription(opts *bind.TransactOpts, _addr common.Address, _newDescription string) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "changeDescription", _addr, _newDescription)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0x1ed6f423.
//
// Solidity: function changeDescription(address _addr, string _newDescription) returns()
func (_SuperMasterNode *SuperMasterNodeSession) ChangeDescription(_addr common.Address, _newDescription string) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.ChangeDescription(&_SuperMasterNode.TransactOpts, _addr, _newDescription)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0x1ed6f423.
//
// Solidity: function changeDescription(address _addr, string _newDescription) returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) ChangeDescription(_addr common.Address, _newDescription string) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.ChangeDescription(&_SuperMasterNode.TransactOpts, _addr, _newDescription)
}

// ChangeIP is a paid mutator transaction binding the contract method 0x46d3b3be.
//
// Solidity: function changeIP(address _addr, string _newIP) returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) ChangeIP(opts *bind.TransactOpts, _addr common.Address, _newIP string) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "changeIP", _addr, _newIP)
}

// ChangeIP is a paid mutator transaction binding the contract method 0x46d3b3be.
//
// Solidity: function changeIP(address _addr, string _newIP) returns()
func (_SuperMasterNode *SuperMasterNodeSession) ChangeIP(_addr common.Address, _newIP string) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.ChangeIP(&_SuperMasterNode.TransactOpts, _addr, _newIP)
}

// ChangeIP is a paid mutator transaction binding the contract method 0x46d3b3be.
//
// Solidity: function changeIP(address _addr, string _newIP) returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) ChangeIP(_addr common.Address, _newIP string) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.ChangeIP(&_SuperMasterNode.TransactOpts, _addr, _newIP)
}

// ChangePubkey is a paid mutator transaction binding the contract method 0xf333e362.
//
// Solidity: function changePubkey(address _addr, string _newPubkey) returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) ChangePubkey(opts *bind.TransactOpts, _addr common.Address, _newPubkey string) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "changePubkey", _addr, _newPubkey)
}

// ChangePubkey is a paid mutator transaction binding the contract method 0xf333e362.
//
// Solidity: function changePubkey(address _addr, string _newPubkey) returns()
func (_SuperMasterNode *SuperMasterNodeSession) ChangePubkey(_addr common.Address, _newPubkey string) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.ChangePubkey(&_SuperMasterNode.TransactOpts, _addr, _newPubkey)
}

// ChangePubkey is a paid mutator transaction binding the contract method 0xf333e362.
//
// Solidity: function changePubkey(address _addr, string _newPubkey) returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) ChangePubkey(_addr common.Address, _newPubkey string) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.ChangePubkey(&_SuperMasterNode.TransactOpts, _addr, _newPubkey)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SuperMasterNode *SuperMasterNodeSession) Initialize() (*types.Transaction, error) {
	return _SuperMasterNode.Contract.Initialize(&_SuperMasterNode.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) Initialize() (*types.Transaction, error) {
	return _SuperMasterNode.Contract.Initialize(&_SuperMasterNode.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xef24f055.
//
// Solidity: function register(address _addr, bool _isUnion, uint256 _lockDay, string _name, string _ip, string _pubkey, string _description, uint256 _creatorIncentive, uint256 _partnerIncentive, uint256 _voterIncentive) payable returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) Register(opts *bind.TransactOpts, _addr common.Address, _isUnion bool, _lockDay *big.Int, _name string, _ip string, _pubkey string, _description string, _creatorIncentive *big.Int, _partnerIncentive *big.Int, _voterIncentive *big.Int) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "register", _addr, _isUnion, _lockDay, _name, _ip, _pubkey, _description, _creatorIncentive, _partnerIncentive, _voterIncentive)
}

// Register is a paid mutator transaction binding the contract method 0xef24f055.
//
// Solidity: function register(address _addr, bool _isUnion, uint256 _lockDay, string _name, string _ip, string _pubkey, string _description, uint256 _creatorIncentive, uint256 _partnerIncentive, uint256 _voterIncentive) payable returns()
func (_SuperMasterNode *SuperMasterNodeSession) Register(_addr common.Address, _isUnion bool, _lockDay *big.Int, _name string, _ip string, _pubkey string, _description string, _creatorIncentive *big.Int, _partnerIncentive *big.Int, _voterIncentive *big.Int) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.Register(&_SuperMasterNode.TransactOpts, _addr, _isUnion, _lockDay, _name, _ip, _pubkey, _description, _creatorIncentive, _partnerIncentive, _voterIncentive)
}

// Register is a paid mutator transaction binding the contract method 0xef24f055.
//
// Solidity: function register(address _addr, bool _isUnion, uint256 _lockDay, string _name, string _ip, string _pubkey, string _description, uint256 _creatorIncentive, uint256 _partnerIncentive, uint256 _voterIncentive) payable returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) Register(_addr common.Address, _isUnion bool, _lockDay *big.Int, _name string, _ip string, _pubkey string, _description string, _creatorIncentive *big.Int, _partnerIncentive *big.Int, _voterIncentive *big.Int) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.Register(&_SuperMasterNode.TransactOpts, _addr, _isUnion, _lockDay, _name, _ip, _pubkey, _description, _creatorIncentive, _partnerIncentive, _voterIncentive)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SuperMasterNode *SuperMasterNodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _SuperMasterNode.Contract.RenounceOwnership(&_SuperMasterNode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SuperMasterNode.Contract.RenounceOwnership(&_SuperMasterNode.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x6353586b.
//
// Solidity: function reward(address _addr) payable returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) Reward(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "reward", _addr)
}

// Reward is a paid mutator transaction binding the contract method 0x6353586b.
//
// Solidity: function reward(address _addr) payable returns()
func (_SuperMasterNode *SuperMasterNodeSession) Reward(_addr common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.Reward(&_SuperMasterNode.TransactOpts, _addr)
}

// Reward is a paid mutator transaction binding the contract method 0x6353586b.
//
// Solidity: function reward(address _addr) payable returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) Reward(_addr common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.Reward(&_SuperMasterNode.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SuperMasterNode *SuperMasterNodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SuperMasterNode *SuperMasterNodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.TransferOwnership(&_SuperMasterNode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SuperMasterNode *SuperMasterNodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SuperMasterNode.Contract.TransferOwnership(&_SuperMasterNode.TransactOpts, newOwner)
}

// SuperMasterNodeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SuperMasterNode contract.
type SuperMasterNodeInitializedIterator struct {
	Event *SuperMasterNodeInitialized // Event containing the contract specifics and raw log

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
func (it *SuperMasterNodeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperMasterNodeInitialized)
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
		it.Event = new(SuperMasterNodeInitialized)
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
func (it *SuperMasterNodeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperMasterNodeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperMasterNodeInitialized represents a Initialized event raised by the SuperMasterNode contract.
type SuperMasterNodeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SuperMasterNode *SuperMasterNodeFilterer) FilterInitialized(opts *bind.FilterOpts) (*SuperMasterNodeInitializedIterator, error) {

	logs, sub, err := _SuperMasterNode.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SuperMasterNodeInitializedIterator{contract: _SuperMasterNode.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SuperMasterNode *SuperMasterNodeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SuperMasterNodeInitialized) (event.Subscription, error) {

	logs, sub, err := _SuperMasterNode.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperMasterNodeInitialized)
				if err := _SuperMasterNode.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SuperMasterNode *SuperMasterNodeFilterer) ParseInitialized(log types.Log) (*SuperMasterNodeInitialized, error) {
	event := new(SuperMasterNodeInitialized)
	if err := _SuperMasterNode.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperMasterNodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SuperMasterNode contract.
type SuperMasterNodeOwnershipTransferredIterator struct {
	Event *SuperMasterNodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SuperMasterNodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperMasterNodeOwnershipTransferred)
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
		it.Event = new(SuperMasterNodeOwnershipTransferred)
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
func (it *SuperMasterNodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperMasterNodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperMasterNodeOwnershipTransferred represents a OwnershipTransferred event raised by the SuperMasterNode contract.
type SuperMasterNodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SuperMasterNode *SuperMasterNodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SuperMasterNodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SuperMasterNode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SuperMasterNodeOwnershipTransferredIterator{contract: _SuperMasterNode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SuperMasterNode *SuperMasterNodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SuperMasterNodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SuperMasterNode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperMasterNodeOwnershipTransferred)
				if err := _SuperMasterNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SuperMasterNode *SuperMasterNodeFilterer) ParseOwnershipTransferred(log types.Log) (*SuperMasterNodeOwnershipTransferred, error) {
	event := new(SuperMasterNodeOwnershipTransferred)
	if err := _SuperMasterNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperMasterNodeSMNAppendRegisterIterator is returned from FilterSMNAppendRegister and is used to iterate over the raw logs and unpacked data for SMNAppendRegister events raised by the SuperMasterNode contract.
type SuperMasterNodeSMNAppendRegisterIterator struct {
	Event *SuperMasterNodeSMNAppendRegister // Event containing the contract specifics and raw log

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
func (it *SuperMasterNodeSMNAppendRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperMasterNodeSMNAppendRegister)
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
		it.Event = new(SuperMasterNodeSMNAppendRegister)
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
func (it *SuperMasterNodeSMNAppendRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperMasterNodeSMNAppendRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperMasterNodeSMNAppendRegister represents a SMNAppendRegister event raised by the SuperMasterNode contract.
type SuperMasterNodeSMNAppendRegister struct {
	Addr     common.Address
	Operator common.Address
	Amount   *big.Int
	LockDay  *big.Int
	RecordID [20]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSMNAppendRegister is a free log retrieval operation binding the contract event 0x57e04eae8e9f0a47b2270ce4ee6e72b46c5b0b89d96db1c3af2e2fdd65e079ca.
//
// Solidity: event SMNAppendRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _recordID)
func (_SuperMasterNode *SuperMasterNodeFilterer) FilterSMNAppendRegister(opts *bind.FilterOpts) (*SuperMasterNodeSMNAppendRegisterIterator, error) {

	logs, sub, err := _SuperMasterNode.contract.FilterLogs(opts, "SMNAppendRegister")
	if err != nil {
		return nil, err
	}
	return &SuperMasterNodeSMNAppendRegisterIterator{contract: _SuperMasterNode.contract, event: "SMNAppendRegister", logs: logs, sub: sub}, nil
}

// WatchSMNAppendRegister is a free log subscription operation binding the contract event 0x57e04eae8e9f0a47b2270ce4ee6e72b46c5b0b89d96db1c3af2e2fdd65e079ca.
//
// Solidity: event SMNAppendRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _recordID)
func (_SuperMasterNode *SuperMasterNodeFilterer) WatchSMNAppendRegister(opts *bind.WatchOpts, sink chan<- *SuperMasterNodeSMNAppendRegister) (event.Subscription, error) {

	logs, sub, err := _SuperMasterNode.contract.WatchLogs(opts, "SMNAppendRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperMasterNodeSMNAppendRegister)
				if err := _SuperMasterNode.contract.UnpackLog(event, "SMNAppendRegister", log); err != nil {
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

// ParseSMNAppendRegister is a log parse operation binding the contract event 0x57e04eae8e9f0a47b2270ce4ee6e72b46c5b0b89d96db1c3af2e2fdd65e079ca.
//
// Solidity: event SMNAppendRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _recordID)
func (_SuperMasterNode *SuperMasterNodeFilterer) ParseSMNAppendRegister(log types.Log) (*SuperMasterNodeSMNAppendRegister, error) {
	event := new(SuperMasterNodeSMNAppendRegister)
	if err := _SuperMasterNode.contract.UnpackLog(event, "SMNAppendRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperMasterNodeSMNRegisterIterator is returned from FilterSMNRegister and is used to iterate over the raw logs and unpacked data for SMNRegister events raised by the SuperMasterNode contract.
type SuperMasterNodeSMNRegisterIterator struct {
	Event *SuperMasterNodeSMNRegister // Event containing the contract specifics and raw log

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
func (it *SuperMasterNodeSMNRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperMasterNodeSMNRegister)
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
		it.Event = new(SuperMasterNodeSMNRegister)
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
func (it *SuperMasterNodeSMNRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperMasterNodeSMNRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperMasterNodeSMNRegister represents a SMNRegister event raised by the SuperMasterNode contract.
type SuperMasterNodeSMNRegister struct {
	Addr     common.Address
	Operator common.Address
	Amount   *big.Int
	LockDay  *big.Int
	ReocrdID [20]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSMNRegister is a free log retrieval operation binding the contract event 0xecd11dca679fcf3154ddadd599fa195c670cabdbe282ac52e908ec97864da313.
//
// Solidity: event SMNRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_SuperMasterNode *SuperMasterNodeFilterer) FilterSMNRegister(opts *bind.FilterOpts) (*SuperMasterNodeSMNRegisterIterator, error) {

	logs, sub, err := _SuperMasterNode.contract.FilterLogs(opts, "SMNRegister")
	if err != nil {
		return nil, err
	}
	return &SuperMasterNodeSMNRegisterIterator{contract: _SuperMasterNode.contract, event: "SMNRegister", logs: logs, sub: sub}, nil
}

// WatchSMNRegister is a free log subscription operation binding the contract event 0xecd11dca679fcf3154ddadd599fa195c670cabdbe282ac52e908ec97864da313.
//
// Solidity: event SMNRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_SuperMasterNode *SuperMasterNodeFilterer) WatchSMNRegister(opts *bind.WatchOpts, sink chan<- *SuperMasterNodeSMNRegister) (event.Subscription, error) {

	logs, sub, err := _SuperMasterNode.contract.WatchLogs(opts, "SMNRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperMasterNodeSMNRegister)
				if err := _SuperMasterNode.contract.UnpackLog(event, "SMNRegister", log); err != nil {
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

// ParseSMNRegister is a log parse operation binding the contract event 0xecd11dca679fcf3154ddadd599fa195c670cabdbe282ac52e908ec97864da313.
//
// Solidity: event SMNRegister(address _addr, address _operator, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_SuperMasterNode *SuperMasterNodeFilterer) ParseSMNRegister(log types.Log) (*SuperMasterNodeSMNRegister, error) {
	event := new(SuperMasterNodeSMNRegister)
	if err := _SuperMasterNode.contract.UnpackLog(event, "SMNRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
