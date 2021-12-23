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

// IMasternodeRegistryV2ABI is the input ABI used to generate the binding from.
const IMasternodeRegistryV2ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"ipv4address\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"enode\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"Announced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"Deactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Denounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Invalidation\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ipv4address\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"enode\",\"type\":\"bytes32[2]\"}],\"name\":\"announce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"canHeartbeat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"can_heartbeat\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"canInvalidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"can_invalidate\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"active\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"active_collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max_of_all_times\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"denounce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerate\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"masternodes\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerateActive\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"masternodes\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"heartbeat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"info\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ipv4address\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"enode\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"announced_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"invalidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"onCollateralUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ownerInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ipv4address\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"enode\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"announced_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token_proxy\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury_proxy\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"validationTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IMasternodeRegistryV2 is an auto generated Go binding around an Ethereum contract.
type IMasternodeRegistryV2 struct {
	IMasternodeRegistryV2Caller     // Read-only binding to the contract
	IMasternodeRegistryV2Transactor // Write-only binding to the contract
	IMasternodeRegistryV2Filterer   // Log filterer for contract events
}

// IMasternodeRegistryV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IMasternodeRegistryV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeRegistryV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IMasternodeRegistryV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeRegistryV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMasternodeRegistryV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeRegistryV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMasternodeRegistryV2Session struct {
	Contract     *IMasternodeRegistryV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMasternodeRegistryV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMasternodeRegistryV2CallerSession struct {
	Contract *IMasternodeRegistryV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IMasternodeRegistryV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMasternodeRegistryV2TransactorSession struct {
	Contract     *IMasternodeRegistryV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IMasternodeRegistryV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IMasternodeRegistryV2Raw struct {
	Contract *IMasternodeRegistryV2 // Generic contract binding to access the raw methods on
}

// IMasternodeRegistryV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMasternodeRegistryV2CallerRaw struct {
	Contract *IMasternodeRegistryV2Caller // Generic read-only contract binding to access the raw methods on
}

// IMasternodeRegistryV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMasternodeRegistryV2TransactorRaw struct {
	Contract *IMasternodeRegistryV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIMasternodeRegistryV2 creates a new instance of IMasternodeRegistryV2, bound to a specific deployed contract.
func NewIMasternodeRegistryV2(address common.Address, backend bind.ContractBackend) (*IMasternodeRegistryV2, error) {
	contract, err := bindIMasternodeRegistryV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryV2{IMasternodeRegistryV2Caller: IMasternodeRegistryV2Caller{contract: contract}, IMasternodeRegistryV2Transactor: IMasternodeRegistryV2Transactor{contract: contract}, IMasternodeRegistryV2Filterer: IMasternodeRegistryV2Filterer{contract: contract}}, nil
}

// NewIMasternodeRegistryV2Caller creates a new read-only instance of IMasternodeRegistryV2, bound to a specific deployed contract.
func NewIMasternodeRegistryV2Caller(address common.Address, caller bind.ContractCaller) (*IMasternodeRegistryV2Caller, error) {
	contract, err := bindIMasternodeRegistryV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryV2Caller{contract: contract}, nil
}

// NewIMasternodeRegistryV2Transactor creates a new write-only instance of IMasternodeRegistryV2, bound to a specific deployed contract.
func NewIMasternodeRegistryV2Transactor(address common.Address, transactor bind.ContractTransactor) (*IMasternodeRegistryV2Transactor, error) {
	contract, err := bindIMasternodeRegistryV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryV2Transactor{contract: contract}, nil
}

// NewIMasternodeRegistryV2Filterer creates a new log filterer instance of IMasternodeRegistryV2, bound to a specific deployed contract.
func NewIMasternodeRegistryV2Filterer(address common.Address, filterer bind.ContractFilterer) (*IMasternodeRegistryV2Filterer, error) {
	contract, err := bindIMasternodeRegistryV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryV2Filterer{contract: contract}, nil
}

// bindIMasternodeRegistryV2 binds a generic wrapper to an already deployed contract.
func bindIMasternodeRegistryV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMasternodeRegistryV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMasternodeRegistryV2.Contract.IMasternodeRegistryV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.IMasternodeRegistryV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.IMasternodeRegistryV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMasternodeRegistryV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.contract.Transact(opts, method, params...)
}

// CanHeartbeat is a free data retrieval call binding the contract method 0xd9966aba.
//
// Solidity: function canHeartbeat(address masternode) constant returns(bool can_heartbeat)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) CanHeartbeat(opts *bind.CallOpts, masternode common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "canHeartbeat", masternode)
	return *ret0, err
}

// CanHeartbeat is a free data retrieval call binding the contract method 0xd9966aba.
//
// Solidity: function canHeartbeat(address masternode) constant returns(bool can_heartbeat)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) CanHeartbeat(masternode common.Address) (bool, error) {
	return _IMasternodeRegistryV2.Contract.CanHeartbeat(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// CanHeartbeat is a free data retrieval call binding the contract method 0xd9966aba.
//
// Solidity: function canHeartbeat(address masternode) constant returns(bool can_heartbeat)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) CanHeartbeat(masternode common.Address) (bool, error) {
	return _IMasternodeRegistryV2.Contract.CanHeartbeat(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// CanInvalidate is a free data retrieval call binding the contract method 0xd13ef5a5.
//
// Solidity: function canInvalidate(address masternode) constant returns(bool can_invalidate)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) CanInvalidate(opts *bind.CallOpts, masternode common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "canInvalidate", masternode)
	return *ret0, err
}

// CanInvalidate is a free data retrieval call binding the contract method 0xd13ef5a5.
//
// Solidity: function canInvalidate(address masternode) constant returns(bool can_invalidate)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) CanInvalidate(masternode common.Address) (bool, error) {
	return _IMasternodeRegistryV2.Contract.CanInvalidate(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// CanInvalidate is a free data retrieval call binding the contract method 0xd13ef5a5.
//
// Solidity: function canInvalidate(address masternode) constant returns(bool can_invalidate)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) CanInvalidate(masternode common.Address) (bool, error) {
	return _IMasternodeRegistryV2.Contract.CanInvalidate(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// CollateralLimits is a free data retrieval call binding the contract method 0xe2cb2195.
//
// Solidity: function collateralLimits() constant returns(uint256 min, uint256 max)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) CollateralLimits(opts *bind.CallOpts) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	ret := new(struct {
		Min *big.Int
		Max *big.Int
	})
	out := ret
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "collateralLimits")
	return *ret, err
}

// CollateralLimits is a free data retrieval call binding the contract method 0xe2cb2195.
//
// Solidity: function collateralLimits() constant returns(uint256 min, uint256 max)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) CollateralLimits() (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _IMasternodeRegistryV2.Contract.CollateralLimits(&_IMasternodeRegistryV2.CallOpts)
}

// CollateralLimits is a free data retrieval call binding the contract method 0xe2cb2195.
//
// Solidity: function collateralLimits() constant returns(uint256 min, uint256 max)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) CollateralLimits() (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _IMasternodeRegistryV2.Contract.CollateralLimits(&_IMasternodeRegistryV2.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 active_collateral, uint256 total_collateral, uint256 max_of_all_times)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) Count(opts *bind.CallOpts) (struct {
	Active           *big.Int
	Total            *big.Int
	ActiveCollateral *big.Int
	TotalCollateral  *big.Int
	MaxOfAllTimes    *big.Int
}, error) {
	ret := new(struct {
		Active           *big.Int
		Total            *big.Int
		ActiveCollateral *big.Int
		TotalCollateral  *big.Int
		MaxOfAllTimes    *big.Int
	})
	out := ret
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "count")
	return *ret, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 active_collateral, uint256 total_collateral, uint256 max_of_all_times)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) Count() (struct {
	Active           *big.Int
	Total            *big.Int
	ActiveCollateral *big.Int
	TotalCollateral  *big.Int
	MaxOfAllTimes    *big.Int
}, error) {
	return _IMasternodeRegistryV2.Contract.Count(&_IMasternodeRegistryV2.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 active_collateral, uint256 total_collateral, uint256 max_of_all_times)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) Count() (struct {
	Active           *big.Int
	Total            *big.Int
	ActiveCollateral *big.Int
	TotalCollateral  *big.Int
	MaxOfAllTimes    *big.Int
}, error) {
	return _IMasternodeRegistryV2.Contract.Count(&_IMasternodeRegistryV2.CallOpts)
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(address[] masternodes)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) Enumerate(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "enumerate")
	return *ret0, err
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(address[] masternodes)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) Enumerate() ([]common.Address, error) {
	return _IMasternodeRegistryV2.Contract.Enumerate(&_IMasternodeRegistryV2.CallOpts)
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(address[] masternodes)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) Enumerate() ([]common.Address, error) {
	return _IMasternodeRegistryV2.Contract.Enumerate(&_IMasternodeRegistryV2.CallOpts)
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(address[] masternodes)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) EnumerateActive(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "enumerateActive")
	return *ret0, err
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(address[] masternodes)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) EnumerateActive() ([]common.Address, error) {
	return _IMasternodeRegistryV2.Contract.EnumerateActive(&_IMasternodeRegistryV2.CallOpts)
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(address[] masternodes)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) EnumerateActive() ([]common.Address, error) {
	return _IMasternodeRegistryV2.Contract.EnumerateActive(&_IMasternodeRegistryV2.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address masternode) constant returns(address owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block, uint256 sw_features)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) Info(opts *bind.CallOpts, masternode common.Address) (struct {
	Owner          common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
	SwFeatures     *big.Int
}, error) {
	ret := new(struct {
		Owner          common.Address
		Ipv4address    uint32
		Enode          [2][32]byte
		Collateral     *big.Int
		AnnouncedBlock *big.Int
		SwFeatures     *big.Int
	})
	out := ret
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "info", masternode)
	return *ret, err
}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address masternode) constant returns(address owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block, uint256 sw_features)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) Info(masternode common.Address) (struct {
	Owner          common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
	SwFeatures     *big.Int
}, error) {
	return _IMasternodeRegistryV2.Contract.Info(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address masternode) constant returns(address owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block, uint256 sw_features)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) Info(masternode common.Address) (struct {
	Owner          common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
	SwFeatures     *big.Int
}, error) {
	return _IMasternodeRegistryV2.Contract.Info(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address masternode) constant returns(bool)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) IsActive(opts *bind.CallOpts, masternode common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "isActive", masternode)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address masternode) constant returns(bool)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) IsActive(masternode common.Address) (bool, error) {
	return _IMasternodeRegistryV2.Contract.IsActive(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address masternode) constant returns(bool)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) IsActive(masternode common.Address) (bool, error) {
	return _IMasternodeRegistryV2.Contract.IsActive(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// OwnerInfo is a free data retrieval call binding the contract method 0xb83e1605.
//
// Solidity: function ownerInfo(address owner) constant returns(address masternode, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block, uint256 sw_features)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) OwnerInfo(opts *bind.CallOpts, owner common.Address) (struct {
	Masternode     common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
	SwFeatures     *big.Int
}, error) {
	ret := new(struct {
		Masternode     common.Address
		Ipv4address    uint32
		Enode          [2][32]byte
		Collateral     *big.Int
		AnnouncedBlock *big.Int
		SwFeatures     *big.Int
	})
	out := ret
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "ownerInfo", owner)
	return *ret, err
}

// OwnerInfo is a free data retrieval call binding the contract method 0xb83e1605.
//
// Solidity: function ownerInfo(address owner) constant returns(address masternode, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block, uint256 sw_features)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) OwnerInfo(owner common.Address) (struct {
	Masternode     common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
	SwFeatures     *big.Int
}, error) {
	return _IMasternodeRegistryV2.Contract.OwnerInfo(&_IMasternodeRegistryV2.CallOpts, owner)
}

// OwnerInfo is a free data retrieval call binding the contract method 0xb83e1605.
//
// Solidity: function ownerInfo(address owner) constant returns(address masternode, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block, uint256 sw_features)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) OwnerInfo(owner common.Address) (struct {
	Masternode     common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
	SwFeatures     *big.Int
}, error) {
	return _IMasternodeRegistryV2.Contract.OwnerInfo(&_IMasternodeRegistryV2.CallOpts, owner)
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) TokenProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "token_proxy")
	return *ret0, err
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) TokenProxy() (common.Address, error) {
	return _IMasternodeRegistryV2.Contract.TokenProxy(&_IMasternodeRegistryV2.CallOpts)
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) TokenProxy() (common.Address, error) {
	return _IMasternodeRegistryV2.Contract.TokenProxy(&_IMasternodeRegistryV2.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) TreasuryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "treasury_proxy")
	return *ret0, err
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) TreasuryProxy() (common.Address, error) {
	return _IMasternodeRegistryV2.Contract.TreasuryProxy(&_IMasternodeRegistryV2.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) TreasuryProxy() (common.Address, error) {
	return _IMasternodeRegistryV2.Contract.TreasuryProxy(&_IMasternodeRegistryV2.CallOpts)
}

// ValidationTarget is a free data retrieval call binding the contract method 0xc3db74d6.
//
// Solidity: function validationTarget(address masternode) constant returns(address target)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Caller) ValidationTarget(opts *bind.CallOpts, masternode common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IMasternodeRegistryV2.contract.Call(opts, out, "validationTarget", masternode)
	return *ret0, err
}

// ValidationTarget is a free data retrieval call binding the contract method 0xc3db74d6.
//
// Solidity: function validationTarget(address masternode) constant returns(address target)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) ValidationTarget(masternode common.Address) (common.Address, error) {
	return _IMasternodeRegistryV2.Contract.ValidationTarget(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// ValidationTarget is a free data retrieval call binding the contract method 0xc3db74d6.
//
// Solidity: function validationTarget(address masternode) constant returns(address target)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2CallerSession) ValidationTarget(masternode common.Address) (common.Address, error) {
	return _IMasternodeRegistryV2.Contract.ValidationTarget(&_IMasternodeRegistryV2.CallOpts, masternode)
}

// Announce is a paid mutator transaction binding the contract method 0xd70d5c30.
//
// Solidity: function announce(address masternode, uint32 ipv4address, bytes32[2] enode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Transactor) Announce(opts *bind.TransactOpts, masternode common.Address, ipv4address uint32, enode [2][32]byte) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.contract.Transact(opts, "announce", masternode, ipv4address, enode)
}

// Announce is a paid mutator transaction binding the contract method 0xd70d5c30.
//
// Solidity: function announce(address masternode, uint32 ipv4address, bytes32[2] enode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) Announce(masternode common.Address, ipv4address uint32, enode [2][32]byte) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.Announce(&_IMasternodeRegistryV2.TransactOpts, masternode, ipv4address, enode)
}

// Announce is a paid mutator transaction binding the contract method 0xd70d5c30.
//
// Solidity: function announce(address masternode, uint32 ipv4address, bytes32[2] enode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2TransactorSession) Announce(masternode common.Address, ipv4address uint32, enode [2][32]byte) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.Announce(&_IMasternodeRegistryV2.TransactOpts, masternode, ipv4address, enode)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Transactor) Denounce(opts *bind.TransactOpts, masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.contract.Transact(opts, "denounce", masternode)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) Denounce(masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.Denounce(&_IMasternodeRegistryV2.TransactOpts, masternode)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2TransactorSession) Denounce(masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.Denounce(&_IMasternodeRegistryV2.TransactOpts, masternode)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Transactor) Heartbeat(opts *bind.TransactOpts, block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.contract.Transact(opts, "heartbeat", block_number, block_hash, sw_features)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) Heartbeat(block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.Heartbeat(&_IMasternodeRegistryV2.TransactOpts, block_number, block_hash, sw_features)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2TransactorSession) Heartbeat(block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.Heartbeat(&_IMasternodeRegistryV2.TransactOpts, block_number, block_hash, sw_features)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Transactor) Invalidate(opts *bind.TransactOpts, masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.contract.Transact(opts, "invalidate", masternode)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) Invalidate(masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.Invalidate(&_IMasternodeRegistryV2.TransactOpts, masternode)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2TransactorSession) Invalidate(masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.Invalidate(&_IMasternodeRegistryV2.TransactOpts, masternode)
}

// OnCollateralUpdate is a paid mutator transaction binding the contract method 0xcdc7d4ad.
//
// Solidity: function onCollateralUpdate(address owner) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Transactor) OnCollateralUpdate(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.contract.Transact(opts, "onCollateralUpdate", owner)
}

// OnCollateralUpdate is a paid mutator transaction binding the contract method 0xcdc7d4ad.
//
// Solidity: function onCollateralUpdate(address owner) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Session) OnCollateralUpdate(owner common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.OnCollateralUpdate(&_IMasternodeRegistryV2.TransactOpts, owner)
}

// OnCollateralUpdate is a paid mutator transaction binding the contract method 0xcdc7d4ad.
//
// Solidity: function onCollateralUpdate(address owner) returns()
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2TransactorSession) OnCollateralUpdate(owner common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistryV2.Contract.OnCollateralUpdate(&_IMasternodeRegistryV2.TransactOpts, owner)
}

// IMasternodeRegistryV2AnnouncedIterator is returned from FilterAnnounced and is used to iterate over the raw logs and unpacked data for Announced events raised by the IMasternodeRegistryV2 contract.
type IMasternodeRegistryV2AnnouncedIterator struct {
	Event *IMasternodeRegistryV2Announced // Event containing the contract specifics and raw log

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
func (it *IMasternodeRegistryV2AnnouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeRegistryV2Announced)
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
		it.Event = new(IMasternodeRegistryV2Announced)
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
func (it *IMasternodeRegistryV2AnnouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeRegistryV2AnnouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeRegistryV2Announced represents a Announced event raised by the IMasternodeRegistryV2 contract.
type IMasternodeRegistryV2Announced struct {
	Masternode  common.Address
	Owner       common.Address
	Ipv4address uint32
	Enode       [2][32]byte
	Collateral  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAnnounced is a free log retrieval operation binding the contract event 0x935a2f33570c4840d82856d75f5d0aafca32c5e6b31db5627552304a7dc82c09.
//
// Solidity: event Announced(address indexed masternode, address indexed owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Filterer) FilterAnnounced(opts *bind.FilterOpts, masternode []common.Address, owner []common.Address) (*IMasternodeRegistryV2AnnouncedIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IMasternodeRegistryV2.contract.FilterLogs(opts, "Announced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryV2AnnouncedIterator{contract: _IMasternodeRegistryV2.contract, event: "Announced", logs: logs, sub: sub}, nil
}

// WatchAnnounced is a free log subscription operation binding the contract event 0x935a2f33570c4840d82856d75f5d0aafca32c5e6b31db5627552304a7dc82c09.
//
// Solidity: event Announced(address indexed masternode, address indexed owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Filterer) WatchAnnounced(opts *bind.WatchOpts, sink chan<- *IMasternodeRegistryV2Announced, masternode []common.Address, owner []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IMasternodeRegistryV2.contract.WatchLogs(opts, "Announced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeRegistryV2Announced)
				if err := _IMasternodeRegistryV2.contract.UnpackLog(event, "Announced", log); err != nil {
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

// IMasternodeRegistryV2DeactivatedIterator is returned from FilterDeactivated and is used to iterate over the raw logs and unpacked data for Deactivated events raised by the IMasternodeRegistryV2 contract.
type IMasternodeRegistryV2DeactivatedIterator struct {
	Event *IMasternodeRegistryV2Deactivated // Event containing the contract specifics and raw log

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
func (it *IMasternodeRegistryV2DeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeRegistryV2Deactivated)
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
		it.Event = new(IMasternodeRegistryV2Deactivated)
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
func (it *IMasternodeRegistryV2DeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeRegistryV2DeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeRegistryV2Deactivated represents a Deactivated event raised by the IMasternodeRegistryV2 contract.
type IMasternodeRegistryV2Deactivated struct {
	Masternode common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeactivated is a free log retrieval operation binding the contract event 0x749cb6b4c510bc468cf6b9c2086d6f0a54d6b18e25d37bf3200e68eab0880c00.
//
// Solidity: event Deactivated(address indexed masternode)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Filterer) FilterDeactivated(opts *bind.FilterOpts, masternode []common.Address) (*IMasternodeRegistryV2DeactivatedIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}

	logs, sub, err := _IMasternodeRegistryV2.contract.FilterLogs(opts, "Deactivated", masternodeRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryV2DeactivatedIterator{contract: _IMasternodeRegistryV2.contract, event: "Deactivated", logs: logs, sub: sub}, nil
}

// WatchDeactivated is a free log subscription operation binding the contract event 0x749cb6b4c510bc468cf6b9c2086d6f0a54d6b18e25d37bf3200e68eab0880c00.
//
// Solidity: event Deactivated(address indexed masternode)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Filterer) WatchDeactivated(opts *bind.WatchOpts, sink chan<- *IMasternodeRegistryV2Deactivated, masternode []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}

	logs, sub, err := _IMasternodeRegistryV2.contract.WatchLogs(opts, "Deactivated", masternodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeRegistryV2Deactivated)
				if err := _IMasternodeRegistryV2.contract.UnpackLog(event, "Deactivated", log); err != nil {
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

// IMasternodeRegistryV2DenouncedIterator is returned from FilterDenounced and is used to iterate over the raw logs and unpacked data for Denounced events raised by the IMasternodeRegistryV2 contract.
type IMasternodeRegistryV2DenouncedIterator struct {
	Event *IMasternodeRegistryV2Denounced // Event containing the contract specifics and raw log

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
func (it *IMasternodeRegistryV2DenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeRegistryV2Denounced)
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
		it.Event = new(IMasternodeRegistryV2Denounced)
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
func (it *IMasternodeRegistryV2DenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeRegistryV2DenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeRegistryV2Denounced represents a Denounced event raised by the IMasternodeRegistryV2 contract.
type IMasternodeRegistryV2Denounced struct {
	Masternode common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDenounced is a free log retrieval operation binding the contract event 0x55faf8e51ab442f8d8510476317b2e313144c3db60adc284affef64140fe8552.
//
// Solidity: event Denounced(address indexed masternode, address indexed owner)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Filterer) FilterDenounced(opts *bind.FilterOpts, masternode []common.Address, owner []common.Address) (*IMasternodeRegistryV2DenouncedIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IMasternodeRegistryV2.contract.FilterLogs(opts, "Denounced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryV2DenouncedIterator{contract: _IMasternodeRegistryV2.contract, event: "Denounced", logs: logs, sub: sub}, nil
}

// WatchDenounced is a free log subscription operation binding the contract event 0x55faf8e51ab442f8d8510476317b2e313144c3db60adc284affef64140fe8552.
//
// Solidity: event Denounced(address indexed masternode, address indexed owner)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Filterer) WatchDenounced(opts *bind.WatchOpts, sink chan<- *IMasternodeRegistryV2Denounced, masternode []common.Address, owner []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IMasternodeRegistryV2.contract.WatchLogs(opts, "Denounced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeRegistryV2Denounced)
				if err := _IMasternodeRegistryV2.contract.UnpackLog(event, "Denounced", log); err != nil {
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

// IMasternodeRegistryV2InvalidationIterator is returned from FilterInvalidation and is used to iterate over the raw logs and unpacked data for Invalidation events raised by the IMasternodeRegistryV2 contract.
type IMasternodeRegistryV2InvalidationIterator struct {
	Event *IMasternodeRegistryV2Invalidation // Event containing the contract specifics and raw log

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
func (it *IMasternodeRegistryV2InvalidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeRegistryV2Invalidation)
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
		it.Event = new(IMasternodeRegistryV2Invalidation)
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
func (it *IMasternodeRegistryV2InvalidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeRegistryV2InvalidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeRegistryV2Invalidation represents a Invalidation event raised by the IMasternodeRegistryV2 contract.
type IMasternodeRegistryV2Invalidation struct {
	Masternode common.Address
	Validator  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidation is a free log retrieval operation binding the contract event 0x389d58799d7eca76264c556a007ffbc7c60caa4e3c8ea0564e791af3a1b9d331.
//
// Solidity: event Invalidation(address indexed masternode, address indexed validator)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Filterer) FilterInvalidation(opts *bind.FilterOpts, masternode []common.Address, validator []common.Address) (*IMasternodeRegistryV2InvalidationIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IMasternodeRegistryV2.contract.FilterLogs(opts, "Invalidation", masternodeRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryV2InvalidationIterator{contract: _IMasternodeRegistryV2.contract, event: "Invalidation", logs: logs, sub: sub}, nil
}

// WatchInvalidation is a free log subscription operation binding the contract event 0x389d58799d7eca76264c556a007ffbc7c60caa4e3c8ea0564e791af3a1b9d331.
//
// Solidity: event Invalidation(address indexed masternode, address indexed validator)
func (_IMasternodeRegistryV2 *IMasternodeRegistryV2Filterer) WatchInvalidation(opts *bind.WatchOpts, sink chan<- *IMasternodeRegistryV2Invalidation, masternode []common.Address, validator []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IMasternodeRegistryV2.contract.WatchLogs(opts, "Invalidation", masternodeRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeRegistryV2Invalidation)
				if err := _IMasternodeRegistryV2.contract.UnpackLog(event, "Invalidation", log); err != nil {
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
