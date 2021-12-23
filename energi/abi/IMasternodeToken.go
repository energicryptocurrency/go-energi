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

// IMasternodeTokenABI is the input ABI used to generate the binding from.
const IMasternodeTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_block\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositCollateral\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IMasternodeToken is an auto generated Go binding around an Ethereum contract.
type IMasternodeToken struct {
	IMasternodeTokenCaller     // Read-only binding to the contract
	IMasternodeTokenTransactor // Write-only binding to the contract
	IMasternodeTokenFilterer   // Log filterer for contract events
}

// IMasternodeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMasternodeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMasternodeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMasternodeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMasternodeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMasternodeTokenSession struct {
	Contract     *IMasternodeToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMasternodeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMasternodeTokenCallerSession struct {
	Contract *IMasternodeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IMasternodeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMasternodeTokenTransactorSession struct {
	Contract     *IMasternodeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IMasternodeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMasternodeTokenRaw struct {
	Contract *IMasternodeToken // Generic contract binding to access the raw methods on
}

// IMasternodeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMasternodeTokenCallerRaw struct {
	Contract *IMasternodeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IMasternodeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMasternodeTokenTransactorRaw struct {
	Contract *IMasternodeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMasternodeToken creates a new instance of IMasternodeToken, bound to a specific deployed contract.
func NewIMasternodeToken(address common.Address, backend bind.ContractBackend) (*IMasternodeToken, error) {
	contract, err := bindIMasternodeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMasternodeToken{IMasternodeTokenCaller: IMasternodeTokenCaller{contract: contract}, IMasternodeTokenTransactor: IMasternodeTokenTransactor{contract: contract}, IMasternodeTokenFilterer: IMasternodeTokenFilterer{contract: contract}}, nil
}

// NewIMasternodeTokenCaller creates a new read-only instance of IMasternodeToken, bound to a specific deployed contract.
func NewIMasternodeTokenCaller(address common.Address, caller bind.ContractCaller) (*IMasternodeTokenCaller, error) {
	contract, err := bindIMasternodeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMasternodeTokenCaller{contract: contract}, nil
}

// NewIMasternodeTokenTransactor creates a new write-only instance of IMasternodeToken, bound to a specific deployed contract.
func NewIMasternodeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IMasternodeTokenTransactor, error) {
	contract, err := bindIMasternodeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMasternodeTokenTransactor{contract: contract}, nil
}

// NewIMasternodeTokenFilterer creates a new log filterer instance of IMasternodeToken, bound to a specific deployed contract.
func NewIMasternodeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IMasternodeTokenFilterer, error) {
	contract, err := bindIMasternodeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMasternodeTokenFilterer{contract: contract}, nil
}

// bindIMasternodeToken binds a generic wrapper to an already deployed contract.
func bindIMasternodeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMasternodeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMasternodeToken *IMasternodeTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMasternodeToken.Contract.IMasternodeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMasternodeToken *IMasternodeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.IMasternodeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMasternodeToken *IMasternodeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.IMasternodeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMasternodeToken *IMasternodeTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMasternodeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMasternodeToken *IMasternodeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMasternodeToken *IMasternodeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_IMasternodeToken *IMasternodeTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IMasternodeToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_IMasternodeToken *IMasternodeTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IMasternodeToken.Contract.Allowance(&_IMasternodeToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_IMasternodeToken *IMasternodeTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _IMasternodeToken.Contract.Allowance(&_IMasternodeToken.CallOpts, _owner, _spender)
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_IMasternodeToken *IMasternodeTokenCaller) BalanceInfo(opts *bind.CallOpts, _tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	ret := new(struct {
		Balance   *big.Int
		LastBlock *big.Int
	})
	out := ret
	err := _IMasternodeToken.contract.Call(opts, out, "balanceInfo", _tokenOwner)
	return *ret, err
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_IMasternodeToken *IMasternodeTokenSession) BalanceInfo(_tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	return _IMasternodeToken.Contract.BalanceInfo(&_IMasternodeToken.CallOpts, _tokenOwner)
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_IMasternodeToken *IMasternodeTokenCallerSession) BalanceInfo(_tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	return _IMasternodeToken.Contract.BalanceInfo(&_IMasternodeToken.CallOpts, _tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_IMasternodeToken *IMasternodeTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IMasternodeToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_IMasternodeToken *IMasternodeTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _IMasternodeToken.Contract.BalanceOf(&_IMasternodeToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_IMasternodeToken *IMasternodeTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _IMasternodeToken.Contract.BalanceOf(&_IMasternodeToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IMasternodeToken *IMasternodeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IMasternodeToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IMasternodeToken *IMasternodeTokenSession) Decimals() (uint8, error) {
	return _IMasternodeToken.Contract.Decimals(&_IMasternodeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IMasternodeToken *IMasternodeTokenCallerSession) Decimals() (uint8, error) {
	return _IMasternodeToken.Contract.Decimals(&_IMasternodeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IMasternodeToken *IMasternodeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IMasternodeToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IMasternodeToken *IMasternodeTokenSession) Name() (string, error) {
	return _IMasternodeToken.Contract.Name(&_IMasternodeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IMasternodeToken *IMasternodeTokenCallerSession) Name() (string, error) {
	return _IMasternodeToken.Contract.Name(&_IMasternodeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IMasternodeToken *IMasternodeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IMasternodeToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IMasternodeToken *IMasternodeTokenSession) Symbol() (string, error) {
	return _IMasternodeToken.Contract.Symbol(&_IMasternodeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IMasternodeToken *IMasternodeTokenCallerSession) Symbol() (string, error) {
	return _IMasternodeToken.Contract.Symbol(&_IMasternodeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IMasternodeToken *IMasternodeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IMasternodeToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IMasternodeToken *IMasternodeTokenSession) TotalSupply() (*big.Int, error) {
	return _IMasternodeToken.Contract.TotalSupply(&_IMasternodeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IMasternodeToken *IMasternodeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IMasternodeToken.Contract.TotalSupply(&_IMasternodeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.Approve(&_IMasternodeToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.Approve(&_IMasternodeToken.TransactOpts, _spender, _value)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_IMasternodeToken *IMasternodeTokenTransactor) DepositCollateral(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMasternodeToken.contract.Transact(opts, "depositCollateral")
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_IMasternodeToken *IMasternodeTokenSession) DepositCollateral() (*types.Transaction, error) {
	return _IMasternodeToken.Contract.DepositCollateral(&_IMasternodeToken.TransactOpts)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_IMasternodeToken *IMasternodeTokenTransactorSession) DepositCollateral() (*types.Transaction, error) {
	return _IMasternodeToken.Contract.DepositCollateral(&_IMasternodeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.Transfer(&_IMasternodeToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.Transfer(&_IMasternodeToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.TransferFrom(&_IMasternodeToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_IMasternodeToken *IMasternodeTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.TransferFrom(&_IMasternodeToken.TransactOpts, _from, _to, _value)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_IMasternodeToken *IMasternodeTokenTransactor) WithdrawCollateral(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.contract.Transact(opts, "withdrawCollateral", _amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_IMasternodeToken *IMasternodeTokenSession) WithdrawCollateral(_amount *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.WithdrawCollateral(&_IMasternodeToken.TransactOpts, _amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_IMasternodeToken *IMasternodeTokenTransactorSession) WithdrawCollateral(_amount *big.Int) (*types.Transaction, error) {
	return _IMasternodeToken.Contract.WithdrawCollateral(&_IMasternodeToken.TransactOpts, _amount)
}

// IMasternodeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IMasternodeToken contract.
type IMasternodeTokenApprovalIterator struct {
	Event *IMasternodeTokenApproval // Event containing the contract specifics and raw log

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
func (it *IMasternodeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeTokenApproval)
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
		it.Event = new(IMasternodeTokenApproval)
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
func (it *IMasternodeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeTokenApproval represents a Approval event raised by the IMasternodeToken contract.
type IMasternodeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMasternodeToken *IMasternodeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IMasternodeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IMasternodeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeTokenApprovalIterator{contract: _IMasternodeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMasternodeToken *IMasternodeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IMasternodeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IMasternodeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeTokenApproval)
				if err := _IMasternodeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IMasternodeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IMasternodeToken contract.
type IMasternodeTokenTransferIterator struct {
	Event *IMasternodeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *IMasternodeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMasternodeTokenTransfer)
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
		it.Event = new(IMasternodeTokenTransfer)
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
func (it *IMasternodeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMasternodeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMasternodeTokenTransfer represents a Transfer event raised by the IMasternodeToken contract.
type IMasternodeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IMasternodeToken *IMasternodeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IMasternodeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMasternodeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IMasternodeTokenTransferIterator{contract: _IMasternodeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IMasternodeToken *IMasternodeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IMasternodeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMasternodeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMasternodeTokenTransfer)
				if err := _IMasternodeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
