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

// IDelegatedPoSABI is the input ABI used to generate the binding from.
const IDelegatedPoSABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"signerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IDelegatedPoS is an auto generated Go binding around an Ethereum contract.
type IDelegatedPoS struct {
	IDelegatedPoSCaller     // Read-only binding to the contract
	IDelegatedPoSTransactor // Write-only binding to the contract
	IDelegatedPoSFilterer   // Log filterer for contract events
}

// IDelegatedPoSCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDelegatedPoSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegatedPoSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDelegatedPoSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegatedPoSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDelegatedPoSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegatedPoSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDelegatedPoSSession struct {
	Contract     *IDelegatedPoS    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDelegatedPoSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDelegatedPoSCallerSession struct {
	Contract *IDelegatedPoSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IDelegatedPoSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDelegatedPoSTransactorSession struct {
	Contract     *IDelegatedPoSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IDelegatedPoSRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDelegatedPoSRaw struct {
	Contract *IDelegatedPoS // Generic contract binding to access the raw methods on
}

// IDelegatedPoSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDelegatedPoSCallerRaw struct {
	Contract *IDelegatedPoSCaller // Generic read-only contract binding to access the raw methods on
}

// IDelegatedPoSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDelegatedPoSTransactorRaw struct {
	Contract *IDelegatedPoSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDelegatedPoS creates a new instance of IDelegatedPoS, bound to a specific deployed contract.
func NewIDelegatedPoS(address common.Address, backend bind.ContractBackend) (*IDelegatedPoS, error) {
	contract, err := bindIDelegatedPoS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDelegatedPoS{IDelegatedPoSCaller: IDelegatedPoSCaller{contract: contract}, IDelegatedPoSTransactor: IDelegatedPoSTransactor{contract: contract}, IDelegatedPoSFilterer: IDelegatedPoSFilterer{contract: contract}}, nil
}

// NewIDelegatedPoSCaller creates a new read-only instance of IDelegatedPoS, bound to a specific deployed contract.
func NewIDelegatedPoSCaller(address common.Address, caller bind.ContractCaller) (*IDelegatedPoSCaller, error) {
	contract, err := bindIDelegatedPoS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDelegatedPoSCaller{contract: contract}, nil
}

// NewIDelegatedPoSTransactor creates a new write-only instance of IDelegatedPoS, bound to a specific deployed contract.
func NewIDelegatedPoSTransactor(address common.Address, transactor bind.ContractTransactor) (*IDelegatedPoSTransactor, error) {
	contract, err := bindIDelegatedPoS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDelegatedPoSTransactor{contract: contract}, nil
}

// NewIDelegatedPoSFilterer creates a new log filterer instance of IDelegatedPoS, bound to a specific deployed contract.
func NewIDelegatedPoSFilterer(address common.Address, filterer bind.ContractFilterer) (*IDelegatedPoSFilterer, error) {
	contract, err := bindIDelegatedPoS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDelegatedPoSFilterer{contract: contract}, nil
}

// bindIDelegatedPoS binds a generic wrapper to an already deployed contract.
func bindIDelegatedPoS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDelegatedPoSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelegatedPoS *IDelegatedPoSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IDelegatedPoS.Contract.IDelegatedPoSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelegatedPoS *IDelegatedPoSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelegatedPoS.Contract.IDelegatedPoSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelegatedPoS *IDelegatedPoSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelegatedPoS.Contract.IDelegatedPoSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelegatedPoS *IDelegatedPoSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IDelegatedPoS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelegatedPoS *IDelegatedPoSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelegatedPoS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelegatedPoS *IDelegatedPoSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelegatedPoS.Contract.contract.Transact(opts, method, params...)
}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() constant returns(address)
func (_IDelegatedPoS *IDelegatedPoSCaller) SignerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IDelegatedPoS.contract.Call(opts, out, "signerAddress")
	return *ret0, err
}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() constant returns(address)
func (_IDelegatedPoS *IDelegatedPoSSession) SignerAddress() (common.Address, error) {
	return _IDelegatedPoS.Contract.SignerAddress(&_IDelegatedPoS.CallOpts)
}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() constant returns(address)
func (_IDelegatedPoS *IDelegatedPoSCallerSession) SignerAddress() (common.Address, error) {
	return _IDelegatedPoS.Contract.SignerAddress(&_IDelegatedPoS.CallOpts)
}
