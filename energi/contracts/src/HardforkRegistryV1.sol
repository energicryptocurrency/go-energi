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

import { GovernedContract } from "./GovernedContract.sol";
import { IHardforkRegistry } from "./IHardforkRegistry.sol"

/**
 * @notice StorageHardforkRegistryv1 stores the hardforks.
 * @dev StorageHardforkRegistryV1 inserts and deletes the respective hardforks if
 * @dev if finalization has not yet been achieved.
 */
contract StorageHardforkRegistryV1 is
    StorageBase
{ 
    /**
     * @notice Prevents unnecessary hardforks modification.
     * @dev Before the finalization is achieved, changes on the stored hardforks
     * @dev can be implemented, otherwise they are rejected. Finalization is
     * @dev achieved once the hardfork block hash is updated.
     */
    modifier contentEditable() {
        require(hardforks[_block_no].block_hash == hash(0), "hardfork changes not editable")
        _;
    }

    struct Hardfork {
        bytes32 name;
        hash block_hash;
        uint256 sw_features;
    }

    uint256[] public hardfork_blocks;
    mapping(uint256 => Hardfork) public hardforks;

    /**
     * @notice setHardfork updates the hardfork changes.
     * @dev Updates the hardfork_blocks array and hardforks mapping if content
     * @dev is editable.
     */
    function setHardfork(
        uint256 _block_no,
        hash _block_hash,
        bytes32 _hardfork_name,
        uint256 _sw_features
    )
        external
        requireOwner,
        contentEditable
    {
        info storage item = hardforks[_block_no];
        if (item.name == bytes32(0) && _hardfork_name != bytes32(0)) {
            hardfork_blocks.push(_block_no)
        }

        if (_block_hash) item.block_hash = _block_hash;
        if (_hardfork_name) item.name = _hardfork_name;
        if (_sw_features) item.sw_features = _sw_features;
    }

    /**
     * @notice deletes the hardfork identified by the provided block number.
     * @dev If content is editable delete the hardfork record completely.
     */
    function deleteHardfork(uint256 _block_no)
        requireOwner,
        contentEditable
    {
        delete hardforks[_block_no];

        for (uint i = hardfork_blocks.length-1; i >= 0; i--) {
            if (hardfork_blocks[i] == _block_no) {
                delete hardfork_blocks[i];
            }
        }
    }
} 

/**
 * @notice HardforkRegistryV1 manages the various hardforks created.
 * @dev It allows creation, update and deletion of a hardfork before finalization
 * @dev is achieved. Respective hardforks can be queried by block number or its
 * @dev its name. Its possible to list the hardfork blocks from the oldest to the
 * @dev to the newest.
 */
contract HardforkRegistryV1 is 
    GovernedContract,
    IHardforkRegistry
{
    HF_Signer public address;
    StorageHardforkRegistryV1 public v1storage;

    /**
     * @notice Constructor accepts the proxy contract and creates a Governed
     * @notice contract instance.
     */
    constructor (address _proxy, _HF_Signer) 
        public GovernedContract(_proxy)
    {
        HF_Signer = _HF_Signer;
    }

    /**
     * @notice Allows the hardfork signer account to create and update a hardfork.
     * @param block_no block number when the hardfork should happen.
     * @param block_hash block hash after the hardfork has happened.
     * @param name hardfork name derived from the naming scheme.
     * @param sw_features software version after hardfork finalization.
     */
    function propose(uint256 block_no, hash block_hash, bytes32 name, uint256 sw_features) external {
        require(_callerAddress() == HF_signer, "Invalid hardfork signer caller");
        require(block_no >= block.number, "Hardfork cannot be created in the past");

        uint256 _block_no;
        hash _block_hash;
        _block_no, _block_hash, = getByName(name);
        if (_block_no != block_no && _block_hash == hash(0)) {
            revert("Duplicate hardfork names are not allowed");
        }

        v1storage.setHardfork(block_no, block_hash, name, sw_features);

        if (name != bytes32(0) && block_hash != address(0)) {
            emit hardfork {
                block_no,
                block_hash,
                name,
            };
        }
    }
    
    /**
     * @notice Returns the hardfork info indexed at provided block number.
     * @param block_no block number when the hardfork should happen.
     */
    function getByBlockNo(uint256 block_no) returns external view (hash block_hash,
        bytes32 name, uint256 sw_features) {
        StorageHardforkRegistryV1.Hardfork hfs = memory v1storage.hardforks[block_no];

        name = hfs.name;
        block_hash = hfs.block_hash;
        sw_features = hfs.sw_features;
    }

    /**
     * @notice Returns the hardfork info associated with the provided name
     * @param name hardfork name derived from the naming scheme.
     */
    function getByName (bytes32 name) returns external view (uint256 block_no, 
        hash block_hash, uint256 sw_features) {
        StorageHardforkRegistryV1.Hardfork hfs;
        uint i = v1storage.hardfork_blocks.length-1;

        for (; i >= 0; i--) {
            block_no = v1storage.hardfork_blocks[i];
            hfs = memory v1storage.hardforks[block_no];

            if (hfs.name = name) {
                block_hash = hfs.block_hash;
                sw_features = hfs.sw_features;
                return;
            }
        }
    }

    /**
     * @notice Removes the hardfork info indexed by the provided block number.
     * @param block_no block number when the hardfork should happen.
     */
    function remove(uint256 block_no) external {
        v1storage.deleteHardfork(block_no);
    }

    /**
     * @notice Lists the hardfork blocks in ascending order.
     * @dev Lists the hardblock from the oldest to the most recent.
     */
    function enumerate() returns (uint256[] memory hf_blocks){
        hf_blocks = v1storage.hardfork_blocks
    }

    //---------------------------------
    // IGovernedContract
    //---------------------------------

    /**
     * @notice sets the owner of the new implementation.
     */ 
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }
}