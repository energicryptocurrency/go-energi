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

// MasternodeTokenV1ABI is the input ABI used to generate the binding from.
const MasternodeTokenV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"v1storage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry_proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositCollateral\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceInfo\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"last_block\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"},{\"name\":\"_registry_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// MasternodeTokenV1Bin is the compiled bytecode used for deploying new contracts.
const MasternodeTokenV1Bin = `608060405234801561001057600080fd5b506040516117163803806117168339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b0319166001600160a01b038416179055604051610063906100ef565b604051809103906000f08015801561007f573d6000803e3d6000fd5b50600280546001600160a01b039283166001600160a01b0319918216179091556003805492841692909116919091179055604080516000808252915182917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a350506100fc565b610420806112f683390190565b6111eb8061010b6000396000f3fe60806040526004361061010d5760003560e01c80636112fe2e116100a557806395d89b4111610074578063ce5494bb11610059578063ce5494bb1461046f578063dd62ed3e146104af578063ec556889146104f75761010d565b806395d89b411461045a578063a9059cbb146102405761010d565b80636112fe2e1461038f5780636f758140146103b957806370a08231146103c157806378c7d979146104015761010d565b806323b872dd116100e157806323b872dd146102c15780632d0593051461031157806330016a421461034f578063313ce567146103645761010d565b8062f55d9d1461017457806306fdde03146101b6578063095ea7b31461024057806318160ddd1461029a575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561018057600080fd5b506101b46004803603602081101561019757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661050c565b005b3480156101c257600080fd5b506101cb6105b4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102055781810151838201526020016101ed565b50505050905090810190601f1680156102325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024c57600080fd5b506102866004803603604081101561026357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356105ec565b604080519115158252519081900360200190f35b3480156102a657600080fd5b506102af610655565b60408051918252519081900360200190f35b3480156102cd57600080fd5b50610286600480360360608110156102e457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356105ec565b34801561031d57600080fd5b5061032661065a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561035b57600080fd5b50610326610676565b34801561037057600080fd5b50610379610692565b6040805160ff9092168252519081900360200190f35b34801561039b57600080fd5b506101b4600480360360208110156103b257600080fd5b5035610697565b6101b4610a88565b3480156103cd57600080fd5b506102af600480360360208110156103e457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610dc4565b34801561040d57600080fd5b506104416004803603602081101561042457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610e6d565b6040805192835260208301919091528051918290030190f35b34801561046657600080fd5b506101cb610f1d565b34801561047b57600080fd5b506101b46004803603602081101561049257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610f54565b3480156104bb57600080fd5b506102af600480360360408110156104d257600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610fe2565b34801561050357600080fd5b50610326610fea565b60005473ffffffffffffffffffffffffffffffffffffffff16331461059257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61059b81611006565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60408051808201909152601581527f4d61737465726e6f646520436f6c6c61746572616c000000000000000000000060208201525b90565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f7765640000000000000000000000000000000000000000006044820152905160009181900360640190fd5b303190565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b601690565b6001541561070657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556000610714611095565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80851660048301529151939450600093919092169163133264e7916024808301926020929190829003018186803b15801561078d57600080fd5b505afa1580156107a1573d6000803e3d6000fd5b505050506040513d60208110156107b757600080fd5b505190508281101561082a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f7420656e6f75676800000000000000000000000000000000000000000000604482015290519081900360640190fd5b829003610836816110c4565b600254604080517fe8a6a28900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590524360448301529151919092169163e8a6a28991606480830192600092919082900301818387803b1580156108b757600080fd5b505af11580156108cb573d6000803e3d6000fd5b50506040805186815290516000935073ffffffffffffffffffffffffffffffffffffffff861692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b15801561098557600080fd5b505afa158015610999573d6000803e3d6000fd5b505050506040513d60208110156109af57600080fd5b5051604080517fcdc7d4ad00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529151919092169163cdc7d4ad91602480830192600092919082900301818387803b158015610a2257600080fd5b505af1158015610a36573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff8516925085156108fc02915085906000818181858888f19350505050158015610a7d573d6000803e3d6000fd5b505060006001555050565b60015415610af757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556000610b05611095565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80851660048301529151939450600093919092169163133264e7916024808301926020929190829003018186803b158015610b7e57600080fd5b505afa158015610b92573d6000803e3d6000fd5b505050506040513d6020811015610ba857600080fd5b505134019050610bb7816110c4565b600254604080517fe8a6a28900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590524360448301529151919092169163e8a6a28991606480830192600092919082900301818387803b158015610c3857600080fd5b505af1158015610c4c573d6000803e3d6000fd5b505060408051348152905173ffffffffffffffffffffffffffffffffffffffff86169350600092507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610d0657600080fd5b505afa158015610d1a573d6000803e3d6000fd5b505050506040513d6020811015610d3057600080fd5b5051604080517fcdc7d4ad00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529151919092169163cdc7d4ad91602480830192600092919082900301818387803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b5050600060015550505050565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529151600093929092169163133264e791602480820192602092909190829003018186803b158015610e3b57600080fd5b505afa158015610e4f573d6000803e3d6000fd5b505050506040513d6020811015610e6557600080fd5b505192915050565b600254604080517f27e235e300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152825160009485949216926327e235e3926024808301939192829003018186803b158015610ee057600080fd5b505afa158015610ef4573d6000803e3d6000fd5b505050506040513d6040811015610f0a57600080fd5b5080516020909101519094909350915050565b60408051808201909152600481527f4d4e475200000000000000000000000000000000000000000000000000000000602082015290565b60005473ffffffffffffffffffffffffffffffffffffffff163314610fda57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610fdf815b50565b600092915050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600254604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b15801561107a57600080fd5b505af115801561108e573d6000803e3d6000fd5b5050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff163314156110bd5750326105e9565b50336105e9565b69152d02c7e14af680000081111561113d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f546f6f206d756368000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b69021e19e0c9bab2400000810615610fdf57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e6f742061206d756c7469706c65000000000000000000000000000000000000604482015290519081900360640190fdfea265627a7a7230582082290b4a4c6fbcafccc0ba92203782cb13d1b71b54d7587a686f57c53b0f859064736f6c634300050900326080604052600080546001600160a01b031916331790556103fb806100256000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806327e235e31161005057806327e235e3146100e657806341c0e1b514610132578063e8a6a2891461013a57610067565b8063133264e71461006c57806313af4035146100b1575b600080fd5b61009f6004803603602081101561008257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610179565b60408051918252519081900360200190f35b6100e4600480360360208110156100c757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101a1565b005b610119600480360360208110156100fc57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661026e565b6040805192835260208301919091528051918290030190f35b6100e4610287565b6100e46004803603606081101561015057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060400135610310565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b60005473ffffffffffffffffffffffffffffffffffffffff16331461022757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6001602081905260009182526040909120805491015482565b60005473ffffffffffffffffffffffffffffffffffffffff16331461030d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b33ff5b60005473ffffffffffffffffffffffffffffffffffffffff16331461039657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff9092166000908152600160208190526040909120918255015556fea265627a7a723058209702fb55b235225c437fc3b1f8dbd2e90528a380070321a8ee05a6762434add464736f6c63430005090032`

// DeployMasternodeTokenV1 deploys a new Ethereum contract, binding an instance of MasternodeTokenV1 to it.
func DeployMasternodeTokenV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _registry_proxy common.Address) (common.Address, *types.Transaction, *MasternodeTokenV1, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeTokenV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MasternodeTokenV1Bin), backend, _proxy, _registry_proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MasternodeTokenV1{MasternodeTokenV1Caller: MasternodeTokenV1Caller{contract: contract}, MasternodeTokenV1Transactor: MasternodeTokenV1Transactor{contract: contract}, MasternodeTokenV1Filterer: MasternodeTokenV1Filterer{contract: contract}}, nil
}

// MasternodeTokenV1Bin is the compiled bytecode of contract after deployment.
const MasternodeTokenV1RuntimeBin = `60806040526004361061010d5760003560e01c80636112fe2e116100a557806395d89b4111610074578063ce5494bb11610059578063ce5494bb1461046f578063dd62ed3e146104af578063ec556889146104f75761010d565b806395d89b411461045a578063a9059cbb146102405761010d565b80636112fe2e1461038f5780636f758140146103b957806370a08231146103c157806378c7d979146104015761010d565b806323b872dd116100e157806323b872dd146102c15780632d0593051461031157806330016a421461034f578063313ce567146103645761010d565b8062f55d9d1461017457806306fdde03146101b6578063095ea7b31461024057806318160ddd1461029a575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561018057600080fd5b506101b46004803603602081101561019757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661050c565b005b3480156101c257600080fd5b506101cb6105b4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102055781810151838201526020016101ed565b50505050905090810190601f1680156102325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024c57600080fd5b506102866004803603604081101561026357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356105ec565b604080519115158252519081900360200190f35b3480156102a657600080fd5b506102af610655565b60408051918252519081900360200190f35b3480156102cd57600080fd5b50610286600480360360608110156102e457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356105ec565b34801561031d57600080fd5b5061032661065a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561035b57600080fd5b50610326610676565b34801561037057600080fd5b50610379610692565b6040805160ff9092168252519081900360200190f35b34801561039b57600080fd5b506101b4600480360360208110156103b257600080fd5b5035610697565b6101b4610a88565b3480156103cd57600080fd5b506102af600480360360208110156103e457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610dc4565b34801561040d57600080fd5b506104416004803603602081101561042457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610e6d565b6040805192835260208301919091528051918290030190f35b34801561046657600080fd5b506101cb610f1d565b34801561047b57600080fd5b506101b46004803603602081101561049257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610f54565b3480156104bb57600080fd5b506102af600480360360408110156104d257600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610fe2565b34801561050357600080fd5b50610326610fea565b60005473ffffffffffffffffffffffffffffffffffffffff16331461059257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61059b81611006565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60408051808201909152601581527f4d61737465726e6f646520436f6c6c61746572616c000000000000000000000060208201525b90565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f7765640000000000000000000000000000000000000000006044820152905160009181900360640190fd5b303190565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b601690565b6001541561070657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556000610714611095565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80851660048301529151939450600093919092169163133264e7916024808301926020929190829003018186803b15801561078d57600080fd5b505afa1580156107a1573d6000803e3d6000fd5b505050506040513d60208110156107b757600080fd5b505190508281101561082a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f7420656e6f75676800000000000000000000000000000000000000000000604482015290519081900360640190fd5b829003610836816110c4565b600254604080517fe8a6a28900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590524360448301529151919092169163e8a6a28991606480830192600092919082900301818387803b1580156108b757600080fd5b505af11580156108cb573d6000803e3d6000fd5b50506040805186815290516000935073ffffffffffffffffffffffffffffffffffffffff861692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b15801561098557600080fd5b505afa158015610999573d6000803e3d6000fd5b505050506040513d60208110156109af57600080fd5b5051604080517fcdc7d4ad00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529151919092169163cdc7d4ad91602480830192600092919082900301818387803b158015610a2257600080fd5b505af1158015610a36573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff8516925085156108fc02915085906000818181858888f19350505050158015610a7d573d6000803e3d6000fd5b505060006001555050565b60015415610af757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600180556000610b05611095565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80851660048301529151939450600093919092169163133264e7916024808301926020929190829003018186803b158015610b7e57600080fd5b505afa158015610b92573d6000803e3d6000fd5b505050506040513d6020811015610ba857600080fd5b505134019050610bb7816110c4565b600254604080517fe8a6a28900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590524360448301529151919092169163e8a6a28991606480830192600092919082900301818387803b158015610c3857600080fd5b505af1158015610c4c573d6000803e3d6000fd5b505060408051348152905173ffffffffffffffffffffffffffffffffffffffff86169350600092507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610d0657600080fd5b505afa158015610d1a573d6000803e3d6000fd5b505050506040513d6020811015610d3057600080fd5b5051604080517fcdc7d4ad00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529151919092169163cdc7d4ad91602480830192600092919082900301818387803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b5050600060015550505050565b600254604080517f133264e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529151600093929092169163133264e791602480820192602092909190829003018186803b158015610e3b57600080fd5b505afa158015610e4f573d6000803e3d6000fd5b505050506040513d6020811015610e6557600080fd5b505192915050565b600254604080517f27e235e300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152825160009485949216926327e235e3926024808301939192829003018186803b158015610ee057600080fd5b505afa158015610ef4573d6000803e3d6000fd5b505050506040513d6040811015610f0a57600080fd5b5080516020909101519094909350915050565b60408051808201909152600481527f4d4e475200000000000000000000000000000000000000000000000000000000602082015290565b60005473ffffffffffffffffffffffffffffffffffffffff163314610fda57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610fdf815b50565b600092915050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600254604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b15801561107a57600080fd5b505af115801561108e573d6000803e3d6000fd5b5050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff163314156110bd5750326105e9565b50336105e9565b69152d02c7e14af680000081111561113d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f546f6f206d756368000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b69021e19e0c9bab2400000810615610fdf57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e6f742061206d756c7469706c65000000000000000000000000000000000000604482015290519081900360640190fdfea265627a7a7230582082290b4a4c6fbcafccc0ba92203782cb13d1b71b54d7587a686f57c53b0f859064736f6c63430005090032`

// MasternodeTokenV1 is an auto generated Go binding around an Ethereum contract.
type MasternodeTokenV1 struct {
	MasternodeTokenV1Caller     // Read-only binding to the contract
	MasternodeTokenV1Transactor // Write-only binding to the contract
	MasternodeTokenV1Filterer   // Log filterer for contract events
}

// MasternodeTokenV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasternodeTokenV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasternodeTokenV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasternodeTokenV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasternodeTokenV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasternodeTokenV1Session struct {
	Contract     *MasternodeTokenV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MasternodeTokenV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasternodeTokenV1CallerSession struct {
	Contract *MasternodeTokenV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MasternodeTokenV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasternodeTokenV1TransactorSession struct {
	Contract     *MasternodeTokenV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MasternodeTokenV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasternodeTokenV1Raw struct {
	Contract *MasternodeTokenV1 // Generic contract binding to access the raw methods on
}

// MasternodeTokenV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasternodeTokenV1CallerRaw struct {
	Contract *MasternodeTokenV1Caller // Generic read-only contract binding to access the raw methods on
}

// MasternodeTokenV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasternodeTokenV1TransactorRaw struct {
	Contract *MasternodeTokenV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasternodeTokenV1 creates a new instance of MasternodeTokenV1, bound to a specific deployed contract.
func NewMasternodeTokenV1(address common.Address, backend bind.ContractBackend) (*MasternodeTokenV1, error) {
	contract, err := bindMasternodeTokenV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1{MasternodeTokenV1Caller: MasternodeTokenV1Caller{contract: contract}, MasternodeTokenV1Transactor: MasternodeTokenV1Transactor{contract: contract}, MasternodeTokenV1Filterer: MasternodeTokenV1Filterer{contract: contract}}, nil
}

// NewMasternodeTokenV1Caller creates a new read-only instance of MasternodeTokenV1, bound to a specific deployed contract.
func NewMasternodeTokenV1Caller(address common.Address, caller bind.ContractCaller) (*MasternodeTokenV1Caller, error) {
	contract, err := bindMasternodeTokenV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1Caller{contract: contract}, nil
}

// NewMasternodeTokenV1Transactor creates a new write-only instance of MasternodeTokenV1, bound to a specific deployed contract.
func NewMasternodeTokenV1Transactor(address common.Address, transactor bind.ContractTransactor) (*MasternodeTokenV1Transactor, error) {
	contract, err := bindMasternodeTokenV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1Transactor{contract: contract}, nil
}

// NewMasternodeTokenV1Filterer creates a new log filterer instance of MasternodeTokenV1, bound to a specific deployed contract.
func NewMasternodeTokenV1Filterer(address common.Address, filterer bind.ContractFilterer) (*MasternodeTokenV1Filterer, error) {
	contract, err := bindMasternodeTokenV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1Filterer{contract: contract}, nil
}

// bindMasternodeTokenV1 binds a generic wrapper to an already deployed contract.
func bindMasternodeTokenV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasternodeTokenV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeTokenV1 *MasternodeTokenV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeTokenV1.Contract.MasternodeTokenV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeTokenV1 *MasternodeTokenV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.MasternodeTokenV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeTokenV1 *MasternodeTokenV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.MasternodeTokenV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasternodeTokenV1 *MasternodeTokenV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MasternodeTokenV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MasternodeTokenV1.Contract.Allowance(&_MasternodeTokenV1.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MasternodeTokenV1.Contract.Allowance(&_MasternodeTokenV1.CallOpts, arg0, arg1)
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) BalanceInfo(opts *bind.CallOpts, _tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	ret := new(struct {
		Balance   *big.Int
		LastBlock *big.Int
	})
	out := ret
	err := _MasternodeTokenV1.contract.Call(opts, out, "balanceInfo", _tokenOwner)
	return *ret, err
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) BalanceInfo(_tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	return _MasternodeTokenV1.Contract.BalanceInfo(&_MasternodeTokenV1.CallOpts, _tokenOwner)
}

// BalanceInfo is a free data retrieval call binding the contract method 0x78c7d979.
//
// Solidity: function balanceInfo(address _tokenOwner) constant returns(uint256 balance, uint256 last_block)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) BalanceInfo(_tokenOwner common.Address) (struct {
	Balance   *big.Int
	LastBlock *big.Int
}, error) {
	return _MasternodeTokenV1.Contract.BalanceInfo(&_MasternodeTokenV1.CallOpts, _tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MasternodeTokenV1.Contract.BalanceOf(&_MasternodeTokenV1.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MasternodeTokenV1.Contract.BalanceOf(&_MasternodeTokenV1.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Decimals() (uint8, error) {
	return _MasternodeTokenV1.Contract.Decimals(&_MasternodeTokenV1.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) Decimals() (uint8, error) {
	return _MasternodeTokenV1.Contract.Decimals(&_MasternodeTokenV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Name() (string, error) {
	return _MasternodeTokenV1.Contract.Name(&_MasternodeTokenV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) Name() (string, error) {
	return _MasternodeTokenV1.Contract.Name(&_MasternodeTokenV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Proxy() (common.Address, error) {
	return _MasternodeTokenV1.Contract.Proxy(&_MasternodeTokenV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) Proxy() (common.Address, error) {
	return _MasternodeTokenV1.Contract.Proxy(&_MasternodeTokenV1.CallOpts)
}

// RegistryProxy is a free data retrieval call binding the contract method 0x30016a42.
//
// Solidity: function registry_proxy() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) RegistryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "registry_proxy")
	return *ret0, err
}

// RegistryProxy is a free data retrieval call binding the contract method 0x30016a42.
//
// Solidity: function registry_proxy() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) RegistryProxy() (common.Address, error) {
	return _MasternodeTokenV1.Contract.RegistryProxy(&_MasternodeTokenV1.CallOpts)
}

// RegistryProxy is a free data retrieval call binding the contract method 0x30016a42.
//
// Solidity: function registry_proxy() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) RegistryProxy() (common.Address, error) {
	return _MasternodeTokenV1.Contract.RegistryProxy(&_MasternodeTokenV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Symbol() (string, error) {
	return _MasternodeTokenV1.Contract.Symbol(&_MasternodeTokenV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) Symbol() (string, error) {
	return _MasternodeTokenV1.Contract.Symbol(&_MasternodeTokenV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) TotalSupply() (*big.Int, error) {
	return _MasternodeTokenV1.Contract.TotalSupply(&_MasternodeTokenV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) TotalSupply() (*big.Int, error) {
	return _MasternodeTokenV1.Contract.TotalSupply(&_MasternodeTokenV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1Caller) V1storage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MasternodeTokenV1.contract.Call(opts, out, "v1storage")
	return *ret0, err
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) V1storage() (common.Address, error) {
	return _MasternodeTokenV1.Contract.V1storage(&_MasternodeTokenV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_MasternodeTokenV1 *MasternodeTokenV1CallerSession) V1storage() (common.Address, error) {
	return _MasternodeTokenV1.Contract.V1storage(&_MasternodeTokenV1.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Approve(&_MasternodeTokenV1.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Approve(&_MasternodeTokenV1.TransactOpts, arg0, arg1)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) DepositCollateral(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "depositCollateral")
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Session) DepositCollateral() (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.DepositCollateral(&_MasternodeTokenV1.TransactOpts)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x6f758140.
//
// Solidity: function depositCollateral() returns()
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) DepositCollateral() (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.DepositCollateral(&_MasternodeTokenV1.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Destroy(&_MasternodeTokenV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Destroy(&_MasternodeTokenV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Migrate(&_MasternodeTokenV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Migrate(&_MasternodeTokenV1.TransactOpts, _oldImpl)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Transfer(&_MasternodeTokenV1.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.Transfer(&_MasternodeTokenV1.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1Session) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.TransferFrom(&_MasternodeTokenV1.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.TransferFrom(&_MasternodeTokenV1.TransactOpts, arg0, arg1, arg2)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Transactor) WithdrawCollateral(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.contract.Transact(opts, "withdrawCollateral", _amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1Session) WithdrawCollateral(_amount *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.WithdrawCollateral(&_MasternodeTokenV1.TransactOpts, _amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 _amount) returns()
func (_MasternodeTokenV1 *MasternodeTokenV1TransactorSession) WithdrawCollateral(_amount *big.Int) (*types.Transaction, error) {
	return _MasternodeTokenV1.Contract.WithdrawCollateral(&_MasternodeTokenV1.TransactOpts, _amount)
}

// MasternodeTokenV1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MasternodeTokenV1 contract.
type MasternodeTokenV1ApprovalIterator struct {
	Event *MasternodeTokenV1Approval // Event containing the contract specifics and raw log

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
func (it *MasternodeTokenV1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasternodeTokenV1Approval)
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
		it.Event = new(MasternodeTokenV1Approval)
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
func (it *MasternodeTokenV1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasternodeTokenV1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasternodeTokenV1Approval represents a Approval event raised by the MasternodeTokenV1 contract.
type MasternodeTokenV1Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MasternodeTokenV1 *MasternodeTokenV1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MasternodeTokenV1ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MasternodeTokenV1.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1ApprovalIterator{contract: _MasternodeTokenV1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MasternodeTokenV1 *MasternodeTokenV1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MasternodeTokenV1Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MasternodeTokenV1.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasternodeTokenV1Approval)
				if err := _MasternodeTokenV1.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MasternodeTokenV1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MasternodeTokenV1 contract.
type MasternodeTokenV1TransferIterator struct {
	Event *MasternodeTokenV1Transfer // Event containing the contract specifics and raw log

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
func (it *MasternodeTokenV1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasternodeTokenV1Transfer)
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
		it.Event = new(MasternodeTokenV1Transfer)
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
func (it *MasternodeTokenV1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasternodeTokenV1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasternodeTokenV1Transfer represents a Transfer event raised by the MasternodeTokenV1 contract.
type MasternodeTokenV1Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MasternodeTokenV1 *MasternodeTokenV1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MasternodeTokenV1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MasternodeTokenV1.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MasternodeTokenV1TransferIterator{contract: _MasternodeTokenV1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MasternodeTokenV1 *MasternodeTokenV1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MasternodeTokenV1Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MasternodeTokenV1.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasternodeTokenV1Transfer)
				if err := _MasternodeTokenV1.contract.UnpackLog(event, "Transfer", log); err != nil {
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
