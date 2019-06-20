// Copyright 2019 The Energi Core Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

// Energi Governance system is the fundamental part of Energi Core.

// NOTE: It's not allowed to change the compiler due to byte-to-byte
//       match requirement.
pragma solidity 0.5.9;
//pragma experimental SMTChecker;

import { IGovernedContract } from "./IGovernedContract.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IProposal } from "./IProposal.sol";
import { ISporkRegistry } from "./ISporkRegistry.sol";
import { NonReentrant } from "./NonReentrant.sol";

/**
 * SC-9: This contract has no chance of being updated. It must be stupid simple.
 *
 * If another upgrade logic is required in the future - it can be done as proxy stage II.
 */
contract GovernedProxy is
    IGovernedContract,
    IGovernedProxy,
    NonReentrant
{
    modifier senderOrigin {
        // Internal calls are expected to use impl directly.
        // That's due to use of call() instead of delegatecall() on purpose.
        // solium-disable-next-line security/no-tx-origin
        require(tx.origin == msg.sender, "Only direct calls are allowed!");
        _;
    }

    IGovernedContract public impl;
    IGovernedProxy public spork_proxy;
    mapping(address => IGovernedContract) public upgrade_proposals;

    constructor(IGovernedContract _impl, IGovernedProxy _sporkProxy) public {
        impl = _impl;
        spork_proxy = _sporkProxy;
    }

    /**
     * Pre-create a new contract first.
     * Then propose upgrade based on that.
     */
    function proposeUpgrade(IGovernedContract _newImpl, uint _period)
        external payable
        senderOrigin
        noReentry
    {
        require(_newImpl != impl, "Already active!");
        require(_newImpl.proxy() == address(this), "Wrong proxy!");

        ISporkRegistry spork_reg = ISporkRegistry(address(spork_proxy.impl()));
        IProposal proposal = spork_reg.createUpgradeProposal.value(msg.value)(_newImpl, _period, msg.sender);

        upgrade_proposals[address(proposal)] = _newImpl;

        emit UpgradeProposal(address(_newImpl), address(proposal));
    }

    /**
     * Once proposal is accepted, anyone can activate that.
     */
    function upgrade(IProposal _proposal)
        external
        noReentry
    {
        IGovernedContract new_impl = upgrade_proposals[address(_proposal)];
        require(new_impl != impl, "Already active!"); // in case it changes in the flight
        require(address(new_impl) != address(0), "Not registered!");
        require(_proposal.isAccepted(), "Not accepted!");

        IGovernedContract old_impl = impl;

        new_impl.migrate(old_impl);
        impl = new_impl;
        old_impl.destroy(new_impl);

        // SECURITY: prevent downgrade attack
        delete upgrade_proposals[address(_proposal)];

        // Return fee ASAP
        _proposal.destroy();

        emit Upgraded(address(new_impl), address(_proposal));
    }

    /**
     * Once proposal is reject, anyone can start collect procedure.
     */
    function collectProposal(IProposal _proposal)
        external
        noReentry
    {
        IGovernedContract new_impl = upgrade_proposals[address(_proposal)];
        require(address(new_impl) != address(0), "Not registered!");
        _proposal.collect();
        delete upgrade_proposals[address(_proposal)];
    }

    /**
     * Related to above
     */
    function proxy() external returns (address) {
        return address(this);
    }

    /**
     * SECURITY: prevent on-behalf-of calls
     */
    function migrate(IGovernedContract) external {
        revert("Good try");
    }

    /**
     * SECURITY: prevent on-behalf-of calls
     */
    function destroy(IGovernedContract) external {
        revert("Good try");
    }

    /**
     * Proxy all other calls to implementation.
     */
    function ()
        external payable
        senderOrigin
    {
        // SECURITY: senderOrigin() modifier is mandatory
        IGovernedContract impl_m = impl;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, 0, calldatasize)

            let res := call(sub(gas, 10000), impl_m, callvalue, ptr, calldatasize, 0, 0)
            // NOTE: returndatasize should allow repeatable calls
            //       what should save one opcode.
            returndatacopy(ptr, 0, returndatasize)

            switch res
            case 0 {
                revert(ptr, returndatasize)
            }
            default {
                return(ptr, returndatasize)
            }
        }
    }
}
