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
        bytes32 block_hash;
        for (uint i = 0; i <hardfork_names.length; i++) {
            Hardfork memory hf = hardforks[hardfork_names[i]];
            if (hf.block_number == _block_no) {
                block_hash = hf.block_hash;
            }
        }
        require(block_hash == bytes32(0), "hardfork changes not editable");
        _;
    }

    struct Hardfork {
        bytes32 name;
        uint256 block_number;
        bytes32 block_hash;
        uint256 sw_features;
    }

    bytes32[] public hardfork_names;
    mapping(bytes32 => Hardfork) public hardforks;

    /**
     * @notice setHardfork updates the hardfork changes.
     * @dev Updates the hardfork_names array and hardforks mapping if content
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
        Hardfork storage item = hardforks[_hardfork_name];
        // Update the hardfork names array once the first hardfork name is set.
        if (item.name == bytes32(0) && _hardfork_name != bytes32(0)) {
            hardfork_names.push(_hardfork_name);
        }

        item.block_number = _block_no;
        if (_sw_features != 0) item.sw_features = _sw_features;
        if (_block_hash != bytes32(0)) item.block_hash = _block_hash;
        if (_hardfork_name != bytes32(0)) item.name = _hardfork_name;
    }

    /**
     * @notice deletes the hardfork identified by the provided hardfork name.
     * @dev If content is editable delete the hardfork record completely.
     */
    function deleteHardfork(bytes32 _hardfork_name, uint256 _block_no)
        external
        requireOwner
        contentEditable(_block_no)
    {
        delete hardforks[_hardfork_name];

        for (uint i = 0; i <hardfork_names.length; i++) {
            if (hardfork_names[i] == _hardfork_name) {
                uint k = i;
                for (; k <hardfork_names.length-1; k++) {
                   hardfork_names[k] = hardfork_names[k+1];
                }

               hardfork_names.pop();
                break;
            }
        }
    }

    /**
     * @notice Returns all the hardfork names available.
     */
    function getHardForkNames() external view returns(bytes32[] memory) {
        return hardfork_names;
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
    GovernedContractAutoProxy,
    IHardforkRegistry,
    NonReentrant
{
    address public HF_signer;
    StorageHardforkRegistryV1 public v1storage;

    /**
     * @dev Finalization interval is a period from x blocks behind the current
     * @dev block number. Its the period after which a hardfork will be considered
     * @dev immutable if set with a block hash or invalid if not.
     */
    uint256 internal HF_FINALIZATION_INTERVAL;

    /**
     * @notice Constructor accepts the proxy contract and creates a Governed
     * @notice contract instance.
     */
    constructor (address _proxy, address _HF_signer, uint256 _HF_finalization_period)
        public GovernedContractAutoProxy(_proxy)
    {
        v1storage = new StorageHardforkRegistryV1();
        HF_signer = _HF_signer;
        HF_FINALIZATION_INTERVAL = _HF_finalization_period;
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
        require(name != bytes32(0), "Hardfork name cannot be empty");

        bytes32 _name;
        uint256 _block_no;
        (_name,_block_no,,) = v1storage.hardforks(name);
        if (_block_no > 0) {
            // Hardfork already exist: Update is currently happening.
            require(_block_no == block_no, "Duplicate hardfork names are not allowed");
            require((block_no + HF_FINALIZATION_INTERVAL) >= block.number, "Hardfork finalization interval exceeded");

            if (block_no < block.number) {
                // During hardfork finalization period, block hash cannot be empty.
                require(block_hash != bytes32(0), "HF finalization block hash cannot be empty");
            }
        } else {
            // Hardfork doesn't exist: new instance will be created.
            require(block_no > (block.number + HF_FINALIZATION_INTERVAL), "Hardfork cannot be scheduled immediately.");
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
    function getHardfork(bytes32 _hardfork_name)
        external
        view
        returns(uint256 block_no, bytes32 block_hash, uint256 sw_features)
    {
        (,block_no, block_hash, sw_features) = v1storage.hardforks(_hardfork_name);
    }

    /**
     * @notice Removes the hardfork info indexed by the provided block number.
     * @param block_no block number when the hardfork should happen.
     */
    function remove(uint256 block_no) external noReentry {
        require(_callerAddress() == HF_signer, "Invalid hardfork signer caller");

        bytes32 _name;
        uint256 _block_no;
        bytes32 _block_hash;
        (_name, _block_no, _block_hash,) = _hardforkInfo(v1storage, block_no);
        require(_name != bytes32(0), "Hardfork name cannot be empty");
        require(_block_hash == bytes32(0), "Finalized hardfork cannot be deleted");

        v1storage.deleteHardfork(_name, _block_no);
    }

    /**
     * @notice Lists the all the hardfork names in the order they were created.
     */
    function enumerateAll() public view returns(bytes32[] memory){
        return v1storage.getHardForkNames();
    }

    /**
     * @notice Lists all the pending hardfork names in the order they were created.
     * @dev Two for-loops used guarantee the final array is not permanently stored.
     */
    function enumeratePending() public view returns(bytes32[] memory){
        uint index;
        uint pending_HFs;
        bytes32[] memory all_names = enumerateAll();
        for (uint i = 0; i < all_names.length; i++) {
            bytes32 name = all_names[i];
            if (!isActive(name)) {
                pending_HFs++;
            }
        }

        bytes32[] memory pending_names = new bytes32[](pending_HFs);
        for (uint i = 0; i < all_names.length; i++) {
            bytes32 name = all_names[i];
            if (!isActive(name)) {
                pending_names[index] = name;
                index++;
            }
        }
        return pending_names;
    }

    /**
     * @notice Lists all the active hardfork names in the order they were created.
     * @dev Two for-loops used guarantee the final array is not permanently stored.
     */
    function enumerateActive() public view returns(bytes32[] memory){
        uint index;
        uint active_HFs;
        bytes32[] memory all_names = enumerateAll();
        for (uint i = 0; i < all_names.length; i++) {
            bytes32 name = all_names[i];
            if (isActive(name)) {
                active_HFs++;
            }
        }

        bytes32[] memory active_names = new bytes32[](active_HFs);
        for (uint i = 0; i < all_names.length; i++) {
            bytes32 name = all_names[i];
            if (isActive(name)) {
                active_names[index] = name;
                index++;
            }
        }
        return active_names;
    }

    /**
     * @notice Checks if the hardfork block has been achieved.
     * @param name hardfork name to be searched.
     */
    function isActive(bytes32 name) public view returns (bool) {
        // if name is empty return false
        if (name == bytes32(0)) return false;

        uint256 block_no;
        (,block_no,,) = v1storage.hardforks(name);
        return (block.number >= block_no && block_no != 0);
    }

    /**
     * @notice Extracts the hardfork information identified by the block number.
     * @dev privately accessed the function returning a memory instance of the hardfork.
     */
    function _hardforkInfo(StorageHardforkRegistryV1 store, uint256 _block_no)
        internal
        view
        returns (bytes32 name, uint256 block_no, bytes32 block_hash, uint256 sw_features)
    {
        bytes32[] memory hf_names_array = store.getHardForkNames();
        for (uint i = 0; i < hf_names_array.length; i++) {
            bytes32 _name;
            uint256 _blockNo;
            bytes32 _block_hash;
            uint256 _sw_features;
            (_name, _blockNo, _block_hash, _sw_features) = store.hardforks(hf_names_array[i]);
            if (_blockNo == _block_no) {
                name = _name;
                block_no = _blockNo;
                block_hash = _block_hash;
                sw_features = _sw_features;
                break;
            }
        }
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
