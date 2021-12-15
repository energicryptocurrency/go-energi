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

// IBlockRewardABI is the input ABI used to generate the binding from.
const IBlockRewardABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IBlockReward is an auto generated Go binding around an Ethereum contract.
type IBlockReward struct {
	IBlockRewardCaller     // Read-only binding to the contract
	IBlockRewardTransactor // Write-only binding to the contract
	IBlockRewardFilterer   // Log filterer for contract events
}

// IBlockRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlockRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlockRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlockRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlockRewardSession struct {
	Contract     *IBlockReward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlockRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlockRewardCallerSession struct {
	Contract *IBlockRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IBlockRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlockRewardTransactorSession struct {
	Contract     *IBlockRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IBlockRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlockRewardRaw struct {
	Contract *IBlockReward // Generic contract binding to access the raw methods on
}

// IBlockRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlockRewardCallerRaw struct {
	Contract *IBlockRewardCaller // Generic read-only contract binding to access the raw methods on
}

// IBlockRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlockRewardTransactorRaw struct {
	Contract *IBlockRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlockReward creates a new instance of IBlockReward, bound to a specific deployed contract.
func NewIBlockReward(address common.Address, backend bind.ContractBackend) (*IBlockReward, error) {
	contract, err := bindIBlockReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlockReward{IBlockRewardCaller: IBlockRewardCaller{contract: contract}, IBlockRewardTransactor: IBlockRewardTransactor{contract: contract}, IBlockRewardFilterer: IBlockRewardFilterer{contract: contract}}, nil
}

// NewIBlockRewardCaller creates a new read-only instance of IBlockReward, bound to a specific deployed contract.
func NewIBlockRewardCaller(address common.Address, caller bind.ContractCaller) (*IBlockRewardCaller, error) {
	contract, err := bindIBlockReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlockRewardCaller{contract: contract}, nil
}

// NewIBlockRewardTransactor creates a new write-only instance of IBlockReward, bound to a specific deployed contract.
func NewIBlockRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlockRewardTransactor, error) {
	contract, err := bindIBlockReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlockRewardTransactor{contract: contract}, nil
}

// NewIBlockRewardFilterer creates a new log filterer instance of IBlockReward, bound to a specific deployed contract.
func NewIBlockRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlockRewardFilterer, error) {
	contract, err := bindIBlockReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlockRewardFilterer{contract: contract}, nil
}

// bindIBlockReward binds a generic wrapper to an already deployed contract.
func bindIBlockReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBlockRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlockReward *IBlockRewardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBlockReward.Contract.IBlockRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlockReward *IBlockRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlockReward.Contract.IBlockRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlockReward *IBlockRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlockReward.Contract.IBlockRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlockReward *IBlockRewardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBlockReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlockReward *IBlockRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlockReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlockReward *IBlockRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlockReward.Contract.contract.Transact(opts, method, params...)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_IBlockReward *IBlockRewardCaller) GetReward(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBlockReward.contract.Call(opts, out, "getReward", _blockNumber)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_IBlockReward *IBlockRewardSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _IBlockReward.Contract.GetReward(&_IBlockReward.CallOpts, _blockNumber)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_IBlockReward *IBlockRewardCallerSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _IBlockReward.Contract.GetReward(&_IBlockReward.CallOpts, _blockNumber)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_IBlockReward *IBlockRewardTransactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlockReward.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_IBlockReward *IBlockRewardSession) Reward() (*types.Transaction, error) {
	return _IBlockReward.Contract.Reward(&_IBlockReward.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_IBlockReward *IBlockRewardTransactorSession) Reward() (*types.Transaction, error) {
	return _IBlockReward.Contract.Reward(&_IBlockReward.TransactOpts)
}
