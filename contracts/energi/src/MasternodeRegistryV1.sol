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
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IBlockReward } from "./IBlockReward.sol";
import { IMasternodeRegistry } from "./IMasternodeRegistry.sol";
import { IMasternodeToken } from "./IMasternodeToken.sol";
import { ITreasury } from "./ITreasury.sol";
import { StorageBase }  from "./StorageBase.sol";


/**
 * Permanent storage of Masternode Registry V1 data.
 */
contract StorageMasternodeRegistryV1 is
    StorageBase
{
    struct Info {
        address payable owner;
        uint64 last_announced;
        uint32 ipv4address;
        bytes32 enode_0;
        bytes32 enode_1;
        uint collateral;

        address prev;
        address next;
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
        uint64 _last_announced,
        uint32 _ipv4address,
        bytes32[2] calldata _enode,
        uint _collateral,
        address _prev,
        address _next
    ) external {
        Info storage item = masternodes[_masternode];
        address old_owner = item.owner;

        if (old_owner != _owner) {
            assert(old_owner == address(0));
            owner_masternodes[_owner] = _masternode;
        }

        item.owner = _owner;
        item.last_announced = _last_announced;
        item.ipv4address = _ipv4address;
        item.enode_0 = _enode[0];
        item.enode_1 = _enode[1];
        item.collateral = _collateral;
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
    ) external {
        Info storage item = masternodes[_masternode];

        if (_set_prev) item.prev = _prev;
        if (_set_next) item.next = _next;
    }

    function deleteMasternode(address _masternode) external {
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
    IMasternodeRegistry
{
    // NOTE: maybe this is too much...
    event Heartbeat(
        address indexed masternode
    );

    enum Config {
        VotePerCycle,
        RequireVoting,
        VotesMax,
        CleanupPeriod
    }

    // Data for migration
    //---------------------------------
    StorageMasternodeRegistryV1 public v1storage;

    IGovernedProxy public token_proxy;
    IGovernedProxy public treasury_proxy;

    uint32 public mn_announced;
    uint64 public mn_total_ever;

    address public current_masternode;
    uint8 public current_payouts;
    uint[4] public config;
    //---------------------------------

    // Not for migration
    struct Status {
        uint256 sw_features;
        uint64 last_heartbeat;
        uint32 validations;
        uint32 votes;
        uint64 inactive_since;
        uint8 seq_payouts;
    }

    uint32 public mn_active;
    mapping(address => Status) public mn_status;
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
        config = _config;
    }

    // IMasternodeRegistry
    //---------------------------------

    enum ValidationStatus {
        MNValid,
        MNCollaterIssue,
        MNNotActive,
        MNHeartbeat
    }

    uint constant internal GAS_RESERVE = 100000;

    // solium-disable security/no-block-members

    // Announcement
    //=================================

    function announce(address masternode, uint32 ipv4address, bytes32[2] calldata enode) external {
        address owner = _callerAddress();
        StorageMasternodeRegistryV1 mn_storage = v1storage;

        // Check collateral
        //---
        uint balance = _announce_checkbalance(owner);

        // Cleanup & checks
        //---
        _announce_clear_old(mn_storage, owner);
        _announce_check_free(mn_storage, masternode);

        // Insert into list
        //---
        (address next, address prev) = _announce_insert(mn_storage, masternode);

        // Save
        //---
        mn_storage.setMasternode(
            masternode,
            address(uint160(owner)),
            uint64(block.timestamp),
            ipv4address,
            enode,
            balance,
            prev,
            next
        );

        Status storage mnstatus = mn_status[masternode];
        mnstatus.last_heartbeat = uint64(block.timestamp);
        mnstatus.seq_payouts = uint8(balance / MN_COLLATERAL_MIN);
        ++mn_active;

        uint32 announced_count = mn_announced + 1;
        mn_announced = announced_count;

        if (announced_count > mn_total_ever) {
            mn_total_ever = announced_count;
        }

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

    function denounce(address masternode) external {
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
        if (mn_status[masternode].seq_payouts > 0) {
            --mn_active;
        }

        delete mn_status[masternode];

        mn_storage.deleteMasternode(masternode);
        --mn_announced;

        //---
        emit Denounced(masternode, mninfo.owner);
    }

    function heartbeat(uint block_number, bytes32 block_hash, uint sw_features) external {
        require((block.number - block_number - 1) <= MN_HEARTBEAT_PAST_BLOCKS, "Too old block");
        require(blockhash(block_number) == block_hash, "Block mismatch");

        address payable masternode = _callerAddress();

        Status storage s = mn_status[masternode];

        require(s.seq_payouts > 0, "Not active");

        uint hearbeat_delay = block.timestamp - s.last_heartbeat;
        require(hearbeat_delay < MN_HEARTBEAT_INTERVAL_MAX, "Too late");
        require(hearbeat_delay > MN_HEARTBEAT_INTERVAL_MIN, "Too early");

        s.last_heartbeat = uint64(block.timestamp);
        s.sw_features = sw_features;

        emit Heartbeat(masternode);
    }

    function validate(address masternode) external {
        address caller = _callerAddress();
        require(caller != masternode, "Vote for self");

        //---
        Status storage cs = mn_status[caller];
        require(_isValid(caller, cs), "Not active caller");
        require(cs.votes > 0, "No more votes");

        cs.votes--;

        //---
        Status storage s = mn_status[masternode];

        require(s.seq_payouts > 0, "Not active target");

        if (s.validations < config[uint(Config.VotesMax)]) {
            s.validations++;
        }

        emit Validation(masternode, caller);
    }

    function isValid(address masternode) external view returns(bool) {
        return _isValid(masternode, mn_status[masternode]);
    }

    //===

    function _isValid(address masternode, Status storage mnstatus) internal view returns(bool) {
        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, masternode);
        return _checkStatus(mnstatus, mninfo) == ValidationStatus.MNValid;
    }

    function _checkStatus(
        Status storage mnstatus,
        StorageMasternodeRegistryV1.Info memory mninfo
    )
        internal view
        returns(ValidationStatus)
    {
        (uint balance, uint age) = IMasternodeToken(address(token_proxy.impl())).balanceInfo(mninfo.owner);

        if (balance != mninfo.collateral) {
            return ValidationStatus.MNCollaterIssue;
        }

        if ((block.timestamp - age) > mninfo.last_announced) {
            return ValidationStatus.MNCollaterIssue;
        }

        if (mnstatus.seq_payouts == 0) {
            return ValidationStatus.MNNotActive;
        }

        if ((block.timestamp - mnstatus.last_heartbeat) >= MN_HEARTBEAT_INTERVAL_MAX) {
            return ValidationStatus.MNHeartbeat;
        }

        return ValidationStatus.MNValid;
    }

    //===

    function count() external view returns(uint active, uint total, uint max_of_all_times) {
        active = mn_active;
        total = mn_announced;
        max_of_all_times = mn_total_ever;
    }

    //===

    function info(address masternode) external view
        returns(address owner, uint32 ipv4address, bytes32[2] memory enode, uint collateral)
    {
        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, masternode);
        require(mninfo.owner != address(0), "Unknown masternode");
        owner = mninfo.owner;
        ipv4address = mninfo.ipv4address;
        enode = [ mninfo.enode_0, mninfo.enode_1 ];
        collateral = mninfo.collateral;
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
            mninfo.owner,
            mninfo.last_announced,
            mninfo.ipv4address,
            mninfo.enode_0,
            mninfo.enode_1,
            mninfo.collateral,
            mninfo.prev,
            mninfo.next
        ) = v1info.masternodes(masternode);
    }

    //===

    function onCollateralUpdate(address owner) external
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

    function enumerate() external view returns(address[] memory masternodes) {
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

    // IGovernedContract
    //---------------------------------
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // IBlockReward
    //---------------------------------
    function reward() external payable {
        // NOTE: ensure to move of remaining from the previous times to Treasury
        //---
        uint diff = address(this).balance - msg.value;

        if (int(diff) > 0) {
            IBlockReward treasury = IBlockReward(address(treasury_proxy.impl()));
            treasury.reward.value(diff)();
        }

        //---
        if (msg.value != 0) {
            assert(gasleft() > GAS_RESERVE);
            assert(msg.value == address(this).balance);

            // solium-disable-next-line no-empty-blocks
            while ((gasleft() > GAS_RESERVE) && !_reward()) {}
        }
    }

    function _reward() internal returns(bool) {
        //---
        address masternode = current_masternode;
        uint8 payouts = current_payouts;

        if (masternode == address(0)) {
            return true;
        }

        StorageMasternodeRegistryV1.Info memory mninfo = _mnInfo(v1storage, masternode);

        Status storage mnstatus = mn_status[masternode];
        mnstatus.votes = uint32(config[uint(Config.VotePerCycle)]);
        uint validations = mnstatus.validations;
        ++payouts;

        if (payouts < mnstatus.seq_payouts) {
            current_payouts = payouts;
        } else {
            mnstatus.validations = 0;
            current_masternode = mninfo.next;
            current_payouts = 0;
        }

        // Reward logic
        //---
        ValidationStatus status = _checkStatus(mnstatus, mninfo);

        if (status == ValidationStatus.MNValid) {
            // solium-disable-next-line security/no-send
            if (!_canReward(validations) || mninfo.owner.send(msg.value)) {
                return true;
            }
        }

        // When not valid
        //---
        if (status == ValidationStatus.MNCollaterIssue) {
            // Immediate
            _denounce(masternode, mninfo.owner);
        } else if (mnstatus.seq_payouts > 0) {
            // Mark as inactive for later auto-cleanup
            mnstatus.seq_payouts = 0;
            mnstatus.inactive_since = uint64(block.timestamp);
            --mn_active;
            mnstatus.validations = 0;
            current_masternode = mninfo.next;
            current_payouts = 0;

            emit Deactivated(masternode);
        } else if ((block.timestamp - mnstatus.inactive_since) > config[uint(Config.CleanupPeriod)]) {
            // Auto-cleanup
            _denounce(masternode, mninfo.owner);
        }

        return false;
    }

    function _canReward(uint validations) internal view returns(bool) {
        uint active = mn_active;

        if (active < config[uint(Config.RequireVoting)]) {
            return true;
        }

        uint required = active / 2;
        uint max = config[uint(Config.VotesMax)];

        if (required > max) {
            required = max;
        }

        return (validations >= required);
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
