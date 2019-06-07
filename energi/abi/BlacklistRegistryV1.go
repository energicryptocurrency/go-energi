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

// BlacklistRegistryV1ABI is the input ABI used to generate the binding from.
const BlacklistRegistryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// BlacklistRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const BlacklistRegistryV1Bin = `6080604052348015600f57600080fd5b5060808061001e6000396000f3fe60806040526004361060255760003560e01c8062f55d9d146027578063ce5494bb146027575b005b348015603257600080fd5b50602560048036036020811015604757600080fd5b505056fea265627a7a7230582033d4b55ff7a9b94ab599d8b27ef87ce8be70c8074440d37b47b0e8bca9ffc74964736f6c63430005090032`

// DeployBlacklistRegistryV1 deploys a new Ethereum contract, binding an instance of BlacklistRegistryV1 to it.
func DeployBlacklistRegistryV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlacklistRegistryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(BlacklistRegistryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlacklistRegistryV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlacklistRegistryV1{BlacklistRegistryV1Caller: BlacklistRegistryV1Caller{contract: contract}, BlacklistRegistryV1Transactor: BlacklistRegistryV1Transactor{contract: contract}, BlacklistRegistryV1Filterer: BlacklistRegistryV1Filterer{contract: contract}}, nil
}

// BlacklistRegistryV1 is an auto generated Go binding around an Ethereum contract.
type BlacklistRegistryV1 struct {
	BlacklistRegistryV1Caller     // Read-only binding to the contract
	BlacklistRegistryV1Transactor // Write-only binding to the contract
	BlacklistRegistryV1Filterer   // Log filterer for contract events
}

// BlacklistRegistryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type BlacklistRegistryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlacklistRegistryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BlacklistRegistryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlacklistRegistryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlacklistRegistryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlacklistRegistryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlacklistRegistryV1Session struct {
	Contract     *BlacklistRegistryV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlacklistRegistryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlacklistRegistryV1CallerSession struct {
	Contract *BlacklistRegistryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BlacklistRegistryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlacklistRegistryV1TransactorSession struct {
	Contract     *BlacklistRegistryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BlacklistRegistryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type BlacklistRegistryV1Raw struct {
	Contract *BlacklistRegistryV1 // Generic contract binding to access the raw methods on
}

// BlacklistRegistryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlacklistRegistryV1CallerRaw struct {
	Contract *BlacklistRegistryV1Caller // Generic read-only contract binding to access the raw methods on
}

// BlacklistRegistryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlacklistRegistryV1TransactorRaw struct {
	Contract *BlacklistRegistryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBlacklistRegistryV1 creates a new instance of BlacklistRegistryV1, bound to a specific deployed contract.
func NewBlacklistRegistryV1(address common.Address, backend bind.ContractBackend) (*BlacklistRegistryV1, error) {
	contract, err := bindBlacklistRegistryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlacklistRegistryV1{BlacklistRegistryV1Caller: BlacklistRegistryV1Caller{contract: contract}, BlacklistRegistryV1Transactor: BlacklistRegistryV1Transactor{contract: contract}, BlacklistRegistryV1Filterer: BlacklistRegistryV1Filterer{contract: contract}}, nil
}

// NewBlacklistRegistryV1Caller creates a new read-only instance of BlacklistRegistryV1, bound to a specific deployed contract.
func NewBlacklistRegistryV1Caller(address common.Address, caller bind.ContractCaller) (*BlacklistRegistryV1Caller, error) {
	contract, err := bindBlacklistRegistryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlacklistRegistryV1Caller{contract: contract}, nil
}

// NewBlacklistRegistryV1Transactor creates a new write-only instance of BlacklistRegistryV1, bound to a specific deployed contract.
func NewBlacklistRegistryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*BlacklistRegistryV1Transactor, error) {
	contract, err := bindBlacklistRegistryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlacklistRegistryV1Transactor{contract: contract}, nil
}

// NewBlacklistRegistryV1Filterer creates a new log filterer instance of BlacklistRegistryV1, bound to a specific deployed contract.
func NewBlacklistRegistryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*BlacklistRegistryV1Filterer, error) {
	contract, err := bindBlacklistRegistryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlacklistRegistryV1Filterer{contract: contract}, nil
}

// bindBlacklistRegistryV1 binds a generic wrapper to an already deployed contract.
func bindBlacklistRegistryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlacklistRegistryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlacklistRegistryV1 *BlacklistRegistryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlacklistRegistryV1.Contract.BlacklistRegistryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlacklistRegistryV1 *BlacklistRegistryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.BlacklistRegistryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlacklistRegistryV1 *BlacklistRegistryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.BlacklistRegistryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlacklistRegistryV1 *BlacklistRegistryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlacklistRegistryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlacklistRegistryV1 *BlacklistRegistryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlacklistRegistryV1 *BlacklistRegistryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.contract.Transact(opts, method, params...)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1Transactor) Destroy(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.contract.Transact(opts, "destroy", arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1Session) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.Destroy(&_BlacklistRegistryV1.TransactOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1TransactorSession) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.Destroy(&_BlacklistRegistryV1.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1Transactor) Migrate(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.contract.Transact(opts, "migrate", arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1Session) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.Migrate(&_BlacklistRegistryV1.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1TransactorSession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.Migrate(&_BlacklistRegistryV1.TransactOpts, arg0)
}
