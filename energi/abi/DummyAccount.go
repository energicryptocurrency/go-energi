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

// DummyAccountABI is the input ABI used to generate the binding from.
const DummyAccountABI = "[{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// DummyAccountBin is the compiled bytecode used for deploying new contracts.
const DummyAccountBin = `6080604052348015600f57600080fd5b50609180601d6000396000f3fe608060408190527f08c379a00000000000000000000000000000000000000000000000000000000081526020608452600560a4527f44756d6d7900000000000000000000000000000000000000000000000000000060c452606490fdfea265627a7a72315820823eba240dbde1120777ee3252b021152be5a461ba9ab0c879d72ca52a284c6264736f6c63430005100032`

// DeployDummyAccount deploys a new Ethereum contract, binding an instance of DummyAccount to it.
func DeployDummyAccount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DummyAccount, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyAccountABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DummyAccountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DummyAccount{DummyAccountCaller: DummyAccountCaller{contract: contract}, DummyAccountTransactor: DummyAccountTransactor{contract: contract}, DummyAccountFilterer: DummyAccountFilterer{contract: contract}}, nil
}

// DummyAccountBin is the compiled bytecode of contract after deployment.
const DummyAccountRuntimeBin = `608060408190527f08c379a00000000000000000000000000000000000000000000000000000000081526020608452600560a4527f44756d6d7900000000000000000000000000000000000000000000000000000060c452606490fdfea265627a7a72315820823eba240dbde1120777ee3252b021152be5a461ba9ab0c879d72ca52a284c6264736f6c63430005100032`

// DummyAccount is an auto generated Go binding around an Ethereum contract.
type DummyAccount struct {
	DummyAccountCaller     // Read-only binding to the contract
	DummyAccountTransactor // Write-only binding to the contract
	DummyAccountFilterer   // Log filterer for contract events
}

// DummyAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyAccountSession struct {
	Contract     *DummyAccount     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyAccountCallerSession struct {
	Contract *DummyAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DummyAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyAccountTransactorSession struct {
	Contract     *DummyAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DummyAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyAccountRaw struct {
	Contract *DummyAccount // Generic contract binding to access the raw methods on
}

// DummyAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyAccountCallerRaw struct {
	Contract *DummyAccountCaller // Generic read-only contract binding to access the raw methods on
}

// DummyAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyAccountTransactorRaw struct {
	Contract *DummyAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyAccount creates a new instance of DummyAccount, bound to a specific deployed contract.
func NewDummyAccount(address common.Address, backend bind.ContractBackend) (*DummyAccount, error) {
	contract, err := bindDummyAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyAccount{DummyAccountCaller: DummyAccountCaller{contract: contract}, DummyAccountTransactor: DummyAccountTransactor{contract: contract}, DummyAccountFilterer: DummyAccountFilterer{contract: contract}}, nil
}

// NewDummyAccountCaller creates a new read-only instance of DummyAccount, bound to a specific deployed contract.
func NewDummyAccountCaller(address common.Address, caller bind.ContractCaller) (*DummyAccountCaller, error) {
	contract, err := bindDummyAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyAccountCaller{contract: contract}, nil
}

// NewDummyAccountTransactor creates a new write-only instance of DummyAccount, bound to a specific deployed contract.
func NewDummyAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyAccountTransactor, error) {
	contract, err := bindDummyAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyAccountTransactor{contract: contract}, nil
}

// NewDummyAccountFilterer creates a new log filterer instance of DummyAccount, bound to a specific deployed contract.
func NewDummyAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyAccountFilterer, error) {
	contract, err := bindDummyAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyAccountFilterer{contract: contract}, nil
}

// bindDummyAccount binds a generic wrapper to an already deployed contract.
func bindDummyAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyAccount *DummyAccountRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DummyAccount.Contract.DummyAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyAccount *DummyAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyAccount.Contract.DummyAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyAccount *DummyAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyAccount.Contract.DummyAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyAccount *DummyAccountCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DummyAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyAccount *DummyAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyAccount *DummyAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyAccount.Contract.contract.Transact(opts, method, params...)
}
