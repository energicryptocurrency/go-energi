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

// IBlacklistRegistryABI is the input ABI used to generate the binding from.
const IBlacklistRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIBlacklistProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"BlacklistProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIBlacklistProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"DrainProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIBlacklistProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"WhitelistProposal\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EBI_signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compensation_fund\",\"outputs\":[{\"internalType\":\"contractITreasury\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"item_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes20\",\"name\":\"owner\",\"type\":\"bytes20\"}],\"name\":\"drainMigration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerateAll\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerateBlocked\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerateDrainable\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isDrainable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"onDrain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"contractIBlacklistProposal\",\"name\":\"enforce\",\"type\":\"address\"},{\"internalType\":\"contractIBlacklistProposal\",\"name\":\"revoke\",\"type\":\"address\"},{\"internalType\":\"contractIBlacklistProposal\",\"name\":\"drain\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"contractIBlacklistProposal\",\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proposeDrain\",\"outputs\":[{\"internalType\":\"contractIBlacklistProposal\",\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proposeRevoke\",\"outputs\":[{\"internalType\":\"contractIBlacklistProposal\",\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IBlacklistRegistry is an auto generated Go binding around an Ethereum contract.
type IBlacklistRegistry struct {
	IBlacklistRegistryCaller     // Read-only binding to the contract
	IBlacklistRegistryTransactor // Write-only binding to the contract
	IBlacklistRegistryFilterer   // Log filterer for contract events
}

// IBlacklistRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlacklistRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlacklistRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlacklistRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlacklistRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlacklistRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlacklistRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlacklistRegistrySession struct {
	Contract     *IBlacklistRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBlacklistRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlacklistRegistryCallerSession struct {
	Contract *IBlacklistRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IBlacklistRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlacklistRegistryTransactorSession struct {
	Contract     *IBlacklistRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IBlacklistRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlacklistRegistryRaw struct {
	Contract *IBlacklistRegistry // Generic contract binding to access the raw methods on
}

// IBlacklistRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlacklistRegistryCallerRaw struct {
	Contract *IBlacklistRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IBlacklistRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlacklistRegistryTransactorRaw struct {
	Contract *IBlacklistRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlacklistRegistry creates a new instance of IBlacklistRegistry, bound to a specific deployed contract.
func NewIBlacklistRegistry(address common.Address, backend bind.ContractBackend) (*IBlacklistRegistry, error) {
	contract, err := bindIBlacklistRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlacklistRegistry{IBlacklistRegistryCaller: IBlacklistRegistryCaller{contract: contract}, IBlacklistRegistryTransactor: IBlacklistRegistryTransactor{contract: contract}, IBlacklistRegistryFilterer: IBlacklistRegistryFilterer{contract: contract}}, nil
}

// NewIBlacklistRegistryCaller creates a new read-only instance of IBlacklistRegistry, bound to a specific deployed contract.
func NewIBlacklistRegistryCaller(address common.Address, caller bind.ContractCaller) (*IBlacklistRegistryCaller, error) {
	contract, err := bindIBlacklistRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlacklistRegistryCaller{contract: contract}, nil
}

// NewIBlacklistRegistryTransactor creates a new write-only instance of IBlacklistRegistry, bound to a specific deployed contract.
func NewIBlacklistRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlacklistRegistryTransactor, error) {
	contract, err := bindIBlacklistRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlacklistRegistryTransactor{contract: contract}, nil
}

// NewIBlacklistRegistryFilterer creates a new log filterer instance of IBlacklistRegistry, bound to a specific deployed contract.
func NewIBlacklistRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlacklistRegistryFilterer, error) {
	contract, err := bindIBlacklistRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlacklistRegistryFilterer{contract: contract}, nil
}

// bindIBlacklistRegistry binds a generic wrapper to an already deployed contract.
func bindIBlacklistRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBlacklistRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlacklistRegistry *IBlacklistRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBlacklistRegistry.Contract.IBlacklistRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlacklistRegistry *IBlacklistRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.IBlacklistRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlacklistRegistry *IBlacklistRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.IBlacklistRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlacklistRegistry *IBlacklistRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBlacklistRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlacklistRegistry *IBlacklistRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlacklistRegistry *IBlacklistRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.contract.Transact(opts, method, params...)
}

// EBISigner is a free data retrieval call binding the contract method 0x94c210fc.
//
// Solidity: function EBI_signer() constant returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryCaller) EBISigner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IBlacklistRegistry.contract.Call(opts, out, "EBI_signer")
	return *ret0, err
}

// EBISigner is a free data retrieval call binding the contract method 0x94c210fc.
//
// Solidity: function EBI_signer() constant returns(address)
func (_IBlacklistRegistry *IBlacklistRegistrySession) EBISigner() (common.Address, error) {
	return _IBlacklistRegistry.Contract.EBISigner(&_IBlacklistRegistry.CallOpts)
}

// EBISigner is a free data retrieval call binding the contract method 0x94c210fc.
//
// Solidity: function EBI_signer() constant returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryCallerSession) EBISigner() (common.Address, error) {
	return _IBlacklistRegistry.Contract.EBISigner(&_IBlacklistRegistry.CallOpts)
}

// CompensationFund is a free data retrieval call binding the contract method 0xf4441152.
//
// Solidity: function compensation_fund() constant returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryCaller) CompensationFund(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IBlacklistRegistry.contract.Call(opts, out, "compensation_fund")
	return *ret0, err
}

// CompensationFund is a free data retrieval call binding the contract method 0xf4441152.
//
// Solidity: function compensation_fund() constant returns(address)
func (_IBlacklistRegistry *IBlacklistRegistrySession) CompensationFund() (common.Address, error) {
	return _IBlacklistRegistry.Contract.CompensationFund(&_IBlacklistRegistry.CallOpts)
}

// CompensationFund is a free data retrieval call binding the contract method 0xf4441152.
//
// Solidity: function compensation_fund() constant returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryCallerSession) CompensationFund() (common.Address, error) {
	return _IBlacklistRegistry.Contract.CompensationFund(&_IBlacklistRegistry.CallOpts)
}

// EnumerateAll is a free data retrieval call binding the contract method 0xbc393afb.
//
// Solidity: function enumerateAll() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistryCaller) EnumerateAll(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IBlacklistRegistry.contract.Call(opts, out, "enumerateAll")
	return *ret0, err
}

// EnumerateAll is a free data retrieval call binding the contract method 0xbc393afb.
//
// Solidity: function enumerateAll() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistrySession) EnumerateAll() ([]common.Address, error) {
	return _IBlacklistRegistry.Contract.EnumerateAll(&_IBlacklistRegistry.CallOpts)
}

// EnumerateAll is a free data retrieval call binding the contract method 0xbc393afb.
//
// Solidity: function enumerateAll() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistryCallerSession) EnumerateAll() ([]common.Address, error) {
	return _IBlacklistRegistry.Contract.EnumerateAll(&_IBlacklistRegistry.CallOpts)
}

// EnumerateBlocked is a free data retrieval call binding the contract method 0x5603125c.
//
// Solidity: function enumerateBlocked() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistryCaller) EnumerateBlocked(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IBlacklistRegistry.contract.Call(opts, out, "enumerateBlocked")
	return *ret0, err
}

// EnumerateBlocked is a free data retrieval call binding the contract method 0x5603125c.
//
// Solidity: function enumerateBlocked() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistrySession) EnumerateBlocked() ([]common.Address, error) {
	return _IBlacklistRegistry.Contract.EnumerateBlocked(&_IBlacklistRegistry.CallOpts)
}

// EnumerateBlocked is a free data retrieval call binding the contract method 0x5603125c.
//
// Solidity: function enumerateBlocked() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistryCallerSession) EnumerateBlocked() ([]common.Address, error) {
	return _IBlacklistRegistry.Contract.EnumerateBlocked(&_IBlacklistRegistry.CallOpts)
}

// EnumerateDrainable is a free data retrieval call binding the contract method 0xee33f9d2.
//
// Solidity: function enumerateDrainable() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistryCaller) EnumerateDrainable(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IBlacklistRegistry.contract.Call(opts, out, "enumerateDrainable")
	return *ret0, err
}

// EnumerateDrainable is a free data retrieval call binding the contract method 0xee33f9d2.
//
// Solidity: function enumerateDrainable() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistrySession) EnumerateDrainable() ([]common.Address, error) {
	return _IBlacklistRegistry.Contract.EnumerateDrainable(&_IBlacklistRegistry.CallOpts)
}

// EnumerateDrainable is a free data retrieval call binding the contract method 0xee33f9d2.
//
// Solidity: function enumerateDrainable() constant returns(address[] addresses)
func (_IBlacklistRegistry *IBlacklistRegistryCallerSession) EnumerateDrainable() ([]common.Address, error) {
	return _IBlacklistRegistry.Contract.EnumerateDrainable(&_IBlacklistRegistry.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) constant returns(bool)
func (_IBlacklistRegistry *IBlacklistRegistryCaller) IsBlacklisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IBlacklistRegistry.contract.Call(opts, out, "isBlacklisted", arg0)
	return *ret0, err
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) constant returns(bool)
func (_IBlacklistRegistry *IBlacklistRegistrySession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _IBlacklistRegistry.Contract.IsBlacklisted(&_IBlacklistRegistry.CallOpts, arg0)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) constant returns(bool)
func (_IBlacklistRegistry *IBlacklistRegistryCallerSession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _IBlacklistRegistry.Contract.IsBlacklisted(&_IBlacklistRegistry.CallOpts, arg0)
}

// IsDrainable is a free data retrieval call binding the contract method 0x3303bbb1.
//
// Solidity: function isDrainable(address ) constant returns(bool)
func (_IBlacklistRegistry *IBlacklistRegistryCaller) IsDrainable(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IBlacklistRegistry.contract.Call(opts, out, "isDrainable", arg0)
	return *ret0, err
}

// IsDrainable is a free data retrieval call binding the contract method 0x3303bbb1.
//
// Solidity: function isDrainable(address ) constant returns(bool)
func (_IBlacklistRegistry *IBlacklistRegistrySession) IsDrainable(arg0 common.Address) (bool, error) {
	return _IBlacklistRegistry.Contract.IsDrainable(&_IBlacklistRegistry.CallOpts, arg0)
}

// IsDrainable is a free data retrieval call binding the contract method 0x3303bbb1.
//
// Solidity: function isDrainable(address ) constant returns(bool)
func (_IBlacklistRegistry *IBlacklistRegistryCallerSession) IsDrainable(arg0 common.Address) (bool, error) {
	return _IBlacklistRegistry.Contract.IsDrainable(&_IBlacklistRegistry.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x3341b445.
//
// Solidity: function proposals(address ) constant returns(address enforce, address revoke, address drain)
func (_IBlacklistRegistry *IBlacklistRegistryCaller) Proposals(opts *bind.CallOpts, arg0 common.Address) (struct {
	Enforce common.Address
	Revoke  common.Address
	Drain   common.Address
}, error) {
	ret := new(struct {
		Enforce common.Address
		Revoke  common.Address
		Drain   common.Address
	})
	out := ret
	err := _IBlacklistRegistry.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x3341b445.
//
// Solidity: function proposals(address ) constant returns(address enforce, address revoke, address drain)
func (_IBlacklistRegistry *IBlacklistRegistrySession) Proposals(arg0 common.Address) (struct {
	Enforce common.Address
	Revoke  common.Address
	Drain   common.Address
}, error) {
	return _IBlacklistRegistry.Contract.Proposals(&_IBlacklistRegistry.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x3341b445.
//
// Solidity: function proposals(address ) constant returns(address enforce, address revoke, address drain)
func (_IBlacklistRegistry *IBlacklistRegistryCallerSession) Proposals(arg0 common.Address) (struct {
	Enforce common.Address
	Revoke  common.Address
	Drain   common.Address
}, error) {
	return _IBlacklistRegistry.Contract.Proposals(&_IBlacklistRegistry.CallOpts, arg0)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address ) returns()
func (_IBlacklistRegistry *IBlacklistRegistryTransactor) Collect(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.contract.Transact(opts, "collect", arg0)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address ) returns()
func (_IBlacklistRegistry *IBlacklistRegistrySession) Collect(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.Collect(&_IBlacklistRegistry.TransactOpts, arg0)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address ) returns()
func (_IBlacklistRegistry *IBlacklistRegistryTransactorSession) Collect(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.Collect(&_IBlacklistRegistry.TransactOpts, arg0)
}

// DrainMigration is a paid mutator transaction binding the contract method 0x3502a130.
//
// Solidity: function drainMigration(uint256 item_id, bytes20 owner) returns()
func (_IBlacklistRegistry *IBlacklistRegistryTransactor) DrainMigration(opts *bind.TransactOpts, item_id *big.Int, owner [20]byte) (*types.Transaction, error) {
	return _IBlacklistRegistry.contract.Transact(opts, "drainMigration", item_id, owner)
}

// DrainMigration is a paid mutator transaction binding the contract method 0x3502a130.
//
// Solidity: function drainMigration(uint256 item_id, bytes20 owner) returns()
func (_IBlacklistRegistry *IBlacklistRegistrySession) DrainMigration(item_id *big.Int, owner [20]byte) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.DrainMigration(&_IBlacklistRegistry.TransactOpts, item_id, owner)
}

// DrainMigration is a paid mutator transaction binding the contract method 0x3502a130.
//
// Solidity: function drainMigration(uint256 item_id, bytes20 owner) returns()
func (_IBlacklistRegistry *IBlacklistRegistryTransactorSession) DrainMigration(item_id *big.Int, owner [20]byte) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.DrainMigration(&_IBlacklistRegistry.TransactOpts, item_id, owner)
}

// OnDrain is a paid mutator transaction binding the contract method 0x79d7bc07.
//
// Solidity: function onDrain(address ) returns()
func (_IBlacklistRegistry *IBlacklistRegistryTransactor) OnDrain(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.contract.Transact(opts, "onDrain", arg0)
}

// OnDrain is a paid mutator transaction binding the contract method 0x79d7bc07.
//
// Solidity: function onDrain(address ) returns()
func (_IBlacklistRegistry *IBlacklistRegistrySession) OnDrain(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.OnDrain(&_IBlacklistRegistry.TransactOpts, arg0)
}

// OnDrain is a paid mutator transaction binding the contract method 0x79d7bc07.
//
// Solidity: function onDrain(address ) returns()
func (_IBlacklistRegistry *IBlacklistRegistryTransactorSession) OnDrain(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.OnDrain(&_IBlacklistRegistry.TransactOpts, arg0)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryTransactor) Propose(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.contract.Transact(opts, "propose", arg0)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistrySession) Propose(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.Propose(&_IBlacklistRegistry.TransactOpts, arg0)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryTransactorSession) Propose(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.Propose(&_IBlacklistRegistry.TransactOpts, arg0)
}

// ProposeDrain is a paid mutator transaction binding the contract method 0xd1e8ebda.
//
// Solidity: function proposeDrain(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryTransactor) ProposeDrain(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.contract.Transact(opts, "proposeDrain", arg0)
}

// ProposeDrain is a paid mutator transaction binding the contract method 0xd1e8ebda.
//
// Solidity: function proposeDrain(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistrySession) ProposeDrain(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.ProposeDrain(&_IBlacklistRegistry.TransactOpts, arg0)
}

// ProposeDrain is a paid mutator transaction binding the contract method 0xd1e8ebda.
//
// Solidity: function proposeDrain(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryTransactorSession) ProposeDrain(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.ProposeDrain(&_IBlacklistRegistry.TransactOpts, arg0)
}

// ProposeRevoke is a paid mutator transaction binding the contract method 0x244fcea5.
//
// Solidity: function proposeRevoke(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryTransactor) ProposeRevoke(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.contract.Transact(opts, "proposeRevoke", arg0)
}

// ProposeRevoke is a paid mutator transaction binding the contract method 0x244fcea5.
//
// Solidity: function proposeRevoke(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistrySession) ProposeRevoke(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.ProposeRevoke(&_IBlacklistRegistry.TransactOpts, arg0)
}

// ProposeRevoke is a paid mutator transaction binding the contract method 0x244fcea5.
//
// Solidity: function proposeRevoke(address ) returns(address)
func (_IBlacklistRegistry *IBlacklistRegistryTransactorSession) ProposeRevoke(arg0 common.Address) (*types.Transaction, error) {
	return _IBlacklistRegistry.Contract.ProposeRevoke(&_IBlacklistRegistry.TransactOpts, arg0)
}

// IBlacklistRegistryBlacklistProposalIterator is returned from FilterBlacklistProposal and is used to iterate over the raw logs and unpacked data for BlacklistProposal events raised by the IBlacklistRegistry contract.
type IBlacklistRegistryBlacklistProposalIterator struct {
	Event *IBlacklistRegistryBlacklistProposal // Event containing the contract specifics and raw log

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
func (it *IBlacklistRegistryBlacklistProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlacklistRegistryBlacklistProposal)
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
		it.Event = new(IBlacklistRegistryBlacklistProposal)
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
func (it *IBlacklistRegistryBlacklistProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlacklistRegistryBlacklistProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlacklistRegistryBlacklistProposal represents a BlacklistProposal event raised by the IBlacklistRegistry contract.
type IBlacklistRegistryBlacklistProposal struct {
	Target   common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlacklistProposal is a free log retrieval operation binding the contract event 0xea0d3b1afc7aa1754f902b277d6f473e2f4a4526e94797814c2042405d692557.
//
// Solidity: event BlacklistProposal(address indexed target, address proposal)
func (_IBlacklistRegistry *IBlacklistRegistryFilterer) FilterBlacklistProposal(opts *bind.FilterOpts, target []common.Address) (*IBlacklistRegistryBlacklistProposalIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IBlacklistRegistry.contract.FilterLogs(opts, "BlacklistProposal", targetRule)
	if err != nil {
		return nil, err
	}
	return &IBlacklistRegistryBlacklistProposalIterator{contract: _IBlacklistRegistry.contract, event: "BlacklistProposal", logs: logs, sub: sub}, nil
}

// WatchBlacklistProposal is a free log subscription operation binding the contract event 0xea0d3b1afc7aa1754f902b277d6f473e2f4a4526e94797814c2042405d692557.
//
// Solidity: event BlacklistProposal(address indexed target, address proposal)
func (_IBlacklistRegistry *IBlacklistRegistryFilterer) WatchBlacklistProposal(opts *bind.WatchOpts, sink chan<- *IBlacklistRegistryBlacklistProposal, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IBlacklistRegistry.contract.WatchLogs(opts, "BlacklistProposal", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlacklistRegistryBlacklistProposal)
				if err := _IBlacklistRegistry.contract.UnpackLog(event, "BlacklistProposal", log); err != nil {
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

// IBlacklistRegistryDrainProposalIterator is returned from FilterDrainProposal and is used to iterate over the raw logs and unpacked data for DrainProposal events raised by the IBlacklistRegistry contract.
type IBlacklistRegistryDrainProposalIterator struct {
	Event *IBlacklistRegistryDrainProposal // Event containing the contract specifics and raw log

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
func (it *IBlacklistRegistryDrainProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlacklistRegistryDrainProposal)
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
		it.Event = new(IBlacklistRegistryDrainProposal)
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
func (it *IBlacklistRegistryDrainProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlacklistRegistryDrainProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlacklistRegistryDrainProposal represents a DrainProposal event raised by the IBlacklistRegistry contract.
type IBlacklistRegistryDrainProposal struct {
	Target   common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDrainProposal is a free log retrieval operation binding the contract event 0xb0163b33033bcea41a78d2d4a9c596c29b0667259543d6d56b8cac8cd92d2cea.
//
// Solidity: event DrainProposal(address indexed target, address proposal)
func (_IBlacklistRegistry *IBlacklistRegistryFilterer) FilterDrainProposal(opts *bind.FilterOpts, target []common.Address) (*IBlacklistRegistryDrainProposalIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IBlacklistRegistry.contract.FilterLogs(opts, "DrainProposal", targetRule)
	if err != nil {
		return nil, err
	}
	return &IBlacklistRegistryDrainProposalIterator{contract: _IBlacklistRegistry.contract, event: "DrainProposal", logs: logs, sub: sub}, nil
}

// WatchDrainProposal is a free log subscription operation binding the contract event 0xb0163b33033bcea41a78d2d4a9c596c29b0667259543d6d56b8cac8cd92d2cea.
//
// Solidity: event DrainProposal(address indexed target, address proposal)
func (_IBlacklistRegistry *IBlacklistRegistryFilterer) WatchDrainProposal(opts *bind.WatchOpts, sink chan<- *IBlacklistRegistryDrainProposal, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IBlacklistRegistry.contract.WatchLogs(opts, "DrainProposal", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlacklistRegistryDrainProposal)
				if err := _IBlacklistRegistry.contract.UnpackLog(event, "DrainProposal", log); err != nil {
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

// IBlacklistRegistryWhitelistProposalIterator is returned from FilterWhitelistProposal and is used to iterate over the raw logs and unpacked data for WhitelistProposal events raised by the IBlacklistRegistry contract.
type IBlacklistRegistryWhitelistProposalIterator struct {
	Event *IBlacklistRegistryWhitelistProposal // Event containing the contract specifics and raw log

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
func (it *IBlacklistRegistryWhitelistProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlacklistRegistryWhitelistProposal)
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
		it.Event = new(IBlacklistRegistryWhitelistProposal)
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
func (it *IBlacklistRegistryWhitelistProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlacklistRegistryWhitelistProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlacklistRegistryWhitelistProposal represents a WhitelistProposal event raised by the IBlacklistRegistry contract.
type IBlacklistRegistryWhitelistProposal struct {
	Target   common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistProposal is a free log retrieval operation binding the contract event 0x2e46fe8e502a44b05a85d4346d3e208fc845c81a81c9b7e4db08a33dca59faff.
//
// Solidity: event WhitelistProposal(address indexed target, address proposal)
func (_IBlacklistRegistry *IBlacklistRegistryFilterer) FilterWhitelistProposal(opts *bind.FilterOpts, target []common.Address) (*IBlacklistRegistryWhitelistProposalIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IBlacklistRegistry.contract.FilterLogs(opts, "WhitelistProposal", targetRule)
	if err != nil {
		return nil, err
	}
	return &IBlacklistRegistryWhitelistProposalIterator{contract: _IBlacklistRegistry.contract, event: "WhitelistProposal", logs: logs, sub: sub}, nil
}

// WatchWhitelistProposal is a free log subscription operation binding the contract event 0x2e46fe8e502a44b05a85d4346d3e208fc845c81a81c9b7e4db08a33dca59faff.
//
// Solidity: event WhitelistProposal(address indexed target, address proposal)
func (_IBlacklistRegistry *IBlacklistRegistryFilterer) WatchWhitelistProposal(opts *bind.WatchOpts, sink chan<- *IBlacklistRegistryWhitelistProposal, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IBlacklistRegistry.contract.WatchLogs(opts, "WhitelistProposal", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlacklistRegistryWhitelistProposal)
				if err := _IBlacklistRegistry.contract.UnpackLog(event, "WhitelistProposal", log); err != nil {
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
