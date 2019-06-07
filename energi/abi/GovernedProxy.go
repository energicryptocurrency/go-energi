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
const GovernedProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgrade_proposals\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_impl\",\"type\":\"address\"},{\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"proposeUpgrade\",\"outputs\":[{\"name\":\"proposal\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_impl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"impl\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// GovernedProxyBin is the compiled bytecode used for deploying new contracts.
const GovernedProxyBin = `608060405234801561001057600080fd5b506040516108073803806108078339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556107a2806100656000396000f3fe6080604052600436106100645760003560e01c80635b6dee4c116100435780635b6dee4c14610202578063ad797cdd1461023b578063ce5494bb1461011757610064565b8062f55d9d146101175780630900f0101461015957806332e3a90514610199575b3233146100d257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f6e6c79206469726563742063616c6c732061726520616c6c6f776564210000604482015290519081900360640190fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff9091169136908237600080368334866127105a03f13d6000833e808015610113573d83f35b3d83fd5b34801561012357600080fd5b506101576004803603602081101561013a57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610250565b005b34801561016557600080fd5b506101576004803603602081101561017c57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166102b7565b3480156101a557600080fd5b506101d9600480360360208110156101bc57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166105f1565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101d96004803603604081101561021857600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610619565b34801561024757600080fd5b506101d9610751565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f476f6f6420747279000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600160205260408120549054908216911681141561035357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920616374697665210000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff81166103d557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f742072656769737465726564210000000000000000000000000000000000604482015290519081900360640190fd5b8173ffffffffffffffffffffffffffffffffffffffff16635051a5ec6040518163ffffffff1660e01b815260040160206040518083038186803b15801561041b57600080fd5b505afa15801561042f573d6000803e3d6000fd5b505050506040513d602081101561044557600080fd5b50516104b257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f742061636365707465642100000000000000000000000000000000000000604482015290519081900360640190fd5b60008054604080517fce5494bb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92831660048201819052915191939285169263ce5494bb9260248084019382900301818387803b15801561052557600080fd5b505af1158015610539573d6000803e3d6000fd5b5050600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8681169182178355604080517ef55d9d000000000000000000000000000000000000000000000000000000008152600481019390935251908616945062f55d9d93506024808301939282900301818387803b1580156105d457600080fd5b505af11580156105e8573d6000803e3d6000fd5b50505050505050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6000805473ffffffffffffffffffffffffffffffffffffffff848116911614156106a457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920616374697665210000000000000000000000000000000000604482015290519081900360640190fd5b604080517f1684f69f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101849052905161030591631684f69f91349160448082019260209290919082900301818588803b15801561071d57600080fd5b505af1158015610731573d6000803e3d6000fd5b50505050506040513d602081101561074857600080fd5b50519392505050565b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a723058203a5f43e0f8aabd2a63006406a3b7863bb4c97b6465317c0ad98069f53c6d1bae64736f6c63430005090032`

// DeployGovernedProxy deploys a new Ethereum contract, binding an instance of GovernedProxy to it.
func DeployGovernedProxy(auth *bind.TransactOpts, backend bind.ContractBackend, impl common.Address) (common.Address, *types.Transaction, *GovernedProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernedProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernedProxyBin), backend, impl)
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

// CurrentImpl is a free data retrieval call binding the contract method 0xad797cdd.
//
// Solidity: function current_impl() constant returns(address)
func (_GovernedProxy *GovernedProxyCaller) CurrentImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "current_impl")
	return *ret0, err
}

// CurrentImpl is a free data retrieval call binding the contract method 0xad797cdd.
//
// Solidity: function current_impl() constant returns(address)
func (_GovernedProxy *GovernedProxySession) CurrentImpl() (common.Address, error) {
	return _GovernedProxy.Contract.CurrentImpl(&_GovernedProxy.CallOpts)
}

// CurrentImpl is a free data retrieval call binding the contract method 0xad797cdd.
//
// Solidity: function current_impl() constant returns(address)
func (_GovernedProxy *GovernedProxyCallerSession) CurrentImpl() (common.Address, error) {
	return _GovernedProxy.Contract.CurrentImpl(&_GovernedProxy.CallOpts)
}

// UpgradeProposals is a free data retrieval call binding the contract method 0x32e3a905.
//
// Solidity: function upgrade_proposals(address ) constant returns(address)
func (_GovernedProxy *GovernedProxyCaller) UpgradeProposals(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "upgrade_proposals", arg0)
	return *ret0, err
}

// UpgradeProposals is a free data retrieval call binding the contract method 0x32e3a905.
//
// Solidity: function upgrade_proposals(address ) constant returns(address)
func (_GovernedProxy *GovernedProxySession) UpgradeProposals(arg0 common.Address) (common.Address, error) {
	return _GovernedProxy.Contract.UpgradeProposals(&_GovernedProxy.CallOpts, arg0)
}

// UpgradeProposals is a free data retrieval call binding the contract method 0x32e3a905.
//
// Solidity: function upgrade_proposals(address ) constant returns(address)
func (_GovernedProxy *GovernedProxyCallerSession) UpgradeProposals(arg0 common.Address) (common.Address, error) {
	return _GovernedProxy.Contract.UpgradeProposals(&_GovernedProxy.CallOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_GovernedProxy *GovernedProxyTransactor) Destroy(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "destroy", arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_GovernedProxy *GovernedProxySession) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Destroy(&_GovernedProxy.TransactOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Destroy(&_GovernedProxy.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_GovernedProxy *GovernedProxyTransactor) Migrate(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "migrate", arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_GovernedProxy *GovernedProxySession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Migrate(&_GovernedProxy.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Migrate(&_GovernedProxy.TransactOpts, arg0)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address new_impl, uint256 period) returns(address proposal)
func (_GovernedProxy *GovernedProxyTransactor) ProposeUpgrade(opts *bind.TransactOpts, new_impl common.Address, period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "proposeUpgrade", new_impl, period)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address new_impl, uint256 period) returns(address proposal)
func (_GovernedProxy *GovernedProxySession) ProposeUpgrade(new_impl common.Address, period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.Contract.ProposeUpgrade(&_GovernedProxy.TransactOpts, new_impl, period)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address new_impl, uint256 period) returns(address proposal)
func (_GovernedProxy *GovernedProxyTransactorSession) ProposeUpgrade(new_impl common.Address, period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.Contract.ProposeUpgrade(&_GovernedProxy.TransactOpts, new_impl, period)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address proposal) returns()
func (_GovernedProxy *GovernedProxyTransactor) Upgrade(opts *bind.TransactOpts, proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "upgrade", proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address proposal) returns()
func (_GovernedProxy *GovernedProxySession) Upgrade(proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Upgrade(&_GovernedProxy.TransactOpts, proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address proposal) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) Upgrade(proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Upgrade(&_GovernedProxy.TransactOpts, proposal)
}
