// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	ethereum "github.com/energicryptocurrency/energi"
	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/event"
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

// IHardforkRegistryABI is the input ABI used to generate the binding from.
const IHardforkRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block_number\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"HardforkCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block_number\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"HardforkFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unindexed_name\",\"type\":\"bytes32\"}],\"name\":\"HardforkRemoved\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"block_number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerate\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerateActive\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumeratePending\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"block_number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IHardforkRegistry is an auto generated Go binding around an Ethereum contract.
type IHardforkRegistry struct {
	IHardforkRegistryCaller     // Read-only binding to the contract
	IHardforkRegistryTransactor // Write-only binding to the contract
	IHardforkRegistryFilterer   // Log filterer for contract events
}

// IHardforkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHardforkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHardforkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHardforkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHardforkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHardforkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHardforkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHardforkRegistrySession struct {
	Contract     *IHardforkRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IHardforkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHardforkRegistryCallerSession struct {
	Contract *IHardforkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IHardforkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHardforkRegistryTransactorSession struct {
	Contract     *IHardforkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IHardforkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHardforkRegistryRaw struct {
	Contract *IHardforkRegistry // Generic contract binding to access the raw methods on
}

// IHardforkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHardforkRegistryCallerRaw struct {
	Contract *IHardforkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IHardforkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHardforkRegistryTransactorRaw struct {
	Contract *IHardforkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHardforkRegistry creates a new instance of IHardforkRegistry, bound to a specific deployed contract.
func NewIHardforkRegistry(address common.Address, backend bind.ContractBackend) (*IHardforkRegistry, error) {
	contract, err := bindIHardforkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHardforkRegistry{IHardforkRegistryCaller: IHardforkRegistryCaller{contract: contract}, IHardforkRegistryTransactor: IHardforkRegistryTransactor{contract: contract}, IHardforkRegistryFilterer: IHardforkRegistryFilterer{contract: contract}}, nil
}

// NewIHardforkRegistryCaller creates a new read-only instance of IHardforkRegistry, bound to a specific deployed contract.
func NewIHardforkRegistryCaller(address common.Address, caller bind.ContractCaller) (*IHardforkRegistryCaller, error) {
	contract, err := bindIHardforkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHardforkRegistryCaller{contract: contract}, nil
}

// NewIHardforkRegistryTransactor creates a new write-only instance of IHardforkRegistry, bound to a specific deployed contract.
func NewIHardforkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IHardforkRegistryTransactor, error) {
	contract, err := bindIHardforkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHardforkRegistryTransactor{contract: contract}, nil
}

// NewIHardforkRegistryFilterer creates a new log filterer instance of IHardforkRegistry, bound to a specific deployed contract.
func NewIHardforkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IHardforkRegistryFilterer, error) {
	contract, err := bindIHardforkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHardforkRegistryFilterer{contract: contract}, nil
}

// bindIHardforkRegistry binds a generic wrapper to an already deployed contract.
func bindIHardforkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IHardforkRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHardforkRegistry *IHardforkRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IHardforkRegistry.Contract.IHardforkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHardforkRegistry *IHardforkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.IHardforkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHardforkRegistry *IHardforkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.IHardforkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHardforkRegistry *IHardforkRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IHardforkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHardforkRegistry *IHardforkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHardforkRegistry *IHardforkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.contract.Transact(opts, method, params...)
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistryCaller) Enumerate(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IHardforkRegistry.contract.Call(opts, out, "enumerate")
	return *ret0, err
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistrySession) Enumerate() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.Enumerate(&_IHardforkRegistry.CallOpts)
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistryCallerSession) Enumerate() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.Enumerate(&_IHardforkRegistry.CallOpts)
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistryCaller) EnumerateActive(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IHardforkRegistry.contract.Call(opts, out, "enumerateActive")
	return *ret0, err
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistrySession) EnumerateActive() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumerateActive(&_IHardforkRegistry.CallOpts)
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistryCallerSession) EnumerateActive() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumerateActive(&_IHardforkRegistry.CallOpts)
}

// EnumeratePending is a free data retrieval call binding the contract method 0xca89ad5f.
//
// Solidity: function enumeratePending() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistryCaller) EnumeratePending(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IHardforkRegistry.contract.Call(opts, out, "enumeratePending")
	return *ret0, err
}

// EnumeratePending is a free data retrieval call binding the contract method 0xca89ad5f.
//
// Solidity: function enumeratePending() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistrySession) EnumeratePending() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumeratePending(&_IHardforkRegistry.CallOpts)
}

// EnumeratePending is a free data retrieval call binding the contract method 0xca89ad5f.
//
// Solidity: function enumeratePending() constant returns(bytes32[])
func (_IHardforkRegistry *IHardforkRegistryCallerSession) EnumeratePending() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumeratePending(&_IHardforkRegistry.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 name) constant returns(uint256 block_number, bytes32 block_hash, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryCaller) Get(opts *bind.CallOpts, name [32]byte) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	SwFeatures  *big.Int
}, error) {
	ret := new(struct {
		BlockNumber *big.Int
		BlockHash   [32]byte
		SwFeatures  *big.Int
	})
	out := ret
	err := _IHardforkRegistry.contract.Call(opts, out, "get", name)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 name) constant returns(uint256 block_number, bytes32 block_hash, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistrySession) Get(name [32]byte) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	SwFeatures  *big.Int
}, error) {
	return _IHardforkRegistry.Contract.Get(&_IHardforkRegistry.CallOpts, name)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 name) constant returns(uint256 block_number, bytes32 block_hash, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryCallerSession) Get(name [32]byte) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	SwFeatures  *big.Int
}, error) {
	return _IHardforkRegistry.Contract.Get(&_IHardforkRegistry.CallOpts, name)
}

// IsActive is a free data retrieval call binding the contract method 0x5c36901c.
//
// Solidity: function isActive(bytes32 name) constant returns(bool)
func (_IHardforkRegistry *IHardforkRegistryCaller) IsActive(opts *bind.CallOpts, name [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IHardforkRegistry.contract.Call(opts, out, "isActive", name)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x5c36901c.
//
// Solidity: function isActive(bytes32 name) constant returns(bool)
func (_IHardforkRegistry *IHardforkRegistrySession) IsActive(name [32]byte) (bool, error) {
	return _IHardforkRegistry.Contract.IsActive(&_IHardforkRegistry.CallOpts, name)
}

// IsActive is a free data retrieval call binding the contract method 0x5c36901c.
//
// Solidity: function isActive(bytes32 name) constant returns(bool)
func (_IHardforkRegistry *IHardforkRegistryCallerSession) IsActive(name [32]byte) (bool, error) {
	return _IHardforkRegistry.Contract.IsActive(&_IHardforkRegistry.CallOpts, name)
}

// Add is a paid mutator transaction binding the contract method 0xaa61604f.
//
// Solidity: function add(bytes32 name, uint256 block_number, uint256 sw_features) returns()
func (_IHardforkRegistry *IHardforkRegistryTransactor) Add(opts *bind.TransactOpts, name [32]byte, block_number *big.Int, sw_features *big.Int) (*types.Transaction, error) {
	return _IHardforkRegistry.contract.Transact(opts, "add", name, block_number, sw_features)
}

// Add is a paid mutator transaction binding the contract method 0xaa61604f.
//
// Solidity: function add(bytes32 name, uint256 block_number, uint256 sw_features) returns()
func (_IHardforkRegistry *IHardforkRegistrySession) Add(name [32]byte, block_number *big.Int, sw_features *big.Int) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Add(&_IHardforkRegistry.TransactOpts, name, block_number, sw_features)
}

// Add is a paid mutator transaction binding the contract method 0xaa61604f.
//
// Solidity: function add(bytes32 name, uint256 block_number, uint256 sw_features) returns()
func (_IHardforkRegistry *IHardforkRegistryTransactorSession) Add(name [32]byte, block_number *big.Int, sw_features *big.Int) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Add(&_IHardforkRegistry.TransactOpts, name, block_number, sw_features)
}

// Finalize is a paid mutator transaction binding the contract method 0x92584d80.
//
// Solidity: function finalize(bytes32 name) returns()
func (_IHardforkRegistry *IHardforkRegistryTransactor) Finalize(opts *bind.TransactOpts, name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.contract.Transact(opts, "finalize", name)
}

// Finalize is a paid mutator transaction binding the contract method 0x92584d80.
//
// Solidity: function finalize(bytes32 name) returns()
func (_IHardforkRegistry *IHardforkRegistrySession) Finalize(name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Finalize(&_IHardforkRegistry.TransactOpts, name)
}

// Finalize is a paid mutator transaction binding the contract method 0x92584d80.
//
// Solidity: function finalize(bytes32 name) returns()
func (_IHardforkRegistry *IHardforkRegistryTransactorSession) Finalize(name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Finalize(&_IHardforkRegistry.TransactOpts, name)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 name) returns(bool)
func (_IHardforkRegistry *IHardforkRegistryTransactor) Remove(opts *bind.TransactOpts, name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.contract.Transact(opts, "remove", name)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 name) returns(bool)
func (_IHardforkRegistry *IHardforkRegistrySession) Remove(name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Remove(&_IHardforkRegistry.TransactOpts, name)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 name) returns(bool)
func (_IHardforkRegistry *IHardforkRegistryTransactorSession) Remove(name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Remove(&_IHardforkRegistry.TransactOpts, name)
}

// IHardforkRegistryHardforkCreatedIterator is returned from FilterHardforkCreated and is used to iterate over the raw logs and unpacked data for HardforkCreated events raised by the IHardforkRegistry contract.
type IHardforkRegistryHardforkCreatedIterator struct {
	Event *IHardforkRegistryHardforkCreated // Event containing the contract specifics and raw log

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
func (it *IHardforkRegistryHardforkCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHardforkRegistryHardforkCreated)
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
		it.Event = new(IHardforkRegistryHardforkCreated)
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
func (it *IHardforkRegistryHardforkCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHardforkRegistryHardforkCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHardforkRegistryHardforkCreated represents a HardforkCreated event raised by the IHardforkRegistry contract.
type IHardforkRegistryHardforkCreated struct {
	Name        [32]byte
	BlockNumber *big.Int
	SwFeatures  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHardforkCreated is a free log retrieval operation binding the contract event 0x33bbb09eb0e71b49dacc2c0e0f73dd640a1314f0d08d5f8efa5c12eac770c4c3.
//
// Solidity: event HardforkCreated(bytes32 indexed name, uint256 block_number, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryFilterer) FilterHardforkCreated(opts *bind.FilterOpts, name [][32]byte) (*IHardforkRegistryHardforkCreatedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _IHardforkRegistry.contract.FilterLogs(opts, "HardforkCreated", nameRule)
	if err != nil {
		return nil, err
	}
	return &IHardforkRegistryHardforkCreatedIterator{contract: _IHardforkRegistry.contract, event: "HardforkCreated", logs: logs, sub: sub}, nil
}

// WatchHardforkCreated is a free log subscription operation binding the contract event 0x33bbb09eb0e71b49dacc2c0e0f73dd640a1314f0d08d5f8efa5c12eac770c4c3.
//
// Solidity: event HardforkCreated(bytes32 indexed name, uint256 block_number, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryFilterer) WatchHardforkCreated(opts *bind.WatchOpts, sink chan<- *IHardforkRegistryHardforkCreated, name [][32]byte) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _IHardforkRegistry.contract.WatchLogs(opts, "HardforkCreated", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHardforkRegistryHardforkCreated)
				if err := _IHardforkRegistry.contract.UnpackLog(event, "HardforkCreated", log); err != nil {
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

// IHardforkRegistryHardforkFinalizedIterator is returned from FilterHardforkFinalized and is used to iterate over the raw logs and unpacked data for HardforkFinalized events raised by the IHardforkRegistry contract.
type IHardforkRegistryHardforkFinalizedIterator struct {
	Event *IHardforkRegistryHardforkFinalized // Event containing the contract specifics and raw log

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
func (it *IHardforkRegistryHardforkFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHardforkRegistryHardforkFinalized)
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
		it.Event = new(IHardforkRegistryHardforkFinalized)
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
func (it *IHardforkRegistryHardforkFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHardforkRegistryHardforkFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHardforkRegistryHardforkFinalized represents a HardforkFinalized event raised by the IHardforkRegistry contract.
type IHardforkRegistryHardforkFinalized struct {
	Name        [32]byte
	BlockNumber *big.Int
	BlockHash   [32]byte
	SwFeatures  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHardforkFinalized is a free log retrieval operation binding the contract event 0x6dc459fd769bc8043e2a9bf76cf8ca708f41158bb7d40566a9f488a8fc6c87da.
//
// Solidity: event HardforkFinalized(bytes32 indexed name, uint256 block_number, bytes32 block_hash, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryFilterer) FilterHardforkFinalized(opts *bind.FilterOpts, name [][32]byte) (*IHardforkRegistryHardforkFinalizedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _IHardforkRegistry.contract.FilterLogs(opts, "HardforkFinalized", nameRule)
	if err != nil {
		return nil, err
	}
	return &IHardforkRegistryHardforkFinalizedIterator{contract: _IHardforkRegistry.contract, event: "HardforkFinalized", logs: logs, sub: sub}, nil
}

// WatchHardforkFinalized is a free log subscription operation binding the contract event 0x6dc459fd769bc8043e2a9bf76cf8ca708f41158bb7d40566a9f488a8fc6c87da.
//
// Solidity: event HardforkFinalized(bytes32 indexed name, uint256 block_number, bytes32 block_hash, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryFilterer) WatchHardforkFinalized(opts *bind.WatchOpts, sink chan<- *IHardforkRegistryHardforkFinalized, name [][32]byte) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _IHardforkRegistry.contract.WatchLogs(opts, "HardforkFinalized", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHardforkRegistryHardforkFinalized)
				if err := _IHardforkRegistry.contract.UnpackLog(event, "HardforkFinalized", log); err != nil {
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

// IHardforkRegistryHardforkRemovedIterator is returned from FilterHardforkRemoved and is used to iterate over the raw logs and unpacked data for HardforkRemoved events raised by the IHardforkRegistry contract.
type IHardforkRegistryHardforkRemovedIterator struct {
	Event *IHardforkRegistryHardforkRemoved // Event containing the contract specifics and raw log

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
func (it *IHardforkRegistryHardforkRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHardforkRegistryHardforkRemoved)
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
		it.Event = new(IHardforkRegistryHardforkRemoved)
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
func (it *IHardforkRegistryHardforkRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHardforkRegistryHardforkRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHardforkRegistryHardforkRemoved represents a HardforkRemoved event raised by the IHardforkRegistry contract.
type IHardforkRegistryHardforkRemoved struct {
	Name          [32]byte
	UnindexedName [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterHardforkRemoved is a free log retrieval operation binding the contract event 0x9e405a3cec329c3a3b1fce0ce501b60d6a7714e185d57047ab1d1970ba9f8fea.
//
// Solidity: event HardforkRemoved(bytes32 indexed name, bytes32 unindexed_name)
func (_IHardforkRegistry *IHardforkRegistryFilterer) FilterHardforkRemoved(opts *bind.FilterOpts, name [][32]byte) (*IHardforkRegistryHardforkRemovedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _IHardforkRegistry.contract.FilterLogs(opts, "HardforkRemoved", nameRule)
	if err != nil {
		return nil, err
	}
	return &IHardforkRegistryHardforkRemovedIterator{contract: _IHardforkRegistry.contract, event: "HardforkRemoved", logs: logs, sub: sub}, nil
}

// WatchHardforkRemoved is a free log subscription operation binding the contract event 0x9e405a3cec329c3a3b1fce0ce501b60d6a7714e185d57047ab1d1970ba9f8fea.
//
// Solidity: event HardforkRemoved(bytes32 indexed name, bytes32 unindexed_name)
func (_IHardforkRegistry *IHardforkRegistryFilterer) WatchHardforkRemoved(opts *bind.WatchOpts, sink chan<- *IHardforkRegistryHardforkRemoved, name [][32]byte) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _IHardforkRegistry.contract.WatchLogs(opts, "HardforkRemoved", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHardforkRegistryHardforkRemoved)
				if err := _IHardforkRegistry.contract.UnpackLog(event, "HardforkRemoved", log); err != nil {
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
