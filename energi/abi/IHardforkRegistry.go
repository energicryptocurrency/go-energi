// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	ethereum "energi.world/core/gen3"
	"energi.world/core/gen3/accounts/abi"
	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/event"
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
const IHardforkRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"Hardfork\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerateActive\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"active_hf_names\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerateAll\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"all_hf_names\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumeratePending\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"pending_hf_names\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hardfork_name\",\"type\":\"bytes32\"}],\"name\":\"getHardfork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hardfork_name\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(bytes32[] active_hf_names)
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
// Solidity: function enumerateActive() constant returns(bytes32[] active_hf_names)
func (_IHardforkRegistry *IHardforkRegistrySession) EnumerateActive() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumerateActive(&_IHardforkRegistry.CallOpts)
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(bytes32[] active_hf_names)
func (_IHardforkRegistry *IHardforkRegistryCallerSession) EnumerateActive() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumerateActive(&_IHardforkRegistry.CallOpts)
}

// EnumerateAll is a free data retrieval call binding the contract method 0xbc393afb.
//
// Solidity: function enumerateAll() constant returns(bytes32[] all_hf_names)
func (_IHardforkRegistry *IHardforkRegistryCaller) EnumerateAll(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IHardforkRegistry.contract.Call(opts, out, "enumerateAll")
	return *ret0, err
}

// EnumerateAll is a free data retrieval call binding the contract method 0xbc393afb.
//
// Solidity: function enumerateAll() constant returns(bytes32[] all_hf_names)
func (_IHardforkRegistry *IHardforkRegistrySession) EnumerateAll() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumerateAll(&_IHardforkRegistry.CallOpts)
}

// EnumerateAll is a free data retrieval call binding the contract method 0xbc393afb.
//
// Solidity: function enumerateAll() constant returns(bytes32[] all_hf_names)
func (_IHardforkRegistry *IHardforkRegistryCallerSession) EnumerateAll() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumerateAll(&_IHardforkRegistry.CallOpts)
}

// EnumeratePending is a free data retrieval call binding the contract method 0xca89ad5f.
//
// Solidity: function enumeratePending() constant returns(bytes32[] pending_hf_names)
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
// Solidity: function enumeratePending() constant returns(bytes32[] pending_hf_names)
func (_IHardforkRegistry *IHardforkRegistrySession) EnumeratePending() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumeratePending(&_IHardforkRegistry.CallOpts)
}

// EnumeratePending is a free data retrieval call binding the contract method 0xca89ad5f.
//
// Solidity: function enumeratePending() constant returns(bytes32[] pending_hf_names)
func (_IHardforkRegistry *IHardforkRegistryCallerSession) EnumeratePending() ([][32]byte, error) {
	return _IHardforkRegistry.Contract.EnumeratePending(&_IHardforkRegistry.CallOpts)
}

// GetHardfork is a free data retrieval call binding the contract method 0x44d396ff.
//
// Solidity: function getHardfork(bytes32 _hardfork_name) constant returns(uint256 block_no, bytes32 block_hash, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryCaller) GetHardfork(opts *bind.CallOpts, _hardfork_name [32]byte) (struct {
	BlockNo    *big.Int
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	ret := new(struct {
		BlockNo    *big.Int
		BlockHash  [32]byte
		SwFeatures *big.Int
	})
	out := ret
	err := _IHardforkRegistry.contract.Call(opts, out, "getHardfork", _hardfork_name)
	return *ret, err
}

// GetHardfork is a free data retrieval call binding the contract method 0x44d396ff.
//
// Solidity: function getHardfork(bytes32 _hardfork_name) constant returns(uint256 block_no, bytes32 block_hash, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistrySession) GetHardfork(_hardfork_name [32]byte) (struct {
	BlockNo    *big.Int
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	return _IHardforkRegistry.Contract.GetHardfork(&_IHardforkRegistry.CallOpts, _hardfork_name)
}

// GetHardfork is a free data retrieval call binding the contract method 0x44d396ff.
//
// Solidity: function getHardfork(bytes32 _hardfork_name) constant returns(uint256 block_no, bytes32 block_hash, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryCallerSession) GetHardfork(_hardfork_name [32]byte) (struct {
	BlockNo    *big.Int
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	return _IHardforkRegistry.Contract.GetHardfork(&_IHardforkRegistry.CallOpts, _hardfork_name)
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

// Propose is a paid mutator transaction binding the contract method 0x072a9823.
//
// Solidity: function propose(uint256 block_no, bytes32 name, bytes32 block_hash, uint256 sw_features) returns()
func (_IHardforkRegistry *IHardforkRegistryTransactor) Propose(opts *bind.TransactOpts, block_no *big.Int, name [32]byte, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IHardforkRegistry.contract.Transact(opts, "propose", block_no, name, block_hash, sw_features)
}

// Propose is a paid mutator transaction binding the contract method 0x072a9823.
//
// Solidity: function propose(uint256 block_no, bytes32 name, bytes32 block_hash, uint256 sw_features) returns()
func (_IHardforkRegistry *IHardforkRegistrySession) Propose(block_no *big.Int, name [32]byte, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Propose(&_IHardforkRegistry.TransactOpts, block_no, name, block_hash, sw_features)
}

// Propose is a paid mutator transaction binding the contract method 0x072a9823.
//
// Solidity: function propose(uint256 block_no, bytes32 name, bytes32 block_hash, uint256 sw_features) returns()
func (_IHardforkRegistry *IHardforkRegistryTransactorSession) Propose(block_no *big.Int, name [32]byte, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Propose(&_IHardforkRegistry.TransactOpts, block_no, name, block_hash, sw_features)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 _hardfork_name) returns()
func (_IHardforkRegistry *IHardforkRegistryTransactor) Remove(opts *bind.TransactOpts, _hardfork_name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.contract.Transact(opts, "remove", _hardfork_name)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 _hardfork_name) returns()
func (_IHardforkRegistry *IHardforkRegistrySession) Remove(_hardfork_name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Remove(&_IHardforkRegistry.TransactOpts, _hardfork_name)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 _hardfork_name) returns()
func (_IHardforkRegistry *IHardforkRegistryTransactorSession) Remove(_hardfork_name [32]byte) (*types.Transaction, error) {
	return _IHardforkRegistry.Contract.Remove(&_IHardforkRegistry.TransactOpts, _hardfork_name)
}

// IHardforkRegistryHardforkIterator is returned from FilterHardfork and is used to iterate over the raw logs and unpacked data for Hardfork events raised by the IHardforkRegistry contract.
type IHardforkRegistryHardforkIterator struct {
	Event *IHardforkRegistryHardfork // Event containing the contract specifics and raw log

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
func (it *IHardforkRegistryHardforkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHardforkRegistryHardfork)
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
		it.Event = new(IHardforkRegistryHardfork)
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
func (it *IHardforkRegistryHardforkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHardforkRegistryHardforkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHardforkRegistryHardfork represents a Hardfork event raised by the IHardforkRegistry contract.
type IHardforkRegistryHardfork struct {
	BlockNo    *big.Int
	BlockHash  [32]byte
	Name       [32]byte
	SwFeatures *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHardfork is a free log retrieval operation binding the contract event 0xb22930fde3f4a5bddcb906dc3248ab30f4549101ef99a59fbfbf6121d74c17c5.
//
// Solidity: event Hardfork(uint256 block_no, bytes32 block_hash, bytes32 name, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryFilterer) FilterHardfork(opts *bind.FilterOpts) (*IHardforkRegistryHardforkIterator, error) {

	logs, sub, err := _IHardforkRegistry.contract.FilterLogs(opts, "Hardfork")
	if err != nil {
		return nil, err
	}
	return &IHardforkRegistryHardforkIterator{contract: _IHardforkRegistry.contract, event: "Hardfork", logs: logs, sub: sub}, nil
}

// WatchHardfork is a free log subscription operation binding the contract event 0xb22930fde3f4a5bddcb906dc3248ab30f4549101ef99a59fbfbf6121d74c17c5.
//
// Solidity: event Hardfork(uint256 block_no, bytes32 block_hash, bytes32 name, uint256 sw_features)
func (_IHardforkRegistry *IHardforkRegistryFilterer) WatchHardfork(opts *bind.WatchOpts, sink chan<- *IHardforkRegistryHardfork) (event.Subscription, error) {

	logs, sub, err := _IHardforkRegistry.contract.WatchLogs(opts, "Hardfork")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHardforkRegistryHardfork)
				if err := _IHardforkRegistry.contract.UnpackLog(event, "Hardfork", log); err != nil {
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
