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

// ISporkRegistryABI is the input ABI used to generate the binding from.
const ISporkRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"consensusGasLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"callGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"xferGas\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_impl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_fee_payer\",\"type\":\"address\"}],\"name\":\"createUpgradeProposal\",\"outputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ISporkRegistry is an auto generated Go binding around an Ethereum contract.
type ISporkRegistry struct {
	ISporkRegistryCaller     // Read-only binding to the contract
	ISporkRegistryTransactor // Write-only binding to the contract
	ISporkRegistryFilterer   // Log filterer for contract events
}

// ISporkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISporkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISporkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISporkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISporkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISporkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISporkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISporkRegistrySession struct {
	Contract     *ISporkRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISporkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISporkRegistryCallerSession struct {
	Contract *ISporkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ISporkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISporkRegistryTransactorSession struct {
	Contract     *ISporkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ISporkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISporkRegistryRaw struct {
	Contract *ISporkRegistry // Generic contract binding to access the raw methods on
}

// ISporkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISporkRegistryCallerRaw struct {
	Contract *ISporkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ISporkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISporkRegistryTransactorRaw struct {
	Contract *ISporkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISporkRegistry creates a new instance of ISporkRegistry, bound to a specific deployed contract.
func NewISporkRegistry(address common.Address, backend bind.ContractBackend) (*ISporkRegistry, error) {
	contract, err := bindISporkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISporkRegistry{ISporkRegistryCaller: ISporkRegistryCaller{contract: contract}, ISporkRegistryTransactor: ISporkRegistryTransactor{contract: contract}, ISporkRegistryFilterer: ISporkRegistryFilterer{contract: contract}}, nil
}

// NewISporkRegistryCaller creates a new read-only instance of ISporkRegistry, bound to a specific deployed contract.
func NewISporkRegistryCaller(address common.Address, caller bind.ContractCaller) (*ISporkRegistryCaller, error) {
	contract, err := bindISporkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISporkRegistryCaller{contract: contract}, nil
}

// NewISporkRegistryTransactor creates a new write-only instance of ISporkRegistry, bound to a specific deployed contract.
func NewISporkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ISporkRegistryTransactor, error) {
	contract, err := bindISporkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISporkRegistryTransactor{contract: contract}, nil
}

// NewISporkRegistryFilterer creates a new log filterer instance of ISporkRegistry, bound to a specific deployed contract.
func NewISporkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ISporkRegistryFilterer, error) {
	contract, err := bindISporkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISporkRegistryFilterer{contract: contract}, nil
}

// bindISporkRegistry binds a generic wrapper to an already deployed contract.
func bindISporkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISporkRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISporkRegistry *ISporkRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISporkRegistry.Contract.ISporkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISporkRegistry *ISporkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISporkRegistry.Contract.ISporkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISporkRegistry *ISporkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISporkRegistry.Contract.ISporkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISporkRegistry *ISporkRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISporkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISporkRegistry *ISporkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISporkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISporkRegistry *ISporkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISporkRegistry.Contract.contract.Transact(opts, method, params...)
}

// ConsensusGasLimits is a free data retrieval call binding the contract method 0xc00ebced.
//
// Solidity: function consensusGasLimits() constant returns(uint256 callGas, uint256 xferGas)
func (_ISporkRegistry *ISporkRegistryCaller) ConsensusGasLimits(opts *bind.CallOpts) (struct {
	CallGas *big.Int
	XferGas *big.Int
}, error) {
	ret := new(struct {
		CallGas *big.Int
		XferGas *big.Int
	})
	out := ret
	err := _ISporkRegistry.contract.Call(opts, out, "consensusGasLimits")
	return *ret, err
}

// ConsensusGasLimits is a free data retrieval call binding the contract method 0xc00ebced.
//
// Solidity: function consensusGasLimits() constant returns(uint256 callGas, uint256 xferGas)
func (_ISporkRegistry *ISporkRegistrySession) ConsensusGasLimits() (struct {
	CallGas *big.Int
	XferGas *big.Int
}, error) {
	return _ISporkRegistry.Contract.ConsensusGasLimits(&_ISporkRegistry.CallOpts)
}

// ConsensusGasLimits is a free data retrieval call binding the contract method 0xc00ebced.
//
// Solidity: function consensusGasLimits() constant returns(uint256 callGas, uint256 xferGas)
func (_ISporkRegistry *ISporkRegistryCallerSession) ConsensusGasLimits() (struct {
	CallGas *big.Int
	XferGas *big.Int
}, error) {
	return _ISporkRegistry.Contract.ConsensusGasLimits(&_ISporkRegistry.CallOpts)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x62877ccd.
//
// Solidity: function createUpgradeProposal(address _impl, uint256 _period, address _fee_payer) returns(address)
func (_ISporkRegistry *ISporkRegistryTransactor) CreateUpgradeProposal(opts *bind.TransactOpts, _impl common.Address, _period *big.Int, _fee_payer common.Address) (*types.Transaction, error) {
	return _ISporkRegistry.contract.Transact(opts, "createUpgradeProposal", _impl, _period, _fee_payer)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x62877ccd.
//
// Solidity: function createUpgradeProposal(address _impl, uint256 _period, address _fee_payer) returns(address)
func (_ISporkRegistry *ISporkRegistrySession) CreateUpgradeProposal(_impl common.Address, _period *big.Int, _fee_payer common.Address) (*types.Transaction, error) {
	return _ISporkRegistry.Contract.CreateUpgradeProposal(&_ISporkRegistry.TransactOpts, _impl, _period, _fee_payer)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x62877ccd.
//
// Solidity: function createUpgradeProposal(address _impl, uint256 _period, address _fee_payer) returns(address)
func (_ISporkRegistry *ISporkRegistryTransactorSession) CreateUpgradeProposal(_impl common.Address, _period *big.Int, _fee_payer common.Address) (*types.Transaction, error) {
	return _ISporkRegistry.Contract.CreateUpgradeProposal(&_ISporkRegistry.TransactOpts, _impl, _period, _fee_payer)
}
