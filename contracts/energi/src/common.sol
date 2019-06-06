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

/**
 * Global constants with no storage space.
 */
contract GlobalConstants {
    address constant internal TREASURY = address(0x301);
    address constant internal MASTERNODE_REGISTRY = address(0x302);
    address constant internal STAKE_REWARD = address(0x303);
    address constant internal BACKBONE_REWARD = address(0x304);
    address constant internal SPORK_REGISTRY = address(0x305);
    address constant internal CHECKPOINT_REGISTRY = address(0x306);
    address constant internal BLACKLIST_REGISTRY = address(0x307);
    address constant internal MIGRATION_CONTRACT = address(0x308);
    address constant internal MASTERNODE_TOKEN = address(0x309);
    address constant internal GEN2_ADDR_RECOVERY = address(0x310);
}

/**
 * Base interface for upgradable contracts
 */
interface IGovernedContract {
    // It must check that the caller is the new instance,
    // and self destruct to the caller address.
    //
    // NOTE: all data migration is assumed to be done at
    //       the new instance construction time.
    function migrate() external;
}

/**
 * Base interface for constructs which receive block rewards
 */
interface IBlockReward {
    // NOTE: it must NEVER fail
    function reward(uint amount) external payable;

    // NOTE: it must NEVER fail
    function getReward(uint block_number) external view returns(uint);
}

