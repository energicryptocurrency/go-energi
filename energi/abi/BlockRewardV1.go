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

// BlockRewardV1ABI is the input ABI used to generate the binding from.
const BlockRewardV1ABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// BlockRewardV1Bin is the compiled bytecode used for deploying new contracts.
const BlockRewardV1Bin = `608060405234801561001057600080fd5b506040516103f13803806103f18339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b031990921691909117905561038c806100656000396000f3fe6080604052600436106100595760003560e01c8063228cb73311610043578063228cb7331461013e578063ce5494bb14610146578063ec5568891461018657610059565b8062f55d9d146100c05780631c4b774b14610102575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100cc57600080fd5b50610100600480360360208110156100e357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101c4565b005b34801561010e57600080fd5b5061012c6004803603602081101561012557600080fd5b503561026c565b60408051918252519081900360200190f35b61010061027d565b34801561015257600080fd5b506101006004803603602081101561016957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166102ac565b34801561019257600080fd5b5061019b61033b565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461024a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610253816102a9565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60008115610278575060015b919050565b60405141903480156108fc02916000818181858888f193505050501580156102a9573d6000803e3d6000fd5b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461033257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6102a9816102a9565b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a72315820a6b12432cc1ef71cf70d1bcd1020ceb40e09f24497f9cab39244e691a617dddf64736f6c634300050b0032`

// DeployBlockRewardV1 deploys a new Ethereum contract, binding an instance of BlockRewardV1 to it.
func DeployBlockRewardV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *BlockRewardV1, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockRewardV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlockRewardV1Bin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockRewardV1{BlockRewardV1Caller: BlockRewardV1Caller{contract: contract}, BlockRewardV1Transactor: BlockRewardV1Transactor{contract: contract}, BlockRewardV1Filterer: BlockRewardV1Filterer{contract: contract}}, nil
}

// BlockRewardV1Bin is the compiled bytecode of contract after deployment.
const BlockRewardV1RuntimeBin = `6080604052600436106100595760003560e01c8063228cb73311610043578063228cb7331461013e578063ce5494bb14610146578063ec5568891461018657610059565b8062f55d9d146100c05780631c4b774b14610102575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100cc57600080fd5b50610100600480360360208110156100e357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101c4565b005b34801561010e57600080fd5b5061012c6004803603602081101561012557600080fd5b503561026c565b60408051918252519081900360200190f35b61010061027d565b34801561015257600080fd5b506101006004803603602081101561016957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166102ac565b34801561019257600080fd5b5061019b61033b565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461024a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610253816102a9565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60008115610278575060015b919050565b60405141903480156108fc02916000818181858888f193505050501580156102a9573d6000803e3d6000fd5b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461033257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6102a9816102a9565b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a72315820a6b12432cc1ef71cf70d1bcd1020ceb40e09f24497f9cab39244e691a617dddf64736f6c634300050b0032`

// BlockRewardV1 is an auto generated Go binding around an Ethereum contract.
type BlockRewardV1 struct {
	BlockRewardV1Caller     // Read-only binding to the contract
	BlockRewardV1Transactor // Write-only binding to the contract
	BlockRewardV1Filterer   // Log filterer for contract events
}

// BlockRewardV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type BlockRewardV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockRewardV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockRewardV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockRewardV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockRewardV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockRewardV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockRewardV1Session struct {
	Contract     *BlockRewardV1    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockRewardV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockRewardV1CallerSession struct {
	Contract *BlockRewardV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BlockRewardV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockRewardV1TransactorSession struct {
	Contract     *BlockRewardV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BlockRewardV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type BlockRewardV1Raw struct {
	Contract *BlockRewardV1 // Generic contract binding to access the raw methods on
}

// BlockRewardV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockRewardV1CallerRaw struct {
	Contract *BlockRewardV1Caller // Generic read-only contract binding to access the raw methods on
}

// BlockRewardV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockRewardV1TransactorRaw struct {
	Contract *BlockRewardV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockRewardV1 creates a new instance of BlockRewardV1, bound to a specific deployed contract.
func NewBlockRewardV1(address common.Address, backend bind.ContractBackend) (*BlockRewardV1, error) {
	contract, err := bindBlockRewardV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockRewardV1{BlockRewardV1Caller: BlockRewardV1Caller{contract: contract}, BlockRewardV1Transactor: BlockRewardV1Transactor{contract: contract}, BlockRewardV1Filterer: BlockRewardV1Filterer{contract: contract}}, nil
}

// NewBlockRewardV1Caller creates a new read-only instance of BlockRewardV1, bound to a specific deployed contract.
func NewBlockRewardV1Caller(address common.Address, caller bind.ContractCaller) (*BlockRewardV1Caller, error) {
	contract, err := bindBlockRewardV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockRewardV1Caller{contract: contract}, nil
}

// NewBlockRewardV1Transactor creates a new write-only instance of BlockRewardV1, bound to a specific deployed contract.
func NewBlockRewardV1Transactor(address common.Address, transactor bind.ContractTransactor) (*BlockRewardV1Transactor, error) {
	contract, err := bindBlockRewardV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockRewardV1Transactor{contract: contract}, nil
}

// NewBlockRewardV1Filterer creates a new log filterer instance of BlockRewardV1, bound to a specific deployed contract.
func NewBlockRewardV1Filterer(address common.Address, filterer bind.ContractFilterer) (*BlockRewardV1Filterer, error) {
	contract, err := bindBlockRewardV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockRewardV1Filterer{contract: contract}, nil
}

// bindBlockRewardV1 binds a generic wrapper to an already deployed contract.
func bindBlockRewardV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockRewardV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockRewardV1 *BlockRewardV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlockRewardV1.Contract.BlockRewardV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockRewardV1 *BlockRewardV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.BlockRewardV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockRewardV1 *BlockRewardV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.BlockRewardV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockRewardV1 *BlockRewardV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlockRewardV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockRewardV1 *BlockRewardV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockRewardV1 *BlockRewardV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.contract.Transact(opts, method, params...)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BlockRewardV1 *BlockRewardV1Caller) GetReward(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BlockRewardV1.contract.Call(opts, out, "getReward", _blockNumber)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BlockRewardV1 *BlockRewardV1Session) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _BlockRewardV1.Contract.GetReward(&_BlockRewardV1.CallOpts, _blockNumber)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BlockRewardV1 *BlockRewardV1CallerSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _BlockRewardV1.Contract.GetReward(&_BlockRewardV1.CallOpts, _blockNumber)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlockRewardV1 *BlockRewardV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlockRewardV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlockRewardV1 *BlockRewardV1Session) Proxy() (common.Address, error) {
	return _BlockRewardV1.Contract.Proxy(&_BlockRewardV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlockRewardV1 *BlockRewardV1CallerSession) Proxy() (common.Address, error) {
	return _BlockRewardV1.Contract.Proxy(&_BlockRewardV1.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlockRewardV1 *BlockRewardV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlockRewardV1 *BlockRewardV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Destroy(&_BlockRewardV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlockRewardV1 *BlockRewardV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Destroy(&_BlockRewardV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlockRewardV1 *BlockRewardV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlockRewardV1 *BlockRewardV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Migrate(&_BlockRewardV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlockRewardV1 *BlockRewardV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Migrate(&_BlockRewardV1.TransactOpts, _oldImpl)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BlockRewardV1 *BlockRewardV1Transactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockRewardV1.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BlockRewardV1 *BlockRewardV1Session) Reward() (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Reward(&_BlockRewardV1.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BlockRewardV1 *BlockRewardV1TransactorSession) Reward() (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Reward(&_BlockRewardV1.TransactOpts)
}
