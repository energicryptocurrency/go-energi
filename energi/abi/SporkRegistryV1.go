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
const SporkRegistryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"createUpgradeProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// SporkRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const SporkRegistryV1Bin = `608060405234801561001057600080fd5b506040516106a93803806106a98339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610644806100656000396000f3fe60806040526004361061003e5760003560e01c8062f55d9d146100405780631684f69f14610080578063ce5494bb14610040578063ec556889146100e2575b005b34801561004c57600080fd5b5061003e6004803603602081101561006357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166100f7565b6100b96004803603604081101561009657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610180565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100ee57600080fd5b506100b9610394565b60005473ffffffffffffffffffffffffffffffffffffffff16331461017d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b50565b600069021e19e0c9bab240000034146101fa57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f46656520616d6f756e7400000000000000000000000000000000000000000000604482015290519081900360640190fd5b6212750082101561026c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f506572696f64206d696e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b6301e133808211156102df57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f506572696f64206d617800000000000000000000000000000000000000000000604482015290519081900360640190fd5b600060338332346040516102f2906103b0565b60ff9094168452602084019290925273ffffffffffffffffffffffffffffffffffffffff1660408084019190915260608301919091525190819003608001906000f080158015610346573d6000803e3d6000fd5b5060405190915073ffffffffffffffffffffffffffffffffffffffff8216903480156108fc02916000818181858888f1935050505015801561038c573d6000803e3d6000fd5b509392505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b610252806103be8339019056fe608060405234801561001057600080fd5b506040516102523803806102528339818101604052608081101561003357600080fd5b5080516020820151604083015160609093015160005542016001556002805460ff909216740100000000000000000000000000000000000000000260ff60a01b196001600160a01b039094166001600160a01b031990931692909217929092161790556101ad806100a56000396000f3fe60806040526004361061005a5760003560e01c80635051a5ec116100435780635051a5ec146100ae578063c40a70f8146100d7578063ddca3f43146101155761005a565b80631703a0181461005c57806329dcb0cf14610087575b005b34801561006857600080fd5b5061007161012a565b6040805160ff9092168252519081900360200190f35b34801561009357600080fd5b5061009c61014b565b60408051918252519081900360200190f35b3480156100ba57600080fd5b506100c3610151565b604080519115158252519081900360200190f35b3480156100e357600080fd5b506100ec610156565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012157600080fd5b5061009c610172565b60025474010000000000000000000000000000000000000000900460ff1681565b60015481565b600090565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b6000548156fea265627a7a7230582020015ec3cf5ec19f469ed9ff8056f22c4261c34f050a7e79b54e51cb2c401aa464736f6c63430005090032a265627a7a723058203d7949ccc4a9d4c9754d689a9948d4afccaa925064002ecb5e3be538529908f964736f6c63430005090032`

// DeploySporkRegistryV1 deploys a new Ethereum contract, binding an instance of SporkRegistryV1 to it.
func DeploySporkRegistryV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *SporkRegistryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(SporkRegistryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SporkRegistryV1Bin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SporkRegistryV1{SporkRegistryV1Caller: SporkRegistryV1Caller{contract: contract}, SporkRegistryV1Transactor: SporkRegistryV1Transactor{contract: contract}, SporkRegistryV1Filterer: SporkRegistryV1Filterer{contract: contract}}, nil
}

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

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SporkRegistryV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Session) Proxy() (common.Address, error) {
	return _SporkRegistryV1.Contract.Proxy(&_SporkRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1CallerSession) Proxy() (common.Address, error) {
	return _SporkRegistryV1.Contract.Proxy(&_SporkRegistryV1.CallOpts)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x1684f69f.
//
// Solidity: function createUpgradeProposal(address , uint256 _period) returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Transactor) CreateUpgradeProposal(opts *bind.TransactOpts, arg0 common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SporkRegistryV1.contract.Transact(opts, "createUpgradeProposal", arg0, _period)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x1684f69f.
//
// Solidity: function createUpgradeProposal(address , uint256 _period) returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Session) CreateUpgradeProposal(arg0 common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.CreateUpgradeProposal(&_SporkRegistryV1.TransactOpts, arg0, _period)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x1684f69f.
//
// Solidity: function createUpgradeProposal(address , uint256 _period) returns(address)
func (_SporkRegistryV1 *SporkRegistryV1TransactorSession) CreateUpgradeProposal(arg0 common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.CreateUpgradeProposal(&_SporkRegistryV1.TransactOpts, arg0, _period)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_SporkRegistryV1 *SporkRegistryV1Transactor) Destroy(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.contract.Transact(opts, "destroy", arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_SporkRegistryV1 *SporkRegistryV1Session) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Destroy(&_SporkRegistryV1.TransactOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_SporkRegistryV1 *SporkRegistryV1TransactorSession) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Destroy(&_SporkRegistryV1.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_SporkRegistryV1 *SporkRegistryV1Transactor) Migrate(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.contract.Transact(opts, "migrate", arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_SporkRegistryV1 *SporkRegistryV1Session) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Migrate(&_SporkRegistryV1.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_SporkRegistryV1 *SporkRegistryV1TransactorSession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Migrate(&_SporkRegistryV1.TransactOpts, arg0)
}
