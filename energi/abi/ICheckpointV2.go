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

// ICheckpointV2ABI is the input ABI used to generate the binding from.
const ICheckpointV2ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"canVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"info\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"since\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"sign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"signature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signatureBase\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"sigbase\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signatures\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"siglist\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ICheckpointV2 is an auto generated Go binding around an Ethereum contract.
type ICheckpointV2 struct {
	ICheckpointV2Caller     // Read-only binding to the contract
	ICheckpointV2Transactor // Write-only binding to the contract
	ICheckpointV2Filterer   // Log filterer for contract events
}

// ICheckpointV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ICheckpointV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ICheckpointV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICheckpointV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICheckpointV2Session struct {
	Contract     *ICheckpointV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICheckpointV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICheckpointV2CallerSession struct {
	Contract *ICheckpointV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ICheckpointV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICheckpointV2TransactorSession struct {
	Contract     *ICheckpointV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ICheckpointV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ICheckpointV2Raw struct {
	Contract *ICheckpointV2 // Generic contract binding to access the raw methods on
}

// ICheckpointV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICheckpointV2CallerRaw struct {
	Contract *ICheckpointV2Caller // Generic read-only contract binding to access the raw methods on
}

// ICheckpointV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICheckpointV2TransactorRaw struct {
	Contract *ICheckpointV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewICheckpointV2 creates a new instance of ICheckpointV2, bound to a specific deployed contract.
func NewICheckpointV2(address common.Address, backend bind.ContractBackend) (*ICheckpointV2, error) {
	contract, err := bindICheckpointV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICheckpointV2{ICheckpointV2Caller: ICheckpointV2Caller{contract: contract}, ICheckpointV2Transactor: ICheckpointV2Transactor{contract: contract}, ICheckpointV2Filterer: ICheckpointV2Filterer{contract: contract}}, nil
}

// NewICheckpointV2Caller creates a new read-only instance of ICheckpointV2, bound to a specific deployed contract.
func NewICheckpointV2Caller(address common.Address, caller bind.ContractCaller) (*ICheckpointV2Caller, error) {
	contract, err := bindICheckpointV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointV2Caller{contract: contract}, nil
}

// NewICheckpointV2Transactor creates a new write-only instance of ICheckpointV2, bound to a specific deployed contract.
func NewICheckpointV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ICheckpointV2Transactor, error) {
	contract, err := bindICheckpointV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointV2Transactor{contract: contract}, nil
}

// NewICheckpointV2Filterer creates a new log filterer instance of ICheckpointV2, bound to a specific deployed contract.
func NewICheckpointV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ICheckpointV2Filterer, error) {
	contract, err := bindICheckpointV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICheckpointV2Filterer{contract: contract}, nil
}

// bindICheckpointV2 binds a generic wrapper to an already deployed contract.
func bindICheckpointV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICheckpointV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointV2 *ICheckpointV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICheckpointV2.Contract.ICheckpointV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointV2 *ICheckpointV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointV2.Contract.ICheckpointV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointV2 *ICheckpointV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointV2.Contract.ICheckpointV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointV2 *ICheckpointV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICheckpointV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointV2 *ICheckpointV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointV2 *ICheckpointV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointV2.Contract.contract.Transact(opts, method, params...)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address masternode) constant returns(bool)
func (_ICheckpointV2 *ICheckpointV2Caller) CanVote(opts *bind.CallOpts, masternode common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ICheckpointV2.contract.Call(opts, out, "canVote", masternode)
	return *ret0, err
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address masternode) constant returns(bool)
func (_ICheckpointV2 *ICheckpointV2Session) CanVote(masternode common.Address) (bool, error) {
	return _ICheckpointV2.Contract.CanVote(&_ICheckpointV2.CallOpts, masternode)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address masternode) constant returns(bool)
func (_ICheckpointV2 *ICheckpointV2CallerSession) CanVote(masternode common.Address) (bool, error) {
	return _ICheckpointV2.Contract.CanVote(&_ICheckpointV2.CallOpts, masternode)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256 number, bytes32 hash, uint256 since)
func (_ICheckpointV2 *ICheckpointV2Caller) Info(opts *bind.CallOpts) (struct {
	Number *big.Int
	Hash   [32]byte
	Since  *big.Int
}, error) {
	ret := new(struct {
		Number *big.Int
		Hash   [32]byte
		Since  *big.Int
	})
	out := ret
	err := _ICheckpointV2.contract.Call(opts, out, "info")
	return *ret, err
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256 number, bytes32 hash, uint256 since)
func (_ICheckpointV2 *ICheckpointV2Session) Info() (struct {
	Number *big.Int
	Hash   [32]byte
	Since  *big.Int
}, error) {
	return _ICheckpointV2.Contract.Info(&_ICheckpointV2.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256 number, bytes32 hash, uint256 since)
func (_ICheckpointV2 *ICheckpointV2CallerSession) Info() (struct {
	Number *big.Int
	Hash   [32]byte
	Since  *big.Int
}, error) {
	return _ICheckpointV2.Contract.Info(&_ICheckpointV2.CallOpts)
}

// Signature is a free data retrieval call binding the contract method 0x2bbe2c88.
//
// Solidity: function signature(address masternode) constant returns(bytes)
func (_ICheckpointV2 *ICheckpointV2Caller) Signature(opts *bind.CallOpts, masternode common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICheckpointV2.contract.Call(opts, out, "signature", masternode)
	return *ret0, err
}

// Signature is a free data retrieval call binding the contract method 0x2bbe2c88.
//
// Solidity: function signature(address masternode) constant returns(bytes)
func (_ICheckpointV2 *ICheckpointV2Session) Signature(masternode common.Address) ([]byte, error) {
	return _ICheckpointV2.Contract.Signature(&_ICheckpointV2.CallOpts, masternode)
}

// Signature is a free data retrieval call binding the contract method 0x2bbe2c88.
//
// Solidity: function signature(address masternode) constant returns(bytes)
func (_ICheckpointV2 *ICheckpointV2CallerSession) Signature(masternode common.Address) ([]byte, error) {
	return _ICheckpointV2.Contract.Signature(&_ICheckpointV2.CallOpts, masternode)
}

// SignatureBase is a free data retrieval call binding the contract method 0x124321c4.
//
// Solidity: function signatureBase() constant returns(bytes32 sigbase)
func (_ICheckpointV2 *ICheckpointV2Caller) SignatureBase(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ICheckpointV2.contract.Call(opts, out, "signatureBase")
	return *ret0, err
}

// SignatureBase is a free data retrieval call binding the contract method 0x124321c4.
//
// Solidity: function signatureBase() constant returns(bytes32 sigbase)
func (_ICheckpointV2 *ICheckpointV2Session) SignatureBase() ([32]byte, error) {
	return _ICheckpointV2.Contract.SignatureBase(&_ICheckpointV2.CallOpts)
}

// SignatureBase is a free data retrieval call binding the contract method 0x124321c4.
//
// Solidity: function signatureBase() constant returns(bytes32 sigbase)
func (_ICheckpointV2 *ICheckpointV2CallerSession) SignatureBase() ([32]byte, error) {
	return _ICheckpointV2.Contract.SignatureBase(&_ICheckpointV2.CallOpts)
}

// Signatures is a free data retrieval call binding the contract method 0xf27959c7.
//
// Solidity: function signatures() constant returns(bytes[] siglist)
func (_ICheckpointV2 *ICheckpointV2Caller) Signatures(opts *bind.CallOpts) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _ICheckpointV2.contract.Call(opts, out, "signatures")
	return *ret0, err
}

// Signatures is a free data retrieval call binding the contract method 0xf27959c7.
//
// Solidity: function signatures() constant returns(bytes[] siglist)
func (_ICheckpointV2 *ICheckpointV2Session) Signatures() ([][]byte, error) {
	return _ICheckpointV2.Contract.Signatures(&_ICheckpointV2.CallOpts)
}

// Signatures is a free data retrieval call binding the contract method 0xf27959c7.
//
// Solidity: function signatures() constant returns(bytes[] siglist)
func (_ICheckpointV2 *ICheckpointV2CallerSession) Signatures() ([][]byte, error) {
	return _ICheckpointV2.Contract.Signatures(&_ICheckpointV2.CallOpts)
}

// Sign is a paid mutator transaction binding the contract method 0x76cd7cbc.
//
// Solidity: function sign(bytes signature) returns()
func (_ICheckpointV2 *ICheckpointV2Transactor) Sign(opts *bind.TransactOpts, signature []byte) (*types.Transaction, error) {
	return _ICheckpointV2.contract.Transact(opts, "sign", signature)
}

// Sign is a paid mutator transaction binding the contract method 0x76cd7cbc.
//
// Solidity: function sign(bytes signature) returns()
func (_ICheckpointV2 *ICheckpointV2Session) Sign(signature []byte) (*types.Transaction, error) {
	return _ICheckpointV2.Contract.Sign(&_ICheckpointV2.TransactOpts, signature)
}

// Sign is a paid mutator transaction binding the contract method 0x76cd7cbc.
//
// Solidity: function sign(bytes signature) returns()
func (_ICheckpointV2 *ICheckpointV2TransactorSession) Sign(signature []byte) (*types.Transaction, error) {
	return _ICheckpointV2.Contract.Sign(&_ICheckpointV2.TransactOpts, signature)
}
