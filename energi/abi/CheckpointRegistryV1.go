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

// CheckpointRegistryV1ABI is the input ABI used to generate the binding from.
const CheckpointRegistryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// CheckpointRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const CheckpointRegistryV1Bin = `608060405234801561001057600080fd5b506040516103463803806103468339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556102e1806100656000396000f3fe6080604052600436106100335760003560e01c8062f55d9d1461009a578063ce5494bb146100dc578063ec5568891461011c575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100a657600080fd5b506100da600480360360208110156100bd57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661015a565b005b3480156100e857600080fd5b506100da600480360360208110156100ff57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610202565b34801561012857600080fd5b50610131610290565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1633146101e057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6101e98161028d565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60005473ffffffffffffffffffffffffffffffffffffffff16331461028857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61028d815b50565b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a7230582053a0d8900e3815d33eaafad9f0418326b00e339a9f70261605a19b7331cb0cf164736f6c63430005090032`

// DeployCheckpointRegistryV1 deploys a new Ethereum contract, binding an instance of CheckpointRegistryV1 to it.
func DeployCheckpointRegistryV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *CheckpointRegistryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointRegistryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CheckpointRegistryV1Bin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CheckpointRegistryV1{CheckpointRegistryV1Caller: CheckpointRegistryV1Caller{contract: contract}, CheckpointRegistryV1Transactor: CheckpointRegistryV1Transactor{contract: contract}, CheckpointRegistryV1Filterer: CheckpointRegistryV1Filterer{contract: contract}}, nil
}

// CheckpointRegistryV1 is an auto generated Go binding around an Ethereum contract.
type CheckpointRegistryV1 struct {
	CheckpointRegistryV1Caller     // Read-only binding to the contract
	CheckpointRegistryV1Transactor // Write-only binding to the contract
	CheckpointRegistryV1Filterer   // Log filterer for contract events
}

// CheckpointRegistryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type CheckpointRegistryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointRegistryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CheckpointRegistryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointRegistryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CheckpointRegistryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointRegistryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CheckpointRegistryV1Session struct {
	Contract     *CheckpointRegistryV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CheckpointRegistryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CheckpointRegistryV1CallerSession struct {
	Contract *CheckpointRegistryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CheckpointRegistryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CheckpointRegistryV1TransactorSession struct {
	Contract     *CheckpointRegistryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CheckpointRegistryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type CheckpointRegistryV1Raw struct {
	Contract *CheckpointRegistryV1 // Generic contract binding to access the raw methods on
}

// CheckpointRegistryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CheckpointRegistryV1CallerRaw struct {
	Contract *CheckpointRegistryV1Caller // Generic read-only contract binding to access the raw methods on
}

// CheckpointRegistryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CheckpointRegistryV1TransactorRaw struct {
	Contract *CheckpointRegistryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckpointRegistryV1 creates a new instance of CheckpointRegistryV1, bound to a specific deployed contract.
func NewCheckpointRegistryV1(address common.Address, backend bind.ContractBackend) (*CheckpointRegistryV1, error) {
	contract, err := bindCheckpointRegistryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CheckpointRegistryV1{CheckpointRegistryV1Caller: CheckpointRegistryV1Caller{contract: contract}, CheckpointRegistryV1Transactor: CheckpointRegistryV1Transactor{contract: contract}, CheckpointRegistryV1Filterer: CheckpointRegistryV1Filterer{contract: contract}}, nil
}

// NewCheckpointRegistryV1Caller creates a new read-only instance of CheckpointRegistryV1, bound to a specific deployed contract.
func NewCheckpointRegistryV1Caller(address common.Address, caller bind.ContractCaller) (*CheckpointRegistryV1Caller, error) {
	contract, err := bindCheckpointRegistryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointRegistryV1Caller{contract: contract}, nil
}

// NewCheckpointRegistryV1Transactor creates a new write-only instance of CheckpointRegistryV1, bound to a specific deployed contract.
func NewCheckpointRegistryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*CheckpointRegistryV1Transactor, error) {
	contract, err := bindCheckpointRegistryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointRegistryV1Transactor{contract: contract}, nil
}

// NewCheckpointRegistryV1Filterer creates a new log filterer instance of CheckpointRegistryV1, bound to a specific deployed contract.
func NewCheckpointRegistryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*CheckpointRegistryV1Filterer, error) {
	contract, err := bindCheckpointRegistryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckpointRegistryV1Filterer{contract: contract}, nil
}

// bindCheckpointRegistryV1 binds a generic wrapper to an already deployed contract.
func bindCheckpointRegistryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointRegistryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointRegistryV1 *CheckpointRegistryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CheckpointRegistryV1.Contract.CheckpointRegistryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointRegistryV1 *CheckpointRegistryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointRegistryV1.Contract.CheckpointRegistryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointRegistryV1 *CheckpointRegistryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointRegistryV1.Contract.CheckpointRegistryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointRegistryV1 *CheckpointRegistryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CheckpointRegistryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointRegistryV1 *CheckpointRegistryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointRegistryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointRegistryV1 *CheckpointRegistryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointRegistryV1.Contract.contract.Transact(opts, method, params...)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_CheckpointRegistryV1 *CheckpointRegistryV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CheckpointRegistryV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_CheckpointRegistryV1 *CheckpointRegistryV1Session) Proxy() (common.Address, error) {
	return _CheckpointRegistryV1.Contract.Proxy(&_CheckpointRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_CheckpointRegistryV1 *CheckpointRegistryV1CallerSession) Proxy() (common.Address, error) {
	return _CheckpointRegistryV1.Contract.Proxy(&_CheckpointRegistryV1.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_CheckpointRegistryV1 *CheckpointRegistryV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _CheckpointRegistryV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_CheckpointRegistryV1 *CheckpointRegistryV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _CheckpointRegistryV1.Contract.Destroy(&_CheckpointRegistryV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_CheckpointRegistryV1 *CheckpointRegistryV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _CheckpointRegistryV1.Contract.Destroy(&_CheckpointRegistryV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_CheckpointRegistryV1 *CheckpointRegistryV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _CheckpointRegistryV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_CheckpointRegistryV1 *CheckpointRegistryV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _CheckpointRegistryV1.Contract.Migrate(&_CheckpointRegistryV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_CheckpointRegistryV1 *CheckpointRegistryV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _CheckpointRegistryV1.Contract.Migrate(&_CheckpointRegistryV1.TransactOpts, _oldImpl)
}
