// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AccountManager

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

// IAccountManagerAccountRecord is an auto generated low-level Go binding around an user-defined struct.
type IAccountManagerAccountRecord struct {
	Id           [20]byte
	Addr         common.Address
	Amount       *big.Int
	LockDay      *big.Int
	StartHeight  *big.Int
	UnlockHeight *big.Int
	BindInfo     IAccountManagerBindInfo
	CreateHeight *big.Int
	UpdateHeight *big.Int
}

// IAccountManagerBindInfo is an auto generated low-level Go binding around an user-defined struct.
type IAccountManagerBindInfo struct {
	BindHeight   *big.Int
	UnbindHeight *big.Int
}

// AccountManagerMetaData contains all meta data concerning the AccountManager contract.
var AccountManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_reocrdID\",\"type\":\"bytes20\"}],\"name\":\"SafeDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"}],\"name\":\"SafeTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SafeWithdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes20[]\",\"name\":\"\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBindAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes20[]\",\"name\":\"\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes20[]\",\"name\":\"\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"}],\"name\":\"getRecordByID\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes20\",\"name\":\"id\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bindHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbindHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIAccountManager.BindInfo\",\"name\":\"bindInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIAccountManager.AccountRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes20\",\"name\":\"id\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bindHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbindHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIAccountManager.BindInfo\",\"name\":\"bindInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIAccountManager.AccountRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes20[]\",\"name\":\"\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_rewardType\",\"type\":\"uint8\"}],\"name\":\"reward\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_bindDay\",\"type\":\"uint256\"}],\"name\":\"setBindDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockDay\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"_recordIDs\",\"type\":\"bytes20[]\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AccountManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountManagerMetaData.ABI instead.
var AccountManagerABI = AccountManagerMetaData.ABI

// AccountManager is an auto generated Go binding around an Ethereum contract.
type AccountManager struct {
	AccountManagerCaller     // Read-only binding to the contract
	AccountManagerTransactor // Write-only binding to the contract
	AccountManagerFilterer   // Log filterer for contract events
}

// AccountManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerSession struct {
	Contract     *AccountManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerCallerSession struct {
	Contract *AccountManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerTransactorSession struct {
	Contract     *AccountManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerRaw struct {
	Contract *AccountManager // Generic contract binding to access the raw methods on
}

// AccountManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerCallerRaw struct {
	Contract *AccountManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerTransactorRaw struct {
	Contract *AccountManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManager creates a new instance of AccountManager, bound to a specific deployed contract.
func NewAccountManager(address common.Address, backend bind.ContractBackend) (*AccountManager, error) {
	contract, err := bindAccountManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// NewAccountManagerCaller creates a new read-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerCaller, error) {
	contract, err := bindAccountManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerCaller{contract: contract}, nil
}

// NewAccountManagerTransactor creates a new write-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerTransactor, error) {
	contract, err := bindAccountManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerTransactor{contract: contract}, nil
}

// NewAccountManagerFilterer creates a new log filterer instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerFilterer, error) {
	contract, err := bindAccountManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFilterer{contract: contract}, nil
}

// bindAccountManager binds a generic wrapper to an already deployed contract.
func bindAccountManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.AccountManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_AccountManager *AccountManagerCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_AccountManager *AccountManagerSession) GetInitializeData() ([]byte, error) {
	return _AccountManager.Contract.GetInitializeData(&_AccountManager.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_AccountManager *AccountManagerCallerSession) GetInitializeData() ([]byte, error) {
	return _AccountManager.Contract.GetInitializeData(&_AccountManager.CallOpts)
}

// GetAvailableAmount is a free data retrieval call binding the contract method 0x7bb476f5.
//
// Solidity: function getAvailableAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerCaller) GetAvailableAmount(opts *bind.CallOpts) (*big.Int, [][20]byte, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getAvailableAmount")

	if err != nil {
		return *new(*big.Int), *new([][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][20]byte)).(*[][20]byte)

	return out0, out1, err

}

// GetAvailableAmount is a free data retrieval call binding the contract method 0x7bb476f5.
//
// Solidity: function getAvailableAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerSession) GetAvailableAmount() (*big.Int, [][20]byte, error) {
	return _AccountManager.Contract.GetAvailableAmount(&_AccountManager.CallOpts)
}

// GetAvailableAmount is a free data retrieval call binding the contract method 0x7bb476f5.
//
// Solidity: function getAvailableAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerCallerSession) GetAvailableAmount() (*big.Int, [][20]byte, error) {
	return _AccountManager.Contract.GetAvailableAmount(&_AccountManager.CallOpts)
}

// GetBindAmount is a free data retrieval call binding the contract method 0x1739b4b3.
//
// Solidity: function getBindAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerCaller) GetBindAmount(opts *bind.CallOpts) (*big.Int, [][20]byte, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getBindAmount")

	if err != nil {
		return *new(*big.Int), *new([][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][20]byte)).(*[][20]byte)

	return out0, out1, err

}

// GetBindAmount is a free data retrieval call binding the contract method 0x1739b4b3.
//
// Solidity: function getBindAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerSession) GetBindAmount() (*big.Int, [][20]byte, error) {
	return _AccountManager.Contract.GetBindAmount(&_AccountManager.CallOpts)
}

// GetBindAmount is a free data retrieval call binding the contract method 0x1739b4b3.
//
// Solidity: function getBindAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerCallerSession) GetBindAmount() (*big.Int, [][20]byte, error) {
	return _AccountManager.Contract.GetBindAmount(&_AccountManager.CallOpts)
}

// GetLockAmount is a free data retrieval call binding the contract method 0xd64c34fc.
//
// Solidity: function getLockAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerCaller) GetLockAmount(opts *bind.CallOpts) (*big.Int, [][20]byte, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getLockAmount")

	if err != nil {
		return *new(*big.Int), *new([][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][20]byte)).(*[][20]byte)

	return out0, out1, err

}

// GetLockAmount is a free data retrieval call binding the contract method 0xd64c34fc.
//
// Solidity: function getLockAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerSession) GetLockAmount() (*big.Int, [][20]byte, error) {
	return _AccountManager.Contract.GetLockAmount(&_AccountManager.CallOpts)
}

// GetLockAmount is a free data retrieval call binding the contract method 0xd64c34fc.
//
// Solidity: function getLockAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerCallerSession) GetLockAmount() (*big.Int, [][20]byte, error) {
	return _AccountManager.Contract.GetLockAmount(&_AccountManager.CallOpts)
}

// GetRecordByID is a free data retrieval call binding the contract method 0x2f468a14.
//
// Solidity: function getRecordByID(bytes20 _recordID) view returns((bytes20,address,uint256,uint256,uint256,uint256,(uint256,uint256),uint256,uint256))
func (_AccountManager *AccountManagerCaller) GetRecordByID(opts *bind.CallOpts, _recordID [20]byte) (IAccountManagerAccountRecord, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getRecordByID", _recordID)

	if err != nil {
		return *new(IAccountManagerAccountRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(IAccountManagerAccountRecord)).(*IAccountManagerAccountRecord)

	return out0, err

}

// GetRecordByID is a free data retrieval call binding the contract method 0x2f468a14.
//
// Solidity: function getRecordByID(bytes20 _recordID) view returns((bytes20,address,uint256,uint256,uint256,uint256,(uint256,uint256),uint256,uint256))
func (_AccountManager *AccountManagerSession) GetRecordByID(_recordID [20]byte) (IAccountManagerAccountRecord, error) {
	return _AccountManager.Contract.GetRecordByID(&_AccountManager.CallOpts, _recordID)
}

// GetRecordByID is a free data retrieval call binding the contract method 0x2f468a14.
//
// Solidity: function getRecordByID(bytes20 _recordID) view returns((bytes20,address,uint256,uint256,uint256,uint256,(uint256,uint256),uint256,uint256))
func (_AccountManager *AccountManagerCallerSession) GetRecordByID(_recordID [20]byte) (IAccountManagerAccountRecord, error) {
	return _AccountManager.Contract.GetRecordByID(&_AccountManager.CallOpts, _recordID)
}

// GetRecords is a free data retrieval call binding the contract method 0x8eff3c29.
//
// Solidity: function getRecords() view returns((bytes20,address,uint256,uint256,uint256,uint256,(uint256,uint256),uint256,uint256)[])
func (_AccountManager *AccountManagerCaller) GetRecords(opts *bind.CallOpts) ([]IAccountManagerAccountRecord, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getRecords")

	if err != nil {
		return *new([]IAccountManagerAccountRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountManagerAccountRecord)).(*[]IAccountManagerAccountRecord)

	return out0, err

}

// GetRecords is a free data retrieval call binding the contract method 0x8eff3c29.
//
// Solidity: function getRecords() view returns((bytes20,address,uint256,uint256,uint256,uint256,(uint256,uint256),uint256,uint256)[])
func (_AccountManager *AccountManagerSession) GetRecords() ([]IAccountManagerAccountRecord, error) {
	return _AccountManager.Contract.GetRecords(&_AccountManager.CallOpts)
}

// GetRecords is a free data retrieval call binding the contract method 0x8eff3c29.
//
// Solidity: function getRecords() view returns((bytes20,address,uint256,uint256,uint256,uint256,(uint256,uint256),uint256,uint256)[])
func (_AccountManager *AccountManagerCallerSession) GetRecords() ([]IAccountManagerAccountRecord, error) {
	return _AccountManager.Contract.GetRecords(&_AccountManager.CallOpts)
}

// GetTotalAmount is a free data retrieval call binding the contract method 0x65ac4341.
//
// Solidity: function getTotalAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerCaller) GetTotalAmount(opts *bind.CallOpts) (*big.Int, [][20]byte, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getTotalAmount")

	if err != nil {
		return *new(*big.Int), *new([][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][20]byte)).(*[][20]byte)

	return out0, out1, err

}

// GetTotalAmount is a free data retrieval call binding the contract method 0x65ac4341.
//
// Solidity: function getTotalAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerSession) GetTotalAmount() (*big.Int, [][20]byte, error) {
	return _AccountManager.Contract.GetTotalAmount(&_AccountManager.CallOpts)
}

// GetTotalAmount is a free data retrieval call binding the contract method 0x65ac4341.
//
// Solidity: function getTotalAmount() view returns(uint256, bytes20[])
func (_AccountManager *AccountManagerCallerSession) GetTotalAmount() (*big.Int, [][20]byte, error) {
	return _AccountManager.Contract.GetTotalAmount(&_AccountManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManager *AccountManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManager *AccountManagerSession) Owner() (common.Address, error) {
	return _AccountManager.Contract.Owner(&_AccountManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManager *AccountManagerCallerSession) Owner() (common.Address, error) {
	return _AccountManager.Contract.Owner(&_AccountManager.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _to, uint256 _lockDay) payable returns(bytes20)
func (_AccountManager *AccountManagerTransactor) Deposit(opts *bind.TransactOpts, _to common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "deposit", _to, _lockDay)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _to, uint256 _lockDay) payable returns(bytes20)
func (_AccountManager *AccountManagerSession) Deposit(_to common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.Contract.Deposit(&_AccountManager.TransactOpts, _to, _lockDay)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _to, uint256 _lockDay) payable returns(bytes20)
func (_AccountManager *AccountManagerTransactorSession) Deposit(_to common.Address, _lockDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.Contract.Deposit(&_AccountManager.TransactOpts, _to, _lockDay)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AccountManager *AccountManagerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AccountManager *AccountManagerSession) Initialize() (*types.Transaction, error) {
	return _AccountManager.Contract.Initialize(&_AccountManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AccountManager *AccountManagerTransactorSession) Initialize() (*types.Transaction, error) {
	return _AccountManager.Contract.Initialize(&_AccountManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManager *AccountManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManager *AccountManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountManager.Contract.RenounceOwnership(&_AccountManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManager *AccountManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountManager.Contract.RenounceOwnership(&_AccountManager.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x3ff1201f.
//
// Solidity: function reward(address _to, uint8 _rewardType) payable returns(bytes20)
func (_AccountManager *AccountManagerTransactor) Reward(opts *bind.TransactOpts, _to common.Address, _rewardType uint8) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "reward", _to, _rewardType)
}

// Reward is a paid mutator transaction binding the contract method 0x3ff1201f.
//
// Solidity: function reward(address _to, uint8 _rewardType) payable returns(bytes20)
func (_AccountManager *AccountManagerSession) Reward(_to common.Address, _rewardType uint8) (*types.Transaction, error) {
	return _AccountManager.Contract.Reward(&_AccountManager.TransactOpts, _to, _rewardType)
}

// Reward is a paid mutator transaction binding the contract method 0x3ff1201f.
//
// Solidity: function reward(address _to, uint8 _rewardType) payable returns(bytes20)
func (_AccountManager *AccountManagerTransactorSession) Reward(_to common.Address, _rewardType uint8) (*types.Transaction, error) {
	return _AccountManager.Contract.Reward(&_AccountManager.TransactOpts, _to, _rewardType)
}

// SetBindDay is a paid mutator transaction binding the contract method 0xdcc8e1ee.
//
// Solidity: function setBindDay(bytes20 _recordID, uint256 _bindDay) returns()
func (_AccountManager *AccountManagerTransactor) SetBindDay(opts *bind.TransactOpts, _recordID [20]byte, _bindDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "setBindDay", _recordID, _bindDay)
}

// SetBindDay is a paid mutator transaction binding the contract method 0xdcc8e1ee.
//
// Solidity: function setBindDay(bytes20 _recordID, uint256 _bindDay) returns()
func (_AccountManager *AccountManagerSession) SetBindDay(_recordID [20]byte, _bindDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.Contract.SetBindDay(&_AccountManager.TransactOpts, _recordID, _bindDay)
}

// SetBindDay is a paid mutator transaction binding the contract method 0xdcc8e1ee.
//
// Solidity: function setBindDay(bytes20 _recordID, uint256 _bindDay) returns()
func (_AccountManager *AccountManagerTransactorSession) SetBindDay(_recordID [20]byte, _bindDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.Contract.SetBindDay(&_AccountManager.TransactOpts, _recordID, _bindDay)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address _to, uint256 _amount, uint256 _lockDay) returns(bytes20)
func (_AccountManager *AccountManagerTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _lockDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "transfer", _to, _amount, _lockDay)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address _to, uint256 _amount, uint256 _lockDay) returns(bytes20)
func (_AccountManager *AccountManagerSession) Transfer(_to common.Address, _amount *big.Int, _lockDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.Contract.Transfer(&_AccountManager.TransactOpts, _to, _amount, _lockDay)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address _to, uint256 _amount, uint256 _lockDay) returns(bytes20)
func (_AccountManager *AccountManagerTransactorSession) Transfer(_to common.Address, _amount *big.Int, _lockDay *big.Int) (*types.Transaction, error) {
	return _AccountManager.Contract.Transfer(&_AccountManager.TransactOpts, _to, _amount, _lockDay)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManager *AccountManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManager *AccountManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.TransferOwnership(&_AccountManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManager *AccountManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.TransferOwnership(&_AccountManager.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_AccountManager *AccountManagerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_AccountManager *AccountManagerSession) Withdraw() (*types.Transaction, error) {
	return _AccountManager.Contract.Withdraw(&_AccountManager.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_AccountManager *AccountManagerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _AccountManager.Contract.Withdraw(&_AccountManager.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x60a2b8f1.
//
// Solidity: function withdraw(bytes20[] _recordIDs) returns(uint256)
func (_AccountManager *AccountManagerTransactor) Withdraw0(opts *bind.TransactOpts, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "withdraw0", _recordIDs)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x60a2b8f1.
//
// Solidity: function withdraw(bytes20[] _recordIDs) returns(uint256)
func (_AccountManager *AccountManagerSession) Withdraw0(_recordIDs [][20]byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Withdraw0(&_AccountManager.TransactOpts, _recordIDs)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x60a2b8f1.
//
// Solidity: function withdraw(bytes20[] _recordIDs) returns(uint256)
func (_AccountManager *AccountManagerTransactorSession) Withdraw0(_recordIDs [][20]byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Withdraw0(&_AccountManager.TransactOpts, _recordIDs)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AccountManager *AccountManagerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AccountManager.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AccountManager *AccountManagerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Fallback(&_AccountManager.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AccountManager *AccountManagerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Fallback(&_AccountManager.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AccountManager *AccountManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AccountManager *AccountManagerSession) Receive() (*types.Transaction, error) {
	return _AccountManager.Contract.Receive(&_AccountManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AccountManager *AccountManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _AccountManager.Contract.Receive(&_AccountManager.TransactOpts)
}

// AccountManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AccountManager contract.
type AccountManagerInitializedIterator struct {
	Event *AccountManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AccountManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerInitialized)
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
		it.Event = new(AccountManagerInitialized)
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
func (it *AccountManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerInitialized represents a Initialized event raised by the AccountManager contract.
type AccountManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AccountManager *AccountManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountManagerInitializedIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountManagerInitializedIterator{contract: _AccountManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AccountManager *AccountManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerInitialized)
				if err := _AccountManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AccountManager *AccountManagerFilterer) ParseInitialized(log types.Log) (*AccountManagerInitialized, error) {
	event := new(AccountManagerInitialized)
	if err := _AccountManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountManager contract.
type AccountManagerOwnershipTransferredIterator struct {
	Event *AccountManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerOwnershipTransferred)
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
		it.Event = new(AccountManagerOwnershipTransferred)
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
func (it *AccountManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AccountManager contract.
type AccountManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountManager *AccountManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerOwnershipTransferredIterator{contract: _AccountManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountManager *AccountManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerOwnershipTransferred)
				if err := _AccountManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccountManager *AccountManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AccountManagerOwnershipTransferred, error) {
	event := new(AccountManagerOwnershipTransferred)
	if err := _AccountManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerSafeDepositIterator is returned from FilterSafeDeposit and is used to iterate over the raw logs and unpacked data for SafeDeposit events raised by the AccountManager contract.
type AccountManagerSafeDepositIterator struct {
	Event *AccountManagerSafeDeposit // Event containing the contract specifics and raw log

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
func (it *AccountManagerSafeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerSafeDeposit)
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
		it.Event = new(AccountManagerSafeDeposit)
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
func (it *AccountManagerSafeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerSafeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerSafeDeposit represents a SafeDeposit event raised by the AccountManager contract.
type AccountManagerSafeDeposit struct {
	Addr     common.Address
	Amount   *big.Int
	LockDay  *big.Int
	ReocrdID [20]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSafeDeposit is a free log retrieval operation binding the contract event 0x41e917cac3031f97b0aa0fd23ba6d7fe4f65834ad44dcd0a39aedccbbe577af9.
//
// Solidity: event SafeDeposit(address _addr, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_AccountManager *AccountManagerFilterer) FilterSafeDeposit(opts *bind.FilterOpts) (*AccountManagerSafeDepositIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "SafeDeposit")
	if err != nil {
		return nil, err
	}
	return &AccountManagerSafeDepositIterator{contract: _AccountManager.contract, event: "SafeDeposit", logs: logs, sub: sub}, nil
}

// WatchSafeDeposit is a free log subscription operation binding the contract event 0x41e917cac3031f97b0aa0fd23ba6d7fe4f65834ad44dcd0a39aedccbbe577af9.
//
// Solidity: event SafeDeposit(address _addr, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_AccountManager *AccountManagerFilterer) WatchSafeDeposit(opts *bind.WatchOpts, sink chan<- *AccountManagerSafeDeposit) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "SafeDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerSafeDeposit)
				if err := _AccountManager.contract.UnpackLog(event, "SafeDeposit", log); err != nil {
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

// ParseSafeDeposit is a log parse operation binding the contract event 0x41e917cac3031f97b0aa0fd23ba6d7fe4f65834ad44dcd0a39aedccbbe577af9.
//
// Solidity: event SafeDeposit(address _addr, uint256 _amount, uint256 _lockDay, bytes20 _reocrdID)
func (_AccountManager *AccountManagerFilterer) ParseSafeDeposit(log types.Log) (*AccountManagerSafeDeposit, error) {
	event := new(AccountManagerSafeDeposit)
	if err := _AccountManager.contract.UnpackLog(event, "SafeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerSafeTransferIterator is returned from FilterSafeTransfer and is used to iterate over the raw logs and unpacked data for SafeTransfer events raised by the AccountManager contract.
type AccountManagerSafeTransferIterator struct {
	Event *AccountManagerSafeTransfer // Event containing the contract specifics and raw log

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
func (it *AccountManagerSafeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerSafeTransfer)
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
		it.Event = new(AccountManagerSafeTransfer)
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
func (it *AccountManagerSafeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerSafeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerSafeTransfer represents a SafeTransfer event raised by the AccountManager contract.
type AccountManagerSafeTransfer struct {
	From    common.Address
	To      common.Address
	Amount  *big.Int
	LockDay *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSafeTransfer is a free log retrieval operation binding the contract event 0xf94345ffe6f5021457656a4d50cd6fc1608b0eb00cde121dc806f833b3a3d04b.
//
// Solidity: event SafeTransfer(address _from, address _to, uint256 _amount, uint256 _lockDay)
func (_AccountManager *AccountManagerFilterer) FilterSafeTransfer(opts *bind.FilterOpts) (*AccountManagerSafeTransferIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "SafeTransfer")
	if err != nil {
		return nil, err
	}
	return &AccountManagerSafeTransferIterator{contract: _AccountManager.contract, event: "SafeTransfer", logs: logs, sub: sub}, nil
}

// WatchSafeTransfer is a free log subscription operation binding the contract event 0xf94345ffe6f5021457656a4d50cd6fc1608b0eb00cde121dc806f833b3a3d04b.
//
// Solidity: event SafeTransfer(address _from, address _to, uint256 _amount, uint256 _lockDay)
func (_AccountManager *AccountManagerFilterer) WatchSafeTransfer(opts *bind.WatchOpts, sink chan<- *AccountManagerSafeTransfer) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "SafeTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerSafeTransfer)
				if err := _AccountManager.contract.UnpackLog(event, "SafeTransfer", log); err != nil {
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

// ParseSafeTransfer is a log parse operation binding the contract event 0xf94345ffe6f5021457656a4d50cd6fc1608b0eb00cde121dc806f833b3a3d04b.
//
// Solidity: event SafeTransfer(address _from, address _to, uint256 _amount, uint256 _lockDay)
func (_AccountManager *AccountManagerFilterer) ParseSafeTransfer(log types.Log) (*AccountManagerSafeTransfer, error) {
	event := new(AccountManagerSafeTransfer)
	if err := _AccountManager.contract.UnpackLog(event, "SafeTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerSafeWithdrawIterator is returned from FilterSafeWithdraw and is used to iterate over the raw logs and unpacked data for SafeWithdraw events raised by the AccountManager contract.
type AccountManagerSafeWithdrawIterator struct {
	Event *AccountManagerSafeWithdraw // Event containing the contract specifics and raw log

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
func (it *AccountManagerSafeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerSafeWithdraw)
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
		it.Event = new(AccountManagerSafeWithdraw)
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
func (it *AccountManagerSafeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerSafeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerSafeWithdraw represents a SafeWithdraw event raised by the AccountManager contract.
type AccountManagerSafeWithdraw struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeWithdraw is a free log retrieval operation binding the contract event 0xac420a5a38642ce5e54e778ed9593117a21c00b6d912e86214979c3da2eb0958.
//
// Solidity: event SafeWithdraw(address _addr, uint256 _amount)
func (_AccountManager *AccountManagerFilterer) FilterSafeWithdraw(opts *bind.FilterOpts) (*AccountManagerSafeWithdrawIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "SafeWithdraw")
	if err != nil {
		return nil, err
	}
	return &AccountManagerSafeWithdrawIterator{contract: _AccountManager.contract, event: "SafeWithdraw", logs: logs, sub: sub}, nil
}

// WatchSafeWithdraw is a free log subscription operation binding the contract event 0xac420a5a38642ce5e54e778ed9593117a21c00b6d912e86214979c3da2eb0958.
//
// Solidity: event SafeWithdraw(address _addr, uint256 _amount)
func (_AccountManager *AccountManagerFilterer) WatchSafeWithdraw(opts *bind.WatchOpts, sink chan<- *AccountManagerSafeWithdraw) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "SafeWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerSafeWithdraw)
				if err := _AccountManager.contract.UnpackLog(event, "SafeWithdraw", log); err != nil {
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

// ParseSafeWithdraw is a log parse operation binding the contract event 0xac420a5a38642ce5e54e778ed9593117a21c00b6d912e86214979c3da2eb0958.
//
// Solidity: event SafeWithdraw(address _addr, uint256 _amount)
func (_AccountManager *AccountManagerFilterer) ParseSafeWithdraw(log types.Log) (*AccountManagerSafeWithdraw, error) {
	event := new(AccountManagerSafeWithdraw)
	if err := _AccountManager.contract.UnpackLog(event, "SafeWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
