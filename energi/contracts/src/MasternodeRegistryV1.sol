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
pragma solidity 0.5.16;
//pragma experimental SMTChecker;

import { GlobalConstants } from "./constants.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IBlockReward } from "./IBlockReward.sol";
import { IMasternodeRegistry } from "./IMasternodeRegistry.sol";
import { IMasternodeToken } from "./IMasternodeToken.sol";
import { ITreasury } from "./ITreasury.sol";
import { NonReentrant } from "./NonReentrant.sol";
import { StorageBase }  from "./StorageBase.sol";


/**
 * Permanent storage of Masternode Registry V1 data.
 */
contract StorageMasternodeRegistryV1 is
    StorageBase
{
    struct Info {
        uint announced_block;
        uint collateral;
        bytes32 enode_0;
        bytes32 enode_1;
        address payable owner;
        address prev;
        address next;
        uint32 ipv4address;
    }

    mapping(address => Info) public masternodes;
    mapping(address => address) public owner_masternodes;

    // NOTE: ABIEncoderV2 is not acceptable at the moment of development!

    /**
     * For initial setup.
     */
    function setMasternode(
        address _masternode,
        address payable _owner,
        uint32 _ipv4address,
        bytes32[2] calldata _enode,
        uint _collateral,
        uint _announced_block,
        address _prev,
        address _next
    )
        external
        requireOwner
    {
        Info storage item = masternodes[_masternode];
        address old_owner = item.owner;

        if (old_owner != _owner) {
            assert(old_owner == address(0));
            owner_masternodes[_owner] = _masternode;
        }

        item.owner = _owner;
        item.ipv4address = _ipv4address;
        item.enode_0 = _enode[0];
        item.enode_1 = _enode[1];
        item.collateral = _collateral;
        item.announced_block = _announced_block;
        item.prev = _prev;
        item.next = _next;
    }

    /**
     * NOTE: Extra booleans are just more failsafe than bitfield or other approaches.
     *       Conditional assignment is required to save gas.
     */
    function setMasternodePos(
        address _masternode,
        bool _set_prev, address _prev,
        bool _set_next, address _next
    )
        external
        requireOwner
    {
        Info storage item = masternodes[_masternode];

        if (_set_prev) item.prev = _prev;
        if (_set_next) item.next = _next;
    }

    function deleteMasternode(address _masternode)
        external
        requireOwner
    {
        delete owner_masternodes[masternodes[_masternode].owner];
        delete masternodes[_masternode];
    }
}


/**
 * MN-2: Genesis hardcoded version of MasternodeRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract MasternodeRegistryV1 is
    GlobalConstants,
    GovernedContract,
    IBlockReward,
    IMasternodeRegistry,
    NonReentrant
{
    // NOTE: maybe this is too much...
    event Heartbeat(
        address indexed masternode
    );

    enum Config {
        RequireValidation,
        ValidationPeriod,
        CleanupPeriod,
        InitialEverCollateral
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
    uint public validation_period;
    uint public cleanup_period;
    //---------------------------------

    // Not for migration
    struct Status {
        uint256 sw_features;
        uint last_heartbeat;
        uint inactive_since;
        uint validator_index;
        uint invalidation_since;
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
    uint last_block_number;
    //---------------------------------

    constructor(
        address _proxy,
        IGovernedProxy _token_proxy,
        IGovernedProxy _treasury_proxy,
        uint[4] memory _config
    )
        public
        GovernedContract(_proxy)
    {
        v1storage = new StorageMasternodeRegistryV1();
        token_proxy = _token_proxy;
        treasury_proxy = _treasury_proxy;

        require_validation = _config[uint(Config.RequireValidation)];
        validation_period = _config[uint(Config.ValidationPeriod)];
        cleanup_period = _config[uint(Config.CleanupPeriod)];

        uint initial_ever_collateral = _config[uint(Config.InitialEverCollateral)];
        mn_ever_collateral = initial_ever_collateral;
        require(initial_ever_collateral >= MN_COLLATERAL_MIN, "Initial collateral");
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
        StorageMasternodeRegistryV1 mn_storage = v1storage;

        // Check collateral
        //---
        uint balance = _announce_checkbalance(owner);

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
        mnstatus.last_heartbeat = block.timestamp;
        mnstatus.seq_payouts = balance / MN_COLLATERAL_MIN;
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
        mnstatus.invalidation_since = block.number;
        mnstatus.validator_index = validator_list.length;
        validator_list.push(masternode);

        // Event
        //---
        emit Announced(masternode, owner, ipv4address, enode, balance);
    }

    function _announce_checkbalance(address owner) internal view returns(uint balance) {
        (balance,) = IMasternodeToken(address(token_proxy.impl())).balanceInfo(owner);
        require(balance >= MN_COLLATERAL_MIN, "Invalid collateral");

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

        uint hearbeat_delay = block.timestamp - s.last_heartbeat;
        require(hearbeat_delay > MN_HEARTBEAT_INTERVAL_MIN, "Too early");

        s.last_heartbeat = block.timestamp;
        s.sw_features = sw_features;

        emit Heartbeat(masternode);
    }

    function invalidate(address masternode)
        external
        noReentry
    {
        address caller = _callerAddress();
        require(caller != masternode, "Invalidation for self");

        uint vote_epoch = block.number / validation_period;

        //---
        Status storage cs = mn_status[caller];
        require(_isActive(caller, cs), "Not active caller");
        require(cs.last_vote_epoch < vote_epoch, "Already invalidated");
        require(validationTarget(caller) == masternode, "Invalid target");

        //---
        Status storage s = mn_status[masternode];

        require(_isActive(masternode, s), "Not active target");

        //---
        cs.last_vote_epoch = vote_epoch;
        s.invalidations++;

        emit Invalidation(masternode, caller);
    }

    function validationTarget(address masternode) public view returns(address target) {
        uint total = validator_list.length;

        uint vperiod = validation_period;
        uint offset = (block.number / vperiod % (total - 1)) + 1;

        uint target_index = mn_status[masternode].validator_index;
        target_index = (target_index + offset) % total;

        return validator_list[target_index];
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
        if (mnstatus.seq_payouts == 0) {
            return ValidationStatus.MNNotActive;
        }

        if ((block.timestamp - mnstatus.last_heartbeat) >= MN_HEARTBEAT_INTERVAL_MAX) {
            return ValidationStatus.MNHeartbeat;
        }

        (uint balance, uint last_block) = IMasternodeToken(address(token_proxy.impl())).balanceInfo(mninfo.owner);

        if (balance != mninfo.collateral) {
            return ValidationStatus.MNCollaterIssue;
        }

        if (last_block > mninfo.announced_block) {
            return ValidationStatus.MNCollaterIssue;
        }

        return ValidationStatus.MNActive;
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

        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, masternode);
        ValidationStatus check = _checkStatus(mn_status[masternode], mninfo);

        if (check == ValidationStatus.MNCollaterIssue) {
            // Only if collateral issue!
            _denounce(masternode, owner);
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

    // IGovernedContract
    //---------------------------------
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // IBlockReward
    //---------------------------------
    function reward()
        external payable
        noReentry
    {
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
            last_block_number = last_block_number;

            assert(gasleft() > GAS_RESERVE);
            assert(msg.value == address(this).balance);

            // solium-disable-next-line no-empty-blocks
            while ((gasleft() > GAS_RESERVE) && !_reward()) {}
        }
    }

    function _reward() internal returns(bool) {
        //---
        address masternode = current_masternode;
        uint payouts = current_payouts;

        if (masternode == address(0)) {
            return true;
        }

        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, masternode);

        Status storage mnstatus = mn_status[masternode];
        uint invalidations = mnstatus.invalidations;
        uint invalidation_since = mnstatus.invalidation_since;
        ++payouts;

        if (payouts < mnstatus.seq_payouts) {
            current_payouts = payouts;
        } else {
            mnstatus.invalidations = 0;
            mnstatus.invalidation_since = block.number;
            current_masternode = mninfo.next;
            current_payouts = 0;
        }

        // Reward logic
        //---
        ValidationStatus status = _checkStatus(mnstatus, mninfo);

        if (status == ValidationStatus.MNActive) {
            // solium-disable security/no-send
            if (!_canReward(invalidations, invalidation_since) ||
                mninfo.owner.send(msg.value)
            ) {
                return true;
            }
            // solium-enable security/no-send
        }

        // When not valid
        //---
        if (status == ValidationStatus.MNCollaterIssue) {
            // Immediate
            _denounce(masternode, mninfo.owner);
        } else if (mnstatus.seq_payouts > 0) {
            // Mark as inactive for later auto-cleanup
            mnstatus.seq_payouts = 0;
            mnstatus.inactive_since = block.timestamp;
            _deactive_common(masternode, mninfo.collateral);
            current_masternode = mninfo.next;
            current_payouts = 0;

            emit Deactivated(masternode);
        } else if ((block.timestamp - mnstatus.inactive_since) > cleanup_period) {
            // Auto-cleanup
            _denounce(masternode, mninfo.owner);
        }

        return false;
    }

    function _canReward(uint invalidations, uint invalidation_since) internal view returns(bool) {
        if (mn_active < require_validation) {
            return true;
        }

        uint threshold = block.number - invalidation_since;
        threshold = (threshold / validation_period) + 1;
        threshold /= 2;

        return (invalidations < threshold);
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
