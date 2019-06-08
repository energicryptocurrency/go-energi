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

// MasternodeRegistryV1ABI is the input ABI used to generate the binding from.
const MasternodeRegistryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"name\":\"active\",\"type\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint256\"},{\"name\":\"max_of_all_times\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"validate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"invalidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"masternode\",\"type\":\"address\"},{\"name\":\"ipv4address\",\"type\":\"uint32\"}],\"name\":\"announce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"heartbeat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"denounce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"isValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// MasternodeRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const MasternodeRegistryV1Bin = `608060405234801561001057600080fd5b5060405161059d38038061059d8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610538806100656000396000f3fe6080604052600436106100c65760003560e01c80636e029ad111610074578063a9fb763c1161004e578063a9fb763c146102fd578063ce5494bb14610313578063ec55688914610353576100c6565b80636e029ad11461026a5780637cc27d45146102945780638b1b925f146102a9576100c6565b8063207c64fb116100a5578063207c64fb146101de57806337a3931f146101de5780633e3e4ac31461021e576100c6565b8062f55d9d1461012d57806306661abd1461016f5780631c4b774b146101a2575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561013957600080fd5b5061016d6004803603602081101561015057600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610391565b005b34801561017b57600080fd5b50610184610430565b60408051938452602084019290925282820152519081900360600190f35b3480156101ae57600080fd5b506101cc600480360360208110156101c557600080fd5b503561043a565b60408051918252519081900360200190f35b3480156101ea57600080fd5b5061016d6004803603602081101561020157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610452565b34801561022a57600080fd5b5061016d6004803603604081101561024157600080fd5b50803573ffffffffffffffffffffffffffffffffffffffff16906020013563ffffffff16610455565b34801561027657600080fd5b5061016d6004803603602081101561028d57600080fd5b5035610452565b3480156102a057600080fd5b5061016d610459565b3480156102b557600080fd5b506102e9600480360360208110156102cc57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661045b565b604080519115158252519081900360200190f35b61016d6004803603602081101561028d57600080fd5b34801561031f57600080fd5b5061016d6004803603602081101561033657600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610461565b34801561035f57600080fd5b506103686104e7565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461041757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b8073ffffffffffffffffffffffffffffffffffffffff16ff5b6000806000909192565b6000811561044d5750677ed7cd92ff1200005b919050565b50565b5050565b565b50600090565b60005473ffffffffffffffffffffffffffffffffffffffff16331461045257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a723058200aa35e2cde5a5771e137df951cc4406195cf975fb84bdc73dd7fd7bf3eb3339664736f6c63430005090032`

// DeployMasternodeRegistryV1 deploys a new Ethereum contract, binding an instance of MasternodeRegistryV1 to it.
func DeployMasternodeRegistryV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *MasternodeRegistryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeRegistryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MasternodeRegistryV1Bin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MasternodeRegistryV1{MasternodeRegistryV1Caller: MasternodeRegistryV1Caller{contract: contract}, MasternodeRegistryV1Transactor: MasternodeRegistryV1Transactor{contract: contract}, MasternodeRegistryV1Filterer: MasternodeRegistryV1Filterer{contract: contract}}, nil
}

// MasternodeRegistryV1 is an auto generated Go binding around an Ethereum contract.
type MasternodeRegistryV1 struct {
	MasternodeRegistryV1Caller     // Read-only binding to the contract
	MasternodeRegistryV1Transactor // Write-only binding to the contract
	MasternodeRegistryV1Filterer   // Log filterer for contract events
}

// MasternodeRegistryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasternodeRegistryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeRegistryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasternodeRegistryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeRegistryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasternodeRegistryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeRegistryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasternodeRegistryV1Session struct {
	Contract     *MasternodeRegistryV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MasternodeRegistryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasternodeRegistryV1CallerSession struct {
	Contract *MasternodeRegistryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MasternodeRegistryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasternodeRegistryV1TransactorSession struct {
	Contract     *MasternodeRegistryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MasternodeRegistryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasternodeRegistryV1Raw struct {
	Contract *MasternodeRegistryV1 // Generic contract binding to access the raw methods on
}

// MasternodeRegistryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasternodeRegistryV1CallerRaw struct {
	Contract *MasternodeRegistryV1Caller // Generic read-only contract binding to access the raw methods on
}

// MasternodeRegistryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasternodeRegistryV1TransactorRaw struct {
	Contract *MasternodeRegistryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasternodeRegistryV1 creates a new instance of MasternodeRegistryV1, bound to a specific deployed contract.
func NewMasternodeRegistryV1(address common.Address, backend bind.ContractBackend) (*MasternodeRegistryV1, error) {
	contract, err := bindMasternodeRegistryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1{MasternodeRegistryV1Caller: MasternodeRegistryV1Caller{contract: contract}, MasternodeRegistryV1Transactor: MasternodeRegistryV1Transactor{contract: contract}, MasternodeRegistryV1Filterer: MasternodeRegistryV1Filterer{contract: contract}}, nil
}

// NewMasternodeRegistryV1Caller creates a new read-only instance of MasternodeRegistryV1, bound to a specific deployed contract.
func NewMasternodeRegistryV1Caller(address common.Address, caller bind.ContractCaller) (*MasternodeRegistryV1Caller, error) {
	contract, err := bindMasternodeRegistryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1Caller{contract: contract}, nil
}

// NewMasternodeRegistryV1Transactor creates a new write-only instance of MasternodeRegistryV1, bound to a specific deployed contract.
func NewMasternodeRegistryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*MasternodeRegistryV1Transactor, error) {
	contract, err := bindMasternodeRegistryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1Transactor{contract: contract}, nil
}

// NewMasternodeRegistryV1Filterer creates a new log filterer instance of MasternodeRegistryV1, bound to a specific deployed contract.
func NewMasternodeRegistryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*MasternodeRegistryV1Filterer, error) {
	contract, err := bindMasternodeRegistryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1Filterer{contract: contract}, nil
}

// bindMasternodeRegistryV1 binds a generic wrapper to an already deployed contract.
func bindMasternodeRegistryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeRegistryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeRegistryV1 *MasternodeRegistryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeRegistryV1.Contract.MasternodeRegistryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeRegistryV1 *MasternodeRegistryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.MasternodeRegistryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeRegistryV1 *MasternodeRegistryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.MasternodeRegistryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeRegistryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 max_of_all_times)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) Count(opts *bind.CallOpts) (struct {
	Active        *big.Int
	Total         *big.Int
	MaxOfAllTimes *big.Int
}, error) {
	ret := new(struct {
		Active        *big.Int
		Total         *big.Int
		MaxOfAllTimes *big.Int
	})
	out := ret
	err := _MasternodeRegistryV1.contract.Call(opts, out, "count")
	return *ret, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 max_of_all_times)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Count() (struct {
	Active        *big.Int
	Total         *big.Int
	MaxOfAllTimes *big.Int
}, error) {
	return _MasternodeRegistryV1.Contract.Count(&_MasternodeRegistryV1.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 max_of_all_times)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) Count() (struct {
	Active        *big.Int
	Total         *big.Int
	MaxOfAllTimes *big.Int
}, error) {
	return _MasternodeRegistryV1.Contract.Count(&_MasternodeRegistryV1.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) GetReward(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "getReward", _blockNumber)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _MasternodeRegistryV1.Contract.GetReward(&_MasternodeRegistryV1.CallOpts, _blockNumber)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _MasternodeRegistryV1.Contract.GetReward(&_MasternodeRegistryV1.CallOpts, _blockNumber)
}

// IsValid is a free data retrieval call binding the contract method 0x8b1b925f.
//
// Solidity: function isValid(address masternode) constant returns(bool)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) IsValid(opts *bind.CallOpts, masternode common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "isValid", masternode)
	return *ret0, err
}

// IsValid is a free data retrieval call binding the contract method 0x8b1b925f.
//
// Solidity: function isValid(address masternode) constant returns(bool)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) IsValid(masternode common.Address) (bool, error) {
	return _MasternodeRegistryV1.Contract.IsValid(&_MasternodeRegistryV1.CallOpts, masternode)
}

// IsValid is a free data retrieval call binding the contract method 0x8b1b925f.
//
// Solidity: function isValid(address masternode) constant returns(bool)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) IsValid(masternode common.Address) (bool, error) {
	return _MasternodeRegistryV1.Contract.IsValid(&_MasternodeRegistryV1.CallOpts, masternode)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Proxy() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.Proxy(&_MasternodeRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) Proxy() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.Proxy(&_MasternodeRegistryV1.CallOpts)
}

// Announce is a paid mutator transaction binding the contract method 0x3e3e4ac3.
//
// Solidity: function announce(address masternode, uint32 ipv4address) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Announce(opts *bind.TransactOpts, masternode common.Address, ipv4address uint32) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "announce", masternode, ipv4address)
}

// Announce is a paid mutator transaction binding the contract method 0x3e3e4ac3.
//
// Solidity: function announce(address masternode, uint32 ipv4address) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Announce(masternode common.Address, ipv4address uint32) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Announce(&_MasternodeRegistryV1.TransactOpts, masternode, ipv4address)
}

// Announce is a paid mutator transaction binding the contract method 0x3e3e4ac3.
//
// Solidity: function announce(address masternode, uint32 ipv4address) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Announce(masternode common.Address, ipv4address uint32) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Announce(&_MasternodeRegistryV1.TransactOpts, masternode, ipv4address)
}

// Denounce is a paid mutator transaction binding the contract method 0x7cc27d45.
//
// Solidity: function denounce() returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Denounce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "denounce")
}

// Denounce is a paid mutator transaction binding the contract method 0x7cc27d45.
//
// Solidity: function denounce() returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Denounce() (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Denounce(&_MasternodeRegistryV1.TransactOpts)
}

// Denounce is a paid mutator transaction binding the contract method 0x7cc27d45.
//
// Solidity: function denounce() returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Denounce() (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Denounce(&_MasternodeRegistryV1.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Destroy(&_MasternodeRegistryV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Destroy(&_MasternodeRegistryV1.TransactOpts, _newImpl)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x6e029ad1.
//
// Solidity: function heartbeat(uint256 sw_features) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Heartbeat(opts *bind.TransactOpts, sw_features *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "heartbeat", sw_features)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x6e029ad1.
//
// Solidity: function heartbeat(uint256 sw_features) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Heartbeat(sw_features *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Heartbeat(&_MasternodeRegistryV1.TransactOpts, sw_features)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x6e029ad1.
//
// Solidity: function heartbeat(uint256 sw_features) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Heartbeat(sw_features *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Heartbeat(&_MasternodeRegistryV1.TransactOpts, sw_features)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Invalidate(opts *bind.TransactOpts, masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "invalidate", masternode)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Invalidate(masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Invalidate(&_MasternodeRegistryV1.TransactOpts, masternode)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Invalidate(masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Invalidate(&_MasternodeRegistryV1.TransactOpts, masternode)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Migrate(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "migrate", arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Migrate(&_MasternodeRegistryV1.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Migrate(&_MasternodeRegistryV1.TransactOpts, arg0)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 ) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Reward(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "reward", arg0)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 ) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Reward(arg0 *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Reward(&_MasternodeRegistryV1.TransactOpts, arg0)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 ) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Reward(arg0 *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Reward(&_MasternodeRegistryV1.TransactOpts, arg0)
}

// Validate is a paid mutator transaction binding the contract method 0x207c64fb.
//
// Solidity: function validate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Validate(opts *bind.TransactOpts, masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "validate", masternode)
}

// Validate is a paid mutator transaction binding the contract method 0x207c64fb.
//
// Solidity: function validate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Validate(masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Validate(&_MasternodeRegistryV1.TransactOpts, masternode)
}

// Validate is a paid mutator transaction binding the contract method 0x207c64fb.
//
// Solidity: function validate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Validate(masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Validate(&_MasternodeRegistryV1.TransactOpts, masternode)
}
