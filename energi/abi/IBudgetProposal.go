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

// IBudgetProposalABI is the input ABI used to generate the binding from.
const IBudgetProposalABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"accepted_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"budgetStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_ref_uuid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_is_accepted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_is_finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_unpaid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"canVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created_block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"distributePayout\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_payer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAccepted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paid_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"payout_address\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposed_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quorum_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ref_uuid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rejected_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setFee\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteAccept\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteReject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBudgetProposal is an auto generated Go binding around an Ethereum contract.
type IBudgetProposal struct {
	IBudgetProposalCaller     // Read-only binding to the contract
	IBudgetProposalTransactor // Write-only binding to the contract
	IBudgetProposalFilterer   // Log filterer for contract events
}

// IBudgetProposalCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBudgetProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBudgetProposalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBudgetProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBudgetProposalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBudgetProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBudgetProposalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBudgetProposalSession struct {
	Contract     *IBudgetProposal  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBudgetProposalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBudgetProposalCallerSession struct {
	Contract *IBudgetProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IBudgetProposalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBudgetProposalTransactorSession struct {
	Contract     *IBudgetProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IBudgetProposalRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBudgetProposalRaw struct {
	Contract *IBudgetProposal // Generic contract binding to access the raw methods on
}

// IBudgetProposalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBudgetProposalCallerRaw struct {
	Contract *IBudgetProposalCaller // Generic read-only contract binding to access the raw methods on
}

// IBudgetProposalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBudgetProposalTransactorRaw struct {
	Contract *IBudgetProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBudgetProposal creates a new instance of IBudgetProposal, bound to a specific deployed contract.
func NewIBudgetProposal(address common.Address, backend bind.ContractBackend) (*IBudgetProposal, error) {
	contract, err := bindIBudgetProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBudgetProposal{IBudgetProposalCaller: IBudgetProposalCaller{contract: contract}, IBudgetProposalTransactor: IBudgetProposalTransactor{contract: contract}, IBudgetProposalFilterer: IBudgetProposalFilterer{contract: contract}}, nil
}

// NewIBudgetProposalCaller creates a new read-only instance of IBudgetProposal, bound to a specific deployed contract.
func NewIBudgetProposalCaller(address common.Address, caller bind.ContractCaller) (*IBudgetProposalCaller, error) {
	contract, err := bindIBudgetProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBudgetProposalCaller{contract: contract}, nil
}

// NewIBudgetProposalTransactor creates a new write-only instance of IBudgetProposal, bound to a specific deployed contract.
func NewIBudgetProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*IBudgetProposalTransactor, error) {
	contract, err := bindIBudgetProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBudgetProposalTransactor{contract: contract}, nil
}

// NewIBudgetProposalFilterer creates a new log filterer instance of IBudgetProposal, bound to a specific deployed contract.
func NewIBudgetProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*IBudgetProposalFilterer, error) {
	contract, err := bindIBudgetProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBudgetProposalFilterer{contract: contract}, nil
}

// bindIBudgetProposal binds a generic wrapper to an already deployed contract.
func bindIBudgetProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBudgetProposalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBudgetProposal *IBudgetProposalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBudgetProposal.Contract.IBudgetProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBudgetProposal *IBudgetProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.Contract.IBudgetProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBudgetProposal *IBudgetProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBudgetProposal.Contract.IBudgetProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBudgetProposal *IBudgetProposalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBudgetProposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBudgetProposal *IBudgetProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBudgetProposal *IBudgetProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBudgetProposal.Contract.contract.Transact(opts, method, params...)
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) AcceptedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "accepted_weight")
	return *ret0, err
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) AcceptedWeight() (*big.Int, error) {
	return _IBudgetProposal.Contract.AcceptedWeight(&_IBudgetProposal.CallOpts)
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) AcceptedWeight() (*big.Int, error) {
	return _IBudgetProposal.Contract.AcceptedWeight(&_IBudgetProposal.CallOpts)
}

// BudgetStatus is a free data retrieval call binding the contract method 0x3b2a1b14.
//
// Solidity: function budgetStatus() constant returns(uint256 _ref_uuid, bool _is_accepted, bool _is_finished, uint256 _unpaid)
func (_IBudgetProposal *IBudgetProposalCaller) BudgetStatus(opts *bind.CallOpts) (struct {
	RefUuid    *big.Int
	IsAccepted bool
	IsFinished bool
	Unpaid     *big.Int
}, error) {
	ret := new(struct {
		RefUuid    *big.Int
		IsAccepted bool
		IsFinished bool
		Unpaid     *big.Int
	})
	out := ret
	err := _IBudgetProposal.contract.Call(opts, out, "budgetStatus")
	return *ret, err
}

// BudgetStatus is a free data retrieval call binding the contract method 0x3b2a1b14.
//
// Solidity: function budgetStatus() constant returns(uint256 _ref_uuid, bool _is_accepted, bool _is_finished, uint256 _unpaid)
func (_IBudgetProposal *IBudgetProposalSession) BudgetStatus() (struct {
	RefUuid    *big.Int
	IsAccepted bool
	IsFinished bool
	Unpaid     *big.Int
}, error) {
	return _IBudgetProposal.Contract.BudgetStatus(&_IBudgetProposal.CallOpts)
}

// BudgetStatus is a free data retrieval call binding the contract method 0x3b2a1b14.
//
// Solidity: function budgetStatus() constant returns(uint256 _ref_uuid, bool _is_accepted, bool _is_finished, uint256 _unpaid)
func (_IBudgetProposal *IBudgetProposalCallerSession) BudgetStatus() (struct {
	RefUuid    *big.Int
	IsAccepted bool
	IsFinished bool
	Unpaid     *big.Int
}, error) {
	return _IBudgetProposal.Contract.BudgetStatus(&_IBudgetProposal.CallOpts)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_IBudgetProposal *IBudgetProposalCaller) CanVote(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "canVote", owner)
	return *ret0, err
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_IBudgetProposal *IBudgetProposalSession) CanVote(owner common.Address) (bool, error) {
	return _IBudgetProposal.Contract.CanVote(&_IBudgetProposal.CallOpts, owner)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_IBudgetProposal *IBudgetProposalCallerSession) CanVote(owner common.Address) (bool, error) {
	return _IBudgetProposal.Contract.CanVote(&_IBudgetProposal.CallOpts, owner)
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) CreatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "created_block")
	return *ret0, err
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) CreatedBlock() (*big.Int, error) {
	return _IBudgetProposal.Contract.CreatedBlock(&_IBudgetProposal.CallOpts)
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) CreatedBlock() (*big.Int, error) {
	return _IBudgetProposal.Contract.CreatedBlock(&_IBudgetProposal.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "deadline")
	return *ret0, err
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) Deadline() (*big.Int, error) {
	return _IBudgetProposal.Contract.Deadline(&_IBudgetProposal.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) Deadline() (*big.Int, error) {
	return _IBudgetProposal.Contract.Deadline(&_IBudgetProposal.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) FeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "fee_amount")
	return *ret0, err
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) FeeAmount() (*big.Int, error) {
	return _IBudgetProposal.Contract.FeeAmount(&_IBudgetProposal.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) FeeAmount() (*big.Int, error) {
	return _IBudgetProposal.Contract.FeeAmount(&_IBudgetProposal.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_IBudgetProposal *IBudgetProposalCaller) FeePayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "fee_payer")
	return *ret0, err
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_IBudgetProposal *IBudgetProposalSession) FeePayer() (common.Address, error) {
	return _IBudgetProposal.Contract.FeePayer(&_IBudgetProposal.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_IBudgetProposal *IBudgetProposalCallerSession) FeePayer() (common.Address, error) {
	return _IBudgetProposal.Contract.FeePayer(&_IBudgetProposal.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_IBudgetProposal *IBudgetProposalCaller) IsAccepted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "isAccepted")
	return *ret0, err
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_IBudgetProposal *IBudgetProposalSession) IsAccepted() (bool, error) {
	return _IBudgetProposal.Contract.IsAccepted(&_IBudgetProposal.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_IBudgetProposal *IBudgetProposalCallerSession) IsAccepted() (bool, error) {
	return _IBudgetProposal.Contract.IsAccepted(&_IBudgetProposal.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_IBudgetProposal *IBudgetProposalCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_IBudgetProposal *IBudgetProposalSession) IsFinished() (bool, error) {
	return _IBudgetProposal.Contract.IsFinished(&_IBudgetProposal.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_IBudgetProposal *IBudgetProposalCallerSession) IsFinished() (bool, error) {
	return _IBudgetProposal.Contract.IsFinished(&_IBudgetProposal.CallOpts)
}

// PaidAmount is a free data retrieval call binding the contract method 0x504881df.
//
// Solidity: function paid_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) PaidAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "paid_amount")
	return *ret0, err
}

// PaidAmount is a free data retrieval call binding the contract method 0x504881df.
//
// Solidity: function paid_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) PaidAmount() (*big.Int, error) {
	return _IBudgetProposal.Contract.PaidAmount(&_IBudgetProposal.CallOpts)
}

// PaidAmount is a free data retrieval call binding the contract method 0x504881df.
//
// Solidity: function paid_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) PaidAmount() (*big.Int, error) {
	return _IBudgetProposal.Contract.PaidAmount(&_IBudgetProposal.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_IBudgetProposal *IBudgetProposalCaller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "parent")
	return *ret0, err
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_IBudgetProposal *IBudgetProposalSession) Parent() (common.Address, error) {
	return _IBudgetProposal.Contract.Parent(&_IBudgetProposal.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_IBudgetProposal *IBudgetProposalCallerSession) Parent() (common.Address, error) {
	return _IBudgetProposal.Contract.Parent(&_IBudgetProposal.CallOpts)
}

// PayoutAddress is a free data retrieval call binding the contract method 0x9d5e6c9d.
//
// Solidity: function payout_address() constant returns(address)
func (_IBudgetProposal *IBudgetProposalCaller) PayoutAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "payout_address")
	return *ret0, err
}

// PayoutAddress is a free data retrieval call binding the contract method 0x9d5e6c9d.
//
// Solidity: function payout_address() constant returns(address)
func (_IBudgetProposal *IBudgetProposalSession) PayoutAddress() (common.Address, error) {
	return _IBudgetProposal.Contract.PayoutAddress(&_IBudgetProposal.CallOpts)
}

// PayoutAddress is a free data retrieval call binding the contract method 0x9d5e6c9d.
//
// Solidity: function payout_address() constant returns(address)
func (_IBudgetProposal *IBudgetProposalCallerSession) PayoutAddress() (common.Address, error) {
	return _IBudgetProposal.Contract.PayoutAddress(&_IBudgetProposal.CallOpts)
}

// ProposedAmount is a free data retrieval call binding the contract method 0x4cafdfb2.
//
// Solidity: function proposed_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) ProposedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "proposed_amount")
	return *ret0, err
}

// ProposedAmount is a free data retrieval call binding the contract method 0x4cafdfb2.
//
// Solidity: function proposed_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) ProposedAmount() (*big.Int, error) {
	return _IBudgetProposal.Contract.ProposedAmount(&_IBudgetProposal.CallOpts)
}

// ProposedAmount is a free data retrieval call binding the contract method 0x4cafdfb2.
//
// Solidity: function proposed_amount() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) ProposedAmount() (*big.Int, error) {
	return _IBudgetProposal.Contract.ProposedAmount(&_IBudgetProposal.CallOpts)
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) QuorumWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "quorum_weight")
	return *ret0, err
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) QuorumWeight() (*big.Int, error) {
	return _IBudgetProposal.Contract.QuorumWeight(&_IBudgetProposal.CallOpts)
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) QuorumWeight() (*big.Int, error) {
	return _IBudgetProposal.Contract.QuorumWeight(&_IBudgetProposal.CallOpts)
}

// RefUuid is a free data retrieval call binding the contract method 0xbd4c1f39.
//
// Solidity: function ref_uuid() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) RefUuid(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "ref_uuid")
	return *ret0, err
}

// RefUuid is a free data retrieval call binding the contract method 0xbd4c1f39.
//
// Solidity: function ref_uuid() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) RefUuid() (*big.Int, error) {
	return _IBudgetProposal.Contract.RefUuid(&_IBudgetProposal.CallOpts)
}

// RefUuid is a free data retrieval call binding the contract method 0xbd4c1f39.
//
// Solidity: function ref_uuid() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) RefUuid() (*big.Int, error) {
	return _IBudgetProposal.Contract.RefUuid(&_IBudgetProposal.CallOpts)
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) RejectedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "rejected_weight")
	return *ret0, err
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) RejectedWeight() (*big.Int, error) {
	return _IBudgetProposal.Contract.RejectedWeight(&_IBudgetProposal.CallOpts)
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) RejectedWeight() (*big.Int, error) {
	return _IBudgetProposal.Contract.RejectedWeight(&_IBudgetProposal.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBudgetProposal.contract.Call(opts, out, "total_weight")
	return *ret0, err
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalSession) TotalWeight() (*big.Int, error) {
	return _IBudgetProposal.Contract.TotalWeight(&_IBudgetProposal.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_IBudgetProposal *IBudgetProposalCallerSession) TotalWeight() (*big.Int, error) {
	return _IBudgetProposal.Contract.TotalWeight(&_IBudgetProposal.CallOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_IBudgetProposal *IBudgetProposalTransactor) Collect(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.contract.Transact(opts, "collect")
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_IBudgetProposal *IBudgetProposalSession) Collect() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.Collect(&_IBudgetProposal.TransactOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_IBudgetProposal *IBudgetProposalTransactorSession) Collect() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.Collect(&_IBudgetProposal.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IBudgetProposal *IBudgetProposalTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IBudgetProposal *IBudgetProposalSession) Destroy() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.Destroy(&_IBudgetProposal.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IBudgetProposal *IBudgetProposalTransactorSession) Destroy() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.Destroy(&_IBudgetProposal.TransactOpts)
}

// DistributePayout is a paid mutator transaction binding the contract method 0x10cac8a5.
//
// Solidity: function distributePayout() returns()
func (_IBudgetProposal *IBudgetProposalTransactor) DistributePayout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.contract.Transact(opts, "distributePayout")
}

// DistributePayout is a paid mutator transaction binding the contract method 0x10cac8a5.
//
// Solidity: function distributePayout() returns()
func (_IBudgetProposal *IBudgetProposalSession) DistributePayout() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.DistributePayout(&_IBudgetProposal.TransactOpts)
}

// DistributePayout is a paid mutator transaction binding the contract method 0x10cac8a5.
//
// Solidity: function distributePayout() returns()
func (_IBudgetProposal *IBudgetProposalTransactorSession) DistributePayout() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.DistributePayout(&_IBudgetProposal.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_IBudgetProposal *IBudgetProposalTransactor) SetFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.contract.Transact(opts, "setFee")
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_IBudgetProposal *IBudgetProposalSession) SetFee() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.SetFee(&_IBudgetProposal.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_IBudgetProposal *IBudgetProposalTransactorSession) SetFee() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.SetFee(&_IBudgetProposal.TransactOpts)
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_IBudgetProposal *IBudgetProposalTransactor) VoteAccept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.contract.Transact(opts, "voteAccept")
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_IBudgetProposal *IBudgetProposalSession) VoteAccept() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.VoteAccept(&_IBudgetProposal.TransactOpts)
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_IBudgetProposal *IBudgetProposalTransactorSession) VoteAccept() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.VoteAccept(&_IBudgetProposal.TransactOpts)
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_IBudgetProposal *IBudgetProposalTransactor) VoteReject(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.contract.Transact(opts, "voteReject")
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_IBudgetProposal *IBudgetProposalSession) VoteReject() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.VoteReject(&_IBudgetProposal.TransactOpts)
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_IBudgetProposal *IBudgetProposalTransactorSession) VoteReject() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.VoteReject(&_IBudgetProposal.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IBudgetProposal *IBudgetProposalTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBudgetProposal.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IBudgetProposal *IBudgetProposalSession) Withdraw() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.Withdraw(&_IBudgetProposal.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IBudgetProposal *IBudgetProposalTransactorSession) Withdraw() (*types.Transaction, error) {
	return _IBudgetProposal.Contract.Withdraw(&_IBudgetProposal.TransactOpts)
}
