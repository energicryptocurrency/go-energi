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
    IBlockReward,
    StorageMasternodeRegistryV1,
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
        require(migration_complete, "cannot upgrade before migration");
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
        uint currentlength = validator_list.length + inactive_count;
        require(currentlength < mn_active, "migration already complete");

        for (uint i = currentlength; i < mn_active; ++i) {
            // limit chunk of MN migrated using gas left
            if (gasleft() <= 500000) break;

            address mn = old_registry.validator_list(i);

            // skip inactive masternodes
            if (!old_registry.isActive(mn) || old_registry.canHeartbeat(mn)) {
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
                , // status.validator_index is reset when adding to the list
                , // status.invalidations not copied (not relevant to mn registry v2.2)
                status.seq_payouts,
                // status.last_vote_epoch not copied (not relevant to mn registry v2.2)
            ) = old_registry.mn_status(mn);

            status.validator_index = validator_list.length;
            validator_list.push(mn);
            mn_status[mn] = status;
        }

        if (validator_list.length >= (mn_active - inactive_count)) {
            mn_active = validator_list.length;
            migration_complete = true;
        }
    }

    /// @notice this function triggered by governance upgrade when this contract is replaced by a newer version
    /// @dev see migrateStatusPartial() - masternode status must be migrated before governance upgrade!
    /// @param _newImpl the new masternode registry that is replacing this one
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    /// @notice the reward() function from IBlockReward is called as part of the block reward loop to pay the masternode
    function reward() external payable noReentry {
        // NOTE: ensure to move of remaining from the previous times to Treasury
        //---
        uint diff = address(this).balance - msg.value;

        if (int(diff) > 0) {
            IBlockReward treasury = IBlockReward(address(treasury_proxy.impl()));
            treasury.reward.value(diff)();
        }

        //---
        // SECURITY: do processing only when reward is exactly as expected
        if (msg.value == REWARD_MASTERNODE_V1) {
            // SECURITY: this check is essential against Masternode skip attacks!
            require(last_block_number < block.number, "Call outside of governance!");
            last_block_number = block.number;

            // Safety checks
            assert(msg.value == address(this).balance);
            uint fractions = payments_per_block;

            for (uint i = fractions; i > 0; --i) {
                assert(gasleft() > GAS_RESERVE);

                // solium-disable-next-line no-empty-blocks
                while ((gasleft() > GAS_RESERVE) && !_reward()) {}
            }
        }
    }

    /// @notice For each payment in a block (payments_per_block) this function is called to pay the next eligible masternode
    function _reward() internal returns(bool) {
        // skip when there's no masternodes
        if (current_masternode == address(0)) {
            return true;
        }

        // get the status of the current masternode
        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, current_masternode);

        // move on to the next masternode if we are done paying
        if (current_payouts >= mn_status[current_masternode].seq_payouts) {
            current_masternode = mninfo.next;
            current_payouts = 0;
            mninfo = _mnInfo(v1storage, current_masternode);
        }

        bool success = false;

        // pay valid masternodes
        ValidationStatus validation = _checkStatus(mn_status[current_masternode], mninfo);
        if (validation == ValidationStatus.MNActive) {
            uint reward_payment = REWARD_MASTERNODE_V1 / payments_per_block;
            // solium-disable-next-line security/no-send
            success = mninfo.owner.send(reward_payment);
            current_payouts++;
        // denounce invalid masternodes if they have a collateral issue or have been around too long
        // solium-disable-next-line security/no-block-members
        } else if ((validation == ValidationStatus.MNCollaterIssue) || ((block.timestamp - mn_status[current_masternode].inactive_since) > cleanup_period)) {
            _denounce(current_masternode, mninfo.owner);
        // deactivate invalid masternodes
        } else if (mn_status[current_masternode].seq_payouts > 0) {
            mn_status[current_masternode].seq_payouts = 0;
            // solium-disable-next-line security/no-block-members
            mn_status[current_masternode].inactive_since = block.timestamp;
            _deactive_common(current_masternode, mninfo.collateral);
            current_masternode = mninfo.next;
            current_payouts = 0;
            emit Deactivated(current_masternode);
        }

        return success;
    }

    /// @notice fallback function not allowed
    function () external payable {
        revert("Not supported");
    }
}
