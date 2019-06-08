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

import { GlobalConstants } from "./constants.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IProposal } from "./IProposal.sol";
import { ISporkRegistry } from "./ISporkRegistry.sol";
import { GenericProposalV1 } from "./GenericProposalV1.sol";

/**
 * Genesis hardcoded version of SporkRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract SporkRegistryV1 is
    GlobalConstants,
    ISporkRegistry,
    GovernedContract
{
    // IGovernedContract
    //---------------------------------
    constructor(address _proxy) public GovernedContract(_proxy) {}
    function migrate(IGovernedContract) external requireProxy {}

    // ISporkRegistry
    //---------------------------------
    function createUpgradeProposal(IGovernedContract, uint _period)
        external payable
        returns (IProposal)
    {
        require(msg.value == FEE_UPGRADE_V1, "Fee amount");
        require(_period >= PERIOD_UPGRADE_MIN, "Period min");
        require(_period <= PERIOD_UPGRADE_MAX, "Period max");

        address payable proposal = address(
            new GenericProposalV1(
                QUORUM_MAJORITY,
                _period,
                // solium-disable-next-line security/no-tx-origin
                tx.origin,
                msg.value
            )
        );

        proposal.transfer(msg.value);

        return IProposal(proposal);
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
