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

// ITreasuryABI is the input ABI used to generate the binding from.
const ITreasuryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ref_uuid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIBudgetProposal\",\"name\":\"proposal\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payout_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BudgetProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Contribution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ref_uuid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIBudgetProposal\",\"name\":\"proposal\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Payout\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"isSuperblock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"listProposals\",\"outputs\":[{\"internalType\":\"contractIBudgetProposal[]\",\"name\":\"proposals\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIBudgetProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"proposal_uuid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ref_uuid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"contractIBudgetProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ref_uuid\",\"type\":\"uint256\"}],\"name\":\"uuid_proposal\",\"outputs\":[{\"internalType\":\"contractIBudgetProposal\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITreasury is an auto generated Go binding around an Ethereum contract.
type ITreasury struct {
	ITreasuryCaller     // Read-only binding to the contract
	ITreasuryTransactor // Write-only binding to the contract
	ITreasuryFilterer   // Log filterer for contract events
}

// ITreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITreasurySession struct {
	Contract     *ITreasury        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITreasuryCallerSession struct {
	Contract *ITreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ITreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITreasuryTransactorSession struct {
	Contract     *ITreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ITreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITreasuryRaw struct {
	Contract *ITreasury // Generic contract binding to access the raw methods on
}

// ITreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITreasuryCallerRaw struct {
	Contract *ITreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// ITreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITreasuryTransactorRaw struct {
	Contract *ITreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITreasury creates a new instance of ITreasury, bound to a specific deployed contract.
func NewITreasury(address common.Address, backend bind.ContractBackend) (*ITreasury, error) {
	contract, err := bindITreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITreasury{ITreasuryCaller: ITreasuryCaller{contract: contract}, ITreasuryTransactor: ITreasuryTransactor{contract: contract}, ITreasuryFilterer: ITreasuryFilterer{contract: contract}}, nil
}

// NewITreasuryCaller creates a new read-only instance of ITreasury, bound to a specific deployed contract.
func NewITreasuryCaller(address common.Address, caller bind.ContractCaller) (*ITreasuryCaller, error) {
	contract, err := bindITreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITreasuryCaller{contract: contract}, nil
}

// NewITreasuryTransactor creates a new write-only instance of ITreasury, bound to a specific deployed contract.
func NewITreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*ITreasuryTransactor, error) {
	contract, err := bindITreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITreasuryTransactor{contract: contract}, nil
}

// NewITreasuryFilterer creates a new log filterer instance of ITreasury, bound to a specific deployed contract.
func NewITreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*ITreasuryFilterer, error) {
	contract, err := bindITreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITreasuryFilterer{contract: contract}, nil
}

// bindITreasury binds a generic wrapper to an already deployed contract.
func bindITreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITreasuryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITreasury *ITreasuryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITreasury.Contract.ITreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITreasury *ITreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITreasury.Contract.ITreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITreasury *ITreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITreasury.Contract.ITreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITreasury *ITreasuryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITreasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITreasury *ITreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITreasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITreasury *ITreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITreasury.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256 amount)
func (_ITreasury *ITreasuryCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ITreasury.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256 amount)
func (_ITreasury *ITreasurySession) Balance() (*big.Int, error) {
	return _ITreasury.Contract.Balance(&_ITreasury.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256 amount)
func (_ITreasury *ITreasuryCallerSession) Balance() (*big.Int, error) {
	return _ITreasury.Contract.Balance(&_ITreasury.CallOpts)
}

// IsSuperblock is a free data retrieval call binding the contract method 0x52782d86.
//
// Solidity: function isSuperblock(uint256 _blockNumber) constant returns(bool)
func (_ITreasury *ITreasuryCaller) IsSuperblock(opts *bind.CallOpts, _blockNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ITreasury.contract.Call(opts, out, "isSuperblock", _blockNumber)
	return *ret0, err
}

// IsSuperblock is a free data retrieval call binding the contract method 0x52782d86.
//
// Solidity: function isSuperblock(uint256 _blockNumber) constant returns(bool)
func (_ITreasury *ITreasurySession) IsSuperblock(_blockNumber *big.Int) (bool, error) {
	return _ITreasury.Contract.IsSuperblock(&_ITreasury.CallOpts, _blockNumber)
}

// IsSuperblock is a free data retrieval call binding the contract method 0x52782d86.
//
// Solidity: function isSuperblock(uint256 _blockNumber) constant returns(bool)
func (_ITreasury *ITreasuryCallerSession) IsSuperblock(_blockNumber *big.Int) (bool, error) {
	return _ITreasury.Contract.IsSuperblock(&_ITreasury.CallOpts, _blockNumber)
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() constant returns(address[] proposals)
func (_ITreasury *ITreasuryCaller) ListProposals(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ITreasury.contract.Call(opts, out, "listProposals")
	return *ret0, err
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() constant returns(address[] proposals)
func (_ITreasury *ITreasurySession) ListProposals() ([]common.Address, error) {
	return _ITreasury.Contract.ListProposals(&_ITreasury.CallOpts)
}

// ListProposals is a free data retrieval call binding the contract method 0x26c7eb1c.
//
// Solidity: function listProposals() constant returns(address[] proposals)
func (_ITreasury *ITreasuryCallerSession) ListProposals() ([]common.Address, error) {
	return _ITreasury.Contract.ListProposals(&_ITreasury.CallOpts)
}

// ProposalUuid is a free data retrieval call binding the contract method 0xcaef1d5a.
//
// Solidity: function proposal_uuid(address proposal) constant returns(uint256)
func (_ITreasury *ITreasuryCaller) ProposalUuid(opts *bind.CallOpts, proposal common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ITreasury.contract.Call(opts, out, "proposal_uuid", proposal)
	return *ret0, err
}

// ProposalUuid is a free data retrieval call binding the contract method 0xcaef1d5a.
//
// Solidity: function proposal_uuid(address proposal) constant returns(uint256)
func (_ITreasury *ITreasurySession) ProposalUuid(proposal common.Address) (*big.Int, error) {
	return _ITreasury.Contract.ProposalUuid(&_ITreasury.CallOpts, proposal)
}

// ProposalUuid is a free data retrieval call binding the contract method 0xcaef1d5a.
//
// Solidity: function proposal_uuid(address proposal) constant returns(uint256)
func (_ITreasury *ITreasuryCallerSession) ProposalUuid(proposal common.Address) (*big.Int, error) {
	return _ITreasury.Contract.ProposalUuid(&_ITreasury.CallOpts, proposal)
}

// UuidProposal is a free data retrieval call binding the contract method 0x5c099215.
//
// Solidity: function uuid_proposal(uint256 _ref_uuid) constant returns(address)
func (_ITreasury *ITreasuryCaller) UuidProposal(opts *bind.CallOpts, _ref_uuid *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ITreasury.contract.Call(opts, out, "uuid_proposal", _ref_uuid)
	return *ret0, err
}

// UuidProposal is a free data retrieval call binding the contract method 0x5c099215.
//
// Solidity: function uuid_proposal(uint256 _ref_uuid) constant returns(address)
func (_ITreasury *ITreasurySession) UuidProposal(_ref_uuid *big.Int) (common.Address, error) {
	return _ITreasury.Contract.UuidProposal(&_ITreasury.CallOpts, _ref_uuid)
}

// UuidProposal is a free data retrieval call binding the contract method 0x5c099215.
//
// Solidity: function uuid_proposal(uint256 _ref_uuid) constant returns(address)
func (_ITreasury *ITreasuryCallerSession) UuidProposal(_ref_uuid *big.Int) (common.Address, error) {
	return _ITreasury.Contract.UuidProposal(&_ITreasury.CallOpts, _ref_uuid)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() returns()
func (_ITreasury *ITreasuryTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITreasury.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() returns()
func (_ITreasury *ITreasurySession) Contribute() (*types.Transaction, error) {
	return _ITreasury.Contract.Contribute(&_ITreasury.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() returns()
func (_ITreasury *ITreasuryTransactorSession) Contribute() (*types.Transaction, error) {
	return _ITreasury.Contract.Contribute(&_ITreasury.TransactOpts)
}

// Propose is a paid mutator transaction binding the contract method 0x40376d56.
//
// Solidity: function propose(uint256 _amount, uint256 _ref_uuid, uint256 _period) returns(address proposal)
func (_ITreasury *ITreasuryTransactor) Propose(opts *bind.TransactOpts, _amount *big.Int, _ref_uuid *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _ITreasury.contract.Transact(opts, "propose", _amount, _ref_uuid, _period)
}

// Propose is a paid mutator transaction binding the contract method 0x40376d56.
//
// Solidity: function propose(uint256 _amount, uint256 _ref_uuid, uint256 _period) returns(address proposal)
func (_ITreasury *ITreasurySession) Propose(_amount *big.Int, _ref_uuid *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _ITreasury.Contract.Propose(&_ITreasury.TransactOpts, _amount, _ref_uuid, _period)
}

// Propose is a paid mutator transaction binding the contract method 0x40376d56.
//
// Solidity: function propose(uint256 _amount, uint256 _ref_uuid, uint256 _period) returns(address proposal)
func (_ITreasury *ITreasuryTransactorSession) Propose(_amount *big.Int, _ref_uuid *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _ITreasury.Contract.Propose(&_ITreasury.TransactOpts, _amount, _ref_uuid, _period)
}

// ITreasuryBudgetProposalIterator is returned from FilterBudgetProposal and is used to iterate over the raw logs and unpacked data for BudgetProposal events raised by the ITreasury contract.
type ITreasuryBudgetProposalIterator struct {
	Event *ITreasuryBudgetProposal // Event containing the contract specifics and raw log

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
func (it *ITreasuryBudgetProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITreasuryBudgetProposal)
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
		it.Event = new(ITreasuryBudgetProposal)
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
func (it *ITreasuryBudgetProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITreasuryBudgetProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITreasuryBudgetProposal represents a BudgetProposal event raised by the ITreasury contract.
type ITreasuryBudgetProposal struct {
	RefUuid       *big.Int
	Proposal      common.Address
	PayoutAddress common.Address
	Amount        *big.Int
	Deadline      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBudgetProposal is a free log retrieval operation binding the contract event 0xa94befa3ce181974130fe36e854a282c70fb75e4decb42bfecdecbd6f88f0f5d.
//
// Solidity: event BudgetProposal(uint256 indexed ref_uuid, address proposal, address payout_address, uint256 amount, uint256 deadline)
func (_ITreasury *ITreasuryFilterer) FilterBudgetProposal(opts *bind.FilterOpts, ref_uuid []*big.Int) (*ITreasuryBudgetProposalIterator, error) {

	var ref_uuidRule []interface{}
	for _, ref_uuidItem := range ref_uuid {
		ref_uuidRule = append(ref_uuidRule, ref_uuidItem)
	}

	logs, sub, err := _ITreasury.contract.FilterLogs(opts, "BudgetProposal", ref_uuidRule)
	if err != nil {
		return nil, err
	}
	return &ITreasuryBudgetProposalIterator{contract: _ITreasury.contract, event: "BudgetProposal", logs: logs, sub: sub}, nil
}

// WatchBudgetProposal is a free log subscription operation binding the contract event 0xa94befa3ce181974130fe36e854a282c70fb75e4decb42bfecdecbd6f88f0f5d.
//
// Solidity: event BudgetProposal(uint256 indexed ref_uuid, address proposal, address payout_address, uint256 amount, uint256 deadline)
func (_ITreasury *ITreasuryFilterer) WatchBudgetProposal(opts *bind.WatchOpts, sink chan<- *ITreasuryBudgetProposal, ref_uuid []*big.Int) (event.Subscription, error) {

	var ref_uuidRule []interface{}
	for _, ref_uuidItem := range ref_uuid {
		ref_uuidRule = append(ref_uuidRule, ref_uuidItem)
	}

	logs, sub, err := _ITreasury.contract.WatchLogs(opts, "BudgetProposal", ref_uuidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITreasuryBudgetProposal)
				if err := _ITreasury.contract.UnpackLog(event, "BudgetProposal", log); err != nil {
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

// ITreasuryContributionIterator is returned from FilterContribution and is used to iterate over the raw logs and unpacked data for Contribution events raised by the ITreasury contract.
type ITreasuryContributionIterator struct {
	Event *ITreasuryContribution // Event containing the contract specifics and raw log

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
func (it *ITreasuryContributionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITreasuryContribution)
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
		it.Event = new(ITreasuryContribution)
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
func (it *ITreasuryContributionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITreasuryContributionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITreasuryContribution represents a Contribution event raised by the ITreasury contract.
type ITreasuryContribution struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContribution is a free log retrieval operation binding the contract event 0x4d154d4aae216bed6d0926db77c00df2b57c6b5ba4eee05775de20facede3a7b.
//
// Solidity: event Contribution(address from, uint256 amount)
func (_ITreasury *ITreasuryFilterer) FilterContribution(opts *bind.FilterOpts) (*ITreasuryContributionIterator, error) {

	logs, sub, err := _ITreasury.contract.FilterLogs(opts, "Contribution")
	if err != nil {
		return nil, err
	}
	return &ITreasuryContributionIterator{contract: _ITreasury.contract, event: "Contribution", logs: logs, sub: sub}, nil
}

// WatchContribution is a free log subscription operation binding the contract event 0x4d154d4aae216bed6d0926db77c00df2b57c6b5ba4eee05775de20facede3a7b.
//
// Solidity: event Contribution(address from, uint256 amount)
func (_ITreasury *ITreasuryFilterer) WatchContribution(opts *bind.WatchOpts, sink chan<- *ITreasuryContribution) (event.Subscription, error) {

	logs, sub, err := _ITreasury.contract.WatchLogs(opts, "Contribution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITreasuryContribution)
				if err := _ITreasury.contract.UnpackLog(event, "Contribution", log); err != nil {
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

// ITreasuryPayoutIterator is returned from FilterPayout and is used to iterate over the raw logs and unpacked data for Payout events raised by the ITreasury contract.
type ITreasuryPayoutIterator struct {
	Event *ITreasuryPayout // Event containing the contract specifics and raw log

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
func (it *ITreasuryPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITreasuryPayout)
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
		it.Event = new(ITreasuryPayout)
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
func (it *ITreasuryPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITreasuryPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITreasuryPayout represents a Payout event raised by the ITreasury contract.
type ITreasuryPayout struct {
	RefUuid  *big.Int
	Proposal common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayout is a free log retrieval operation binding the contract event 0x5f7341a552ae2d452b071917104c05fbac3663936a69be768a05c40605056e7d.
//
// Solidity: event Payout(uint256 indexed ref_uuid, address proposal, uint256 amount)
func (_ITreasury *ITreasuryFilterer) FilterPayout(opts *bind.FilterOpts, ref_uuid []*big.Int) (*ITreasuryPayoutIterator, error) {

	var ref_uuidRule []interface{}
	for _, ref_uuidItem := range ref_uuid {
		ref_uuidRule = append(ref_uuidRule, ref_uuidItem)
	}

	logs, sub, err := _ITreasury.contract.FilterLogs(opts, "Payout", ref_uuidRule)
	if err != nil {
		return nil, err
	}
	return &ITreasuryPayoutIterator{contract: _ITreasury.contract, event: "Payout", logs: logs, sub: sub}, nil
}

// WatchPayout is a free log subscription operation binding the contract event 0x5f7341a552ae2d452b071917104c05fbac3663936a69be768a05c40605056e7d.
//
// Solidity: event Payout(uint256 indexed ref_uuid, address proposal, uint256 amount)
func (_ITreasury *ITreasuryFilterer) WatchPayout(opts *bind.WatchOpts, sink chan<- *ITreasuryPayout, ref_uuid []*big.Int) (event.Subscription, error) {

	var ref_uuidRule []interface{}
	for _, ref_uuidItem := range ref_uuid {
		ref_uuidRule = append(ref_uuidRule, ref_uuidItem)
	}

	logs, sub, err := _ITreasury.contract.WatchLogs(opts, "Payout", ref_uuidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITreasuryPayout)
				if err := _ITreasury.contract.UnpackLog(event, "Payout", log); err != nil {
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
