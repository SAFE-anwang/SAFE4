// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SMNVote

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

// SMNVoteMetaData contains all meta data concerning the SMNVote contract.
var SMNVoteMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyAddr\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"}],\"name\":\"approval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyAddr\",\"type\":\"address\"},{\"internalType\":\"bytes20[]\",\"name\":\"_recordIDs\",\"type\":\"bytes20[]\"}],\"name\":\"approval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"decreaseRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedRecords4Voter\",\"outputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"retIDs\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxies4Voter\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"retAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"retNums\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVoteNum4Proxy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_smnAddr\",\"type\":\"address\"}],\"name\":\"getVoteNum4SMN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotedRecords4Voter\",\"outputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"retIDs\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotedSMN4Voter\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"retAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"retNums\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVoters4Proxy\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_smnAddr\",\"type\":\"address\"}],\"name\":\"getVoters4SMN\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"retAddrs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_smnAddr\",\"type\":\"address\"}],\"name\":\"proxyVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"}],\"name\":\"removeApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"_recordIDs\",\"type\":\"bytes20[]\"}],\"name\":\"removeApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"}],\"name\":\"removeRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"}],\"name\":\"removeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"_recordIDs\",\"type\":\"bytes20[]\"}],\"name\":\"removeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_smnAddr\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_recordID\",\"type\":\"bytes20\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_smnAddr\",\"type\":\"address\"},{\"internalType\":\"bytes20[]\",\"name\":\"_recordIDs\",\"type\":\"bytes20[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SMNVoteABI is the input ABI used to generate the binding from.
// Deprecated: Use SMNVoteMetaData.ABI instead.
var SMNVoteABI = SMNVoteMetaData.ABI

// SMNVote is an auto generated Go binding around an Ethereum contract.
type SMNVote struct {
	SMNVoteCaller     // Read-only binding to the contract
	SMNVoteTransactor // Write-only binding to the contract
	SMNVoteFilterer   // Log filterer for contract events
}

// SMNVoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type SMNVoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMNVoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SMNVoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMNVoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SMNVoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMNVoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SMNVoteSession struct {
	Contract     *SMNVote          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SMNVoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SMNVoteCallerSession struct {
	Contract *SMNVoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SMNVoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SMNVoteTransactorSession struct {
	Contract     *SMNVoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SMNVoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type SMNVoteRaw struct {
	Contract *SMNVote // Generic contract binding to access the raw methods on
}

// SMNVoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SMNVoteCallerRaw struct {
	Contract *SMNVoteCaller // Generic read-only contract binding to access the raw methods on
}

// SMNVoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SMNVoteTransactorRaw struct {
	Contract *SMNVoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSMNVote creates a new instance of SMNVote, bound to a specific deployed contract.
func NewSMNVote(address common.Address, backend bind.ContractBackend) (*SMNVote, error) {
	contract, err := bindSMNVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SMNVote{SMNVoteCaller: SMNVoteCaller{contract: contract}, SMNVoteTransactor: SMNVoteTransactor{contract: contract}, SMNVoteFilterer: SMNVoteFilterer{contract: contract}}, nil
}

// NewSMNVoteCaller creates a new read-only instance of SMNVote, bound to a specific deployed contract.
func NewSMNVoteCaller(address common.Address, caller bind.ContractCaller) (*SMNVoteCaller, error) {
	contract, err := bindSMNVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SMNVoteCaller{contract: contract}, nil
}

// NewSMNVoteTransactor creates a new write-only instance of SMNVote, bound to a specific deployed contract.
func NewSMNVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*SMNVoteTransactor, error) {
	contract, err := bindSMNVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SMNVoteTransactor{contract: contract}, nil
}

// NewSMNVoteFilterer creates a new log filterer instance of SMNVote, bound to a specific deployed contract.
func NewSMNVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*SMNVoteFilterer, error) {
	contract, err := bindSMNVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SMNVoteFilterer{contract: contract}, nil
}

// bindSMNVote binds a generic wrapper to an already deployed contract.
func bindSMNVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SMNVoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SMNVote *SMNVoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SMNVote.Contract.SMNVoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SMNVote *SMNVoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMNVote.Contract.SMNVoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SMNVote *SMNVoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SMNVote.Contract.SMNVoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SMNVote *SMNVoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SMNVote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SMNVote *SMNVoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMNVote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SMNVote *SMNVoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SMNVote.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SMNVote *SMNVoteCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SMNVote *SMNVoteSession) GetInitializeData() ([]byte, error) {
	return _SMNVote.Contract.GetInitializeData(&_SMNVote.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_SMNVote *SMNVoteCallerSession) GetInitializeData() ([]byte, error) {
	return _SMNVote.Contract.GetInitializeData(&_SMNVote.CallOpts)
}

// GetProxiedRecords4Voter is a free data retrieval call binding the contract method 0x77970201.
//
// Solidity: function getProxiedRecords4Voter() view returns(bytes20[] retIDs)
func (_SMNVote *SMNVoteCaller) GetProxiedRecords4Voter(opts *bind.CallOpts) ([][20]byte, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "getProxiedRecords4Voter")

	if err != nil {
		return *new([][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][20]byte)).(*[][20]byte)

	return out0, err

}

// GetProxiedRecords4Voter is a free data retrieval call binding the contract method 0x77970201.
//
// Solidity: function getProxiedRecords4Voter() view returns(bytes20[] retIDs)
func (_SMNVote *SMNVoteSession) GetProxiedRecords4Voter() ([][20]byte, error) {
	return _SMNVote.Contract.GetProxiedRecords4Voter(&_SMNVote.CallOpts)
}

// GetProxiedRecords4Voter is a free data retrieval call binding the contract method 0x77970201.
//
// Solidity: function getProxiedRecords4Voter() view returns(bytes20[] retIDs)
func (_SMNVote *SMNVoteCallerSession) GetProxiedRecords4Voter() ([][20]byte, error) {
	return _SMNVote.Contract.GetProxiedRecords4Voter(&_SMNVote.CallOpts)
}

// GetProxies4Voter is a free data retrieval call binding the contract method 0x30616728.
//
// Solidity: function getProxies4Voter() view returns(address[] retAddrs, uint256[] retNums)
func (_SMNVote *SMNVoteCaller) GetProxies4Voter(opts *bind.CallOpts) (struct {
	RetAddrs []common.Address
	RetNums  []*big.Int
}, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "getProxies4Voter")

	outstruct := new(struct {
		RetAddrs []common.Address
		RetNums  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RetAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.RetNums = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetProxies4Voter is a free data retrieval call binding the contract method 0x30616728.
//
// Solidity: function getProxies4Voter() view returns(address[] retAddrs, uint256[] retNums)
func (_SMNVote *SMNVoteSession) GetProxies4Voter() (struct {
	RetAddrs []common.Address
	RetNums  []*big.Int
}, error) {
	return _SMNVote.Contract.GetProxies4Voter(&_SMNVote.CallOpts)
}

// GetProxies4Voter is a free data retrieval call binding the contract method 0x30616728.
//
// Solidity: function getProxies4Voter() view returns(address[] retAddrs, uint256[] retNums)
func (_SMNVote *SMNVoteCallerSession) GetProxies4Voter() (struct {
	RetAddrs []common.Address
	RetNums  []*big.Int
}, error) {
	return _SMNVote.Contract.GetProxies4Voter(&_SMNVote.CallOpts)
}

// GetVoteNum4Proxy is a free data retrieval call binding the contract method 0xfc46f7f1.
//
// Solidity: function getVoteNum4Proxy() view returns(uint256)
func (_SMNVote *SMNVoteCaller) GetVoteNum4Proxy(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "getVoteNum4Proxy")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoteNum4Proxy is a free data retrieval call binding the contract method 0xfc46f7f1.
//
// Solidity: function getVoteNum4Proxy() view returns(uint256)
func (_SMNVote *SMNVoteSession) GetVoteNum4Proxy() (*big.Int, error) {
	return _SMNVote.Contract.GetVoteNum4Proxy(&_SMNVote.CallOpts)
}

// GetVoteNum4Proxy is a free data retrieval call binding the contract method 0xfc46f7f1.
//
// Solidity: function getVoteNum4Proxy() view returns(uint256)
func (_SMNVote *SMNVoteCallerSession) GetVoteNum4Proxy() (*big.Int, error) {
	return _SMNVote.Contract.GetVoteNum4Proxy(&_SMNVote.CallOpts)
}

// GetVoteNum4SMN is a free data retrieval call binding the contract method 0x2c572616.
//
// Solidity: function getVoteNum4SMN(address _smnAddr) view returns(uint256)
func (_SMNVote *SMNVoteCaller) GetVoteNum4SMN(opts *bind.CallOpts, _smnAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "getVoteNum4SMN", _smnAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoteNum4SMN is a free data retrieval call binding the contract method 0x2c572616.
//
// Solidity: function getVoteNum4SMN(address _smnAddr) view returns(uint256)
func (_SMNVote *SMNVoteSession) GetVoteNum4SMN(_smnAddr common.Address) (*big.Int, error) {
	return _SMNVote.Contract.GetVoteNum4SMN(&_SMNVote.CallOpts, _smnAddr)
}

// GetVoteNum4SMN is a free data retrieval call binding the contract method 0x2c572616.
//
// Solidity: function getVoteNum4SMN(address _smnAddr) view returns(uint256)
func (_SMNVote *SMNVoteCallerSession) GetVoteNum4SMN(_smnAddr common.Address) (*big.Int, error) {
	return _SMNVote.Contract.GetVoteNum4SMN(&_SMNVote.CallOpts, _smnAddr)
}

// GetVotedRecords4Voter is a free data retrieval call binding the contract method 0x86b13aa8.
//
// Solidity: function getVotedRecords4Voter() view returns(bytes20[] retIDs)
func (_SMNVote *SMNVoteCaller) GetVotedRecords4Voter(opts *bind.CallOpts) ([][20]byte, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "getVotedRecords4Voter")

	if err != nil {
		return *new([][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][20]byte)).(*[][20]byte)

	return out0, err

}

// GetVotedRecords4Voter is a free data retrieval call binding the contract method 0x86b13aa8.
//
// Solidity: function getVotedRecords4Voter() view returns(bytes20[] retIDs)
func (_SMNVote *SMNVoteSession) GetVotedRecords4Voter() ([][20]byte, error) {
	return _SMNVote.Contract.GetVotedRecords4Voter(&_SMNVote.CallOpts)
}

// GetVotedRecords4Voter is a free data retrieval call binding the contract method 0x86b13aa8.
//
// Solidity: function getVotedRecords4Voter() view returns(bytes20[] retIDs)
func (_SMNVote *SMNVoteCallerSession) GetVotedRecords4Voter() ([][20]byte, error) {
	return _SMNVote.Contract.GetVotedRecords4Voter(&_SMNVote.CallOpts)
}

// GetVotedSMN4Voter is a free data retrieval call binding the contract method 0xafb75b94.
//
// Solidity: function getVotedSMN4Voter() view returns(address[] retAddrs, uint256[] retNums)
func (_SMNVote *SMNVoteCaller) GetVotedSMN4Voter(opts *bind.CallOpts) (struct {
	RetAddrs []common.Address
	RetNums  []*big.Int
}, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "getVotedSMN4Voter")

	outstruct := new(struct {
		RetAddrs []common.Address
		RetNums  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RetAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.RetNums = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetVotedSMN4Voter is a free data retrieval call binding the contract method 0xafb75b94.
//
// Solidity: function getVotedSMN4Voter() view returns(address[] retAddrs, uint256[] retNums)
func (_SMNVote *SMNVoteSession) GetVotedSMN4Voter() (struct {
	RetAddrs []common.Address
	RetNums  []*big.Int
}, error) {
	return _SMNVote.Contract.GetVotedSMN4Voter(&_SMNVote.CallOpts)
}

// GetVotedSMN4Voter is a free data retrieval call binding the contract method 0xafb75b94.
//
// Solidity: function getVotedSMN4Voter() view returns(address[] retAddrs, uint256[] retNums)
func (_SMNVote *SMNVoteCallerSession) GetVotedSMN4Voter() (struct {
	RetAddrs []common.Address
	RetNums  []*big.Int
}, error) {
	return _SMNVote.Contract.GetVotedSMN4Voter(&_SMNVote.CallOpts)
}

// GetVoters4Proxy is a free data retrieval call binding the contract method 0xaf373fbf.
//
// Solidity: function getVoters4Proxy() view returns(address[])
func (_SMNVote *SMNVoteCaller) GetVoters4Proxy(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "getVoters4Proxy")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVoters4Proxy is a free data retrieval call binding the contract method 0xaf373fbf.
//
// Solidity: function getVoters4Proxy() view returns(address[])
func (_SMNVote *SMNVoteSession) GetVoters4Proxy() ([]common.Address, error) {
	return _SMNVote.Contract.GetVoters4Proxy(&_SMNVote.CallOpts)
}

// GetVoters4Proxy is a free data retrieval call binding the contract method 0xaf373fbf.
//
// Solidity: function getVoters4Proxy() view returns(address[])
func (_SMNVote *SMNVoteCallerSession) GetVoters4Proxy() ([]common.Address, error) {
	return _SMNVote.Contract.GetVoters4Proxy(&_SMNVote.CallOpts)
}

// GetVoters4SMN is a free data retrieval call binding the contract method 0x2947932b.
//
// Solidity: function getVoters4SMN(address _smnAddr) view returns(address[] retAddrs)
func (_SMNVote *SMNVoteCaller) GetVoters4SMN(opts *bind.CallOpts, _smnAddr common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "getVoters4SMN", _smnAddr)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVoters4SMN is a free data retrieval call binding the contract method 0x2947932b.
//
// Solidity: function getVoters4SMN(address _smnAddr) view returns(address[] retAddrs)
func (_SMNVote *SMNVoteSession) GetVoters4SMN(_smnAddr common.Address) ([]common.Address, error) {
	return _SMNVote.Contract.GetVoters4SMN(&_SMNVote.CallOpts, _smnAddr)
}

// GetVoters4SMN is a free data retrieval call binding the contract method 0x2947932b.
//
// Solidity: function getVoters4SMN(address _smnAddr) view returns(address[] retAddrs)
func (_SMNVote *SMNVoteCallerSession) GetVoters4SMN(_smnAddr common.Address) ([]common.Address, error) {
	return _SMNVote.Contract.GetVoters4SMN(&_SMNVote.CallOpts, _smnAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SMNVote *SMNVoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SMNVote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SMNVote *SMNVoteSession) Owner() (common.Address, error) {
	return _SMNVote.Contract.Owner(&_SMNVote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SMNVote *SMNVoteCallerSession) Owner() (common.Address, error) {
	return _SMNVote.Contract.Owner(&_SMNVote.CallOpts)
}

// Approval is a paid mutator transaction binding the contract method 0x1e1192c5.
//
// Solidity: function approval(address _proxyAddr, bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactor) Approval(opts *bind.TransactOpts, _proxyAddr common.Address, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "approval", _proxyAddr, _recordID)
}

// Approval is a paid mutator transaction binding the contract method 0x1e1192c5.
//
// Solidity: function approval(address _proxyAddr, bytes20 _recordID) returns()
func (_SMNVote *SMNVoteSession) Approval(_proxyAddr common.Address, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.Approval(&_SMNVote.TransactOpts, _proxyAddr, _recordID)
}

// Approval is a paid mutator transaction binding the contract method 0x1e1192c5.
//
// Solidity: function approval(address _proxyAddr, bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactorSession) Approval(_proxyAddr common.Address, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.Approval(&_SMNVote.TransactOpts, _proxyAddr, _recordID)
}

// Approval0 is a paid mutator transaction binding the contract method 0xde2594cd.
//
// Solidity: function approval(address _proxyAddr, bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteTransactor) Approval0(opts *bind.TransactOpts, _proxyAddr common.Address, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "approval0", _proxyAddr, _recordIDs)
}

// Approval0 is a paid mutator transaction binding the contract method 0xde2594cd.
//
// Solidity: function approval(address _proxyAddr, bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteSession) Approval0(_proxyAddr common.Address, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.Approval0(&_SMNVote.TransactOpts, _proxyAddr, _recordIDs)
}

// Approval0 is a paid mutator transaction binding the contract method 0xde2594cd.
//
// Solidity: function approval(address _proxyAddr, bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteTransactorSession) Approval0(_proxyAddr common.Address, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.Approval0(&_SMNVote.TransactOpts, _proxyAddr, _recordIDs)
}

// DecreaseRecord is a paid mutator transaction binding the contract method 0x6a04edb7.
//
// Solidity: function decreaseRecord(bytes20 _recordID, uint256 _amount, uint256 _num) returns()
func (_SMNVote *SMNVoteTransactor) DecreaseRecord(opts *bind.TransactOpts, _recordID [20]byte, _amount *big.Int, _num *big.Int) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "decreaseRecord", _recordID, _amount, _num)
}

// DecreaseRecord is a paid mutator transaction binding the contract method 0x6a04edb7.
//
// Solidity: function decreaseRecord(bytes20 _recordID, uint256 _amount, uint256 _num) returns()
func (_SMNVote *SMNVoteSession) DecreaseRecord(_recordID [20]byte, _amount *big.Int, _num *big.Int) (*types.Transaction, error) {
	return _SMNVote.Contract.DecreaseRecord(&_SMNVote.TransactOpts, _recordID, _amount, _num)
}

// DecreaseRecord is a paid mutator transaction binding the contract method 0x6a04edb7.
//
// Solidity: function decreaseRecord(bytes20 _recordID, uint256 _amount, uint256 _num) returns()
func (_SMNVote *SMNVoteTransactorSession) DecreaseRecord(_recordID [20]byte, _amount *big.Int, _num *big.Int) (*types.Transaction, error) {
	return _SMNVote.Contract.DecreaseRecord(&_SMNVote.TransactOpts, _recordID, _amount, _num)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SMNVote *SMNVoteTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SMNVote *SMNVoteSession) Initialize() (*types.Transaction, error) {
	return _SMNVote.Contract.Initialize(&_SMNVote.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SMNVote *SMNVoteTransactorSession) Initialize() (*types.Transaction, error) {
	return _SMNVote.Contract.Initialize(&_SMNVote.TransactOpts)
}

// ProxyVote is a paid mutator transaction binding the contract method 0xdd742cff.
//
// Solidity: function proxyVote(address _smnAddr) returns()
func (_SMNVote *SMNVoteTransactor) ProxyVote(opts *bind.TransactOpts, _smnAddr common.Address) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "proxyVote", _smnAddr)
}

// ProxyVote is a paid mutator transaction binding the contract method 0xdd742cff.
//
// Solidity: function proxyVote(address _smnAddr) returns()
func (_SMNVote *SMNVoteSession) ProxyVote(_smnAddr common.Address) (*types.Transaction, error) {
	return _SMNVote.Contract.ProxyVote(&_SMNVote.TransactOpts, _smnAddr)
}

// ProxyVote is a paid mutator transaction binding the contract method 0xdd742cff.
//
// Solidity: function proxyVote(address _smnAddr) returns()
func (_SMNVote *SMNVoteTransactorSession) ProxyVote(_smnAddr common.Address) (*types.Transaction, error) {
	return _SMNVote.Contract.ProxyVote(&_SMNVote.TransactOpts, _smnAddr)
}

// RemoveApproval is a paid mutator transaction binding the contract method 0xd311a62a.
//
// Solidity: function removeApproval(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactor) RemoveApproval(opts *bind.TransactOpts, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "removeApproval", _recordID)
}

// RemoveApproval is a paid mutator transaction binding the contract method 0xd311a62a.
//
// Solidity: function removeApproval(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteSession) RemoveApproval(_recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveApproval(&_SMNVote.TransactOpts, _recordID)
}

// RemoveApproval is a paid mutator transaction binding the contract method 0xd311a62a.
//
// Solidity: function removeApproval(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactorSession) RemoveApproval(_recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveApproval(&_SMNVote.TransactOpts, _recordID)
}

// RemoveApproval0 is a paid mutator transaction binding the contract method 0xf0bde029.
//
// Solidity: function removeApproval(bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteTransactor) RemoveApproval0(opts *bind.TransactOpts, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "removeApproval0", _recordIDs)
}

// RemoveApproval0 is a paid mutator transaction binding the contract method 0xf0bde029.
//
// Solidity: function removeApproval(bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteSession) RemoveApproval0(_recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveApproval0(&_SMNVote.TransactOpts, _recordIDs)
}

// RemoveApproval0 is a paid mutator transaction binding the contract method 0xf0bde029.
//
// Solidity: function removeApproval(bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteTransactorSession) RemoveApproval0(_recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveApproval0(&_SMNVote.TransactOpts, _recordIDs)
}

// RemoveRecord is a paid mutator transaction binding the contract method 0x5ae399e5.
//
// Solidity: function removeRecord(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactor) RemoveRecord(opts *bind.TransactOpts, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "removeRecord", _recordID)
}

// RemoveRecord is a paid mutator transaction binding the contract method 0x5ae399e5.
//
// Solidity: function removeRecord(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteSession) RemoveRecord(_recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveRecord(&_SMNVote.TransactOpts, _recordID)
}

// RemoveRecord is a paid mutator transaction binding the contract method 0x5ae399e5.
//
// Solidity: function removeRecord(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactorSession) RemoveRecord(_recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveRecord(&_SMNVote.TransactOpts, _recordID)
}

// RemoveVote is a paid mutator transaction binding the contract method 0x5ce30d97.
//
// Solidity: function removeVote(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactor) RemoveVote(opts *bind.TransactOpts, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "removeVote", _recordID)
}

// RemoveVote is a paid mutator transaction binding the contract method 0x5ce30d97.
//
// Solidity: function removeVote(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteSession) RemoveVote(_recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveVote(&_SMNVote.TransactOpts, _recordID)
}

// RemoveVote is a paid mutator transaction binding the contract method 0x5ce30d97.
//
// Solidity: function removeVote(bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactorSession) RemoveVote(_recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveVote(&_SMNVote.TransactOpts, _recordID)
}

// RemoveVote0 is a paid mutator transaction binding the contract method 0x730b11dc.
//
// Solidity: function removeVote(bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteTransactor) RemoveVote0(opts *bind.TransactOpts, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "removeVote0", _recordIDs)
}

// RemoveVote0 is a paid mutator transaction binding the contract method 0x730b11dc.
//
// Solidity: function removeVote(bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteSession) RemoveVote0(_recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveVote0(&_SMNVote.TransactOpts, _recordIDs)
}

// RemoveVote0 is a paid mutator transaction binding the contract method 0x730b11dc.
//
// Solidity: function removeVote(bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteTransactorSession) RemoveVote0(_recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.RemoveVote0(&_SMNVote.TransactOpts, _recordIDs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SMNVote *SMNVoteTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SMNVote *SMNVoteSession) RenounceOwnership() (*types.Transaction, error) {
	return _SMNVote.Contract.RenounceOwnership(&_SMNVote.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SMNVote *SMNVoteTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SMNVote.Contract.RenounceOwnership(&_SMNVote.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SMNVote *SMNVoteTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SMNVote *SMNVoteSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SMNVote.Contract.TransferOwnership(&_SMNVote.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SMNVote *SMNVoteTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SMNVote.Contract.TransferOwnership(&_SMNVote.TransactOpts, newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0x35cca783.
//
// Solidity: function vote(address _smnAddr, bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactor) Vote(opts *bind.TransactOpts, _smnAddr common.Address, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "vote", _smnAddr, _recordID)
}

// Vote is a paid mutator transaction binding the contract method 0x35cca783.
//
// Solidity: function vote(address _smnAddr, bytes20 _recordID) returns()
func (_SMNVote *SMNVoteSession) Vote(_smnAddr common.Address, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.Vote(&_SMNVote.TransactOpts, _smnAddr, _recordID)
}

// Vote is a paid mutator transaction binding the contract method 0x35cca783.
//
// Solidity: function vote(address _smnAddr, bytes20 _recordID) returns()
func (_SMNVote *SMNVoteTransactorSession) Vote(_smnAddr common.Address, _recordID [20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.Vote(&_SMNVote.TransactOpts, _smnAddr, _recordID)
}

// Vote0 is a paid mutator transaction binding the contract method 0xfb3abaf6.
//
// Solidity: function vote(address _smnAddr, bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteTransactor) Vote0(opts *bind.TransactOpts, _smnAddr common.Address, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.contract.Transact(opts, "vote0", _smnAddr, _recordIDs)
}

// Vote0 is a paid mutator transaction binding the contract method 0xfb3abaf6.
//
// Solidity: function vote(address _smnAddr, bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteSession) Vote0(_smnAddr common.Address, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.Vote0(&_SMNVote.TransactOpts, _smnAddr, _recordIDs)
}

// Vote0 is a paid mutator transaction binding the contract method 0xfb3abaf6.
//
// Solidity: function vote(address _smnAddr, bytes20[] _recordIDs) returns()
func (_SMNVote *SMNVoteTransactorSession) Vote0(_smnAddr common.Address, _recordIDs [][20]byte) (*types.Transaction, error) {
	return _SMNVote.Contract.Vote0(&_SMNVote.TransactOpts, _smnAddr, _recordIDs)
}

// SMNVoteInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SMNVote contract.
type SMNVoteInitializedIterator struct {
	Event *SMNVoteInitialized // Event containing the contract specifics and raw log

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
func (it *SMNVoteInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMNVoteInitialized)
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
		it.Event = new(SMNVoteInitialized)
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
func (it *SMNVoteInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMNVoteInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMNVoteInitialized represents a Initialized event raised by the SMNVote contract.
type SMNVoteInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SMNVote *SMNVoteFilterer) FilterInitialized(opts *bind.FilterOpts) (*SMNVoteInitializedIterator, error) {

	logs, sub, err := _SMNVote.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SMNVoteInitializedIterator{contract: _SMNVote.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SMNVote *SMNVoteFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SMNVoteInitialized) (event.Subscription, error) {

	logs, sub, err := _SMNVote.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMNVoteInitialized)
				if err := _SMNVote.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SMNVote *SMNVoteFilterer) ParseInitialized(log types.Log) (*SMNVoteInitialized, error) {
	event := new(SMNVoteInitialized)
	if err := _SMNVote.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SMNVoteOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SMNVote contract.
type SMNVoteOwnershipTransferredIterator struct {
	Event *SMNVoteOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SMNVoteOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMNVoteOwnershipTransferred)
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
		it.Event = new(SMNVoteOwnershipTransferred)
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
func (it *SMNVoteOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMNVoteOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMNVoteOwnershipTransferred represents a OwnershipTransferred event raised by the SMNVote contract.
type SMNVoteOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SMNVote *SMNVoteFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SMNVoteOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SMNVote.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SMNVoteOwnershipTransferredIterator{contract: _SMNVote.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SMNVote *SMNVoteFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SMNVoteOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SMNVote.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMNVoteOwnershipTransferred)
				if err := _SMNVote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SMNVote *SMNVoteFilterer) ParseOwnershipTransferred(log types.Log) (*SMNVoteOwnershipTransferred, error) {
	event := new(SMNVoteOwnershipTransferred)
	if err := _SMNVote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
