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

/// @title Hardfork Registry Interface
/// @notice Provides events and function signatures for a Hardfork Registry
interface IHardforkRegistry {
    event HardforkCreated (
        bytes32 indexed name,
        uint256 block_number,
        uint256 sw_features
    );

    event HardforkFinalized (
        bytes32 indexed name,
        uint256 block_number,
        bytes32 block_hash,
        uint256 sw_features
    );

    event HardforkRemoved (bytes32 indexed name);

    /// @notice add a new hard fork to the registry, or update an existing hard fork
    /// @dev may only be called by the hard fork signer
    /// @dev hard forks which are active or finalized may not be updated
    /// @dev emits HardforkCreated if a new hard fork was added to the registry
    /// @param name The name of the hard fork to add or update
    /// @param block_number The block number when the hard fork will go into effect
    /// @param sw_features A version integer describing the minimum software required for the hard fork
    function add(bytes32 name, uint256 block_number, uint256 sw_features) external;

    /// @notice finalize a hard fork
    /// @dev may only be called by the hard fork signer
    /// @dev may only be called on a hard fork that has been active for some number of confirmations
    /// @dev emits HardforkFinalized when successful
    /// @param name The name of the hard fork to finalize
    function finalize(bytes32 name) external;

    /// @notice remove a hard fork from the registry
    /// @dev hard forks which are active or finalized may not be removed
    /// @dev emits HardforkRemoved if a hard fork was removed from the registry
    /// @param name The name of the hard fork to remove
    /// @return true when the hard fork was removed, false otherwise
    function remove(bytes32 name) external returns(bool);

    /// @notice get the information for a hard fork
    /// @param name The name of the hard fork to look up
    /// @return state the state of the hard fork: -1: no hard fork, 0: hard fork pending, 1: hard fork active, 2: hard fork final
    /// @return block_number the block number on which the hard fork will become active
    /// @return block_hash the hash of the block on which a finalized hard fork became active
    /// @return sw_fetaures A version integer describing the minimum software required for the hard fork
    function get(bytes32 name) external view returns(int state, uint256 block_number, bytes32 block_hash, uint256 sw_features);

    /// @notice get the names of all the hard forks
    /// @return A list of hard fork names
    function enumerate() external view returns(bytes32[] memory);

    /// @notice Get the names of pending hard forks
    /// @return A list of pending hard fork names
    function enumeratePending() external view returns (bytes32[] memory);

    /// @notice Get the names of active hard forks
    /// @return A list of active hard fork names
    function enumerateActive() external view returns (bytes32[] memory);

    /// @notice check whether a hard fork is active or not
    /// @dev A hard fork is considered active if the current block number is greater than or equal to the hard fork block number
    /// @param name The name of the hard fork to check whether or not it is active
    /// @return (block.number >= hf.block_number)
    function isActive(bytes32 name) external view returns(bool);
}
