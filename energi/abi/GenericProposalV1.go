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

// GenericProposalV1ABI is the input ABI used to generate the binding from.
const GenericProposalV1ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAccepted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_payer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_quorum\",\"type\":\"uint8\"},{\"name\":\"_period\",\"type\":\"uint256\"},{\"name\":\"_fee_payer\",\"type\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// GenericProposalV1Bin is the compiled bytecode used for deploying new contracts.
const GenericProposalV1Bin = `608060405234801561001057600080fd5b506040516102523803806102528339818101604052608081101561003357600080fd5b5080516020820151604083015160609093015160005542016001556002805460ff909216740100000000000000000000000000000000000000000260ff60a01b196001600160a01b039094166001600160a01b031990931692909217929092161790556101ad806100a56000396000f3fe60806040526004361061005a5760003560e01c80635051a5ec116100435780635051a5ec146100ae578063c40a70f8146100d7578063ddca3f43146101155761005a565b80631703a0181461005c57806329dcb0cf14610087575b005b34801561006857600080fd5b5061007161012a565b6040805160ff9092168252519081900360200190f35b34801561009357600080fd5b5061009c61014b565b60408051918252519081900360200190f35b3480156100ba57600080fd5b506100c3610151565b604080519115158252519081900360200190f35b3480156100e357600080fd5b506100ec610156565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012157600080fd5b5061009c610172565b60025474010000000000000000000000000000000000000000900460ff1681565b60015481565b600090565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b6000548156fea265627a7a72305820e6532d93d98857256a35cec991392469bcd9c34f05436f686b0eaa7827be9d3f64736f6c63430005090032`

// DeployGenericProposalV1 deploys a new Ethereum contract, binding an instance of GenericProposalV1 to it.
func DeployGenericProposalV1(auth *bind.TransactOpts, backend bind.ContractBackend, _quorum uint8, _period *big.Int, _fee_payer common.Address, _fee *big.Int) (common.Address, *types.Transaction, *GenericProposalV1, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericProposalV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericProposalV1Bin), backend, _quorum, _period, _fee_payer, _fee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericProposalV1{GenericProposalV1Caller: GenericProposalV1Caller{contract: contract}, GenericProposalV1Transactor: GenericProposalV1Transactor{contract: contract}, GenericProposalV1Filterer: GenericProposalV1Filterer{contract: contract}}, nil
}

// GenericProposalV1 is an auto generated Go binding around an Ethereum contract.
type GenericProposalV1 struct {
	GenericProposalV1Caller     // Read-only binding to the contract
	GenericProposalV1Transactor // Write-only binding to the contract
	GenericProposalV1Filterer   // Log filterer for contract events
}

// GenericProposalV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type GenericProposalV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericProposalV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericProposalV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericProposalV1Session struct {
	Contract     *GenericProposalV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GenericProposalV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericProposalV1CallerSession struct {
	Contract *GenericProposalV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GenericProposalV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericProposalV1TransactorSession struct {
	Contract     *GenericProposalV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GenericProposalV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type GenericProposalV1Raw struct {
	Contract *GenericProposalV1 // Generic contract binding to access the raw methods on
}

// GenericProposalV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericProposalV1CallerRaw struct {
	Contract *GenericProposalV1Caller // Generic read-only contract binding to access the raw methods on
}

// GenericProposalV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericProposalV1TransactorRaw struct {
	Contract *GenericProposalV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericProposalV1 creates a new instance of GenericProposalV1, bound to a specific deployed contract.
func NewGenericProposalV1(address common.Address, backend bind.ContractBackend) (*GenericProposalV1, error) {
	contract, err := bindGenericProposalV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericProposalV1{GenericProposalV1Caller: GenericProposalV1Caller{contract: contract}, GenericProposalV1Transactor: GenericProposalV1Transactor{contract: contract}, GenericProposalV1Filterer: GenericProposalV1Filterer{contract: contract}}, nil
}

// NewGenericProposalV1Caller creates a new read-only instance of GenericProposalV1, bound to a specific deployed contract.
func NewGenericProposalV1Caller(address common.Address, caller bind.ContractCaller) (*GenericProposalV1Caller, error) {
	contract, err := bindGenericProposalV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericProposalV1Caller{contract: contract}, nil
}

// NewGenericProposalV1Transactor creates a new write-only instance of GenericProposalV1, bound to a specific deployed contract.
func NewGenericProposalV1Transactor(address common.Address, transactor bind.ContractTransactor) (*GenericProposalV1Transactor, error) {
	contract, err := bindGenericProposalV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericProposalV1Transactor{contract: contract}, nil
}

// NewGenericProposalV1Filterer creates a new log filterer instance of GenericProposalV1, bound to a specific deployed contract.
func NewGenericProposalV1Filterer(address common.Address, filterer bind.ContractFilterer) (*GenericProposalV1Filterer, error) {
	contract, err := bindGenericProposalV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericProposalV1Filterer{contract: contract}, nil
}

// bindGenericProposalV1 binds a generic wrapper to an already deployed contract.
func bindGenericProposalV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericProposalV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericProposalV1 *GenericProposalV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericProposalV1.Contract.GenericProposalV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericProposalV1 *GenericProposalV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV1.Contract.GenericProposalV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericProposalV1 *GenericProposalV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericProposalV1.Contract.GenericProposalV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericProposalV1 *GenericProposalV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericProposalV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericProposalV1 *GenericProposalV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericProposalV1 *GenericProposalV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericProposalV1.Contract.contract.Transact(opts, method, params...)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposalV1 *GenericProposalV1Caller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV1.contract.Call(opts, out, "deadline")
	return *ret0, err
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposalV1 *GenericProposalV1Session) Deadline() (*big.Int, error) {
	return _GenericProposalV1.Contract.Deadline(&_GenericProposalV1.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposalV1 *GenericProposalV1CallerSession) Deadline() (*big.Int, error) {
	return _GenericProposalV1.Contract.Deadline(&_GenericProposalV1.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_GenericProposalV1 *GenericProposalV1Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV1.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_GenericProposalV1 *GenericProposalV1Session) Fee() (*big.Int, error) {
	return _GenericProposalV1.Contract.Fee(&_GenericProposalV1.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_GenericProposalV1 *GenericProposalV1CallerSession) Fee() (*big.Int, error) {
	return _GenericProposalV1.Contract.Fee(&_GenericProposalV1.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposalV1 *GenericProposalV1Caller) FeePayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericProposalV1.contract.Call(opts, out, "fee_payer")
	return *ret0, err
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposalV1 *GenericProposalV1Session) FeePayer() (common.Address, error) {
	return _GenericProposalV1.Contract.FeePayer(&_GenericProposalV1.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposalV1 *GenericProposalV1CallerSession) FeePayer() (common.Address, error) {
	return _GenericProposalV1.Contract.FeePayer(&_GenericProposalV1.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposalV1 *GenericProposalV1Caller) IsAccepted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GenericProposalV1.contract.Call(opts, out, "isAccepted")
	return *ret0, err
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposalV1 *GenericProposalV1Session) IsAccepted() (bool, error) {
	return _GenericProposalV1.Contract.IsAccepted(&_GenericProposalV1.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposalV1 *GenericProposalV1CallerSession) IsAccepted() (bool, error) {
	return _GenericProposalV1.Contract.IsAccepted(&_GenericProposalV1.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint8)
func (_GenericProposalV1 *GenericProposalV1Caller) Quorum(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GenericProposalV1.contract.Call(opts, out, "quorum")
	return *ret0, err
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint8)
func (_GenericProposalV1 *GenericProposalV1Session) Quorum() (uint8, error) {
	return _GenericProposalV1.Contract.Quorum(&_GenericProposalV1.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() constant returns(uint8)
func (_GenericProposalV1 *GenericProposalV1CallerSession) Quorum() (uint8, error) {
	return _GenericProposalV1.Contract.Quorum(&_GenericProposalV1.CallOpts)
}
