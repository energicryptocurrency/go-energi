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

// ICheckpointABI is the input ABI used to generate the binding from.
const ICheckpointABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"signatureBase\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"sigbase\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masternode\",\"type\":\"address\"}],\"name\":\"signature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"info\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"since\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"sign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signatures\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"siglist\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ICheckpoint is an auto generated Go binding around an Ethereum contract.
type ICheckpoint struct {
	ICheckpointCaller     // Read-only binding to the contract
	ICheckpointTransactor // Write-only binding to the contract
	ICheckpointFilterer   // Log filterer for contract events
}

// ICheckpointCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICheckpointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICheckpointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICheckpointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICheckpointSession struct {
	Contract     *ICheckpoint      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICheckpointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICheckpointCallerSession struct {
	Contract *ICheckpointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ICheckpointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICheckpointTransactorSession struct {
	Contract     *ICheckpointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ICheckpointRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICheckpointRaw struct {
	Contract *ICheckpoint // Generic contract binding to access the raw methods on
}

// ICheckpointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICheckpointCallerRaw struct {
	Contract *ICheckpointCaller // Generic read-only contract binding to access the raw methods on
}

// ICheckpointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICheckpointTransactorRaw struct {
	Contract *ICheckpointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICheckpoint creates a new instance of ICheckpoint, bound to a specific deployed contract.
func NewICheckpoint(address common.Address, backend bind.ContractBackend) (*ICheckpoint, error) {
	contract, err := bindICheckpoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICheckpoint{ICheckpointCaller: ICheckpointCaller{contract: contract}, ICheckpointTransactor: ICheckpointTransactor{contract: contract}, ICheckpointFilterer: ICheckpointFilterer{contract: contract}}, nil
}

// NewICheckpointCaller creates a new read-only instance of ICheckpoint, bound to a specific deployed contract.
func NewICheckpointCaller(address common.Address, caller bind.ContractCaller) (*ICheckpointCaller, error) {
	contract, err := bindICheckpoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointCaller{contract: contract}, nil
}

// NewICheckpointTransactor creates a new write-only instance of ICheckpoint, bound to a specific deployed contract.
func NewICheckpointTransactor(address common.Address, transactor bind.ContractTransactor) (*ICheckpointTransactor, error) {
	contract, err := bindICheckpoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTransactor{contract: contract}, nil
}

// NewICheckpointFilterer creates a new log filterer instance of ICheckpoint, bound to a specific deployed contract.
func NewICheckpointFilterer(address common.Address, filterer bind.ContractFilterer) (*ICheckpointFilterer, error) {
	contract, err := bindICheckpoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICheckpointFilterer{contract: contract}, nil
}

// bindICheckpoint binds a generic wrapper to an already deployed contract.
func bindICheckpoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICheckpointABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpoint *ICheckpointRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICheckpoint.Contract.ICheckpointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpoint *ICheckpointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpoint.Contract.ICheckpointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpoint *ICheckpointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpoint.Contract.ICheckpointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpoint *ICheckpointCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICheckpoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpoint *ICheckpointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpoint *ICheckpointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpoint.Contract.contract.Transact(opts, method, params...)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256 number, bytes32 hash, uint256 since)
func (_ICheckpoint *ICheckpointCaller) Info(opts *bind.CallOpts) (struct {
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
	err := _ICheckpoint.contract.Call(opts, out, "info")
	return *ret, err
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256 number, bytes32 hash, uint256 since)
func (_ICheckpoint *ICheckpointSession) Info() (struct {
	Number *big.Int
	Hash   [32]byte
	Since  *big.Int
}, error) {
	return _ICheckpoint.Contract.Info(&_ICheckpoint.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256 number, bytes32 hash, uint256 since)
func (_ICheckpoint *ICheckpointCallerSession) Info() (struct {
	Number *big.Int
	Hash   [32]byte
	Since  *big.Int
}, error) {
	return _ICheckpoint.Contract.Info(&_ICheckpoint.CallOpts)
}

// Signature is a free data retrieval call binding the contract method 0x2bbe2c88.
//
// Solidity: function signature(address masternode) constant returns(bytes)
func (_ICheckpoint *ICheckpointCaller) Signature(opts *bind.CallOpts, masternode common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICheckpoint.contract.Call(opts, out, "signature", masternode)
	return *ret0, err
}

// Signature is a free data retrieval call binding the contract method 0x2bbe2c88.
//
// Solidity: function signature(address masternode) constant returns(bytes)
func (_ICheckpoint *ICheckpointSession) Signature(masternode common.Address) ([]byte, error) {
	return _ICheckpoint.Contract.Signature(&_ICheckpoint.CallOpts, masternode)
}

// Signature is a free data retrieval call binding the contract method 0x2bbe2c88.
//
// Solidity: function signature(address masternode) constant returns(bytes)
func (_ICheckpoint *ICheckpointCallerSession) Signature(masternode common.Address) ([]byte, error) {
	return _ICheckpoint.Contract.Signature(&_ICheckpoint.CallOpts, masternode)
}

// SignatureBase is a free data retrieval call binding the contract method 0x124321c4.
//
// Solidity: function signatureBase() constant returns(bytes32 sigbase)
func (_ICheckpoint *ICheckpointCaller) SignatureBase(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ICheckpoint.contract.Call(opts, out, "signatureBase")
	return *ret0, err
}

// SignatureBase is a free data retrieval call binding the contract method 0x124321c4.
//
// Solidity: function signatureBase() constant returns(bytes32 sigbase)
func (_ICheckpoint *ICheckpointSession) SignatureBase() ([32]byte, error) {
	return _ICheckpoint.Contract.SignatureBase(&_ICheckpoint.CallOpts)
}

// SignatureBase is a free data retrieval call binding the contract method 0x124321c4.
//
// Solidity: function signatureBase() constant returns(bytes32 sigbase)
func (_ICheckpoint *ICheckpointCallerSession) SignatureBase() ([32]byte, error) {
	return _ICheckpoint.Contract.SignatureBase(&_ICheckpoint.CallOpts)
}

// Signatures is a free data retrieval call binding the contract method 0xf27959c7.
//
// Solidity: function signatures() constant returns(bytes[] siglist)
func (_ICheckpoint *ICheckpointCaller) Signatures(opts *bind.CallOpts) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _ICheckpoint.contract.Call(opts, out, "signatures")
	return *ret0, err
}

// Signatures is a free data retrieval call binding the contract method 0xf27959c7.
//
// Solidity: function signatures() constant returns(bytes[] siglist)
func (_ICheckpoint *ICheckpointSession) Signatures() ([][]byte, error) {
	return _ICheckpoint.Contract.Signatures(&_ICheckpoint.CallOpts)
}

// Signatures is a free data retrieval call binding the contract method 0xf27959c7.
//
// Solidity: function signatures() constant returns(bytes[] siglist)
func (_ICheckpoint *ICheckpointCallerSession) Signatures() ([][]byte, error) {
	return _ICheckpoint.Contract.Signatures(&_ICheckpoint.CallOpts)
}

// Sign is a paid mutator transaction binding the contract method 0x76cd7cbc.
//
// Solidity: function sign(bytes signature) returns()
func (_ICheckpoint *ICheckpointTransactor) Sign(opts *bind.TransactOpts, signature []byte) (*types.Transaction, error) {
	return _ICheckpoint.contract.Transact(opts, "sign", signature)
}

// Sign is a paid mutator transaction binding the contract method 0x76cd7cbc.
//
// Solidity: function sign(bytes signature) returns()
func (_ICheckpoint *ICheckpointSession) Sign(signature []byte) (*types.Transaction, error) {
	return _ICheckpoint.Contract.Sign(&_ICheckpoint.TransactOpts, signature)
}

// Sign is a paid mutator transaction binding the contract method 0x76cd7cbc.
//
// Solidity: function sign(bytes signature) returns()
func (_ICheckpoint *ICheckpointTransactorSession) Sign(signature []byte) (*types.Transaction, error) {
	return _ICheckpoint.Contract.Sign(&_ICheckpoint.TransactOpts, signature)
}
