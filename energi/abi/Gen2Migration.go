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

// Gen2MigrationABI is the input ABI used to generate the binding from.
const Gen2MigrationABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"hashToSign\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chain_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_item_id\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"sig_v\",\"type\":\"uint8\"},{\"name\":\"sig_r\",\"type\":\"bytes32\"},{\"name\":\"sig_s\",\"type\":\"bytes32\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"itemCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury_proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"name\":\"owner\",\"type\":\"bytes20\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"bytes20[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setSnapshot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_item_id\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"sig_v\",\"type\":\"uint8\"},{\"name\":\"sig_r\",\"type\":\"bytes32\"},{\"name\":\"sig_s\",\"type\":\"bytes32\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_treasury_proxy\",\"type\":\"address\"},{\"name\":\"_chain_id\",\"type\":\"uint256\"},{\"name\":\"_signer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"item_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Migrated\",\"type\":\"event\"}]"

// Gen2MigrationBin is the compiled bytecode used for deploying new contracts.
const Gen2MigrationBin = `608060405234801561001057600080fd5b50604051610dc0380380610dc08339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b039384166001600160a01b031991821617909155600193909355600280549290911691909216179055610d3c806100846000396000f3fe6080604052600436106100bc5760003560e01c80636bfb0d0111610074578063c66106571161004e578063c661065714610279578063e38e6e5e146102df578063f7121490146103ae576100bc565b80636bfb0d01146102385780639890220b1461024d578063a273178414610264576100bc565b80633af973b1116100a55780633af973b11461018a578063476ce0c31461019f5780635b7633d0146101fa576100bc565b80630a96cb49146101235780631a39d8ef14610175575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561012f57600080fd5b506101636004803603602081101561014657600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610409565b60408051918252519081900360200190f35b34801561018157600080fd5b50610163610489565b34801561019657600080fd5b5061016361048f565b3480156101ab57600080fd5b50610163600480360360a08110156101c257600080fd5b5080359073ffffffffffffffffffffffffffffffffffffffff6020820135169060ff6040820135169060608101359060800135610495565b34801561020657600080fd5b5061020f61066e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561024457600080fd5b5061016361068a565b34801561025957600080fd5b50610262610691565b005b34801561027057600080fd5b5061020f610827565b34801561028557600080fd5b506102a36004803603602081101561029c57600080fd5b5035610843565b604080517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000909316835260208301919091528051918290030190f35b3480156102eb57600080fd5b506102626004803603604081101561030257600080fd5b81019060208101813564010000000081111561031d57600080fd5b82018360208201111561032f57600080fd5b8035906020019184602083028401116401000000008311171561035157600080fd5b91939092909160208101903564010000000081111561036f57600080fd5b82018360208201111561038157600080fd5b803590602001918460208302840111640100000000831117156103a357600080fd5b509092509050610873565b3480156103ba57600080fd5b50610262600480360360a08110156103d157600080fd5b5080359073ffffffffffffffffffffffffffffffffffffffff6020820135169060ff6040820135169060608101359060800135610b46565b6001546040805160609390931b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166020808501919091527f7c7c456e657267692047656e2032206d6967726174696f6e20636c61696d7c7c60348501526054808501939093528151808503909301835260749093019052805191012090565b60035481565b60015481565b600454600090861061050857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f496e76616c696420494400000000000000000000000000000000000000000000604482015290519081900360640190fd5b600061051386610409565b9050600060018287878760405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610574573d6000803e3d6000fd5b5050506020604051035160601b9050806bffffffffffffffffffffffff1916600489815481106105a057fe5b600091825260209091206002909102015460601b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000161461064257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f496e76616c6964207369676e6174757265000000000000000000000000000000604482015290519081900360640190fd5b6004888154811061064f57fe5b9060005260206000209060020201600101549250505095945050505050565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b6004545b90565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b1580156106fa57600080fd5b505afa15801561070e573d6000803e3d6000fd5b505050506040513d602081101561072457600080fd5b505190503373ffffffffffffffffffffffffffffffffffffffff8216146107ac57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742074726561737572790000000000000000000000000000000000000000604482015290519081900360640190fd5b8073ffffffffffffffffffffffffffffffffffffffff1663d7bb99ba3073ffffffffffffffffffffffffffffffffffffffff16316040518263ffffffff1660e01b81526004016000604051808303818588803b15801561080b57600080fd5b505af115801561081f573d6000803e3d6000fd5b505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6004818154811061085057fe5b60009182526020909120600290910201805460019091015460609190911b915082565b600454156108e257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f416c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b60025473ffffffffffffffffffffffffffffffffffffffff16331461096857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c69642073656e646572000000000000000000000000000000000000604482015290519081900360640190fd5b8281146109d657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6d61746368206c656e6774680000000000000000000000000000000000000000604482015290519081900360640190fd5b82610a4257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f6861732064617461000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b82610a4e600482610c8f565b506000835b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019015610b3c57858582818110610a8857fe5b905060200201356bffffffffffffffffffffffff191660048281548110610aab57fe5b6000918252602090912060029091020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001660609290921c919091179055838382818110610af757fe5b9050602002013560048281548110610b0b57fe5b906000526020600020906002020160010181905550838382818110610b2c57fe5b9050602002013582019150610a53565b5060035550505050565b6000610b558686868686610495565b905060008111610bc657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c7265616479207370656e7400000000000000000000000000000000000000604482015290519081900360640190fd5b600060048781548110610bd557fe5b6000918252602091829020600160029092020101919091556040805188815273ffffffffffffffffffffffffffffffffffffffff881692810192909252818101839052517ffd90f074a575cd9336850f79afca1e89f5ca1bf434d82a21ca5d6d4a87a724a49181900360600190a160405173ffffffffffffffffffffffffffffffffffffffff86169082156108fc029083906000818181858888f19350505050158015610c86573d6000803e3d6000fd5b50505050505050565b815481835581811115610cbb57600202816002028360005260206000209182019101610cbb9190610cc0565b505050565b61068e91905b80821115610d035780547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560006001820155600201610cc6565b509056fea265627a7a7230582061059e90c70cedee50de6bf1053d42546a48d9436ecdc292a479e5542391881064736f6c63430005090032`

// DeployGen2Migration deploys a new Ethereum contract, binding an instance of Gen2Migration to it.
func DeployGen2Migration(auth *bind.TransactOpts, backend bind.ContractBackend, _treasury_proxy common.Address, _chain_id *big.Int, _signer common.Address) (common.Address, *types.Transaction, *Gen2Migration, error) {
	parsed, err := abi.JSON(strings.NewReader(Gen2MigrationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Gen2MigrationBin), backend, _treasury_proxy, _chain_id, _signer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gen2Migration{Gen2MigrationCaller: Gen2MigrationCaller{contract: contract}, Gen2MigrationTransactor: Gen2MigrationTransactor{contract: contract}, Gen2MigrationFilterer: Gen2MigrationFilterer{contract: contract}}, nil
}

// Gen2MigrationBin is the compiled bytecode of contract after deployment.
const Gen2MigrationRuntimeBin = `6080604052600436106100bc5760003560e01c80636bfb0d0111610074578063c66106571161004e578063c661065714610279578063e38e6e5e146102df578063f7121490146103ae576100bc565b80636bfb0d01146102385780639890220b1461024d578063a273178414610264576100bc565b80633af973b1116100a55780633af973b11461018a578063476ce0c31461019f5780635b7633d0146101fa576100bc565b80630a96cb49146101235780631a39d8ef14610175575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561012f57600080fd5b506101636004803603602081101561014657600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610409565b60408051918252519081900360200190f35b34801561018157600080fd5b50610163610489565b34801561019657600080fd5b5061016361048f565b3480156101ab57600080fd5b50610163600480360360a08110156101c257600080fd5b5080359073ffffffffffffffffffffffffffffffffffffffff6020820135169060ff6040820135169060608101359060800135610495565b34801561020657600080fd5b5061020f61066e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561024457600080fd5b5061016361068a565b34801561025957600080fd5b50610262610691565b005b34801561027057600080fd5b5061020f610827565b34801561028557600080fd5b506102a36004803603602081101561029c57600080fd5b5035610843565b604080517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000909316835260208301919091528051918290030190f35b3480156102eb57600080fd5b506102626004803603604081101561030257600080fd5b81019060208101813564010000000081111561031d57600080fd5b82018360208201111561032f57600080fd5b8035906020019184602083028401116401000000008311171561035157600080fd5b91939092909160208101903564010000000081111561036f57600080fd5b82018360208201111561038157600080fd5b803590602001918460208302840111640100000000831117156103a357600080fd5b509092509050610873565b3480156103ba57600080fd5b50610262600480360360a08110156103d157600080fd5b5080359073ffffffffffffffffffffffffffffffffffffffff6020820135169060ff6040820135169060608101359060800135610b46565b6001546040805160609390931b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166020808501919091527f7c7c456e657267692047656e2032206d6967726174696f6e20636c61696d7c7c60348501526054808501939093528151808503909301835260749093019052805191012090565b60035481565b60015481565b600454600090861061050857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f496e76616c696420494400000000000000000000000000000000000000000000604482015290519081900360640190fd5b600061051386610409565b9050600060018287878760405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610574573d6000803e3d6000fd5b5050506020604051035160601b9050806bffffffffffffffffffffffff1916600489815481106105a057fe5b600091825260209091206002909102015460601b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000161461064257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f496e76616c6964207369676e6174757265000000000000000000000000000000604482015290519081900360640190fd5b6004888154811061064f57fe5b9060005260206000209060020201600101549250505095945050505050565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b6004545b90565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638abf60776040518163ffffffff1660e01b815260040160206040518083038186803b1580156106fa57600080fd5b505afa15801561070e573d6000803e3d6000fd5b505050506040513d602081101561072457600080fd5b505190503373ffffffffffffffffffffffffffffffffffffffff8216146107ac57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e6f742074726561737572790000000000000000000000000000000000000000604482015290519081900360640190fd5b8073ffffffffffffffffffffffffffffffffffffffff1663d7bb99ba3073ffffffffffffffffffffffffffffffffffffffff16316040518263ffffffff1660e01b81526004016000604051808303818588803b15801561080b57600080fd5b505af115801561081f573d6000803e3d6000fd5b505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6004818154811061085057fe5b60009182526020909120600290910201805460019091015460609190911b915082565b600454156108e257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f416c726561647920736574000000000000000000000000000000000000000000604482015290519081900360640190fd5b60025473ffffffffffffffffffffffffffffffffffffffff16331461096857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c69642073656e646572000000000000000000000000000000000000604482015290519081900360640190fd5b8281146109d657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6d61746368206c656e6774680000000000000000000000000000000000000000604482015290519081900360640190fd5b82610a4257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f6861732064617461000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b82610a4e600482610c8f565b506000835b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019015610b3c57858582818110610a8857fe5b905060200201356bffffffffffffffffffffffff191660048281548110610aab57fe5b6000918252602090912060029091020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001660609290921c919091179055838382818110610af757fe5b9050602002013560048281548110610b0b57fe5b906000526020600020906002020160010181905550838382818110610b2c57fe5b9050602002013582019150610a53565b5060035550505050565b6000610b558686868686610495565b905060008111610bc657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f416c7265616479207370656e7400000000000000000000000000000000000000604482015290519081900360640190fd5b600060048781548110610bd557fe5b6000918252602091829020600160029092020101919091556040805188815273ffffffffffffffffffffffffffffffffffffffff881692810192909252818101839052517ffd90f074a575cd9336850f79afca1e89f5ca1bf434d82a21ca5d6d4a87a724a49181900360600190a160405173ffffffffffffffffffffffffffffffffffffffff86169082156108fc029083906000818181858888f19350505050158015610c86573d6000803e3d6000fd5b50505050505050565b815481835581811115610cbb57600202816002028360005260206000209182019101610cbb9190610cc0565b505050565b61068e91905b80821115610d035780547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560006001820155600201610cc6565b509056fea265627a7a7230582061059e90c70cedee50de6bf1053d42546a48d9436ecdc292a479e5542391881064736f6c63430005090032`

// Gen2Migration is an auto generated Go binding around an Ethereum contract.
type Gen2Migration struct {
	Gen2MigrationCaller     // Read-only binding to the contract
	Gen2MigrationTransactor // Write-only binding to the contract
	Gen2MigrationFilterer   // Log filterer for contract events
}

// Gen2MigrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type Gen2MigrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gen2MigrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Gen2MigrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gen2MigrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Gen2MigrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gen2MigrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Gen2MigrationSession struct {
	Contract     *Gen2Migration    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Gen2MigrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Gen2MigrationCallerSession struct {
	Contract *Gen2MigrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Gen2MigrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Gen2MigrationTransactorSession struct {
	Contract     *Gen2MigrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Gen2MigrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type Gen2MigrationRaw struct {
	Contract *Gen2Migration // Generic contract binding to access the raw methods on
}

// Gen2MigrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Gen2MigrationCallerRaw struct {
	Contract *Gen2MigrationCaller // Generic read-only contract binding to access the raw methods on
}

// Gen2MigrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Gen2MigrationTransactorRaw struct {
	Contract *Gen2MigrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGen2Migration creates a new instance of Gen2Migration, bound to a specific deployed contract.
func NewGen2Migration(address common.Address, backend bind.ContractBackend) (*Gen2Migration, error) {
	contract, err := bindGen2Migration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gen2Migration{Gen2MigrationCaller: Gen2MigrationCaller{contract: contract}, Gen2MigrationTransactor: Gen2MigrationTransactor{contract: contract}, Gen2MigrationFilterer: Gen2MigrationFilterer{contract: contract}}, nil
}

// NewGen2MigrationCaller creates a new read-only instance of Gen2Migration, bound to a specific deployed contract.
func NewGen2MigrationCaller(address common.Address, caller bind.ContractCaller) (*Gen2MigrationCaller, error) {
	contract, err := bindGen2Migration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Gen2MigrationCaller{contract: contract}, nil
}

// NewGen2MigrationTransactor creates a new write-only instance of Gen2Migration, bound to a specific deployed contract.
func NewGen2MigrationTransactor(address common.Address, transactor bind.ContractTransactor) (*Gen2MigrationTransactor, error) {
	contract, err := bindGen2Migration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Gen2MigrationTransactor{contract: contract}, nil
}

// NewGen2MigrationFilterer creates a new log filterer instance of Gen2Migration, bound to a specific deployed contract.
func NewGen2MigrationFilterer(address common.Address, filterer bind.ContractFilterer) (*Gen2MigrationFilterer, error) {
	contract, err := bindGen2Migration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Gen2MigrationFilterer{contract: contract}, nil
}

// bindGen2Migration binds a generic wrapper to an already deployed contract.
func bindGen2Migration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Gen2MigrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gen2Migration *Gen2MigrationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gen2Migration.Contract.Gen2MigrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gen2Migration *Gen2MigrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gen2Migration.Contract.Gen2MigrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gen2Migration *Gen2MigrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gen2Migration.Contract.Gen2MigrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gen2Migration *Gen2MigrationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gen2Migration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gen2Migration *Gen2MigrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gen2Migration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gen2Migration *Gen2MigrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gen2Migration.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gen2Migration.contract.Call(opts, out, "chain_id")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationSession) ChainId() (*big.Int, error) {
	return _Gen2Migration.Contract.ChainId(&_Gen2Migration.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationCallerSession) ChainId() (*big.Int, error) {
	return _Gen2Migration.Contract.ChainId(&_Gen2Migration.CallOpts)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) constant returns(bytes20 owner, uint256 amount)
func (_Gen2Migration *Gen2MigrationCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner  [20]byte
	Amount *big.Int
}, error) {
	ret := new(struct {
		Owner  [20]byte
		Amount *big.Int
	})
	out := ret
	err := _Gen2Migration.contract.Call(opts, out, "coins", arg0)
	return *ret, err
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) constant returns(bytes20 owner, uint256 amount)
func (_Gen2Migration *Gen2MigrationSession) Coins(arg0 *big.Int) (struct {
	Owner  [20]byte
	Amount *big.Int
}, error) {
	return _Gen2Migration.Contract.Coins(&_Gen2Migration.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) constant returns(bytes20 owner, uint256 amount)
func (_Gen2Migration *Gen2MigrationCallerSession) Coins(arg0 *big.Int) (struct {
	Owner  [20]byte
	Amount *big.Int
}, error) {
	return _Gen2Migration.Contract.Coins(&_Gen2Migration.CallOpts, arg0)
}

// HashToSign is a free data retrieval call binding the contract method 0x0a96cb49.
//
// Solidity: function hashToSign(address _destination) constant returns(bytes32)
func (_Gen2Migration *Gen2MigrationCaller) HashToSign(opts *bind.CallOpts, _destination common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Gen2Migration.contract.Call(opts, out, "hashToSign", _destination)
	return *ret0, err
}

// HashToSign is a free data retrieval call binding the contract method 0x0a96cb49.
//
// Solidity: function hashToSign(address _destination) constant returns(bytes32)
func (_Gen2Migration *Gen2MigrationSession) HashToSign(_destination common.Address) ([32]byte, error) {
	return _Gen2Migration.Contract.HashToSign(&_Gen2Migration.CallOpts, _destination)
}

// HashToSign is a free data retrieval call binding the contract method 0x0a96cb49.
//
// Solidity: function hashToSign(address _destination) constant returns(bytes32)
func (_Gen2Migration *Gen2MigrationCallerSession) HashToSign(_destination common.Address) ([32]byte, error) {
	return _Gen2Migration.Contract.HashToSign(&_Gen2Migration.CallOpts, _destination)
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationCaller) ItemCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gen2Migration.contract.Call(opts, out, "itemCount")
	return *ret0, err
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationSession) ItemCount() (*big.Int, error) {
	return _Gen2Migration.Contract.ItemCount(&_Gen2Migration.CallOpts)
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationCallerSession) ItemCount() (*big.Int, error) {
	return _Gen2Migration.Contract.ItemCount(&_Gen2Migration.CallOpts)
}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() constant returns(address)
func (_Gen2Migration *Gen2MigrationCaller) SignerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gen2Migration.contract.Call(opts, out, "signerAddress")
	return *ret0, err
}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() constant returns(address)
func (_Gen2Migration *Gen2MigrationSession) SignerAddress() (common.Address, error) {
	return _Gen2Migration.Contract.SignerAddress(&_Gen2Migration.CallOpts)
}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() constant returns(address)
func (_Gen2Migration *Gen2MigrationCallerSession) SignerAddress() (common.Address, error) {
	return _Gen2Migration.Contract.SignerAddress(&_Gen2Migration.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gen2Migration.contract.Call(opts, out, "totalAmount")
	return *ret0, err
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationSession) TotalAmount() (*big.Int, error) {
	return _Gen2Migration.Contract.TotalAmount(&_Gen2Migration.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Gen2Migration *Gen2MigrationCallerSession) TotalAmount() (*big.Int, error) {
	return _Gen2Migration.Contract.TotalAmount(&_Gen2Migration.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_Gen2Migration *Gen2MigrationCaller) TreasuryProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gen2Migration.contract.Call(opts, out, "treasury_proxy")
	return *ret0, err
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_Gen2Migration *Gen2MigrationSession) TreasuryProxy() (common.Address, error) {
	return _Gen2Migration.Contract.TreasuryProxy(&_Gen2Migration.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0xa2731784.
//
// Solidity: function treasury_proxy() constant returns(address)
func (_Gen2Migration *Gen2MigrationCallerSession) TreasuryProxy() (common.Address, error) {
	return _Gen2Migration.Contract.TreasuryProxy(&_Gen2Migration.CallOpts)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x476ce0c3.
//
// Solidity: function verifyClaim(uint256 _item_id, address _destination, uint8 sig_v, bytes32 sig_r, bytes32 sig_s) constant returns(uint256 amount)
func (_Gen2Migration *Gen2MigrationCaller) VerifyClaim(opts *bind.CallOpts, _item_id *big.Int, _destination common.Address, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gen2Migration.contract.Call(opts, out, "verifyClaim", _item_id, _destination, sig_v, sig_r, sig_s)
	return *ret0, err
}

// VerifyClaim is a free data retrieval call binding the contract method 0x476ce0c3.
//
// Solidity: function verifyClaim(uint256 _item_id, address _destination, uint8 sig_v, bytes32 sig_r, bytes32 sig_s) constant returns(uint256 amount)
func (_Gen2Migration *Gen2MigrationSession) VerifyClaim(_item_id *big.Int, _destination common.Address, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*big.Int, error) {
	return _Gen2Migration.Contract.VerifyClaim(&_Gen2Migration.CallOpts, _item_id, _destination, sig_v, sig_r, sig_s)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x476ce0c3.
//
// Solidity: function verifyClaim(uint256 _item_id, address _destination, uint8 sig_v, bytes32 sig_r, bytes32 sig_s) constant returns(uint256 amount)
func (_Gen2Migration *Gen2MigrationCallerSession) VerifyClaim(_item_id *big.Int, _destination common.Address, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*big.Int, error) {
	return _Gen2Migration.Contract.VerifyClaim(&_Gen2Migration.CallOpts, _item_id, _destination, sig_v, sig_r, sig_s)
}

// Claim is a paid mutator transaction binding the contract method 0xf7121490.
//
// Solidity: function claim(uint256 _item_id, address _destination, uint8 sig_v, bytes32 sig_r, bytes32 sig_s) returns()
func (_Gen2Migration *Gen2MigrationTransactor) Claim(opts *bind.TransactOpts, _item_id *big.Int, _destination common.Address, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _Gen2Migration.contract.Transact(opts, "claim", _item_id, _destination, sig_v, sig_r, sig_s)
}

// Claim is a paid mutator transaction binding the contract method 0xf7121490.
//
// Solidity: function claim(uint256 _item_id, address _destination, uint8 sig_v, bytes32 sig_r, bytes32 sig_s) returns()
func (_Gen2Migration *Gen2MigrationSession) Claim(_item_id *big.Int, _destination common.Address, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _Gen2Migration.Contract.Claim(&_Gen2Migration.TransactOpts, _item_id, _destination, sig_v, sig_r, sig_s)
}

// Claim is a paid mutator transaction binding the contract method 0xf7121490.
//
// Solidity: function claim(uint256 _item_id, address _destination, uint8 sig_v, bytes32 sig_r, bytes32 sig_s) returns()
func (_Gen2Migration *Gen2MigrationTransactorSession) Claim(_item_id *big.Int, _destination common.Address, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _Gen2Migration.Contract.Claim(&_Gen2Migration.TransactOpts, _item_id, _destination, sig_v, sig_r, sig_s)
}

// Drain is a paid mutator transaction binding the contract method 0x9890220b.
//
// Solidity: function drain() returns()
func (_Gen2Migration *Gen2MigrationTransactor) Drain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gen2Migration.contract.Transact(opts, "drain")
}

// Drain is a paid mutator transaction binding the contract method 0x9890220b.
//
// Solidity: function drain() returns()
func (_Gen2Migration *Gen2MigrationSession) Drain() (*types.Transaction, error) {
	return _Gen2Migration.Contract.Drain(&_Gen2Migration.TransactOpts)
}

// Drain is a paid mutator transaction binding the contract method 0x9890220b.
//
// Solidity: function drain() returns()
func (_Gen2Migration *Gen2MigrationTransactorSession) Drain() (*types.Transaction, error) {
	return _Gen2Migration.Contract.Drain(&_Gen2Migration.TransactOpts)
}

// SetSnapshot is a paid mutator transaction binding the contract method 0xe38e6e5e.
//
// Solidity: function setSnapshot(bytes20[] _owners, uint256[] _amounts) returns()
func (_Gen2Migration *Gen2MigrationTransactor) SetSnapshot(opts *bind.TransactOpts, _owners [][20]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _Gen2Migration.contract.Transact(opts, "setSnapshot", _owners, _amounts)
}

// SetSnapshot is a paid mutator transaction binding the contract method 0xe38e6e5e.
//
// Solidity: function setSnapshot(bytes20[] _owners, uint256[] _amounts) returns()
func (_Gen2Migration *Gen2MigrationSession) SetSnapshot(_owners [][20]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _Gen2Migration.Contract.SetSnapshot(&_Gen2Migration.TransactOpts, _owners, _amounts)
}

// SetSnapshot is a paid mutator transaction binding the contract method 0xe38e6e5e.
//
// Solidity: function setSnapshot(bytes20[] _owners, uint256[] _amounts) returns()
func (_Gen2Migration *Gen2MigrationTransactorSession) SetSnapshot(_owners [][20]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _Gen2Migration.Contract.SetSnapshot(&_Gen2Migration.TransactOpts, _owners, _amounts)
}

// Gen2MigrationMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the Gen2Migration contract.
type Gen2MigrationMigratedIterator struct {
	Event *Gen2MigrationMigrated // Event containing the contract specifics and raw log

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
func (it *Gen2MigrationMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Gen2MigrationMigrated)
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
		it.Event = new(Gen2MigrationMigrated)
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
func (it *Gen2MigrationMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Gen2MigrationMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Gen2MigrationMigrated represents a Migrated event raised by the Gen2Migration contract.
type Gen2MigrationMigrated struct {
	ItemId      *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xfd90f074a575cd9336850f79afca1e89f5ca1bf434d82a21ca5d6d4a87a724a4.
//
// Solidity: event Migrated(uint256 item_id, address destination, uint256 amount)
func (_Gen2Migration *Gen2MigrationFilterer) FilterMigrated(opts *bind.FilterOpts) (*Gen2MigrationMigratedIterator, error) {

	logs, sub, err := _Gen2Migration.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &Gen2MigrationMigratedIterator{contract: _Gen2Migration.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xfd90f074a575cd9336850f79afca1e89f5ca1bf434d82a21ca5d6d4a87a724a4.
//
// Solidity: event Migrated(uint256 item_id, address destination, uint256 amount)
func (_Gen2Migration *Gen2MigrationFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *Gen2MigrationMigrated) (event.Subscription, error) {

	logs, sub, err := _Gen2Migration.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Gen2MigrationMigrated)
				if err := _Gen2Migration.contract.UnpackLog(event, "Migrated", log); err != nil {
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
