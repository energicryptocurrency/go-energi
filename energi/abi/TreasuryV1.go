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
const TreasuryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"superblock_cycle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"v1storage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"isSuperblock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"},{\"name\":\"_superblock_cycle\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// TreasuryV1Bin is the compiled bytecode used for deploying new contracts.
const TreasuryV1Bin = `608060405234801561001057600080fd5b506040516107ad3803806107ad8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b0319166001600160a01b038416179055604051610063906100b3565b604051809103906000f08015801561007f573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b03929092169190911790556002819055806100ac57fe5b50506100c0565b6102288061058583390190565b6104b6806100cf6000396000f3fe60806040526004361061007a5760003560e01c80632d0593051161004e5780632d0593051461017c57806352782d86146101ba578063ce5494bb146101f8578063ec556889146102385761007a565b8062f55d9d146100e15780630ef34745146101235780631c4b774b1461014a578063228cb73314610174575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100ed57600080fd5b506101216004803603602081101561010457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661024d565b005b34801561012f57600080fd5b506101386102f5565b60408051918252519081900360200190f35b34801561015657600080fd5b506101386004803603602081101561016d57600080fd5b50356102fb565b610121610315565b34801561018857600080fd5b50610191610317565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156101c657600080fd5b506101e4600480360360208110156101dd57600080fd5b5035610333565b604080519115158252519081900360200190f35b34801561020457600080fd5b506101216004803603602081101561021b57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610348565b34801561024457600080fd5b506101916103d6565b60005473ffffffffffffffffffffffffffffffffffffffff1633146102d357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6102dc816103f2565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60025481565b6000811561031057506926f6a8f4e638030000005b919050565b565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6000600254828161034057fe5b061592915050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103ce57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6103d3815b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600154604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b15801561046657600080fd5b505af115801561047a573d6000803e3d6000fd5b505050505056fea265627a7a72305820a9f3d9795efdbd059dedba0d6b0be030c1f7931f622dd5ec7b8bdf9e674a669464736f6c634300050900326080604052600080546001600160a01b03191633179055610203806100256000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806313af40351461003b57806341c0e1b514610070575b600080fd5b61006e6004803603602081101561005157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610078565b005b61006e610145565b60005473ffffffffffffffffffffffffffffffffffffffff1633146100fe57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101cb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b33fffea265627a7a723058207b05c108059409986123a397489da27379264c3009c659a142ec87fa82338cbb64736f6c63430005090032`

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
