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

// ICheckpointRegistryABI is the input ABI used to generate the binding from.
const ICheckpointRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractICheckpoint\",\"name\":\"checkpoint\",\"type\":\"address\"}],\"name\":\"Checkpoint\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CPP_signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"contractICheckpoint[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"contractICheckpoint\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractICheckpoint\",\"name\":\"checkpoint\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"sign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"signatureBase\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"sigbase\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ICheckpointRegistry is an auto generated Go binding around an Ethereum contract.
type ICheckpointRegistry struct {
	ICheckpointRegistryCaller     // Read-only binding to the contract
	ICheckpointRegistryTransactor // Write-only binding to the contract
	ICheckpointRegistryFilterer   // Log filterer for contract events
}

// ICheckpointRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICheckpointRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICheckpointRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICheckpointRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICheckpointRegistrySession struct {
	Contract     *ICheckpointRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICheckpointRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICheckpointRegistryCallerSession struct {
	Contract *ICheckpointRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ICheckpointRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICheckpointRegistryTransactorSession struct {
	Contract     *ICheckpointRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ICheckpointRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICheckpointRegistryRaw struct {
	Contract *ICheckpointRegistry // Generic contract binding to access the raw methods on
}

// ICheckpointRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICheckpointRegistryCallerRaw struct {
	Contract *ICheckpointRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ICheckpointRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICheckpointRegistryTransactorRaw struct {
	Contract *ICheckpointRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICheckpointRegistry creates a new instance of ICheckpointRegistry, bound to a specific deployed contract.
func NewICheckpointRegistry(address common.Address, backend bind.ContractBackend) (*ICheckpointRegistry, error) {
	contract, err := bindICheckpointRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistry{ICheckpointRegistryCaller: ICheckpointRegistryCaller{contract: contract}, ICheckpointRegistryTransactor: ICheckpointRegistryTransactor{contract: contract}, ICheckpointRegistryFilterer: ICheckpointRegistryFilterer{contract: contract}}, nil
}

// NewICheckpointRegistryCaller creates a new read-only instance of ICheckpointRegistry, bound to a specific deployed contract.
func NewICheckpointRegistryCaller(address common.Address, caller bind.ContractCaller) (*ICheckpointRegistryCaller, error) {
	contract, err := bindICheckpointRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryCaller{contract: contract}, nil
}

// NewICheckpointRegistryTransactor creates a new write-only instance of ICheckpointRegistry, bound to a specific deployed contract.
func NewICheckpointRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ICheckpointRegistryTransactor, error) {
	contract, err := bindICheckpointRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryTransactor{contract: contract}, nil
}

// NewICheckpointRegistryFilterer creates a new log filterer instance of ICheckpointRegistry, bound to a specific deployed contract.
func NewICheckpointRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ICheckpointRegistryFilterer, error) {
	contract, err := bindICheckpointRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryFilterer{contract: contract}, nil
}

// bindICheckpointRegistry binds a generic wrapper to an already deployed contract.
func bindICheckpointRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICheckpointRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointRegistry *ICheckpointRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICheckpointRegistry.Contract.ICheckpointRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointRegistry *ICheckpointRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointRegistry.Contract.ICheckpointRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointRegistry *ICheckpointRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointRegistry.Contract.ICheckpointRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointRegistry *ICheckpointRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICheckpointRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointRegistry *ICheckpointRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointRegistry *ICheckpointRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointRegistry.Contract.contract.Transact(opts, method, params...)
}

// CPPSigner is a free data retrieval call binding the contract method 0xd59f1758.
//
// Solidity: function CPP_signer() constant returns(address)
func (_ICheckpointRegistry *ICheckpointRegistryCaller) CPPSigner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ICheckpointRegistry.contract.Call(opts, out, "CPP_signer")
	return *ret0, err
}

// CPPSigner is a free data retrieval call binding the contract method 0xd59f1758.
//
// Solidity: function CPP_signer() constant returns(address)
func (_ICheckpointRegistry *ICheckpointRegistrySession) CPPSigner() (common.Address, error) {
	return _ICheckpointRegistry.Contract.CPPSigner(&_ICheckpointRegistry.CallOpts)
}

// CPPSigner is a free data retrieval call binding the contract method 0xd59f1758.
//
// Solidity: function CPP_signer() constant returns(address)
func (_ICheckpointRegistry *ICheckpointRegistryCallerSession) CPPSigner() (common.Address, error) {
	return _ICheckpointRegistry.Contract.CPPSigner(&_ICheckpointRegistry.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0x5a48c0b0.
//
// Solidity: function checkpoints() constant returns(address[])
func (_ICheckpointRegistry *ICheckpointRegistryCaller) Checkpoints(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ICheckpointRegistry.contract.Call(opts, out, "checkpoints")
	return *ret0, err
}

// Checkpoints is a free data retrieval call binding the contract method 0x5a48c0b0.
//
// Solidity: function checkpoints() constant returns(address[])
func (_ICheckpointRegistry *ICheckpointRegistrySession) Checkpoints() ([]common.Address, error) {
	return _ICheckpointRegistry.Contract.Checkpoints(&_ICheckpointRegistry.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0x5a48c0b0.
//
// Solidity: function checkpoints() constant returns(address[])
func (_ICheckpointRegistry *ICheckpointRegistryCallerSession) Checkpoints() ([]common.Address, error) {
	return _ICheckpointRegistry.Contract.Checkpoints(&_ICheckpointRegistry.CallOpts)
}

// SignatureBase is a free data retrieval call binding the contract method 0x851f2209.
//
// Solidity: function signatureBase(uint256 number, bytes32 hash) constant returns(bytes32 sigbase)
func (_ICheckpointRegistry *ICheckpointRegistryCaller) SignatureBase(opts *bind.CallOpts, number *big.Int, hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ICheckpointRegistry.contract.Call(opts, out, "signatureBase", number, hash)
	return *ret0, err
}

// SignatureBase is a free data retrieval call binding the contract method 0x851f2209.
//
// Solidity: function signatureBase(uint256 number, bytes32 hash) constant returns(bytes32 sigbase)
func (_ICheckpointRegistry *ICheckpointRegistrySession) SignatureBase(number *big.Int, hash [32]byte) ([32]byte, error) {
	return _ICheckpointRegistry.Contract.SignatureBase(&_ICheckpointRegistry.CallOpts, number, hash)
}

// SignatureBase is a free data retrieval call binding the contract method 0x851f2209.
//
// Solidity: function signatureBase(uint256 number, bytes32 hash) constant returns(bytes32 sigbase)
func (_ICheckpointRegistry *ICheckpointRegistryCallerSession) SignatureBase(number *big.Int, hash [32]byte) ([32]byte, error) {
	return _ICheckpointRegistry.Contract.SignatureBase(&_ICheckpointRegistry.CallOpts, number, hash)
}

// Propose is a paid mutator transaction binding the contract method 0xc20fa2ee.
//
// Solidity: function propose(uint256 number, bytes32 hash, bytes signature) returns(address)
func (_ICheckpointRegistry *ICheckpointRegistryTransactor) Propose(opts *bind.TransactOpts, number *big.Int, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistry.contract.Transact(opts, "propose", number, hash, signature)
}

// Propose is a paid mutator transaction binding the contract method 0xc20fa2ee.
//
// Solidity: function propose(uint256 number, bytes32 hash, bytes signature) returns(address)
func (_ICheckpointRegistry *ICheckpointRegistrySession) Propose(number *big.Int, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistry.Contract.Propose(&_ICheckpointRegistry.TransactOpts, number, hash, signature)
}

// Propose is a paid mutator transaction binding the contract method 0xc20fa2ee.
//
// Solidity: function propose(uint256 number, bytes32 hash, bytes signature) returns(address)
func (_ICheckpointRegistry *ICheckpointRegistryTransactorSession) Propose(number *big.Int, hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistry.Contract.Propose(&_ICheckpointRegistry.TransactOpts, number, hash, signature)
}

// Sign is a paid mutator transaction binding the contract method 0x51fae959.
//
// Solidity: function sign(address checkpoint, bytes signature) returns()
func (_ICheckpointRegistry *ICheckpointRegistryTransactor) Sign(opts *bind.TransactOpts, checkpoint common.Address, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistry.contract.Transact(opts, "sign", checkpoint, signature)
}

// Sign is a paid mutator transaction binding the contract method 0x51fae959.
//
// Solidity: function sign(address checkpoint, bytes signature) returns()
func (_ICheckpointRegistry *ICheckpointRegistrySession) Sign(checkpoint common.Address, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistry.Contract.Sign(&_ICheckpointRegistry.TransactOpts, checkpoint, signature)
}

// Sign is a paid mutator transaction binding the contract method 0x51fae959.
//
// Solidity: function sign(address checkpoint, bytes signature) returns()
func (_ICheckpointRegistry *ICheckpointRegistryTransactorSession) Sign(checkpoint common.Address, signature []byte) (*types.Transaction, error) {
	return _ICheckpointRegistry.Contract.Sign(&_ICheckpointRegistry.TransactOpts, checkpoint, signature)
}

// ICheckpointRegistryCheckpointIterator is returned from FilterCheckpoint and is used to iterate over the raw logs and unpacked data for Checkpoint events raised by the ICheckpointRegistry contract.
type ICheckpointRegistryCheckpointIterator struct {
	Event *ICheckpointRegistryCheckpoint // Event containing the contract specifics and raw log

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
func (it *ICheckpointRegistryCheckpointIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICheckpointRegistryCheckpoint)
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
		it.Event = new(ICheckpointRegistryCheckpoint)
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
func (it *ICheckpointRegistryCheckpointIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICheckpointRegistryCheckpointIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICheckpointRegistryCheckpoint represents a Checkpoint event raised by the ICheckpointRegistry contract.
type ICheckpointRegistryCheckpoint struct {
	Number     *big.Int
	Hash       [32]byte
	Checkpoint common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCheckpoint is a free log retrieval operation binding the contract event 0x7f582e7a234c68b135245365b0c69d608f9235392a1ef801b46f04ca5d6cdad4.
//
// Solidity: event Checkpoint(uint256 indexed number, bytes32 hash, address checkpoint)
func (_ICheckpointRegistry *ICheckpointRegistryFilterer) FilterCheckpoint(opts *bind.FilterOpts, number []*big.Int) (*ICheckpointRegistryCheckpointIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}

	logs, sub, err := _ICheckpointRegistry.contract.FilterLogs(opts, "Checkpoint", numberRule)
	if err != nil {
		return nil, err
	}
	return &ICheckpointRegistryCheckpointIterator{contract: _ICheckpointRegistry.contract, event: "Checkpoint", logs: logs, sub: sub}, nil
}

// WatchCheckpoint is a free log subscription operation binding the contract event 0x7f582e7a234c68b135245365b0c69d608f9235392a1ef801b46f04ca5d6cdad4.
//
// Solidity: event Checkpoint(uint256 indexed number, bytes32 hash, address checkpoint)
func (_ICheckpointRegistry *ICheckpointRegistryFilterer) WatchCheckpoint(opts *bind.WatchOpts, sink chan<- *ICheckpointRegistryCheckpoint, number []*big.Int) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}

	logs, sub, err := _ICheckpointRegistry.contract.WatchLogs(opts, "Checkpoint", numberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICheckpointRegistryCheckpoint)
				if err := _ICheckpointRegistry.contract.UnpackLog(event, "Checkpoint", log); err != nil {
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
