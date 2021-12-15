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

// MasternodeTokenV2ABI is the input ABI used to generate the binding from.
const MasternodeTokenV2ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"contractIGovernedProxy\",\"name\":\"_registry_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_block\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositCollateral\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry_proxy\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"v1storage\",\"outputs\":[{\"internalType\":\"contractStorageMasternodeTokenV1\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MasternodeTokenV2Bin is the compiled bytecode used for deploying new contracts.
const MasternodeTokenV2Bin = `608060405234801561001057600080fd5b5060405161185a38038061185a8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b0319166001600160a01b03841617905560405182908290610067906100f5565b604051809103906000f080158015610083573d6000803e3d6000fd5b50600280546001600160a01b039283166001600160a01b0319918216179091556003805492841692909116919091179055604080516000808252915182917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a350505050610102565b6104208061143a83390190565b611329806101116000396000f3fe60806040526004361061010d5760003560e01c80636112fe2e116100a557806395d89b4111610074578063ce5494bb11610059578063ce5494bb1461046f578063dd62ed3e146104af578063ec556889146104f75761010d565b806395d89b411461045a578063a9059cbb146102405761010d565b80636112fe2e1461038f5780636f758140146103b957806370a08231146103c157806378c7d979146104015761010d565b806323b872dd116100e157806323b872dd146102c15780632d0593051461031157806330016a421461034f578063313ce567146103645761010d565b8062f55d9d1461017457806306fdde03146101b6578063095ea7b31461024057806318160ddd1461029a575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561018057600080fd5b506101b46004803603602081101561019757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661050c565b005b3480156101c257600080fd5b506101cb6105b4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102055781810151838201526020016101ed565b50505050905090810190601f1680156102325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024c57600080fd5b506102866004803603604081101561026357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356105ec565b604080519115158252519081900360200190f35b3480156102a657600080fd5b506102af610655565b60408051918252519081900360200190f35b3480156102cd57600080fd5b50610286600480360360608110156102e457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356105ec565b34801561031d57600080fd5b5061032661065a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561035b57600080fd5b50610326610676565b34801561037057600080fd5b50610379610692565b6040805160ff9092168252519081900360200190f35b34801561039b57600080fd5b506101b4600480360360208110156103b257600080fd5b5035610697565b6101b4610a88565b3480156103cd57600080fd5b506102af600480360360208110156103e457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610dc4565b34801561040d57600080fd5b506104416004803603602081101561042457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610e6d565b6040805192835260208301919091528051918290030190f35b34801561046657600080fd5b506101cb610f1d565b34801561047b57600080fd5b506101b46004803603602081101561049257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610f54565b3480156104bb57600080fd5b506102af600480360360408110156104d257600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610fe6565b34801561050357600080fd5b50610326610fee565b60005473ffffffffffffffffffffffffffffffffffffffff16331461059257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61059b8161100a565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60408051808201909152601581527f4d61737465726e6f646520436f6c6c61746572616c000000000000000000000060208201525b90565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f7765640000000000000000000000000000000000000000006044820152905160009181900360640190fd5b303190565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b601290565b6001541561070657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556000610714611099565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80851660048301529151939450600093919092169163133264e7916024808301926020929190829003018186803b15801561078d57600080fd5b505afa1580156107a1573d6000803e3d6000fd5b505050506040513d60208110156107b757600080fd5b505190508281101561082a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f7420656e6f75676800000000000000000000000000000000000000000000604482015290519081900360640190fd5b829003610836816110c8565b600254604080517fe8a6a28900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590524360448301529151919092169163e8a6a28991606480830192600092919082900301818387803b1580156108b757600080fd5b505af11580156108cb573d6000803e3d6000fd5b50506040805186815290516000935073ffffffffffffffffffffffffffffffffffffffff861692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b15801561098557600080fd5b505afa158015610999573d6000803e3d6000fd5b505050506040513d60208110156109af57600080fd5b5051604080517fcdc7d4ad00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529151919092169163cdc7d4ad91602480830192600092919082900301818387803b158015610a2257600080fd5b505af1158015610a36573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff8516925085156108fc02915085906000818181858888f19350505050158015610a7d573d6000803e3d6000fd5b505060006001555050565b60015415610af757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556000610b05611099565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80851660048301529151939450600093919092169163133264e7916024808301926020929190829003018186803b158015610b7e57600080fd5b505afa158015610b92573d6000803e3d6000fd5b505050506040513d6020811015610ba857600080fd5b505134019050610bb7816110c8565b600254604080517fe8a6a28900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590524360448301529151919092169163e8a6a28991606480830192600092919082900301818387803b158015610c3857600080fd5b505af1158015610c4c573d6000803e3d6000fd5b505060408051348152905173ffffffffffffffffffffffffffffffffffffffff86169350600092507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610d0657600080fd5b505afa158015610d1a573d6000803e3d6000fd5b505050506040513d6020811015610d3057600080fd5b5051604080517fcdc7d4ad00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529151919092169163cdc7d4ad91602480830192600092919082900301818387803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b5050600060015550505050565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529151600093929092169163133264e791602480820192602092909190829003018186803b158015610e3b57600080fd5b505afa158015610e4f573d6000803e3d6000fd5b505050506040513d6020811015610e6557600080fd5b505192915050565b600254604080517f27e235e300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152825160009485949216926327e235e3926024808301939192829003018186803b158015610ee057600080fd5b505afa158015610ef4573d6000803e3d6000fd5b505050506040513d6040811015610f0a57600080fd5b5080516020909101519094909350915050565b60408051808201909152600481527f4d4e524700000000000000000000000000000000000000000000000000000000602082015290565b60005473ffffffffffffffffffffffffffffffffffffffff163314610fda57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610fe3816111b9565b50565b600092915050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600254604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b15801561107e57600080fd5b505af1158015611092573d6000803e3d6000fd5b5050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff163314156110c15750326105e9565b50336105e9565b69152d02c7e14af680000081111561114157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f546f6f206d756368000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b683635c9adc5dea00000810615610fe357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e6f742061206d756c7469706c65000000000000000000000000000000000000604482015290519081900360640190fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166341c0e1b56040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561122357600080fd5b505af1158015611237573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16632d0593056040518163ffffffff1660e01b815260040160206040518083038186803b15801561128157600080fd5b505afa158015611295573d6000803e3d6000fd5b505050506040513d60208110156112ab57600080fd5b5051600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790555056fea265627a7a723158205905081457b9ec909f2e836a5604261e77031781d35ade3d39d9962a7846397c64736f6c634300051000326080604052600080546001600160a01b031916331790556103fb806100256000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806327e235e31161005057806327e235e3146100e657806341c0e1b514610132578063e8a6a2891461013a57610067565b8063133264e71461006c57806313af4035146100b1575b600080fd5b61009f6004803603602081101561008257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610179565b60408051918252519081900360200190f35b6100e4600480360360208110156100c757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101a1565b005b610119600480360360208110156100fc57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661026e565b6040805192835260208301919091528051918290030190f35b6100e4610287565b6100e46004803603606081101561015057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060400135610310565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b60005473ffffffffffffffffffffffffffffffffffffffff16331461022757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6001602081905260009182526040909120805491015482565b60005473ffffffffffffffffffffffffffffffffffffffff16331461030d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b33ff5b60005473ffffffffffffffffffffffffffffffffffffffff16331461039657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff9092166000908152600160208190526040909120918255015556fea265627a7a723158204f7087aac63529a55dabfac7688c5f4cd0d61e297f1ec358aa8ae46e937743c964736f6c63430005100032`

// DeployMasternodeTokenV2 deploys a new Ethereum contract, binding an instance of MasternodeTokenV2 to it.
func DeployMasternodeTokenV2(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _registry_proxy common.Address) (common.Address, *types.Transaction, *MasternodeTokenV2, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeTokenV2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MasternodeTokenV2Bin), backend, _proxy, _registry_proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MasternodeTokenV2{MasternodeTokenV2Caller: MasternodeTokenV2Caller{contract: contract}, MasternodeTokenV2Transactor: MasternodeTokenV2Transactor{contract: contract}, MasternodeTokenV2Filterer: MasternodeTokenV2Filterer{contract: contract}}, nil
}

// MasternodeTokenV2Bin is the compiled bytecode of contract after deployment.
const MasternodeTokenV2RuntimeBin = `60806040526004361061010d5760003560e01c80636112fe2e116100a557806395d89b4111610074578063ce5494bb11610059578063ce5494bb1461046f578063dd62ed3e146104af578063ec556889146104f75761010d565b806395d89b411461045a578063a9059cbb146102405761010d565b80636112fe2e1461038f5780636f758140146103b957806370a08231146103c157806378c7d979146104015761010d565b806323b872dd116100e157806323b872dd146102c15780632d0593051461031157806330016a421461034f578063313ce567146103645761010d565b8062f55d9d1461017457806306fdde03146101b6578063095ea7b31461024057806318160ddd1461029a575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561018057600080fd5b506101b46004803603602081101561019757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661050c565b005b3480156101c257600080fd5b506101cb6105b4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102055781810151838201526020016101ed565b50505050905090810190601f1680156102325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024c57600080fd5b506102866004803603604081101561026357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356105ec565b604080519115158252519081900360200190f35b3480156102a657600080fd5b506102af610655565b60408051918252519081900360200190f35b3480156102cd57600080fd5b50610286600480360360608110156102e457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356105ec565b34801561031d57600080fd5b5061032661065a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561035b57600080fd5b50610326610676565b34801561037057600080fd5b50610379610692565b6040805160ff9092168252519081900360200190f35b34801561039b57600080fd5b506101b4600480360360208110156103b257600080fd5b5035610697565b6101b4610a88565b3480156103cd57600080fd5b506102af600480360360208110156103e457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610dc4565b34801561040d57600080fd5b506104416004803603602081101561042457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610e6d565b6040805192835260208301919091528051918290030190f35b34801561046657600080fd5b506101cb610f1d565b34801561047b57600080fd5b506101b46004803603602081101561049257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610f54565b3480156104bb57600080fd5b506102af600480360360408110156104d257600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610fe6565b34801561050357600080fd5b50610326610fee565b60005473ffffffffffffffffffffffffffffffffffffffff16331461059257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61059b8161100a565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60408051808201909152601581527f4d61737465726e6f646520436f6c6c61746572616c000000000000000000000060208201525b90565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f7765640000000000000000000000000000000000000000006044820152905160009181900360640190fd5b303190565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b601290565b6001541561070657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556000610714611099565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80851660048301529151939450600093919092169163133264e7916024808301926020929190829003018186803b15801561078d57600080fd5b505afa1580156107a1573d6000803e3d6000fd5b505050506040513d60208110156107b757600080fd5b505190508281101561082a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f7420656e6f75676800000000000000000000000000000000000000000000604482015290519081900360640190fd5b829003610836816110c8565b600254604080517fe8a6a28900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590524360448301529151919092169163e8a6a28991606480830192600092919082900301818387803b1580156108b757600080fd5b505af11580156108cb573d6000803e3d6000fd5b50506040805186815290516000935073ffffffffffffffffffffffffffffffffffffffff861692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b15801561098557600080fd5b505afa158015610999573d6000803e3d6000fd5b505050506040513d60208110156109af57600080fd5b5051604080517fcdc7d4ad00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529151919092169163cdc7d4ad91602480830192600092919082900301818387803b158015610a2257600080fd5b505af1158015610a36573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff8516925085156108fc02915085906000818181858888f19350505050158015610a7d573d6000803e3d6000fd5b505060006001555050565b60015415610af757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556000610b05611099565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80851660048301529151939450600093919092169163133264e7916024808301926020929190829003018186803b158015610b7e57600080fd5b505afa158015610b92573d6000803e3d6000fd5b505050506040513d6020811015610ba857600080fd5b505134019050610bb7816110c8565b600254604080517fe8a6a28900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590524360448301529151919092169163e8a6a28991606480830192600092919082900301818387803b158015610c3857600080fd5b505af1158015610c4c573d6000803e3d6000fd5b505060408051348152905173ffffffffffffffffffffffffffffffffffffffff86169350600092507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610d0657600080fd5b505afa158015610d1a573d6000803e3d6000fd5b505050506040513d6020811015610d3057600080fd5b5051604080517fcdc7d4ad00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529151919092169163cdc7d4ad91602480830192600092919082900301818387803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b5050600060015550505050565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529151600093929092169163133264e791602480820192602092909190829003018186803b158015610e3b57600080fd5b505afa158015610e4f573d6000803e3d6000fd5b505050506040513d6020811015610e6557600080fd5b505192915050565b600254604080517f27e235e300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152825160009485949216926327e235e3926024808301939192829003018186803b158015610ee057600080fd5b505afa158015610ef4573d6000803e3d6000fd5b505050506040513d6040811015610f0a57600080fd5b5080516020909101519094909350915050565b60408051808201909152600481527f4d4e524700000000000000000000000000000000000000000000000000000000602082015290565b60005473ffffffffffffffffffffffffffffffffffffffff163314610fda57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610fe3816111b9565b50565b600092915050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600254604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b15801561107e57600080fd5b505af1158015611092573d6000803e3d6000fd5b5050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff163314156110c15750326105e9565b50336105e9565b69152d02c7e14af680000081111561114157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f546f6f206d756368000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b683635c9adc5dea00000810615610fe357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e6f742061206d756c7469706c65000000000000000000000000000000000000604482015290519081900360640190fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166341c0e1b56040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561122357600080fd5b505af1158015611237573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16632d0593056040518163ffffffff1660e01b815260040160206040518083038186803b15801561128157600080fd5b505afa158015611295573d6000803e3d6000fd5b505050506040513d60208110156112ab57600080fd5b5051600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790555056fea265627a7a723158205905081457b9ec909f2e836a5604261e77031781d35ade3d39d9962a7846397c64736f6c63430005100032`

// MasternodeTokenV2 is an auto generated Go binding around an Ethereum contract.
type MasternodeTokenV2 struct {
	MasternodeTokenV2Caller     // Read-only binding to the contract
	MasternodeTokenV2Transactor // Write-only binding to the contract
	MasternodeTokenV2Filterer   // Log filterer for contract events
}

// MasternodeTokenV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasternodeTokenV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasternodeTokenV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasternodeTokenV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasternodeTokenV2Session struct {
	Contract     *MasternodeTokenV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MasternodeTokenV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasternodeTokenV2CallerSession struct {
	Contract *MasternodeTokenV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MasternodeTokenV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasternodeTokenV2TransactorSession struct {
	Contract     *MasternodeTokenV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MasternodeTokenV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasternodeTokenV2Raw struct {
	Contract *MasternodeTokenV2 // Generic contract binding to access the raw methods on
}

// MasternodeTokenV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasternodeTokenV2CallerRaw struct {
	Contract *MasternodeTokenV2Caller // Generic read-only contract binding to access the raw methods on
}

// MasternodeTokenV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasternodeTokenV2TransactorRaw struct {
	Contract *MasternodeTokenV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasternodeTokenV2 creates a new instance of MasternodeTokenV2, bound to a specific deployed contract.
func NewMasternodeTokenV2(address common.Address, backend bind.ContractBackend) (*MasternodeTokenV2, error) {
	contract, err := bindMasternodeTokenV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV2{MasternodeTokenV2Caller: MasternodeTokenV2Caller{contract: contract}, MasternodeTokenV2Transactor: MasternodeTokenV2Transactor{contract: contract}, MasternodeTokenV2Filterer: MasternodeTokenV2Filterer{contract: contract}}, nil
}

// NewMasternodeTokenV2Caller creates a new read-only instance of MasternodeTokenV2, bound to a specific deployed contract.
func NewMasternodeTokenV2Caller(address common.Address, caller bind.ContractCaller) (*MasternodeTokenV2Caller, error) {
	contract, err := bindMasternodeTokenV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV2Caller{contract: contract}, nil
}

// NewMasternodeTokenV2Transactor creates a new write-only instance of MasternodeTokenV2, bound to a specific deployed contract.
func NewMasternodeTokenV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MasternodeTokenV2Transactor, error) {
	contract, err := bindMasternodeTokenV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV2Transactor{contract: contract}, nil
}

// NewMasternodeTokenV2Filterer creates a new log filterer instance of MasternodeTokenV2, bound to a specific deployed contract.
func NewMasternodeTokenV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MasternodeTokenV2Filterer, error) {
	contract, err := bindMasternodeTokenV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV2Filterer{contract: contract}, nil
}

// bindMasternodeTokenV2 binds a generic wrapper to an already deployed contract.
func bindMasternodeTokenV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeTokenV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeTokenV2 *MasternodeTokenV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeTokenV2.Contract.MasternodeTokenV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeTokenV2 *MasternodeTokenV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.MasternodeTokenV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeTokenV2 *MasternodeTokenV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.MasternodeTokenV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeTokenV2 *MasternodeTokenV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeTokenV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MasternodeTokenV2.Contract.Allowance(&_MasternodeTokenV2.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MasternodeTokenV2.Contract.Allowance(&_MasternodeTokenV2.CallOpts, arg0, arg1)
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) BalanceInfo(opts *bind.CallOpts, _tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	ret := new(struct {
		Balance   *big.Int
		LastBlock *big.Int
	})
	out := ret
	err := _MasternodeTokenV2.contract.Call(opts, out, "balanceInfo", _tokenOwner)
	return *ret, err
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) BalanceInfo(_tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	return _MasternodeTokenV2.Contract.BalanceInfo(&_MasternodeTokenV2.CallOpts, _tokenOwner)
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) BalanceInfo(_tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	return _MasternodeTokenV2.Contract.BalanceInfo(&_MasternodeTokenV2.CallOpts, _tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MasternodeTokenV2.Contract.BalanceOf(&_MasternodeTokenV2.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MasternodeTokenV2.Contract.BalanceOf(&_MasternodeTokenV2.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Decimals() (uint8, error) {
	return _MasternodeTokenV2.Contract.Decimals(&_MasternodeTokenV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) Decimals() (uint8, error) {
	return _MasternodeTokenV2.Contract.Decimals(&_MasternodeTokenV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Name() (string, error) {
	return _MasternodeTokenV2.Contract.Name(&_MasternodeTokenV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) Name() (string, error) {
	return _MasternodeTokenV2.Contract.Name(&_MasternodeTokenV2.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Proxy() (common.Address, error) {
	return _MasternodeTokenV2.Contract.Proxy(&_MasternodeTokenV2.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) Proxy() (common.Address, error) {
	return _MasternodeTokenV2.Contract.Proxy(&_MasternodeTokenV2.CallOpts)
}

// RegistryProxy is a free data retrieval call binding the contract method 0x30016a42.
//
// Solidity: function registry_proxy() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) RegistryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "registry_proxy")
	return *ret0, err
}

// RegistryProxy is a free data retrieval call binding the contract method 0x30016a42.
//
// Solidity: function registry_proxy() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) RegistryProxy() (common.Address, error) {
	return _MasternodeTokenV2.Contract.RegistryProxy(&_MasternodeTokenV2.CallOpts)
}

// RegistryProxy is a free data retrieval call binding the contract method 0x30016a42.
//
// Solidity: function registry_proxy() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) RegistryProxy() (common.Address, error) {
	return _MasternodeTokenV2.Contract.RegistryProxy(&_MasternodeTokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Symbol() (string, error) {
	return _MasternodeTokenV2.Contract.Symbol(&_MasternodeTokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) Symbol() (string, error) {
	return _MasternodeTokenV2.Contract.Symbol(&_MasternodeTokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) TotalSupply() (*big.Int, error) {
	return _MasternodeTokenV2.Contract.TotalSupply(&_MasternodeTokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) TotalSupply() (*big.Int, error) {
	return _MasternodeTokenV2.Contract.TotalSupply(&_MasternodeTokenV2.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2Caller) V1storage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeTokenV2.contract.Call(opts, out, "v1storage")
	return *ret0, err
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) V1storage() (common.Address, error) {
	return _MasternodeTokenV2.Contract.V1storage(&_MasternodeTokenV2.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeTokenV2 *MasternodeTokenV2CallerSession) V1storage() (common.Address, error) {
	return _MasternodeTokenV2.Contract.V1storage(&_MasternodeTokenV2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2Transactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.Approve(&_MasternodeTokenV2.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.Approve(&_MasternodeTokenV2.TransactOpts, arg0, arg1)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_MasternodeTokenV2 *MasternodeTokenV2Transactor) DepositCollateral(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeTokenV2.contract.Transact(opts, "depositCollateral")
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_MasternodeTokenV2 *MasternodeTokenV2Session) DepositCollateral() (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.DepositCollateral(&_MasternodeTokenV2.TransactOpts)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorSession) DepositCollateral() (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.DepositCollateral(&_MasternodeTokenV2.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV2.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.Destroy(&_MasternodeTokenV2.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.Destroy(&_MasternodeTokenV2.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV2.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.Migrate(&_MasternodeTokenV2.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.Migrate(&_MasternodeTokenV2.TransactOpts, _oldImpl)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2Transactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.Transfer(&_MasternodeTokenV2.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.Transfer(&_MasternodeTokenV2.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2Transactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2Session) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.TransferFrom(&_MasternodeTokenV2.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.TransferFrom(&_MasternodeTokenV2.TransactOpts, arg0, arg1, arg2)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2Transactor) WithdrawCollateral(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.contract.Transact(opts, "withdrawCollateral", _amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2Session) WithdrawCollateral(_amount *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.WithdrawCollateral(&_MasternodeTokenV2.TransactOpts, _amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_MasternodeTokenV2 *MasternodeTokenV2TransactorSession) WithdrawCollateral(_amount *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV2.Contract.WithdrawCollateral(&_MasternodeTokenV2.TransactOpts, _amount)
}

// MasternodeTokenV2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MasternodeTokenV2 contract.
type MasternodeTokenV2ApprovalIterator struct {
	Event *MasternodeTokenV2Approval // Event containing the contract specifics and raw log

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
func (it *MasternodeTokenV2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasternodeTokenV2Approval)
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
		it.Event = new(MasternodeTokenV2Approval)
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
func (it *MasternodeTokenV2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasternodeTokenV2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasternodeTokenV2Approval represents a Approval event raised by the MasternodeTokenV2 contract.
type MasternodeTokenV2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MasternodeTokenV2 *MasternodeTokenV2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MasternodeTokenV2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MasternodeTokenV2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV2ApprovalIterator{contract: _MasternodeTokenV2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MasternodeTokenV2 *MasternodeTokenV2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MasternodeTokenV2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MasternodeTokenV2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasternodeTokenV2Approval)
				if err := _MasternodeTokenV2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MasternodeTokenV2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MasternodeTokenV2 contract.
type MasternodeTokenV2TransferIterator struct {
	Event *MasternodeTokenV2Transfer // Event containing the contract specifics and raw log

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
func (it *MasternodeTokenV2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasternodeTokenV2Transfer)
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
		it.Event = new(MasternodeTokenV2Transfer)
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
func (it *MasternodeTokenV2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasternodeTokenV2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasternodeTokenV2Transfer represents a Transfer event raised by the MasternodeTokenV2 contract.
type MasternodeTokenV2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MasternodeTokenV2 *MasternodeTokenV2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MasternodeTokenV2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MasternodeTokenV2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV2TransferIterator{contract: _MasternodeTokenV2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MasternodeTokenV2 *MasternodeTokenV2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MasternodeTokenV2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MasternodeTokenV2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasternodeTokenV2Transfer)
				if err := _MasternodeTokenV2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
