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

// GovernedProxyABI is the input ABI used to generate the binding from.
const GovernedProxyABI = "[{\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_impl\",\"type\":\"address\"},{\"internalType\":\"contractIGovernedProxy\",\"name\":\"_sporkProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIGovernedContract\",\"name\":\"impl\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIUpgradeProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"UpgradeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIGovernedContract\",\"name\":\"impl\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIUpgradeProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"collectUpgradeProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"impl\",\"outputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"listUpgradeProposals\",\"outputs\":[{\"internalType\":\"contractIUpgradeProposal[]\",\"name\":\"proposals\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_newImpl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"proposeUpgrade\",\"outputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spork_proxy\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"upgradeProposalImpl\",\"outputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"new_impl\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"upgrade_proposal_list\",\"outputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgrade_proposals\",\"outputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GovernedProxyBin is the compiled bytecode used for deploying new contracts.
const GovernedProxyBin = `608060405234801561001057600080fd5b506040516112ea3803806112ea8339818101604052604081101561003357600080fd5b508051602090910151600180546001600160a01b039384166001600160a01b031991821617909155600280549390921692169190911790556112708061007a6000396000f3fe6080604052600436106100c65760003560e01c80638abf607711610074578063ce5494bb1161004e578063ce5494bb14610179578063dd6a851d146103c1578063ec556889146103d6576100c6565b80638abf607714610307578063a1b0e4761461031c578063b364595e1461035c576100c6565b80635b6dee4c116100a55780635b6dee4c146102645780636d5b6c441461029d5780636fa09ab0146102dd576100c6565b8062f55d9d146101795780630900f010146101bb57806332e3a905146101fb575b32331461013457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f6e6c79206469726563742063616c6c732061726520616c6c6f776564210000604482015290519081900360640190fd5b60015460405173ffffffffffffffffffffffffffffffffffffffff909116903660008237600080368334866127105a03f13d6000833e808015610175573d83f35b3d83fd5b34801561018557600080fd5b506101b96004803603602081101561019c57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166103eb565b005b3480156101c757600080fd5b506101b9600480360360208110156101de57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610452565b34801561020757600080fd5b5061023b6004803603602081101561021e57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166108c5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61023b6004803603604081101561027a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356108ed565b3480156102a957600080fd5b5061023b600480360360208110156102c057600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610d84565b3480156102e957600080fd5b5061023b6004803603602081101561030057600080fd5b5035610daf565b34801561031357600080fd5b5061023b610de3565b34801561032857600080fd5b506101b96004803603602081101561033f57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610dff565b34801561036857600080fd5b50610371610fc3565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103ad578181015183820152602001610395565b505050509050019250505060405180910390f35b3480156103cd57600080fd5b5061023b611077565b3480156103e257600080fd5b5061023b611093565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f476f6f6420747279000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600054156104c157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600081815573ffffffffffffffffffffffffffffffffffffffff808416825260036020526040909120549154918116911681141561056257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920616374697665210000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff81166105e457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f742072656769737465726564210000000000000000000000000000000000604482015290519081900360640190fd5b8173ffffffffffffffffffffffffffffffffffffffff16635051a5ec6040518163ffffffff1660e01b815260040160206040518083038186803b15801561062a57600080fd5b505afa15801561063e573d6000803e3d6000fd5b505050506040513d602081101561065457600080fd5b50516106c157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f742061636365707465642100000000000000000000000000000000000000604482015290519081900360640190fd5b600154604080517fce5494bb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff928316600482018190529151919284169163ce5494bb9160248082019260009290919082900301818387803b15801561073857600080fd5b505af115801561074c573d6000803e3d6000fd5b5050600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff868116918217909255604080517ef55d9d000000000000000000000000000000000000000000000000000000008152600481019290925251918516935062f55d9d925060248082019260009290919082900301818387803b1580156107ed57600080fd5b505af1158015610801573d6000803e3d6000fd5b5050505061080e83611097565b8273ffffffffffffffffffffffffffffffffffffffff166383197ef06040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561085657600080fd5b505af115801561086a573d6000803e3d6000fd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8781168252915191861693507f5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7925081900360200190a250506000805550565b60036020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b600032331461095d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f6e6c79206469726563742063616c6c732061726520616c6c6f776564210000604482015290519081900360640190fd5b600054156109cc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160008190555473ffffffffffffffffffffffffffffffffffffffff84811691161415610a5b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920616374697665210000000000000000000000000000000000604482015290519081900360640190fd5b3073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1663ec5568896040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610aba57600080fd5b505af1158015610ace573d6000803e3d6000fd5b505050506040513d6020811015610ae457600080fd5b505173ffffffffffffffffffffffffffffffffffffffff1614610b6857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f57726f6e672070726f7879210000000000000000000000000000000000000000604482015290519081900360640190fd5b600254604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691638abf6077916004808301926020929190829003018186803b158015610bd357600080fd5b505afa158015610be7573d6000803e3d6000fd5b505050506040513d6020811015610bfd57600080fd5b5051604080517f62877ccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790523360448301529151929350600092918416916362877ccd913491606480830192602092919082900301818588803b158015610c8457600080fd5b505af1158015610c98573d6000803e3d6000fd5b50505050506040513d6020811015610caf57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff80821660008181526003602090815260408083208054958c167fffffffffffffffffffffffff000000000000000000000000000000000000000096871681179091556004805460018101825594527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9093018054909516841790945583519283529251939450927f812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763929181900390910190a260008055949350505050565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600360205260409020541690565b60048181548110610dbc57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60005415610e6e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600090815573ffffffffffffffffffffffffffffffffffffffff808316825260036020526040909120541680610f0757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f742072656769737465726564210000000000000000000000000000000000604482015290519081900360640190fd5b8173ffffffffffffffffffffffffffffffffffffffff1663e52253816040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610f4f57600080fd5b505af1158015610f63573d6000803e3d6000fd5b50505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016905550610fbb82611097565b505060008055565b60045460408051828152602080840282010190915260609190818015610ff3578160200160208202803883390190505b50915060005b81811015611072576004818154811061100e57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683828151811061104557fe5b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610ff9565b505090565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b3090565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260036020526040812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055600454905b81811015611236578273ffffffffffffffffffffffffffffffffffffffff166004828154811061111357fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16141561122e576004600183038154811061114b57fe5b6000918252602090912001546004805473ffffffffffffffffffffffffffffffffffffffff909216918390811061117e57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060048054806111d157fe5b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055611236565b6001016110e7565b50505056fea265627a7a723158204304b03a5396fe82ee304678afd1117363d8bf64e4ea1478069ae0ce99c2319a64736f6c63430005100032`

// DeployGovernedProxy deploys a new Ethereum contract, binding an instance of GovernedProxy to it.
func DeployGovernedProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _impl common.Address, _sporkProxy common.Address) (common.Address, *types.Transaction, *GovernedProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernedProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernedProxyBin), backend, _impl, _sporkProxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernedProxy{GovernedProxyCaller: GovernedProxyCaller{contract: contract}, GovernedProxyTransactor: GovernedProxyTransactor{contract: contract}, GovernedProxyFilterer: GovernedProxyFilterer{contract: contract}}, nil
}

// GovernedProxyBin is the compiled bytecode of contract after deployment.
const GovernedProxyRuntimeBin = `6080604052600436106100c65760003560e01c80638abf607711610074578063ce5494bb1161004e578063ce5494bb14610179578063dd6a851d146103c1578063ec556889146103d6576100c6565b80638abf607714610307578063a1b0e4761461031c578063b364595e1461035c576100c6565b80635b6dee4c116100a55780635b6dee4c146102645780636d5b6c441461029d5780636fa09ab0146102dd576100c6565b8062f55d9d146101795780630900f010146101bb57806332e3a905146101fb575b32331461013457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f6e6c79206469726563742063616c6c732061726520616c6c6f776564210000604482015290519081900360640190fd5b60015460405173ffffffffffffffffffffffffffffffffffffffff909116903660008237600080368334866127105a03f13d6000833e808015610175573d83f35b3d83fd5b34801561018557600080fd5b506101b96004803603602081101561019c57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166103eb565b005b3480156101c757600080fd5b506101b9600480360360208110156101de57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610452565b34801561020757600080fd5b5061023b6004803603602081101561021e57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166108c5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61023b6004803603604081101561027a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356108ed565b3480156102a957600080fd5b5061023b600480360360208110156102c057600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610d84565b3480156102e957600080fd5b5061023b6004803603602081101561030057600080fd5b5035610daf565b34801561031357600080fd5b5061023b610de3565b34801561032857600080fd5b506101b96004803603602081101561033f57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610dff565b34801561036857600080fd5b50610371610fc3565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103ad578181015183820152602001610395565b505050509050019250505060405180910390f35b3480156103cd57600080fd5b5061023b611077565b3480156103e257600080fd5b5061023b611093565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f476f6f6420747279000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600054156104c157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600081815573ffffffffffffffffffffffffffffffffffffffff808416825260036020526040909120549154918116911681141561056257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920616374697665210000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff81166105e457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f742072656769737465726564210000000000000000000000000000000000604482015290519081900360640190fd5b8173ffffffffffffffffffffffffffffffffffffffff16635051a5ec6040518163ffffffff1660e01b815260040160206040518083038186803b15801561062a57600080fd5b505afa15801561063e573d6000803e3d6000fd5b505050506040513d602081101561065457600080fd5b50516106c157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f742061636365707465642100000000000000000000000000000000000000604482015290519081900360640190fd5b600154604080517fce5494bb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff928316600482018190529151919284169163ce5494bb9160248082019260009290919082900301818387803b15801561073857600080fd5b505af115801561074c573d6000803e3d6000fd5b5050600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff868116918217909255604080517ef55d9d000000000000000000000000000000000000000000000000000000008152600481019290925251918516935062f55d9d925060248082019260009290919082900301818387803b1580156107ed57600080fd5b505af1158015610801573d6000803e3d6000fd5b5050505061080e83611097565b8273ffffffffffffffffffffffffffffffffffffffff166383197ef06040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561085657600080fd5b505af115801561086a573d6000803e3d6000fd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8781168252915191861693507f5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7925081900360200190a250506000805550565b60036020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b600032331461095d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f6e6c79206469726563742063616c6c732061726520616c6c6f776564210000604482015290519081900360640190fd5b600054156109cc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160008190555473ffffffffffffffffffffffffffffffffffffffff84811691161415610a5b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f416c726561647920616374697665210000000000000000000000000000000000604482015290519081900360640190fd5b3073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1663ec5568896040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610aba57600080fd5b505af1158015610ace573d6000803e3d6000fd5b505050506040513d6020811015610ae457600080fd5b505173ffffffffffffffffffffffffffffffffffffffff1614610b6857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f57726f6e672070726f7879210000000000000000000000000000000000000000604482015290519081900360640190fd5b600254604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691638abf6077916004808301926020929190829003018186803b158015610bd357600080fd5b505afa158015610be7573d6000803e3d6000fd5b505050506040513d6020811015610bfd57600080fd5b5051604080517f62877ccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790523360448301529151929350600092918416916362877ccd913491606480830192602092919082900301818588803b158015610c8457600080fd5b505af1158015610c98573d6000803e3d6000fd5b50505050506040513d6020811015610caf57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff80821660008181526003602090815260408083208054958c167fffffffffffffffffffffffff000000000000000000000000000000000000000096871681179091556004805460018101825594527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9093018054909516841790945583519283529251939450927f812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763929181900390910190a260008055949350505050565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600360205260409020541690565b60048181548110610dbc57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60005415610e6e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600090815573ffffffffffffffffffffffffffffffffffffffff808316825260036020526040909120541680610f0757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f742072656769737465726564210000000000000000000000000000000000604482015290519081900360640190fd5b8173ffffffffffffffffffffffffffffffffffffffff1663e52253816040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610f4f57600080fd5b505af1158015610f63573d6000803e3d6000fd5b50505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016905550610fbb82611097565b505060008055565b60045460408051828152602080840282010190915260609190818015610ff3578160200160208202803883390190505b50915060005b81811015611072576004818154811061100e57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683828151811061104557fe5b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610ff9565b505090565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b3090565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260036020526040812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055600454905b81811015611236578273ffffffffffffffffffffffffffffffffffffffff166004828154811061111357fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16141561122e576004600183038154811061114b57fe5b6000918252602090912001546004805473ffffffffffffffffffffffffffffffffffffffff909216918390811061117e57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060048054806111d157fe5b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055611236565b6001016110e7565b50505056fea265627a7a723158204304b03a5396fe82ee304678afd1117363d8bf64e4ea1478069ae0ce99c2319a64736f6c63430005100032`

// GovernedProxy is an auto generated Go binding around an Ethereum contract.
type GovernedProxy struct {
	GovernedProxyCaller     // Read-only binding to the contract
	GovernedProxyTransactor // Write-only binding to the contract
	GovernedProxyFilterer   // Log filterer for contract events
}

// GovernedProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernedProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernedProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernedProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernedProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernedProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernedProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernedProxySession struct {
	Contract     *GovernedProxy    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernedProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernedProxyCallerSession struct {
	Contract *GovernedProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GovernedProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernedProxyTransactorSession struct {
	Contract     *GovernedProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GovernedProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernedProxyRaw struct {
	Contract *GovernedProxy // Generic contract binding to access the raw methods on
}

// GovernedProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernedProxyCallerRaw struct {
	Contract *GovernedProxyCaller // Generic read-only contract binding to access the raw methods on
}

// GovernedProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernedProxyTransactorRaw struct {
	Contract *GovernedProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernedProxy creates a new instance of GovernedProxy, bound to a specific deployed contract.
func NewGovernedProxy(address common.Address, backend bind.ContractBackend) (*GovernedProxy, error) {
	contract, err := bindGovernedProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernedProxy{GovernedProxyCaller: GovernedProxyCaller{contract: contract}, GovernedProxyTransactor: GovernedProxyTransactor{contract: contract}, GovernedProxyFilterer: GovernedProxyFilterer{contract: contract}}, nil
}

// NewGovernedProxyCaller creates a new read-only instance of GovernedProxy, bound to a specific deployed contract.
func NewGovernedProxyCaller(address common.Address, caller bind.ContractCaller) (*GovernedProxyCaller, error) {
	contract, err := bindGovernedProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyCaller{contract: contract}, nil
}

// NewGovernedProxyTransactor creates a new write-only instance of GovernedProxy, bound to a specific deployed contract.
func NewGovernedProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernedProxyTransactor, error) {
	contract, err := bindGovernedProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyTransactor{contract: contract}, nil
}

// NewGovernedProxyFilterer creates a new log filterer instance of GovernedProxy, bound to a specific deployed contract.
func NewGovernedProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernedProxyFilterer, error) {
	contract, err := bindGovernedProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyFilterer{contract: contract}, nil
}

// bindGovernedProxy binds a generic wrapper to an already deployed contract.
func bindGovernedProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernedProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernedProxy *GovernedProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernedProxy.Contract.GovernedProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernedProxy *GovernedProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernedProxy.Contract.GovernedProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernedProxy *GovernedProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernedProxy.Contract.GovernedProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernedProxy *GovernedProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernedProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernedProxy *GovernedProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernedProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernedProxy *GovernedProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernedProxy.Contract.contract.Transact(opts, method, params...)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_GovernedProxy *GovernedProxyCaller) Impl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "impl")
	return *ret0, err
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_GovernedProxy *GovernedProxySession) Impl() (common.Address, error) {
	return _GovernedProxy.Contract.Impl(&_GovernedProxy.CallOpts)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_GovernedProxy *GovernedProxyCallerSession) Impl() (common.Address, error) {
	return _GovernedProxy.Contract.Impl(&_GovernedProxy.CallOpts)
}

// ListUpgradeProposals is a free data retrieval call binding the contract method 0xb364595e.
//
// Solidity: function listUpgradeProposals() constant returns(address[] proposals)
func (_GovernedProxy *GovernedProxyCaller) ListUpgradeProposals(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "listUpgradeProposals")
	return *ret0, err
}

// ListUpgradeProposals is a free data retrieval call binding the contract method 0xb364595e.
//
// Solidity: function listUpgradeProposals() constant returns(address[] proposals)
func (_GovernedProxy *GovernedProxySession) ListUpgradeProposals() ([]common.Address, error) {
	return _GovernedProxy.Contract.ListUpgradeProposals(&_GovernedProxy.CallOpts)
}

// ListUpgradeProposals is a free data retrieval call binding the contract method 0xb364595e.
//
// Solidity: function listUpgradeProposals() constant returns(address[] proposals)
func (_GovernedProxy *GovernedProxyCallerSession) ListUpgradeProposals() ([]common.Address, error) {
	return _GovernedProxy.Contract.ListUpgradeProposals(&_GovernedProxy.CallOpts)
}

// SporkProxy is a free data retrieval call binding the contract method 0xdd6a851d.
//
// Solidity: function spork_proxy() constant returns(address)
func (_GovernedProxy *GovernedProxyCaller) SporkProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "spork_proxy")
	return *ret0, err
}

// SporkProxy is a free data retrieval call binding the contract method 0xdd6a851d.
//
// Solidity: function spork_proxy() constant returns(address)
func (_GovernedProxy *GovernedProxySession) SporkProxy() (common.Address, error) {
	return _GovernedProxy.Contract.SporkProxy(&_GovernedProxy.CallOpts)
}

// SporkProxy is a free data retrieval call binding the contract method 0xdd6a851d.
//
// Solidity: function spork_proxy() constant returns(address)
func (_GovernedProxy *GovernedProxyCallerSession) SporkProxy() (common.Address, error) {
	return _GovernedProxy.Contract.SporkProxy(&_GovernedProxy.CallOpts)
}

// UpgradeProposalImpl is a free data retrieval call binding the contract method 0x6d5b6c44.
//
// Solidity: function upgradeProposalImpl(address _proposal) constant returns(address new_impl)
func (_GovernedProxy *GovernedProxyCaller) UpgradeProposalImpl(opts *bind.CallOpts, _proposal common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "upgradeProposalImpl", _proposal)
	return *ret0, err
}

// UpgradeProposalImpl is a free data retrieval call binding the contract method 0x6d5b6c44.
//
// Solidity: function upgradeProposalImpl(address _proposal) constant returns(address new_impl)
func (_GovernedProxy *GovernedProxySession) UpgradeProposalImpl(_proposal common.Address) (common.Address, error) {
	return _GovernedProxy.Contract.UpgradeProposalImpl(&_GovernedProxy.CallOpts, _proposal)
}

// UpgradeProposalImpl is a free data retrieval call binding the contract method 0x6d5b6c44.
//
// Solidity: function upgradeProposalImpl(address _proposal) constant returns(address new_impl)
func (_GovernedProxy *GovernedProxyCallerSession) UpgradeProposalImpl(_proposal common.Address) (common.Address, error) {
	return _GovernedProxy.Contract.UpgradeProposalImpl(&_GovernedProxy.CallOpts, _proposal)
}

// UpgradeProposalList is a free data retrieval call binding the contract method 0x6fa09ab0.
//
// Solidity: function upgrade_proposal_list(uint256 ) constant returns(address)
func (_GovernedProxy *GovernedProxyCaller) UpgradeProposalList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "upgrade_proposal_list", arg0)
	return *ret0, err
}

// UpgradeProposalList is a free data retrieval call binding the contract method 0x6fa09ab0.
//
// Solidity: function upgrade_proposal_list(uint256 ) constant returns(address)
func (_GovernedProxy *GovernedProxySession) UpgradeProposalList(arg0 *big.Int) (common.Address, error) {
	return _GovernedProxy.Contract.UpgradeProposalList(&_GovernedProxy.CallOpts, arg0)
}

// UpgradeProposalList is a free data retrieval call binding the contract method 0x6fa09ab0.
//
// Solidity: function upgrade_proposal_list(uint256 ) constant returns(address)
func (_GovernedProxy *GovernedProxyCallerSession) UpgradeProposalList(arg0 *big.Int) (common.Address, error) {
	return _GovernedProxy.Contract.UpgradeProposalList(&_GovernedProxy.CallOpts, arg0)
}

// UpgradeProposals is a free data retrieval call binding the contract method 0x32e3a905.
//
// Solidity: function upgrade_proposals(address ) constant returns(address)
func (_GovernedProxy *GovernedProxyCaller) UpgradeProposals(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernedProxy.contract.Call(opts, out, "upgrade_proposals", arg0)
	return *ret0, err
}

// UpgradeProposals is a free data retrieval call binding the contract method 0x32e3a905.
//
// Solidity: function upgrade_proposals(address ) constant returns(address)
func (_GovernedProxy *GovernedProxySession) UpgradeProposals(arg0 common.Address) (common.Address, error) {
	return _GovernedProxy.Contract.UpgradeProposals(&_GovernedProxy.CallOpts, arg0)
}

// UpgradeProposals is a free data retrieval call binding the contract method 0x32e3a905.
//
// Solidity: function upgrade_proposals(address ) constant returns(address)
func (_GovernedProxy *GovernedProxyCallerSession) UpgradeProposals(arg0 common.Address) (common.Address, error) {
	return _GovernedProxy.Contract.UpgradeProposals(&_GovernedProxy.CallOpts, arg0)
}

// CollectUpgradeProposal is a paid mutator transaction binding the contract method 0xa1b0e476.
//
// Solidity: function collectUpgradeProposal(address _proposal) returns()
func (_GovernedProxy *GovernedProxyTransactor) CollectUpgradeProposal(opts *bind.TransactOpts, _proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "collectUpgradeProposal", _proposal)
}

// CollectUpgradeProposal is a paid mutator transaction binding the contract method 0xa1b0e476.
//
// Solidity: function collectUpgradeProposal(address _proposal) returns()
func (_GovernedProxy *GovernedProxySession) CollectUpgradeProposal(_proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.CollectUpgradeProposal(&_GovernedProxy.TransactOpts, _proposal)
}

// CollectUpgradeProposal is a paid mutator transaction binding the contract method 0xa1b0e476.
//
// Solidity: function collectUpgradeProposal(address _proposal) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) CollectUpgradeProposal(_proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.CollectUpgradeProposal(&_GovernedProxy.TransactOpts, _proposal)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_GovernedProxy *GovernedProxyTransactor) Destroy(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "destroy", arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_GovernedProxy *GovernedProxySession) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Destroy(&_GovernedProxy.TransactOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address ) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) Destroy(arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Destroy(&_GovernedProxy.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_GovernedProxy *GovernedProxyTransactor) Migrate(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "migrate", arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_GovernedProxy *GovernedProxySession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Migrate(&_GovernedProxy.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Migrate(&_GovernedProxy.TransactOpts, arg0)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns(address)
func (_GovernedProxy *GovernedProxyTransactor) ProposeUpgrade(opts *bind.TransactOpts, _newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "proposeUpgrade", _newImpl, _period)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns(address)
func (_GovernedProxy *GovernedProxySession) ProposeUpgrade(_newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.Contract.ProposeUpgrade(&_GovernedProxy.TransactOpts, _newImpl, _period)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns(address)
func (_GovernedProxy *GovernedProxyTransactorSession) ProposeUpgrade(_newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _GovernedProxy.Contract.ProposeUpgrade(&_GovernedProxy.TransactOpts, _newImpl, _period)
}

// Proxy is a paid mutator transaction binding the contract method 0xec556889.
//
// Solidity: function proxy() returns(address)
func (_GovernedProxy *GovernedProxyTransactor) Proxy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "proxy")
}

// Proxy is a paid mutator transaction binding the contract method 0xec556889.
//
// Solidity: function proxy() returns(address)
func (_GovernedProxy *GovernedProxySession) Proxy() (*types.Transaction, error) {
	return _GovernedProxy.Contract.Proxy(&_GovernedProxy.TransactOpts)
}

// Proxy is a paid mutator transaction binding the contract method 0xec556889.
//
// Solidity: function proxy() returns(address)
func (_GovernedProxy *GovernedProxyTransactorSession) Proxy() (*types.Transaction, error) {
	return _GovernedProxy.Contract.Proxy(&_GovernedProxy.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_GovernedProxy *GovernedProxyTransactor) Upgrade(opts *bind.TransactOpts, _proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.contract.Transact(opts, "upgrade", _proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_GovernedProxy *GovernedProxySession) Upgrade(_proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Upgrade(&_GovernedProxy.TransactOpts, _proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_GovernedProxy *GovernedProxyTransactorSession) Upgrade(_proposal common.Address) (*types.Transaction, error) {
	return _GovernedProxy.Contract.Upgrade(&_GovernedProxy.TransactOpts, _proposal)
}

// GovernedProxyUpgradeProposalIterator is returned from FilterUpgradeProposal and is used to iterate over the raw logs and unpacked data for UpgradeProposal events raised by the GovernedProxy contract.
type GovernedProxyUpgradeProposalIterator struct {
	Event *GovernedProxyUpgradeProposal // Event containing the contract specifics and raw log

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
func (it *GovernedProxyUpgradeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernedProxyUpgradeProposal)
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
		it.Event = new(GovernedProxyUpgradeProposal)
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
func (it *GovernedProxyUpgradeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernedProxyUpgradeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernedProxyUpgradeProposal represents a UpgradeProposal event raised by the GovernedProxy contract.
type GovernedProxyUpgradeProposal struct {
	Impl     common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradeProposal is a free log retrieval operation binding the contract event 0x812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763.
//
// Solidity: event UpgradeProposal(address indexed impl, address proposal)
func (_GovernedProxy *GovernedProxyFilterer) FilterUpgradeProposal(opts *bind.FilterOpts, impl []common.Address) (*GovernedProxyUpgradeProposalIterator, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _GovernedProxy.contract.FilterLogs(opts, "UpgradeProposal", implRule)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyUpgradeProposalIterator{contract: _GovernedProxy.contract, event: "UpgradeProposal", logs: logs, sub: sub}, nil
}

// WatchUpgradeProposal is a free log subscription operation binding the contract event 0x812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763.
//
// Solidity: event UpgradeProposal(address indexed impl, address proposal)
func (_GovernedProxy *GovernedProxyFilterer) WatchUpgradeProposal(opts *bind.WatchOpts, sink chan<- *GovernedProxyUpgradeProposal, impl []common.Address) (event.Subscription, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _GovernedProxy.contract.WatchLogs(opts, "UpgradeProposal", implRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernedProxyUpgradeProposal)
				if err := _GovernedProxy.contract.UnpackLog(event, "UpgradeProposal", log); err != nil {
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

// GovernedProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GovernedProxy contract.
type GovernedProxyUpgradedIterator struct {
	Event *GovernedProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *GovernedProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernedProxyUpgraded)
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
		it.Event = new(GovernedProxyUpgraded)
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
func (it *GovernedProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernedProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernedProxyUpgraded represents a Upgraded event raised by the GovernedProxy contract.
type GovernedProxyUpgraded struct {
	Impl     common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed impl, address proposal)
func (_GovernedProxy *GovernedProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, impl []common.Address) (*GovernedProxyUpgradedIterator, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _GovernedProxy.contract.FilterLogs(opts, "Upgraded", implRule)
	if err != nil {
		return nil, err
	}
	return &GovernedProxyUpgradedIterator{contract: _GovernedProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed impl, address proposal)
func (_GovernedProxy *GovernedProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GovernedProxyUpgraded, impl []common.Address) (event.Subscription, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _GovernedProxy.contract.WatchLogs(opts, "Upgraded", implRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernedProxyUpgraded)
				if err := _GovernedProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
