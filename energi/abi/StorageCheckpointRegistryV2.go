// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	ethereum "energi.world/core/gen3"
	"energi.world/core/gen3/accounts/abi"
	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/event"
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

// StorageCheckpointRegistryV2ABI is the input ABI used to generate the binding from.
const StorageCheckpointRegistryV2ABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"contractICheckpoint\",\"name\":\"cp\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"contractICheckpoint\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"listCheckpoints\",\"outputs\":[{\"internalType\":\"contractICheckpoint[]\",\"name\":\"res\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractICheckpoint\",\"name\":\"cp\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StorageCheckpointRegistryV2Bin is the compiled bytecode used for deploying new contracts.
const StorageCheckpointRegistryV2Bin = `6080604052600080546001600160a01b03191633179055610c7a806100256000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806341c0e1b51161005b57806341c0e1b5146100d3578063a4ece52c146100db578063b8a24252146100e3578063d9592ead146101035761007d565b80630a3b0a4f1461008257806313af40351461009757806329092d0e146100aa575b600080fd5b610095610090366004610a3d565b610118565b005b6100956100a5366004610a3d565b61030a565b6100bd6100b8366004610a3d565b6103a2565b6040516100ca9190610ba4565b60405180910390f35b610095610726565b61009561077a565b6100f66100f1366004610a63565b6108c1565b6040516100ca9190610bb2565b61010b6108e9565b6040516100ca9190610b8c565b60005473ffffffffffffffffffffffffffffffffffffffff163314610172576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b60405180910390fd5b6002546fffffffffffffffffffffffffffffffff70010000000000000000000000000000000090910416600a141561026e57600280546fffffffffffffffffffffffffffffffff908116600090815260016020819052604080832080547fffffffffffffffffffffffff000000000000000000000000000000000000000090811690915585547001000000000000000000000000000000008104861690861601851684529220805490921673ffffffffffffffffffffffffffffffffffffffff86161790915582547fffffffffffffffffffffffffffffffff000000000000000000000000000000008116908316909101909116179055610307565b600280546fffffffffffffffffffffffffffffffff8082167001000000000000000000000000000000009283900482160181166000908152600160208190526040909120805473ffffffffffffffffffffffffffffffffffffffff87167fffffffffffffffffffffffff000000000000000000000000000000000000000090911617905583548381048316909101821690920291161790555b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461035b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000805473ffffffffffffffffffffffffffffffffffffffff1633146103f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b60008091506000808473ffffffffffffffffffffffffffffffffffffffff1663370158ea6040518163ffffffff1660e01b815260040160606040518083038186803b15801561044257600080fd5b505afa158015610456573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061047a9190810190610a81565b506002546fffffffffffffffffffffffffffffffff16945090925090505b6002546fffffffffffffffffffffffffffffffff808216700100000000000000000000000000000000909204811691909101168310156105b1576000838152600160205260408082205481517f370158ea0000000000000000000000000000000000000000000000000000000081529151839273ffffffffffffffffffffffffffffffffffffffff9092169163370158ea916004808301926060929190829003018186803b15801561054957600080fd5b505afa15801561055d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105819190810190610a81565b5091509150818414801561059457508083145b156105a4576001955050506105b1565b5050600190920191610498565b6001841515141561071e57825b6002547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6fffffffffffffffffffffffffffffffff808316700100000000000000000000000000000000909304811692909201011681101561067e576001818101600081815260209290925260408083205493835290912080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091556105be565b50600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6fffffffffffffffffffffffffffffffff8083167001000000000000000000000000000000009384900482160182018116600090815260016020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016905583548381048216909201811690920291161790555b505050919050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b33ff5b60005473ffffffffffffffffffffffffffffffffffffffff1633146107cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b60025470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff166107fd576108bf565b600280546fffffffffffffffffffffffffffffffff908116600090815260016020819052604090912080547fffffffffffffffffffffffff000000000000000000000000000000000000000016905582547fffffffffffffffffffffffffffffffff0000000000000000000000000000000081169083169091018216178082167001000000000000000000000000000000009182900483167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01909216021790555b565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6060600260109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16604051908082528060200260200182016040528015610945578160200160208202803883390190505b506002549091506fffffffffffffffffffffffffffffffff165b6002546fffffffffffffffffffffffffffffffff80821670010000000000000000000000000000000090920481169190910116811015610a1257600081815260016020526040902054600254835173ffffffffffffffffffffffffffffffffffffffff9092169184916fffffffffffffffffffffffffffffffff1684039081106109e557fe5b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015260010161095f565b5090565b8051610a2181610c1a565b92915050565b8035610a2181610c2e565b8035610a2181610c1a565b600060208284031215610a4f57600080fd5b6000610a5b8484610a27565b949350505050565b600060208284031215610a7557600080fd5b6000610a5b8484610a32565b600080600060608486031215610a9657600080fd5b6000610aa28686610a16565b9350506020610ab386828701610a16565b9250506040610ac486828701610a16565b9150509250925092565b6000610ada8383610b4a565b505060200190565b6000610aed82610bd6565b610af78185610bda565b9350610b0283610bd0565b8060005b83811015610b30578151610b1a8882610ace565b9750610b2583610bd0565b925050600101610b06565b509495945050505050565b610b4481610bee565b82525050565b610b4481610bf6565b6000610b60600a83610bda565b7f4e6f74206f776e65722100000000000000000000000000000000000000000000815260200192915050565b60208082528101610b9d8184610ae2565b9392505050565b60208101610a218284610b3b565b60208101610a218284610b4a565b60208082528101610a2181610b53565b60200190565b5190565b90815260200190565b6000610a2182610c01565b151590565b90565b6000610a2182610be3565b73ffffffffffffffffffffffffffffffffffffffff1690565b610c2381610bf3565b811461030757600080fd5b610c2381610bf656fea365627a7a723158203bd923588e93a109a1e411932047a38e886334bfa4aaed9369d341e6c171e9c86c6578706572696d656e74616cf564736f6c63430005100040`

// DeployStorageCheckpointRegistryV2 deploys a new Ethereum contract, binding an instance of StorageCheckpointRegistryV2 to it.
func DeployStorageCheckpointRegistryV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageCheckpointRegistryV2, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageCheckpointRegistryV2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StorageCheckpointRegistryV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageCheckpointRegistryV2{StorageCheckpointRegistryV2Caller: StorageCheckpointRegistryV2Caller{contract: contract}, StorageCheckpointRegistryV2Transactor: StorageCheckpointRegistryV2Transactor{contract: contract}, StorageCheckpointRegistryV2Filterer: StorageCheckpointRegistryV2Filterer{contract: contract}}, nil
}

// StorageCheckpointRegistryV2Bin is the compiled bytecode of contract after deployment.
const StorageCheckpointRegistryV2RuntimeBin = `608060405234801561001057600080fd5b506004361061007d5760003560e01c806341c0e1b51161005b57806341c0e1b5146100d3578063a4ece52c146100db578063b8a24252146100e3578063d9592ead146101035761007d565b80630a3b0a4f1461008257806313af40351461009757806329092d0e146100aa575b600080fd5b610095610090366004610a3d565b610118565b005b6100956100a5366004610a3d565b61030a565b6100bd6100b8366004610a3d565b6103a2565b6040516100ca9190610ba4565b60405180910390f35b610095610726565b61009561077a565b6100f66100f1366004610a63565b6108c1565b6040516100ca9190610bb2565b61010b6108e9565b6040516100ca9190610b8c565b60005473ffffffffffffffffffffffffffffffffffffffff163314610172576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b60405180910390fd5b6002546fffffffffffffffffffffffffffffffff70010000000000000000000000000000000090910416600a141561026e57600280546fffffffffffffffffffffffffffffffff908116600090815260016020819052604080832080547fffffffffffffffffffffffff000000000000000000000000000000000000000090811690915585547001000000000000000000000000000000008104861690861601851684529220805490921673ffffffffffffffffffffffffffffffffffffffff86161790915582547fffffffffffffffffffffffffffffffff000000000000000000000000000000008116908316909101909116179055610307565b600280546fffffffffffffffffffffffffffffffff8082167001000000000000000000000000000000009283900482160181166000908152600160208190526040909120805473ffffffffffffffffffffffffffffffffffffffff87167fffffffffffffffffffffffff000000000000000000000000000000000000000090911617905583548381048316909101821690920291161790555b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461035b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000805473ffffffffffffffffffffffffffffffffffffffff1633146103f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b60008091506000808473ffffffffffffffffffffffffffffffffffffffff1663370158ea6040518163ffffffff1660e01b815260040160606040518083038186803b15801561044257600080fd5b505afa158015610456573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061047a9190810190610a81565b506002546fffffffffffffffffffffffffffffffff16945090925090505b6002546fffffffffffffffffffffffffffffffff808216700100000000000000000000000000000000909204811691909101168310156105b1576000838152600160205260408082205481517f370158ea0000000000000000000000000000000000000000000000000000000081529151839273ffffffffffffffffffffffffffffffffffffffff9092169163370158ea916004808301926060929190829003018186803b15801561054957600080fd5b505afa15801561055d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105819190810190610a81565b5091509150818414801561059457508083145b156105a4576001955050506105b1565b5050600190920191610498565b6001841515141561071e57825b6002547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6fffffffffffffffffffffffffffffffff808316700100000000000000000000000000000000909304811692909201011681101561067e576001818101600081815260209290925260408083205493835290912080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091556105be565b50600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6fffffffffffffffffffffffffffffffff8083167001000000000000000000000000000000009384900482160182018116600090815260016020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016905583548381048216909201811690920291161790555b505050919050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b33ff5b60005473ffffffffffffffffffffffffffffffffffffffff1633146107cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016990610bc0565b60025470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff166107fd576108bf565b600280546fffffffffffffffffffffffffffffffff908116600090815260016020819052604090912080547fffffffffffffffffffffffff000000000000000000000000000000000000000016905582547fffffffffffffffffffffffffffffffff0000000000000000000000000000000081169083169091018216178082167001000000000000000000000000000000009182900483167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01909216021790555b565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6060600260109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16604051908082528060200260200182016040528015610945578160200160208202803883390190505b506002549091506fffffffffffffffffffffffffffffffff165b6002546fffffffffffffffffffffffffffffffff80821670010000000000000000000000000000000090920481169190910116811015610a1257600081815260016020526040902054600254835173ffffffffffffffffffffffffffffffffffffffff9092169184916fffffffffffffffffffffffffffffffff1684039081106109e557fe5b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015260010161095f565b5090565b8051610a2181610c1a565b92915050565b8035610a2181610c2e565b8035610a2181610c1a565b600060208284031215610a4f57600080fd5b6000610a5b8484610a27565b949350505050565b600060208284031215610a7557600080fd5b6000610a5b8484610a32565b600080600060608486031215610a9657600080fd5b6000610aa28686610a16565b9350506020610ab386828701610a16565b9250506040610ac486828701610a16565b9150509250925092565b6000610ada8383610b4a565b505060200190565b6000610aed82610bd6565b610af78185610bda565b9350610b0283610bd0565b8060005b83811015610b30578151610b1a8882610ace565b9750610b2583610bd0565b925050600101610b06565b509495945050505050565b610b4481610bee565b82525050565b610b4481610bf6565b6000610b60600a83610bda565b7f4e6f74206f776e65722100000000000000000000000000000000000000000000815260200192915050565b60208082528101610b9d8184610ae2565b9392505050565b60208101610a218284610b3b565b60208101610a218284610b4a565b60208082528101610a2181610b53565b60200190565b5190565b90815260200190565b6000610a2182610c01565b151590565b90565b6000610a2182610be3565b73ffffffffffffffffffffffffffffffffffffffff1690565b610c2381610bf3565b811461030757600080fd5b610c2381610bf656fea365627a7a723158203bd923588e93a109a1e411932047a38e886334bfa4aaed9369d341e6c171e9c86c6578706572696d656e74616cf564736f6c63430005100040`

// StorageCheckpointRegistryV2 is an auto generated Go binding around an Ethereum contract.
type StorageCheckpointRegistryV2 struct {
	StorageCheckpointRegistryV2Caller     // Read-only binding to the contract
	StorageCheckpointRegistryV2Transactor // Write-only binding to the contract
	StorageCheckpointRegistryV2Filterer   // Log filterer for contract events
}

// StorageCheckpointRegistryV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCheckpointRegistryV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageCheckpointRegistryV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageCheckpointRegistryV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageCheckpointRegistryV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageCheckpointRegistryV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageCheckpointRegistryV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageCheckpointRegistryV2Session struct {
	Contract     *StorageCheckpointRegistryV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StorageCheckpointRegistryV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCheckpointRegistryV2CallerSession struct {
	Contract *StorageCheckpointRegistryV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// StorageCheckpointRegistryV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageCheckpointRegistryV2TransactorSession struct {
	Contract     *StorageCheckpointRegistryV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// StorageCheckpointRegistryV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type StorageCheckpointRegistryV2Raw struct {
	Contract *StorageCheckpointRegistryV2 // Generic contract binding to access the raw methods on
}

// StorageCheckpointRegistryV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCheckpointRegistryV2CallerRaw struct {
	Contract *StorageCheckpointRegistryV2Caller // Generic read-only contract binding to access the raw methods on
}

// StorageCheckpointRegistryV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageCheckpointRegistryV2TransactorRaw struct {
	Contract *StorageCheckpointRegistryV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageCheckpointRegistryV2 creates a new instance of StorageCheckpointRegistryV2, bound to a specific deployed contract.
func NewStorageCheckpointRegistryV2(address common.Address, backend bind.ContractBackend) (*StorageCheckpointRegistryV2, error) {
	contract, err := bindStorageCheckpointRegistryV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageCheckpointRegistryV2{StorageCheckpointRegistryV2Caller: StorageCheckpointRegistryV2Caller{contract: contract}, StorageCheckpointRegistryV2Transactor: StorageCheckpointRegistryV2Transactor{contract: contract}, StorageCheckpointRegistryV2Filterer: StorageCheckpointRegistryV2Filterer{contract: contract}}, nil
}

// NewStorageCheckpointRegistryV2Caller creates a new read-only instance of StorageCheckpointRegistryV2, bound to a specific deployed contract.
func NewStorageCheckpointRegistryV2Caller(address common.Address, caller bind.ContractCaller) (*StorageCheckpointRegistryV2Caller, error) {
	contract, err := bindStorageCheckpointRegistryV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCheckpointRegistryV2Caller{contract: contract}, nil
}

// NewStorageCheckpointRegistryV2Transactor creates a new write-only instance of StorageCheckpointRegistryV2, bound to a specific deployed contract.
func NewStorageCheckpointRegistryV2Transactor(address common.Address, transactor bind.ContractTransactor) (*StorageCheckpointRegistryV2Transactor, error) {
	contract, err := bindStorageCheckpointRegistryV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCheckpointRegistryV2Transactor{contract: contract}, nil
}

// NewStorageCheckpointRegistryV2Filterer creates a new log filterer instance of StorageCheckpointRegistryV2, bound to a specific deployed contract.
func NewStorageCheckpointRegistryV2Filterer(address common.Address, filterer bind.ContractFilterer) (*StorageCheckpointRegistryV2Filterer, error) {
	contract, err := bindStorageCheckpointRegistryV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageCheckpointRegistryV2Filterer{contract: contract}, nil
}

// bindStorageCheckpointRegistryV2 binds a generic wrapper to an already deployed contract.
func bindStorageCheckpointRegistryV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageCheckpointRegistryV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StorageCheckpointRegistryV2.Contract.StorageCheckpointRegistryV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.StorageCheckpointRegistryV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.StorageCheckpointRegistryV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StorageCheckpointRegistryV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.contract.Transact(opts, method, params...)
}

// Checkpoints is a free data retrieval call binding the contract method 0xb8a24252.
//
// Solidity: function checkpoints(uint256 ) constant returns(address)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Caller) Checkpoints(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StorageCheckpointRegistryV2.contract.Call(opts, out, "checkpoints", arg0)
	return *ret0, err
}

// Checkpoints is a free data retrieval call binding the contract method 0xb8a24252.
//
// Solidity: function checkpoints(uint256 ) constant returns(address)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Session) Checkpoints(arg0 *big.Int) (common.Address, error) {
	return _StorageCheckpointRegistryV2.Contract.Checkpoints(&_StorageCheckpointRegistryV2.CallOpts, arg0)
}

// Checkpoints is a free data retrieval call binding the contract method 0xb8a24252.
//
// Solidity: function checkpoints(uint256 ) constant returns(address)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2CallerSession) Checkpoints(arg0 *big.Int) (common.Address, error) {
	return _StorageCheckpointRegistryV2.Contract.Checkpoints(&_StorageCheckpointRegistryV2.CallOpts, arg0)
}

// ListCheckpoints is a free data retrieval call binding the contract method 0xd9592ead.
//
// Solidity: function listCheckpoints() constant returns(address[] res)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Caller) ListCheckpoints(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _StorageCheckpointRegistryV2.contract.Call(opts, out, "listCheckpoints")
	return *ret0, err
}

// ListCheckpoints is a free data retrieval call binding the contract method 0xd9592ead.
//
// Solidity: function listCheckpoints() constant returns(address[] res)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Session) ListCheckpoints() ([]common.Address, error) {
	return _StorageCheckpointRegistryV2.Contract.ListCheckpoints(&_StorageCheckpointRegistryV2.CallOpts)
}

// ListCheckpoints is a free data retrieval call binding the contract method 0xd9592ead.
//
// Solidity: function listCheckpoints() constant returns(address[] res)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2CallerSession) ListCheckpoints() ([]common.Address, error) {
	return _StorageCheckpointRegistryV2.Contract.ListCheckpoints(&_StorageCheckpointRegistryV2.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address cp) returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Transactor) Add(opts *bind.TransactOpts, cp common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.contract.Transact(opts, "add", cp)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address cp) returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Session) Add(cp common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.Add(&_StorageCheckpointRegistryV2.TransactOpts, cp)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address cp) returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2TransactorSession) Add(cp common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.Add(&_StorageCheckpointRegistryV2.TransactOpts, cp)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Transactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Session) Kill() (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.Kill(&_StorageCheckpointRegistryV2.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2TransactorSession) Kill() (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.Kill(&_StorageCheckpointRegistryV2.TransactOpts)
}

// Pop is a paid mutator transaction binding the contract method 0xa4ece52c.
//
// Solidity: function pop() returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Transactor) Pop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.contract.Transact(opts, "pop")
}

// Pop is a paid mutator transaction binding the contract method 0xa4ece52c.
//
// Solidity: function pop() returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Session) Pop() (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.Pop(&_StorageCheckpointRegistryV2.TransactOpts)
}

// Pop is a paid mutator transaction binding the contract method 0xa4ece52c.
//
// Solidity: function pop() returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2TransactorSession) Pop() (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.Pop(&_StorageCheckpointRegistryV2.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address cp) returns(bool found)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Transactor) Remove(opts *bind.TransactOpts, cp common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.contract.Transact(opts, "remove", cp)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address cp) returns(bool found)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Session) Remove(cp common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.Remove(&_StorageCheckpointRegistryV2.TransactOpts, cp)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address cp) returns(bool found)
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2TransactorSession) Remove(cp common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.Remove(&_StorageCheckpointRegistryV2.TransactOpts, cp)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Transactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2Session) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.SetOwner(&_StorageCheckpointRegistryV2.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_StorageCheckpointRegistryV2 *StorageCheckpointRegistryV2TransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _StorageCheckpointRegistryV2.Contract.SetOwner(&_StorageCheckpointRegistryV2.TransactOpts, _newOwner)
}
