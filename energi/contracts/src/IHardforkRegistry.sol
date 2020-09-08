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
 * IHardforkRegistry defines the HardforkRegistry public interface.
 */
interface IHardforkRegistry {
    event Hardfork (
        uint256 block_no,
        bytes32 block_hash,
        bytes32 name
    );

    function propose(uint256 block_no, bytes32 name, bytes32 block_hash, uint256 sw_features) external;
    function getHardfork(uint256 block_no) external view returns(bytes32 name,
        bytes32 block_hash, uint256 sw_features);
    function remove(uint256 block_no) external;
    function enumerateAll() external view returns(bytes32[] memory all_hf_names);
    function enumeratePending() external view returns(bytes32[] memory pending_hf_names);
    function enumerateActive() external view returns(bytes32[] memory active_hf_names);
    function isActive(bytes32 name) external view returns (bool);
}