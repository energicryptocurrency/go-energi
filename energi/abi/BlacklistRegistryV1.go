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
const BlacklistRegistryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"v1storage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// BlacklistRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const BlacklistRegistryV1Bin = `608060405234801561001057600080fd5b5060405161068f38038061068f8339818101604052602081101561003357600080fd5b5051600080546001600160a01b0319166001600160a01b03831617905560405161005c9061009f565b604051809103906000f080158015610078573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b0392909216919091179055506100ac565b6102288061046783390190565b6103ac806100bb6000396000f3fe60806040526004361061003e5760003560e01c8062f55d9d146100a55780632d059305146100e7578063ce5494bb14610125578063ec55688914610165575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100b157600080fd5b506100e5600480360360208110156100c857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661017a565b005b3480156100f357600080fd5b506100fc610222565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561013157600080fd5b506100e56004803603602081101561014857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661023e565b34801561017157600080fd5b506100fc6102cc565b60005473ffffffffffffffffffffffffffffffffffffffff16331461020057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610209816102e8565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1633146102c457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6102c9815b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600154604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b15801561035c57600080fd5b505af1158015610370573d6000803e3d6000fd5b505050505056fea265627a7a72305820925523634e30f8bddee3a5ee10af56144f19edb5be9455c17db11be255f70f4764736f6c634300050900326080604052600080546001600160a01b03191633179055610203806100256000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806313af40351461003b57806341c0e1b514610070575b600080fd5b61006e6004803603602081101561005157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610078565b005b61006e610145565b60005473ffffffffffffffffffffffffffffffffffffffff1633146100fe57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101cb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b33fffea265627a7a723058200a51211df62fac93f054c317ec518fa71b4705ecc5233f20f988b4d8a593991464736f6c63430005090032`

// DeployBlacklistRegistryV1 deploys a new Ethereum contract, binding an instance of BlacklistRegistryV1 to it.
func DeployBlacklistRegistryV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *BlacklistRegistryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(BlacklistRegistryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlacklistRegistryV1Bin), backend, _proxy)
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

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlacklistRegistryV1 *BlacklistRegistryV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlacklistRegistryV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlacklistRegistryV1 *BlacklistRegistryV1Session) Proxy() (common.Address, error) {
	return _BlacklistRegistryV1.Contract.Proxy(&_BlacklistRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlacklistRegistryV1 *BlacklistRegistryV1CallerSession) Proxy() (common.Address, error) {
	return _BlacklistRegistryV1.Contract.Proxy(&_BlacklistRegistryV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_BlacklistRegistryV1 *BlacklistRegistryV1Caller) V1storage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlacklistRegistryV1.contract.Call(opts, out, "v1storage")
	return *ret0, err
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_BlacklistRegistryV1 *BlacklistRegistryV1Session) V1storage() (common.Address, error) {
	return _BlacklistRegistryV1.Contract.V1storage(&_BlacklistRegistryV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_BlacklistRegistryV1 *BlacklistRegistryV1CallerSession) V1storage() (common.Address, error) {
	return _BlacklistRegistryV1.Contract.V1storage(&_BlacklistRegistryV1.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.Destroy(&_BlacklistRegistryV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.Destroy(&_BlacklistRegistryV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.Migrate(&_BlacklistRegistryV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlacklistRegistryV1 *BlacklistRegistryV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _BlacklistRegistryV1.Contract.Migrate(&_BlacklistRegistryV1.TransactOpts, _oldImpl)
}
