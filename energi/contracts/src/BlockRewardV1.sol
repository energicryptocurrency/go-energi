// Copyright 2021 The Energi Core Authors
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
pragma solidity 0.5.16;
//pragma experimental SMTChecker;

import { GovernedContract } from "./GovernedContract.sol";
import { IBlockReward } from "./IBlockReward.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { NonReentrant } from "./NonReentrant.sol";

/**
 * Genesis hardcoded version of SporkReward
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract BlockRewardV1 is
    GovernedContract,
    IBlockReward,
    NonReentrant
{
    uint constant internal GET_GAS = 10000;

    IGovernedProxy[] public reward_proxies;

    // IGovernedContract
    //---------------------------------
    constructor(address _proxy, IGovernedProxy[] memory _reward_proxies)
        public
        GovernedContract(_proxy)
    {
        for (uint i = 0; i < _reward_proxies.length; ++i) {
            reward_proxies.push(_reward_proxies[i]);
        }
    }

    // IBlockReward
    //---------------------------------
    function reward()
        external payable
        noReentry
    {
        uint len = reward_proxies.length;
        uint gas_per_reward = (gasleft() / len) - GET_GAS;

        for (uint i = 0; i < len; ++i) {
            IBlockReward impl = IBlockReward(address(reward_proxies[i].impl()));

            uint amount = impl.getReward.gas(GET_GAS)(block.number);

            // solium-disable-next-line security/no-call-value
            address(impl).call.value(amount).gas(gas_per_reward)(abi.encode(impl.reward.selector));
        }
    }

    function getReward(uint _blockNumber)
        external view
        returns(uint amount)
    {
        for (uint i = reward_proxies.length; i-- > 0;) {
            IBlockReward impl = IBlockReward(address(reward_proxies[i].impl()));
            amount += impl.getReward(_blockNumber);
        }
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
