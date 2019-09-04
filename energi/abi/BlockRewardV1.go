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

// BlockRewardV1ABI is the input ABI used to generate the binding from.
const BlockRewardV1ABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reward_proxies\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"contractIGovernedProxy[]\",\"name\":\"_reward_proxies\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// BlockRewardV1Bin is the compiled bytecode used for deploying new contracts.
const BlockRewardV1Bin = `608060405234801561001057600080fd5b506040516109b83803806109b88339818101604052604081101561003357600080fd5b81516020830180516040519294929383019291908464010000000082111561005a57600080fd5b90830190602082018581111561006f57600080fd5b825186602082028301116401000000008211171561008c57600080fd5b82525081516020918201928201910280838360005b838110156100b95781810151838201526020016100a1565b50505050919091016040525050600080546001600160a01b0319166001600160a01b038616179055505080515b60001981019015610145578181815181106100fd57fe5b60200260200101516002828154811061011257fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506100e6565b505050610861806101576000396000f3fe6080604052600436106100645760003560e01c80633c92fb74116100435780633c92fb7414610151578063ce5494bb146101a4578063ec556889146101e457610064565b8062f55d9d146100cb5780631c4b774b1461010d578063228cb73314610149575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100d757600080fd5b5061010b600480360360208110156100ee57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101f9565b005b34801561011957600080fd5b506101376004803603602081101561013057600080fd5b50356102a1565b60408051918252519081900360200190f35b61010b610429565b34801561015d57600080fd5b5061017b6004803603602081101561017457600080fd5b503561074e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156101b057600080fd5b5061010b600480360360208110156101c757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610782565b3480156101f057600080fd5b5061017b610810565b60005473ffffffffffffffffffffffffffffffffffffffff16331461027f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6102888161080d565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b6002546000905b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019015610423576000600282815481106102e057fe5b60009182526020918290200154604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921692638abf607792600480840193829003018186803b15801561034f57600080fd5b505afa158015610363573d6000803e3d6000fd5b505050506040513d602081101561037957600080fd5b5051604080517f1c4b774b00000000000000000000000000000000000000000000000000000000815260048101879052905191925073ffffffffffffffffffffffffffffffffffffffff831691631c4b774b91602480820192602092909190829003018186803b1580156103ec57600080fd5b505afa158015610400573d6000803e3d6000fd5b505050506040513d602081101561041657600080fd5b50519290920191506102a8565b50919050565b6001541561049857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556002546000610a8c825a816104ad57fe5b04039050815b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019015610744576000600282815481106104eb57fe5b60009182526020918290200154604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921692638abf607792600480840193829003018186803b15801561055a57600080fd5b505afa15801561056e573d6000803e3d6000fd5b505050506040513d602081101561058457600080fd5b5051604080517f1c4b774b000000000000000000000000000000000000000000000000000000008152436004820152905191925060009173ffffffffffffffffffffffffffffffffffffffff841691631c4b774b91610a8c91602480820192602092909190829003018187803b1580156105fd57600080fd5b5086fa158015610611573d6000803e3d6000fd5b50505050506040513d602081101561062857600080fd5b5051604080517f228cb7330000000000000000000000000000000000000000000000000000000060208281019190915282518083038201815291830192839052815193945073ffffffffffffffffffffffffffffffffffffffff8616938893869392909182918401908083835b602083106106d257805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610695565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381858888f193505050503d8060008114610735576040519150601f19603f3d011682016040523d82523d6000602084013e61073a565b606091505b50505050506104b3565b5050600060015550565b6002818154811061075b57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60005473ffffffffffffffffffffffffffffffffffffffff16331461080857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61080d815b50565b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a72315820151de96baaf42bc3bf6683e9ef0d887135aefe283395f3dcd5b5b6aee42deec064736f6c634300050b0032`

// DeployBlockRewardV1 deploys a new Ethereum contract, binding an instance of BlockRewardV1 to it.
func DeployBlockRewardV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _reward_proxies []common.Address) (common.Address, *types.Transaction, *BlockRewardV1, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockRewardV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlockRewardV1Bin), backend, _proxy, _reward_proxies)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockRewardV1{BlockRewardV1Caller: BlockRewardV1Caller{contract: contract}, BlockRewardV1Transactor: BlockRewardV1Transactor{contract: contract}, BlockRewardV1Filterer: BlockRewardV1Filterer{contract: contract}}, nil
}

// BlockRewardV1Bin is the compiled bytecode of contract after deployment.
const BlockRewardV1RuntimeBin = `6080604052600436106100645760003560e01c80633c92fb74116100435780633c92fb7414610151578063ce5494bb146101a4578063ec556889146101e457610064565b8062f55d9d146100cb5780631c4b774b1461010d578063228cb73314610149575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100d757600080fd5b5061010b600480360360208110156100ee57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101f9565b005b34801561011957600080fd5b506101376004803603602081101561013057600080fd5b50356102a1565b60408051918252519081900360200190f35b61010b610429565b34801561015d57600080fd5b5061017b6004803603602081101561017457600080fd5b503561074e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156101b057600080fd5b5061010b600480360360208110156101c757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610782565b3480156101f057600080fd5b5061017b610810565b60005473ffffffffffffffffffffffffffffffffffffffff16331461027f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6102888161080d565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b6002546000905b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019015610423576000600282815481106102e057fe5b60009182526020918290200154604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921692638abf607792600480840193829003018186803b15801561034f57600080fd5b505afa158015610363573d6000803e3d6000fd5b505050506040513d602081101561037957600080fd5b5051604080517f1c4b774b00000000000000000000000000000000000000000000000000000000815260048101879052905191925073ffffffffffffffffffffffffffffffffffffffff831691631c4b774b91602480820192602092909190829003018186803b1580156103ec57600080fd5b505afa158015610400573d6000803e3d6000fd5b505050506040513d602081101561041657600080fd5b50519290920191506102a8565b50919050565b6001541561049857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556002546000610a8c825a816104ad57fe5b04039050815b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019015610744576000600282815481106104eb57fe5b60009182526020918290200154604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921692638abf607792600480840193829003018186803b15801561055a57600080fd5b505afa15801561056e573d6000803e3d6000fd5b505050506040513d602081101561058457600080fd5b5051604080517f1c4b774b000000000000000000000000000000000000000000000000000000008152436004820152905191925060009173ffffffffffffffffffffffffffffffffffffffff841691631c4b774b91610a8c91602480820192602092909190829003018187803b1580156105fd57600080fd5b5086fa158015610611573d6000803e3d6000fd5b50505050506040513d602081101561062857600080fd5b5051604080517f228cb7330000000000000000000000000000000000000000000000000000000060208281019190915282518083038201815291830192839052815193945073ffffffffffffffffffffffffffffffffffffffff8616938893869392909182918401908083835b602083106106d257805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610695565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381858888f193505050503d8060008114610735576040519150601f19603f3d011682016040523d82523d6000602084013e61073a565b606091505b50505050506104b3565b5050600060015550565b6002818154811061075b57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60005473ffffffffffffffffffffffffffffffffffffffff16331461080857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61080d815b50565b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a72315820151de96baaf42bc3bf6683e9ef0d887135aefe283395f3dcd5b5b6aee42deec064736f6c634300050b0032`

// BlockRewardV1 is an auto generated Go binding around an Ethereum contract.
type BlockRewardV1 struct {
	BlockRewardV1Caller     // Read-only binding to the contract
	BlockRewardV1Transactor // Write-only binding to the contract
	BlockRewardV1Filterer   // Log filterer for contract events
}

// BlockRewardV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type BlockRewardV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockRewardV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockRewardV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockRewardV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockRewardV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockRewardV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockRewardV1Session struct {
	Contract     *BlockRewardV1    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockRewardV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockRewardV1CallerSession struct {
	Contract *BlockRewardV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BlockRewardV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockRewardV1TransactorSession struct {
	Contract     *BlockRewardV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BlockRewardV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type BlockRewardV1Raw struct {
	Contract *BlockRewardV1 // Generic contract binding to access the raw methods on
}

// BlockRewardV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockRewardV1CallerRaw struct {
	Contract *BlockRewardV1Caller // Generic read-only contract binding to access the raw methods on
}

// BlockRewardV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockRewardV1TransactorRaw struct {
	Contract *BlockRewardV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockRewardV1 creates a new instance of BlockRewardV1, bound to a specific deployed contract.
func NewBlockRewardV1(address common.Address, backend bind.ContractBackend) (*BlockRewardV1, error) {
	contract, err := bindBlockRewardV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockRewardV1{BlockRewardV1Caller: BlockRewardV1Caller{contract: contract}, BlockRewardV1Transactor: BlockRewardV1Transactor{contract: contract}, BlockRewardV1Filterer: BlockRewardV1Filterer{contract: contract}}, nil
}

// NewBlockRewardV1Caller creates a new read-only instance of BlockRewardV1, bound to a specific deployed contract.
func NewBlockRewardV1Caller(address common.Address, caller bind.ContractCaller) (*BlockRewardV1Caller, error) {
	contract, err := bindBlockRewardV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockRewardV1Caller{contract: contract}, nil
}

// NewBlockRewardV1Transactor creates a new write-only instance of BlockRewardV1, bound to a specific deployed contract.
func NewBlockRewardV1Transactor(address common.Address, transactor bind.ContractTransactor) (*BlockRewardV1Transactor, error) {
	contract, err := bindBlockRewardV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockRewardV1Transactor{contract: contract}, nil
}

// NewBlockRewardV1Filterer creates a new log filterer instance of BlockRewardV1, bound to a specific deployed contract.
func NewBlockRewardV1Filterer(address common.Address, filterer bind.ContractFilterer) (*BlockRewardV1Filterer, error) {
	contract, err := bindBlockRewardV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockRewardV1Filterer{contract: contract}, nil
}

// bindBlockRewardV1 binds a generic wrapper to an already deployed contract.
func bindBlockRewardV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockRewardV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockRewardV1 *BlockRewardV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlockRewardV1.Contract.BlockRewardV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockRewardV1 *BlockRewardV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.BlockRewardV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockRewardV1 *BlockRewardV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.BlockRewardV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockRewardV1 *BlockRewardV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlockRewardV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockRewardV1 *BlockRewardV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockRewardV1 *BlockRewardV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.contract.Transact(opts, method, params...)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BlockRewardV1 *BlockRewardV1Caller) GetReward(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BlockRewardV1.contract.Call(opts, out, "getReward", _blockNumber)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BlockRewardV1 *BlockRewardV1Session) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _BlockRewardV1.Contract.GetReward(&_BlockRewardV1.CallOpts, _blockNumber)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_BlockRewardV1 *BlockRewardV1CallerSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _BlockRewardV1.Contract.GetReward(&_BlockRewardV1.CallOpts, _blockNumber)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlockRewardV1 *BlockRewardV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlockRewardV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlockRewardV1 *BlockRewardV1Session) Proxy() (common.Address, error) {
	return _BlockRewardV1.Contract.Proxy(&_BlockRewardV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_BlockRewardV1 *BlockRewardV1CallerSession) Proxy() (common.Address, error) {
	return _BlockRewardV1.Contract.Proxy(&_BlockRewardV1.CallOpts)
}

// RewardProxies is a free data retrieval call binding the contract method 0x3c92fb74.
//
// Solidity: function reward_proxies(uint256 ) constant returns(address)
func (_BlockRewardV1 *BlockRewardV1Caller) RewardProxies(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlockRewardV1.contract.Call(opts, out, "reward_proxies", arg0)
	return *ret0, err
}

// RewardProxies is a free data retrieval call binding the contract method 0x3c92fb74.
//
// Solidity: function reward_proxies(uint256 ) constant returns(address)
func (_BlockRewardV1 *BlockRewardV1Session) RewardProxies(arg0 *big.Int) (common.Address, error) {
	return _BlockRewardV1.Contract.RewardProxies(&_BlockRewardV1.CallOpts, arg0)
}

// RewardProxies is a free data retrieval call binding the contract method 0x3c92fb74.
//
// Solidity: function reward_proxies(uint256 ) constant returns(address)
func (_BlockRewardV1 *BlockRewardV1CallerSession) RewardProxies(arg0 *big.Int) (common.Address, error) {
	return _BlockRewardV1.Contract.RewardProxies(&_BlockRewardV1.CallOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlockRewardV1 *BlockRewardV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlockRewardV1 *BlockRewardV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Destroy(&_BlockRewardV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_BlockRewardV1 *BlockRewardV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Destroy(&_BlockRewardV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlockRewardV1 *BlockRewardV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlockRewardV1 *BlockRewardV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Migrate(&_BlockRewardV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_BlockRewardV1 *BlockRewardV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Migrate(&_BlockRewardV1.TransactOpts, _oldImpl)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BlockRewardV1 *BlockRewardV1Transactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockRewardV1.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BlockRewardV1 *BlockRewardV1Session) Reward() (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Reward(&_BlockRewardV1.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_BlockRewardV1 *BlockRewardV1TransactorSession) Reward() (*types.Transaction, error) {
	return _BlockRewardV1.Contract.Reward(&_BlockRewardV1.TransactOpts)
}
