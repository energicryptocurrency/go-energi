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

// SporkRegistryV1ABI is the input ABI used to generate the binding from.
const SporkRegistryV1ABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SporkRegistryV1 is an auto generated Go binding around an Ethereum contract.
type SporkRegistryV1 struct {
	SporkRegistryV1Caller     // Read-only binding to the contract
	SporkRegistryV1Transactor // Write-only binding to the contract
	SporkRegistryV1Filterer   // Log filterer for contract events
}

// SporkRegistryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SporkRegistryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SporkRegistryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SporkRegistryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SporkRegistryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SporkRegistryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SporkRegistryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SporkRegistryV1Session struct {
	Contract     *SporkRegistryV1  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SporkRegistryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SporkRegistryV1CallerSession struct {
	Contract *SporkRegistryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SporkRegistryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SporkRegistryV1TransactorSession struct {
	Contract     *SporkRegistryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SporkRegistryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SporkRegistryV1Raw struct {
	Contract *SporkRegistryV1 // Generic contract binding to access the raw methods on
}

// SporkRegistryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SporkRegistryV1CallerRaw struct {
	Contract *SporkRegistryV1Caller // Generic read-only contract binding to access the raw methods on
}

// SporkRegistryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SporkRegistryV1TransactorRaw struct {
	Contract *SporkRegistryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSporkRegistryV1 creates a new instance of SporkRegistryV1, bound to a specific deployed contract.
func NewSporkRegistryV1(address common.Address, backend bind.ContractBackend) (*SporkRegistryV1, error) {
	contract, err := bindSporkRegistryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SporkRegistryV1{SporkRegistryV1Caller: SporkRegistryV1Caller{contract: contract}, SporkRegistryV1Transactor: SporkRegistryV1Transactor{contract: contract}, SporkRegistryV1Filterer: SporkRegistryV1Filterer{contract: contract}}, nil
}

// NewSporkRegistryV1Caller creates a new read-only instance of SporkRegistryV1, bound to a specific deployed contract.
func NewSporkRegistryV1Caller(address common.Address, caller bind.ContractCaller) (*SporkRegistryV1Caller, error) {
	contract, err := bindSporkRegistryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SporkRegistryV1Caller{contract: contract}, nil
}

// NewSporkRegistryV1Transactor creates a new write-only instance of SporkRegistryV1, bound to a specific deployed contract.
func NewSporkRegistryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SporkRegistryV1Transactor, error) {
	contract, err := bindSporkRegistryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SporkRegistryV1Transactor{contract: contract}, nil
}

// NewSporkRegistryV1Filterer creates a new log filterer instance of SporkRegistryV1, bound to a specific deployed contract.
func NewSporkRegistryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SporkRegistryV1Filterer, error) {
	contract, err := bindSporkRegistryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SporkRegistryV1Filterer{contract: contract}, nil
}

// bindSporkRegistryV1 binds a generic wrapper to an already deployed contract.
func bindSporkRegistryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SporkRegistryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SporkRegistryV1 *SporkRegistryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SporkRegistryV1.Contract.SporkRegistryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SporkRegistryV1 *SporkRegistryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.SporkRegistryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SporkRegistryV1 *SporkRegistryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.SporkRegistryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SporkRegistryV1 *SporkRegistryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SporkRegistryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SporkRegistryV1 *SporkRegistryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SporkRegistryV1 *SporkRegistryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.contract.Transact(opts, method, params...)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SporkRegistryV1 *SporkRegistryV1Transactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SporkRegistryV1.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SporkRegistryV1 *SporkRegistryV1Session) Migrate() (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Migrate(&_SporkRegistryV1.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SporkRegistryV1 *SporkRegistryV1TransactorSession) Migrate() (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Migrate(&_SporkRegistryV1.TransactOpts)
}
