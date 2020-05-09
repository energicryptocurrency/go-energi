// Copyright 2019-2020 The Energi Core Authors
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

import { GlobalConstants } from "./constants.sol";
import { GlobalConstantsV2 } from "./constantsV2.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IBlockReward } from "./IBlockReward.sol";
import { IMasternodeRegistryV2 } from "./IMasternodeRegistryV2.sol";
import { IMasternodeToken } from "./IMasternodeToken.sol";
import { ITreasury } from "./ITreasury.sol";
import { NonReentrant } from "./NonReentrant.sol";
import { StorageBase }  from "./StorageBase.sol";
import {
    MasternodeRegistryV1,
    StorageMasternodeRegistryV1
} from "./MasternodeRegistryV1.sol";
import { MasternodeRegistryV2 } from "./MasternodeRegistryV2.sol";

/**
 * MN-2: Genesis hardcoded version of MasternodeRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract MasternodeRegistryV2_1 is
    MasternodeRegistryV2
{
    constructor(
        address _proxy,
        IGovernedProxy _token_proxy,
        IGovernedProxy _treasury_proxy,
        uint[5] memory _config
    )
        public
        MasternodeRegistryV2(_proxy, _token_proxy, _treasury_proxy, _config)
    {

    }

    function invalidate(address /*masternode*/)
        external
        noReentry
    {
        require(false, "invalidations disabled");
    }

    // IGovernedContract
    //---------------------------------
    function _migrate(IGovernedContract _oldImpl) internal {
        // Dispose
        v1storage.kill();

        MasternodeRegistryV2 oldinstance = MasternodeRegistryV2(address(_oldImpl));
        v1storage = oldinstance.v1storage();

        // Migration data
        mn_announced = oldinstance.mn_announced();
        current_masternode = oldinstance.current_masternode();
        current_payouts = oldinstance.current_payouts();

        // Other data
        mn_ever_collateral = oldinstance.mn_ever_collateral();
        mn_active_collateral = oldinstance.mn_active_collateral();
        mn_announced_collateral = oldinstance.mn_announced_collateral();
        mn_active = oldinstance.mn_active();
        last_block_number = block.number;

        // Restore the mn status information.
        // NOTE: Only active masternodes (validator_list()) are considered for migration
        // NOTE: this may be a serious gas consumption problem due to open limit.
        for (uint i = 0; i < mn_active; ++i) {
            address mn = oldinstance.validator_list(i);
            Status memory status;
            (
                status.sw_features,
                status.next_heartbeat,
                status.inactive_since,
                status.validator_index,
                status.invalidations,
                status.seq_payouts,
                status.last_vote_epoch
            ) = oldinstance.mn_status(mn);

            validator_list.push(mn);
            mn_status[mn] = status;
        }

        _processValidationEpoch();
    }

    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
