// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Property

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

// IPropertyPropertyInfo is an auto generated low-level Go binding around an user-defined struct.
type IPropertyPropertyInfo struct {
	Name         string
	Value        *big.Int
	Description  string
	CreateHeight *big.Int
	UpdateHeight *big.Int
}

// IPropertyUnconfirmedPropertyInfo is an auto generated low-level Go binding around an user-defined struct.
type IPropertyUnconfirmedPropertyInfo struct {
	Name        string
	Value       *big.Int
	Applicant   common.Address
	Voters      []common.Address
	VoteResults []*big.Int
	Reason      string
	ApplyHeight *big.Int
}

// PropertyMetaData contains all meta data concerning the Property contract.
var PropertyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"PropertyAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldValue\",\"type\":\"uint256\"}],\"name\":\"PropertyApplyUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"}],\"name\":\"PropertyUpdateAgree\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"}],\"name\":\"PropertyUpdateReject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_voteResult\",\"type\":\"uint256\"}],\"name\":\"PropertyUpdateVote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"applyUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIProperty.PropertyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getUnconfirmedInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"applicant\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"voteResults\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"applyHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIProperty.UnconfirmedPropertyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_voteResult\",\"type\":\"uint256\"}],\"name\":\"vote4Update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PropertyABI is the input ABI used to generate the binding from.
// Deprecated: Use PropertyMetaData.ABI instead.
var PropertyABI = PropertyMetaData.ABI

// Property is an auto generated Go binding around an Ethereum contract.
type Property struct {
	PropertyCaller     // Read-only binding to the contract
	PropertyTransactor // Write-only binding to the contract
	PropertyFilterer   // Log filterer for contract events
}

// PropertyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropertyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropertyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PropertyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropertySession struct {
	Contract     *Property         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PropertyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropertyCallerSession struct {
	Contract *PropertyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PropertyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropertyTransactorSession struct {
	Contract     *PropertyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PropertyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropertyRaw struct {
	Contract *Property // Generic contract binding to access the raw methods on
}

// PropertyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropertyCallerRaw struct {
	Contract *PropertyCaller // Generic read-only contract binding to access the raw methods on
}

// PropertyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropertyTransactorRaw struct {
	Contract *PropertyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProperty creates a new instance of Property, bound to a specific deployed contract.
func NewProperty(address common.Address, backend bind.ContractBackend) (*Property, error) {
	contract, err := bindProperty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Property{PropertyCaller: PropertyCaller{contract: contract}, PropertyTransactor: PropertyTransactor{contract: contract}, PropertyFilterer: PropertyFilterer{contract: contract}}, nil
}

// NewPropertyCaller creates a new read-only instance of Property, bound to a specific deployed contract.
func NewPropertyCaller(address common.Address, caller bind.ContractCaller) (*PropertyCaller, error) {
	contract, err := bindProperty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyCaller{contract: contract}, nil
}

// NewPropertyTransactor creates a new write-only instance of Property, bound to a specific deployed contract.
func NewPropertyTransactor(address common.Address, transactor bind.ContractTransactor) (*PropertyTransactor, error) {
	contract, err := bindProperty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyTransactor{contract: contract}, nil
}

// NewPropertyFilterer creates a new log filterer instance of Property, bound to a specific deployed contract.
func NewPropertyFilterer(address common.Address, filterer bind.ContractFilterer) (*PropertyFilterer, error) {
	contract, err := bindProperty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PropertyFilterer{contract: contract}, nil
}

// bindProperty binds a generic wrapper to an already deployed contract.
func bindProperty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PropertyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Property *PropertyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Property.Contract.PropertyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Property *PropertyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Property.Contract.PropertyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Property *PropertyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Property.Contract.PropertyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Property *PropertyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Property.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Property *PropertyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Property.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Property *PropertyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Property.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Property *PropertyCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Property.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Property *PropertySession) GetInitializeData() ([]byte, error) {
	return _Property.Contract.GetInitializeData(&_Property.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Property *PropertyCallerSession) GetInitializeData() ([]byte, error) {
	return _Property.Contract.GetInitializeData(&_Property.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x7ee005ab.
//
// Solidity: function getInfo(string _name) view returns((string,uint256,string,uint256,uint256))
func (_Property *PropertyCaller) GetInfo(opts *bind.CallOpts, _name string) (IPropertyPropertyInfo, error) {
	var out []interface{}
	err := _Property.contract.Call(opts, &out, "getInfo", _name)

	if err != nil {
		return *new(IPropertyPropertyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPropertyPropertyInfo)).(*IPropertyPropertyInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x7ee005ab.
//
// Solidity: function getInfo(string _name) view returns((string,uint256,string,uint256,uint256))
func (_Property *PropertySession) GetInfo(_name string) (IPropertyPropertyInfo, error) {
	return _Property.Contract.GetInfo(&_Property.CallOpts, _name)
}

// GetInfo is a free data retrieval call binding the contract method 0x7ee005ab.
//
// Solidity: function getInfo(string _name) view returns((string,uint256,string,uint256,uint256))
func (_Property *PropertyCallerSession) GetInfo(_name string) (IPropertyPropertyInfo, error) {
	return _Property.Contract.GetInfo(&_Property.CallOpts, _name)
}

// GetUnconfirmedInfo is a free data retrieval call binding the contract method 0x71b0340d.
//
// Solidity: function getUnconfirmedInfo(string _name) view returns((string,uint256,address,address[],uint256[],string,uint256))
func (_Property *PropertyCaller) GetUnconfirmedInfo(opts *bind.CallOpts, _name string) (IPropertyUnconfirmedPropertyInfo, error) {
	var out []interface{}
	err := _Property.contract.Call(opts, &out, "getUnconfirmedInfo", _name)

	if err != nil {
		return *new(IPropertyUnconfirmedPropertyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPropertyUnconfirmedPropertyInfo)).(*IPropertyUnconfirmedPropertyInfo)

	return out0, err

}

// GetUnconfirmedInfo is a free data retrieval call binding the contract method 0x71b0340d.
//
// Solidity: function getUnconfirmedInfo(string _name) view returns((string,uint256,address,address[],uint256[],string,uint256))
func (_Property *PropertySession) GetUnconfirmedInfo(_name string) (IPropertyUnconfirmedPropertyInfo, error) {
	return _Property.Contract.GetUnconfirmedInfo(&_Property.CallOpts, _name)
}

// GetUnconfirmedInfo is a free data retrieval call binding the contract method 0x71b0340d.
//
// Solidity: function getUnconfirmedInfo(string _name) view returns((string,uint256,address,address[],uint256[],string,uint256))
func (_Property *PropertyCallerSession) GetUnconfirmedInfo(_name string) (IPropertyUnconfirmedPropertyInfo, error) {
	return _Property.Contract.GetUnconfirmedInfo(&_Property.CallOpts, _name)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string _name) view returns(uint256)
func (_Property *PropertyCaller) GetValue(opts *bind.CallOpts, _name string) (*big.Int, error) {
	var out []interface{}
	err := _Property.contract.Call(opts, &out, "getValue", _name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string _name) view returns(uint256)
func (_Property *PropertySession) GetValue(_name string) (*big.Int, error) {
	return _Property.Contract.GetValue(&_Property.CallOpts, _name)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string _name) view returns(uint256)
func (_Property *PropertyCallerSession) GetValue(_name string) (*big.Int, error) {
	return _Property.Contract.GetValue(&_Property.CallOpts, _name)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Property *PropertyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Property.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Property *PropertySession) Owner() (common.Address, error) {
	return _Property.Contract.Owner(&_Property.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Property *PropertyCallerSession) Owner() (common.Address, error) {
	return _Property.Contract.Owner(&_Property.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f15e547.
//
// Solidity: function add(string _name, uint256 _value, string _description) returns()
func (_Property *PropertyTransactor) Add(opts *bind.TransactOpts, _name string, _value *big.Int, _description string) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "add", _name, _value, _description)
}

// Add is a paid mutator transaction binding the contract method 0x4f15e547.
//
// Solidity: function add(string _name, uint256 _value, string _description) returns()
func (_Property *PropertySession) Add(_name string, _value *big.Int, _description string) (*types.Transaction, error) {
	return _Property.Contract.Add(&_Property.TransactOpts, _name, _value, _description)
}

// Add is a paid mutator transaction binding the contract method 0x4f15e547.
//
// Solidity: function add(string _name, uint256 _value, string _description) returns()
func (_Property *PropertyTransactorSession) Add(_name string, _value *big.Int, _description string) (*types.Transaction, error) {
	return _Property.Contract.Add(&_Property.TransactOpts, _name, _value, _description)
}

// ApplyUpdate is a paid mutator transaction binding the contract method 0xa0f2379c.
//
// Solidity: function applyUpdate(string _name, uint256 _value, string _reason) returns()
func (_Property *PropertyTransactor) ApplyUpdate(opts *bind.TransactOpts, _name string, _value *big.Int, _reason string) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "applyUpdate", _name, _value, _reason)
}

// ApplyUpdate is a paid mutator transaction binding the contract method 0xa0f2379c.
//
// Solidity: function applyUpdate(string _name, uint256 _value, string _reason) returns()
func (_Property *PropertySession) ApplyUpdate(_name string, _value *big.Int, _reason string) (*types.Transaction, error) {
	return _Property.Contract.ApplyUpdate(&_Property.TransactOpts, _name, _value, _reason)
}

// ApplyUpdate is a paid mutator transaction binding the contract method 0xa0f2379c.
//
// Solidity: function applyUpdate(string _name, uint256 _value, string _reason) returns()
func (_Property *PropertyTransactorSession) ApplyUpdate(_name string, _value *big.Int, _reason string) (*types.Transaction, error) {
	return _Property.Contract.ApplyUpdate(&_Property.TransactOpts, _name, _value, _reason)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Property *PropertyTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Property *PropertySession) Initialize() (*types.Transaction, error) {
	return _Property.Contract.Initialize(&_Property.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Property *PropertyTransactorSession) Initialize() (*types.Transaction, error) {
	return _Property.Contract.Initialize(&_Property.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Property *PropertyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Property *PropertySession) RenounceOwnership() (*types.Transaction, error) {
	return _Property.Contract.RenounceOwnership(&_Property.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Property *PropertyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Property.Contract.RenounceOwnership(&_Property.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Property *PropertyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Property *PropertySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Property.Contract.TransferOwnership(&_Property.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Property *PropertyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Property.Contract.TransferOwnership(&_Property.TransactOpts, newOwner)
}

// Vote4Update is a paid mutator transaction binding the contract method 0xa15b0d52.
//
// Solidity: function vote4Update(string _name, uint256 _voteResult) returns()
func (_Property *PropertyTransactor) Vote4Update(opts *bind.TransactOpts, _name string, _voteResult *big.Int) (*types.Transaction, error) {
	return _Property.contract.Transact(opts, "vote4Update", _name, _voteResult)
}

// Vote4Update is a paid mutator transaction binding the contract method 0xa15b0d52.
//
// Solidity: function vote4Update(string _name, uint256 _voteResult) returns()
func (_Property *PropertySession) Vote4Update(_name string, _voteResult *big.Int) (*types.Transaction, error) {
	return _Property.Contract.Vote4Update(&_Property.TransactOpts, _name, _voteResult)
}

// Vote4Update is a paid mutator transaction binding the contract method 0xa15b0d52.
//
// Solidity: function vote4Update(string _name, uint256 _voteResult) returns()
func (_Property *PropertyTransactorSession) Vote4Update(_name string, _voteResult *big.Int) (*types.Transaction, error) {
	return _Property.Contract.Vote4Update(&_Property.TransactOpts, _name, _voteResult)
}

// PropertyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Property contract.
type PropertyInitializedIterator struct {
	Event *PropertyInitialized // Event containing the contract specifics and raw log

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
func (it *PropertyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyInitialized)
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
		it.Event = new(PropertyInitialized)
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
func (it *PropertyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyInitialized represents a Initialized event raised by the Property contract.
type PropertyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Property *PropertyFilterer) FilterInitialized(opts *bind.FilterOpts) (*PropertyInitializedIterator, error) {

	logs, sub, err := _Property.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PropertyInitializedIterator{contract: _Property.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Property *PropertyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PropertyInitialized) (event.Subscription, error) {

	logs, sub, err := _Property.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyInitialized)
				if err := _Property.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Property *PropertyFilterer) ParseInitialized(log types.Log) (*PropertyInitialized, error) {
	event := new(PropertyInitialized)
	if err := _Property.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Property contract.
type PropertyOwnershipTransferredIterator struct {
	Event *PropertyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PropertyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyOwnershipTransferred)
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
		it.Event = new(PropertyOwnershipTransferred)
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
func (it *PropertyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyOwnershipTransferred represents a OwnershipTransferred event raised by the Property contract.
type PropertyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Property *PropertyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PropertyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Property.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PropertyOwnershipTransferredIterator{contract: _Property.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Property *PropertyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PropertyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Property.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyOwnershipTransferred)
				if err := _Property.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Property *PropertyFilterer) ParseOwnershipTransferred(log types.Log) (*PropertyOwnershipTransferred, error) {
	event := new(PropertyOwnershipTransferred)
	if err := _Property.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyPropertyAddIterator is returned from FilterPropertyAdd and is used to iterate over the raw logs and unpacked data for PropertyAdd events raised by the Property contract.
type PropertyPropertyAddIterator struct {
	Event *PropertyPropertyAdd // Event containing the contract specifics and raw log

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
func (it *PropertyPropertyAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyPropertyAdd)
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
		it.Event = new(PropertyPropertyAdd)
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
func (it *PropertyPropertyAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyPropertyAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyPropertyAdd represents a PropertyAdd event raised by the Property contract.
type PropertyPropertyAdd struct {
	Name  string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPropertyAdd is a free log retrieval operation binding the contract event 0x03cd47ea9d0993f3697def99757c19581e28a706ca3b8c60578551fb51aee8c5.
//
// Solidity: event PropertyAdd(string _name, uint256 _value)
func (_Property *PropertyFilterer) FilterPropertyAdd(opts *bind.FilterOpts) (*PropertyPropertyAddIterator, error) {

	logs, sub, err := _Property.contract.FilterLogs(opts, "PropertyAdd")
	if err != nil {
		return nil, err
	}
	return &PropertyPropertyAddIterator{contract: _Property.contract, event: "PropertyAdd", logs: logs, sub: sub}, nil
}

// WatchPropertyAdd is a free log subscription operation binding the contract event 0x03cd47ea9d0993f3697def99757c19581e28a706ca3b8c60578551fb51aee8c5.
//
// Solidity: event PropertyAdd(string _name, uint256 _value)
func (_Property *PropertyFilterer) WatchPropertyAdd(opts *bind.WatchOpts, sink chan<- *PropertyPropertyAdd) (event.Subscription, error) {

	logs, sub, err := _Property.contract.WatchLogs(opts, "PropertyAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyPropertyAdd)
				if err := _Property.contract.UnpackLog(event, "PropertyAdd", log); err != nil {
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

// ParsePropertyAdd is a log parse operation binding the contract event 0x03cd47ea9d0993f3697def99757c19581e28a706ca3b8c60578551fb51aee8c5.
//
// Solidity: event PropertyAdd(string _name, uint256 _value)
func (_Property *PropertyFilterer) ParsePropertyAdd(log types.Log) (*PropertyPropertyAdd, error) {
	event := new(PropertyPropertyAdd)
	if err := _Property.contract.UnpackLog(event, "PropertyAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyPropertyApplyUpdateIterator is returned from FilterPropertyApplyUpdate and is used to iterate over the raw logs and unpacked data for PropertyApplyUpdate events raised by the Property contract.
type PropertyPropertyApplyUpdateIterator struct {
	Event *PropertyPropertyApplyUpdate // Event containing the contract specifics and raw log

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
func (it *PropertyPropertyApplyUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyPropertyApplyUpdate)
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
		it.Event = new(PropertyPropertyApplyUpdate)
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
func (it *PropertyPropertyApplyUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyPropertyApplyUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyPropertyApplyUpdate represents a PropertyApplyUpdate event raised by the Property contract.
type PropertyPropertyApplyUpdate struct {
	Name     string
	NewValue *big.Int
	OldValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPropertyApplyUpdate is a free log retrieval operation binding the contract event 0x3ff2f2225de641a4f1c78be010ee1229d36b11b29d3d124b741db0a7bdec0e15.
//
// Solidity: event PropertyApplyUpdate(string _name, uint256 _newValue, uint256 _oldValue)
func (_Property *PropertyFilterer) FilterPropertyApplyUpdate(opts *bind.FilterOpts) (*PropertyPropertyApplyUpdateIterator, error) {

	logs, sub, err := _Property.contract.FilterLogs(opts, "PropertyApplyUpdate")
	if err != nil {
		return nil, err
	}
	return &PropertyPropertyApplyUpdateIterator{contract: _Property.contract, event: "PropertyApplyUpdate", logs: logs, sub: sub}, nil
}

// WatchPropertyApplyUpdate is a free log subscription operation binding the contract event 0x3ff2f2225de641a4f1c78be010ee1229d36b11b29d3d124b741db0a7bdec0e15.
//
// Solidity: event PropertyApplyUpdate(string _name, uint256 _newValue, uint256 _oldValue)
func (_Property *PropertyFilterer) WatchPropertyApplyUpdate(opts *bind.WatchOpts, sink chan<- *PropertyPropertyApplyUpdate) (event.Subscription, error) {

	logs, sub, err := _Property.contract.WatchLogs(opts, "PropertyApplyUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyPropertyApplyUpdate)
				if err := _Property.contract.UnpackLog(event, "PropertyApplyUpdate", log); err != nil {
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

// ParsePropertyApplyUpdate is a log parse operation binding the contract event 0x3ff2f2225de641a4f1c78be010ee1229d36b11b29d3d124b741db0a7bdec0e15.
//
// Solidity: event PropertyApplyUpdate(string _name, uint256 _newValue, uint256 _oldValue)
func (_Property *PropertyFilterer) ParsePropertyApplyUpdate(log types.Log) (*PropertyPropertyApplyUpdate, error) {
	event := new(PropertyPropertyApplyUpdate)
	if err := _Property.contract.UnpackLog(event, "PropertyApplyUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyPropertyUpdateAgreeIterator is returned from FilterPropertyUpdateAgree and is used to iterate over the raw logs and unpacked data for PropertyUpdateAgree events raised by the Property contract.
type PropertyPropertyUpdateAgreeIterator struct {
	Event *PropertyPropertyUpdateAgree // Event containing the contract specifics and raw log

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
func (it *PropertyPropertyUpdateAgreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyPropertyUpdateAgree)
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
		it.Event = new(PropertyPropertyUpdateAgree)
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
func (it *PropertyPropertyUpdateAgreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyPropertyUpdateAgreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyPropertyUpdateAgree represents a PropertyUpdateAgree event raised by the Property contract.
type PropertyPropertyUpdateAgree struct {
	Name     string
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPropertyUpdateAgree is a free log retrieval operation binding the contract event 0x41e98b278cb6e3d11bdd2d89f0e73284bc731871dd52a7c1f36b397268d0e3c5.
//
// Solidity: event PropertyUpdateAgree(string _name, uint256 _newValue)
func (_Property *PropertyFilterer) FilterPropertyUpdateAgree(opts *bind.FilterOpts) (*PropertyPropertyUpdateAgreeIterator, error) {

	logs, sub, err := _Property.contract.FilterLogs(opts, "PropertyUpdateAgree")
	if err != nil {
		return nil, err
	}
	return &PropertyPropertyUpdateAgreeIterator{contract: _Property.contract, event: "PropertyUpdateAgree", logs: logs, sub: sub}, nil
}

// WatchPropertyUpdateAgree is a free log subscription operation binding the contract event 0x41e98b278cb6e3d11bdd2d89f0e73284bc731871dd52a7c1f36b397268d0e3c5.
//
// Solidity: event PropertyUpdateAgree(string _name, uint256 _newValue)
func (_Property *PropertyFilterer) WatchPropertyUpdateAgree(opts *bind.WatchOpts, sink chan<- *PropertyPropertyUpdateAgree) (event.Subscription, error) {

	logs, sub, err := _Property.contract.WatchLogs(opts, "PropertyUpdateAgree")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyPropertyUpdateAgree)
				if err := _Property.contract.UnpackLog(event, "PropertyUpdateAgree", log); err != nil {
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

// ParsePropertyUpdateAgree is a log parse operation binding the contract event 0x41e98b278cb6e3d11bdd2d89f0e73284bc731871dd52a7c1f36b397268d0e3c5.
//
// Solidity: event PropertyUpdateAgree(string _name, uint256 _newValue)
func (_Property *PropertyFilterer) ParsePropertyUpdateAgree(log types.Log) (*PropertyPropertyUpdateAgree, error) {
	event := new(PropertyPropertyUpdateAgree)
	if err := _Property.contract.UnpackLog(event, "PropertyUpdateAgree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyPropertyUpdateRejectIterator is returned from FilterPropertyUpdateReject and is used to iterate over the raw logs and unpacked data for PropertyUpdateReject events raised by the Property contract.
type PropertyPropertyUpdateRejectIterator struct {
	Event *PropertyPropertyUpdateReject // Event containing the contract specifics and raw log

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
func (it *PropertyPropertyUpdateRejectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyPropertyUpdateReject)
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
		it.Event = new(PropertyPropertyUpdateReject)
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
func (it *PropertyPropertyUpdateRejectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyPropertyUpdateRejectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyPropertyUpdateReject represents a PropertyUpdateReject event raised by the Property contract.
type PropertyPropertyUpdateReject struct {
	Name     string
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPropertyUpdateReject is a free log retrieval operation binding the contract event 0x85e07acf81367fcdbc2186e9e7392b10042a5bd46f5c98ea8d92cbe331aa8a73.
//
// Solidity: event PropertyUpdateReject(string _name, uint256 _newValue)
func (_Property *PropertyFilterer) FilterPropertyUpdateReject(opts *bind.FilterOpts) (*PropertyPropertyUpdateRejectIterator, error) {

	logs, sub, err := _Property.contract.FilterLogs(opts, "PropertyUpdateReject")
	if err != nil {
		return nil, err
	}
	return &PropertyPropertyUpdateRejectIterator{contract: _Property.contract, event: "PropertyUpdateReject", logs: logs, sub: sub}, nil
}

// WatchPropertyUpdateReject is a free log subscription operation binding the contract event 0x85e07acf81367fcdbc2186e9e7392b10042a5bd46f5c98ea8d92cbe331aa8a73.
//
// Solidity: event PropertyUpdateReject(string _name, uint256 _newValue)
func (_Property *PropertyFilterer) WatchPropertyUpdateReject(opts *bind.WatchOpts, sink chan<- *PropertyPropertyUpdateReject) (event.Subscription, error) {

	logs, sub, err := _Property.contract.WatchLogs(opts, "PropertyUpdateReject")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyPropertyUpdateReject)
				if err := _Property.contract.UnpackLog(event, "PropertyUpdateReject", log); err != nil {
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

// ParsePropertyUpdateReject is a log parse operation binding the contract event 0x85e07acf81367fcdbc2186e9e7392b10042a5bd46f5c98ea8d92cbe331aa8a73.
//
// Solidity: event PropertyUpdateReject(string _name, uint256 _newValue)
func (_Property *PropertyFilterer) ParsePropertyUpdateReject(log types.Log) (*PropertyPropertyUpdateReject, error) {
	event := new(PropertyPropertyUpdateReject)
	if err := _Property.contract.UnpackLog(event, "PropertyUpdateReject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyPropertyUpdateVoteIterator is returned from FilterPropertyUpdateVote and is used to iterate over the raw logs and unpacked data for PropertyUpdateVote events raised by the Property contract.
type PropertyPropertyUpdateVoteIterator struct {
	Event *PropertyPropertyUpdateVote // Event containing the contract specifics and raw log

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
func (it *PropertyPropertyUpdateVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyPropertyUpdateVote)
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
		it.Event = new(PropertyPropertyUpdateVote)
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
func (it *PropertyPropertyUpdateVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyPropertyUpdateVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyPropertyUpdateVote represents a PropertyUpdateVote event raised by the Property contract.
type PropertyPropertyUpdateVote struct {
	Name       string
	NewValue   *big.Int
	Voter      common.Address
	VoteResult *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPropertyUpdateVote is a free log retrieval operation binding the contract event 0x9887df82232047c477af2f6734744bdf006cd6694dc3a931b66b1b7a71b586a4.
//
// Solidity: event PropertyUpdateVote(string _name, uint256 _newValue, address _voter, uint256 _voteResult)
func (_Property *PropertyFilterer) FilterPropertyUpdateVote(opts *bind.FilterOpts) (*PropertyPropertyUpdateVoteIterator, error) {

	logs, sub, err := _Property.contract.FilterLogs(opts, "PropertyUpdateVote")
	if err != nil {
		return nil, err
	}
	return &PropertyPropertyUpdateVoteIterator{contract: _Property.contract, event: "PropertyUpdateVote", logs: logs, sub: sub}, nil
}

// WatchPropertyUpdateVote is a free log subscription operation binding the contract event 0x9887df82232047c477af2f6734744bdf006cd6694dc3a931b66b1b7a71b586a4.
//
// Solidity: event PropertyUpdateVote(string _name, uint256 _newValue, address _voter, uint256 _voteResult)
func (_Property *PropertyFilterer) WatchPropertyUpdateVote(opts *bind.WatchOpts, sink chan<- *PropertyPropertyUpdateVote) (event.Subscription, error) {

	logs, sub, err := _Property.contract.WatchLogs(opts, "PropertyUpdateVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyPropertyUpdateVote)
				if err := _Property.contract.UnpackLog(event, "PropertyUpdateVote", log); err != nil {
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

// ParsePropertyUpdateVote is a log parse operation binding the contract event 0x9887df82232047c477af2f6734744bdf006cd6694dc3a931b66b1b7a71b586a4.
//
// Solidity: event PropertyUpdateVote(string _name, uint256 _newValue, address _voter, uint256 _voteResult)
func (_Property *PropertyFilterer) ParsePropertyUpdateVote(log types.Log) (*PropertyPropertyUpdateVote, error) {
	event := new(PropertyPropertyUpdateVote)
	if err := _Property.contract.UnpackLog(event, "PropertyUpdateVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
