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

// BudgetProposalV2ABI is the input ABI used to generate the binding from.
const BudgetProposalV2ABI = "[{\"inputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"_mnregistry_proxy\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_payout_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ref_uuid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proposed_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"accepted_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"budgetStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"is_accepted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"is_finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"unpaid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"canVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created_block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"distributePayout\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee_payer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finish_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAccepted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mnregistry_proxy\",\"outputs\":[{\"internalType\":\"contractIGovernedProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paid_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"payout_address\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposed_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quorum_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ref_uuid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rejected_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setFee\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supermajority\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteAccept\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voteReject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BudgetProposalV2Bin is the compiled bytecode used for deploying new contracts.
const BudgetProposalV2Bin = `60806040523480156200001157600080fd5b50604051620015ff380380620015ff833981810160405260a08110156200003757600080fd5b5080516020808301516040808501516060860151608090960151600180546001600160a01b0319908116331790915543600255600080546001600160a01b03808a16918416821783554285016003556004805491891691909416178355600b805460ff1916604290811790915586517f8abf60770000000000000000000000000000000000000000000000000000000081529651999a9799959894968b96600a96899693958d95909485949193638abf60779383830193909290829003018186803b1580156200010657600080fd5b505afa1580156200011b573d6000803e3d6000fd5b505050506040513d60208110156200013257600080fd5b5051604080517f06661abd00000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216916306661abd9160048082019260a092909190829003018186803b1580156200019057600080fd5b505afa158015620001a5573d6000803e3d6000fd5b505050506040513d60a0811015620001bc57600080fd5b5060408101516080909101519092509050806200023a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4e6f7420726561647920666f722070726f706f73616c73000000000000000000604482015290519081900360640190fd5b60028104821015620002ad57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f41637469766520776569676874203c20312f3220657665722077656967687400604482015290519081900360640190fd5b600160ff871610156200032157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f51756f72756d206d696e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b606460ff871611156200039557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f51756f72756d206d617800000000000000000000000000000000000000000000604482015290519081900360640190fd5b6008829055606460ff8716830204600955603360ff871610620003be57600954600a55620003c9565b60646033830204600a555b6000600954116200043b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f51756f72756d2077656967687400000000000000000000000000000000000000604482015290519081900360640190fd5b6000600a5411620004ad57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f46696e6973682077656967687400000000000000000000000000000000000000604482015290519081900360640190fd5b505050600f96909655505050600e919091555050505061112c80620004d36000396000f3fe6080604052600436106101b75760003560e01c80637639b1eb116100ec578063aec2ccae1161008a578063c40a70f811610064578063c40a70f8146104df578063c86e6c15146104f4578063e522538114610509578063fe7334e81461051e576101b7565b8063aec2ccae14610475578063bd4c1f39146104b5578063c2472ef8146104ca576101b7565b806391840a6b116100c657806391840a6b146103f6578063990a663b1461040b5780639d5e6c9d14610420578063adfaa72e14610435576101b7565b80637639b1eb146103b75780637b352962146103cc57806383197ef0146103e1576101b7565b80634cafdfb21161015957806356c2a0a11161013357806356c2a0a1146103245780635c31f2201461033957806360f96a8f1461036457806375df0f99146103a2576101b7565b80634cafdfb2146102d1578063504881df146102e65780635051a5ec146102fb576101b7565b80632ded3227116101955780632ded3227146102645780633b2a1b141461026c5780633ccfd60b146102a75780633d1db3e9146102bc576101b7565b80630b62be451461021e57806310cac8a51461024557806329dcb0cf1461024f575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f776564000000000000000000000000000000000000000000604482015290519081900360640190fd5b34801561022a57600080fd5b50610233610533565b60408051918252519081900360200190f35b61024d610539565b005b34801561025b57600080fd5b50610233610550565b61024d610556565b34801561027857600080fd5b506102816105e6565b604080519485529215156020850152901515838301526060830152519081900360800190f35b3480156102b357600080fd5b5061024d61061e565b3480156102c857600080fd5b506102336106db565b3480156102dd57600080fd5b506102336106e1565b3480156102f257600080fd5b506102336106e7565b34801561030757600080fd5b506103106106ed565b604080519115158252519081900360200190f35b34801561033057600080fd5b5061024d610746565b34801561034557600080fd5b5061034e610759565b6040805160ff9092168252519081900360200190f35b34801561037057600080fd5b50610379610762565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156103ae57600080fd5b5061023361077e565b3480156103c357600080fd5b50610233610784565b3480156103d857600080fd5b5061031061078a565b3480156103ed57600080fd5b5061024d6107b4565b34801561040257600080fd5b50610233610855565b34801561041757600080fd5b5061023361085b565b34801561042c57600080fd5b50610379610861565b34801561044157600080fd5b506103106004803603602081101561045857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661087d565b34801561048157600080fd5b506103106004803603602081101561049857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610a17565b3480156104c157600080fd5b50610233610a2c565b3480156104d657600080fd5b5061024d610a32565b3480156104eb57600080fd5b50610379610a45565b34801561050057600080fd5b50610233610a61565b34801561051557600080fd5b5061024d610a67565b34801561052a57600080fd5b50610379610dbc565b60025481565b600d8054340190819055600e54101561054e57fe5b565b60035481565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105dc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6005805434019055565b600f54600080806105f56106ed565b92506105ff61078a565b9150600e54600d54111561060f57fe5b600d54600e5403905090919293565b6106266106ed565b61069157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742061636365707465640000000000000000000000000000000000000000604482015290519081900360640190fd5b60045460405173ffffffffffffffffffffffffffffffffffffffff90911690303180156108fc02916000818181858888f193505050501580156106d8573d6000803e3d6000fd5b50565b600a5481565b600e5481565b600d5481565b6000600a546006541061070257506001610743565b61070a61078a565b61071657506000610743565b60095460075460065401101561072e57506000610743565b600b5460095460649160ff1602046006541190505b90565b61074e610dd8565b600780549091019055565b600b5460ff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60065481565b6000426003541115806107a15750600a5460065410155b806107af5750600a54600754115b905090565b60015473ffffffffffffffffffffffffffffffffffffffff16331461083a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60045473ffffffffffffffffffffffffffffffffffffffff16ff5b60085481565b60055481565b60045473ffffffffffffffffffffffffffffffffffffffff1690565b60008054604080517f8abf60770000000000000000000000000000000000000000000000000000000081529051839273ffffffffffffffffffffffffffffffffffffffff1691638abf6077916004808301926020929190829003018186803b1580156108e857600080fd5b505afa1580156108fc573d6000803e3d6000fd5b505050506040513d602081101561091257600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015291519293506000929184169163b83e16059160248082019260e092909190829003018186803b15801561098a57600080fd5b505afa15801561099e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e08110156109c357600080fd5b5060a00151600354909150421080156109dd575060025481105b8015610a0f575073ffffffffffffffffffffffffffffffffffffffff84166000908152600c602052604090205460ff16155b949350505050565b600c6020526000908152604090205460ff1681565b600f5481565b610a3a610dd8565b600680549091019055565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b610a6f61078a565b8015610a805750610a7e6106ed565b155b610aeb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f7420636f6c6c65637461626c650000000000000000000000000000000000604482015290519081900360640190fd5b60015473ffffffffffffffffffffffffffffffffffffffff163314610b7157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610bda57600080fd5b505afa158015610bee573d6000803e3d6000fd5b505050506040513d6020811015610c0457600080fd5b5051604080517fa2731784000000000000000000000000000000000000000000000000000000008152905191925060009173ffffffffffffffffffffffffffffffffffffffff84169163a2731784916004808301926020929190829003018186803b158015610c7257600080fd5b505afa158015610c86573d6000803e3d6000fd5b505050506040513d6020811015610c9c57600080fd5b5051604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921691638abf607791600480820192602092909190829003018186803b158015610d0657600080fd5b505afa158015610d1a573d6000803e3d6000fd5b505050506040513d6020811015610d3057600080fd5b5051604080517fd7bb99ba000000000000000000000000000000000000000000000000000000008152905191925073ffffffffffffffffffffffffffffffffffffffff83169163d7bb99ba91303191600480830192600092919082900301818588803b158015610d9f57600080fd5b505af1158015610db3573d6000803e3d6000fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60004260035411610e4a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f46696e6973686564000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610eb357600080fd5b505afa158015610ec7573d6000803e3d6000fd5b505050506040513d6020811015610edd57600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815233600482018190529151929350909160009173ffffffffffffffffffffffffffffffffffffffff85169163b83e16059160248082019260e092909190829003018186803b158015610f5657600080fd5b505afa158015610f6a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e0811015610f8f57600080fd5b50608081015160a0909101516002549195509150811061101057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f7420656c696769626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600c602052604090205460ff16156110a557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b5073ffffffffffffffffffffffffffffffffffffffff166000908152600c6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055509056fea265627a7a72315820e5e97d2670fa57167c7905e4e899751e87aa3624c34289ac3f5005da752f470164736f6c63430005100032`

// DeployBudgetProposalV2 deploys a new Ethereum contract, binding an instance of BudgetProposalV2 to it.
func DeployBudgetProposalV2(auth *bind.TransactOpts, backend bind.ContractBackend, _mnregistry_proxy common.Address, _payout_address common.Address, _ref_uuid *big.Int, _proposed_amount *big.Int, _period *big.Int) (common.Address, *types.Transaction, *BudgetProposalV2, error) {
	parsed, err := abi.JSON(strings.NewReader(BudgetProposalV2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BudgetProposalV2Bin), backend, _mnregistry_proxy, _payout_address, _ref_uuid, _proposed_amount, _period)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BudgetProposalV2{BudgetProposalV2Caller: BudgetProposalV2Caller{contract: contract}, BudgetProposalV2Transactor: BudgetProposalV2Transactor{contract: contract}, BudgetProposalV2Filterer: BudgetProposalV2Filterer{contract: contract}}, nil
}

// BudgetProposalV2Bin is the compiled bytecode of contract after deployment.
const BudgetProposalV2RuntimeBin = `6080604052600436106101b75760003560e01c80637639b1eb116100ec578063aec2ccae1161008a578063c40a70f811610064578063c40a70f8146104df578063c86e6c15146104f4578063e522538114610509578063fe7334e81461051e576101b7565b8063aec2ccae14610475578063bd4c1f39146104b5578063c2472ef8146104ca576101b7565b806391840a6b116100c657806391840a6b146103f6578063990a663b1461040b5780639d5e6c9d14610420578063adfaa72e14610435576101b7565b80637639b1eb146103b75780637b352962146103cc57806383197ef0146103e1576101b7565b80634cafdfb21161015957806356c2a0a11161013357806356c2a0a1146103245780635c31f2201461033957806360f96a8f1461036457806375df0f99146103a2576101b7565b80634cafdfb2146102d1578063504881df146102e65780635051a5ec146102fb576101b7565b80632ded3227116101955780632ded3227146102645780633b2a1b141461026c5780633ccfd60b146102a75780633d1db3e9146102bc576101b7565b80630b62be451461021e57806310cac8a51461024557806329dcb0cf1461024f575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420616c6c6f776564000000000000000000000000000000000000000000604482015290519081900360640190fd5b34801561022a57600080fd5b50610233610533565b60408051918252519081900360200190f35b61024d610539565b005b34801561025b57600080fd5b50610233610550565b61024d610556565b34801561027857600080fd5b506102816105e6565b604080519485529215156020850152901515838301526060830152519081900360800190f35b3480156102b357600080fd5b5061024d61061e565b3480156102c857600080fd5b506102336106db565b3480156102dd57600080fd5b506102336106e1565b3480156102f257600080fd5b506102336106e7565b34801561030757600080fd5b506103106106ed565b604080519115158252519081900360200190f35b34801561033057600080fd5b5061024d610746565b34801561034557600080fd5b5061034e610759565b6040805160ff9092168252519081900360200190f35b34801561037057600080fd5b50610379610762565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156103ae57600080fd5b5061023361077e565b3480156103c357600080fd5b50610233610784565b3480156103d857600080fd5b5061031061078a565b3480156103ed57600080fd5b5061024d6107b4565b34801561040257600080fd5b50610233610855565b34801561041757600080fd5b5061023361085b565b34801561042c57600080fd5b50610379610861565b34801561044157600080fd5b506103106004803603602081101561045857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661087d565b34801561048157600080fd5b506103106004803603602081101561049857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610a17565b3480156104c157600080fd5b50610233610a2c565b3480156104d657600080fd5b5061024d610a32565b3480156104eb57600080fd5b50610379610a45565b34801561050057600080fd5b50610233610a61565b34801561051557600080fd5b5061024d610a67565b34801561052a57600080fd5b50610379610dbc565b60025481565b600d8054340190819055600e54101561054e57fe5b565b60035481565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105dc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6005805434019055565b600f54600080806105f56106ed565b92506105ff61078a565b9150600e54600d54111561060f57fe5b600d54600e5403905090919293565b6106266106ed565b61069157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742061636365707465640000000000000000000000000000000000000000604482015290519081900360640190fd5b60045460405173ffffffffffffffffffffffffffffffffffffffff90911690303180156108fc02916000818181858888f193505050501580156106d8573d6000803e3d6000fd5b50565b600a5481565b600e5481565b600d5481565b6000600a546006541061070257506001610743565b61070a61078a565b61071657506000610743565b60095460075460065401101561072e57506000610743565b600b5460095460649160ff1602046006541190505b90565b61074e610dd8565b600780549091019055565b600b5460ff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60065481565b6000426003541115806107a15750600a5460065410155b806107af5750600a54600754115b905090565b60015473ffffffffffffffffffffffffffffffffffffffff16331461083a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60045473ffffffffffffffffffffffffffffffffffffffff16ff5b60085481565b60055481565b60045473ffffffffffffffffffffffffffffffffffffffff1690565b60008054604080517f8abf60770000000000000000000000000000000000000000000000000000000081529051839273ffffffffffffffffffffffffffffffffffffffff1691638abf6077916004808301926020929190829003018186803b1580156108e857600080fd5b505afa1580156108fc573d6000803e3d6000fd5b505050506040513d602081101561091257600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015291519293506000929184169163b83e16059160248082019260e092909190829003018186803b15801561098a57600080fd5b505afa15801561099e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e08110156109c357600080fd5b5060a00151600354909150421080156109dd575060025481105b8015610a0f575073ffffffffffffffffffffffffffffffffffffffff84166000908152600c602052604090205460ff16155b949350505050565b600c6020526000908152604090205460ff1681565b600f5481565b610a3a610dd8565b600680549091019055565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b610a6f61078a565b8015610a805750610a7e6106ed565b155b610aeb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f7420636f6c6c65637461626c650000000000000000000000000000000000604482015290519081900360640190fd5b60015473ffffffffffffffffffffffffffffffffffffffff163314610b7157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4f6e6c7920706172656e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610bda57600080fd5b505afa158015610bee573d6000803e3d6000fd5b505050506040513d6020811015610c0457600080fd5b5051604080517fa2731784000000000000000000000000000000000000000000000000000000008152905191925060009173ffffffffffffffffffffffffffffffffffffffff84169163a2731784916004808301926020929190829003018186803b158015610c7257600080fd5b505afa158015610c86573d6000803e3d6000fd5b505050506040513d6020811015610c9c57600080fd5b5051604080517f8abf6077000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff90921691638abf607791600480820192602092909190829003018186803b158015610d0657600080fd5b505afa158015610d1a573d6000803e3d6000fd5b505050506040513d6020811015610d3057600080fd5b5051604080517fd7bb99ba000000000000000000000000000000000000000000000000000000008152905191925073ffffffffffffffffffffffffffffffffffffffff83169163d7bb99ba91303191600480830192600092919082900301818588803b158015610d9f57600080fd5b505af1158015610db3573d6000803e3d6000fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60004260035411610e4a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f46696e6973686564000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b158015610eb357600080fd5b505afa158015610ec7573d6000803e3d6000fd5b505050506040513d6020811015610edd57600080fd5b5051604080517fb83e160500000000000000000000000000000000000000000000000000000000815233600482018190529151929350909160009173ffffffffffffffffffffffffffffffffffffffff85169163b83e16059160248082019260e092909190829003018186803b158015610f5657600080fd5b505afa158015610f6a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e0811015610f8f57600080fd5b50608081015160a0909101516002549195509150811061101057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f7420656c696769626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600c602052604090205460ff16156110a557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b5073ffffffffffffffffffffffffffffffffffffffff166000908152600c6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055509056fea265627a7a72315820e5e97d2670fa57167c7905e4e899751e87aa3624c34289ac3f5005da752f470164736f6c63430005100032`

// BudgetProposalV2 is an auto generated Go binding around an Ethereum contract.
type BudgetProposalV2 struct {
	BudgetProposalV2Caller     // Read-only binding to the contract
	BudgetProposalV2Transactor // Write-only binding to the contract
	BudgetProposalV2Filterer   // Log filterer for contract events
}

// BudgetProposalV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type BudgetProposalV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BudgetProposalV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BudgetProposalV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BudgetProposalV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BudgetProposalV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BudgetProposalV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BudgetProposalV2Session struct {
	Contract     *BudgetProposalV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BudgetProposalV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BudgetProposalV2CallerSession struct {
	Contract *BudgetProposalV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BudgetProposalV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BudgetProposalV2TransactorSession struct {
	Contract     *BudgetProposalV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BudgetProposalV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type BudgetProposalV2Raw struct {
	Contract *BudgetProposalV2 // Generic contract binding to access the raw methods on
}

// BudgetProposalV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BudgetProposalV2CallerRaw struct {
	Contract *BudgetProposalV2Caller // Generic read-only contract binding to access the raw methods on
}

// BudgetProposalV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BudgetProposalV2TransactorRaw struct {
	Contract *BudgetProposalV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBudgetProposalV2 creates a new instance of BudgetProposalV2, bound to a specific deployed contract.
func NewBudgetProposalV2(address common.Address, backend bind.ContractBackend) (*BudgetProposalV2, error) {
	contract, err := bindBudgetProposalV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BudgetProposalV2{BudgetProposalV2Caller: BudgetProposalV2Caller{contract: contract}, BudgetProposalV2Transactor: BudgetProposalV2Transactor{contract: contract}, BudgetProposalV2Filterer: BudgetProposalV2Filterer{contract: contract}}, nil
}

// NewBudgetProposalV2Caller creates a new read-only instance of BudgetProposalV2, bound to a specific deployed contract.
func NewBudgetProposalV2Caller(address common.Address, caller bind.ContractCaller) (*BudgetProposalV2Caller, error) {
	contract, err := bindBudgetProposalV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BudgetProposalV2Caller{contract: contract}, nil
}

// NewBudgetProposalV2Transactor creates a new write-only instance of BudgetProposalV2, bound to a specific deployed contract.
func NewBudgetProposalV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BudgetProposalV2Transactor, error) {
	contract, err := bindBudgetProposalV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BudgetProposalV2Transactor{contract: contract}, nil
}

// NewBudgetProposalV2Filterer creates a new log filterer instance of BudgetProposalV2, bound to a specific deployed contract.
func NewBudgetProposalV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BudgetProposalV2Filterer, error) {
	contract, err := bindBudgetProposalV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BudgetProposalV2Filterer{contract: contract}, nil
}

// bindBudgetProposalV2 binds a generic wrapper to an already deployed contract.
func bindBudgetProposalV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BudgetProposalV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BudgetProposalV2 *BudgetProposalV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BudgetProposalV2.Contract.BudgetProposalV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BudgetProposalV2 *BudgetProposalV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.BudgetProposalV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BudgetProposalV2 *BudgetProposalV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.BudgetProposalV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BudgetProposalV2 *BudgetProposalV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BudgetProposalV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BudgetProposalV2 *BudgetProposalV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BudgetProposalV2 *BudgetProposalV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.contract.Transact(opts, method, params...)
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) AcceptedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "accepted_weight")
	return *ret0, err
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) AcceptedWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.AcceptedWeight(&_BudgetProposalV2.CallOpts)
}

// AcceptedWeight is a free data retrieval call binding the contract method 0x7639b1eb.
//
// Solidity: function accepted_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) AcceptedWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.AcceptedWeight(&_BudgetProposalV2.CallOpts)
}

// BudgetStatus is a free data retrieval call binding the contract method 0x3b2a1b14.
//
// Solidity: function budgetStatus() constant returns(uint256 uuid, bool is_accepted, bool is_finished, uint256 unpaid)
func (_BudgetProposalV2 *BudgetProposalV2Caller) BudgetStatus(opts *bind.CallOpts) (struct {
	Uuid       *big.Int
	IsAccepted bool
	IsFinished bool
	Unpaid     *big.Int
}, error) {
	ret := new(struct {
		Uuid       *big.Int
		IsAccepted bool
		IsFinished bool
		Unpaid     *big.Int
	})
	out := ret
	err := _BudgetProposalV2.contract.Call(opts, out, "budgetStatus")
	return *ret, err
}

// BudgetStatus is a free data retrieval call binding the contract method 0x3b2a1b14.
//
// Solidity: function budgetStatus() constant returns(uint256 uuid, bool is_accepted, bool is_finished, uint256 unpaid)
func (_BudgetProposalV2 *BudgetProposalV2Session) BudgetStatus() (struct {
	Uuid       *big.Int
	IsAccepted bool
	IsFinished bool
	Unpaid     *big.Int
}, error) {
	return _BudgetProposalV2.Contract.BudgetStatus(&_BudgetProposalV2.CallOpts)
}

// BudgetStatus is a free data retrieval call binding the contract method 0x3b2a1b14.
//
// Solidity: function budgetStatus() constant returns(uint256 uuid, bool is_accepted, bool is_finished, uint256 unpaid)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) BudgetStatus() (struct {
	Uuid       *big.Int
	IsAccepted bool
	IsFinished bool
	Unpaid     *big.Int
}, error) {
	return _BudgetProposalV2.Contract.BudgetStatus(&_BudgetProposalV2.CallOpts)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2Caller) CanVote(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "canVote", owner)
	return *ret0, err
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2Session) CanVote(owner common.Address) (bool, error) {
	return _BudgetProposalV2.Contract.CanVote(&_BudgetProposalV2.CallOpts, owner)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address owner) constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) CanVote(owner common.Address) (bool, error) {
	return _BudgetProposalV2.Contract.CanVote(&_BudgetProposalV2.CallOpts, owner)
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) CreatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "created_block")
	return *ret0, err
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) CreatedBlock() (*big.Int, error) {
	return _BudgetProposalV2.Contract.CreatedBlock(&_BudgetProposalV2.CallOpts)
}

// CreatedBlock is a free data retrieval call binding the contract method 0x0b62be45.
//
// Solidity: function created_block() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) CreatedBlock() (*big.Int, error) {
	return _BudgetProposalV2.Contract.CreatedBlock(&_BudgetProposalV2.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "deadline")
	return *ret0, err
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) Deadline() (*big.Int, error) {
	return _BudgetProposalV2.Contract.Deadline(&_BudgetProposalV2.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) Deadline() (*big.Int, error) {
	return _BudgetProposalV2.Contract.Deadline(&_BudgetProposalV2.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) FeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "fee_amount")
	return *ret0, err
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) FeeAmount() (*big.Int, error) {
	return _BudgetProposalV2.Contract.FeeAmount(&_BudgetProposalV2.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x990a663b.
//
// Solidity: function fee_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) FeeAmount() (*big.Int, error) {
	return _BudgetProposalV2.Contract.FeeAmount(&_BudgetProposalV2.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2Caller) FeePayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "fee_payer")
	return *ret0, err
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2Session) FeePayer() (common.Address, error) {
	return _BudgetProposalV2.Contract.FeePayer(&_BudgetProposalV2.CallOpts)
}

// FeePayer is a free data retrieval call binding the contract method 0xc40a70f8.
//
// Solidity: function fee_payer() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) FeePayer() (common.Address, error) {
	return _BudgetProposalV2.Contract.FeePayer(&_BudgetProposalV2.CallOpts)
}

// FinishWeight is a free data retrieval call binding the contract method 0x3d1db3e9.
//
// Solidity: function finish_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) FinishWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "finish_weight")
	return *ret0, err
}

// FinishWeight is a free data retrieval call binding the contract method 0x3d1db3e9.
//
// Solidity: function finish_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) FinishWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.FinishWeight(&_BudgetProposalV2.CallOpts)
}

// FinishWeight is a free data retrieval call binding the contract method 0x3d1db3e9.
//
// Solidity: function finish_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) FinishWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.FinishWeight(&_BudgetProposalV2.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2Caller) IsAccepted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "isAccepted")
	return *ret0, err
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2Session) IsAccepted() (bool, error) {
	return _BudgetProposalV2.Contract.IsAccepted(&_BudgetProposalV2.CallOpts)
}

// IsAccepted is a free data retrieval call binding the contract method 0x5051a5ec.
//
// Solidity: function isAccepted() constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) IsAccepted() (bool, error) {
	return _BudgetProposalV2.Contract.IsAccepted(&_BudgetProposalV2.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2Caller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2Session) IsFinished() (bool, error) {
	return _BudgetProposalV2.Contract.IsFinished(&_BudgetProposalV2.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) IsFinished() (bool, error) {
	return _BudgetProposalV2.Contract.IsFinished(&_BudgetProposalV2.CallOpts)
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2Caller) MnregistryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "mnregistry_proxy")
	return *ret0, err
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2Session) MnregistryProxy() (common.Address, error) {
	return _BudgetProposalV2.Contract.MnregistryProxy(&_BudgetProposalV2.CallOpts)
}

// MnregistryProxy is a free data retrieval call binding the contract method 0xfe7334e8.
//
// Solidity: function mnregistry_proxy() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) MnregistryProxy() (common.Address, error) {
	return _BudgetProposalV2.Contract.MnregistryProxy(&_BudgetProposalV2.CallOpts)
}

// PaidAmount is a free data retrieval call binding the contract method 0x504881df.
//
// Solidity: function paid_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) PaidAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "paid_amount")
	return *ret0, err
}

// PaidAmount is a free data retrieval call binding the contract method 0x504881df.
//
// Solidity: function paid_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) PaidAmount() (*big.Int, error) {
	return _BudgetProposalV2.Contract.PaidAmount(&_BudgetProposalV2.CallOpts)
}

// PaidAmount is a free data retrieval call binding the contract method 0x504881df.
//
// Solidity: function paid_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) PaidAmount() (*big.Int, error) {
	return _BudgetProposalV2.Contract.PaidAmount(&_BudgetProposalV2.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2Caller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "parent")
	return *ret0, err
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2Session) Parent() (common.Address, error) {
	return _BudgetProposalV2.Contract.Parent(&_BudgetProposalV2.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) Parent() (common.Address, error) {
	return _BudgetProposalV2.Contract.Parent(&_BudgetProposalV2.CallOpts)
}

// PayoutAddress is a free data retrieval call binding the contract method 0x9d5e6c9d.
//
// Solidity: function payout_address() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2Caller) PayoutAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "payout_address")
	return *ret0, err
}

// PayoutAddress is a free data retrieval call binding the contract method 0x9d5e6c9d.
//
// Solidity: function payout_address() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2Session) PayoutAddress() (common.Address, error) {
	return _BudgetProposalV2.Contract.PayoutAddress(&_BudgetProposalV2.CallOpts)
}

// PayoutAddress is a free data retrieval call binding the contract method 0x9d5e6c9d.
//
// Solidity: function payout_address() constant returns(address)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) PayoutAddress() (common.Address, error) {
	return _BudgetProposalV2.Contract.PayoutAddress(&_BudgetProposalV2.CallOpts)
}

// ProposedAmount is a free data retrieval call binding the contract method 0x4cafdfb2.
//
// Solidity: function proposed_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) ProposedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "proposed_amount")
	return *ret0, err
}

// ProposedAmount is a free data retrieval call binding the contract method 0x4cafdfb2.
//
// Solidity: function proposed_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) ProposedAmount() (*big.Int, error) {
	return _BudgetProposalV2.Contract.ProposedAmount(&_BudgetProposalV2.CallOpts)
}

// ProposedAmount is a free data retrieval call binding the contract method 0x4cafdfb2.
//
// Solidity: function proposed_amount() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) ProposedAmount() (*big.Int, error) {
	return _BudgetProposalV2.Contract.ProposedAmount(&_BudgetProposalV2.CallOpts)
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) QuorumWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "quorum_weight")
	return *ret0, err
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) QuorumWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.QuorumWeight(&_BudgetProposalV2.CallOpts)
}

// QuorumWeight is a free data retrieval call binding the contract method 0x75df0f99.
//
// Solidity: function quorum_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) QuorumWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.QuorumWeight(&_BudgetProposalV2.CallOpts)
}

// RefUuid is a free data retrieval call binding the contract method 0xbd4c1f39.
//
// Solidity: function ref_uuid() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) RefUuid(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "ref_uuid")
	return *ret0, err
}

// RefUuid is a free data retrieval call binding the contract method 0xbd4c1f39.
//
// Solidity: function ref_uuid() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) RefUuid() (*big.Int, error) {
	return _BudgetProposalV2.Contract.RefUuid(&_BudgetProposalV2.CallOpts)
}

// RefUuid is a free data retrieval call binding the contract method 0xbd4c1f39.
//
// Solidity: function ref_uuid() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) RefUuid() (*big.Int, error) {
	return _BudgetProposalV2.Contract.RefUuid(&_BudgetProposalV2.CallOpts)
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) RejectedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "rejected_weight")
	return *ret0, err
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) RejectedWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.RejectedWeight(&_BudgetProposalV2.CallOpts)
}

// RejectedWeight is a free data retrieval call binding the contract method 0xc86e6c15.
//
// Solidity: function rejected_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) RejectedWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.RejectedWeight(&_BudgetProposalV2.CallOpts)
}

// Supermajority is a free data retrieval call binding the contract method 0x5c31f220.
//
// Solidity: function supermajority() constant returns(uint8)
func (_BudgetProposalV2 *BudgetProposalV2Caller) Supermajority(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "supermajority")
	return *ret0, err
}

// Supermajority is a free data retrieval call binding the contract method 0x5c31f220.
//
// Solidity: function supermajority() constant returns(uint8)
func (_BudgetProposalV2 *BudgetProposalV2Session) Supermajority() (uint8, error) {
	return _BudgetProposalV2.Contract.Supermajority(&_BudgetProposalV2.CallOpts)
}

// Supermajority is a free data retrieval call binding the contract method 0x5c31f220.
//
// Solidity: function supermajority() constant returns(uint8)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) Supermajority() (uint8, error) {
	return _BudgetProposalV2.Contract.Supermajority(&_BudgetProposalV2.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Caller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "total_weight")
	return *ret0, err
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2Session) TotalWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.TotalWeight(&_BudgetProposalV2.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x91840a6b.
//
// Solidity: function total_weight() constant returns(uint256)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) TotalWeight() (*big.Int, error) {
	return _BudgetProposalV2.Contract.TotalWeight(&_BudgetProposalV2.CallOpts)
}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2Caller) Voted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BudgetProposalV2.contract.Call(opts, out, "voted", arg0)
	return *ret0, err
}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2Session) Voted(arg0 common.Address) (bool, error) {
	return _BudgetProposalV2.Contract.Voted(&_BudgetProposalV2.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) constant returns(bool)
func (_BudgetProposalV2 *BudgetProposalV2CallerSession) Voted(arg0 common.Address) (bool, error) {
	return _BudgetProposalV2.Contract.Voted(&_BudgetProposalV2.CallOpts, arg0)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_BudgetProposalV2 *BudgetProposalV2Transactor) Collect(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.contract.Transact(opts, "collect")
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_BudgetProposalV2 *BudgetProposalV2Session) Collect() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.Collect(&_BudgetProposalV2.TransactOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_BudgetProposalV2 *BudgetProposalV2TransactorSession) Collect() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.Collect(&_BudgetProposalV2.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BudgetProposalV2 *BudgetProposalV2Transactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BudgetProposalV2 *BudgetProposalV2Session) Destroy() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.Destroy(&_BudgetProposalV2.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BudgetProposalV2 *BudgetProposalV2TransactorSession) Destroy() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.Destroy(&_BudgetProposalV2.TransactOpts)
}

// DistributePayout is a paid mutator transaction binding the contract method 0x10cac8a5.
//
// Solidity: function distributePayout() returns()
func (_BudgetProposalV2 *BudgetProposalV2Transactor) DistributePayout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.contract.Transact(opts, "distributePayout")
}

// DistributePayout is a paid mutator transaction binding the contract method 0x10cac8a5.
//
// Solidity: function distributePayout() returns()
func (_BudgetProposalV2 *BudgetProposalV2Session) DistributePayout() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.DistributePayout(&_BudgetProposalV2.TransactOpts)
}

// DistributePayout is a paid mutator transaction binding the contract method 0x10cac8a5.
//
// Solidity: function distributePayout() returns()
func (_BudgetProposalV2 *BudgetProposalV2TransactorSession) DistributePayout() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.DistributePayout(&_BudgetProposalV2.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_BudgetProposalV2 *BudgetProposalV2Transactor) SetFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.contract.Transact(opts, "setFee")
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_BudgetProposalV2 *BudgetProposalV2Session) SetFee() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.SetFee(&_BudgetProposalV2.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x2ded3227.
//
// Solidity: function setFee() returns()
func (_BudgetProposalV2 *BudgetProposalV2TransactorSession) SetFee() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.SetFee(&_BudgetProposalV2.TransactOpts)
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_BudgetProposalV2 *BudgetProposalV2Transactor) VoteAccept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.contract.Transact(opts, "voteAccept")
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_BudgetProposalV2 *BudgetProposalV2Session) VoteAccept() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.VoteAccept(&_BudgetProposalV2.TransactOpts)
}

// VoteAccept is a paid mutator transaction binding the contract method 0xc2472ef8.
//
// Solidity: function voteAccept() returns()
func (_BudgetProposalV2 *BudgetProposalV2TransactorSession) VoteAccept() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.VoteAccept(&_BudgetProposalV2.TransactOpts)
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_BudgetProposalV2 *BudgetProposalV2Transactor) VoteReject(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.contract.Transact(opts, "voteReject")
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_BudgetProposalV2 *BudgetProposalV2Session) VoteReject() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.VoteReject(&_BudgetProposalV2.TransactOpts)
}

// VoteReject is a paid mutator transaction binding the contract method 0x56c2a0a1.
//
// Solidity: function voteReject() returns()
func (_BudgetProposalV2 *BudgetProposalV2TransactorSession) VoteReject() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.VoteReject(&_BudgetProposalV2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BudgetProposalV2 *BudgetProposalV2Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BudgetProposalV2.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BudgetProposalV2 *BudgetProposalV2Session) Withdraw() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.Withdraw(&_BudgetProposalV2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BudgetProposalV2 *BudgetProposalV2TransactorSession) Withdraw() (*types.Transaction, error) {
	return _BudgetProposalV2.Contract.Withdraw(&_BudgetProposalV2.TransactOpts)
}
