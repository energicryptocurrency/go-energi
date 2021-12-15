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

// ICheckpointRegistryV2ABI is the input ABI used to generate the binding from.
const ICheckpointRegistryV2ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractICheckpoint\",\"name\":\"checkpoint\",\"type\":\"address\"}],\"name\":\"Checkpoint\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CPP_signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"contractICheckpoint[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"contractICheckpoint\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractICheckpoint\",\"name\":\"checkpoint\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"sign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"signatureBase\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"sigbase\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ICheckpointRegistryV2 is an auto generated Go binding around an Ethereum contract.
type ICheckpointRegistryV2 struct {
	ICheckpointRegistryV2Caller     // Read-only binding to the contract
	ICheckpointRegistryV2Transactor // Write-only binding to the contract
	ICheckpointRegistryV2Filterer   // Log filterer for contract events
}

// ICheckpointRegistryV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ICheckpointRegistryV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointRegistryV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ICheckpointRegistryV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointRegistryV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICheckpointRegistryV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointRegistryV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICheckpointRegistryV2Session struct {
	Contract     *ICheckpointRegistryV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ICheckpointRegistryV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICheckpointRegistryV2CallerSession struct {
	Contract *ICheckpointRegistryV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ICheckpointRegistryV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICheckpointRegistryV2TransactorSession struct {
	Contract     *ICheckpointRegistryV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ICheckpointRegistryV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ICheckpointRegistryV2Raw struct {
	Contract *ICheckpointRegistryV2 // Generic contract binding to access the raw methods on
}

// ICheckpointRegistryV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICheckpointRegistryV2CallerRaw struct {
	Contract *ICheckpointRegistryV2Caller // Generic read-only contract binding to access the raw methods on
}

// ICheckpointRegistryV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICheckpointRegistryV2TransactorRaw struct {
	Contract *ICheckpointRegistryV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewICheckpointRegistryV2 creates a new instance of ICheckpointRegistryV2, bound to a specific deployed contract.
func NewICheckpointRegistryV2(address common.Address, backend bind.ContractBackend) (*ICheckpointRegistryV2, error) {
	contract, err := bindICheckpointRegistryV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryV2{ICheckpointRegistryV2Caller: ICheckpointRegistryV2Caller{contract: contract}, ICheckpointRegistryV2Transactor: ICheckpointRegistryV2Transactor{contract: contract}, ICheckpointRegistryV2Filterer: ICheckpointRegistryV2Filterer{contract: contract}}, nil
}

// NewICheckpointRegistryV2Caller creates a new read-only instance of ICheckpointRegistryV2, bound to a specific deployed contract.
func NewICheckpointRegistryV2Caller(address common.Address, caller bind.ContractCaller) (*ICheckpointRegistryV2Caller, error) {
	contract, err := bindICheckpointRegistryV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryV2Caller{contract: contract}, nil
}

// NewICheckpointRegistryV2Transactor creates a new write-only instance of ICheckpointRegistryV2, bound to a specific deployed contract.
func NewICheckpointRegistryV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ICheckpointRegistryV2Transactor, error) {
	contract, err := bindICheckpointRegistryV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryV2Transactor{contract: contract}, nil
}

// NewICheckpointRegistryV2Filterer creates a new log filterer instance of ICheckpointRegistryV2, bound to a specific deployed contract.
func NewICheckpointRegistryV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ICheckpointRegistryV2Filterer, error) {
	contract, err := bindICheckpointRegistryV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryV2Filterer{contract: contract}, nil
}

// bindICheckpointRegistryV2 binds a generic wrapper to an already deployed contract.
func bindICheckpointRegistryV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICheckpointRegistryV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICheckpointRegistryV2.Contract.ICheckpointRegistryV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.ICheckpointRegistryV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.ICheckpointRegistryV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICheckpointRegistryV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.contract.Transact(opts, method, params...)
}

// CPPSigner is a free data retrieval call binding the contract method 0xd59f1758.
//
// Solidity: function CPP_signer() constant returns(address)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Caller) CPPSigner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ICheckpointRegistryV2.contract.Call(opts, out, "CPP_signer")
	return *ret0, err
}

// CPPSigner is a free data retrieval call binding the contract method 0xd59f1758.
//
// Solidity: function CPP_signer() constant returns(address)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Session) CPPSigner() (common.Address, error) {
	return _ICheckpointRegistryV2.Contract.CPPSigner(&_ICheckpointRegistryV2.CallOpts)
}

// CPPSigner is a free data retrieval call binding the contract method 0xd59f1758.
//
// Solidity: function CPP_signer() constant returns(address)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2CallerSession) CPPSigner() (common.Address, error) {
	return _ICheckpointRegistryV2.Contract.CPPSigner(&_ICheckpointRegistryV2.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0x5a48c0b0.
//
// Solidity: function checkpoints() constant returns(address[])
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Caller) Checkpoints(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ICheckpointRegistryV2.contract.Call(opts, out, "checkpoints")
	return *ret0, err
}

// Checkpoints is a free data retrieval call binding the contract method 0x5a48c0b0.
//
// Solidity: function checkpoints() constant returns(address[])
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Session) Checkpoints() ([]common.Address, error) {
	return _ICheckpointRegistryV2.Contract.Checkpoints(&_ICheckpointRegistryV2.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0x5a48c0b0.
//
// Solidity: function checkpoints() constant returns(address[])
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2CallerSession) Checkpoints() ([]common.Address, error) {
	return _ICheckpointRegistryV2.Contract.Checkpoints(&_ICheckpointRegistryV2.CallOpts)
}

// SignatureBase is a free data retrieval call binding the contract method 0x851f2209.
//
// Solidity: function signatureBase(uint256 number, bytes32 hash) constant returns(bytes32 sigbase)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Caller) SignatureBase(opts *bind.CallOpts, number *big.Int, hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ICheckpointRegistryV2.contract.Call(opts, out, "signatureBase", number, hash)
	return *ret0, err
}

// SignatureBase is a free data retrieval call binding the contract method 0x851f2209.
//
// Solidity: function signatureBase(uint256 number, bytes32 hash) constant returns(bytes32 sigbase)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Session) SignatureBase(number *big.Int, hash [32]byte) ([32]byte, error) {
	return _ICheckpointRegistryV2.Contract.SignatureBase(&_ICheckpointRegistryV2.CallOpts, number, hash)
}

// SignatureBase is a free data retrieval call binding the contract method 0x851f2209.
//
// Solidity: function signatureBase(uint256 number, bytes32 hash) constant returns(bytes32 sigbase)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2CallerSession) SignatureBase(number *big.Int, hash [32]byte) ([32]byte, error) {
	return _ICheckpointRegistryV2.Contract.SignatureBase(&_ICheckpointRegistryV2.CallOpts, number, hash)
}

// Propose is a paid mutator transaction binding the contract method 0xc20fa2ee.
//
// Solidity: function propose(uint256 number, bytes32 hash, bytes signature) returns(address)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Transactor) Propose(opts *bind.TransactOpts, number *big.Int, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.contract.Transact(opts, "propose", number, hash, signature)
}

// Propose is a paid mutator transaction binding the contract method 0xc20fa2ee.
//
// Solidity: function propose(uint256 number, bytes32 hash, bytes signature) returns(address)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Session) Propose(number *big.Int, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.Propose(&_ICheckpointRegistryV2.TransactOpts, number, hash, signature)
}

// Propose is a paid mutator transaction binding the contract method 0xc20fa2ee.
//
// Solidity: function propose(uint256 number, bytes32 hash, bytes signature) returns(address)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2TransactorSession) Propose(number *big.Int, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.Propose(&_ICheckpointRegistryV2.TransactOpts, number, hash, signature)
}

// Remove is a paid mutator transaction binding the contract method 0x28b0558b.
//
// Solidity: function remove(uint256 number, bytes32 hash) returns(bool deleted)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Transactor) Remove(opts *bind.TransactOpts, number *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.contract.Transact(opts, "remove", number, hash)
}

// Remove is a paid mutator transaction binding the contract method 0x28b0558b.
//
// Solidity: function remove(uint256 number, bytes32 hash) returns(bool deleted)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Session) Remove(number *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.Remove(&_ICheckpointRegistryV2.TransactOpts, number, hash)
}

// Remove is a paid mutator transaction binding the contract method 0x28b0558b.
//
// Solidity: function remove(uint256 number, bytes32 hash) returns(bool deleted)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2TransactorSession) Remove(number *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.Remove(&_ICheckpointRegistryV2.TransactOpts, number, hash)
}

// Sign is a paid mutator transaction binding the contract method 0x51fae959.
//
// Solidity: function sign(address checkpoint, bytes signature) returns()
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Transactor) Sign(opts *bind.TransactOpts, checkpoint common.Address, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.contract.Transact(opts, "sign", checkpoint, signature)
}

// Sign is a paid mutator transaction binding the contract method 0x51fae959.
//
// Solidity: function sign(address checkpoint, bytes signature) returns()
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Session) Sign(checkpoint common.Address, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.Sign(&_ICheckpointRegistryV2.TransactOpts, checkpoint, signature)
}

// Sign is a paid mutator transaction binding the contract method 0x51fae959.
//
// Solidity: function sign(address checkpoint, bytes signature) returns()
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2TransactorSession) Sign(checkpoint common.Address, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistryV2.Contract.Sign(&_ICheckpointRegistryV2.TransactOpts, checkpoint, signature)
}

// ICheckpointRegistryV2CheckpointIterator is returned from FilterCheckpoint and is used to iterate over the raw logs and unpacked data for Checkpoint events raised by the ICheckpointRegistryV2 contract.
type ICheckpointRegistryV2CheckpointIterator struct {
	Event *ICheckpointRegistryV2Checkpoint // Event containing the contract specifics and raw log

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
func (it *ICheckpointRegistryV2CheckpointIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICheckpointRegistryV2Checkpoint)
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
		it.Event = new(ICheckpointRegistryV2Checkpoint)
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
func (it *ICheckpointRegistryV2CheckpointIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICheckpointRegistryV2CheckpointIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICheckpointRegistryV2Checkpoint represents a Checkpoint event raised by the ICheckpointRegistryV2 contract.
type ICheckpointRegistryV2Checkpoint struct {
	Number     *big.Int
	Hash       [32]byte
	Checkpoint common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCheckpoint is a free log retrieval operation binding the contract event 0x7f582e7a234c68b135245365b0c69d608f9235392a1ef801b46f04ca5d6cdad4.
//
// Solidity: event Checkpoint(uint256 indexed number, bytes32 hash, address checkpoint)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Filterer) FilterCheckpoint(opts *bind.FilterOpts, number []*big.Int) (*ICheckpointRegistryV2CheckpointIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}

	logs, sub, err := _ICheckpointRegistryV2.contract.FilterLogs(opts, "Checkpoint", numberRule)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryV2CheckpointIterator{contract: _ICheckpointRegistryV2.contract, event: "Checkpoint", logs: logs, sub: sub}, nil
}

// WatchCheckpoint is a free log subscription operation binding the contract event 0x7f582e7a234c68b135245365b0c69d608f9235392a1ef801b46f04ca5d6cdad4.
//
// Solidity: event Checkpoint(uint256 indexed number, bytes32 hash, address checkpoint)
func (_ICheckpointRegistryV2 *ICheckpointRegistryV2Filterer) WatchCheckpoint(opts *bind.WatchOpts, sink chan<- *ICheckpointRegistryV2Checkpoint, number []*big.Int) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}

	logs, sub, err := _ICheckpointRegistryV2.contract.WatchLogs(opts, "Checkpoint", numberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICheckpointRegistryV2Checkpoint)
				if err := _ICheckpointRegistryV2.contract.UnpackLog(event, "Checkpoint", log); err != nil {
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
