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
import { GovernedContract } from "./GovernedContract.sol";
import { IHardforkRegistry } from "./IHardforkRegistry.sol";
import { NonReentrant } from "./NonReentrant.sol";
import { StorageBase }  from "./StorageBase.sol";

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
    modifier contentEditable(uint256 _block_no) {
        require(hardforks[_block_no].block_hash == bytes32(0), "hardfork changes not editable");
        _;
    }

    struct Hardfork {
        bytes32 name;
        bytes32 block_hash;
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
        bytes32 _block_hash,
        bytes32 _hardfork_name,
        uint256 _sw_features
    )
        external
        requireOwner
        contentEditable(_block_no)
    {
        Hardfork storage item = hardforks[_block_no];
        // Update the hardfork blocks once the first hardfork name is set
        if (item.name == bytes32(0) && _hardfork_name != bytes32(0)) {
            hardfork_blocks.push(_block_no);
        }

        if (_block_hash != bytes32(0)) item.block_hash = _block_hash;
        if (_hardfork_name != bytes32(0)) item.name = _hardfork_name;
        if (_sw_features != 0) item.sw_features = _sw_features;
    }

    /**
     * @notice deletes the hardfork identified by the provided block number.
     * @dev If content is editable delete the hardfork record completely.
     */
    function deleteHardfork(uint256 _block_no)
        external
        requireOwner
        contentEditable(_block_no)
    {
        delete hardforks[_block_no];

        for (uint i = hardfork_blocks.length-1; i >= 0; i--) {
            if (hardfork_blocks[i] == _block_no) {
                delete hardfork_blocks[i];
            }
        }
    }

    /**
     * @notice Return the hardfork blocks.
     */
    function getHardForkBlocks() external view returns(uint256[] memory) {
        return hardfork_blocks;
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
    IHardforkRegistry,
    NonReentrant
{
    address public HF_signer;
    StorageHardforkRegistryV1 public v1storage;

    /**
     * @notice Constructor accepts the proxy contract and creates a Governed
     * @notice contract instance.
     */
    constructor (address _proxy, address _HF_signer)
        public GovernedContract(_proxy)
    {
        v1storage = new StorageHardforkRegistryV1();
        HF_signer = _HF_signer;
    }

    /**
     * @notice Allows the hardfork signer account to create and update a hardfork.
     * @param block_no block number when the hardfork should happen.
     * @param name hardfork name derived from the naming scheme.
     * @param block_hash block hash after the hardfork has happened.
     * @param sw_features software version after hardfork finalization.
     */
    function propose(uint256 block_no, bytes32 name, bytes32 block_hash, uint256 sw_features)
        external
        noReentry
    {
        require(_callerAddress() == HF_signer, "Invalid hardfork signer caller");
        require(block_no >= block.number, "Hardfork cannot be created in the past");

        uint256 _block_no;
        bytes32 _block_hash;
        (_block_no, _block_hash,) = getByName(name);
        if (_block_no != block_no && _block_hash == bytes32(0)) {
            revert("Duplicate hardfork names are not allowed");
        }

        v1storage.setHardfork(block_no, block_hash, name, sw_features);

        if (name != bytes32(0) && block_hash != bytes32(0)) {
            emit Hardfork (
                block_no,
                block_hash,
                name
            );
        }
    }

    /**
     * @notice Returns the hardfork info indexed at provided block number.
     * @param block_no block number when the hardfork should happen.
     */
    function getByBlockNo(uint256 block_no)
        external
        view
        returns(bytes32 name, bytes32 block_hash, uint256 sw_features)
    {
        StorageHardforkRegistryV1.Hardfork memory hfs = _hardforkInfo(v1storage, block_no);

        name = hfs.name;
        block_hash = hfs.block_hash;
        sw_features = hfs.sw_features;
    }

    /**
     * @notice Returns the hardfork info associated with the provided name
     * @param name hardfork name derived from the naming scheme.
     */
    function getByName (bytes32 name)
        public
        view
        returns(uint256 block_no, bytes32 block_hash, uint256 sw_features)
    {
        StorageHardforkRegistryV1.Hardfork memory hfs;
        uint i = v1storage.getHardForkBlocks().length - 1;

        for (; i >= 0; i--) {
            block_no = v1storage.hardfork_blocks(i);
            hfs = _hardforkInfo(v1storage, block_no);

            if (hfs.name == name) {
                block_hash = hfs.block_hash;
                sw_features = hfs.sw_features;
                return (block_no, block_hash, sw_features);
            }
        }
    }

    /**
     * @notice Removes the hardfork info indexed by the provided block number.
     * @param block_no block number when the hardfork should happen.
     */
    function remove(uint256 block_no) external noReentry {
        StorageHardforkRegistryV1.Hardfork memory hfs = _hardforkInfo(v1storage, block_no);
        require(hfs.block_hash == bytes32(0), "Finalized hardfork cannot be deleted");

        v1storage.deleteHardfork(block_no);
    }

    /**
     * @notice Lists the hardfork blocks in ascending order.
     * @dev Lists the hardblock from the oldest to the most recent.
     */
    function enumerate() external returns(uint256[] memory hf_blocks){
        hf_blocks = v1storage.getHardForkBlocks();
    }

    /**
     * @notice Extracts the hardfork information identified by the block number.
     * @dev privately accessed the function returning a memory instance of the hardfork.
     */
    function _hardforkInfo(StorageHardforkRegistryV1 store, uint256 _block_no)
        internal
        view
        returns (StorageHardforkRegistryV1.Hardfork memory hfs)
    {
        (
            hfs.name,
            hfs.block_hash,
            hfs.sw_features
        ) = store.hardforks(_block_no);
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

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}