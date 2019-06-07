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

// Gen2MigrationABI is the input ABI used to generate the binding from.
const Gen2MigrationABI = "[]"

// Gen2MigrationBin is the compiled bytecode used for deploying new contracts.
const Gen2MigrationBin = `6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723058208f6e3a5d66b12cc6a160f0fae73dd1923adde39308e8ac8b866b69120d4b4e4664736f6c63430005090032`

// DeployGen2Migration deploys a new Ethereum contract, binding an instance of Gen2Migration to it.
func DeployGen2Migration(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gen2Migration, error) {
	parsed, err := abi.JSON(strings.NewReader(Gen2MigrationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Gen2MigrationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gen2Migration{Gen2MigrationCaller: Gen2MigrationCaller{contract: contract}, Gen2MigrationTransactor: Gen2MigrationTransactor{contract: contract}, Gen2MigrationFilterer: Gen2MigrationFilterer{contract: contract}}, nil
}

// Gen2Migration is an auto generated Go binding around an Ethereum contract.
type Gen2Migration struct {
	Gen2MigrationCaller     // Read-only binding to the contract
	Gen2MigrationTransactor // Write-only binding to the contract
	Gen2MigrationFilterer   // Log filterer for contract events
}

// Gen2MigrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type Gen2MigrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gen2MigrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Gen2MigrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gen2MigrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Gen2MigrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gen2MigrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Gen2MigrationSession struct {
	Contract     *Gen2Migration    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Gen2MigrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Gen2MigrationCallerSession struct {
	Contract *Gen2MigrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Gen2MigrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Gen2MigrationTransactorSession struct {
	Contract     *Gen2MigrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Gen2MigrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type Gen2MigrationRaw struct {
	Contract *Gen2Migration // Generic contract binding to access the raw methods on
}

// Gen2MigrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Gen2MigrationCallerRaw struct {
	Contract *Gen2MigrationCaller // Generic read-only contract binding to access the raw methods on
}

// Gen2MigrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Gen2MigrationTransactorRaw struct {
	Contract *Gen2MigrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGen2Migration creates a new instance of Gen2Migration, bound to a specific deployed contract.
func NewGen2Migration(address common.Address, backend bind.ContractBackend) (*Gen2Migration, error) {
	contract, err := bindGen2Migration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gen2Migration{Gen2MigrationCaller: Gen2MigrationCaller{contract: contract}, Gen2MigrationTransactor: Gen2MigrationTransactor{contract: contract}, Gen2MigrationFilterer: Gen2MigrationFilterer{contract: contract}}, nil
}

// NewGen2MigrationCaller creates a new read-only instance of Gen2Migration, bound to a specific deployed contract.
func NewGen2MigrationCaller(address common.Address, caller bind.ContractCaller) (*Gen2MigrationCaller, error) {
	contract, err := bindGen2Migration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Gen2MigrationCaller{contract: contract}, nil
}

// NewGen2MigrationTransactor creates a new write-only instance of Gen2Migration, bound to a specific deployed contract.
func NewGen2MigrationTransactor(address common.Address, transactor bind.ContractTransactor) (*Gen2MigrationTransactor, error) {
	contract, err := bindGen2Migration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Gen2MigrationTransactor{contract: contract}, nil
}

// NewGen2MigrationFilterer creates a new log filterer instance of Gen2Migration, bound to a specific deployed contract.
func NewGen2MigrationFilterer(address common.Address, filterer bind.ContractFilterer) (*Gen2MigrationFilterer, error) {
	contract, err := bindGen2Migration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Gen2MigrationFilterer{contract: contract}, nil
}

// bindGen2Migration binds a generic wrapper to an already deployed contract.
func bindGen2Migration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Gen2MigrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gen2Migration *Gen2MigrationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gen2Migration.Contract.Gen2MigrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gen2Migration *Gen2MigrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gen2Migration.Contract.Gen2MigrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gen2Migration *Gen2MigrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gen2Migration.Contract.Gen2MigrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gen2Migration *Gen2MigrationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gen2Migration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gen2Migration *Gen2MigrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gen2Migration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gen2Migration *Gen2MigrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gen2Migration.Contract.contract.Transact(opts, method, params...)
}
