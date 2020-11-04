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
import { StorageMasternodeRegistryV1 } from "./MasternodeRegistryV1.sol";

/// @notice MasternodeRegistryV2_2 is a mostly trivial update to MasternodeRegistryV2
/// @dev MasternodeRegistryV2_2 disables the proof of service invalidations due to a chain-split vulnerability
contract MasternodeRegistryV2_2 is
    MasternodeRegistryV2
{
    bool public migration_complete;         // true when there are no more masternodes to migrate
    uint public migration_inactive_count;   // count of inactive masternodes skipped during migration
    uint public migration_progress;         // current validator_index of migration
    uint public migration_mn_active;        // number of active masternodes to migrate

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
        migration_inactive_count = 0;
        migration_progress = 0;
        migration_mn_active = 0;
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
    function _migrate(IGovernedContract /*_oldImpl*/) internal {
        require(migration_complete);
        last_block_number = block.number;
    }

    /// @notice migrate masternode statuses from the current masternode registry
    /// @dev We migrate the available masternodes till gas left is less than or equal to 10000,
    /// @dev so this function will use the gas limit to determine how many masternodes
    /// @dev that will be migrated at a time.
    function migrateStatusPartial() external noReentry {
        require(!migration_complete, "migration already done");

        // address(uint160()) cast converts from non-payable address to allow cast to IGovernedProxy()
        IGovernedContract current_mnreg_impl = IGovernedProxy(address(uint160(proxy))).impl();
        require(address(current_mnreg_impl) != address(this), "cannot migrate from self");

        MasternodeRegistryV2 old_registry = MasternodeRegistryV2(address(current_mnreg_impl));
        StorageMasternodeRegistryV1 old_storage = StorageMasternodeRegistryV1(address(old_registry.v1storage()));
        if (migration_progress == 0) {
            migration_mn_active = old_registry.mn_active();
        }
        require(migration_progress < migration_mn_active, "migration already complete");

        for (uint i = migration_progress; i < migration_mn_active; ++i) {
            // limit chunk of MN migrated using gas left
            if (gasleft() <= 2500000) break;

            address mn = old_registry.validator_list(i);

            // skip inactive masternodes
            if (!old_registry.isActive(mn) || old_registry.canHeartbeat(mn)) {
                migration_inactive_count++;
                continue;
            } else if (current_masternode == address(0)) {
                // set current_masternode to the first active masternode
                current_masternode = mn;
            }

            // migrate validator_list and mn_status variables
            Status memory status;
            (
                status.sw_features,
                status.next_heartbeat,
                status.inactive_since,
                , // status.validator_index is reset when adding to the list
                , // status.invalidations not copied (not relevant to mn registry v2.2)
                status.seq_payouts,
                // status.last_vote_epoch not copied (not relevant to mn registry v2.2)
            ) = old_registry.mn_status(mn);
            status.validator_index = validator_list.length;
            validator_list.push(mn);
            mn_status[mn] = status;

            // migrate the mninfo
            StorageMasternodeRegistryV1.Info memory mninfo;
            (
                mninfo.announced_block,
                mninfo.collateral,
                mninfo.enode_0,
                mninfo.enode_1,
                mninfo.owner,
                mninfo.prev,
                mninfo.next,
                mninfo.ipv4address
            ) = old_storage.masternodes(mn);
            bytes32[2] memory enode;
            enode[0] = mninfo.enode_0;
            enode[1] = mninfo.enode_1;
            v1storage.setMasternode(
                mn,
                mninfo.owner,
                mninfo.ipv4address,
                enode,
                mninfo.collateral,
                mninfo.announced_block,
                address(0), // previous
                address(0) // next
            );

            // update masternode statistics
            mn_active++;
            mn_announced++;
            mn_ever_collateral += mninfo.collateral;
            mn_active_collateral += mninfo.collateral;
            mn_announced_collateral += mninfo.collateral;

            // connect the previous masternode to the current masternode
            if (validator_list.length > 0) {
                address mn_prev = validator_list[i - 1];
                // prevInfo.next = mn
                v1storage.setMasternodePos(mn_prev, false, address(0), true, mn);
                // mninfo.prev = mn_prev
                v1storage.setMasternodePos(mn, true, mn_prev, false, address(0));
            }
        }

        // update migration progress
        migration_progress = validator_list.length + migration_inactive_count;

        // check if the migration is complete
        if (validator_list.length >= (migration_mn_active - migration_inactive_count)) {
            // TODO: connect the list circularly
            address mn_first = validator_list[0];
            address mn_last = validator_list[validator_list.length - 1];
            // mn_first.prev = mn_last
            v1storage.setMasternodePos(mn_first, true, mn_last, false, address(0));
            // mn_last.next = mn_first
            v1storage.setMasternodePos(mn_last, false, address(0), true, mn_first);
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
