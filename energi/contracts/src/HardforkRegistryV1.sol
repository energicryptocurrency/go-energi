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

import { IGovernedContract } from "./IGovernedContract.sol";
import { GovernedContractAutoProxy } from "./GovernedContractAutoProxy.sol";
import { IHardforkRegistry } from "./IHardforkRegistry.sol";
import { NonReentrant } from "./NonReentrant.sol";
import { StorageBase }  from "./StorageBase.sol";

/// @title Storage for the HardforkRegistry
/// @notice Manages HardforkRegistry storage operations.
/// @dev All storage operations require the hard fork is not in effect or finalized.
contract StorageHardforkRegistryV1 is StorageBase
{
    struct Hardfork {
        bytes32 name;
        uint256 block_number;
        bytes32 block_hash;
        uint256 sw_features;
    }

    bytes32[] public hardfork_names;
    mapping(bytes32 => Hardfork) public hardforks;

    /// @notice A hard fork which is in effect or finalized may not be modified.
    /// @param name The name of the hardfork to check if it is still pending
    modifier requirePending(bytes32 name)
    {
        require(name != bytes32(0), "Hardfork name cannot be empty");
        Hardfork storage hf = hardforks[name];
        // once a hard fork block number happens, any change to the hard fork would be dangerous
        require(hf.block_number < block.number, "Hardfork is already in effect");
        // once the hard fork is finalized we disallow any changes
        require(hf.block_hash == bytes32(0), "Hardfork is already finalized");
        _;
    }

    /// @notice adds a hardfork to the registry or updates a pending hardfork
    /// @param name The name of the hardfork to create or update. Cannot be empty.
    /// @param block_number The block number when the hardfork will go into effect.
    /// @param sw_features A software version number describing the minimum software needed for the hardfork.
    function set(
        bytes32 name,
        uint256 block_number,
        uint256 sw_features
    )
        external
        requireOwner
        requirePending(name)
    {
        require(block_number < block.number, "Hardfork is too soon");

        Hardfork storage hf = hardforks[name];

        // save new mapping key
        if (hf.name == bytes32(0)) {
            hardfork_names.push(name);
        }

        hf.name = name;
        hf.block_number = block_number;
        hf.sw_features = sw_features;
    }

    /// @notice Once a hard fork has been finalized, it is no longer possible to change.
    /// @dev Once a hardfork has gone into effect, if there have been at least finalization_confirmations
    /// @dev blocks since the hard fork block, we store the hash of the hardfork block. Hard forks with
    /// @dev a block hash associated to them are considered final and may no longer be changed in any way.
    /// @param name The name of the hardfork to finalize
    /// @param finalization_confirmations The number of block confirmations before a hardfork is eligible to finalize.
    function finalize(bytes32 name, uint256 finalization_confirmations)
        external
        requireOwner
        requirePending(name)
    {
        Hardfork storage hf = hardforks(name);
        require(hf.name != bytes32(0), "Hardfork doesn't exist");
        require(block.number > (hf.block_number + finalization_confirmations), "Hardfork not eligible for finalizing");
        bytes32 hardfork_block = block.blockhash(hf.block_number);
        require(hardfork_block != bytes32(0), "No block hash to finalize");
        hf.block_hash = hardfork_block;
    }

    /// @notice removes a hard fork from the registry
    /// @param name The name of the hardfork to remove
    /// @return true when the hardfork is removed, false otherwise
    function remove(bytes32 name)
        external
        requireOwner
        requirePending(name)
        returns (bool)
    {
        bool found=false;

        for (uint i = 0; i < hardfork_names.length; i++) {
            if (hardfork_names[i] == name) {
                found=true;

                // remove the name from the hardfork_names array
                for (uint k = i; k < (hardfork_names.length - 1); k++) {
                   hardfork_names[k] = hardfork_names[k + 1];
                }
                hardfork_names.pop();

                // remove the hardfork data from the mapping
                delete hardforks[name];
                break;
            }
        }

        return found;
    }
}

/// @title Hardfork Registry
/// @notice Manages and enumerates hardforks.
/// @dev Any storage operations must be done by HF_signer
contract HardforkRegistryV1 is
    GovernedContractAutoProxy,
    IHardforkRegistry,
    NonReentrant
{
    address public hf_signer;
    StorageHardforkRegistryV1 public v1storage;
    uint256 finalization_confirmations;

    /// @notice Only hf_signer is allowed to update the HardforkRegistry
    modifier requireHardforkSigner()
    {
        require(_callerAddress() == hf_signer, "only hf_signer is allowed to call this function");
        _;
    }

    /// @notice Construct a new HardforkRegistry
    /// @param _proxy The proxy address of the HardforkRegistry. If address(0) is used, a new proxy will be created.
    /// @param _hf_signer The address of the key responsible for managing hardforks.
    /// @param _finalization_confirmations The number of block confirmations before a hardfork is eligible to finalize.
    constructor(address _proxy, address _HardforkSigner, uint256 _finalization_confirmations)
        public GovernedContractAutoProxy(_proxy)
    {
        v1storage = new StorageHardforkRegistryV1();
        hf_signer = _hf_signer;
        finalization_confirmations = _finalization_confirmations;
    }

    /// @notice add a new hard fork to the registry, or update an existing hard fork
    /// @dev may only be called by the hard fork signer
    /// @dev hard forks which are active or finalized may not be updated
    /// @dev emits HardforkCreated if a new hard fork was added to the registry
    /// @param name The name of the hard fork to add or update
    /// @param block_number The block number when the hard fork will go into effect
    /// @param sw_features A version integer describing the minimum software required for the hard fork
    function add(bytes32 name, uint256 block_number, uint256 sw_features) external requireHardforkSigner
    {
        v1storage.set(name, block_number, sw_features);
    }

    /// @notice finalize a hard fork
    /// @dev may only be called by the hard fork signer
    /// @dev may only be called on a hard fork that has been active for some number of confirmations
    /// @dev emits HardforkFinalized when successful
    /// @param name The name of the hard fork to finalize
    function finalize(bytes32 name) external requireHardforkSigner
    {
        v1storage.finalize(name);
    }

    /// @notice remove a hard fork from the registry
    /// @dev hard forks which are active or finalized may not be removed
    /// @dev emits HardforkRemoved if a hard fork was removed from the registry
    /// @param name The name of the hard fork to remove
    /// @return true when the hard fork was removed, false otherwise
    function remove(bytes32 name) returns(bool) external requireHardforkSigner
    {
        return v1storage.remove(name);
    }

    /// @notice get the information for a hard fork
    /// @param name The name of the hard fork to look up
    /// @return state the state of the hard fork: -1: no hard fork, 0: hard fork pending, 1: hard fork active, 2: hard fork final
    /// @return block_number the block number on which the hard fork will become active
    /// @return block_hash the hash of the block on which a finalized hard fork became active
    /// @return sw_fetaures A version integer describing the minimum software required for the hard fork
    function get(bytes32 name) external view returns(int state, uint256 block_number, bytes32 block_hash, uint256 sw_features)
    {
        // default state of -1 unless we can find this hard fork
        state = -1;

        // look up the hard fork
        bytes32 _name;
        (_name, block_number, block_hash, sw_features) = v1storage.hardforks(name);

        // check if the hard fork is found
        if (_name != bytes32(0)) {
            state = 0;
        }
        // check if the hard fork is active
        if (block_number >= block.number) {
            state = 1;
        }
        // check if the hard fork is finalized
        if (block_hash != bytes32(0)) {
            state = 2;
        }
    }

    /// @notice get the names of all the hard forks
    /// @return A list of hard fork names
    function enumerate() external view returns(bytes32[] memory)
    {
        return v1storage.hardfork_names();
    }

    /// @notice Get the names of pending hard forks
    /// @return A list of pending hard fork names
    function enumeratePending() external view returns (bytes32[] memory)
    {
        bytes32[] storage hf_names = v1storage.hardfork_names();
        uint names_count = hf_names.length;

        bytes32[] memory pending;
        for (uint i = 0; i < names_count; i++) {
            (, block_number, ,) = v1storage.hardforks(hf_names[i]);
            if (block.number < block_number) {
                pending.push(hf_names[i]);
            }
        }

        return pending;
    }

    /// @notice Get the names of active hard forks
    /// @return A list of active hard fork names
    function enumerateActive() external view returns (bytes32[] memory)
    {
        bytes32[] storage hf_names = v1storage.hardfork_names();
        uint names_count = hf_names.length;

        bytes32[] memory active;
        for (uint i = 0; i < names_count; i++) {
            (, block_number, ,) = v1storage.hardforks(hf_names[i]);
            if (block.number >= block_number) {
                active.push(hf_names[i]);
            }
        }

        return active;
    }

    /// @notice check whether a hard fork is active or not
    /// @dev A hard fork is considered active if the current block number is greater than or equal to the hard fork block number
    /// @param name The name of the hard fork to check whether or not it is active
    /// @return (block.number >= hf.block_number)
    function isActive(bytes32 name) external view returns(bool)
    {
        bytes32 _name;
        uint256 block_number;
        (_name, block_number, ,) = v1storage.hardforks(names);
        return ((block.number >= block_number) && (_name != bytes32(0));
    }

    /// @notice move data to new hardfork registry during upgrade
    function _migrate(IGovernedContract _newimpl) internal
    {
        // TODO: implement me
    }

    /// @notice called to finalize a governance upgrade
    function _destroy(IGovernedContract _newImpl) internal
    {
        v1storage.setOwner(_newImpl);
    }

    /// @notice fallback function not allowed
    function () external payable
    {
        revert("Not supported");
    }
}
