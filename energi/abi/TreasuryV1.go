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

// TreasuryV1ABI is the input ABI used to generate the binding from.
const TreasuryV1ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"block_number\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// TreasuryV1 is an auto generated Go binding around an Ethereum contract.
type TreasuryV1 struct {
	TreasuryV1Caller     // Read-only binding to the contract
	TreasuryV1Transactor // Write-only binding to the contract
	TreasuryV1Filterer   // Log filterer for contract events
}

// TreasuryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasuryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasuryV1Session struct {
	Contract     *TreasuryV1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryV1CallerSession struct {
	Contract *TreasuryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TreasuryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryV1TransactorSession struct {
	Contract     *TreasuryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TreasuryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryV1Raw struct {
	Contract *TreasuryV1 // Generic contract binding to access the raw methods on
}

// TreasuryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryV1CallerRaw struct {
	Contract *TreasuryV1Caller // Generic read-only contract binding to access the raw methods on
}

// TreasuryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryV1TransactorRaw struct {
	Contract *TreasuryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasuryV1 creates a new instance of TreasuryV1, bound to a specific deployed contract.
func NewTreasuryV1(address common.Address, backend bind.ContractBackend) (*TreasuryV1, error) {
	contract, err := bindTreasuryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasuryV1{TreasuryV1Caller: TreasuryV1Caller{contract: contract}, TreasuryV1Transactor: TreasuryV1Transactor{contract: contract}, TreasuryV1Filterer: TreasuryV1Filterer{contract: contract}}, nil
}

// NewTreasuryV1Caller creates a new read-only instance of TreasuryV1, bound to a specific deployed contract.
func NewTreasuryV1Caller(address common.Address, caller bind.ContractCaller) (*TreasuryV1Caller, error) {
	contract, err := bindTreasuryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryV1Caller{contract: contract}, nil
}

// NewTreasuryV1Transactor creates a new write-only instance of TreasuryV1, bound to a specific deployed contract.
func NewTreasuryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryV1Transactor, error) {
	contract, err := bindTreasuryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryV1Transactor{contract: contract}, nil
}

// NewTreasuryV1Filterer creates a new log filterer instance of TreasuryV1, bound to a specific deployed contract.
func NewTreasuryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryV1Filterer, error) {
	contract, err := bindTreasuryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryV1Filterer{contract: contract}, nil
}

// bindTreasuryV1 binds a generic wrapper to an already deployed contract.
func bindTreasuryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryV1 *TreasuryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TreasuryV1.Contract.TreasuryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryV1 *TreasuryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryV1.Contract.TreasuryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryV1 *TreasuryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryV1.Contract.TreasuryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryV1 *TreasuryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TreasuryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryV1 *TreasuryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryV1 *TreasuryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryV1.Contract.contract.Transact(opts, method, params...)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1Caller) GetReward(opts *bind.CallOpts, block_number *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TreasuryV1.contract.Call(opts, out, "getReward", block_number)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1Session) GetReward(block_number *big.Int) (*big.Int, error) {
	return _TreasuryV1.Contract.GetReward(&_TreasuryV1.CallOpts, block_number)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1CallerSession) GetReward(block_number *big.Int) (*big.Int, error) {
	return _TreasuryV1.Contract.GetReward(&_TreasuryV1.CallOpts, block_number)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_TreasuryV1 *TreasuryV1Transactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryV1.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_TreasuryV1 *TreasuryV1Session) Migrate() (*types.Transaction, error) {
	return _TreasuryV1.Contract.Migrate(&_TreasuryV1.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_TreasuryV1 *TreasuryV1TransactorSession) Migrate() (*types.Transaction, error) {
	return _TreasuryV1.Contract.Migrate(&_TreasuryV1.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_TreasuryV1 *TreasuryV1Transactor) Reward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TreasuryV1.contract.Transact(opts, "reward", amount)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_TreasuryV1 *TreasuryV1Session) Reward(amount *big.Int) (*types.Transaction, error) {
	return _TreasuryV1.Contract.Reward(&_TreasuryV1.TransactOpts, amount)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_TreasuryV1 *TreasuryV1TransactorSession) Reward(amount *big.Int) (*types.Transaction, error) {
	return _TreasuryV1.Contract.Reward(&_TreasuryV1.TransactOpts, amount)
}
