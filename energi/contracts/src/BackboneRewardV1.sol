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
pragma solidity 0.5.10;
//pragma experimental SMTChecker;

import { GlobalConstants } from "./constants.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IBlockReward } from "./IBlockReward.sol";
import { NonReentrant } from "./NonReentrant.sol";

/**
 * Genesis hardcoded version of BackboneReward
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract BackboneRewardV1 is
    GlobalConstants,
    GovernedContract,
    IBlockReward,
    NonReentrant
{
    address payable public backbone_address;

    // IGovernedContract
    //---------------------------------
    constructor(address _proxy, address payable _backbone_address) public GovernedContract(_proxy) {
        backbone_address = _backbone_address;
    }

    // IBlockReward
    //---------------------------------
    function reward()
        external payable
        noReentry
    {
        if (msg.value > 0) {
            backbone_address.transfer(msg.value);
        }
    }

    function getReward(uint _blockNumber)
        external view
        returns(uint amount)
    {
        if ((_blockNumber > 0) && (backbone_address != address(0))) {
            amount = REWARD_BACKBONE_V1;
        }
    }

    // Safety
    //---
    function () external payable {
        revert("Not supported");
    }
}

