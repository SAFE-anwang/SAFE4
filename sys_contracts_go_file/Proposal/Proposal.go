// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Proposal

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

// IProposalProposalInfo is an auto generated low-level Go binding around an user-defined struct.
type IProposalProposalInfo struct {
	Id           *big.Int
	Creator      common.Address
	Title        string
	PayAmount    *big.Int
	PayTimes     *big.Int
	StartPayTime *big.Int
	EndPayTime   *big.Int
	Description  string
	Detail       string
	Voters       []common.Address
	VoteResults  []*big.Int
	State        *big.Int
	CreateHeight *big.Int
	UpdateHeight *big.Int
}

// ProposalMetaData contains all meta data concerning the Proposal contract.
var ProposalMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"}],\"name\":\"ProposalAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_state\",\"type\":\"uint256\"}],\"name\":\"ProposalState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_voteResult\",\"type\":\"uint256\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"}],\"name\":\"changTitile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"changeDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_detail\",\"type\":\"string\"}],\"name\":\"changeDetail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endPayTime\",\"type\":\"uint256\"}],\"name\":\"changeEndPayTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_payAmount\",\"type\":\"uint256\"}],\"name\":\"changePayAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_payTimes\",\"type\":\"uint256\"}],\"name\":\"changePayTimes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startPayTime\",\"type\":\"uint256\"}],\"name\":\"changeStartPayTimes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_payAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_payTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startPayTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endPayTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_detail\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAll\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPayTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPayTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"voteResults\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIProposal.ProposalInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPayTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPayTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"voteResults\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIProposal.ProposalInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMine\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPayTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPayTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"voteResults\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structIProposal.ProposalInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_voteResult\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProposalABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalMetaData.ABI instead.
var ProposalABI = ProposalMetaData.ABI

// Proposal is an auto generated Go binding around an Ethereum contract.
type Proposal struct {
	ProposalCaller     // Read-only binding to the contract
	ProposalTransactor // Write-only binding to the contract
	ProposalFilterer   // Log filterer for contract events
}

// ProposalCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalSession struct {
	Contract     *Proposal         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalCallerSession struct {
	Contract *ProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProposalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalTransactorSession struct {
	Contract     *ProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProposalRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalRaw struct {
	Contract *Proposal // Generic contract binding to access the raw methods on
}

// ProposalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalCallerRaw struct {
	Contract *ProposalCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalTransactorRaw struct {
	Contract *ProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposal creates a new instance of Proposal, bound to a specific deployed contract.
func NewProposal(address common.Address, backend bind.ContractBackend) (*Proposal, error) {
	contract, err := bindProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proposal{ProposalCaller: ProposalCaller{contract: contract}, ProposalTransactor: ProposalTransactor{contract: contract}, ProposalFilterer: ProposalFilterer{contract: contract}}, nil
}

// NewProposalCaller creates a new read-only instance of Proposal, bound to a specific deployed contract.
func NewProposalCaller(address common.Address, caller bind.ContractCaller) (*ProposalCaller, error) {
	contract, err := bindProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalCaller{contract: contract}, nil
}

// NewProposalTransactor creates a new write-only instance of Proposal, bound to a specific deployed contract.
func NewProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalTransactor, error) {
	contract, err := bindProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalTransactor{contract: contract}, nil
}

// NewProposalFilterer creates a new log filterer instance of Proposal, bound to a specific deployed contract.
func NewProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalFilterer, error) {
	contract, err := bindProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalFilterer{contract: contract}, nil
}

// bindProposal binds a generic wrapper to an already deployed contract.
func bindProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProposalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposal *ProposalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proposal.Contract.ProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposal *ProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.Contract.ProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposal *ProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposal.Contract.ProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposal *ProposalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposal *ProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposal *ProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposal.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Proposal *ProposalCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Proposal *ProposalSession) GetInitializeData() ([]byte, error) {
	return _Proposal.Contract.GetInitializeData(&_Proposal.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Proposal *ProposalCallerSession) GetInitializeData() ([]byte, error) {
	return _Proposal.Contract.GetInitializeData(&_Proposal.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256)[])
func (_Proposal *ProposalCaller) GetAll(opts *bind.CallOpts) ([]IProposalProposalInfo, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "getAll")

	if err != nil {
		return *new([]IProposalProposalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IProposalProposalInfo)).(*[]IProposalProposalInfo)

	return out0, err

}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256)[])
func (_Proposal *ProposalSession) GetAll() ([]IProposalProposalInfo, error) {
	return _Proposal.Contract.GetAll(&_Proposal.CallOpts)
}

// GetAll is a free data retrieval call binding the contract method 0x53ed5143.
//
// Solidity: function getAll() view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256)[])
func (_Proposal *ProposalCallerSession) GetAll() ([]IProposalProposalInfo, error) {
	return _Proposal.Contract.GetAll(&_Proposal.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _id) view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256))
func (_Proposal *ProposalCaller) GetInfo(opts *bind.CallOpts, _id *big.Int) (IProposalProposalInfo, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "getInfo", _id)

	if err != nil {
		return *new(IProposalProposalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IProposalProposalInfo)).(*IProposalProposalInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _id) view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256))
func (_Proposal *ProposalSession) GetInfo(_id *big.Int) (IProposalProposalInfo, error) {
	return _Proposal.Contract.GetInfo(&_Proposal.CallOpts, _id)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _id) view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256))
func (_Proposal *ProposalCallerSession) GetInfo(_id *big.Int) (IProposalProposalInfo, error) {
	return _Proposal.Contract.GetInfo(&_Proposal.CallOpts, _id)
}

// GetMine is a free data retrieval call binding the contract method 0xcd008f1a.
//
// Solidity: function getMine() view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256)[])
func (_Proposal *ProposalCaller) GetMine(opts *bind.CallOpts) ([]IProposalProposalInfo, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "getMine")

	if err != nil {
		return *new([]IProposalProposalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IProposalProposalInfo)).(*[]IProposalProposalInfo)

	return out0, err

}

// GetMine is a free data retrieval call binding the contract method 0xcd008f1a.
//
// Solidity: function getMine() view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256)[])
func (_Proposal *ProposalSession) GetMine() ([]IProposalProposalInfo, error) {
	return _Proposal.Contract.GetMine(&_Proposal.CallOpts)
}

// GetMine is a free data retrieval call binding the contract method 0xcd008f1a.
//
// Solidity: function getMine() view returns((uint256,address,string,uint256,uint256,uint256,uint256,string,string,address[],uint256[],uint256,uint256,uint256)[])
func (_Proposal *ProposalCallerSession) GetMine() ([]IProposalProposalInfo, error) {
	return _Proposal.Contract.GetMine(&_Proposal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proposal *ProposalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proposal *ProposalSession) Owner() (common.Address, error) {
	return _Proposal.Contract.Owner(&_Proposal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proposal *ProposalCallerSession) Owner() (common.Address, error) {
	return _Proposal.Contract.Owner(&_Proposal.CallOpts)
}

// ChangTitile is a paid mutator transaction binding the contract method 0x9cfc01cb.
//
// Solidity: function changTitile(uint256 _id, string _title) returns()
func (_Proposal *ProposalTransactor) ChangTitile(opts *bind.TransactOpts, _id *big.Int, _title string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "changTitile", _id, _title)
}

// ChangTitile is a paid mutator transaction binding the contract method 0x9cfc01cb.
//
// Solidity: function changTitile(uint256 _id, string _title) returns()
func (_Proposal *ProposalSession) ChangTitile(_id *big.Int, _title string) (*types.Transaction, error) {
	return _Proposal.Contract.ChangTitile(&_Proposal.TransactOpts, _id, _title)
}

// ChangTitile is a paid mutator transaction binding the contract method 0x9cfc01cb.
//
// Solidity: function changTitile(uint256 _id, string _title) returns()
func (_Proposal *ProposalTransactorSession) ChangTitile(_id *big.Int, _title string) (*types.Transaction, error) {
	return _Proposal.Contract.ChangTitile(&_Proposal.TransactOpts, _id, _title)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0xde6a06ed.
//
// Solidity: function changeDescription(uint256 _id, string _description) returns()
func (_Proposal *ProposalTransactor) ChangeDescription(opts *bind.TransactOpts, _id *big.Int, _description string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "changeDescription", _id, _description)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0xde6a06ed.
//
// Solidity: function changeDescription(uint256 _id, string _description) returns()
func (_Proposal *ProposalSession) ChangeDescription(_id *big.Int, _description string) (*types.Transaction, error) {
	return _Proposal.Contract.ChangeDescription(&_Proposal.TransactOpts, _id, _description)
}

// ChangeDescription is a paid mutator transaction binding the contract method 0xde6a06ed.
//
// Solidity: function changeDescription(uint256 _id, string _description) returns()
func (_Proposal *ProposalTransactorSession) ChangeDescription(_id *big.Int, _description string) (*types.Transaction, error) {
	return _Proposal.Contract.ChangeDescription(&_Proposal.TransactOpts, _id, _description)
}

// ChangeDetail is a paid mutator transaction binding the contract method 0x36d35abd.
//
// Solidity: function changeDetail(uint256 _id, string _detail) returns()
func (_Proposal *ProposalTransactor) ChangeDetail(opts *bind.TransactOpts, _id *big.Int, _detail string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "changeDetail", _id, _detail)
}

// ChangeDetail is a paid mutator transaction binding the contract method 0x36d35abd.
//
// Solidity: function changeDetail(uint256 _id, string _detail) returns()
func (_Proposal *ProposalSession) ChangeDetail(_id *big.Int, _detail string) (*types.Transaction, error) {
	return _Proposal.Contract.ChangeDetail(&_Proposal.TransactOpts, _id, _detail)
}

// ChangeDetail is a paid mutator transaction binding the contract method 0x36d35abd.
//
// Solidity: function changeDetail(uint256 _id, string _detail) returns()
func (_Proposal *ProposalTransactorSession) ChangeDetail(_id *big.Int, _detail string) (*types.Transaction, error) {
	return _Proposal.Contract.ChangeDetail(&_Proposal.TransactOpts, _id, _detail)
}

// ChangeEndPayTime is a paid mutator transaction binding the contract method 0x70de9df3.
//
// Solidity: function changeEndPayTime(uint256 _id, uint256 _endPayTime) returns()
func (_Proposal *ProposalTransactor) ChangeEndPayTime(opts *bind.TransactOpts, _id *big.Int, _endPayTime *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "changeEndPayTime", _id, _endPayTime)
}

// ChangeEndPayTime is a paid mutator transaction binding the contract method 0x70de9df3.
//
// Solidity: function changeEndPayTime(uint256 _id, uint256 _endPayTime) returns()
func (_Proposal *ProposalSession) ChangeEndPayTime(_id *big.Int, _endPayTime *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.ChangeEndPayTime(&_Proposal.TransactOpts, _id, _endPayTime)
}

// ChangeEndPayTime is a paid mutator transaction binding the contract method 0x70de9df3.
//
// Solidity: function changeEndPayTime(uint256 _id, uint256 _endPayTime) returns()
func (_Proposal *ProposalTransactorSession) ChangeEndPayTime(_id *big.Int, _endPayTime *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.ChangeEndPayTime(&_Proposal.TransactOpts, _id, _endPayTime)
}

// ChangePayAmount is a paid mutator transaction binding the contract method 0x2cad4e9f.
//
// Solidity: function changePayAmount(uint256 _id, uint256 _payAmount) returns()
func (_Proposal *ProposalTransactor) ChangePayAmount(opts *bind.TransactOpts, _id *big.Int, _payAmount *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "changePayAmount", _id, _payAmount)
}

// ChangePayAmount is a paid mutator transaction binding the contract method 0x2cad4e9f.
//
// Solidity: function changePayAmount(uint256 _id, uint256 _payAmount) returns()
func (_Proposal *ProposalSession) ChangePayAmount(_id *big.Int, _payAmount *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.ChangePayAmount(&_Proposal.TransactOpts, _id, _payAmount)
}

// ChangePayAmount is a paid mutator transaction binding the contract method 0x2cad4e9f.
//
// Solidity: function changePayAmount(uint256 _id, uint256 _payAmount) returns()
func (_Proposal *ProposalTransactorSession) ChangePayAmount(_id *big.Int, _payAmount *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.ChangePayAmount(&_Proposal.TransactOpts, _id, _payAmount)
}

// ChangePayTimes is a paid mutator transaction binding the contract method 0x1944b342.
//
// Solidity: function changePayTimes(uint256 _id, uint256 _payTimes) returns()
func (_Proposal *ProposalTransactor) ChangePayTimes(opts *bind.TransactOpts, _id *big.Int, _payTimes *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "changePayTimes", _id, _payTimes)
}

// ChangePayTimes is a paid mutator transaction binding the contract method 0x1944b342.
//
// Solidity: function changePayTimes(uint256 _id, uint256 _payTimes) returns()
func (_Proposal *ProposalSession) ChangePayTimes(_id *big.Int, _payTimes *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.ChangePayTimes(&_Proposal.TransactOpts, _id, _payTimes)
}

// ChangePayTimes is a paid mutator transaction binding the contract method 0x1944b342.
//
// Solidity: function changePayTimes(uint256 _id, uint256 _payTimes) returns()
func (_Proposal *ProposalTransactorSession) ChangePayTimes(_id *big.Int, _payTimes *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.ChangePayTimes(&_Proposal.TransactOpts, _id, _payTimes)
}

// ChangeStartPayTimes is a paid mutator transaction binding the contract method 0xa03e5ebf.
//
// Solidity: function changeStartPayTimes(uint256 _id, uint256 _startPayTime) returns()
func (_Proposal *ProposalTransactor) ChangeStartPayTimes(opts *bind.TransactOpts, _id *big.Int, _startPayTime *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "changeStartPayTimes", _id, _startPayTime)
}

// ChangeStartPayTimes is a paid mutator transaction binding the contract method 0xa03e5ebf.
//
// Solidity: function changeStartPayTimes(uint256 _id, uint256 _startPayTime) returns()
func (_Proposal *ProposalSession) ChangeStartPayTimes(_id *big.Int, _startPayTime *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.ChangeStartPayTimes(&_Proposal.TransactOpts, _id, _startPayTime)
}

// ChangeStartPayTimes is a paid mutator transaction binding the contract method 0xa03e5ebf.
//
// Solidity: function changeStartPayTimes(uint256 _id, uint256 _startPayTime) returns()
func (_Proposal *ProposalTransactorSession) ChangeStartPayTimes(_id *big.Int, _startPayTime *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.ChangeStartPayTimes(&_Proposal.TransactOpts, _id, _startPayTime)
}

// Create is a paid mutator transaction binding the contract method 0xa1facc40.
//
// Solidity: function create(string _title, uint256 _payAmount, uint256 _payTimes, uint256 _startPayTime, uint256 _endPayTime, string _description, string _detail) payable returns(uint256)
func (_Proposal *ProposalTransactor) Create(opts *bind.TransactOpts, _title string, _payAmount *big.Int, _payTimes *big.Int, _startPayTime *big.Int, _endPayTime *big.Int, _description string, _detail string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "create", _title, _payAmount, _payTimes, _startPayTime, _endPayTime, _description, _detail)
}

// Create is a paid mutator transaction binding the contract method 0xa1facc40.
//
// Solidity: function create(string _title, uint256 _payAmount, uint256 _payTimes, uint256 _startPayTime, uint256 _endPayTime, string _description, string _detail) payable returns(uint256)
func (_Proposal *ProposalSession) Create(_title string, _payAmount *big.Int, _payTimes *big.Int, _startPayTime *big.Int, _endPayTime *big.Int, _description string, _detail string) (*types.Transaction, error) {
	return _Proposal.Contract.Create(&_Proposal.TransactOpts, _title, _payAmount, _payTimes, _startPayTime, _endPayTime, _description, _detail)
}

// Create is a paid mutator transaction binding the contract method 0xa1facc40.
//
// Solidity: function create(string _title, uint256 _payAmount, uint256 _payTimes, uint256 _startPayTime, uint256 _endPayTime, string _description, string _detail) payable returns(uint256)
func (_Proposal *ProposalTransactorSession) Create(_title string, _payAmount *big.Int, _payTimes *big.Int, _startPayTime *big.Int, _endPayTime *big.Int, _description string, _detail string) (*types.Transaction, error) {
	return _Proposal.Contract.Create(&_Proposal.TransactOpts, _title, _payAmount, _payTimes, _startPayTime, _endPayTime, _description, _detail)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Proposal *ProposalTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Proposal *ProposalSession) Initialize() (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Proposal *ProposalTransactorSession) Initialize() (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proposal *ProposalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proposal *ProposalSession) RenounceOwnership() (*types.Transaction, error) {
	return _Proposal.Contract.RenounceOwnership(&_Proposal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proposal *ProposalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Proposal.Contract.RenounceOwnership(&_Proposal.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proposal *ProposalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proposal *ProposalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.TransferOwnership(&_Proposal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proposal *ProposalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.TransferOwnership(&_Proposal.TransactOpts, newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _id, uint256 _voteResult) returns()
func (_Proposal *ProposalTransactor) Vote(opts *bind.TransactOpts, _id *big.Int, _voteResult *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "vote", _id, _voteResult)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _id, uint256 _voteResult) returns()
func (_Proposal *ProposalSession) Vote(_id *big.Int, _voteResult *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.Vote(&_Proposal.TransactOpts, _id, _voteResult)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _id, uint256 _voteResult) returns()
func (_Proposal *ProposalTransactorSession) Vote(_id *big.Int, _voteResult *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.Vote(&_Proposal.TransactOpts, _id, _voteResult)
}

// ProposalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Proposal contract.
type ProposalInitializedIterator struct {
	Event *ProposalInitialized // Event containing the contract specifics and raw log

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
func (it *ProposalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalInitialized)
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
		it.Event = new(ProposalInitialized)
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
func (it *ProposalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalInitialized represents a Initialized event raised by the Proposal contract.
type ProposalInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Proposal *ProposalFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProposalInitializedIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProposalInitializedIterator{contract: _Proposal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Proposal *ProposalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProposalInitialized) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalInitialized)
				if err := _Proposal.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseInitialized(log types.Log) (*ProposalInitialized, error) {
	event := new(ProposalInitialized)
	if err := _Proposal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Proposal contract.
type ProposalOwnershipTransferredIterator struct {
	Event *ProposalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProposalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalOwnershipTransferred)
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
		it.Event = new(ProposalOwnershipTransferred)
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
func (it *ProposalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalOwnershipTransferred represents a OwnershipTransferred event raised by the Proposal contract.
type ProposalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proposal *ProposalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProposalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProposalOwnershipTransferredIterator{contract: _Proposal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proposal *ProposalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProposalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalOwnershipTransferred)
				if err := _Proposal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseOwnershipTransferred(log types.Log) (*ProposalOwnershipTransferred, error) {
	event := new(ProposalOwnershipTransferred)
	if err := _Proposal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalProposalAddIterator is returned from FilterProposalAdd and is used to iterate over the raw logs and unpacked data for ProposalAdd events raised by the Proposal contract.
type ProposalProposalAddIterator struct {
	Event *ProposalProposalAdd // Event containing the contract specifics and raw log

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
func (it *ProposalProposalAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalProposalAdd)
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
		it.Event = new(ProposalProposalAdd)
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
func (it *ProposalProposalAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalProposalAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalProposalAdd represents a ProposalAdd event raised by the Proposal contract.
type ProposalProposalAdd struct {
	Id    *big.Int
	Title string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProposalAdd is a free log retrieval operation binding the contract event 0x7b57fcf35b1bcd5bb9218d359ee897d9b58f7313b66f60b7ab7be17e81508793.
//
// Solidity: event ProposalAdd(uint256 _id, string _title)
func (_Proposal *ProposalFilterer) FilterProposalAdd(opts *bind.FilterOpts) (*ProposalProposalAddIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "ProposalAdd")
	if err != nil {
		return nil, err
	}
	return &ProposalProposalAddIterator{contract: _Proposal.contract, event: "ProposalAdd", logs: logs, sub: sub}, nil
}

// WatchProposalAdd is a free log subscription operation binding the contract event 0x7b57fcf35b1bcd5bb9218d359ee897d9b58f7313b66f60b7ab7be17e81508793.
//
// Solidity: event ProposalAdd(uint256 _id, string _title)
func (_Proposal *ProposalFilterer) WatchProposalAdd(opts *bind.WatchOpts, sink chan<- *ProposalProposalAdd) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "ProposalAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalProposalAdd)
				if err := _Proposal.contract.UnpackLog(event, "ProposalAdd", log); err != nil {
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

// ParseProposalAdd is a log parse operation binding the contract event 0x7b57fcf35b1bcd5bb9218d359ee897d9b58f7313b66f60b7ab7be17e81508793.
//
// Solidity: event ProposalAdd(uint256 _id, string _title)
func (_Proposal *ProposalFilterer) ParseProposalAdd(log types.Log) (*ProposalProposalAdd, error) {
	event := new(ProposalProposalAdd)
	if err := _Proposal.contract.UnpackLog(event, "ProposalAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalProposalStateIterator is returned from FilterProposalState and is used to iterate over the raw logs and unpacked data for ProposalState events raised by the Proposal contract.
type ProposalProposalStateIterator struct {
	Event *ProposalProposalState // Event containing the contract specifics and raw log

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
func (it *ProposalProposalStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalProposalState)
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
		it.Event = new(ProposalProposalState)
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
func (it *ProposalProposalStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalProposalStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalProposalState represents a ProposalState event raised by the Proposal contract.
type ProposalProposalState struct {
	Id    *big.Int
	State *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProposalState is a free log retrieval operation binding the contract event 0x55ed2de72a4c56af374ebeca99cbcc42204b4370d5da9983d2e71d9aaf3f0b77.
//
// Solidity: event ProposalState(uint256 _id, uint256 _state)
func (_Proposal *ProposalFilterer) FilterProposalState(opts *bind.FilterOpts) (*ProposalProposalStateIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "ProposalState")
	if err != nil {
		return nil, err
	}
	return &ProposalProposalStateIterator{contract: _Proposal.contract, event: "ProposalState", logs: logs, sub: sub}, nil
}

// WatchProposalState is a free log subscription operation binding the contract event 0x55ed2de72a4c56af374ebeca99cbcc42204b4370d5da9983d2e71d9aaf3f0b77.
//
// Solidity: event ProposalState(uint256 _id, uint256 _state)
func (_Proposal *ProposalFilterer) WatchProposalState(opts *bind.WatchOpts, sink chan<- *ProposalProposalState) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "ProposalState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalProposalState)
				if err := _Proposal.contract.UnpackLog(event, "ProposalState", log); err != nil {
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

// ParseProposalState is a log parse operation binding the contract event 0x55ed2de72a4c56af374ebeca99cbcc42204b4370d5da9983d2e71d9aaf3f0b77.
//
// Solidity: event ProposalState(uint256 _id, uint256 _state)
func (_Proposal *ProposalFilterer) ParseProposalState(log types.Log) (*ProposalProposalState, error) {
	event := new(ProposalProposalState)
	if err := _Proposal.contract.UnpackLog(event, "ProposalState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the Proposal contract.
type ProposalProposalVoteIterator struct {
	Event *ProposalProposalVote // Event containing the contract specifics and raw log

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
func (it *ProposalProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalProposalVote)
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
		it.Event = new(ProposalProposalVote)
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
func (it *ProposalProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalProposalVote represents a ProposalVote event raised by the Proposal contract.
type ProposalProposalVote struct {
	Id         *big.Int
	Voter      common.Address
	VoteResult *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0xb7d6a4487a51b4cce7f38e1a0e194825fc0c04a006a1495b338344295671cdae.
//
// Solidity: event ProposalVote(uint256 _id, address _voter, uint256 _voteResult)
func (_Proposal *ProposalFilterer) FilterProposalVote(opts *bind.FilterOpts) (*ProposalProposalVoteIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return &ProposalProposalVoteIterator{contract: _Proposal.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0xb7d6a4487a51b4cce7f38e1a0e194825fc0c04a006a1495b338344295671cdae.
//
// Solidity: event ProposalVote(uint256 _id, address _voter, uint256 _voteResult)
func (_Proposal *ProposalFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *ProposalProposalVote) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalProposalVote)
				if err := _Proposal.contract.UnpackLog(event, "ProposalVote", log); err != nil {
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

// ParseProposalVote is a log parse operation binding the contract event 0xb7d6a4487a51b4cce7f38e1a0e194825fc0c04a006a1495b338344295671cdae.
//
// Solidity: event ProposalVote(uint256 _id, address _voter, uint256 _voteResult)
func (_Proposal *ProposalFilterer) ParseProposalVote(log types.Log) (*ProposalProposalVote, error) {
	event := new(ProposalProposalVote)
	if err := _Proposal.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
