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

// GenericProposalV2ABI is the input ABI used to generate the binding from.
const GenericProposalV2ABI = "[{\"inputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"_mnregistry_proxy\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_quorum\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_supermajority\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"_feePayer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"accepted_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"canVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created_block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_payer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finish_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAccepted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mnregistry_proxy\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quorum_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rejected_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setFee\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supermajority\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteAccept\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteReject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericProposalV2Bin is the compiled bytecode used for deploying new contracts.
const GenericProposalV2Bin = `60806040523480156200001157600080fd5b506040516200148d3803806200148d833981810160405260a08110156200003757600080fd5b508051602080830151604080850151606086015160809096015160018054336001600160a01b031991821617909155436002556000805482166001600160a01b03808a16918217835542860160035560048054909416908516178355600b805460ff191660ff8c1617905585517f8abf607700000000000000000000000000000000000000000000000000000000815295519899969894979395919485949193638abf60779380820193929190829003018186803b158015620000f957600080fd5b505afa1580156200010e573d6000803e3d6000fd5b505050506040513d60208110156200012557600080fd5b5051604080517f06661abd00000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216916306661abd9160048082019260a092909190829003018186803b1580156200018357600080fd5b505afa15801562000198573d6000803e3d6000fd5b505050506040513d60a0811015620001af57600080fd5b5060408101516080909101519092509050806200022d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4e6f7420726561647920666f722070726f706f73616c73000000000000000000604482015290519081900360640190fd5b60028104821015620002a057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f41637469766520776569676874203c20312f3220657665722077656967687400604482015290519081900360640190fd5b600160ff871610156200031457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f51756f72756d206d696e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b606460ff871611156200038857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f51756f72756d206d617800000000000000000000000000000000000000000000604482015290519081900360640190fd5b6008829055606460ff8716830204600955603360ff871610620003b157600954600a55620003bc565b60646033830204600a555b6000600954116200042e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f51756f72756d2077656967687400000000000000000000000000000000000000604482015290519081900360640190fd5b6000600a5411620004a057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f46696e6973682077656967687400000000000000000000000000000000000000604482015290519081900360640190fd5b50505050505050610fd680620004b76000396000f3fe6080604052600436106101755760003560e01c80637b352962116100cb578063aec2ccae1161007f578063c86e6c1511610059578063c86e6c151461041b578063e522538114610430578063fe7334e81461044557610175565b8063aec2ccae146103b1578063c2472ef8146103f1578063c40a70f81461040657610175565b806391840a6b116100b057806391840a6b14610347578063990a663b1461035c578063adfaa72e1461037157610175565b80637b3529621461031d57806383197ef01461033257610175565b80635051a5ec1161012d57806360f96a8f1161010757806360f96a8f146102b557806375df0f99146102f35780637639b1eb1461030857610175565b80635051a5ec1461024c57806356c2a0a1146102755780635c31f2201461028a57610175565b80632ded32271161015e5780632ded3227146102185780633ccfd60b146102225780633d1db3e91461023757610175565b80630b62be45146101dc57806329dcb0cf14610203575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f776564000000000000000000000000000000000000000000604482015290519081900360640190fd5b3480156101e857600080fd5b506101f161045a565b60408051918252519081900360200190f35b34801561020f57600080fd5b506101f1610460565b610220610466565b005b34801561022e57600080fd5b506102206104f6565b34801561024357600080fd5b506101f16105b3565b34801561025857600080fd5b506102616105b9565b604080519115158252519081900360200190f35b34801561028157600080fd5b50610220610612565b34801561029657600080fd5b5061029f610625565b6040805160ff9092168252519081900360200190f35b3480156102c157600080fd5b506102ca61062e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156102ff57600080fd5b506101f161064a565b34801561031457600080fd5b506101f1610650565b34801561032957600080fd5b50610261610656565b34801561033e57600080fd5b50610220610680565b34801561035357600080fd5b506101f1610721565b34801561036857600080fd5b506101f1610727565b34801561037d57600080fd5b506102616004803603602081101561039457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661072d565b3480156103bd57600080fd5b50610261600480360360208110156103d457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166108c7565b3480156103fd57600080fd5b506102206108dc565b34801561041257600080fd5b506102ca6108ef565b34801561042757600080fd5b506101f161090b565b34801561043c57600080fd5b50610220610911565b34801561045157600080fd5b506102ca610c66565b60025481565b60035481565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104ec57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6005805434019055565b6104fe6105b9565b61056957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742061636365707465640000000000000000000000000000000000000000604482015290519081900360640190fd5b60045460405173ffffffffffffffffffffffffffffffffffffffff90911690303180156108fc02916000818181858888f193505050501580156105b0573d6000803e3d6000fd5b50565b600a5481565b6000600a54600654106105ce5750600161060f565b6105d6610656565b6105e25750600061060f565b6009546007546006540110156105fa5750600061060f565b600b5460095460649160ff1602046006541190505b90565b61061a610c82565b600780549091019055565b600b5460ff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60065481565b60004260035411158061066d5750600a5460065410155b8061067b5750600a54600754115b905090565b60015473ffffffffffffffffffffffffffffffffffffffff16331461070657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60045473ffffffffffffffffffffffffffffffffffffffff16ff5b60085481565b60055481565b60008054604080517f8abf60770000000000000000000000000000000000000000000000000000000081529051839273ffffffffffffffffffffffffffffffffffffffff1691638abf6077916004808301926020929190829003018186803b15801561079857600080fd5b505afa1580156107ac573d6000803e3d6000fd5b505050506040513d60208110156107c257600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015291519293506000929184169163b83e16059160248082019260e092909190829003018186803b15801561083a57600080fd5b505afa15801561084e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e081101561087357600080fd5b5060a001516003549091504210801561088d575060025481105b80156108bf575073ffffffffffffffffffffffffffffffffffffffff84166000908152600c602052604090205460ff16155b949350505050565b600c6020526000908152604090205460ff1681565b6108e4610c82565b600680549091019055565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b610919610656565b801561092a57506109286105b9565b155b61099557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f7420636f6c6c65637461626c650000000000000000000000000000000000604482015290519081900360640190fd5b60015473ffffffffffffffffffffffffffffffffffffffff163314610a1b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610a8457600080fd5b505afa158015610a98573d6000803e3d6000fd5b505050506040513d6020811015610aae57600080fd5b5051604080517fa2731784000000000000000000000000000000000000000000000000000000008152905191925060009173ffffffffffffffffffffffffffffffffffffffff84169163a2731784916004808301926020929190829003018186803b158015610b1c57600080fd5b505afa158015610b30573d6000803e3d6000fd5b505050506040513d6020811015610b4657600080fd5b5051604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921691638abf607791600480820192602092909190829003018186803b158015610bb057600080fd5b505afa158015610bc4573d6000803e3d6000fd5b505050506040513d6020811015610bda57600080fd5b5051604080517fd7bb99ba000000000000000000000000000000000000000000000000000000008152905191925073ffffffffffffffffffffffffffffffffffffffff83169163d7bb99ba91303191600480830192600092919082900301818588803b158015610c4957600080fd5b505af1158015610c5d573d6000803e3d6000fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60004260035411610cf457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f46696e6973686564000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610d5d57600080fd5b505afa158015610d71573d6000803e3d6000fd5b505050506040513d6020811015610d8757600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815233600482018190529151929350909160009173ffffffffffffffffffffffffffffffffffffffff85169163b83e16059160248082019260e092909190829003018186803b158015610e0057600080fd5b505afa158015610e14573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e0811015610e3957600080fd5b50608081015160a09091015160025491955091508110610eba57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f7420656c696769626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600c602052604090205460ff1615610f4f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b5073ffffffffffffffffffffffffffffffffffffffff166000908152600c6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055509056fea265627a7a72315820aacb5dce7820da55069c0b3f57aad110313eabd5a295f0a57de5050dd0e877fe64736f6c63430005100032`

// DeployGenericProposalV2 deploys a new Ethereum contract, binding an instance of GenericProposalV2 to it.
func DeployGenericProposalV2(auth *bind.TransactOpts, backend bind.ContractBackend, _mnregistry_proxy common.Address, _quorum uint8, _period *big.Int, _supermajority uint8, _feePayer common.Address) (common.Address, *types.Transaction, *GenericProposalV2, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericProposalV2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericProposalV2Bin), backend, _mnregistry_proxy, _quorum, _period, _supermajority, _feePayer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericProposalV2{GenericProposalV2Caller: GenericProposalV2Caller{contract: contract}, GenericProposalV2Transactor: GenericProposalV2Transactor{contract: contract}, GenericProposalV2Filterer: GenericProposalV2Filterer{contract: contract}}, nil
}

// GenericProposalV2Bin is the compiled bytecode of contract after deployment.
const GenericProposalV2RuntimeBin = `6080604052600436106101755760003560e01c80637b352962116100cb578063aec2ccae1161007f578063c86e6c1511610059578063c86e6c151461041b578063e522538114610430578063fe7334e81461044557610175565b8063aec2ccae146103b1578063c2472ef8146103f1578063c40a70f81461040657610175565b806391840a6b116100b057806391840a6b14610347578063990a663b1461035c578063adfaa72e1461037157610175565b80637b3529621461031d57806383197ef01461033257610175565b80635051a5ec1161012d57806360f96a8f1161010757806360f96a8f146102b557806375df0f99146102f35780637639b1eb1461030857610175565b80635051a5ec1461024c57806356c2a0a1146102755780635c31f2201461028a57610175565b80632ded32271161015e5780632ded3227146102185780633ccfd60b146102225780633d1db3e91461023757610175565b80630b62be45146101dc57806329dcb0cf14610203575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f776564000000000000000000000000000000000000000000604482015290519081900360640190fd5b3480156101e857600080fd5b506101f161045a565b60408051918252519081900360200190f35b34801561020f57600080fd5b506101f1610460565b610220610466565b005b34801561022e57600080fd5b506102206104f6565b34801561024357600080fd5b506101f16105b3565b34801561025857600080fd5b506102616105b9565b604080519115158252519081900360200190f35b34801561028157600080fd5b50610220610612565b34801561029657600080fd5b5061029f610625565b6040805160ff9092168252519081900360200190f35b3480156102c157600080fd5b506102ca61062e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156102ff57600080fd5b506101f161064a565b34801561031457600080fd5b506101f1610650565b34801561032957600080fd5b50610261610656565b34801561033e57600080fd5b50610220610680565b34801561035357600080fd5b506101f1610721565b34801561036857600080fd5b506101f1610727565b34801561037d57600080fd5b506102616004803603602081101561039457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661072d565b3480156103bd57600080fd5b50610261600480360360208110156103d457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166108c7565b3480156103fd57600080fd5b506102206108dc565b34801561041257600080fd5b506102ca6108ef565b34801561042757600080fd5b506101f161090b565b34801561043c57600080fd5b50610220610911565b34801561045157600080fd5b506102ca610c66565b60025481565b60035481565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104ec57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6005805434019055565b6104fe6105b9565b61056957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742061636365707465640000000000000000000000000000000000000000604482015290519081900360640190fd5b60045460405173ffffffffffffffffffffffffffffffffffffffff90911690303180156108fc02916000818181858888f193505050501580156105b0573d6000803e3d6000fd5b50565b600a5481565b6000600a54600654106105ce5750600161060f565b6105d6610656565b6105e25750600061060f565b6009546007546006540110156105fa5750600061060f565b600b5460095460649160ff1602046006541190505b90565b61061a610c82565b600780549091019055565b600b5460ff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60065481565b60004260035411158061066d5750600a5460065410155b8061067b5750600a54600754115b905090565b60015473ffffffffffffffffffffffffffffffffffffffff16331461070657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60045473ffffffffffffffffffffffffffffffffffffffff16ff5b60085481565b60055481565b60008054604080517f8abf60770000000000000000000000000000000000000000000000000000000081529051839273ffffffffffffffffffffffffffffffffffffffff1691638abf6077916004808301926020929190829003018186803b15801561079857600080fd5b505afa1580156107ac573d6000803e3d6000fd5b505050506040513d60208110156107c257600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015291519293506000929184169163b83e16059160248082019260e092909190829003018186803b15801561083a57600080fd5b505afa15801561084e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e081101561087357600080fd5b5060a001516003549091504210801561088d575060025481105b80156108bf575073ffffffffffffffffffffffffffffffffffffffff84166000908152600c602052604090205460ff16155b949350505050565b600c6020526000908152604090205460ff1681565b6108e4610c82565b600680549091019055565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b610919610656565b801561092a57506109286105b9565b155b61099557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f7420636f6c6c65637461626c650000000000000000000000000000000000604482015290519081900360640190fd5b60015473ffffffffffffffffffffffffffffffffffffffff163314610a1b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610a8457600080fd5b505afa158015610a98573d6000803e3d6000fd5b505050506040513d6020811015610aae57600080fd5b5051604080517fa2731784000000000000000000000000000000000000000000000000000000008152905191925060009173ffffffffffffffffffffffffffffffffffffffff84169163a2731784916004808301926020929190829003018186803b158015610b1c57600080fd5b505afa158015610b30573d6000803e3d6000fd5b505050506040513d6020811015610b4657600080fd5b5051604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921691638abf607791600480820192602092909190829003018186803b158015610bb057600080fd5b505afa158015610bc4573d6000803e3d6000fd5b505050506040513d6020811015610bda57600080fd5b5051604080517fd7bb99ba000000000000000000000000000000000000000000000000000000008152905191925073ffffffffffffffffffffffffffffffffffffffff83169163d7bb99ba91303191600480830192600092919082900301818588803b158015610c4957600080fd5b505af1158015610c5d573d6000803e3d6000fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60004260035411610cf457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f46696e6973686564000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610d5d57600080fd5b505afa158015610d71573d6000803e3d6000fd5b505050506040513d6020811015610d8757600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815233600482018190529151929350909160009173ffffffffffffffffffffffffffffffffffffffff85169163b83e16059160248082019260e092909190829003018186803b158015610e0057600080fd5b505afa158015610e14573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e0811015610e3957600080fd5b50608081015160a09091015160025491955091508110610eba57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f7420656c696769626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600c602052604090205460ff1615610f4f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b5073ffffffffffffffffffffffffffffffffffffffff166000908152600c6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055509056fea265627a7a72315820aacb5dce7820da55069c0b3f57aad110313eabd5a295f0a57de5050dd0e877fe64736f6c63430005100032`

// GenericProposalV2 is an auto generated Go binding around an Ethereum contract.
type GenericProposalV2 struct {
	GenericProposalV2Caller     // Read-only binding to the contract
	GenericProposalV2Transactor // Write-only binding to the contract
	GenericProposalV2Filterer   // Log filterer for contract events
}

// GenericProposalV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type GenericProposalV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericProposalV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericProposalV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericProposalV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericProposalV2Session struct {
	Contract     *GenericProposalV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GenericProposalV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericProposalV2CallerSession struct {
	Contract *GenericProposalV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GenericProposalV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericProposalV2TransactorSession struct {
	Contract     *GenericProposalV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GenericProposalV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type GenericProposalV2Raw struct {
	Contract *GenericProposalV2 // Generic contract binding to access the raw methods on
}

// GenericProposalV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericProposalV2CallerRaw struct {
	Contract *GenericProposalV2Caller // Generic read-only contract binding to access the raw methods on
}

// GenericProposalV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericProposalV2TransactorRaw struct {
	Contract *GenericProposalV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericProposalV2 creates a new instance of GenericProposalV2, bound to a specific deployed contract.
func NewGenericProposalV2(address common.Address, backend bind.ContractBackend) (*GenericProposalV2, error) {
	contract, err := bindGenericProposalV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericProposalV2{GenericProposalV2Caller: GenericProposalV2Caller{contract: contract}, GenericProposalV2Transactor: GenericProposalV2Transactor{contract: contract}, GenericProposalV2Filterer: GenericProposalV2Filterer{contract: contract}}, nil
}

// NewGenericProposalV2Caller creates a new read-only instance of GenericProposalV2, bound to a specific deployed contract.
func NewGenericProposalV2Caller(address common.Address, caller bind.ContractCaller) (*GenericProposalV2Caller, error) {
	contract, err := bindGenericProposalV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericProposalV2Caller{contract: contract}, nil
}

// NewGenericProposalV2Transactor creates a new write-only instance of GenericProposalV2, bound to a specific deployed contract.
func NewGenericProposalV2Transactor(address common.Address, transactor bind.ContractTransactor) (*GenericProposalV2Transactor, error) {
	contract, err := bindGenericProposalV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericProposalV2Transactor{contract: contract}, nil
}

// NewGenericProposalV2Filterer creates a new log filterer instance of GenericProposalV2, bound to a specific deployed contract.
func NewGenericProposalV2Filterer(address common.Address, filterer bind.ContractFilterer) (*GenericProposalV2Filterer, error) {
	contract, err := bindGenericProposalV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericProposalV2Filterer{contract: contract}, nil
}

// bindGenericProposalV2 binds a generic wrapper to an already deployed contract.
func bindGenericProposalV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericProposalV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericProposalV2 *GenericProposalV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericProposalV2.Contract.GenericProposalV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericProposalV2 *GenericProposalV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV2.Contract.GenericProposalV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericProposalV2 *GenericProposalV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericProposalV2.Contract.GenericProposalV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericProposalV2 *GenericProposalV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericProposalV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericProposalV2 *GenericProposalV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericProposalV2 *GenericProposalV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericProposalV2.Contract.contract.Transact(opts, method, params...)
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Caller) AcceptedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "accepted_weight")
	return *ret0, err
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Session) AcceptedWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.AcceptedWeight(&_GenericProposalV2.CallOpts)
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2CallerSession) AcceptedWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.AcceptedWeight(&_GenericProposalV2.CallOpts)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2Caller) CanVote(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "canVote", owner)
	return *ret0, err
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2Session) CanVote(owner common.Address) (bool, error) {
	return _GenericProposalV2.Contract.CanVote(&_GenericProposalV2.CallOpts, owner)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2CallerSession) CanVote(owner common.Address) (bool, error) {
	return _GenericProposalV2.Contract.CanVote(&_GenericProposalV2.CallOpts, owner)
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Caller) CreatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "created_block")
	return *ret0, err
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Session) CreatedBlock() (*big.Int, error) {
	return _GenericProposalV2.Contract.CreatedBlock(&_GenericProposalV2.CallOpts)
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2CallerSession) CreatedBlock() (*big.Int, error) {
	return _GenericProposalV2.Contract.CreatedBlock(&_GenericProposalV2.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Caller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "deadline")
	return *ret0, err
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Session) Deadline() (*big.Int, error) {
	return _GenericProposalV2.Contract.Deadline(&_GenericProposalV2.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2CallerSession) Deadline() (*big.Int, error) {
	return _GenericProposalV2.Contract.Deadline(&_GenericProposalV2.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Caller) FeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "fee_amount")
	return *ret0, err
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Session) FeeAmount() (*big.Int, error) {
	return _GenericProposalV2.Contract.FeeAmount(&_GenericProposalV2.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2CallerSession) FeeAmount() (*big.Int, error) {
	return _GenericProposalV2.Contract.FeeAmount(&_GenericProposalV2.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2Caller) FeePayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "fee_payer")
	return *ret0, err
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2Session) FeePayer() (common.Address, error) {
	return _GenericProposalV2.Contract.FeePayer(&_GenericProposalV2.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2CallerSession) FeePayer() (common.Address, error) {
	return _GenericProposalV2.Contract.FeePayer(&_GenericProposalV2.CallOpts)
}

// FinishWeight is a free data retrieval call binding the contract method 0x3d1db3e9.
//
// Solidity: function finish_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Caller) FinishWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "finish_weight")
	return *ret0, err
}

// FinishWeight is a free data retrieval call binding the contract method 0x3d1db3e9.
//
// Solidity: function finish_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Session) FinishWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.FinishWeight(&_GenericProposalV2.CallOpts)
}

// FinishWeight is a free data retrieval call binding the contract method 0x3d1db3e9.
//
// Solidity: function finish_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2CallerSession) FinishWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.FinishWeight(&_GenericProposalV2.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2Caller) IsAccepted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "isAccepted")
	return *ret0, err
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2Session) IsAccepted() (bool, error) {
	return _GenericProposalV2.Contract.IsAccepted(&_GenericProposalV2.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2CallerSession) IsAccepted() (bool, error) {
	return _GenericProposalV2.Contract.IsAccepted(&_GenericProposalV2.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2Caller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2Session) IsFinished() (bool, error) {
	return _GenericProposalV2.Contract.IsFinished(&_GenericProposalV2.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2CallerSession) IsFinished() (bool, error) {
	return _GenericProposalV2.Contract.IsFinished(&_GenericProposalV2.CallOpts)
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2Caller) MnregistryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "mnregistry_proxy")
	return *ret0, err
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2Session) MnregistryProxy() (common.Address, error) {
	return _GenericProposalV2.Contract.MnregistryProxy(&_GenericProposalV2.CallOpts)
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2CallerSession) MnregistryProxy() (common.Address, error) {
	return _GenericProposalV2.Contract.MnregistryProxy(&_GenericProposalV2.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2Caller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "parent")
	return *ret0, err
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2Session) Parent() (common.Address, error) {
	return _GenericProposalV2.Contract.Parent(&_GenericProposalV2.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_GenericProposalV2 *GenericProposalV2CallerSession) Parent() (common.Address, error) {
	return _GenericProposalV2.Contract.Parent(&_GenericProposalV2.CallOpts)
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Caller) QuorumWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "quorum_weight")
	return *ret0, err
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Session) QuorumWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.QuorumWeight(&_GenericProposalV2.CallOpts)
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2CallerSession) QuorumWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.QuorumWeight(&_GenericProposalV2.CallOpts)
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Caller) RejectedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "rejected_weight")
	return *ret0, err
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Session) RejectedWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.RejectedWeight(&_GenericProposalV2.CallOpts)
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2CallerSession) RejectedWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.RejectedWeight(&_GenericProposalV2.CallOpts)
}

// Supermajority is a free data retrieval call binding the contract method 0x5c31f220.
//
// Solidity: function supermajority() constant returns(uint8)
func (_GenericProposalV2 *GenericProposalV2Caller) Supermajority(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "supermajority")
	return *ret0, err
}

// Supermajority is a free data retrieval call binding the contract method 0x5c31f220.
//
// Solidity: function supermajority() constant returns(uint8)
func (_GenericProposalV2 *GenericProposalV2Session) Supermajority() (uint8, error) {
	return _GenericProposalV2.Contract.Supermajority(&_GenericProposalV2.CallOpts)
}

// Supermajority is a free data retrieval call binding the contract method 0x5c31f220.
//
// Solidity: function supermajority() constant returns(uint8)
func (_GenericProposalV2 *GenericProposalV2CallerSession) Supermajority() (uint8, error) {
	return _GenericProposalV2.Contract.Supermajority(&_GenericProposalV2.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Caller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "total_weight")
	return *ret0, err
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2Session) TotalWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.TotalWeight(&_GenericProposalV2.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_GenericProposalV2 *GenericProposalV2CallerSession) TotalWeight() (*big.Int, error) {
	return _GenericProposalV2.Contract.TotalWeight(&_GenericProposalV2.CallOpts)
}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2Caller) Voted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GenericProposalV2.contract.Call(opts, out, "voted", arg0)
	return *ret0, err
}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2Session) Voted(arg0 common.Address) (bool, error) {
	return _GenericProposalV2.Contract.Voted(&_GenericProposalV2.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) constant returns(bool)
func (_GenericProposalV2 *GenericProposalV2CallerSession) Voted(arg0 common.Address) (bool, error) {
	return _GenericProposalV2.Contract.Voted(&_GenericProposalV2.CallOpts, arg0)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_GenericProposalV2 *GenericProposalV2Transactor) Collect(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV2.contract.Transact(opts, "collect")
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_GenericProposalV2 *GenericProposalV2Session) Collect() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.Collect(&_GenericProposalV2.TransactOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_GenericProposalV2 *GenericProposalV2TransactorSession) Collect() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.Collect(&_GenericProposalV2.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_GenericProposalV2 *GenericProposalV2Transactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV2.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_GenericProposalV2 *GenericProposalV2Session) Destroy() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.Destroy(&_GenericProposalV2.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_GenericProposalV2 *GenericProposalV2TransactorSession) Destroy() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.Destroy(&_GenericProposalV2.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_GenericProposalV2 *GenericProposalV2Transactor) SetFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV2.contract.Transact(opts, "setFee")
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_GenericProposalV2 *GenericProposalV2Session) SetFee() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.SetFee(&_GenericProposalV2.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_GenericProposalV2 *GenericProposalV2TransactorSession) SetFee() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.SetFee(&_GenericProposalV2.TransactOpts)
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_GenericProposalV2 *GenericProposalV2Transactor) VoteAccept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV2.contract.Transact(opts, "voteAccept")
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_GenericProposalV2 *GenericProposalV2Session) VoteAccept() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.VoteAccept(&_GenericProposalV2.TransactOpts)
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_GenericProposalV2 *GenericProposalV2TransactorSession) VoteAccept() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.VoteAccept(&_GenericProposalV2.TransactOpts)
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_GenericProposalV2 *GenericProposalV2Transactor) VoteReject(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV2.contract.Transact(opts, "voteReject")
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_GenericProposalV2 *GenericProposalV2Session) VoteReject() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.VoteReject(&_GenericProposalV2.TransactOpts)
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_GenericProposalV2 *GenericProposalV2TransactorSession) VoteReject() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.VoteReject(&_GenericProposalV2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GenericProposalV2 *GenericProposalV2Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericProposalV2.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GenericProposalV2 *GenericProposalV2Session) Withdraw() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.Withdraw(&_GenericProposalV2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GenericProposalV2 *GenericProposalV2TransactorSession) Withdraw() (*types.Transaction, error) {
	return _GenericProposalV2.Contract.Withdraw(&_GenericProposalV2.TransactOpts)
}
