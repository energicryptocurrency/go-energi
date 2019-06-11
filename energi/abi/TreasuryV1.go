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

// TreasuryV1ABI is the input ABI used to generate the binding from.
const TreasuryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"superblock_cycle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"v1storage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"isSuperblock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"},{\"name\":\"_superblock_cycle\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Contribution\",\"type\":\"event\"}]"

// TreasuryV1Bin is the compiled bytecode used for deploying new contracts.
const TreasuryV1Bin = `608060405234801561001057600080fd5b506040516108da3803806108da8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b0319166001600160a01b038416179055604051610063906100b3565b604051809103906000f08015801561007f573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b03929092169190911790556002819055806100ac57fe5b50506100c0565b610228806106b283390190565b6105e3806100cf6000396000f3fe6080604052600436106100bb5760003560e01c80632d05930511610074578063ce5494bb1161004e578063ce5494bb1461028e578063d7bb99ba146102ce578063ec556889146102d6576100bb565b80632d059305146101fd57806352782d861461023b578063b69ef8a814610279576100bb565b80630ef34745116100a55780630ef34745146101a45780631c4b774b146101cb578063228cb733146101f5576100bb565b8062f55d9d1461012257806306ec16f814610164575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561012e57600080fd5b506101626004803603602081101561014557600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166102eb565b005b34801561017057600080fd5b506101626004803603602081101561018757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610393565b3480156101b057600080fd5b506101b96103f6565b60408051918252519081900360200190f35b3480156101d757600080fd5b506101b9600480360360208110156101ee57600080fd5b50356103fc565b610162610416565b34801561020957600080fd5b50610212610418565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561024757600080fd5b506102656004803603602081101561025e57600080fd5b5035610434565b604080519115158252519081900360200190f35b34801561028557600080fd5b506101b9610449565b34801561029a57600080fd5b50610162600480360360208110156102b157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661044e565b6101626104dc565b3480156102e257600080fd5b5061021261051e565b60005473ffffffffffffffffffffffffffffffffffffffff16331461037157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61037a8161053a565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b8073ffffffffffffffffffffffffffffffffffffffff1663e52253816040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156103db57600080fd5b505af11580156103ef573d6000803e3d6000fd5b5050505050565b60025481565b6000811561041157506926f6a8f4e638030000005b919050565b565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6000600254828161044157fe5b061592915050565b303190565b60005473ffffffffffffffffffffffffffffffffffffffff1633146104d457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6104d9815b50565b3415610416576040805133815234602082015281517f4d154d4aae216bed6d0926db77c00df2b57c6b5ba4eee05775de20facede3a7b929181900390910190a1565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600154604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b1580156103db57600080fdfea265627a7a72305820d56b2ab8e3ace05e37cb2f1f8e1e74f7bea8e14a3c4dded6fc81890dd1d9a9c864736f6c634300050900326080604052600080546001600160a01b03191633179055610203806100256000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806313af40351461003b57806341c0e1b514610070575b600080fd5b61006e6004803603602081101561005157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610078565b005b61006e610145565b60005473ffffffffffffffffffffffffffffffffffffffff1633146100fe57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101cb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b33fffea265627a7a72305820ce2aa02fd153709e8547b34d6f3fdbc5fe09a8e727a8ae50ed881b46906855c364736f6c63430005090032`

// DeployTreasuryV1 deploys a new Ethereum contract, binding an instance of TreasuryV1 to it.
func DeployTreasuryV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _superblock_cycle *big.Int) (common.Address, *types.Transaction, *TreasuryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TreasuryV1Bin), backend, _proxy, _superblock_cycle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TreasuryV1{TreasuryV1Caller: TreasuryV1Caller{contract: contract}, TreasuryV1Transactor: TreasuryV1Transactor{contract: contract}, TreasuryV1Filterer: TreasuryV1Filterer{contract: contract}}, nil
}

// TreasuryV1 is an auto generated Go binding around an Ethereum contract.
type TreasuryV1 struct {
	TreasuryV1Caller     // Read-only binding to the contract
	TreasuryV1Transactor // Write-only binding to the contract
	TreasuryV1Filterer   // Log filterer for contract events
}

// TreasuryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasuryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasuryV1Session struct {
	Contract     *TreasuryV1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryV1CallerSession struct {
	Contract *TreasuryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TreasuryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryV1TransactorSession struct {
	Contract     *TreasuryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TreasuryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryV1Raw struct {
	Contract *TreasuryV1 // Generic contract binding to access the raw methods on
}

// TreasuryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryV1CallerRaw struct {
	Contract *TreasuryV1Caller // Generic read-only contract binding to access the raw methods on
}

// TreasuryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryV1TransactorRaw struct {
	Contract *TreasuryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasuryV1 creates a new instance of TreasuryV1, bound to a specific deployed contract.
func NewTreasuryV1(address common.Address, backend bind.ContractBackend) (*TreasuryV1, error) {
	contract, err := bindTreasuryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasuryV1{TreasuryV1Caller: TreasuryV1Caller{contract: contract}, TreasuryV1Transactor: TreasuryV1Transactor{contract: contract}, TreasuryV1Filterer: TreasuryV1Filterer{contract: contract}}, nil
}

// NewTreasuryV1Caller creates a new read-only instance of TreasuryV1, bound to a specific deployed contract.
func NewTreasuryV1Caller(address common.Address, caller bind.ContractCaller) (*TreasuryV1Caller, error) {
	contract, err := bindTreasuryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryV1Caller{contract: contract}, nil
}

// NewTreasuryV1Transactor creates a new write-only instance of TreasuryV1, bound to a specific deployed contract.
func NewTreasuryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryV1Transactor, error) {
	contract, err := bindTreasuryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryV1Transactor{contract: contract}, nil
}

// NewTreasuryV1Filterer creates a new log filterer instance of TreasuryV1, bound to a specific deployed contract.
func NewTreasuryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryV1Filterer, error) {
	contract, err := bindTreasuryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryV1Filterer{contract: contract}, nil
}

// bindTreasuryV1 binds a generic wrapper to an already deployed contract.
func bindTreasuryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryV1 *TreasuryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TreasuryV1.Contract.TreasuryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryV1 *TreasuryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryV1.Contract.TreasuryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryV1 *TreasuryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryV1.Contract.TreasuryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryV1 *TreasuryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TreasuryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryV1 *TreasuryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryV1 *TreasuryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryV1.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1Caller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TreasuryV1.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1Session) Balance() (*big.Int, error) {
	return _TreasuryV1.Contract.Balance(&_TreasuryV1.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1CallerSession) Balance() (*big.Int, error) {
	return _TreasuryV1.Contract.Balance(&_TreasuryV1.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1Caller) GetReward(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TreasuryV1.contract.Call(opts, out, "getReward", _blockNumber)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1Session) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _TreasuryV1.Contract.GetReward(&_TreasuryV1.CallOpts, _blockNumber)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_TreasuryV1 *TreasuryV1CallerSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _TreasuryV1.Contract.GetReward(&_TreasuryV1.CallOpts, _blockNumber)
}

// IsSuperblock is a free data retrieval call binding the contract method 0x52782d86.
//
// Solidity: function isSuperblock(uint256 _blockNumber) constant returns(bool)
func (_TreasuryV1 *TreasuryV1Caller) IsSuperblock(opts *bind.CallOpts, _blockNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TreasuryV1.contract.Call(opts, out, "isSuperblock", _blockNumber)
	return *ret0, err
}

// IsSuperblock is a free data retrieval call binding the contract method 0x52782d86.
//
// Solidity: function isSuperblock(uint256 _blockNumber) constant returns(bool)
func (_TreasuryV1 *TreasuryV1Session) IsSuperblock(_blockNumber *big.Int) (bool, error) {
	return _TreasuryV1.Contract.IsSuperblock(&_TreasuryV1.CallOpts, _blockNumber)
}

// IsSuperblock is a free data retrieval call binding the contract method 0x52782d86.
//
// Solidity: function isSuperblock(uint256 _blockNumber) constant returns(bool)
func (_TreasuryV1 *TreasuryV1CallerSession) IsSuperblock(_blockNumber *big.Int) (bool, error) {
	return _TreasuryV1.Contract.IsSuperblock(&_TreasuryV1.CallOpts, _blockNumber)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_TreasuryV1 *TreasuryV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TreasuryV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_TreasuryV1 *TreasuryV1Session) Proxy() (common.Address, error) {
	return _TreasuryV1.Contract.Proxy(&_TreasuryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_TreasuryV1 *TreasuryV1CallerSession) Proxy() (common.Address, error) {
	return _TreasuryV1.Contract.Proxy(&_TreasuryV1.CallOpts)
}

// SuperblockCycle is a free data retrieval call binding the contract method 0x0ef34745.
//
// Solidity: function superblock_cycle() constant returns(uint256)
func (_TreasuryV1 *TreasuryV1Caller) SuperblockCycle(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TreasuryV1.contract.Call(opts, out, "superblock_cycle")
	return *ret0, err
}

// SuperblockCycle is a free data retrieval call binding the contract method 0x0ef34745.
//
// Solidity: function superblock_cycle() constant returns(uint256)
func (_TreasuryV1 *TreasuryV1Session) SuperblockCycle() (*big.Int, error) {
	return _TreasuryV1.Contract.SuperblockCycle(&_TreasuryV1.CallOpts)
}

// SuperblockCycle is a free data retrieval call binding the contract method 0x0ef34745.
//
// Solidity: function superblock_cycle() constant returns(uint256)
func (_TreasuryV1 *TreasuryV1CallerSession) SuperblockCycle() (*big.Int, error) {
	return _TreasuryV1.Contract.SuperblockCycle(&_TreasuryV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_TreasuryV1 *TreasuryV1Caller) V1storage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TreasuryV1.contract.Call(opts, out, "v1storage")
	return *ret0, err
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_TreasuryV1 *TreasuryV1Session) V1storage() (common.Address, error) {
	return _TreasuryV1.Contract.V1storage(&_TreasuryV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_TreasuryV1 *TreasuryV1CallerSession) V1storage() (common.Address, error) {
	return _TreasuryV1.Contract.V1storage(&_TreasuryV1.CallOpts)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address proposal) returns()
func (_TreasuryV1 *TreasuryV1Transactor) Collect(opts *bind.TransactOpts, proposal common.Address) (*types.Transaction, error) {
	return _TreasuryV1.contract.Transact(opts, "collect", proposal)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address proposal) returns()
func (_TreasuryV1 *TreasuryV1Session) Collect(proposal common.Address) (*types.Transaction, error) {
	return _TreasuryV1.Contract.Collect(&_TreasuryV1.TransactOpts, proposal)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address proposal) returns()
func (_TreasuryV1 *TreasuryV1TransactorSession) Collect(proposal common.Address) (*types.Transaction, error) {
	return _TreasuryV1.Contract.Collect(&_TreasuryV1.TransactOpts, proposal)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() returns()
func (_TreasuryV1 *TreasuryV1Transactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryV1.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() returns()
func (_TreasuryV1 *TreasuryV1Session) Contribute() (*types.Transaction, error) {
	return _TreasuryV1.Contract.Contribute(&_TreasuryV1.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() returns()
func (_TreasuryV1 *TreasuryV1TransactorSession) Contribute() (*types.Transaction, error) {
	return _TreasuryV1.Contract.Contribute(&_TreasuryV1.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_TreasuryV1 *TreasuryV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _TreasuryV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_TreasuryV1 *TreasuryV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _TreasuryV1.Contract.Destroy(&_TreasuryV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_TreasuryV1 *TreasuryV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _TreasuryV1.Contract.Destroy(&_TreasuryV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_TreasuryV1 *TreasuryV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _TreasuryV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_TreasuryV1 *TreasuryV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _TreasuryV1.Contract.Migrate(&_TreasuryV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_TreasuryV1 *TreasuryV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _TreasuryV1.Contract.Migrate(&_TreasuryV1.TransactOpts, _oldImpl)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_TreasuryV1 *TreasuryV1Transactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryV1.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_TreasuryV1 *TreasuryV1Session) Reward() (*types.Transaction, error) {
	return _TreasuryV1.Contract.Reward(&_TreasuryV1.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_TreasuryV1 *TreasuryV1TransactorSession) Reward() (*types.Transaction, error) {
	return _TreasuryV1.Contract.Reward(&_TreasuryV1.TransactOpts)
}

// TreasuryV1ContributionIterator is returned from FilterContribution and is used to iterate over the raw logs and unpacked data for Contribution events raised by the TreasuryV1 contract.
type TreasuryV1ContributionIterator struct {
	Event *TreasuryV1Contribution // Event containing the contract specifics and raw log

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
func (it *TreasuryV1ContributionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryV1Contribution)
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
		it.Event = new(TreasuryV1Contribution)
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
func (it *TreasuryV1ContributionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryV1ContributionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryV1Contribution represents a Contribution event raised by the TreasuryV1 contract.
type TreasuryV1Contribution struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContribution is a free log retrieval operation binding the contract event 0x4d154d4aae216bed6d0926db77c00df2b57c6b5ba4eee05775de20facede3a7b.
//
// Solidity: event Contribution(address from, uint256 amount)
func (_TreasuryV1 *TreasuryV1Filterer) FilterContribution(opts *bind.FilterOpts) (*TreasuryV1ContributionIterator, error) {

	logs, sub, err := _TreasuryV1.contract.FilterLogs(opts, "Contribution")
	if err != nil {
		return nil, err
	}
	return &TreasuryV1ContributionIterator{contract: _TreasuryV1.contract, event: "Contribution", logs: logs, sub: sub}, nil
}

// WatchContribution is a free log subscription operation binding the contract event 0x4d154d4aae216bed6d0926db77c00df2b57c6b5ba4eee05775de20facede3a7b.
//
// Solidity: event Contribution(address from, uint256 amount)
func (_TreasuryV1 *TreasuryV1Filterer) WatchContribution(opts *bind.WatchOpts, sink chan<- *TreasuryV1Contribution) (event.Subscription, error) {

	logs, sub, err := _TreasuryV1.contract.WatchLogs(opts, "Contribution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryV1Contribution)
				if err := _TreasuryV1.contract.UnpackLog(event, "Contribution", log); err != nil {
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
