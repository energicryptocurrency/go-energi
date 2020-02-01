// Copyright 2020 The Energi Core Authors
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
pragma solidity 0.5.11;
//pragma experimental SMTChecker;

import { GlobalConstantsV2 } from "./constantsV2.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { MasternodeTokenV1 } from "./MasternodeTokenV1.sol";

/**
 * MN-1: Genesis hardcoded version of MasternodeToken.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract MasternodeTokenV2 is
    MasternodeTokenV1,
    GlobalConstantsV2
{
    constructor(address _proxy, IGovernedProxy _registry_proxy)
        public
        MasternodeTokenV1(_proxy, _registry_proxy)
    {}

    // ERC20
    //---------------------------------
    function symbol() external view returns (string memory) {
        return "MNRG";
    }

    function decimals() external view returns (uint8) {
        return 18;
    }

    // IGovernedContract
    //---------------------------------
    function _migrate(IGovernedContract _oldImpl) internal {
        v1storage.kill();
        v1storage = MasternodeTokenV1(address(_oldImpl)).v1storage();
    }

    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    //---------------------------------
    function _validateBalance(uint256 _amount) internal pure {
        // NOTE: "Too small" check makes no sense as it would be just zero.

        if (_amount > MN_COLLATERAL_MAX) {
            revert("Too much");
        }

        if ((_amount % MN_COLLATERAL_V2_MIN) != 0) {
            revert("Not a multiple");
        }
    }

    // Safety
    //---
    function () external payable {
        revert("Not supported");
    }
}
