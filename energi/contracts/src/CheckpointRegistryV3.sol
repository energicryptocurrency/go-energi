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

import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { ICheckpoint } from "./ICheckpoint.sol";
import { ICheckpointRegistryV2 } from "./ICheckpointRegistryV2.sol";
import { StorageCheckpointRegistryV2 } from "./StorageCheckpointRegistryV2.sol";
import { CheckpointV2 } from "./CheckpointRegistryV2.sol";

// solium-disable-next-line no-empty-blocks


/**
 * Genesis hardcoded version of CheckpointRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract CheckpointRegistryV3 is GovernedContract, ICheckpointRegistryV2  {

      // Data for migration
      //---------------------------------
      // main storage contract (for current version queue style) that stores checkpoints
      StorageCheckpointRegistryV2 public v2storage;
      // Igoverned proxy registry that is checked to be active when creating new checkpoints
      IGovernedProxy public mnregistry_proxy;
      // address that is expected to be making signatures for propose or removal of signatures
      address public CPP_signer;


      constructor(address _proxy, IGovernedProxy _mnregistry_proxy, address _cpp_signer) public GovernedContract(_proxy) {
          v2storage = new StorageCheckpointRegistryV2();
          mnregistry_proxy = _mnregistry_proxy;
          CPP_signer = _cpp_signer;
      }

      // IGovernedContract
      //---------------------------------
      function _destroy(IGovernedContract _newImpl) internal {
          v2storage.setOwner(_newImpl);
      }

      // ICheckpointRegistry
      //---------------------------------
      function signatureBase(uint number, bytes32 hash) public view returns(bytes32 sigbase) {
          sigbase = keccak256(
              abi.encodePacked(
                  "||Energi Blockchain Checkpoint||",
                  number,
                  hash
              )
          );
      }

      // ICheckpointRegistry
      //---------------------------------
      function propose(uint number, bytes32 hash, bytes calldata signature) external returns(ICheckpoint checkpoint) {
          bytes32 sigbase = signatureBase(number, hash);
          require(signature.length == 65, "Invalid signature length");
          (bytes32 r, bytes32 s) = abi.decode(signature, (bytes32, bytes32));
          require(ecrecover(sigbase, uint8(signature[64]), r, s) == CPP_signer, "Invalid signer");

          checkpoint = new CheckpointV2(mnregistry_proxy, number, hash, sigbase, signature);
          v2storage.add(checkpoint);

          emit Checkpoint(
              number,
              hash,
              checkpoint
          );
      }

      // Remove checkpoint from storage (always succeeds)
      function remove(uint number, bytes32 hash, bytes calldata signature) external returns(bool deleted) {
          // Allow to remove checkpoint by any caller as far as signature is correct.
          // require(_callerAddress() == CPP_signer, "Invalid caller");

          // validation
          bytes32 sigbase = signatureBase(number, hash);
          require(signature.length == 65, "Invalid signature length");
          (bytes32 r, bytes32 s) = abi.decode(signature, (bytes32, bytes32));
          require(ecrecover(sigbase, uint8(signature[64]), r, s) == CPP_signer, "Invalid signer");

          // remove checkpoint from storage
          deleted = v2storage.remove(new CheckpointV2(mnregistry_proxy, number, hash, sigbase, signature));
      }

      function checkpoints() external view returns(ICheckpoint[] memory) {
          return v2storage.listCheckpoints();
      }

      function sign(ICheckpoint checkpoint, bytes calldata signature) external {
          checkpoint.sign(signature);
      }




      // Safety
      //---------------------------------
      function () external payable {
          revert("Not supported");
      }
}
