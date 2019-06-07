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

// MasternodeTokenV1ABI is the input ABI used to generate the binding from.
const MasternodeTokenV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// MasternodeTokenV1Bin is the compiled bytecode used for deploying new contracts.
const MasternodeTokenV1Bin = `6080604052348015600f57600080fd5b5060808061001e6000396000f3fe60806040526004361060255760003560e01c8062f55d9d146027578063ce5494bb146027575b005b348015603257600080fd5b50602560048036036020811015604757600080fd5b505056fea265627a7a72305820e11e752471db3ebbd1cd430af499c48bef65d4fbf01edc323bc766ecacd190b464736f6c63430005090032`

// DeployMasternodeTokenV1 deploys a new Ethereum contract, binding an instance of MasternodeTokenV1 to it.
func DeployMasternodeTokenV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MasternodeTokenV1, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeTokenV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MasternodeTokenV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MasternodeTokenV1{MasternodeTokenV1Caller: MasternodeTokenV1Caller{contract: contract}, MasternodeTokenV1Transactor: MasternodeTokenV1Transactor{contract: contract}, MasternodeTokenV1Filterer: MasternodeTokenV1Filterer{contract: contract}}, nil
}

// MasternodeTokenV1 is an auto generated Go binding around an Ethereum contract.
type MasternodeTokenV1 struct {
	MasternodeTokenV1Caller     // Read-only binding to the contract
	MasternodeTokenV1Transactor // Write-only binding to the contract
	MasternodeTokenV1Filterer   // Log filterer for contract events
}

// MasternodeTokenV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasternodeTokenV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasternodeTokenV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasternodeTokenV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasternodeTokenV1Session struct {
	Contract     *MasternodeTokenV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MasternodeTokenV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasternodeTokenV1CallerSession struct {
	Contract *MasternodeTokenV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MasternodeTokenV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasternodeTokenV1TransactorSession struct {
	Contract     *MasternodeTokenV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MasternodeTokenV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasternodeTokenV1Raw struct {
	Contract *MasternodeTokenV1 // Generic contract binding to access the raw methods on
}

// MasternodeTokenV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasternodeTokenV1CallerRaw struct {
	Contract *MasternodeTokenV1Caller // Generic read-only contract binding to access the raw methods on
}

// MasternodeTokenV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasternodeTokenV1TransactorRaw struct {
	Contract *MasternodeTokenV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasternodeTokenV1 creates a new instance of MasternodeTokenV1, bound to a specific deployed contract.
func NewMasternodeTokenV1(address common.Address, backend bind.ContractBackend) (*MasternodeTokenV1, error) {
	contract, err := bindMasternodeTokenV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1{MasternodeTokenV1Caller: MasternodeTokenV1Caller{contract: contract}, MasternodeTokenV1Transactor: MasternodeTokenV1Transactor{contract: contract}, MasternodeTokenV1Filterer: MasternodeTokenV1Filterer{contract: contract}}, nil
}

// NewMasternodeTokenV1Caller creates a new read-only instance of MasternodeTokenV1, bound to a specific deployed contract.
func NewMasternodeTokenV1Caller(address common.Address, caller bind.ContractCaller) (*MasternodeTokenV1Caller, error) {
	contract, err := bindMasternodeTokenV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1Caller{contract: contract}, nil
}

// NewMasternodeTokenV1Transactor creates a new write-only instance of MasternodeTokenV1, bound to a specific deployed contract.
func NewMasternodeTokenV1Transactor(address common.Address, transactor bind.ContractTransactor) (*MasternodeTokenV1Transactor, error) {
	contract, err := bindMasternodeTokenV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1Transactor{contract: contract}, nil
}

// NewMasternodeTokenV1Filterer creates a new log filterer instance of MasternodeTokenV1, bound to a specific deployed contract.
func NewMasternodeTokenV1Filterer(address common.Address, filterer bind.ContractFilterer) (*MasternodeTokenV1Filterer, error) {
	contract, err := bindMasternodeTokenV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1Filterer{contract: contract}, nil
}

// bindMasternodeTokenV1 binds a generic wrapper to an already deployed contract.
func bindMasternodeTokenV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeTokenV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeTokenV1 *MasternodeTokenV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeTokenV1.Contract.MasternodeTokenV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeTokenV1 *MasternodeTokenV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.MasternodeTokenV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeTokenV1 *MasternodeTokenV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.MasternodeTokenV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeTokenV1 *MasternodeTokenV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeTokenV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.contract.Transact(opts, method, params...)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) Destroy(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "destroy", arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Destroy(&_MasternodeTokenV1.TransactOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Destroy(&_MasternodeTokenV1.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) Migrate(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "migrate", arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Migrate(&_MasternodeTokenV1.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Migrate(&_MasternodeTokenV1.TransactOpts, arg0)
}
