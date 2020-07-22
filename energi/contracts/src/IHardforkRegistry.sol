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
pragma solidity 0.5.16;

/**
 * IHardforkRegistry defines the public HardforkRegistry public interface.
 */
interface IHardforkRegistry {
    event hardfork {
        uint256 block_no;
        hash block_hash;
        bytes32 name;
    }

    function propose(uint256 block_no, hash block_hash, bytes32 name, uint256 sw_features) external;
    function getByBlockNo(uint256 block_no) returns external view (hash block_hash,
        bytes32 name, uint256 sw_features);
    function getByName (bytes32 name) returns external view (uint256 block_no, 
        hash block_hash, uint256 sw_features);
    function remove(uint256 block_no) external;
    function enumerate() returns (uint256[] memory hf_blocks);
}