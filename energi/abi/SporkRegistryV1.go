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
const SporkRegistryV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"_period\",\"type\":\"uint256\"},{\"name\":\"_fee_payer\",\"type\":\"address\"}],\"name\":\"createUpgradeProposal\",\"outputs\":[{\"name\":\"proposal\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mnregistry_proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"},{\"name\":\"_mnregistry_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// SporkRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const SporkRegistryV1Bin = `608060405234801561001057600080fd5b506040516116f13803806116f18339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556116778061007a6000396000f3fe6080604052600436106100595760003560e01c8063ce5494bb11610043578063ce5494bb1461016e578063ec556889146101ae578063fe7334e8146101c357610059565b8062f55d9d146100c057806362877ccd14610102575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100cc57600080fd5b50610100600480360360208110156100e357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101d8565b005b6101456004803603606081101561011857600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101359160409091013516610280565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561017a57600080fd5b506101006004803603602081101561019157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166104cd565b3480156101ba57600080fd5b5061014561055b565b3480156101cf57600080fd5b50610145610577565b60005473ffffffffffffffffffffffffffffffffffffffff16331461025e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61026781610558565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b600069021e19e0c9bab240000034146102fa57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f496e76616c696420666565000000000000000000000000000000000000000000604482015290519081900360640190fd5b6212750083101561036c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f506572696f64206d696e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b6301e133808311156103df57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f506572696f64206d617800000000000000000000000000000000000000000000604482015290519081900360640190fd5b60015460405173ffffffffffffffffffffffffffffffffffffffff909116906033908590859061040e90610593565b73ffffffffffffffffffffffffffffffffffffffff948516815260ff90931660208401526040808401929092529092166060820152905190819003608001906000f080158015610462573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff16632ded3227346040518263ffffffff1660e01b81526004016000604051808303818588803b1580156104ad57600080fd5b505af11580156104c1573d6000803e3d6000fd5b50505050509392505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461055357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610558815b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6110a2806105a18339019056fe608060405234801561001057600080fd5b506040516110a23803806110a28339818101604052608081101561003357600080fd5b50805160208083015160408085015160609095015160018054336001600160a01b031991821617909155436002556000805482166001600160a01b038089169182178355428a016003556004805490941690851617835584517f8abf607700000000000000000000000000000000000000000000000000000000815294519798959795969395919485949193638abf6077938282019392909190829003018186803b1580156100e157600080fd5b505afa1580156100f5573d6000803e3d6000fd5b505050506040513d602081101561010b57600080fd5b5051604080517f06661abd00000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216916306661abd9160048082019260a092909190829003018186803b15801561016857600080fd5b505afa15801561017c573d6000803e3d6000fd5b505050506040513d60a081101561019257600080fd5b50604081015160809091015190925090508061020f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4e6f7420726561647920666f722070726f706f73616c73000000000000000000604482015290519081900360640190fd5b6002810482101561028157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f41637469766520776569676874203c20312f3220657665722077656967687400604482015290519081900360640190fd5b600160ff8616101561028f57fe5b606460ff8616111561029d57fe5b6008829055606460ff8616830204600955603360ff8616106102c457600954600a556102cf565b60646033830204600a555b6000600954116102db57fe5b6000600a54116102e757fe5b505050505050610da6806102fc6000396000f3fe60806040526004361061015f5760003560e01c80637b352962116100c0578063c2472ef811610074578063c86e6c1511610059578063c86e6c151461039a578063e5225381146103af578063fe7334e8146103c45761015f565b8063c2472ef814610370578063c40a70f8146103855761015f565b806391840a6b116100a557806391840a6b14610306578063990a663b1461031b578063aec2ccae146103305761015f565b80637b352962146102dc57806383197ef0146102f15761015f565b80635051a5ec1161011757806360f96a8f116100fc57806360f96a8f1461027457806375df0f99146102b25780637639b1eb146102c75761015f565b80635051a5ec1461023657806356c2a0a11461025f5761015f565b80632ded3227116101485780632ded3227146102025780633ccfd60b1461020c5780633d1db3e9146102215761015f565b80630b62be45146101c657806329dcb0cf146101ed575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f776564000000000000000000000000000000000000000000604482015290519081900360640190fd5b3480156101d257600080fd5b506101db6103d9565b60408051918252519081900360200190f35b3480156101f957600080fd5b506101db6103df565b61020a6103e5565b005b34801561021857600080fd5b5061020a610475565b34801561022d57600080fd5b506101db610532565b34801561024257600080fd5b5061024b610538565b604080519115158252519081900360200190f35b34801561026b57600080fd5b5061020a610585565b34801561028057600080fd5b50610289610598565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156102be57600080fd5b506101db6105b4565b3480156102d357600080fd5b506101db6105ba565b3480156102e857600080fd5b5061024b6105c0565b3480156102fd57600080fd5b5061020a6105ea565b34801561031257600080fd5b506101db61068b565b34801561032757600080fd5b506101db610691565b34801561033c57600080fd5b5061024b6004803603602081101561035357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610697565b34801561037c57600080fd5b5061020a6106ac565b34801561039157600080fd5b506102896106bf565b3480156103a657600080fd5b506101db6106db565b3480156103bb57600080fd5b5061020a6106e1565b3480156103d057600080fd5b50610289610a36565b60025481565b60035481565b60015473ffffffffffffffffffffffffffffffffffffffff16331461046b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6005805434019055565b61047d610538565b6104e857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742061636365707465640000000000000000000000000000000000000000604482015290519081900360640190fd5b60045460405173ffffffffffffffffffffffffffffffffffffffff90911690303180156108fc02916000818181858888f1935050505015801561052f573d6000803e3d6000fd5b50565b600a5481565b6000600a546006541061054d57506001610582565b6105556105c0565b61056157506000610582565b60095460075460065401101561057957506000610582565b50600754600654115b90565b61058d610a52565b600780549091019055565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60065481565b6000426003541115806105d75750600a5460065410155b806105e55750600a54600754115b905090565b60015473ffffffffffffffffffffffffffffffffffffffff16331461067057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60045473ffffffffffffffffffffffffffffffffffffffff16ff5b60085481565b60055481565b600b6020526000908152604090205460ff1681565b6106b4610a52565b600680549091019055565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b6106e96105c0565b80156106fa57506106f8610538565b155b61076557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f7420636f6c6c65637461626c650000000000000000000000000000000000604482015290519081900360640190fd5b60015473ffffffffffffffffffffffffffffffffffffffff1633146107eb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b15801561085457600080fd5b505afa158015610868573d6000803e3d6000fd5b505050506040513d602081101561087e57600080fd5b5051604080517fa2731784000000000000000000000000000000000000000000000000000000008152905191925060009173ffffffffffffffffffffffffffffffffffffffff84169163a2731784916004808301926020929190829003018186803b1580156108ec57600080fd5b505afa158015610900573d6000803e3d6000fd5b505050506040513d602081101561091657600080fd5b5051604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921691638abf607791600480820192602092909190829003018186803b15801561098057600080fd5b505afa158015610994573d6000803e3d6000fd5b505050506040513d60208110156109aa57600080fd5b5051604080517fd7bb99ba000000000000000000000000000000000000000000000000000000008152905191925073ffffffffffffffffffffffffffffffffffffffff83169163d7bb99ba91303191600480830192600092919082900301818588803b158015610a1957600080fd5b505af1158015610a2d573d6000803e3d6000fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60004260035411610ac457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f46696e6973686564000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610b2d57600080fd5b505afa158015610b41573d6000803e3d6000fd5b505050506040513d6020811015610b5757600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815233600482018190529151929350909160009173ffffffffffffffffffffffffffffffffffffffff85169163b83e16059160248082019260c092909190829003018186803b158015610bd057600080fd5b505afa158015610be4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c0811015610c0957600080fd5b50608081015160a09091015160025491955091508110610c8a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f7420656c696769626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600b602052604090205460ff1615610d1f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b5073ffffffffffffffffffffffffffffffffffffffff166000908152600b6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055509056fea265627a7a72305820b876e4fad8dcf1f657bd31da04e9214170c2333fbe81818d97e3b9c662fa85c564736f6c63430005090032a265627a7a72305820ac154f5a62ff27f6d5de64c312e9005c4564152c43955a824d9053383cf85d2064736f6c63430005090032`

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

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x62877ccd.
//
// Solidity: function createUpgradeProposal(address , uint256 _period, address _fee_payer) returns(address proposal)
func (_SporkRegistryV1 *SporkRegistryV1Transactor) CreateUpgradeProposal(opts *bind.TransactOpts, arg0 common.Address, _period *big.Int, _fee_payer common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.contract.Transact(opts, "createUpgradeProposal", arg0, _period, _fee_payer)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x62877ccd.
//
// Solidity: function createUpgradeProposal(address , uint256 _period, address _fee_payer) returns(address proposal)
func (_SporkRegistryV1 *SporkRegistryV1Session) CreateUpgradeProposal(arg0 common.Address, _period *big.Int, _fee_payer common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.CreateUpgradeProposal(&_SporkRegistryV1.TransactOpts, arg0, _period, _fee_payer)
}

// CreateUpgradeProposal is a paid mutator transaction binding the contract method 0x62877ccd.
//
// Solidity: function createUpgradeProposal(address , uint256 _period, address _fee_payer) returns(address proposal)
func (_SporkRegistryV1 *SporkRegistryV1TransactorSession) CreateUpgradeProposal(arg0 common.Address, _period *big.Int, _fee_payer common.Address) (*types.Transaction, error) {
	return _SporkRegistryV1.Contract.CreateUpgradeProposal(&_SporkRegistryV1.TransactOpts, arg0, _period, _fee_payer)
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
