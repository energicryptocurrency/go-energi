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

import {
    GlobalConstants,
    IGovernedContract
} from "./common.sol";
import { GlobalConstants, IGovernedContract } from "./common.sol";

import {
    IProposal,
    ISporkRegistry
} from "./SporkRegistryV1.sol";

/**
 * This contract has no chance of being updated. It must be stupid simple.
 *
 * If another upgrade logic is required in the future - it can be done as proxy stage II.
 */
contract GovernedProxy is
    GlobalConstants
{
    IGovernedContract public current_impl;
    mapping(address => IGovernedContract) public upgrade_proposals;

    constructor(IGovernedContract impl) public {
        current_impl = impl;
    }

    /**
     * Pre-create a new contract first.
     * Then propose upgrade based on that.
     */
    function proposeUpgrade(IGovernedContract new_impl, uint period) external payable
        returns(IProposal proposal)
    {
        require(new_impl != current_impl, "Already active!");
        return ISporkRegistry(SPORK_REGISTRY).createUpgradeProposal.value(msg.value)(new_impl, period);
    }

    /**
     * Once proposal is accepted, anyone can activate that.
     */
    function upgrade(IProposal proposal) external {
        IGovernedContract new_impl = upgrade_proposals[address(proposal)];
        require(new_impl != current_impl, "Already active!"); // in case it changes in the flight
        require(address(new_impl) != address(0), "Not registered!");
        require(proposal.isAccepted(), "Not accepted!");

        IGovernedContract old_impl = current_impl;

        new_impl.migrate(old_impl);
        current_impl = new_impl;
        old_impl.destroy(new_impl);
    }

    /**
     * Proxy all other calls to implementation.
     */
    function () external payable {
        // Internal calls are expected to use current_impl directly.
        // That's due to use of call() instead of delegatecall() on purpose.
        // solium-disable-next-line security/no-tx-origin
        require(tx.origin == msg.sender, "Only direct calls are allowed!");

        IGovernedContract impl = current_impl;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, 0, calldatasize)

            let res := call(sub(gas, 10000), impl, callvalue, ptr, calldatasize, 0, 0)
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
