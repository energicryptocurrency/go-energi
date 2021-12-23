// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	ethereum "github.com/energicryptocurrency/energi"
	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IProposalABI is the input ABI used to generate the binding from.
const IProposalABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"accepted_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"canVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created_block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_payer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAccepted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quorum_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rejected_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setFee\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteAccept\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteReject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IProposal is an auto generated Go binding around an Ethereum contract.
type IProposal struct {
	IProposalCaller     // Read-only binding to the contract
	IProposalTransactor // Write-only binding to the contract
	IProposalFilterer   // Log filterer for contract events
}

// IProposalCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProposalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProposalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProposalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProposalSession struct {
	Contract     *IProposal        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProposalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProposalCallerSession struct {
	Contract *IProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IProposalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProposalTransactorSession struct {
	Contract     *IProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IProposalRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProposalRaw struct {
	Contract *IProposal // Generic contract binding to access the raw methods on
}

// IProposalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProposalCallerRaw struct {
	Contract *IProposalCaller // Generic read-only contract binding to access the raw methods on
}

// IProposalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProposalTransactorRaw struct {
	Contract *IProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProposal creates a new instance of IProposal, bound to a specific deployed contract.
func NewIProposal(address common.Address, backend bind.ContractBackend) (*IProposal, error) {
	contract, err := bindIProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProposal{IProposalCaller: IProposalCaller{contract: contract}, IProposalTransactor: IProposalTransactor{contract: contract}, IProposalFilterer: IProposalFilterer{contract: contract}}, nil
}

// NewIProposalCaller creates a new read-only instance of IProposal, bound to a specific deployed contract.
func NewIProposalCaller(address common.Address, caller bind.ContractCaller) (*IProposalCaller, error) {
	contract, err := bindIProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProposalCaller{contract: contract}, nil
}

// NewIProposalTransactor creates a new write-only instance of IProposal, bound to a specific deployed contract.
func NewIProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*IProposalTransactor, error) {
	contract, err := bindIProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProposalTransactor{contract: contract}, nil
}

// NewIProposalFilterer creates a new log filterer instance of IProposal, bound to a specific deployed contract.
func NewIProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*IProposalFilterer, error) {
	contract, err := bindIProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProposalFilterer{contract: contract}, nil
}

// bindIProposal binds a generic wrapper to an already deployed contract.
func bindIProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IProposalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProposal *IProposalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IProposal.Contract.IProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProposal *IProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposal.Contract.IProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProposal *IProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProposal.Contract.IProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProposal *IProposalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IProposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProposal *IProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProposal *IProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProposal.Contract.contract.Transact(opts, method, params...)
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_IProposal *IProposalCaller) AcceptedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "accepted_weight")
	return *ret0, err
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_IProposal *IProposalSession) AcceptedWeight() (*big.Int, error) {
	return _IProposal.Contract.AcceptedWeight(&_IProposal.CallOpts)
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_IProposal *IProposalCallerSession) AcceptedWeight() (*big.Int, error) {
	return _IProposal.Contract.AcceptedWeight(&_IProposal.CallOpts)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_IProposal *IProposalCaller) CanVote(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "canVote", owner)
	return *ret0, err
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_IProposal *IProposalSession) CanVote(owner common.Address) (bool, error) {
	return _IProposal.Contract.CanVote(&_IProposal.CallOpts, owner)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_IProposal *IProposalCallerSession) CanVote(owner common.Address) (bool, error) {
	return _IProposal.Contract.CanVote(&_IProposal.CallOpts, owner)
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_IProposal *IProposalCaller) CreatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "created_block")
	return *ret0, err
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_IProposal *IProposalSession) CreatedBlock() (*big.Int, error) {
	return _IProposal.Contract.CreatedBlock(&_IProposal.CallOpts)
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_IProposal *IProposalCallerSession) CreatedBlock() (*big.Int, error) {
	return _IProposal.Contract.CreatedBlock(&_IProposal.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_IProposal *IProposalCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "deadline")
	return *ret0, err
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_IProposal *IProposalSession) Deadline() (*big.Int, error) {
	return _IProposal.Contract.Deadline(&_IProposal.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_IProposal *IProposalCallerSession) Deadline() (*big.Int, error) {
	return _IProposal.Contract.Deadline(&_IProposal.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_IProposal *IProposalCaller) FeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "fee_amount")
	return *ret0, err
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_IProposal *IProposalSession) FeeAmount() (*big.Int, error) {
	return _IProposal.Contract.FeeAmount(&_IProposal.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_IProposal *IProposalCallerSession) FeeAmount() (*big.Int, error) {
	return _IProposal.Contract.FeeAmount(&_IProposal.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_IProposal *IProposalCaller) FeePayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "fee_payer")
	return *ret0, err
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_IProposal *IProposalSession) FeePayer() (common.Address, error) {
	return _IProposal.Contract.FeePayer(&_IProposal.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_IProposal *IProposalCallerSession) FeePayer() (common.Address, error) {
	return _IProposal.Contract.FeePayer(&_IProposal.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_IProposal *IProposalCaller) IsAccepted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "isAccepted")
	return *ret0, err
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_IProposal *IProposalSession) IsAccepted() (bool, error) {
	return _IProposal.Contract.IsAccepted(&_IProposal.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_IProposal *IProposalCallerSession) IsAccepted() (bool, error) {
	return _IProposal.Contract.IsAccepted(&_IProposal.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_IProposal *IProposalCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_IProposal *IProposalSession) IsFinished() (bool, error) {
	return _IProposal.Contract.IsFinished(&_IProposal.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_IProposal *IProposalCallerSession) IsFinished() (bool, error) {
	return _IProposal.Contract.IsFinished(&_IProposal.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_IProposal *IProposalCaller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "parent")
	return *ret0, err
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_IProposal *IProposalSession) Parent() (common.Address, error) {
	return _IProposal.Contract.Parent(&_IProposal.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_IProposal *IProposalCallerSession) Parent() (common.Address, error) {
	return _IProposal.Contract.Parent(&_IProposal.CallOpts)
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_IProposal *IProposalCaller) QuorumWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "quorum_weight")
	return *ret0, err
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_IProposal *IProposalSession) QuorumWeight() (*big.Int, error) {
	return _IProposal.Contract.QuorumWeight(&_IProposal.CallOpts)
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_IProposal *IProposalCallerSession) QuorumWeight() (*big.Int, error) {
	return _IProposal.Contract.QuorumWeight(&_IProposal.CallOpts)
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_IProposal *IProposalCaller) RejectedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "rejected_weight")
	return *ret0, err
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_IProposal *IProposalSession) RejectedWeight() (*big.Int, error) {
	return _IProposal.Contract.RejectedWeight(&_IProposal.CallOpts)
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_IProposal *IProposalCallerSession) RejectedWeight() (*big.Int, error) {
	return _IProposal.Contract.RejectedWeight(&_IProposal.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_IProposal *IProposalCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IProposal.contract.Call(opts, out, "total_weight")
	return *ret0, err
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_IProposal *IProposalSession) TotalWeight() (*big.Int, error) {
	return _IProposal.Contract.TotalWeight(&_IProposal.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_IProposal *IProposalCallerSession) TotalWeight() (*big.Int, error) {
	return _IProposal.Contract.TotalWeight(&_IProposal.CallOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_IProposal *IProposalTransactor) Collect(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposal.contract.Transact(opts, "collect")
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_IProposal *IProposalSession) Collect() (*types.Transaction, error) {
	return _IProposal.Contract.Collect(&_IProposal.TransactOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_IProposal *IProposalTransactorSession) Collect() (*types.Transaction, error) {
	return _IProposal.Contract.Collect(&_IProposal.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IProposal *IProposalTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposal.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IProposal *IProposalSession) Destroy() (*types.Transaction, error) {
	return _IProposal.Contract.Destroy(&_IProposal.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IProposal *IProposalTransactorSession) Destroy() (*types.Transaction, error) {
	return _IProposal.Contract.Destroy(&_IProposal.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_IProposal *IProposalTransactor) SetFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposal.contract.Transact(opts, "setFee")
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_IProposal *IProposalSession) SetFee() (*types.Transaction, error) {
	return _IProposal.Contract.SetFee(&_IProposal.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_IProposal *IProposalTransactorSession) SetFee() (*types.Transaction, error) {
	return _IProposal.Contract.SetFee(&_IProposal.TransactOpts)
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_IProposal *IProposalTransactor) VoteAccept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposal.contract.Transact(opts, "voteAccept")
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_IProposal *IProposalSession) VoteAccept() (*types.Transaction, error) {
	return _IProposal.Contract.VoteAccept(&_IProposal.TransactOpts)
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_IProposal *IProposalTransactorSession) VoteAccept() (*types.Transaction, error) {
	return _IProposal.Contract.VoteAccept(&_IProposal.TransactOpts)
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_IProposal *IProposalTransactor) VoteReject(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposal.contract.Transact(opts, "voteReject")
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_IProposal *IProposalSession) VoteReject() (*types.Transaction, error) {
	return _IProposal.Contract.VoteReject(&_IProposal.TransactOpts)
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_IProposal *IProposalTransactorSession) VoteReject() (*types.Transaction, error) {
	return _IProposal.Contract.VoteReject(&_IProposal.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IProposal *IProposalTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposal.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IProposal *IProposalSession) Withdraw() (*types.Transaction, error) {
	return _IProposal.Contract.Withdraw(&_IProposal.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IProposal *IProposalTransactorSession) Withdraw() (*types.Transaction, error) {
	return _IProposal.Contract.Withdraw(&_IProposal.TransactOpts)
}
