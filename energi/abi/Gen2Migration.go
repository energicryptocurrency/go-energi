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
const Gen2MigrationABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"block_number\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

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

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256)
func (_Gen2Migration *Gen2MigrationCaller) GetReward(opts *bind.CallOpts, block_number *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gen2Migration.contract.Call(opts, out, "getReward", block_number)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256)
func (_Gen2Migration *Gen2MigrationSession) GetReward(block_number *big.Int) (*big.Int, error) {
	return _Gen2Migration.Contract.GetReward(&_Gen2Migration.CallOpts, block_number)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256)
func (_Gen2Migration *Gen2MigrationCallerSession) GetReward(block_number *big.Int) (*big.Int, error) {
	return _Gen2Migration.Contract.GetReward(&_Gen2Migration.CallOpts, block_number)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Gen2Migration *Gen2MigrationTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gen2Migration.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Gen2Migration *Gen2MigrationSession) Migrate() (*types.Transaction, error) {
	return _Gen2Migration.Contract.Migrate(&_Gen2Migration.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Gen2Migration *Gen2MigrationTransactorSession) Migrate() (*types.Transaction, error) {
	return _Gen2Migration.Contract.Migrate(&_Gen2Migration.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_Gen2Migration *Gen2MigrationTransactor) Reward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gen2Migration.contract.Transact(opts, "reward", amount)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_Gen2Migration *Gen2MigrationSession) Reward(amount *big.Int) (*types.Transaction, error) {
	return _Gen2Migration.Contract.Reward(&_Gen2Migration.TransactOpts, amount)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_Gen2Migration *Gen2MigrationTransactorSession) Reward(amount *big.Int) (*types.Transaction, error) {
	return _Gen2Migration.Contract.Reward(&_Gen2Migration.TransactOpts, amount)
}
