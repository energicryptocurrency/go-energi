// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakerRewardV1ABI is the input ABI used to generate the binding from.
const StakerRewardV1ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"block_number\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// StakerRewardV1 is an auto generated Go binding around an Ethereum contract.
type StakerRewardV1 struct {
	StakerRewardV1Caller     // Read-only binding to the contract
	StakerRewardV1Transactor // Write-only binding to the contract
	StakerRewardV1Filterer   // Log filterer for contract events
}

// StakerRewardV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type StakerRewardV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerRewardV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StakerRewardV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerRewardV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakerRewardV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerRewardV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakerRewardV1Session struct {
	Contract     *StakerRewardV1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerRewardV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakerRewardV1CallerSession struct {
	Contract *StakerRewardV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakerRewardV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakerRewardV1TransactorSession struct {
	Contract     *StakerRewardV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakerRewardV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type StakerRewardV1Raw struct {
	Contract *StakerRewardV1 // Generic contract binding to access the raw methods on
}

// StakerRewardV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakerRewardV1CallerRaw struct {
	Contract *StakerRewardV1Caller // Generic read-only contract binding to access the raw methods on
}

// StakerRewardV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakerRewardV1TransactorRaw struct {
	Contract *StakerRewardV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStakerRewardV1 creates a new instance of StakerRewardV1, bound to a specific deployed contract.
func NewStakerRewardV1(address common.Address, backend bind.ContractBackend) (*StakerRewardV1, error) {
	contract, err := bindStakerRewardV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakerRewardV1{StakerRewardV1Caller: StakerRewardV1Caller{contract: contract}, StakerRewardV1Transactor: StakerRewardV1Transactor{contract: contract}, StakerRewardV1Filterer: StakerRewardV1Filterer{contract: contract}}, nil
}

// NewStakerRewardV1Caller creates a new read-only instance of StakerRewardV1, bound to a specific deployed contract.
func NewStakerRewardV1Caller(address common.Address, caller bind.ContractCaller) (*StakerRewardV1Caller, error) {
	contract, err := bindStakerRewardV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakerRewardV1Caller{contract: contract}, nil
}

// NewStakerRewardV1Transactor creates a new write-only instance of StakerRewardV1, bound to a specific deployed contract.
func NewStakerRewardV1Transactor(address common.Address, transactor bind.ContractTransactor) (*StakerRewardV1Transactor, error) {
	contract, err := bindStakerRewardV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakerRewardV1Transactor{contract: contract}, nil
}

// NewStakerRewardV1Filterer creates a new log filterer instance of StakerRewardV1, bound to a specific deployed contract.
func NewStakerRewardV1Filterer(address common.Address, filterer bind.ContractFilterer) (*StakerRewardV1Filterer, error) {
	contract, err := bindStakerRewardV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakerRewardV1Filterer{contract: contract}, nil
}

// bindStakerRewardV1 binds a generic wrapper to an already deployed contract.
func bindStakerRewardV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakerRewardV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakerRewardV1 *StakerRewardV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakerRewardV1.Contract.StakerRewardV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakerRewardV1 *StakerRewardV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.StakerRewardV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakerRewardV1 *StakerRewardV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.StakerRewardV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakerRewardV1 *StakerRewardV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakerRewardV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakerRewardV1 *StakerRewardV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakerRewardV1 *StakerRewardV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.contract.Transact(opts, method, params...)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_StakerRewardV1 *StakerRewardV1Caller) GetReward(opts *bind.CallOpts, block_number *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakerRewardV1.contract.Call(opts, out, "getReward", block_number)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_StakerRewardV1 *StakerRewardV1Session) GetReward(block_number *big.Int) (*big.Int, error) {
	return _StakerRewardV1.Contract.GetReward(&_StakerRewardV1.CallOpts, block_number)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_StakerRewardV1 *StakerRewardV1CallerSession) GetReward(block_number *big.Int) (*big.Int, error) {
	return _StakerRewardV1.Contract.GetReward(&_StakerRewardV1.CallOpts, block_number)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_StakerRewardV1 *StakerRewardV1Transactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerRewardV1.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_StakerRewardV1 *StakerRewardV1Session) Migrate() (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Migrate(&_StakerRewardV1.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_StakerRewardV1 *StakerRewardV1TransactorSession) Migrate() (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Migrate(&_StakerRewardV1.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_StakerRewardV1 *StakerRewardV1Transactor) Reward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakerRewardV1.contract.Transact(opts, "reward", amount)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_StakerRewardV1 *StakerRewardV1Session) Reward(amount *big.Int) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Reward(&_StakerRewardV1.TransactOpts, amount)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_StakerRewardV1 *StakerRewardV1TransactorSession) Reward(amount *big.Int) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Reward(&_StakerRewardV1.TransactOpts, amount)
}
