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

// IMasternodeRegistryABI is the input ABI used to generate the binding from.
const IMasternodeRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"active\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"active_collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max_of_all_times\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"info\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ipv4address\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"enode\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"announced_block\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"invalidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token_proxy\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"heartbeat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury_proxy\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ownerInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ipv4address\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"enode\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"announced_block\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"validationTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"denounce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"onCollateralUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ipv4address\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"enode\",\"type\":\"bytes32[2]\"}],\"name\":\"announce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerateActive\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"masternodes\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerate\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"masternodes\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"ipv4address\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"enode\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"Announced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Denounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"Invalidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"Deactivated\",\"type\":\"event\"}]"

// IMasternodeRegistry is an auto generated Go binding around an Ethereum contract.
type IMasternodeRegistry struct {
	IMasternodeRegistryCaller     // Read-only binding to the contract
	IMasternodeRegistryTransactor // Write-only binding to the contract
	IMasternodeRegistryFilterer   // Log filterer for contract events
}

// IMasternodeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMasternodeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMasternodeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMasternodeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMasternodeRegistrySession struct {
	Contract     *IMasternodeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMasternodeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMasternodeRegistryCallerSession struct {
	Contract *IMasternodeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMasternodeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMasternodeRegistryTransactorSession struct {
	Contract     *IMasternodeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMasternodeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMasternodeRegistryRaw struct {
	Contract *IMasternodeRegistry // Generic contract binding to access the raw methods on
}

// IMasternodeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMasternodeRegistryCallerRaw struct {
	Contract *IMasternodeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IMasternodeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMasternodeRegistryTransactorRaw struct {
	Contract *IMasternodeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMasternodeRegistry creates a new instance of IMasternodeRegistry, bound to a specific deployed contract.
func NewIMasternodeRegistry(address common.Address, backend bind.ContractBackend) (*IMasternodeRegistry, error) {
	contract, err := bindIMasternodeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistry{IMasternodeRegistryCaller: IMasternodeRegistryCaller{contract: contract}, IMasternodeRegistryTransactor: IMasternodeRegistryTransactor{contract: contract}, IMasternodeRegistryFilterer: IMasternodeRegistryFilterer{contract: contract}}, nil
}

// NewIMasternodeRegistryCaller creates a new read-only instance of IMasternodeRegistry, bound to a specific deployed contract.
func NewIMasternodeRegistryCaller(address common.Address, caller bind.ContractCaller) (*IMasternodeRegistryCaller, error) {
	contract, err := bindIMasternodeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryCaller{contract: contract}, nil
}

// NewIMasternodeRegistryTransactor creates a new write-only instance of IMasternodeRegistry, bound to a specific deployed contract.
func NewIMasternodeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IMasternodeRegistryTransactor, error) {
	contract, err := bindIMasternodeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryTransactor{contract: contract}, nil
}

// NewIMasternodeRegistryFilterer creates a new log filterer instance of IMasternodeRegistry, bound to a specific deployed contract.
func NewIMasternodeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IMasternodeRegistryFilterer, error) {
	contract, err := bindIMasternodeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryFilterer{contract: contract}, nil
}

// bindIMasternodeRegistry binds a generic wrapper to an already deployed contract.
func bindIMasternodeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMasternodeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMasternodeRegistry *IMasternodeRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMasternodeRegistry.Contract.IMasternodeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMasternodeRegistry *IMasternodeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.IMasternodeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMasternodeRegistry *IMasternodeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.IMasternodeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMasternodeRegistry *IMasternodeRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMasternodeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMasternodeRegistry *IMasternodeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMasternodeRegistry *IMasternodeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 active_collateral, uint256 total_collateral, uint256 max_of_all_times)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) Count(opts *bind.CallOpts) (struct {
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
	err := _IMasternodeRegistry.contract.Call(opts, out, "count")
	return *ret, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 active_collateral, uint256 total_collateral, uint256 max_of_all_times)
func (_IMasternodeRegistry *IMasternodeRegistrySession) Count() (struct {
	Active           *big.Int
	Total            *big.Int
	ActiveCollateral *big.Int
	TotalCollateral  *big.Int
	MaxOfAllTimes    *big.Int
}, error) {
	return _IMasternodeRegistry.Contract.Count(&_IMasternodeRegistry.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256 active, uint256 total, uint256 active_collateral, uint256 total_collateral, uint256 max_of_all_times)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) Count() (struct {
	Active           *big.Int
	Total            *big.Int
	ActiveCollateral *big.Int
	TotalCollateral  *big.Int
	MaxOfAllTimes    *big.Int
}, error) {
	return _IMasternodeRegistry.Contract.Count(&_IMasternodeRegistry.CallOpts)
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(address[] masternodes)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) Enumerate(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IMasternodeRegistry.contract.Call(opts, out, "enumerate")
	return *ret0, err
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(address[] masternodes)
func (_IMasternodeRegistry *IMasternodeRegistrySession) Enumerate() ([]common.Address, error) {
	return _IMasternodeRegistry.Contract.Enumerate(&_IMasternodeRegistry.CallOpts)
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(address[] masternodes)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) Enumerate() ([]common.Address, error) {
	return _IMasternodeRegistry.Contract.Enumerate(&_IMasternodeRegistry.CallOpts)
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(address[] masternodes)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) EnumerateActive(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IMasternodeRegistry.contract.Call(opts, out, "enumerateActive")
	return *ret0, err
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(address[] masternodes)
func (_IMasternodeRegistry *IMasternodeRegistrySession) EnumerateActive() ([]common.Address, error) {
	return _IMasternodeRegistry.Contract.EnumerateActive(&_IMasternodeRegistry.CallOpts)
}

// EnumerateActive is a free data retrieval call binding the contract method 0xe1d6f43a.
//
// Solidity: function enumerateActive() constant returns(address[] masternodes)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) EnumerateActive() ([]common.Address, error) {
	return _IMasternodeRegistry.Contract.EnumerateActive(&_IMasternodeRegistry.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address masternode) constant returns(address owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) Info(opts *bind.CallOpts, masternode common.Address) (struct {
	Owner          common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
}, error) {
	ret := new(struct {
		Owner          common.Address
		Ipv4address    uint32
		Enode          [2][32]byte
		Collateral     *big.Int
		AnnouncedBlock *big.Int
	})
	out := ret
	err := _IMasternodeRegistry.contract.Call(opts, out, "info", masternode)
	return *ret, err
}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address masternode) constant returns(address owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block)
func (_IMasternodeRegistry *IMasternodeRegistrySession) Info(masternode common.Address) (struct {
	Owner          common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
}, error) {
	return _IMasternodeRegistry.Contract.Info(&_IMasternodeRegistry.CallOpts, masternode)
}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address masternode) constant returns(address owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) Info(masternode common.Address) (struct {
	Owner          common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
}, error) {
	return _IMasternodeRegistry.Contract.Info(&_IMasternodeRegistry.CallOpts, masternode)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address masternode) constant returns(bool)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) IsActive(opts *bind.CallOpts, masternode common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IMasternodeRegistry.contract.Call(opts, out, "isActive", masternode)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address masternode) constant returns(bool)
func (_IMasternodeRegistry *IMasternodeRegistrySession) IsActive(masternode common.Address) (bool, error) {
	return _IMasternodeRegistry.Contract.IsActive(&_IMasternodeRegistry.CallOpts, masternode)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address masternode) constant returns(bool)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) IsActive(masternode common.Address) (bool, error) {
	return _IMasternodeRegistry.Contract.IsActive(&_IMasternodeRegistry.CallOpts, masternode)
}

// OwnerInfo is a free data retrieval call binding the contract method 0xb83e1605.
//
// Solidity: function ownerInfo(address owner) constant returns(address masternode, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) OwnerInfo(opts *bind.CallOpts, owner common.Address) (struct {
	Masternode     common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
}, error) {
	ret := new(struct {
		Masternode     common.Address
		Ipv4address    uint32
		Enode          [2][32]byte
		Collateral     *big.Int
		AnnouncedBlock *big.Int
	})
	out := ret
	err := _IMasternodeRegistry.contract.Call(opts, out, "ownerInfo", owner)
	return *ret, err
}

// OwnerInfo is a free data retrieval call binding the contract method 0xb83e1605.
//
// Solidity: function ownerInfo(address owner) constant returns(address masternode, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block)
func (_IMasternodeRegistry *IMasternodeRegistrySession) OwnerInfo(owner common.Address) (struct {
	Masternode     common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
}, error) {
	return _IMasternodeRegistry.Contract.OwnerInfo(&_IMasternodeRegistry.CallOpts, owner)
}

// OwnerInfo is a free data retrieval call binding the contract method 0xb83e1605.
//
// Solidity: function ownerInfo(address owner) constant returns(address masternode, uint32 ipv4address, bytes32[2] enode, uint256 collateral, uint256 announced_block)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) OwnerInfo(owner common.Address) (struct {
	Masternode     common.Address
	Ipv4address    uint32
	Enode          [2][32]byte
	Collateral     *big.Int
	AnnouncedBlock *big.Int
}, error) {
	return _IMasternodeRegistry.Contract.OwnerInfo(&_IMasternodeRegistry.CallOpts, owner)
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) TokenProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IMasternodeRegistry.contract.Call(opts, out, "token_proxy")
	return *ret0, err
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_IMasternodeRegistry *IMasternodeRegistrySession) TokenProxy() (common.Address, error) {
	return _IMasternodeRegistry.Contract.TokenProxy(&_IMasternodeRegistry.CallOpts)
}

// TokenProxy is a free data retrieval call binding the contract method 0x84afd47f.
//
// Solidity: function token_proxy() constant returns(address)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) TokenProxy() (common.Address, error) {
	return _IMasternodeRegistry.Contract.TokenProxy(&_IMasternodeRegistry.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) TreasuryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IMasternodeRegistry.contract.Call(opts, out, "treasury_proxy")
	return *ret0, err
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_IMasternodeRegistry *IMasternodeRegistrySession) TreasuryProxy() (common.Address, error) {
	return _IMasternodeRegistry.Contract.TreasuryProxy(&_IMasternodeRegistry.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) TreasuryProxy() (common.Address, error) {
	return _IMasternodeRegistry.Contract.TreasuryProxy(&_IMasternodeRegistry.CallOpts)
}

// ValidationTarget is a free data retrieval call binding the contract method 0xc3db74d6.
//
// Solidity: function validationTarget(address masternode) constant returns(address target)
func (_IMasternodeRegistry *IMasternodeRegistryCaller) ValidationTarget(opts *bind.CallOpts, masternode common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IMasternodeRegistry.contract.Call(opts, out, "validationTarget", masternode)
	return *ret0, err
}

// ValidationTarget is a free data retrieval call binding the contract method 0xc3db74d6.
//
// Solidity: function validationTarget(address masternode) constant returns(address target)
func (_IMasternodeRegistry *IMasternodeRegistrySession) ValidationTarget(masternode common.Address) (common.Address, error) {
	return _IMasternodeRegistry.Contract.ValidationTarget(&_IMasternodeRegistry.CallOpts, masternode)
}

// ValidationTarget is a free data retrieval call binding the contract method 0xc3db74d6.
//
// Solidity: function validationTarget(address masternode) constant returns(address target)
func (_IMasternodeRegistry *IMasternodeRegistryCallerSession) ValidationTarget(masternode common.Address) (common.Address, error) {
	return _IMasternodeRegistry.Contract.ValidationTarget(&_IMasternodeRegistry.CallOpts, masternode)
}

// Announce is a paid mutator transaction binding the contract method 0xd70d5c30.
//
// Solidity: function announce(address masternode, uint32 ipv4address, bytes32[2] enode) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactor) Announce(opts *bind.TransactOpts, masternode common.Address, ipv4address uint32, enode [2][32]byte) (*types.Transaction, error) {
	return _IMasternodeRegistry.contract.Transact(opts, "announce", masternode, ipv4address, enode)
}

// Announce is a paid mutator transaction binding the contract method 0xd70d5c30.
//
// Solidity: function announce(address masternode, uint32 ipv4address, bytes32[2] enode) returns()
func (_IMasternodeRegistry *IMasternodeRegistrySession) Announce(masternode common.Address, ipv4address uint32, enode [2][32]byte) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.Announce(&_IMasternodeRegistry.TransactOpts, masternode, ipv4address, enode)
}

// Announce is a paid mutator transaction binding the contract method 0xd70d5c30.
//
// Solidity: function announce(address masternode, uint32 ipv4address, bytes32[2] enode) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactorSession) Announce(masternode common.Address, ipv4address uint32, enode [2][32]byte) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.Announce(&_IMasternodeRegistry.TransactOpts, masternode, ipv4address, enode)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactor) Denounce(opts *bind.TransactOpts, masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.contract.Transact(opts, "denounce", masternode)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_IMasternodeRegistry *IMasternodeRegistrySession) Denounce(masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.Denounce(&_IMasternodeRegistry.TransactOpts, masternode)
}

// Denounce is a paid mutator transaction binding the contract method 0xca0e551f.
//
// Solidity: function denounce(address masternode) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactorSession) Denounce(masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.Denounce(&_IMasternodeRegistry.TransactOpts, masternode)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactor) Heartbeat(opts *bind.TransactOpts, block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IMasternodeRegistry.contract.Transact(opts, "heartbeat", block_number, block_hash, sw_features)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_IMasternodeRegistry *IMasternodeRegistrySession) Heartbeat(block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.Heartbeat(&_IMasternodeRegistry.TransactOpts, block_number, block_hash, sw_features)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x91ceac96.
//
// Solidity: function heartbeat(uint256 block_number, bytes32 block_hash, uint256 sw_features) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactorSession) Heartbeat(block_number *big.Int, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.Heartbeat(&_IMasternodeRegistry.TransactOpts, block_number, block_hash, sw_features)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactor) Invalidate(opts *bind.TransactOpts, masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.contract.Transact(opts, "invalidate", masternode)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_IMasternodeRegistry *IMasternodeRegistrySession) Invalidate(masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.Invalidate(&_IMasternodeRegistry.TransactOpts, masternode)
}

// Invalidate is a paid mutator transaction binding the contract method 0x37a3931f.
//
// Solidity: function invalidate(address masternode) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactorSession) Invalidate(masternode common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.Invalidate(&_IMasternodeRegistry.TransactOpts, masternode)
}

// OnCollateralUpdate is a paid mutator transaction binding the contract method 0xcdc7d4ad.
//
// Solidity: function onCollateralUpdate(address owner) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactor) OnCollateralUpdate(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.contract.Transact(opts, "onCollateralUpdate", owner)
}

// OnCollateralUpdate is a paid mutator transaction binding the contract method 0xcdc7d4ad.
//
// Solidity: function onCollateralUpdate(address owner) returns()
func (_IMasternodeRegistry *IMasternodeRegistrySession) OnCollateralUpdate(owner common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.OnCollateralUpdate(&_IMasternodeRegistry.TransactOpts, owner)
}

// OnCollateralUpdate is a paid mutator transaction binding the contract method 0xcdc7d4ad.
//
// Solidity: function onCollateralUpdate(address owner) returns()
func (_IMasternodeRegistry *IMasternodeRegistryTransactorSession) OnCollateralUpdate(owner common.Address) (*types.Transaction, error) {
	return _IMasternodeRegistry.Contract.OnCollateralUpdate(&_IMasternodeRegistry.TransactOpts, owner)
}

// IMasternodeRegistryAnnouncedIterator is returned from FilterAnnounced and is used to iterate over the raw logs and unpacked data for Announced events raised by the IMasternodeRegistry contract.
type IMasternodeRegistryAnnouncedIterator struct {
	Event *IMasternodeRegistryAnnounced // Event containing the contract specifics and raw log

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
func (it *IMasternodeRegistryAnnouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeRegistryAnnounced)
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
		it.Event = new(IMasternodeRegistryAnnounced)
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
func (it *IMasternodeRegistryAnnouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeRegistryAnnouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeRegistryAnnounced represents a Announced event raised by the IMasternodeRegistry contract.
type IMasternodeRegistryAnnounced struct {
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
func (_IMasternodeRegistry *IMasternodeRegistryFilterer) FilterAnnounced(opts *bind.FilterOpts, masternode []common.Address, owner []common.Address) (*IMasternodeRegistryAnnouncedIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IMasternodeRegistry.contract.FilterLogs(opts, "Announced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryAnnouncedIterator{contract: _IMasternodeRegistry.contract, event: "Announced", logs: logs, sub: sub}, nil
}

// WatchAnnounced is a free log subscription operation binding the contract event 0x935a2f33570c4840d82856d75f5d0aafca32c5e6b31db5627552304a7dc82c09.
//
// Solidity: event Announced(address indexed masternode, address indexed owner, uint32 ipv4address, bytes32[2] enode, uint256 collateral)
func (_IMasternodeRegistry *IMasternodeRegistryFilterer) WatchAnnounced(opts *bind.WatchOpts, sink chan<- *IMasternodeRegistryAnnounced, masternode []common.Address, owner []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IMasternodeRegistry.contract.WatchLogs(opts, "Announced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeRegistryAnnounced)
				if err := _IMasternodeRegistry.contract.UnpackLog(event, "Announced", log); err != nil {
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

// IMasternodeRegistryDeactivatedIterator is returned from FilterDeactivated and is used to iterate over the raw logs and unpacked data for Deactivated events raised by the IMasternodeRegistry contract.
type IMasternodeRegistryDeactivatedIterator struct {
	Event *IMasternodeRegistryDeactivated // Event containing the contract specifics and raw log

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
func (it *IMasternodeRegistryDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeRegistryDeactivated)
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
		it.Event = new(IMasternodeRegistryDeactivated)
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
func (it *IMasternodeRegistryDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeRegistryDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeRegistryDeactivated represents a Deactivated event raised by the IMasternodeRegistry contract.
type IMasternodeRegistryDeactivated struct {
	Masternode common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeactivated is a free log retrieval operation binding the contract event 0x749cb6b4c510bc468cf6b9c2086d6f0a54d6b18e25d37bf3200e68eab0880c00.
//
// Solidity: event Deactivated(address indexed masternode)
func (_IMasternodeRegistry *IMasternodeRegistryFilterer) FilterDeactivated(opts *bind.FilterOpts, masternode []common.Address) (*IMasternodeRegistryDeactivatedIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}

	logs, sub, err := _IMasternodeRegistry.contract.FilterLogs(opts, "Deactivated", masternodeRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryDeactivatedIterator{contract: _IMasternodeRegistry.contract, event: "Deactivated", logs: logs, sub: sub}, nil
}

// WatchDeactivated is a free log subscription operation binding the contract event 0x749cb6b4c510bc468cf6b9c2086d6f0a54d6b18e25d37bf3200e68eab0880c00.
//
// Solidity: event Deactivated(address indexed masternode)
func (_IMasternodeRegistry *IMasternodeRegistryFilterer) WatchDeactivated(opts *bind.WatchOpts, sink chan<- *IMasternodeRegistryDeactivated, masternode []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}

	logs, sub, err := _IMasternodeRegistry.contract.WatchLogs(opts, "Deactivated", masternodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeRegistryDeactivated)
				if err := _IMasternodeRegistry.contract.UnpackLog(event, "Deactivated", log); err != nil {
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

// IMasternodeRegistryDenouncedIterator is returned from FilterDenounced and is used to iterate over the raw logs and unpacked data for Denounced events raised by the IMasternodeRegistry contract.
type IMasternodeRegistryDenouncedIterator struct {
	Event *IMasternodeRegistryDenounced // Event containing the contract specifics and raw log

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
func (it *IMasternodeRegistryDenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeRegistryDenounced)
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
		it.Event = new(IMasternodeRegistryDenounced)
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
func (it *IMasternodeRegistryDenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeRegistryDenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeRegistryDenounced represents a Denounced event raised by the IMasternodeRegistry contract.
type IMasternodeRegistryDenounced struct {
	Masternode common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDenounced is a free log retrieval operation binding the contract event 0x55faf8e51ab442f8d8510476317b2e313144c3db60adc284affef64140fe8552.
//
// Solidity: event Denounced(address indexed masternode, address indexed owner)
func (_IMasternodeRegistry *IMasternodeRegistryFilterer) FilterDenounced(opts *bind.FilterOpts, masternode []common.Address, owner []common.Address) (*IMasternodeRegistryDenouncedIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IMasternodeRegistry.contract.FilterLogs(opts, "Denounced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryDenouncedIterator{contract: _IMasternodeRegistry.contract, event: "Denounced", logs: logs, sub: sub}, nil
}

// WatchDenounced is a free log subscription operation binding the contract event 0x55faf8e51ab442f8d8510476317b2e313144c3db60adc284affef64140fe8552.
//
// Solidity: event Denounced(address indexed masternode, address indexed owner)
func (_IMasternodeRegistry *IMasternodeRegistryFilterer) WatchDenounced(opts *bind.WatchOpts, sink chan<- *IMasternodeRegistryDenounced, masternode []common.Address, owner []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IMasternodeRegistry.contract.WatchLogs(opts, "Denounced", masternodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeRegistryDenounced)
				if err := _IMasternodeRegistry.contract.UnpackLog(event, "Denounced", log); err != nil {
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

// IMasternodeRegistryInvalidationIterator is returned from FilterInvalidation and is used to iterate over the raw logs and unpacked data for Invalidation events raised by the IMasternodeRegistry contract.
type IMasternodeRegistryInvalidationIterator struct {
	Event *IMasternodeRegistryInvalidation // Event containing the contract specifics and raw log

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
func (it *IMasternodeRegistryInvalidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeRegistryInvalidation)
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
		it.Event = new(IMasternodeRegistryInvalidation)
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
func (it *IMasternodeRegistryInvalidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeRegistryInvalidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeRegistryInvalidation represents a Invalidation event raised by the IMasternodeRegistry contract.
type IMasternodeRegistryInvalidation struct {
	Masternode common.Address
	Validator  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidation is a free log retrieval operation binding the contract event 0x389d58799d7eca76264c556a007ffbc7c60caa4e3c8ea0564e791af3a1b9d331.
//
// Solidity: event Invalidation(address indexed masternode, address indexed validator)
func (_IMasternodeRegistry *IMasternodeRegistryFilterer) FilterInvalidation(opts *bind.FilterOpts, masternode []common.Address, validator []common.Address) (*IMasternodeRegistryInvalidationIterator, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IMasternodeRegistry.contract.FilterLogs(opts, "Invalidation", masternodeRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeRegistryInvalidationIterator{contract: _IMasternodeRegistry.contract, event: "Invalidation", logs: logs, sub: sub}, nil
}

// WatchInvalidation is a free log subscription operation binding the contract event 0x389d58799d7eca76264c556a007ffbc7c60caa4e3c8ea0564e791af3a1b9d331.
//
// Solidity: event Invalidation(address indexed masternode, address indexed validator)
func (_IMasternodeRegistry *IMasternodeRegistryFilterer) WatchInvalidation(opts *bind.WatchOpts, sink chan<- *IMasternodeRegistryInvalidation, masternode []common.Address, validator []common.Address) (event.Subscription, error) {

	var masternodeRule []interface{}
	for _, masternodeItem := range masternode {
		masternodeRule = append(masternodeRule, masternodeItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IMasternodeRegistry.contract.WatchLogs(opts, "Invalidation", masternodeRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeRegistryInvalidation)
				if err := _IMasternodeRegistry.contract.UnpackLog(event, "Invalidation", log); err != nil {
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
