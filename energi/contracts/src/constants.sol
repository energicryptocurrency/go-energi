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

/**
 * Global constants with no storage space.
 * NOTE: it could be a library, but Solidity does not support such case.
 */
contract GlobalConstants {
    address payable constant internal TREASURY = address(0x301);
    address payable constant internal MASTERNODE_REGISTRY = address(0x302);
    address payable constant internal STAKE_REWARD = address(0x303);
    address payable constant internal BACKBONE_REWARD = address(0x304);
    address payable constant internal SPORK_REGISTRY = address(0x305);
    address payable constant internal CHECKPOINT_REGISTRY = address(0x306);
    address payable constant internal BLACKLIST_REGISTRY = address(0x307);
    address constant internal MIGRATION_CONTRACT = address(0x308);
    address payable constant internal MASTERNODE_TOKEN = address(0x309);
    address constant internal SYSTEM_FAUCET = address(0x310);

    uint constant internal FEE_UPGRADE_V1 = 10000 ether;
    uint constant internal FEE_BUDGET_V1 = 100 ether;
    uint constant internal FEE_CHECKPOINT_V1 = 1000 ether;
    uint constant internal FEE_BLACKLIST_V1 = 1000 ether;
    uint constant internal FEE_BLACKLIST_REVOKE_V1 = 100 ether;
    uint constant internal FEE_BLACKLIST_DRAIN_V1 = 100 ether;

    uint constant internal PERIOD_UPGRADE_MIN = 2 weeks;
    uint constant internal PERIOD_UPGRADE_MAX = 365 days;
    uint constant internal PERIOD_BUDGET_MIN = 2 weeks;
    uint constant internal PERIOD_BUDGET_MAX = 30 days;
    uint constant internal PERIOD_CHECKPOINT = 1 weeks;
    uint constant internal PERIOD_BLACKLIST = 1 weeks;

    uint8 constant internal QUORUM_MIN = 1;
    uint8 constant internal QUORUM_MAJORITY = 51;
    uint8 constant internal QUORUM_MAX = 100;

    uint constant internal REWARD_STAKER_V1 = 2.28 ether;
    uint constant internal REWARD_BACKBONE_V1 = 2.28 ether;
    uint constant internal REWARD_MASTERNODE_V1 = 9.14 ether;
    uint constant internal REWARD_TREASURY_V1 = 184000 ether;

    uint constant internal MN_COLLATERAL_MIN = 10000 ether;
    uint constant internal MN_COLLATERAL_MAX = 100000 ether;

    uint constant internal MN_HEARTBEAT_INTERVAL = 60 minutes;
    uint constant internal MN_HEARTBEAT_INTERVAL_MIN = MN_HEARTBEAT_INTERVAL / 2;
    uint constant internal MN_HEARTBEAT_INTERVAL_MAX = MN_HEARTBEAT_INTERVAL * 2;
    uint constant internal MN_HEARTBEAT_PAST_BLOCKS = 10;

    uint constant internal BUDGET_AMOUNT_MIN = FEE_BUDGET_V1;
    uint constant internal BUDGET_AMOUNT_MAX = REWARD_TREASURY_V1;
    uint constant internal BUDGET_PROPOSAL_MAX = 100;
}
