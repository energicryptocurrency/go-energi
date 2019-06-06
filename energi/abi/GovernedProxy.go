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

// GovernedProxyABI is the input ABI used to generate the binding from.
const GovernedProxyABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GovernedProxyBin is the compiled bytecode used for deploying new contracts.
const GovernedProxyBin = `6080604052348015600f57600080fd5b50606c80601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80638fd3ab8014602d575b600080fd5b60336035565b005b56fea265627a7a7230582028e48d6d42662dcb11add9f787438b2d43c821de6572ec717e857dbdb15c644164736f6c63430005090032`

// DeployGovernedProxy deploys a new Ethereum contract, binding an instance of GovernedProxy to it.
func DeployGovernedProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GovernedProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernedProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernedProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernedProxy{GovernedProxyCaller: GovernedProxyCaller{contract: contract}, GovernedProxyTransactor: GovernedProxyTransactor{contract: contract}, GovernedProxyFilterer: GovernedProxyFilterer{contract: contract}}, nil
}

// GovernedProxy is an auto generated Go binding around an Ethereum contract.
type GovernedProxy struct {
	GovernedProxyCaller     // Read-only binding to the contract
	GovernedProxyTransactor // Write-only binding to the contract
	GovernedProxyFilterer   // Log filterer for contract events
}

// GovernedProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernedProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernedProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernedProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernedProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernedProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernedProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernedProxySession struct {
	Contract     *GovernedProxy    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernedProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernedProxyCallerSession struct {
	Contract *GovernedProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GovernedProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernedProxyTransactorSession struct {
	Contract     *GovernedProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GovernedProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernedProxyRaw struct {
	Contract *GovernedProxy // Generic contract binding to access the raw methods on
}

// GovernedProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernedProxyCallerRaw struct {
	Contract *GovernedProxyCaller // Generic read-only contract binding to access the raw methods on
}

// GovernedProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernedProxyTransactorRaw struct {
	Contract *GovernedProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernedProxy creates a new instance of GovernedProxy, bound to a specific deployed contract.
func NewGovernedProxy(address common.Address, backend bind.ContractBackend) (*GovernedProxy, error) {
	contract, err := bindGovernedProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernedProxy{GovernedProxyCaller: GovernedProxyCaller{contract: contract}, GovernedProxyTransactor: GovernedProxyTransactor{contract: contract}, GovernedProxyFilterer: GovernedProxyFilterer{contract: contract}}, nil
}

// NewGovernedProxyCaller creates a new read-only instance of GovernedProxy, bound to a specific deployed contract.
func NewGovernedProxyCaller(address common.Address, caller bind.ContractCaller) (*GovernedProxyCaller, error) {
	contract, err := bindGovernedProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyCaller{contract: contract}, nil
}

// NewGovernedProxyTransactor creates a new write-only instance of GovernedProxy, bound to a specific deployed contract.
func NewGovernedProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernedProxyTransactor, error) {
	contract, err := bindGovernedProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyTransactor{contract: contract}, nil
}

// NewGovernedProxyFilterer creates a new log filterer instance of GovernedProxy, bound to a specific deployed contract.
func NewGovernedProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernedProxyFilterer, error) {
	contract, err := bindGovernedProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyFilterer{contract: contract}, nil
}

// bindGovernedProxy binds a generic wrapper to an already deployed contract.
func bindGovernedProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernedProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernedProxy *GovernedProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernedProxy.Contract.GovernedProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernedProxy *GovernedProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernedProxy.Contract.GovernedProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernedProxy *GovernedProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernedProxy.Contract.GovernedProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernedProxy *GovernedProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernedProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernedProxy *GovernedProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernedProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernedProxy *GovernedProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernedProxy.Contract.contract.Transact(opts, method, params...)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_GovernedProxy *GovernedProxyTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_GovernedProxy *GovernedProxySession) Migrate() (*types.Transaction, error) {
	return _GovernedProxy.Contract.Migrate(&_GovernedProxy.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_GovernedProxy *GovernedProxyTransactorSession) Migrate() (*types.Transaction, error) {
	return _GovernedProxy.Contract.Migrate(&_GovernedProxy.TransactOpts)
}
