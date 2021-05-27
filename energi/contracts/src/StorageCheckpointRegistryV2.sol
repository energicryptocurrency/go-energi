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
// pragma experimental SMTChecker;
pragma experimental ABIEncoderV2;


import { ICheckpoint } from "./ICheckpoint.sol";
import { StorageBase }  from "./StorageBase.sol";

/**
 * Permanent storage of Checkpoint Registry V3 data.
 */
contract StorageCheckpointRegistryV2 is StorageBase {

    // main storage data structure for queue implementation
    mapping(uint => ICheckpoint) public checkpoints;
    // starting index(key) for checkpoints
    uint128 startingKeyIndex;
    // number of checkpoints currently stored
    uint128 size;
    // number of stored checkpoints' limit
    uint constant maxSize = 10;



    // push new checkpoint
    function add(ICheckpoint cp) external requireOwner {
        // if queue is full and needs first element to be deleted
        if (size == maxSize)  {
          delete checkpoints[startingKeyIndex];
          checkpoints[startingKeyIndex + size] = cp;
          startingKeyIndex++;
        } else {
          checkpoints[startingKeyIndex + size] = cp;
          size++;
        }
    }

    // pop first element
    function pop() external requireOwner {
      // nothing to pop
      if (size == 0) return;

      // remove last element
      delete checkpoints[startingKeyIndex];
      startingKeyIndex++;
      size--;
    }


    // for removal we find the checkpoint and move the right part of the queue to the left
    function remove(ICheckpoint cp) external  requireOwner returns(bool found) {
      uint foundCpIndex;
      found = false;
      // find the cp in map
      (uint number_1, bytes32 hash_1,  ) = cp.info();
      for (foundCpIndex = startingKeyIndex; foundCpIndex < startingKeyIndex + size; foundCpIndex++) {
          (uint number_2, bytes32 hash_2,  ) = checkpoints[foundCpIndex].info();
          if (number_1 == number_2 && hash_1 == hash_2) {
            found = true;
            break;
          }

      }

      // if we found the checkpoint
      if (found == true) {
        // shift every element after index to the left by one
        for (uint i = foundCpIndex; i < startingKeyIndex + size - 1; i++) {
            checkpoints[i] = checkpoints[i + 1];
        }
        // remove last element
        delete checkpoints[startingKeyIndex + size - 1];
        size--;
      }
      return found;
    }


    // return checkpoinst  [startingKeyIndex, startingKeyIndex+size) from map
    function listCheckpoints() external view returns(ICheckpoint[] memory res) {
        res = new ICheckpoint[](size);
        for (uint i = startingKeyIndex; i < startingKeyIndex + size; i++) {
            res[i - startingKeyIndex] = checkpoints[i];
        }
    }

}
