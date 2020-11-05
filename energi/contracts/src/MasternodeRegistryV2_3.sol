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
import { StorageMasternodeRegistryV1 } from "./MasternodeRegistryV1.sol";

/**
 * MN-2: Genesis hardcoded version of MasternodeRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract MasternodeRegistryV2_3 is
    GlobalConstants,
    GlobalConstantsV2,
    GovernedContract,
    IBlockReward,
    IMasternodeRegistryV2,
    NonReentrant
{
    // MN-4 approximation logic:
    // - the target is 1000 hearbeats per hour
    // - {MN count} / 1000/3600 ~ {MN count} * 4
    uint constant internal TARGET_HEARTBEATS_COEF = 4;

    enum Config {
        RequireValidation,
        ValidationPeriods,
        CleanupPeriod,
        InitialEverCollateral,
        PaymentsPerBlock
    }

    // Data for migration
    //---------------------------------
    StorageMasternodeRegistryV1 public v1storage;

    IGovernedProxy public token_proxy;
    IGovernedProxy public treasury_proxy;

    uint public mn_announced;

    address public current_masternode;
    uint public current_payouts;
    uint public require_validation;
    uint public validation_periods;
    uint public cleanup_period;
    uint public payments_per_block;
    //---------------------------------

    // Not for migration
    struct Status {
        uint256 sw_features;
        uint next_heartbeat;
        uint inactive_since;
        uint validator_index;
        uint invalidations;
        uint seq_payouts;
        uint last_vote_epoch;
    }

    uint public mn_ever_collateral;
    uint public mn_active_collateral;
    uint public mn_announced_collateral;

    uint public mn_active;
    mapping(address => Status) public mn_status;
    address[] public validator_list;
    uint public last_block_number;

    uint public curr_validation_ends;
    uint public curr_validation_offset;

    bool migration_complete;
    uint inactive_count;
    //---------------------------------

    constructor(
        address _proxy,
        IGovernedProxy _token_proxy,
        IGovernedProxy _treasury_proxy,
        uint[5] memory _config
    )
        public
        GovernedContract(_proxy)
    {
        v1storage = new StorageMasternodeRegistryV1();
        token_proxy = _token_proxy;
        treasury_proxy = _treasury_proxy;

        require_validation = _config[uint(Config.RequireValidation)];
        validation_periods = _config[uint(Config.ValidationPeriods)];
        cleanup_period = _config[uint(Config.CleanupPeriod)];
        payments_per_block = _config[uint(Config.PaymentsPerBlock)];

        require(validation_periods <= require_validation, "Validations > Require");

        uint initial_ever_collateral = _config[uint(Config.InitialEverCollateral)];
        mn_ever_collateral = initial_ever_collateral;
        require(initial_ever_collateral >= MN_COLLATERAL_V2_MIN, "Initial collateral");
    }

    // IMasternodeRegistry
    //---------------------------------

    enum ValidationStatus {
        MNActive,
        MNCollaterIssue,
        MNNotActive,
        MNHeartbeat
    }

    uint constant internal GAS_RESERVE = 100000;

    // solium-disable security/no-block-members

    // Announcement
    //=================================

    function announce(address masternode, uint32 ipv4address, bytes32[2] calldata enode)
        external
        noReentry
    {
        address owner = _callerAddress();

        // Check collateral
        //---
        uint balance = _announce_checkbalance(owner);

        _announce(masternode, owner, balance, ipv4address, enode);
    }

    function _announce(
        address masternode,
        address owner,
        uint balance,
        uint32 ipv4address,
        bytes32[2] memory enode
    ) internal {
        StorageMasternodeRegistryV1 mn_storage = v1storage;

        // Cleanup & checks
        //---
        _announce_clear_old(mn_storage, owner);
        _announce_check_free(mn_storage, masternode);
        _announce_check_ipv4(ipv4address);

        // Insert into list
        //---
        (address next, address prev) = _announce_insert(mn_storage, masternode);

        // Save
        //---
        mn_storage.setMasternode(
            masternode,
            address(uint160(owner)),
            ipv4address,
            enode,
            balance,
            block.number,
            prev,
            next
        );

        Status storage mnstatus = mn_status[masternode];
        mnstatus.next_heartbeat = block.timestamp;
        mnstatus.seq_payouts = balance / MN_COLLATERAL_V2_MIN;
        ++mn_active;
        ++mn_announced;

        mn_active_collateral += balance;
        uint announced_collateral = mn_announced_collateral;
        announced_collateral += balance;
        mn_announced_collateral = announced_collateral;

        if (announced_collateral > mn_ever_collateral) {
            mn_ever_collateral = announced_collateral;
        }

        // Validator logic is de-coupled for easier changes
        //---
        mnstatus.validator_index = validator_list.length;
        validator_list.push(masternode);

        // Event
        //---
        emit Announced(masternode, owner, ipv4address, enode, balance);
    }

    function _announce_checkbalance(address owner) internal view returns(uint balance) {
        (balance,) = _getCollateralInfo(owner);
        require(balance >= MN_COLLATERAL_V2_MIN, "Invalid collateral");
    }

    function _announce_clear_old(StorageMasternodeRegistryV1 mn_storage, address owner) internal {
        address old_masternode = mn_storage.owner_masternodes(owner);

        // Regardless if it is re-announcement
        if (old_masternode != address(0)) {
            _denounce(old_masternode, owner);
        }
    }

    function _announce_check_free(StorageMasternodeRegistryV1 mn_storage, address masternode)
        internal view
    {
        // SECURITY: there is an option of seizing a foreign MN address at cost of collateral.
        //           The mitigation is regeneration of such address by a victim.
        //           MN should refuse to operate in such condition.
        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(mn_storage, masternode);
        require(mninfo.owner == address(0), "Invalid owner");
    }

    function _announce_check_ipv4(uint32 ipv4address) internal pure {
        uint a = ipv4address & 0xFF000000;
        uint b = ipv4address & 0x00FF0000;
        uint c = ipv4address & 0x0000FF00;
        // solium-disable operator-whitespace
        require(
            // 127.0.0.0/8
            (a != (127 << 24)) &&
            // 10.0.0.0/8
            (a != (10 << 24)) &&
            // 192.168.0.0/16
            !((a == (192 << 24)) && (b == (168 << 16))) &&
            // 172.16.0.0/12
            !((a == (172 << 24)) && ((b & 0x00F00000) == (16 << 16))) &&
            // 0.0.0.0/8
            (a != 0) &&
            // 100.64.0.0/10
            !((a == (100 << 24)) && ((b & 0x00C00000) == (64 << 16))) &&
            // 169.254.0.0/16
            !((a == (169 << 24)) && (b == (254 << 16))) &&
            // 198.18.0.0/15
            !((a == (198 << 24)) && ((b & 0x00FE0000) == (18 << 16))) &&
            // 198.51.100.0/24
            !((a == (198 << 24)) && (b == (51 << 16)) && (c == (100 << 8))) &&
            // 203.0.113.0/24
            !((a == (203 << 24)) && (b == (0 << 16)) && (c == (113 << 8))) &&
            // 224.0.0.0/4
            ((a & 0xF0000000) != (224 << 24)) &&
            // 240.0.0.0/4
            ((a & 0xF0000000) != (240 << 24)) &&
            // 255.255.255.255/32
            (ipv4address != 0xFFFFFFFF),
            "Wrong IP");
        // solium-enable operator-whitespace
    }

    function _announce_insert(StorageMasternodeRegistryV1 mn_storage, address masternode)
        internal
        returns(address next, address prev)
    {
        // NOTE: always insert as the last - before the current one
        next = current_masternode;

        if (next != address(0)) {
            StorageMasternodeRegistryV1.Info memory nextinfo = _mnInfo(mn_storage, next);

            prev = nextinfo.prev;

            // Not effective for the second one, but reliable
            mn_storage.setMasternodePos(
                nextinfo.prev,
                false, address(0),
                true, masternode
            );
            mn_storage.setMasternodePos(
                next,
                true, masternode,
                false, address(0)
            );
        } else {
            // The first one
            current_masternode = masternode;
            current_payouts = 0;
            prev = masternode;
            next = masternode;
        }
    }

    //=================================

    function denounce(address masternode)
        external
        noReentry
    {
        _denounce(masternode, _callerAddress());
    }

    function _denounce(address masternode, address owner) internal {
        // Check masternode ownership, if already registered.
        //---
        StorageMasternodeRegistryV1 mn_storage = v1storage;
        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(mn_storage, masternode);

        if (mninfo.owner == address(0)) {
            return;
        }

        require((owner == mninfo.owner), "Invalid owner");

        // Remove from list
        //---
        if (mninfo.next == masternode) {
            // the last one
            current_masternode = address(0);
        } else {
            if (current_masternode == masternode) {
                current_masternode = mninfo.next;
                current_payouts = 0;
            }

            mn_storage.setMasternodePos(
                mninfo.prev,
                false, address(0),
                true, mninfo.next
            );
            mn_storage.setMasternodePos(
                mninfo.next,
                true, mninfo.prev,
                false, address(0)
            );
        }

        // Delete
        //---

        mn_announced_collateral -= mninfo.collateral;

        if (mn_status[masternode].seq_payouts > 0) {
            _deactive_common(masternode, mninfo.collateral);
        }

        delete mn_status[masternode];

        mn_storage.deleteMasternode(masternode);
        --mn_announced;

        //---
        emit Denounced(masternode, mninfo.owner);
    }

    function _deactive_common(address masternode, uint collateral) internal {
        // Remove from validators
        address last_validator = validator_list[validator_list.length - 1];
        uint validator_index = mn_status[masternode].validator_index;

        mn_status[last_validator].validator_index = validator_index;
        validator_list[validator_index] = last_validator;
        validator_list.pop();

        //--
        --mn_active;
        mn_active_collateral -= collateral;
    }

    function heartbeat(uint block_number, bytes32 block_hash, uint sw_features)
        external
        noReentry
    {
        require((block.number - block_number - 1) <= MN_HEARTBEAT_PAST_BLOCKS, "Too old block");
        require(blockhash(block_number) == block_hash, "Block mismatch");

        address payable masternode = _callerAddress();

        Status storage s = mn_status[masternode];

        require(_isActive(masternode, s), "Not active");

        require(s.next_heartbeat <= block.timestamp, "Too early");

        s.next_heartbeat = block.timestamp + _newHeartbeatInterval();
        s.sw_features = sw_features;
    }

    function _newHeartbeatInterval() internal view returns(uint delay) {
        delay = mn_active * TARGET_HEARTBEATS_COEF;

        if (delay < MN_HEARTBEAT_INTERVAL_MIN) {
            delay = MN_HEARTBEAT_INTERVAL_MIN;
        }
    }

    /// @notice proof of service invalidation
    /// @dev this is disabled due to chain split vulnerability in previous versions
    /// @dev masternode address is the masternode to invalidate.
    function invalidate(address /*masternode*/) external noReentry {
        require(false, "invalidations disabled");
    }

    function validationTarget(address /*masternode*/) public view returns(address) {
        return address(0);
    }

    function isActive(address masternode) external view returns(bool) {
        return _isActive(masternode, mn_status[masternode]);
    }

    //===

    function _isActive(address masternode, Status storage mnstatus) internal view returns(bool) {
        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, masternode);
        return _checkStatus(mnstatus, mninfo) == ValidationStatus.MNActive;
    }

    function _checkStatus(
        Status storage mnstatus,
        StorageMasternodeRegistryV1.Info memory mninfo
    )
        internal view
        returns(ValidationStatus)
    {
        (uint balance, uint last_block) = _getCollateralInfo(mninfo.owner);
        return _checkStatus(mnstatus, mninfo, balance, last_block);
    }

    function _checkStatus(
        Status storage mnstatus,
        StorageMasternodeRegistryV1.Info memory mninfo,
        uint balance,
        uint last_block
    )
        internal view
        returns(ValidationStatus)
    {
        if (mnstatus.seq_payouts == 0) {
            return ValidationStatus.MNNotActive;
        }

        if (block.timestamp > (mnstatus.next_heartbeat + MN_HEARTBEAT_INTERVAL_MAX)) {
            return ValidationStatus.MNHeartbeat;
        }

        if (balance != mninfo.collateral) {
            return ValidationStatus.MNCollaterIssue;
        }

        if (last_block > mninfo.announced_block) {
            return ValidationStatus.MNCollaterIssue;
        }

        return ValidationStatus.MNActive;
    }

    function _getCollateralInfo(address owner)
        internal view
        returns(
            uint balance,
            uint last_block
        )
    {
        (balance, last_block) = IMasternodeToken(address(token_proxy.impl())).balanceInfo(owner);
    }

    //===

    function count() external view
        returns(
            uint active, uint total,
            uint active_collateral, uint total_collateral,
            uint max_of_all_times
        )
    {
        active = mn_active;
        total = mn_announced;
        active_collateral = mn_active_collateral;
        total_collateral = mn_announced_collateral;
        max_of_all_times = mn_ever_collateral;
    }

    //===
    function info(address masternode) external view
        returns(
            address owner, uint32 ipv4address, bytes32[2] memory enode,
            uint collateral, uint announced_block, uint sw_features
        )
    {
        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, masternode);
        require(mninfo.owner != address(0), "Unknown masternode");
        owner = mninfo.owner;
        ipv4address = mninfo.ipv4address;
        enode = [ mninfo.enode_0, mninfo.enode_1 ];
        collateral = mninfo.collateral;
        announced_block = mninfo.announced_block;

        sw_features = mn_status[masternode].sw_features;
    }

    function ownerInfo(address owner) external view
        returns(
            address masternode, uint32 ipv4address, bytes32[2] memory enode,
            uint collateral, uint announced_block, uint sw_features
        )
    {
        StorageMasternodeRegistryV1 mnstorage = v1storage;

        masternode = mnstorage.owner_masternodes(owner);
        require(masternode != address(0), "Unknown owner");

        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(mnstorage, masternode);
        masternode = masternode;
        ipv4address = mninfo.ipv4address;
        enode = [ mninfo.enode_0, mninfo.enode_1 ];
        collateral = mninfo.collateral;
        announced_block = mninfo.announced_block;

        sw_features = mn_status[masternode].sw_features;
    }

    function _mnInfo(
        StorageMasternodeRegistryV1 v1info,
        address masternode
    )
        internal view
        returns (StorageMasternodeRegistryV1.Info memory mninfo)
    {
        // NOTE: no ABIv2 encoding is enabled
        (
            mninfo.announced_block,
            mninfo.collateral,
            mninfo.enode_0,
            mninfo.enode_1,
            mninfo.owner,
            mninfo.prev,
            mninfo.next,
            mninfo.ipv4address
        ) = v1info.masternodes(masternode);
    }

    //===

    function onCollateralUpdate(address owner)
        external
        noReentry
    {
        // SECURITY: this is a callback for MasternodeToken, but
        //           it must be safe to be called by ANYONE.

        StorageMasternodeRegistryV1 mn_storage = v1storage;
        address masternode = mn_storage.owner_masternodes(owner);

        if (masternode == address(0)) {
            return;
        }

        (uint balance, uint last_block) = _getCollateralInfo(owner);

        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, masternode);
        ValidationStatus check = _checkStatus(mn_status[masternode], mninfo, balance, last_block);

        if (check == ValidationStatus.MNCollaterIssue) {
            // Re-announce, if there is collateral left.
            if (balance >= MN_COLLATERAL_V2_MIN) {
                uint32 ipv4address = mninfo.ipv4address;
                bytes32[2] memory enode = [mninfo.enode_0, mninfo.enode_1];

                _announce(masternode, owner, balance, ipv4address, enode);
            } else {
                _denounce(masternode, owner);
            }
        }
    }

    function enumerate()
        external view
        returns(address[] memory masternodes)
    {
        // NOTE: it should be OK for 0
        masternodes = new address[](mn_announced);
        address curr_mn = current_masternode;

        if (curr_mn == address(0)) {
            return masternodes;
        }

        address next = curr_mn;
        StorageMasternodeRegistryV1.Info memory mninfo;
        uint i = 0;

        do {
            masternodes[i] = next;
            mninfo = _mnInfo(v1storage, next);
            next = mninfo.next;
            ++i;
        } while (next != curr_mn);
    }

    function enumerateActive()
        external view
        returns(address[] memory masternodes)
    {
        // NOTE: this API is targeted at fast consensus execution
        masternodes = new address[](mn_active);

        for (uint i = 0; i < masternodes.length; ++i) {
            masternodes[i] = validator_list[i];
        }
    }

    // IMasternodeRegistryV2
    //---------------------------------
    function collateralLimits() external pure returns (uint min, uint max) {
        min = MN_COLLATERAL_V2_MIN;
        max = MN_COLLATERAL_MAX;
    }

    function canHeartbeat(address masternode) external view returns(bool can_heartbeat) {
        Status storage s = mn_status[masternode];

        return _isActive(masternode, s) && (s.next_heartbeat <= block.timestamp);
    }

    function canInvalidate(address /*masternode*/) external view returns(bool) {
        return false;
    }

    /// @notice this migration function triggered by governance upgrade when replacing another version
    /// @dev see migrateStatusPartial() - masternode status must be migrated before governance upgrade!
    /// @param _oldImpl the previous masternode registry being migrated
    function _migrate(IGovernedContract _oldImpl) internal {
        require(migration_complete, "cannot upgrade before migration");
        // Dispose
        v1storage.kill();

        MasternodeRegistryV2_3 oldinstance = MasternodeRegistryV2_3(address(_oldImpl));
        v1storage = oldinstance.v1storage();

        // Migration data
        current_masternode = oldinstance.current_masternode();
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

        MasternodeRegistryV2_3 old_registry = MasternodeRegistryV2_3(address(current_mnreg_impl));
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
        Status storage mns = mn_status[current_masternode];

        // move on to the next masternode if we are done paying
        if (current_payouts >= mns.seq_payouts) {
            current_masternode = mninfo.next;
            current_payouts = 0;
            mninfo = _mnInfo(v1storage, current_masternode);
        }

        bool success = false;

        // pay valid masternodes
        ValidationStatus validation = _checkStatus(mns, mninfo);
        if (validation == ValidationStatus.MNActive) {
            uint reward_payment = REWARD_MASTERNODE_V1 / payments_per_block;
            success = mninfo.owner.send(reward_payment);
            current_payouts++;
        // denounce invalid masternodes if they have a collateral issue or have been around too long
        } else if ((validation == ValidationStatus.MNCollaterIssue) || ((block.timestamp - mns.inactive_since) > cleanup_period)) {
            _denounce(current_masternode, mninfo.owner);
        // deactivate invalid masternodes
        } else if (mns.seq_payouts > 0) {
            mns.seq_payouts = 0;
            mns.inactive_since = block.timestamp;
            _deactive_common(current_masternode, mninfo.collateral);
            current_masternode = mninfo.next;
            current_payouts = 0;
            emit Deactivated(current_masternode);
        } else {
            current_masternode = mninfo.next;
            current_payouts = 0;
        }

        return success;
    }

    //===

    function getReward(uint _blockNumber)
        external view
        returns(uint amount)
    {
        ITreasury treasury = ITreasury(address(treasury_proxy.impl()));

        if ((_blockNumber > 0) && !treasury.isSuperblock(_blockNumber)) {
            amount = REWARD_MASTERNODE_V1;
        }
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
