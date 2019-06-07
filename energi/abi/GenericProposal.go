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

// GenericProposalABI is the input ABI used to generate the binding from.
const GenericProposalABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAccepted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_payer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_quorum\",\"type\":\"uint8\"},{\"name\":\"_period\",\"type\":\"uint256\"},{\"name\":\"_fee_payer\",\"type\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// GenericProposalBin is the compiled bytecode used for deploying new contracts.
const GenericProposalBin = `608060405234801561001057600080fd5b506040516102523803806102528339818101604052608081101561003357600080fd5b5080516020820151604083015160609093015160005542016001556002805460ff909216740100000000000000000000000000000000000000000260ff60a01b196001600160a01b039094166001600160a01b031990931692909217929092161790556101ad806100a56000396000f3fe60806040526004361061005a5760003560e01c80635051a5ec116100435780635051a5ec146100ae578063c40a70f8146100d7578063ddca3f43146101155761005a565b80631703a0181461005c57806329dcb0cf14610087575b005b34801561006857600080fd5b5061007161012a565b6040805160ff9092168252519081900360200190f35b34801561009357600080fd5b5061009c61014b565b60408051918252519081900360200190f35b3480156100ba57600080fd5b506100c3610151565b604080519115158252519081900360200190f35b3480156100e357600080fd5b506100ec610156565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012157600080fd5b5061009c610172565b60025474010000000000000000000000000000000000000000900460ff1681565b60015481565b600090565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b6000548156fea265627a7a723058202c36fc05f8683895af5d0e9f063271e4ebb6dc1b1a07c1f4ec5bc996f2677f5864736f6c63430005090032`

// DeployGenericProposal deploys a new Ethereum contract, binding an instance of GenericProposal to it.
func DeployGenericProposal(auth *bind.TransactOpts, backend bind.ContractBackend, _quorum uint8, _period *big.Int, _fee_payer common.Address, _fee *big.Int) (common.Address, *types.Transaction, *GenericProposal, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericProposalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericProposalBin), backend, _quorum, _period, _fee_payer, _fee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericProposal{GenericProposalCaller: GenericProposalCaller{contract: contract}, GenericProposalTransactor: GenericProposalTransactor{contract: contract}, GenericProposalFilterer: GenericProposalFilterer{contract: contract}}, nil
}

// GenericProposal is an auto generated Go binding around an Ethereum contract.
type GenericProposal struct {
	GenericProposalCaller     // Read-only binding to the contract
	GenericProposalTransactor // Write-only binding to the contract
	GenericProposalFilterer   // Log filterer for contract events
}

// GenericProposalCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericProposalSession struct {
	Contract     *GenericProposal  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericProposalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericProposalCallerSession struct {
	Contract *GenericProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GenericProposalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericProposalTransactorSession struct {
	Contract     *GenericProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GenericProposalRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericProposalRaw struct {
	Contract *GenericProposal // Generic contract binding to access the raw methods on
}

// GenericProposalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericProposalCallerRaw struct {
	Contract *GenericProposalCaller // Generic read-only contract binding to access the raw methods on
}

// GenericProposalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericProposalTransactorRaw struct {
	Contract *GenericProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericProposal creates a new instance of GenericProposal, bound to a specific deployed contract.
func NewGenericProposal(address common.Address, backend bind.ContractBackend) (*GenericProposal, error) {
	contract, err := bindGenericProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericProposal{GenericProposalCaller: GenericProposalCaller{contract: contract}, GenericProposalTransactor: GenericProposalTransactor{contract: contract}, GenericProposalFilterer: GenericProposalFilterer{contract: contract}}, nil
}

// NewGenericProposalCaller creates a new read-only instance of GenericProposal, bound to a specific deployed contract.
func NewGenericProposalCaller(address common.Address, caller bind.ContractCaller) (*GenericProposalCaller, error) {
	contract, err := bindGenericProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericProposalCaller{contract: contract}, nil
}

// NewGenericProposalTransactor creates a new write-only instance of GenericProposal, bound to a specific deployed contract.
func NewGenericProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericProposalTransactor, error) {
	contract, err := bindGenericProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericProposalTransactor{contract: contract}, nil
}

// NewGenericProposalFilterer creates a new log filterer instance of GenericProposal, bound to a specific deployed contract.
func NewGenericProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericProposalFilterer, error) {
	contract, err := bindGenericProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericProposalFilterer{contract: contract}, nil
}

// bindGenericProposal binds a generic wrapper to an already deployed contract.
func bindGenericProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericProposalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericProposal *GenericProposalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericProposal.Contract.GenericProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericProposal *GenericProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposal.Contract.GenericProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericProposal *GenericProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericProposal.Contract.GenericProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericProposal *GenericProposalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericProposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericProposal *GenericProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericProposal *GenericProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericProposal.Contract.contract.Transact(opts, method, params...)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposal *GenericProposalCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposal.contract.Call(opts, out, "deadline")
	return *ret0, err
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposal *GenericProposalSession) Deadline() (*big.Int, error) {
	return _GenericProposal.Contract.Deadline(&_GenericProposal.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposal *GenericProposalCallerSession) Deadline() (*big.Int, error) {
	return _GenericProposal.Contract.Deadline(&_GenericProposal.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_GenericProposal *GenericProposalCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposal.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_GenericProposal *GenericProposalSession) Fee() (*big.Int, error) {
	return _GenericProposal.Contract.Fee(&_GenericProposal.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_GenericProposal *GenericProposalCallerSession) Fee() (*big.Int, error) {
	return _GenericProposal.Contract.Fee(&_GenericProposal.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposal *GenericProposalCaller) FeePayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericProposal.contract.Call(opts, out, "fee_payer")
	return *ret0, err
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposal *GenericProposalSession) FeePayer() (common.Address, error) {
	return _GenericProposal.Contract.FeePayer(&_GenericProposal.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposal *GenericProposalCallerSession) FeePayer() (common.Address, error) {
	return _GenericProposal.Contract.FeePayer(&_GenericProposal.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposal *GenericProposalCaller) IsAccepted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GenericProposal.contract.Call(opts, out, "isAccepted")
	return *ret0, err
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposal *GenericProposalSession) IsAccepted() (bool, error) {
	return _GenericProposal.Contract.IsAccepted(&_GenericProposal.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposal *GenericProposalCallerSession) IsAccepted() (bool, error) {
	return _GenericProposal.Contract.IsAccepted(&_GenericProposal.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint8)
func (_GenericProposal *GenericProposalCaller) Quorum(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GenericProposal.contract.Call(opts, out, "quorum")
	return *ret0, err
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint8)
func (_GenericProposal *GenericProposalSession) Quorum() (uint8, error) {
	return _GenericProposal.Contract.Quorum(&_GenericProposal.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint8)
func (_GenericProposal *GenericProposalCallerSession) Quorum() (uint8, error) {
	return _GenericProposal.Contract.Quorum(&_GenericProposal.CallOpts)
}
