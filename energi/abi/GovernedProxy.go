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
const GovernedProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgrade_proposals\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"},{\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"proposeUpgrade\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"impl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spork_proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_impl\",\"type\":\"address\"},{\"name\":\"_sporkProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newImpl\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"UpgradeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newImpl\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]"

// GovernedProxyBin is the compiled bytecode used for deploying new contracts.
const GovernedProxyBin = `608060405234801561001057600080fd5b50604051610ba3380380610ba38339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b03199182161790915560018054939092169216919091179055610b298061007a6000396000f3fe60806040526004361061007a5760003560e01c80638abf60771161004e5780638abf607714610251578063ce5494bb1461012d578063dd6a851d14610266578063ec5568891461027b5761007a565b8062f55d9d1461012d5780630900f0101461016f57806332e3a905146101af5780635b6dee4c14610218575b3233146100e857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f6e6c79206469726563742063616c6c732061726520616c6c6f776564210000604482015290519081900360640190fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff9091169136908237600080368334866127105a03f13d6000833e808015610129573d83f35b3d83fd5b34801561013957600080fd5b5061016d6004803603602081101561015057600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610290565b005b34801561017b57600080fd5b5061016d6004803603602081101561019257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166102f7565b3480156101bb57600080fd5b506101ef600480360360208110156101d257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166106b5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61016d6004803603604081101561022e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356106dd565b34801561025d57600080fd5b506101ef610ab8565b34801561027257600080fd5b506101ef610ad4565b34801561028757600080fd5b506101ef610af0565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f476f6f6420747279000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600260205260408120549054908216911681141561039357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920616374697665210000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff811661041557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f742072656769737465726564210000000000000000000000000000000000604482015290519081900360640190fd5b8173ffffffffffffffffffffffffffffffffffffffff16635051a5ec6040518163ffffffff1660e01b815260040160206040518083038186803b15801561045b57600080fd5b505afa15801561046f573d6000803e3d6000fd5b505050506040513d602081101561048557600080fd5b50516104f257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f742061636365707465642100000000000000000000000000000000000000604482015290519081900360640190fd5b60008054604080517fce5494bb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92831660048201819052915191939285169263ce5494bb9260248084019382900301818387803b15801561056557600080fd5b505af1158015610579573d6000803e3d6000fd5b5050600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8681169182178355604080517ef55d9d000000000000000000000000000000000000000000000000000000008152600481019390935251908616945062f55d9d93506024808301939282900301818387803b15801561061457600080fd5b505af1158015610628573d6000803e3d6000fd5b50505073ffffffffffffffffffffffffffffffffffffffff80851660008181526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690558151928352905192861693507f5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c792918290030190a2505050565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b32331461074b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f6e6c79206469726563742063616c6c732061726520616c6c6f776564210000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff838116911614156107d557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920616374697665210000000000000000000000000000000000604482015290519081900360640190fd5b3073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663ec5568896040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561083457600080fd5b505af1158015610848573d6000803e3d6000fd5b505050506040513d602081101561085e57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff16146108e257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f57726f6e672070726f7879210000000000000000000000000000000000000000604482015290519081900360640190fd5b600154604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691638abf6077916004808301926020929190829003018186803b15801561094d57600080fd5b505afa158015610961573d6000803e3d6000fd5b505050506040513d602081101561097757600080fd5b5051604080517f1684f69f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201869052915192935060009291841691631684f69f913491604480830192602092919082900301818588803b1580156109f857600080fd5b505af1158015610a0c573d6000803e3d6000fd5b50505050506040513d6020811015610a2357600080fd5b505173ffffffffffffffffffffffffffffffffffffffff80821660008181526002602090815260409182902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016948a1694851790558151928352905193945091927f812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763929181900390910190a250505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b309056fea265627a7a72305820d632d0bb3f732d6a0499716f83b814834ed71a8d1ab0bffc57a66f2e8377a50f64736f6c63430005090032`

// DeployGovernedProxy deploys a new Ethereum contract, binding an instance of GovernedProxy to it.
func DeployGovernedProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _impl common.Address, _sporkProxy common.Address) (common.Address, *types.Transaction, *GovernedProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernedProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernedProxyBin), backend, _impl, _sporkProxy)
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

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_GovernedProxy *GovernedProxyCaller) Impl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "impl")
	return *ret0, err
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_GovernedProxy *GovernedProxySession) Impl() (common.Address, error) {
	return _GovernedProxy.Contract.Impl(&_GovernedProxy.CallOpts)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_GovernedProxy *GovernedProxyCallerSession) Impl() (common.Address, error) {
	return _GovernedProxy.Contract.Impl(&_GovernedProxy.CallOpts)
}

// SporkProxy is a free data retrieval call binding the contract method 0xdd6a851d.
//
// Solidity: function spork_proxy() constant returns(address)
func (_GovernedProxy *GovernedProxyCaller) SporkProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "spork_proxy")
	return *ret0, err
}

// SporkProxy is a free data retrieval call binding the contract method 0xdd6a851d.
//
// Solidity: function spork_proxy() constant returns(address)
func (_GovernedProxy *GovernedProxySession) SporkProxy() (common.Address, error) {
	return _GovernedProxy.Contract.SporkProxy(&_GovernedProxy.CallOpts)
}

// SporkProxy is a free data retrieval call binding the contract method 0xdd6a851d.
//
// Solidity: function spork_proxy() constant returns(address)
func (_GovernedProxy *GovernedProxyCallerSession) SporkProxy() (common.Address, error) {
	return _GovernedProxy.Contract.SporkProxy(&_GovernedProxy.CallOpts)
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
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns()
func (_GovernedProxy *GovernedProxyTransactor) ProposeUpgrade(opts *bind.TransactOpts, _newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "proposeUpgrade", _newImpl, _period)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns()
func (_GovernedProxy *GovernedProxySession) ProposeUpgrade(_newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.Contract.ProposeUpgrade(&_GovernedProxy.TransactOpts, _newImpl, _period)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) ProposeUpgrade(_newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.Contract.ProposeUpgrade(&_GovernedProxy.TransactOpts, _newImpl, _period)
}

// Proxy is a paid mutator transaction binding the contract method 0xec556889.
//
// Solidity: function proxy() returns(address)
func (_GovernedProxy *GovernedProxyTransactor) Proxy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "proxy")
}

// Proxy is a paid mutator transaction binding the contract method 0xec556889.
//
// Solidity: function proxy() returns(address)
func (_GovernedProxy *GovernedProxySession) Proxy() (*types.Transaction, error) {
	return _GovernedProxy.Contract.Proxy(&_GovernedProxy.TransactOpts)
}

// Proxy is a paid mutator transaction binding the contract method 0xec556889.
//
// Solidity: function proxy() returns(address)
func (_GovernedProxy *GovernedProxyTransactorSession) Proxy() (*types.Transaction, error) {
	return _GovernedProxy.Contract.Proxy(&_GovernedProxy.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_GovernedProxy *GovernedProxyTransactor) Upgrade(opts *bind.TransactOpts, _proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "upgrade", _proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_GovernedProxy *GovernedProxySession) Upgrade(_proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Upgrade(&_GovernedProxy.TransactOpts, _proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) Upgrade(_proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Upgrade(&_GovernedProxy.TransactOpts, _proposal)
}

// GovernedProxyUpgradeProposalIterator is returned from FilterUpgradeProposal and is used to iterate over the raw logs and unpacked data for UpgradeProposal events raised by the GovernedProxy contract.
type GovernedProxyUpgradeProposalIterator struct {
	Event *GovernedProxyUpgradeProposal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernedProxyUpgradeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernedProxyUpgradeProposal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernedProxyUpgradeProposal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernedProxyUpgradeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernedProxyUpgradeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernedProxyUpgradeProposal represents a UpgradeProposal event raised by the GovernedProxy contract.
type GovernedProxyUpgradeProposal struct {
	NewImpl  common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradeProposal is a free log retrieval operation binding the contract event 0x812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763.
//
// Solidity: event UpgradeProposal(address indexed newImpl, address proposal)
func (_GovernedProxy *GovernedProxyFilterer) FilterUpgradeProposal(opts *bind.FilterOpts, newImpl []common.Address) (*GovernedProxyUpgradeProposalIterator, error) {

	var newImplRule []interface{}
	for _, newImplItem := range newImpl {
		newImplRule = append(newImplRule, newImplItem)
	}

	logs, sub, err := _GovernedProxy.contract.FilterLogs(opts, "UpgradeProposal", newImplRule)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyUpgradeProposalIterator{contract: _GovernedProxy.contract, event: "UpgradeProposal", logs: logs, sub: sub}, nil
}

// WatchUpgradeProposal is a free log subscription operation binding the contract event 0x812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763.
//
// Solidity: event UpgradeProposal(address indexed newImpl, address proposal)
func (_GovernedProxy *GovernedProxyFilterer) WatchUpgradeProposal(opts *bind.WatchOpts, sink chan<- *GovernedProxyUpgradeProposal, newImpl []common.Address) (event.Subscription, error) {

	var newImplRule []interface{}
	for _, newImplItem := range newImpl {
		newImplRule = append(newImplRule, newImplItem)
	}

	logs, sub, err := _GovernedProxy.contract.WatchLogs(opts, "UpgradeProposal", newImplRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernedProxyUpgradeProposal)
				if err := _GovernedProxy.contract.UnpackLog(event, "UpgradeProposal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GovernedProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GovernedProxy contract.
type GovernedProxyUpgradedIterator struct {
	Event *GovernedProxyUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernedProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernedProxyUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernedProxyUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernedProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernedProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernedProxyUpgraded represents a Upgraded event raised by the GovernedProxy contract.
type GovernedProxyUpgraded struct {
	NewImpl  common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed newImpl, address proposal)
func (_GovernedProxy *GovernedProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, newImpl []common.Address) (*GovernedProxyUpgradedIterator, error) {

	var newImplRule []interface{}
	for _, newImplItem := range newImpl {
		newImplRule = append(newImplRule, newImplItem)
	}

	logs, sub, err := _GovernedProxy.contract.FilterLogs(opts, "Upgraded", newImplRule)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyUpgradedIterator{contract: _GovernedProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed newImpl, address proposal)
func (_GovernedProxy *GovernedProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GovernedProxyUpgraded, newImpl []common.Address) (event.Subscription, error) {

	var newImplRule []interface{}
	for _, newImplItem := range newImpl {
		newImplRule = append(newImplRule, newImplItem)
	}

	logs, sub, err := _GovernedProxy.contract.WatchLogs(opts, "Upgraded", newImplRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernedProxyUpgraded)
				if err := _GovernedProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
