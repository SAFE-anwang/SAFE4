// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NodeState

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

// NodeStateMetaData contains all meta data concerning the NodeState contract.
var NodeStateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_states\",\"type\":\"uint8[]\"}],\"name\":\"uploadState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NodeStateABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeStateMetaData.ABI instead.
var NodeStateABI = NodeStateMetaData.ABI

// NodeState is an auto generated Go binding around an Ethereum contract.
type NodeState struct {
	NodeStateCaller     // Read-only binding to the contract
	NodeStateTransactor // Write-only binding to the contract
	NodeStateFilterer   // Log filterer for contract events
}

// NodeStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeStateSession struct {
	Contract     *NodeState        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeStateCallerSession struct {
	Contract *NodeStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NodeStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeStateTransactorSession struct {
	Contract     *NodeStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NodeStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeStateRaw struct {
	Contract *NodeState // Generic contract binding to access the raw methods on
}

// NodeStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeStateCallerRaw struct {
	Contract *NodeStateCaller // Generic read-only contract binding to access the raw methods on
}

// NodeStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeStateTransactorRaw struct {
	Contract *NodeStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeState creates a new instance of NodeState, bound to a specific deployed contract.
func NewNodeState(address common.Address, backend bind.ContractBackend) (*NodeState, error) {
	contract, err := bindNodeState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeState{NodeStateCaller: NodeStateCaller{contract: contract}, NodeStateTransactor: NodeStateTransactor{contract: contract}, NodeStateFilterer: NodeStateFilterer{contract: contract}}, nil
}

// NewNodeStateCaller creates a new read-only instance of NodeState, bound to a specific deployed contract.
func NewNodeStateCaller(address common.Address, caller bind.ContractCaller) (*NodeStateCaller, error) {
	contract, err := bindNodeState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeStateCaller{contract: contract}, nil
}

// NewNodeStateTransactor creates a new write-only instance of NodeState, bound to a specific deployed contract.
func NewNodeStateTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeStateTransactor, error) {
	contract, err := bindNodeState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeStateTransactor{contract: contract}, nil
}

// NewNodeStateFilterer creates a new log filterer instance of NodeState, bound to a specific deployed contract.
func NewNodeStateFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeStateFilterer, error) {
	contract, err := bindNodeState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeStateFilterer{contract: contract}, nil
}

// bindNodeState binds a generic wrapper to an already deployed contract.
func bindNodeState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeState *NodeStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeState.Contract.NodeStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeState *NodeStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeState.Contract.NodeStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeState *NodeStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeState.Contract.NodeStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeState *NodeStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeState *NodeStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeState *NodeStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeState.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_NodeState *NodeStateCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _NodeState.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_NodeState *NodeStateSession) GetInitializeData() ([]byte, error) {
	return _NodeState.Contract.GetInitializeData(&_NodeState.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_NodeState *NodeStateCallerSession) GetInitializeData() ([]byte, error) {
	return _NodeState.Contract.GetInitializeData(&_NodeState.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeState *NodeStateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeState.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeState *NodeStateSession) Owner() (common.Address, error) {
	return _NodeState.Contract.Owner(&_NodeState.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeState *NodeStateCallerSession) Owner() (common.Address, error) {
	return _NodeState.Contract.Owner(&_NodeState.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NodeState *NodeStateTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeState.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NodeState *NodeStateSession) Initialize() (*types.Transaction, error) {
	return _NodeState.Contract.Initialize(&_NodeState.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NodeState *NodeStateTransactorSession) Initialize() (*types.Transaction, error) {
	return _NodeState.Contract.Initialize(&_NodeState.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeState *NodeStateTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeState.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeState *NodeStateSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeState.Contract.RenounceOwnership(&_NodeState.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeState *NodeStateTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeState.Contract.RenounceOwnership(&_NodeState.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeState *NodeStateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NodeState.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeState *NodeStateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeState.Contract.TransferOwnership(&_NodeState.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeState *NodeStateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeState.Contract.TransferOwnership(&_NodeState.TransactOpts, newOwner)
}

// UploadState is a paid mutator transaction binding the contract method 0x27676765.
//
// Solidity: function uploadState(uint256[] _ids, uint8[] _states) returns()
func (_NodeState *NodeStateTransactor) UploadState(opts *bind.TransactOpts, _ids []*big.Int, _states []uint8) (*types.Transaction, error) {
	return _NodeState.contract.Transact(opts, "uploadState", _ids, _states)
}

// UploadState is a paid mutator transaction binding the contract method 0x27676765.
//
// Solidity: function uploadState(uint256[] _ids, uint8[] _states) returns()
func (_NodeState *NodeStateSession) UploadState(_ids []*big.Int, _states []uint8) (*types.Transaction, error) {
	return _NodeState.Contract.UploadState(&_NodeState.TransactOpts, _ids, _states)
}

// UploadState is a paid mutator transaction binding the contract method 0x27676765.
//
// Solidity: function uploadState(uint256[] _ids, uint8[] _states) returns()
func (_NodeState *NodeStateTransactorSession) UploadState(_ids []*big.Int, _states []uint8) (*types.Transaction, error) {
	return _NodeState.Contract.UploadState(&_NodeState.TransactOpts, _ids, _states)
}

// NodeStateInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NodeState contract.
type NodeStateInitializedIterator struct {
	Event *NodeStateInitialized // Event containing the contract specifics and raw log

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
func (it *NodeStateInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStateInitialized)
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
		it.Event = new(NodeStateInitialized)
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
func (it *NodeStateInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStateInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStateInitialized represents a Initialized event raised by the NodeState contract.
type NodeStateInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeState *NodeStateFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodeStateInitializedIterator, error) {

	logs, sub, err := _NodeState.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodeStateInitializedIterator{contract: _NodeState.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeState *NodeStateFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodeStateInitialized) (event.Subscription, error) {

	logs, sub, err := _NodeState.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStateInitialized)
				if err := _NodeState.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NodeState *NodeStateFilterer) ParseInitialized(log types.Log) (*NodeStateInitialized, error) {
	event := new(NodeStateInitialized)
	if err := _NodeState.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NodeState contract.
type NodeStateOwnershipTransferredIterator struct {
	Event *NodeStateOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeStateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStateOwnershipTransferred)
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
		it.Event = new(NodeStateOwnershipTransferred)
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
func (it *NodeStateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStateOwnershipTransferred represents a OwnershipTransferred event raised by the NodeState contract.
type NodeStateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeState *NodeStateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeStateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeState.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeStateOwnershipTransferredIterator{contract: _NodeState.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeState *NodeStateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeStateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeState.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStateOwnershipTransferred)
				if err := _NodeState.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NodeState *NodeStateFilterer) ParseOwnershipTransferred(log types.Log) (*NodeStateOwnershipTransferred, error) {
	event := new(NodeStateOwnershipTransferred)
	if err := _NodeState.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
