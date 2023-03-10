// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package System

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

// SystemMetaData contains all meta data concerning the System contract.
var SystemMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SystemABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemMetaData.ABI instead.
var SystemABI = SystemMetaData.ABI

// System is an auto generated Go binding around an Ethereum contract.
type System struct {
	SystemCaller     // Read-only binding to the contract
	SystemTransactor // Write-only binding to the contract
	SystemFilterer   // Log filterer for contract events
}

// SystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemSession struct {
	Contract     *System           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemCallerSession struct {
	Contract *SystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemTransactorSession struct {
	Contract     *SystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRaw struct {
	Contract *System // Generic contract binding to access the raw methods on
}

// SystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemCallerRaw struct {
	Contract *SystemCaller // Generic read-only contract binding to access the raw methods on
}

// SystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemTransactorRaw struct {
	Contract *SystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystem creates a new instance of System, bound to a specific deployed contract.
func NewSystem(address common.Address, backend bind.ContractBackend) (*System, error) {
	contract, err := bindSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// NewSystemCaller creates a new read-only instance of System, bound to a specific deployed contract.
func NewSystemCaller(address common.Address, caller bind.ContractCaller) (*SystemCaller, error) {
	contract, err := bindSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemCaller{contract: contract}, nil
}

// NewSystemTransactor creates a new write-only instance of System, bound to a specific deployed contract.
func NewSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemTransactor, error) {
	contract, err := bindSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemTransactor{contract: contract}, nil
}

// NewSystemFilterer creates a new log filterer instance of System, bound to a specific deployed contract.
func NewSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemFilterer, error) {
	contract, err := bindSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemFilterer{contract: contract}, nil
}

// bindSystem binds a generic wrapper to an already deployed contract.
func bindSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.SystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_System *SystemCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_System *SystemSession) GetInitializeData() ([]byte, error) {
	return _System.Contract.GetInitializeData(&_System.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_System *SystemCallerSession) GetInitializeData() ([]byte, error) {
	return _System.Contract.GetInitializeData(&_System.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_System *SystemCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_System *SystemSession) Owner() (common.Address, error) {
	return _System.Contract.Owner(&_System.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_System *SystemCallerSession) Owner() (common.Address, error) {
	return _System.Contract.Owner(&_System.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_System *SystemTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_System *SystemSession) Initialize() (*types.Transaction, error) {
	return _System.Contract.Initialize(&_System.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_System *SystemTransactorSession) Initialize() (*types.Transaction, error) {
	return _System.Contract.Initialize(&_System.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_System *SystemTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_System *SystemSession) RenounceOwnership() (*types.Transaction, error) {
	return _System.Contract.RenounceOwnership(&_System.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_System *SystemTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _System.Contract.RenounceOwnership(&_System.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_System *SystemTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _System.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_System *SystemSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _System.Contract.TransferOwnership(&_System.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_System *SystemTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _System.Contract.TransferOwnership(&_System.TransactOpts, newOwner)
}

// SystemInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the System contract.
type SystemInitializedIterator struct {
	Event *SystemInitialized // Event containing the contract specifics and raw log

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
func (it *SystemInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemInitialized)
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
		it.Event = new(SystemInitialized)
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
func (it *SystemInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemInitialized represents a Initialized event raised by the System contract.
type SystemInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_System *SystemFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemInitializedIterator, error) {

	logs, sub, err := _System.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemInitializedIterator{contract: _System.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_System *SystemFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemInitialized) (event.Subscription, error) {

	logs, sub, err := _System.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemInitialized)
				if err := _System.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_System *SystemFilterer) ParseInitialized(log types.Log) (*SystemInitialized, error) {
	event := new(SystemInitialized)
	if err := _System.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the System contract.
type SystemOwnershipTransferredIterator struct {
	Event *SystemOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemOwnershipTransferred)
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
		it.Event = new(SystemOwnershipTransferred)
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
func (it *SystemOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemOwnershipTransferred represents a OwnershipTransferred event raised by the System contract.
type SystemOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_System *SystemFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _System.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemOwnershipTransferredIterator{contract: _System.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_System *SystemFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _System.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemOwnershipTransferred)
				if err := _System.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_System *SystemFilterer) ParseOwnershipTransferred(log types.Log) (*SystemOwnershipTransferred, error) {
	event := new(SystemOwnershipTransferred)
	if err := _System.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
