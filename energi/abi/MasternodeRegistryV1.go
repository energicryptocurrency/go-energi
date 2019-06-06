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

// MasternodeRegistryV1ABI is the input ABI used to generate the binding from.
const MasternodeRegistryV1ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"block_number\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// MasternodeRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const MasternodeRegistryV1Bin = `608060405234801561001057600080fd5b5060ec8061001f6000396000f3fe60806040526004361060305760003560e01c80631c4b774b1460355780638fd3ab8014606d578063a9fb763c146081575b600080fd5b348015604057600080fd5b50605b60048036036020811015605557600080fd5b5035609b565b60408051918252519081900360200190f35b348015607857600080fd5b50607f60b2565b005b607f60048036036020811015609557600080fd5b503560b4565b6000811560ad5750677ed7cd92ff1200005b919050565b565b5056fea265627a7a72305820af485dbe7fbb0e3d16211552f37d8d13adb802e1c250da69238cf0144b0d745264736f6c63430005090032`

// DeployMasternodeRegistryV1 deploys a new Ethereum contract, binding an instance of MasternodeRegistryV1 to it.
func DeployMasternodeRegistryV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MasternodeRegistryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeRegistryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MasternodeRegistryV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MasternodeRegistryV1{MasternodeRegistryV1Caller: MasternodeRegistryV1Caller{contract: contract}, MasternodeRegistryV1Transactor: MasternodeRegistryV1Transactor{contract: contract}, MasternodeRegistryV1Filterer: MasternodeRegistryV1Filterer{contract: contract}}, nil
}

// MasternodeRegistryV1 is an auto generated Go binding around an Ethereum contract.
type MasternodeRegistryV1 struct {
	MasternodeRegistryV1Caller     // Read-only binding to the contract
	MasternodeRegistryV1Transactor // Write-only binding to the contract
	MasternodeRegistryV1Filterer   // Log filterer for contract events
}

// MasternodeRegistryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasternodeRegistryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeRegistryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasternodeRegistryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeRegistryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasternodeRegistryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeRegistryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasternodeRegistryV1Session struct {
	Contract     *MasternodeRegistryV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MasternodeRegistryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasternodeRegistryV1CallerSession struct {
	Contract *MasternodeRegistryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MasternodeRegistryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasternodeRegistryV1TransactorSession struct {
	Contract     *MasternodeRegistryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MasternodeRegistryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasternodeRegistryV1Raw struct {
	Contract *MasternodeRegistryV1 // Generic contract binding to access the raw methods on
}

// MasternodeRegistryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasternodeRegistryV1CallerRaw struct {
	Contract *MasternodeRegistryV1Caller // Generic read-only contract binding to access the raw methods on
}

// MasternodeRegistryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasternodeRegistryV1TransactorRaw struct {
	Contract *MasternodeRegistryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasternodeRegistryV1 creates a new instance of MasternodeRegistryV1, bound to a specific deployed contract.
func NewMasternodeRegistryV1(address common.Address, backend bind.ContractBackend) (*MasternodeRegistryV1, error) {
	contract, err := bindMasternodeRegistryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1{MasternodeRegistryV1Caller: MasternodeRegistryV1Caller{contract: contract}, MasternodeRegistryV1Transactor: MasternodeRegistryV1Transactor{contract: contract}, MasternodeRegistryV1Filterer: MasternodeRegistryV1Filterer{contract: contract}}, nil
}

// NewMasternodeRegistryV1Caller creates a new read-only instance of MasternodeRegistryV1, bound to a specific deployed contract.
func NewMasternodeRegistryV1Caller(address common.Address, caller bind.ContractCaller) (*MasternodeRegistryV1Caller, error) {
	contract, err := bindMasternodeRegistryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1Caller{contract: contract}, nil
}

// NewMasternodeRegistryV1Transactor creates a new write-only instance of MasternodeRegistryV1, bound to a specific deployed contract.
func NewMasternodeRegistryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*MasternodeRegistryV1Transactor, error) {
	contract, err := bindMasternodeRegistryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1Transactor{contract: contract}, nil
}

// NewMasternodeRegistryV1Filterer creates a new log filterer instance of MasternodeRegistryV1, bound to a specific deployed contract.
func NewMasternodeRegistryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*MasternodeRegistryV1Filterer, error) {
	contract, err := bindMasternodeRegistryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1Filterer{contract: contract}, nil
}

// bindMasternodeRegistryV1 binds a generic wrapper to an already deployed contract.
func bindMasternodeRegistryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeRegistryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeRegistryV1 *MasternodeRegistryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeRegistryV1.Contract.MasternodeRegistryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeRegistryV1 *MasternodeRegistryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.MasternodeRegistryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeRegistryV1 *MasternodeRegistryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.MasternodeRegistryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeRegistryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.contract.Transact(opts, method, params...)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) GetReward(opts *bind.CallOpts, block_number *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "getReward", block_number)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) GetReward(block_number *big.Int) (*big.Int, error) {
	return _MasternodeRegistryV1.Contract.GetReward(&_MasternodeRegistryV1.CallOpts, block_number)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 block_number) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) GetReward(block_number *big.Int) (*big.Int, error) {
	return _MasternodeRegistryV1.Contract.GetReward(&_MasternodeRegistryV1.CallOpts, block_number)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Migrate() (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Migrate(&_MasternodeRegistryV1.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Migrate() (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Migrate(&_MasternodeRegistryV1.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Reward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "reward", amount)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Reward(amount *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Reward(&_MasternodeRegistryV1.TransactOpts, amount)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 amount) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Reward(amount *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Reward(&_MasternodeRegistryV1.TransactOpts, amount)
}
