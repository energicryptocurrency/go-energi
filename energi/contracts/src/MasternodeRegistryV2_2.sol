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

pragma solidity 0.5.16;

import {
    IGovernedProxy,
    IGovernedContract,
    MasternodeRegistryV2
} from "./MasternodeRegistryV2.sol";

/// @notice MasternodeRegistryV2_2 is a mostly trivial update to MasternodeRegistryV2
/// @dev MasternodeRegistryV2_2 disables the proof of service invalidations due to a chain-split vulnerability
contract MasternodeRegistryV2_2 is
    MasternodeRegistryV2
{
    bool migration_complete;
    uint inactive_count;

    /// @notice construct a new MasternodeRegistryV2_2
    /// @param _proxy The MasternodeRegistry proxy address
    /// @param _token_proxy The Masternode Token (MNRG) proxy address
    /// @param _treasury_proxy The Treasury proxy address
    /// @param _config MasternodeRegistry configuration ( MNRequireValidation, MNValidationPeriod, MNCleanupPeriod, MNEverCollateral, MNRewardsPerBlock )
    constructor(
        address _proxy,
        IGovernedProxy _token_proxy,
        IGovernedProxy _treasury_proxy,
        uint[5] memory _config
    )
        public
        MasternodeRegistryV2(_proxy, _token_proxy, _treasury_proxy, _config)
    {
        migration_complete = false;
        inactive_count = 0;
        current_masternode = address(0);
        current_payouts = 0;
    }

    /// @notice proof of service invalidation
    /// @dev this is disabled due to chain split vulnerability in previous versions
    /// @dev masternode address is the masternode to invalidate.
    function invalidate(address /*masternode*/) external noReentry {
        require(false, "invalidations disabled");
    }

    /// @notice this migration function triggered by governance upgrade when replacing another version
    /// @dev see migrateStatusPartial() - masternode status must be migrated before governance upgrade!
    /// @param _oldImpl the previous masternode registry being migrated
    function _migrate(IGovernedContract _oldImpl) internal {
        // Dispose
        v1storage.kill();

        MasternodeRegistryV2 oldinstance = MasternodeRegistryV2(address(_oldImpl));
        v1storage = oldinstance.v1storage();

        // Migration data
        mn_announced = oldinstance.mn_announced();
        if (current_masternode == oldinstance.current_masternode()) {
            current_payouts = oldinstance.current_payouts();
        }

        // Other data
        mn_ever_collateral = oldinstance.mn_ever_collateral();
        mn_active_collateral = oldinstance.mn_active_collateral();
        mn_announced_collateral = oldinstance.mn_announced_collateral();
        last_block_number = block.number;
    }

    /// @notice migrate masternode statuses from the current masternode registry
    /// @dev We migrate the available masternodes till gas left is less than or equal to 10000,
    /// @dev so this function will use the gas limit to determine how many masternodes
    /// @dev that will be migrated at a ago.
    function migrateStatusPartial() external noReentry {
        require(!migration_complete, "migration already done");

        // address(uint160()) cast converts from non-payable address to allow cast to IGovernedProxy()
        IGovernedContract current_mnreg_impl = IGovernedProxy(address(uint160(proxy))).impl();
        require(address(current_mnreg_impl) != address(this), "cannot migrate from self");

        MasternodeRegistryV2 old_registry = MasternodeRegistryV2(address(current_mnreg_impl));
        mn_active = old_registry.mn_active();
        uint currentlength = validator_list.length;
        require(currentlength < mn_active, "migration already complete");

        for (uint i = currentlength; i < mn_active; ++i) {
            // limit chunk of MN migrated using gas left
            if (gasleft() <= 500000) break;

            address mn = old_registry.validator_list(i);

            // skip inactive masternodes
            if (!old_registry.isActive(mn)) {
                inactive_count++;
                continue;
            } else if (current_masternode == address(0)) {
                current_masternode = mn;
            }

            Status memory status;
            (
                status.sw_features,
                status.next_heartbeat,
                status.inactive_since,
                status.validator_index,
                , // status.invalidations not copied (not relevant to mn registry v2.2)
                status.seq_payouts,
                // status.last_vote_epoch not copied (not relevant to mn registry v2.2)
            ) = old_registry.mn_status(mn);

            validator_list.push(mn);
            mn_status[mn] = status;
        }

        if (validator_list.length >= (mn_active - inactive_count)) {
            mn_active = old_registry.mn_active() - inactive_count;
            migration_complete = true;
        }
    }

    /// @notice this function triggered by governance upgrade when this contract is replaced by a newer version
    /// @dev see migrateStatusPartial() - masternode status must be migrated before governance upgrade!
    /// @param _newImpl the new masternode registry that is replacing this one
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    /// @notice fallback function not allowed
    function () external payable {
        revert("Not supported");
    }
}
