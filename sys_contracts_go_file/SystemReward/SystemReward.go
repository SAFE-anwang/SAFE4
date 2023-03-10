// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SystemReward

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

// SystemRewardMetaData contains all meta data concerning the SystemReward contract.
var SystemRewardMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_smnAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_smnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mnAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mnAmount\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SystemRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemRewardMetaData.ABI instead.
var SystemRewardABI = SystemRewardMetaData.ABI

// SystemReward is an auto generated Go binding around an Ethereum contract.
type SystemReward struct {
	SystemRewardCaller     // Read-only binding to the contract
	SystemRewardTransactor // Write-only binding to the contract
	SystemRewardFilterer   // Log filterer for contract events
}

// SystemRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemRewardSession struct {
	Contract     *SystemReward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemRewardCallerSession struct {
	Contract *SystemRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SystemRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemRewardTransactorSession struct {
	Contract     *SystemRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SystemRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRewardRaw struct {
	Contract *SystemReward // Generic contract binding to access the raw methods on
}

// SystemRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemRewardCallerRaw struct {
	Contract *SystemRewardCaller // Generic read-only contract binding to access the raw methods on
}

// SystemRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemRewardTransactorRaw struct {
	Contract *SystemRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemReward creates a new instance of SystemReward, bound to a specific deployed contract.
func NewSystemReward(address common.Address, backend bind.ContractBackend) (*SystemReward, error) {
	contract, err := bindSystemReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemReward{SystemRewardCaller: SystemRewardCaller{contract: contract}, SystemRewardTransactor: SystemRewardTransactor{contract: contract}, SystemRewardFilterer: SystemRewardFilterer{contract: contract}}, nil
}

// NewSystemRewardCaller creates a new read-only instance of SystemReward, bound to a specific deployed contract.
func NewSystemRewardCaller(address common.Address, caller bind.ContractCaller) (*SystemRewardCaller, error) {
	contract, err := bindSystemReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRewardCaller{contract: contract}, nil
}

// NewSystemRewardTransactor creates a new write-only instance of SystemReward, bound to a specific deployed contract.
func NewSystemRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemRewardTransactor, error) {
	contract, err := bindSystemReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRewardTransactor{contract: contract}, nil
}

// NewSystemRewardFilterer creates a new log filterer instance of SystemReward, bound to a specific deployed contract.
func NewSystemRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemRewardFilterer, error) {
	contract, err := bindSystemReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemRewardFilterer{contract: contract}, nil
}

// bindSystemReward binds a generic wrapper to an already deployed contract.
func bindSystemReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemReward *SystemRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemReward.Contract.SystemRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemReward *SystemRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReward.Contract.SystemRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemReward *SystemRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemReward.Contract.SystemRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemReward *SystemRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemReward *SystemRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemReward *SystemRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemReward.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SystemReward *SystemRewardCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SystemReward.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SystemReward *SystemRewardSession) GetInitializeData() ([]byte, error) {
	return _SystemReward.Contract.GetInitializeData(&_SystemReward.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SystemReward *SystemRewardCallerSession) GetInitializeData() ([]byte, error) {
	return _SystemReward.Contract.GetInitializeData(&_SystemReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemReward *SystemRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemReward *SystemRewardSession) Owner() (common.Address, error) {
	return _SystemReward.Contract.Owner(&_SystemReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemReward *SystemRewardCallerSession) Owner() (common.Address, error) {
	return _SystemReward.Contract.Owner(&_SystemReward.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SystemReward *SystemRewardTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReward.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SystemReward *SystemRewardSession) Initialize() (*types.Transaction, error) {
	return _SystemReward.Contract.Initialize(&_SystemReward.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SystemReward *SystemRewardTransactorSession) Initialize() (*types.Transaction, error) {
	return _SystemReward.Contract.Initialize(&_SystemReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemReward *SystemRewardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReward.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemReward *SystemRewardSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemReward.Contract.RenounceOwnership(&_SystemReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemReward *SystemRewardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemReward.Contract.RenounceOwnership(&_SystemReward.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x26240f10.
//
// Solidity: function reward(address _smnAddr, uint256 _smnAmount, address _mnAddr, uint256 _mnAmount) payable returns()
func (_SystemReward *SystemRewardTransactor) Reward(opts *bind.TransactOpts, _smnAddr common.Address, _smnAmount *big.Int, _mnAddr common.Address, _mnAmount *big.Int) (*types.Transaction, error) {
	return _SystemReward.contract.Transact(opts, "reward", _smnAddr, _smnAmount, _mnAddr, _mnAmount)
}

// Reward is a paid mutator transaction binding the contract method 0x26240f10.
//
// Solidity: function reward(address _smnAddr, uint256 _smnAmount, address _mnAddr, uint256 _mnAmount) payable returns()
func (_SystemReward *SystemRewardSession) Reward(_smnAddr common.Address, _smnAmount *big.Int, _mnAddr common.Address, _mnAmount *big.Int) (*types.Transaction, error) {
	return _SystemReward.Contract.Reward(&_SystemReward.TransactOpts, _smnAddr, _smnAmount, _mnAddr, _mnAmount)
}

// Reward is a paid mutator transaction binding the contract method 0x26240f10.
//
// Solidity: function reward(address _smnAddr, uint256 _smnAmount, address _mnAddr, uint256 _mnAmount) payable returns()
func (_SystemReward *SystemRewardTransactorSession) Reward(_smnAddr common.Address, _smnAmount *big.Int, _mnAddr common.Address, _mnAmount *big.Int) (*types.Transaction, error) {
	return _SystemReward.Contract.Reward(&_SystemReward.TransactOpts, _smnAddr, _smnAmount, _mnAddr, _mnAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemReward *SystemRewardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemReward.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemReward *SystemRewardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemReward.Contract.TransferOwnership(&_SystemReward.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemReward *SystemRewardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemReward.Contract.TransferOwnership(&_SystemReward.TransactOpts, newOwner)
}

// SystemRewardInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemReward contract.
type SystemRewardInitializedIterator struct {
	Event *SystemRewardInitialized // Event containing the contract specifics and raw log

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
func (it *SystemRewardInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRewardInitialized)
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
		it.Event = new(SystemRewardInitialized)
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
func (it *SystemRewardInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRewardInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRewardInitialized represents a Initialized event raised by the SystemReward contract.
type SystemRewardInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemReward *SystemRewardFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemRewardInitializedIterator, error) {

	logs, sub, err := _SystemReward.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemRewardInitializedIterator{contract: _SystemReward.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemReward *SystemRewardFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemRewardInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemReward.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRewardInitialized)
				if err := _SystemReward.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SystemReward *SystemRewardFilterer) ParseInitialized(log types.Log) (*SystemRewardInitialized, error) {
	event := new(SystemRewardInitialized)
	if err := _SystemReward.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRewardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemReward contract.
type SystemRewardOwnershipTransferredIterator struct {
	Event *SystemRewardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemRewardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRewardOwnershipTransferred)
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
		it.Event = new(SystemRewardOwnershipTransferred)
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
func (it *SystemRewardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRewardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRewardOwnershipTransferred represents a OwnershipTransferred event raised by the SystemReward contract.
type SystemRewardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemReward *SystemRewardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemRewardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemReward.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemRewardOwnershipTransferredIterator{contract: _SystemReward.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemReward *SystemRewardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemRewardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemReward.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRewardOwnershipTransferred)
				if err := _SystemReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SystemReward *SystemRewardFilterer) ParseOwnershipTransferred(log types.Log) (*SystemRewardOwnershipTransferred, error) {
	event := new(SystemRewardOwnershipTransferred)
	if err := _SystemReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
