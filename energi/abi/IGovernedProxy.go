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

// IGovernedProxyABI is the input ABI used to generate the binding from.
const IGovernedProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIGovernedContract\",\"name\":\"impl\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIUpgradeProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"UpgradeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIGovernedContract\",\"name\":\"impl\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIUpgradeProposal\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"collectUpgradeProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"impl\",\"outputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"listUpgradeProposals\",\"outputs\":[{\"internalType\":\"contractIUpgradeProposal[]\",\"name\":\"proposals\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_newImpl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"proposeUpgrade\",\"outputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIUpgradeProposal\",\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"upgradeProposalImpl\",\"outputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"new_impl\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IGovernedProxy is an auto generated Go binding around an Ethereum contract.
type IGovernedProxy struct {
	IGovernedProxyCaller     // Read-only binding to the contract
	IGovernedProxyTransactor // Write-only binding to the contract
	IGovernedProxyFilterer   // Log filterer for contract events
}

// IGovernedProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGovernedProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovernedProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGovernedProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovernedProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGovernedProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovernedProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGovernedProxySession struct {
	Contract     *IGovernedProxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGovernedProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGovernedProxyCallerSession struct {
	Contract *IGovernedProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IGovernedProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGovernedProxyTransactorSession struct {
	Contract     *IGovernedProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IGovernedProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGovernedProxyRaw struct {
	Contract *IGovernedProxy // Generic contract binding to access the raw methods on
}

// IGovernedProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGovernedProxyCallerRaw struct {
	Contract *IGovernedProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IGovernedProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGovernedProxyTransactorRaw struct {
	Contract *IGovernedProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGovernedProxy creates a new instance of IGovernedProxy, bound to a specific deployed contract.
func NewIGovernedProxy(address common.Address, backend bind.ContractBackend) (*IGovernedProxy, error) {
	contract, err := bindIGovernedProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGovernedProxy{IGovernedProxyCaller: IGovernedProxyCaller{contract: contract}, IGovernedProxyTransactor: IGovernedProxyTransactor{contract: contract}, IGovernedProxyFilterer: IGovernedProxyFilterer{contract: contract}}, nil
}

// NewIGovernedProxyCaller creates a new read-only instance of IGovernedProxy, bound to a specific deployed contract.
func NewIGovernedProxyCaller(address common.Address, caller bind.ContractCaller) (*IGovernedProxyCaller, error) {
	contract, err := bindIGovernedProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGovernedProxyCaller{contract: contract}, nil
}

// NewIGovernedProxyTransactor creates a new write-only instance of IGovernedProxy, bound to a specific deployed contract.
func NewIGovernedProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IGovernedProxyTransactor, error) {
	contract, err := bindIGovernedProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGovernedProxyTransactor{contract: contract}, nil
}

// NewIGovernedProxyFilterer creates a new log filterer instance of IGovernedProxy, bound to a specific deployed contract.
func NewIGovernedProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IGovernedProxyFilterer, error) {
	contract, err := bindIGovernedProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGovernedProxyFilterer{contract: contract}, nil
}

// bindIGovernedProxy binds a generic wrapper to an already deployed contract.
func bindIGovernedProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGovernedProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGovernedProxy *IGovernedProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGovernedProxy.Contract.IGovernedProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGovernedProxy *IGovernedProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.IGovernedProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGovernedProxy *IGovernedProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.IGovernedProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGovernedProxy *IGovernedProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGovernedProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGovernedProxy *IGovernedProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGovernedProxy *IGovernedProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.contract.Transact(opts, method, params...)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_IGovernedProxy *IGovernedProxyCaller) Impl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IGovernedProxy.contract.Call(opts, out, "impl")
	return *ret0, err
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_IGovernedProxy *IGovernedProxySession) Impl() (common.Address, error) {
	return _IGovernedProxy.Contract.Impl(&_IGovernedProxy.CallOpts)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_IGovernedProxy *IGovernedProxyCallerSession) Impl() (common.Address, error) {
	return _IGovernedProxy.Contract.Impl(&_IGovernedProxy.CallOpts)
}

// ListUpgradeProposals is a free data retrieval call binding the contract method 0xb364595e.
//
// Solidity: function listUpgradeProposals() constant returns(address[] proposals)
func (_IGovernedProxy *IGovernedProxyCaller) ListUpgradeProposals(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IGovernedProxy.contract.Call(opts, out, "listUpgradeProposals")
	return *ret0, err
}

// ListUpgradeProposals is a free data retrieval call binding the contract method 0xb364595e.
//
// Solidity: function listUpgradeProposals() constant returns(address[] proposals)
func (_IGovernedProxy *IGovernedProxySession) ListUpgradeProposals() ([]common.Address, error) {
	return _IGovernedProxy.Contract.ListUpgradeProposals(&_IGovernedProxy.CallOpts)
}

// ListUpgradeProposals is a free data retrieval call binding the contract method 0xb364595e.
//
// Solidity: function listUpgradeProposals() constant returns(address[] proposals)
func (_IGovernedProxy *IGovernedProxyCallerSession) ListUpgradeProposals() ([]common.Address, error) {
	return _IGovernedProxy.Contract.ListUpgradeProposals(&_IGovernedProxy.CallOpts)
}

// UpgradeProposalImpl is a free data retrieval call binding the contract method 0x6d5b6c44.
//
// Solidity: function upgradeProposalImpl(address _proposal) constant returns(address new_impl)
func (_IGovernedProxy *IGovernedProxyCaller) UpgradeProposalImpl(opts *bind.CallOpts, _proposal common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IGovernedProxy.contract.Call(opts, out, "upgradeProposalImpl", _proposal)
	return *ret0, err
}

// UpgradeProposalImpl is a free data retrieval call binding the contract method 0x6d5b6c44.
//
// Solidity: function upgradeProposalImpl(address _proposal) constant returns(address new_impl)
func (_IGovernedProxy *IGovernedProxySession) UpgradeProposalImpl(_proposal common.Address) (common.Address, error) {
	return _IGovernedProxy.Contract.UpgradeProposalImpl(&_IGovernedProxy.CallOpts, _proposal)
}

// UpgradeProposalImpl is a free data retrieval call binding the contract method 0x6d5b6c44.
//
// Solidity: function upgradeProposalImpl(address _proposal) constant returns(address new_impl)
func (_IGovernedProxy *IGovernedProxyCallerSession) UpgradeProposalImpl(_proposal common.Address) (common.Address, error) {
	return _IGovernedProxy.Contract.UpgradeProposalImpl(&_IGovernedProxy.CallOpts, _proposal)
}

// CollectUpgradeProposal is a paid mutator transaction binding the contract method 0xa1b0e476.
//
// Solidity: function collectUpgradeProposal(address _proposal) returns()
func (_IGovernedProxy *IGovernedProxyTransactor) CollectUpgradeProposal(opts *bind.TransactOpts, _proposal common.Address) (*types.Transaction, error) {
	return _IGovernedProxy.contract.Transact(opts, "collectUpgradeProposal", _proposal)
}

// CollectUpgradeProposal is a paid mutator transaction binding the contract method 0xa1b0e476.
//
// Solidity: function collectUpgradeProposal(address _proposal) returns()
func (_IGovernedProxy *IGovernedProxySession) CollectUpgradeProposal(_proposal common.Address) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.CollectUpgradeProposal(&_IGovernedProxy.TransactOpts, _proposal)
}

// CollectUpgradeProposal is a paid mutator transaction binding the contract method 0xa1b0e476.
//
// Solidity: function collectUpgradeProposal(address _proposal) returns()
func (_IGovernedProxy *IGovernedProxyTransactorSession) CollectUpgradeProposal(_proposal common.Address) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.CollectUpgradeProposal(&_IGovernedProxy.TransactOpts, _proposal)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns(address)
func (_IGovernedProxy *IGovernedProxyTransactor) ProposeUpgrade(opts *bind.TransactOpts, _newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _IGovernedProxy.contract.Transact(opts, "proposeUpgrade", _newImpl, _period)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns(address)
func (_IGovernedProxy *IGovernedProxySession) ProposeUpgrade(_newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.ProposeUpgrade(&_IGovernedProxy.TransactOpts, _newImpl, _period)
}

// ProposeUpgrade is a paid mutator transaction binding the contract method 0x5b6dee4c.
//
// Solidity: function proposeUpgrade(address _newImpl, uint256 _period) returns(address)
func (_IGovernedProxy *IGovernedProxyTransactorSession) ProposeUpgrade(_newImpl common.Address, _period *big.Int) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.ProposeUpgrade(&_IGovernedProxy.TransactOpts, _newImpl, _period)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_IGovernedProxy *IGovernedProxyTransactor) Upgrade(opts *bind.TransactOpts, _proposal common.Address) (*types.Transaction, error) {
	return _IGovernedProxy.contract.Transact(opts, "upgrade", _proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_IGovernedProxy *IGovernedProxySession) Upgrade(_proposal common.Address) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.Upgrade(&_IGovernedProxy.TransactOpts, _proposal)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proposal) returns()
func (_IGovernedProxy *IGovernedProxyTransactorSession) Upgrade(_proposal common.Address) (*types.Transaction, error) {
	return _IGovernedProxy.Contract.Upgrade(&_IGovernedProxy.TransactOpts, _proposal)
}

// IGovernedProxyUpgradeProposalIterator is returned from FilterUpgradeProposal and is used to iterate over the raw logs and unpacked data for UpgradeProposal events raised by the IGovernedProxy contract.
type IGovernedProxyUpgradeProposalIterator struct {
	Event *IGovernedProxyUpgradeProposal // Event containing the contract specifics and raw log

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
func (it *IGovernedProxyUpgradeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovernedProxyUpgradeProposal)
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
		it.Event = new(IGovernedProxyUpgradeProposal)
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
func (it *IGovernedProxyUpgradeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovernedProxyUpgradeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovernedProxyUpgradeProposal represents a UpgradeProposal event raised by the IGovernedProxy contract.
type IGovernedProxyUpgradeProposal struct {
	Impl     common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradeProposal is a free log retrieval operation binding the contract event 0x812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763.
//
// Solidity: event UpgradeProposal(address indexed impl, address proposal)
func (_IGovernedProxy *IGovernedProxyFilterer) FilterUpgradeProposal(opts *bind.FilterOpts, impl []common.Address) (*IGovernedProxyUpgradeProposalIterator, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _IGovernedProxy.contract.FilterLogs(opts, "UpgradeProposal", implRule)
	if err != nil {
		return nil, err
	}
	return &IGovernedProxyUpgradeProposalIterator{contract: _IGovernedProxy.contract, event: "UpgradeProposal", logs: logs, sub: sub}, nil
}

// WatchUpgradeProposal is a free log subscription operation binding the contract event 0x812eb2689eecf94cfb55caf4a123ea76c6d93eef07dd54a5273b7a4949f7d763.
//
// Solidity: event UpgradeProposal(address indexed impl, address proposal)
func (_IGovernedProxy *IGovernedProxyFilterer) WatchUpgradeProposal(opts *bind.WatchOpts, sink chan<- *IGovernedProxyUpgradeProposal, impl []common.Address) (event.Subscription, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _IGovernedProxy.contract.WatchLogs(opts, "UpgradeProposal", implRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovernedProxyUpgradeProposal)
				if err := _IGovernedProxy.contract.UnpackLog(event, "UpgradeProposal", log); err != nil {
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

// IGovernedProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IGovernedProxy contract.
type IGovernedProxyUpgradedIterator struct {
	Event *IGovernedProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *IGovernedProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovernedProxyUpgraded)
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
		it.Event = new(IGovernedProxyUpgraded)
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
func (it *IGovernedProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovernedProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovernedProxyUpgraded represents a Upgraded event raised by the IGovernedProxy contract.
type IGovernedProxyUpgraded struct {
	Impl     common.Address
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed impl, address proposal)
func (_IGovernedProxy *IGovernedProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, impl []common.Address) (*IGovernedProxyUpgradedIterator, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _IGovernedProxy.contract.FilterLogs(opts, "Upgraded", implRule)
	if err != nil {
		return nil, err
	}
	return &IGovernedProxyUpgradedIterator{contract: _IGovernedProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x5d611f318680d00598bb735d61bacf0c514c6b50e1e5ad30040a4df2b12791c7.
//
// Solidity: event Upgraded(address indexed impl, address proposal)
func (_IGovernedProxy *IGovernedProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IGovernedProxyUpgraded, impl []common.Address) (event.Subscription, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}

	logs, sub, err := _IGovernedProxy.contract.WatchLogs(opts, "Upgraded", implRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovernedProxyUpgraded)
				if err := _IGovernedProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
