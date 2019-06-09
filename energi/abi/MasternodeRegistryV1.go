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
const MasternodeRegistryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"name\":\"active\",\"type\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint256\"},{\"name\":\"max_of_all_times\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_masternode\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"validate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"v1storage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"masternode\",\"type\":\"address\"},{\"name\":\"ipv4address\",\"type\":\"uint32\"}],\"name\":\"announce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mn_status\",\"outputs\":[{\"name\":\"sw_features\",\"type\":\"uint256\"},{\"name\":\"last_heartbeat\",\"type\":\"uint64\"},{\"name\":\"validations\",\"type\":\"uint32\"},{\"name\":\"is_active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mn_active\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"denounce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token_proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"isValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"block_number\",\"type\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"heartbeat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_payouts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury_proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mn_announced\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"denounce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mn_total_ever\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"},{\"name\":\"_token_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"Announced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Denounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Validation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"Heartbeat\",\"type\":\"event\"}]"

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

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 max_of_all_times)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) Count(opts *bind.CallOpts) (struct {
	Active        *big.Int
	Total         *big.Int
	MaxOfAllTimes *big.Int
}, error) {
	ret := new(struct {
		Active        *big.Int
		Total         *big.Int
		MaxOfAllTimes *big.Int
	})
	out := ret
	err := _MasternodeRegistryV1.contract.Call(opts, out, "count")
	return *ret, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 max_of_all_times)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Count() (struct {
	Active        *big.Int
	Total         *big.Int
	MaxOfAllTimes *big.Int
}, error) {
	return _MasternodeRegistryV1.Contract.Count(&_MasternodeRegistryV1.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 max_of_all_times)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) Count() (struct {
	Active        *big.Int
	Total         *big.Int
	MaxOfAllTimes *big.Int
}, error) {
	return _MasternodeRegistryV1.Contract.Count(&_MasternodeRegistryV1.CallOpts)
}

// CurrentMasternode is a free data retrieval call binding the contract method 0x1a26763c.
//
// Solidity: function current_masternode() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) CurrentMasternode(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "current_masternode")
	return *ret0, err
}

// CurrentMasternode is a free data retrieval call binding the contract method 0x1a26763c.
//
// Solidity: function current_masternode() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) CurrentMasternode() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.CurrentMasternode(&_MasternodeRegistryV1.CallOpts)
}

// CurrentMasternode is a free data retrieval call binding the contract method 0x1a26763c.
//
// Solidity: function current_masternode() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) CurrentMasternode() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.CurrentMasternode(&_MasternodeRegistryV1.CallOpts)
}

// CurrentPayouts is a free data retrieval call binding the contract method 0x9c5e8ae2.
//
// Solidity: function current_payouts() constant returns(uint8)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) CurrentPayouts(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "current_payouts")
	return *ret0, err
}

// CurrentPayouts is a free data retrieval call binding the contract method 0x9c5e8ae2.
//
// Solidity: function current_payouts() constant returns(uint8)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) CurrentPayouts() (uint8, error) {
	return _MasternodeRegistryV1.Contract.CurrentPayouts(&_MasternodeRegistryV1.CallOpts)
}

// CurrentPayouts is a free data retrieval call binding the contract method 0x9c5e8ae2.
//
// Solidity: function current_payouts() constant returns(uint8)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) CurrentPayouts() (uint8, error) {
	return _MasternodeRegistryV1.Contract.CurrentPayouts(&_MasternodeRegistryV1.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) GetReward(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "getReward", _blockNumber)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _MasternodeRegistryV1.Contract.GetReward(&_MasternodeRegistryV1.CallOpts, _blockNumber)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _MasternodeRegistryV1.Contract.GetReward(&_MasternodeRegistryV1.CallOpts, _blockNumber)
}

// IsValid is a free data retrieval call binding the contract method 0x8b1b925f.
//
// Solidity: function isValid(address masternode) constant returns(bool)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) IsValid(opts *bind.CallOpts, masternode common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "isValid", masternode)
	return *ret0, err
}

// IsValid is a free data retrieval call binding the contract method 0x8b1b925f.
//
// Solidity: function isValid(address masternode) constant returns(bool)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) IsValid(masternode common.Address) (bool, error) {
	return _MasternodeRegistryV1.Contract.IsValid(&_MasternodeRegistryV1.CallOpts, masternode)
}

// IsValid is a free data retrieval call binding the contract method 0x8b1b925f.
//
// Solidity: function isValid(address masternode) constant returns(bool)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) IsValid(masternode common.Address) (bool, error) {
	return _MasternodeRegistryV1.Contract.IsValid(&_MasternodeRegistryV1.CallOpts, masternode)
}

// MnActive is a free data retrieval call binding the contract method 0x521f284d.
//
// Solidity: function mn_active() constant returns(uint32)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) MnActive(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "mn_active")
	return *ret0, err
}

// MnActive is a free data retrieval call binding the contract method 0x521f284d.
//
// Solidity: function mn_active() constant returns(uint32)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) MnActive() (uint32, error) {
	return _MasternodeRegistryV1.Contract.MnActive(&_MasternodeRegistryV1.CallOpts)
}

// MnActive is a free data retrieval call binding the contract method 0x521f284d.
//
// Solidity: function mn_active() constant returns(uint32)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) MnActive() (uint32, error) {
	return _MasternodeRegistryV1.Contract.MnActive(&_MasternodeRegistryV1.CallOpts)
}

// MnAnnounced is a free data retrieval call binding the contract method 0xab00fdc4.
//
// Solidity: function mn_announced() constant returns(uint32)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) MnAnnounced(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "mn_announced")
	return *ret0, err
}

// MnAnnounced is a free data retrieval call binding the contract method 0xab00fdc4.
//
// Solidity: function mn_announced() constant returns(uint32)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) MnAnnounced() (uint32, error) {
	return _MasternodeRegistryV1.Contract.MnAnnounced(&_MasternodeRegistryV1.CallOpts)
}

// MnAnnounced is a free data retrieval call binding the contract method 0xab00fdc4.
//
// Solidity: function mn_announced() constant returns(uint32)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) MnAnnounced() (uint32, error) {
	return _MasternodeRegistryV1.Contract.MnAnnounced(&_MasternodeRegistryV1.CallOpts)
}

// MnStatus is a free data retrieval call binding the contract method 0x4d1b4dae.
//
// Solidity: function mn_status(address ) constant returns(uint256 sw_features, uint64 last_heartbeat, uint32 validations, bool is_active)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) MnStatus(opts *bind.CallOpts, arg0 common.Address) (struct {
	SwFeatures    *big.Int
	LastHeartbeat uint64
	Validations   uint32
	IsActive      bool
}, error) {
	ret := new(struct {
		SwFeatures    *big.Int
		LastHeartbeat uint64
		Validations   uint32
		IsActive      bool
	})
	out := ret
	err := _MasternodeRegistryV1.contract.Call(opts, out, "mn_status", arg0)
	return *ret, err
}

// MnStatus is a free data retrieval call binding the contract method 0x4d1b4dae.
//
// Solidity: function mn_status(address ) constant returns(uint256 sw_features, uint64 last_heartbeat, uint32 validations, bool is_active)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) MnStatus(arg0 common.Address) (struct {
	SwFeatures    *big.Int
	LastHeartbeat uint64
	Validations   uint32
	IsActive      bool
}, error) {
	return _MasternodeRegistryV1.Contract.MnStatus(&_MasternodeRegistryV1.CallOpts, arg0)
}

// MnStatus is a free data retrieval call binding the contract method 0x4d1b4dae.
//
// Solidity: function mn_status(address ) constant returns(uint256 sw_features, uint64 last_heartbeat, uint32 validations, bool is_active)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) MnStatus(arg0 common.Address) (struct {
	SwFeatures    *big.Int
	LastHeartbeat uint64
	Validations   uint32
	IsActive      bool
}, error) {
	return _MasternodeRegistryV1.Contract.MnStatus(&_MasternodeRegistryV1.CallOpts, arg0)
}

// MnTotalEver is a free data retrieval call binding the contract method 0xdb937a82.
//
// Solidity: function mn_total_ever() constant returns(uint64)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) MnTotalEver(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "mn_total_ever")
	return *ret0, err
}

// MnTotalEver is a free data retrieval call binding the contract method 0xdb937a82.
//
// Solidity: function mn_total_ever() constant returns(uint64)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) MnTotalEver() (uint64, error) {
	return _MasternodeRegistryV1.Contract.MnTotalEver(&_MasternodeRegistryV1.CallOpts)
}

// MnTotalEver is a free data retrieval call binding the contract method 0xdb937a82.
//
// Solidity: function mn_total_ever() constant returns(uint64)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) MnTotalEver() (uint64, error) {
	return _MasternodeRegistryV1.Contract.MnTotalEver(&_MasternodeRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Proxy() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.Proxy(&_MasternodeRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) Proxy() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.Proxy(&_MasternodeRegistryV1.CallOpts)
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) TokenProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "token_proxy")
	return *ret0, err
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) TokenProxy() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.TokenProxy(&_MasternodeRegistryV1.CallOpts)
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) TokenProxy() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.TokenProxy(&_MasternodeRegistryV1.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) TreasuryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "treasury_proxy")
	return *ret0, err
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) TreasuryProxy() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.TreasuryProxy(&_MasternodeRegistryV1.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) TreasuryProxy() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.TreasuryProxy(&_MasternodeRegistryV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Caller) V1storage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeRegistryV1.contract.Call(opts, out, "v1storage")
	return *ret0, err
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) V1storage() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.V1storage(&_MasternodeRegistryV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeRegistryV1 *MasternodeRegistryV1CallerSession) V1storage() (common.Address, error) {
	return _MasternodeRegistryV1.Contract.V1storage(&_MasternodeRegistryV1.CallOpts)
}

// Announce is a paid mutator transaction binding the contract method 0x3e3e4ac3.
//
// Solidity: function announce(address masternode, uint32 ipv4address) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Announce(opts *bind.TransactOpts, masternode common.Address, ipv4address uint32) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "announce", masternode, ipv4address)
}

// Announce is a paid mutator transaction binding the contract method 0x3e3e4ac3.
//
// Solidity: function announce(address masternode, uint32 ipv4address) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Announce(masternode common.Address, ipv4address uint32) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Announce(&_MasternodeRegistryV1.TransactOpts, masternode, ipv4address)
}

// Announce is a paid mutator transaction binding the contract method 0x3e3e4ac3.
//
// Solidity: function announce(address masternode, uint32 ipv4address) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Announce(masternode common.Address, ipv4address uint32) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Announce(&_MasternodeRegistryV1.TransactOpts, masternode, ipv4address)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Denounce(opts *bind.TransactOpts, masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "denounce", masternode)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Denounce(masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Denounce(&_MasternodeRegistryV1.TransactOpts, masternode)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Denounce(masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Denounce(&_MasternodeRegistryV1.TransactOpts, masternode)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Destroy(&_MasternodeRegistryV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Destroy(&_MasternodeRegistryV1.TransactOpts, _newImpl)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Heartbeat(opts *bind.TransactOpts, block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "heartbeat", block_number, block_hash, sw_features)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Heartbeat(block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Heartbeat(&_MasternodeRegistryV1.TransactOpts, block_number, block_hash, sw_features)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Heartbeat(block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Heartbeat(&_MasternodeRegistryV1.TransactOpts, block_number, block_hash, sw_features)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Migrate(&_MasternodeRegistryV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Migrate(&_MasternodeRegistryV1.TransactOpts, _oldImpl)
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

// Validate is a paid mutator transaction binding the contract method 0x207c64fb.
//
// Solidity: function validate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Transactor) Validate(opts *bind.TransactOpts, masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.contract.Transact(opts, "validate", masternode)
}

// Validate is a paid mutator transaction binding the contract method 0x207c64fb.
//
// Solidity: function validate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1Session) Validate(masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Validate(&_MasternodeRegistryV1.TransactOpts, masternode)
}

// Validate is a paid mutator transaction binding the contract method 0x207c64fb.
//
// Solidity: function validate(address masternode) returns()
func (_MasternodeRegistryV1 *MasternodeRegistryV1TransactorSession) Validate(masternode common.Address) (*types.Transaction, error) {
	return _MasternodeRegistryV1.Contract.Validate(&_MasternodeRegistryV1.TransactOpts, masternode)
}

// MasternodeRegistryV1AnnouncedIterator is returned from FilterAnnounced and is used to iterate over the raw logs and unpacked data for Announced events raised by the MasternodeRegistryV1 contract.
type MasternodeRegistryV1AnnouncedIterator struct {
	Event *MasternodeRegistryV1Announced // Event containing the contract specifics and raw log

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
func (it *MasternodeRegistryV1AnnouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasternodeRegistryV1Announced)
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
		it.Event = new(MasternodeRegistryV1Announced)
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
func (it *MasternodeRegistryV1AnnouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasternodeRegistryV1AnnouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasternodeRegistryV1Announced represents a Announced event raised by the MasternodeRegistryV1 contract.
type MasternodeRegistryV1Announced struct {
	Masternode common.Address
	Owner      common.Address
	Collateral *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAnnounced is a free log retrieval operation binding the contract event 0x5fc4086a510706fb705f95d9eab5c5aeb798ead7091e5817dae8c837a63ea43d.
//
// Solidity: event Announced(address indexed masternode, address indexed owner, uint256 collateral)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Filterer) FilterAnnounced(opts *bind.FilterOpts, masternode []common.Address, owner []common.Address) (*MasternodeRegistryV1AnnouncedIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MasternodeRegistryV1.contract.FilterLogs(opts, "Announced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1AnnouncedIterator{contract: _MasternodeRegistryV1.contract, event: "Announced", logs: logs, sub: sub}, nil
}

// WatchAnnounced is a free log subscription operation binding the contract event 0x5fc4086a510706fb705f95d9eab5c5aeb798ead7091e5817dae8c837a63ea43d.
//
// Solidity: event Announced(address indexed masternode, address indexed owner, uint256 collateral)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Filterer) WatchAnnounced(opts *bind.WatchOpts, sink chan<- *MasternodeRegistryV1Announced, masternode []common.Address, owner []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MasternodeRegistryV1.contract.WatchLogs(opts, "Announced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasternodeRegistryV1Announced)
				if err := _MasternodeRegistryV1.contract.UnpackLog(event, "Announced", log); err != nil {
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

// MasternodeRegistryV1DenouncedIterator is returned from FilterDenounced and is used to iterate over the raw logs and unpacked data for Denounced events raised by the MasternodeRegistryV1 contract.
type MasternodeRegistryV1DenouncedIterator struct {
	Event *MasternodeRegistryV1Denounced // Event containing the contract specifics and raw log

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
func (it *MasternodeRegistryV1DenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasternodeRegistryV1Denounced)
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
		it.Event = new(MasternodeRegistryV1Denounced)
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
func (it *MasternodeRegistryV1DenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasternodeRegistryV1DenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasternodeRegistryV1Denounced represents a Denounced event raised by the MasternodeRegistryV1 contract.
type MasternodeRegistryV1Denounced struct {
	Masternode common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDenounced is a free log retrieval operation binding the contract event 0x55faf8e51ab442f8d8510476317b2e313144c3db60adc284affef64140fe8552.
//
// Solidity: event Denounced(address indexed masternode, address indexed owner)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Filterer) FilterDenounced(opts *bind.FilterOpts, masternode []common.Address, owner []common.Address) (*MasternodeRegistryV1DenouncedIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MasternodeRegistryV1.contract.FilterLogs(opts, "Denounced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1DenouncedIterator{contract: _MasternodeRegistryV1.contract, event: "Denounced", logs: logs, sub: sub}, nil
}

// WatchDenounced is a free log subscription operation binding the contract event 0x55faf8e51ab442f8d8510476317b2e313144c3db60adc284affef64140fe8552.
//
// Solidity: event Denounced(address indexed masternode, address indexed owner)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Filterer) WatchDenounced(opts *bind.WatchOpts, sink chan<- *MasternodeRegistryV1Denounced, masternode []common.Address, owner []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MasternodeRegistryV1.contract.WatchLogs(opts, "Denounced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasternodeRegistryV1Denounced)
				if err := _MasternodeRegistryV1.contract.UnpackLog(event, "Denounced", log); err != nil {
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

// MasternodeRegistryV1HeartbeatIterator is returned from FilterHeartbeat and is used to iterate over the raw logs and unpacked data for Heartbeat events raised by the MasternodeRegistryV1 contract.
type MasternodeRegistryV1HeartbeatIterator struct {
	Event *MasternodeRegistryV1Heartbeat // Event containing the contract specifics and raw log

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
func (it *MasternodeRegistryV1HeartbeatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasternodeRegistryV1Heartbeat)
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
		it.Event = new(MasternodeRegistryV1Heartbeat)
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
func (it *MasternodeRegistryV1HeartbeatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasternodeRegistryV1HeartbeatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasternodeRegistryV1Heartbeat represents a Heartbeat event raised by the MasternodeRegistryV1 contract.
type MasternodeRegistryV1Heartbeat struct {
	Masternode common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHeartbeat is a free log retrieval operation binding the contract event 0x76fd25e9b1ccb2a2eb85da234dd15c82d9eec18877d3e6fc916eb7330fe04a64.
//
// Solidity: event Heartbeat(address indexed masternode)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Filterer) FilterHeartbeat(opts *bind.FilterOpts, masternode []common.Address) (*MasternodeRegistryV1HeartbeatIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}

	logs, sub, err := _MasternodeRegistryV1.contract.FilterLogs(opts, "Heartbeat", masternodeRule)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1HeartbeatIterator{contract: _MasternodeRegistryV1.contract, event: "Heartbeat", logs: logs, sub: sub}, nil
}

// WatchHeartbeat is a free log subscription operation binding the contract event 0x76fd25e9b1ccb2a2eb85da234dd15c82d9eec18877d3e6fc916eb7330fe04a64.
//
// Solidity: event Heartbeat(address indexed masternode)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Filterer) WatchHeartbeat(opts *bind.WatchOpts, sink chan<- *MasternodeRegistryV1Heartbeat, masternode []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}

	logs, sub, err := _MasternodeRegistryV1.contract.WatchLogs(opts, "Heartbeat", masternodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasternodeRegistryV1Heartbeat)
				if err := _MasternodeRegistryV1.contract.UnpackLog(event, "Heartbeat", log); err != nil {
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

// MasternodeRegistryV1ValidationIterator is returned from FilterValidation and is used to iterate over the raw logs and unpacked data for Validation events raised by the MasternodeRegistryV1 contract.
type MasternodeRegistryV1ValidationIterator struct {
	Event *MasternodeRegistryV1Validation // Event containing the contract specifics and raw log

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
func (it *MasternodeRegistryV1ValidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasternodeRegistryV1Validation)
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
		it.Event = new(MasternodeRegistryV1Validation)
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
func (it *MasternodeRegistryV1ValidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasternodeRegistryV1ValidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasternodeRegistryV1Validation represents a Validation event raised by the MasternodeRegistryV1 contract.
type MasternodeRegistryV1Validation struct {
	Masternode common.Address
	Validator  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidation is a free log retrieval operation binding the contract event 0xecf3d22a6c5bee8410ef008f28f974ab2b99754941406d83f2e79834bab78b6d.
//
// Solidity: event Validation(address indexed masternode, address validator)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Filterer) FilterValidation(opts *bind.FilterOpts, masternode []common.Address) (*MasternodeRegistryV1ValidationIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}

	logs, sub, err := _MasternodeRegistryV1.contract.FilterLogs(opts, "Validation", masternodeRule)
	if err != nil {
		return nil, err
	}
	return &MasternodeRegistryV1ValidationIterator{contract: _MasternodeRegistryV1.contract, event: "Validation", logs: logs, sub: sub}, nil
}

// WatchValidation is a free log subscription operation binding the contract event 0xecf3d22a6c5bee8410ef008f28f974ab2b99754941406d83f2e79834bab78b6d.
//
// Solidity: event Validation(address indexed masternode, address validator)
func (_MasternodeRegistryV1 *MasternodeRegistryV1Filterer) WatchValidation(opts *bind.WatchOpts, sink chan<- *MasternodeRegistryV1Validation, masternode []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}

	logs, sub, err := _MasternodeRegistryV1.contract.WatchLogs(opts, "Validation", masternodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasternodeRegistryV1Validation)
				if err := _MasternodeRegistryV1.contract.UnpackLog(event, "Validation", log); err != nil {
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
