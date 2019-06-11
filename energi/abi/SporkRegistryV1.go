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

// SporkRegistryV1ABI is the input ABI used to generate the binding from.
const SporkRegistryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"createUpgradeProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mnregistry_proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"},{\"name\":\"_mnregistry_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// SporkRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const SporkRegistryV1Bin = `608060405234801561001057600080fd5b506040516115b83803806115b88339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b0319918216179091556001805493909216921691909117905561153e8061007a6000396000f3fe6080604052600436106100595760003560e01c8063ce5494bb11610043578063ce5494bb14610164578063ec556889146101a4578063fe7334e8146101b957610059565b8062f55d9d146100c05780631684f69f14610102575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100cc57600080fd5b50610100600480360360208110156100e357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101ce565b005b61013b6004803603604081101561011857600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610276565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561017057600080fd5b506101006004803603602081101561018757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166104b0565b3480156101b057600080fd5b5061013b61053e565b3480156101c557600080fd5b5061013b61055a565b60005473ffffffffffffffffffffffffffffffffffffffff16331461025457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61025d8161053b565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b600069021e19e0c9bab240000034146102f057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f46656520616d6f756e7400000000000000000000000000000000000000000000604482015290519081900360640190fd5b6212750082101561036257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f506572696f64206d696e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b6301e133808211156103d557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f506572696f64206d617800000000000000000000000000000000000000000000604482015290519081900360640190fd5b60015460405160009173ffffffffffffffffffffffffffffffffffffffff169060339085903290349061040790610576565b73ffffffffffffffffffffffffffffffffffffffff958616815260ff909416602085015260408085019390935293166060830152608082019290925290519081900360a001906000f080158015610462573d6000803e3d6000fd5b5060405190915073ffffffffffffffffffffffffffffffffffffffff8216903480156108fc02916000818181858888f193505050501580156104a8573d6000803e3d6000fd5b509392505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461053657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61053b815b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b610f86806105848339019056fe608060405234801561001057600080fd5b50604051610f86380380610f86833981810160405260a081101561003357600080fd5b508051602080830151604080850151606086015160809096015160018054336001600160a01b031991821617909155436002556000805482166001600160a01b03808a169182178355600585905542860160035560048054909416908b1617835585517f8abf607700000000000000000000000000000000000000000000000000000000815295519899969894979395919485949193638abf60779380820193929190829003018186803b1580156100ea57600080fd5b505afa1580156100fe573d6000803e3d6000fd5b505050506040513d602081101561011457600080fd5b5051604080517f06661abd00000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216916306661abd9160048082019260a092909190829003018186803b15801561017157600080fd5b505afa158015610185573d6000803e3d6000fd5b505050506040513d60a081101561019b57600080fd5b50604081015160809091015190925090506002810482101561021e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f41637469766520776569676874203c20312f3220657665722077656967687400604482015290519081900360640190fd5b600160ff8716101561022c57fe5b606460ff8716111561023a57fe5b6008829055606460ff8716830204600955603360ff87161061026157600954600a5561026c565b60646033830204600a555b50505050505050610d04806102826000396000f3fe6080604052600436106101445760003560e01c80637b352962116100c0578063c2472ef811610074578063c86e6c1511610059578063c86e6c1514610396578063e5225381146103ab578063fe7334e8146103c057610144565b8063c2472ef81461036c578063c40a70f81461038157610144565b806391840a6b116100a557806391840a6b14610302578063990a663b14610317578063aec2ccae1461032c57610144565b80637b352962146102d857806383197ef0146102ed57610144565b80635051a5ec1161011757806360f96a8f116100fc57806360f96a8f1461027057806375df0f99146102ae5780637639b1eb146102c357610144565b80635051a5ec1461023257806356c2a0a11461025b57610144565b80630b62be45146101cc57806329dcb0cf146101f35780633ccfd60b146102085780633d1db3e91461021d575b60015473ffffffffffffffffffffffffffffffffffffffff1633146101ca57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b005b3480156101d857600080fd5b506101e16103d5565b60408051918252519081900360200190f35b3480156101ff57600080fd5b506101e16103db565b34801561021457600080fd5b506101ca6103e1565b34801561022957600080fd5b506101e161049e565b34801561023e57600080fd5b506102476104a4565b604080519115158252519081900360200190f35b34801561026757600080fd5b506101ca6104f1565b34801561027c57600080fd5b50610285610504565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156102ba57600080fd5b506101e1610520565b3480156102cf57600080fd5b506101e1610526565b3480156102e457600080fd5b5061024761052c565b3480156102f957600080fd5b506101ca610556565b34801561030e57600080fd5b506101e16105f7565b34801561032357600080fd5b506101e16105fd565b34801561033857600080fd5b506102476004803603602081101561034f57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610603565b34801561037857600080fd5b506101ca610618565b34801561038d57600080fd5b5061028561062b565b3480156103a257600080fd5b506101e1610647565b3480156103b757600080fd5b506101ca61064d565b3480156103cc57600080fd5b50610285610994565b60025481565b60035481565b6103e96104a4565b61045457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742061636365707465640000000000000000000000000000000000000000604482015290519081900360640190fd5b60045460405173ffffffffffffffffffffffffffffffffffffffff90911690303180156108fc02916000818181858888f1935050505015801561049b573d6000803e3d6000fd5b50565b600a5481565b6000600a54600654106104b9575060016104ee565b6104c161052c565b6104cd575060006104ee565b6009546007546006540110156104e5575060006104ee565b50600754600654115b90565b6104f96109b0565b600780549091019055565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60065481565b6000426003541115806105435750600a5460065410155b806105515750600a54600754115b905090565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105dc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f7420706172656e7400000000000000000000000000000000000000000000604482015290519081900360640190fd5b60045473ffffffffffffffffffffffffffffffffffffffff16ff5b60085481565b60055481565b600b6020526000908152604090205460ff1681565b6106206109b0565b600680549091019055565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b61065561052c565b801561066657506106646104a4565b155b6106d157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f7420636f6c6c65637461626c650000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b15801561073a57600080fd5b505afa15801561074e573d6000803e3d6000fd5b505050506040513d602081101561076457600080fd5b5051604080517fa2731784000000000000000000000000000000000000000000000000000000008152905191925060009173ffffffffffffffffffffffffffffffffffffffff84169163a2731784916004808301926020929190829003018186803b1580156107d257600080fd5b505afa1580156107e6573d6000803e3d6000fd5b505050506040513d60208110156107fc57600080fd5b5051604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921691638abf607791600480820192602092909190829003018186803b15801561086657600080fd5b505afa15801561087a573d6000803e3d6000fd5b505050506040513d602081101561089057600080fd5b505190503373ffffffffffffffffffffffffffffffffffffffff82161461091857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742074726561737572790000000000000000000000000000000000000000604482015290519081900360640190fd5b8073ffffffffffffffffffffffffffffffffffffffff1663d7bb99ba3073ffffffffffffffffffffffffffffffffffffffff16316040518263ffffffff1660e01b81526004016000604051808303818588803b15801561097757600080fd5b505af115801561098b573d6000803e3d6000fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60004260035411610a2257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f46696e6973686564000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610a8b57600080fd5b505afa158015610a9f573d6000803e3d6000fd5b505050506040513d6020811015610ab557600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815233600482018190529151929350909160009173ffffffffffffffffffffffffffffffffffffffff85169163b83e16059160248082019260c092909190829003018186803b158015610b2e57600080fd5b505afa158015610b42573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c0811015610b6757600080fd5b50608081015160a09091015160025491955091508110610be857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f7420656c696769626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600b602052604090205460ff1615610c7d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b5073ffffffffffffffffffffffffffffffffffffffff166000908152600b6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055509056fea265627a7a72305820a87ad64d6c170a670b3d724a09140219efd30396f8e2d2990942049a7e11f33764736f6c63430005090032a265627a7a723058209b5927cd6012c0d88656b819a1f9a810223f75a031ae73a7d66bcbcc16b1ffae64736f6c63430005090032`

// DeploySporkRegistryV1 deploys a new Ethereum contract, binding an instance of SporkRegistryV1 to it.
func DeploySporkRegistryV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _mnregistry_proxy common.Address) (common.Address, *types.Transaction, *SporkRegistryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(SporkRegistryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SporkRegistryV1Bin), backend, _proxy, _mnregistry_proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SporkRegistryV1{SporkRegistryV1Caller: SporkRegistryV1Caller{contract: contract}, SporkRegistryV1Transactor: SporkRegistryV1Transactor{contract: contract}, SporkRegistryV1Filterer: SporkRegistryV1Filterer{contract: contract}}, nil
}

// SporkRegistryV1 is an auto generated Go binding around an Ethereum contract.
type SporkRegistryV1 struct {
	SporkRegistryV1Caller     // Read-only binding to the contract
	SporkRegistryV1Transactor // Write-only binding to the contract
	SporkRegistryV1Filterer   // Log filterer for contract events
}

// SporkRegistryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SporkRegistryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SporkRegistryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SporkRegistryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SporkRegistryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SporkRegistryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SporkRegistryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SporkRegistryV1Session struct {
	Contract     *SporkRegistryV1  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SporkRegistryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SporkRegistryV1CallerSession struct {
	Contract *SporkRegistryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SporkRegistryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SporkRegistryV1TransactorSession struct {
	Contract     *SporkRegistryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SporkRegistryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SporkRegistryV1Raw struct {
	Contract *SporkRegistryV1 // Generic contract binding to access the raw methods on
}

// SporkRegistryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SporkRegistryV1CallerRaw struct {
	Contract *SporkRegistryV1Caller // Generic read-only contract binding to access the raw methods on
}

// SporkRegistryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SporkRegistryV1TransactorRaw struct {
	Contract *SporkRegistryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSporkRegistryV1 creates a new instance of SporkRegistryV1, bound to a specific deployed contract.
func NewSporkRegistryV1(address common.Address, backend bind.ContractBackend) (*SporkRegistryV1, error) {
	contract, err := bindSporkRegistryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SporkRegistryV1{SporkRegistryV1Caller: SporkRegistryV1Caller{contract: contract}, SporkRegistryV1Transactor: SporkRegistryV1Transactor{contract: contract}, SporkRegistryV1Filterer: SporkRegistryV1Filterer{contract: contract}}, nil
}

// NewSporkRegistryV1Caller creates a new read-only instance of SporkRegistryV1, bound to a specific deployed contract.
func NewSporkRegistryV1Caller(address common.Address, caller bind.ContractCaller) (*SporkRegistryV1Caller, error) {
	contract, err := bindSporkRegistryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SporkRegistryV1Caller{contract: contract}, nil
}

// NewSporkRegistryV1Transactor creates a new write-only instance of SporkRegistryV1, bound to a specific deployed contract.
func NewSporkRegistryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SporkRegistryV1Transactor, error) {
	contract, err := bindSporkRegistryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SporkRegistryV1Transactor{contract: contract}, nil
}

// NewSporkRegistryV1Filterer creates a new log filterer instance of SporkRegistryV1, bound to a specific deployed contract.
func NewSporkRegistryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SporkRegistryV1Filterer, error) {
	contract, err := bindSporkRegistryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SporkRegistryV1Filterer{contract: contract}, nil
}

// bindSporkRegistryV1 binds a generic wrapper to an already deployed contract.
func bindSporkRegistryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SporkRegistryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SporkRegistryV1 *SporkRegistryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SporkRegistryV1.Contract.SporkRegistryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SporkRegistryV1 *SporkRegistryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.SporkRegistryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SporkRegistryV1 *SporkRegistryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.SporkRegistryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SporkRegistryV1 *SporkRegistryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SporkRegistryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SporkRegistryV1 *SporkRegistryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SporkRegistryV1 *SporkRegistryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.contract.Transact(opts, method, params...)
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Caller) MnregistryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SporkRegistryV1.contract.Call(opts, out, "mnregistry_proxy")
	return *ret0, err
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Session) MnregistryProxy() (common.Address, error) {
	return _SporkRegistryV1.Contract.MnregistryProxy(&_SporkRegistryV1.CallOpts)
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1CallerSession) MnregistryProxy() (common.Address, error) {
	return _SporkRegistryV1.Contract.MnregistryProxy(&_SporkRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SporkRegistryV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Session) Proxy() (common.Address, error) {
	return _SporkRegistryV1.Contract.Proxy(&_SporkRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_SporkRegistryV1 *SporkRegistryV1CallerSession) Proxy() (common.Address, error) {
	return _SporkRegistryV1.Contract.Proxy(&_SporkRegistryV1.CallOpts)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x1684f69f.
//
// Solidity: function createUpgradeProposal(address , uint256 _period) returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Transactor) CreateUpgradeProposal(opts *bind.TransactOpts, arg0 common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SporkRegistryV1.contract.Transact(opts, "createUpgradeProposal", arg0, _period)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x1684f69f.
//
// Solidity: function createUpgradeProposal(address , uint256 _period) returns(address)
func (_SporkRegistryV1 *SporkRegistryV1Session) CreateUpgradeProposal(arg0 common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.CreateUpgradeProposal(&_SporkRegistryV1.TransactOpts, arg0, _period)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x1684f69f.
//
// Solidity: function createUpgradeProposal(address , uint256 _period) returns(address)
func (_SporkRegistryV1 *SporkRegistryV1TransactorSession) CreateUpgradeProposal(arg0 common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.CreateUpgradeProposal(&_SporkRegistryV1.TransactOpts, arg0, _period)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_SporkRegistryV1 *SporkRegistryV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_SporkRegistryV1 *SporkRegistryV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Destroy(&_SporkRegistryV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_SporkRegistryV1 *SporkRegistryV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Destroy(&_SporkRegistryV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_SporkRegistryV1 *SporkRegistryV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_SporkRegistryV1 *SporkRegistryV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Migrate(&_SporkRegistryV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_SporkRegistryV1 *SporkRegistryV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.Migrate(&_SporkRegistryV1.TransactOpts, _oldImpl)
}
