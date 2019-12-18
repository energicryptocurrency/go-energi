// Copyright 2019 The Energi Core Authors
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
pragma solidity 0.5.11;
//pragma experimental SMTChecker;
pragma experimental ABIEncoderV2;

import { IGovernedContract } from "./GovernedContract.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { ICheckpoint } from "./ICheckpoint.sol";

import {
    StorageCheckpointRegistryV1,
    CheckpointV1,
    CheckpointRegistryV1
} from "./CheckpointRegistryV1.sol";

// solium-disable-next-line no-empty-blocks

/**
 * Checkpoint V2 object
 */
contract CheckpointV2 is CheckpointV1 {
    constructor(
        IGovernedProxy _mnregistry_proxy,
        uint _number,
        bytes32 _hash,
        bytes32 _sigbase,
        bytes memory _cpp_sig
    )
        public
        CheckpointV1(_mnregistry_proxy, _number, _hash, _sigbase)
    {
        require(_cpp_sig.length == 65, "Invalid signature length");
        (bytes32 r, bytes32 s) = abi.decode(_cpp_sig, (bytes32, bytes32));
        address signer = ecrecover(_sigbase, uint8(_cpp_sig[64]), r, s);

        signature_list.push(_cpp_sig);
        signers[signer] = signature_list.length;
    }

    function canVote(address masternode) external view returns(bool) {
        if (signers[masternode] != 0) {
            return false;
        }

        if ((block.number - since) >= SIGNING_PERIOD) {
            return false;
        }

        return true;
    }
}

/**
 * Genesis hardcoded version of CheckpointRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract CheckpointRegistryV2 is CheckpointRegistryV1 {
    constructor(address _proxy, IGovernedProxy _mnregistry_proxy, address _cpp_signer)
        public CheckpointRegistryV1(_proxy, _mnregistry_proxy, _cpp_signer)
    {}

    // IGovernedContract
    //---------------------------------
    function _migrate(IGovernedContract _oldImpl) internal {
        v1storage.kill();
        v1storage = CheckpointRegistryV1(address(_oldImpl)).v1storage();
    }

    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // ICheckpointRegistry
    //---------------------------------
    function propose(uint number, bytes32 hash, bytes calldata signature) external returns(ICheckpoint checkpoint) {
        bytes32 sigbase = signatureBase(number, hash);
        require(signature.length == 65, "Invalid signature length");
        (bytes32 r, bytes32 s) = abi.decode(signature, (bytes32, bytes32));
        require(ecrecover(sigbase, uint8(signature[64]), r, s) == CPP_signer, "Invalid signer");

        checkpoint = new CheckpointV2(mnregistry_proxy, number, hash, sigbase, signature);
        v1storage.add(checkpoint);

        emit Checkpoint(
            number,
            hash,
            checkpoint
        );
    }
}
