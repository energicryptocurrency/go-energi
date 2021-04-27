// Copyright 2021 The Energi Core Authors
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
        uint block_number;
        bytes32 block_hash;
        uint sw_features;
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
        if (hf.name != bytes32(0)) require(hf.block_number > block.number, "Hardfork is already in effect or doesn't exist");
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
        uint block_number,
        uint sw_features
    )
        external
        requireOwner
        requirePending(name)
        returns(bool new_hardfork)
    {
        require(block_number > block.number, "Hardfork is too soon");

        //name associated hardfork
        Hardfork storage hf = hardforks[name];

        //if Hardfork name is not present push as a new hardfork name
        if (hf.name == bytes32(0)) {
            hardfork_names.push(name);
            hf.name = name;
            new_hardfork = true;
        } else {
            require(hf.block_number > block.number, "Cannot modify active Hardfork");
            new_hardfork = false;
        }

        //set/modify hardfork activation block/version
        hf.block_number = block_number;
        hf.sw_features = sw_features;

    }

    /// @notice Once a hard fork has been finalized, it is no longer possible to change.
    /// @dev Once a hardfork has gone into effect, if there have been at least finalization_confirmations
    /// @dev blocks since the hard fork block, we store the hash of the hardfork block. Hard forks with
    /// @dev a block hash associated to them are considered final and may no longer be changed in any way.
    /// @param name The name of the hardfork to finalize
    /// @param finalization_confirmations The number of block confirmations before a hardfork is eligible to finalize.
    function finalize(bytes32 name, uint finalization_confirmations)
        external
        requireOwner
    {
        Hardfork storage hf = hardforks[name];
        require(hf.name != bytes32(0), "Hardfork doesn't exist");
        require(block.number > (hf.block_number + finalization_confirmations), "Hardfork not eligible for finalizing");
        require(hf.block_hash == bytes32(0), "Hardfork already finalized");
        hf.block_hash = blockhash(hf.block_number);
        require(hf.block_hash != bytes32(0), "No block hash to finalize");
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
        bool found = false;

        for (uint i = 0; i < hardfork_names.length; i++) {
            if (hardfork_names[i] == name) {
                found = true;

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

    /// @notice get the hardfork names array
    /// @return an array of hardfork names
    function get_names() external view requireOwner returns(bytes32[] memory)
    {
        return hardfork_names;
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
    uint finalization_confirmations;

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
    constructor(address _proxy, address _hf_signer, uint _finalization_confirmations)
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
    function add(bytes32 name, uint block_number, uint sw_features) external requireHardforkSigner
    {
        if (v1storage.set(name, block_number, sw_features)) {
            emit HardforkCreated(name, block_number, sw_features);
        }
    }

    /// @notice finalize a hard fork
    /// @dev may only be called by the hard fork signer
    /// @dev may only be called on a hard fork that has been active for some number of confirmations
    /// @dev emits HardforkFinalized when successful
    /// @param name The name of the hard fork to finalize
    function finalize(bytes32 name) external requireHardforkSigner
    {
        v1storage.finalize(name, finalization_confirmations);
        uint block_number;
        bytes32 block_hash;
        uint sw_features;
        (, block_number, block_hash, sw_features) = v1storage.hardforks(name);
        emit HardforkFinalized(name, block_number, block_hash, sw_features);
    }

    /// @notice remove a hard fork from the registry
    /// @dev hard forks which are active or finalized may not be removed
    /// @dev emits HardforkRemoved if a hard fork was removed from the registry
    /// @param name The name of the hard fork to remove
    /// @return true when the hard fork was removed, false otherwise
    function remove(bytes32 name) external requireHardforkSigner returns(bool)
    {
        if (v1storage.remove(name)) {
            emit HardforkRemoved(name);
            return true;
        }
        return false;
    }

    /// @notice get the information for a hard fork
    /// @param name The name of the hard fork to look up
    /// @return state the state of the hard fork: -1: no hard fork, 0: hard fork pending, 1: hard fork active, 2: hard fork final
    /// @return block_number the block number on which the hard fork will become active
    /// @return block_hash the hash of the block on which a finalized hard fork became active
    /// @return sw_fetaures A version integer describing the minimum software required for the hard fork
    function get(bytes32 name) external view returns(uint block_number, bytes32 block_hash, uint sw_features)
    {
        bytes32 _name;
        (_name, block_number, block_hash, sw_features) = v1storage.hardforks(name);
        require(_name != bytes32(0), "no such hard fork");
    }

    /// @notice get the names of all the hard forks
    /// @return A list of hard fork names
    function enumerate() external view returns(bytes32[] memory)
    {
        return v1storage.get_names();
    }

    /// @notice Get the names of pending hard forks
    /// @return A list of pending hard fork names
    function enumeratePending() external view returns (bytes32[] memory)
    {
        bytes32[] memory hf_names = v1storage.get_names();
        uint names_count = hf_names.length;

        //count the number of pending hfs
        uint pendingNum = 0;
        for (uint i = 0; i < names_count; i++) {
            uint block_number;
            (, block_number, ,) = v1storage.hardforks(hf_names[i]);
            if (block.number < block_number) {
                pendingNum++;
            }
        }

        //collect pending hfs
        bytes32[] memory pending = new bytes32[](pendingNum);
        uint ind = 0;
        for (uint i = 0; i < names_count; i++) {
            uint block_number;
            (, block_number, ,) = v1storage.hardforks(hf_names[i]);
            if (block.number < block_number) {
                pending[ind] = hf_names[i];
                ind++;
            }
        }

        return pending;
    }

    /// @notice Get the names of active hard forks
    /// @return A list of active hard fork names
    function enumerateActive() external view returns (bytes32[] memory)
    {
        bytes32[] memory hf_names = v1storage.get_names();
        uint names_count = hf_names.length;

        //count the number of active hfs
        uint activeNum = 0;
        for (uint i = 0; i < names_count; i++) {
            uint block_number;
            (, block_number, ,) = v1storage.hardforks(hf_names[i]);
            if (block.number >= block_number) {
                activeNum++;
            }
        }

        //collect active hfs
        bytes32[] memory active = new bytes32[](activeNum);
        uint ind = 0;
        for (uint i = 0; i < names_count; i++) {
            uint block_number;
            (, block_number, ,) = v1storage.hardforks(hf_names[i]);
            if (block.number >= block_number) {
                active[ind] = hf_names[i];
                ind++;
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
        uint block_number;
        (_name, block_number, ,) = v1storage.hardforks(name);

        return ((block.number >= block_number) && (_name != bytes32(0)));
    }

    /// @notice move data to new hardfork registry during upgrade
    function _migrate(IGovernedContract _oldImpl) internal
    {
        // move v1storage from oldImpl to this impl
        // other storage variables finalization period and HF signer are set by the constructor
        v1storage.kill();
        HardforkRegistryV1 oldinstance = HardforkRegistryV1(address(uint160(address(_oldImpl))));
        v1storage = oldinstance.v1storage();
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
