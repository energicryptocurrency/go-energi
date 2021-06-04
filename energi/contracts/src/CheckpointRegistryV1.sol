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
pragma solidity 0.5.16;
//pragma experimental SMTChecker;
pragma experimental ABIEncoderV2;

import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IMasternodeRegistry } from "./IMasternodeRegistry.sol";
import { ICheckpoint } from "./ICheckpoint.sol";
import { ICheckpointRegistry } from "./ICheckpointRegistry.sol";
import { StorageBase }  from "./StorageBase.sol";

/**
 * Permanent storage of Checkpoint Registry V1 data.
 */
// solium-disable-next-line no-empty-blocks
contract StorageCheckpointRegistryV1 is
    StorageBase
{
    // NOTE: ABIEncoderV2 is not acceptable at the moment of development!

    ICheckpoint[] public checkpoints;

    function add(ICheckpoint cp)
        external
        requireOwner
    {
        checkpoints.push(cp);
    }

    function listCheckpoints()
        external view
        returns(ICheckpoint[] memory res)
    {
        uint len = checkpoints.length;
        res = new ICheckpoint[](len);
        for (uint i = len; i-- > 0;) {
            res[i] = checkpoints[i];
        }
    }
}

/**
 * Checkpoint object
 */
contract CheckpointV1 is ICheckpoint {
    uint constant internal SIGNING_PERIOD = 24 * 60;

    IGovernedProxy internal mnregistry_proxy;
    uint internal since;

    uint internal number;
    bytes32 internal hash;
    bytes32 public signatureBase;
    mapping(address => uint) internal signers;

    bytes[] internal signature_list;

    constructor(IGovernedProxy _mnregistry_proxy, uint _number, bytes32 _hash, bytes32 _sigbase) public {
        mnregistry_proxy = _mnregistry_proxy;
        since = block.number;
        number = _number;
        hash = _hash;
        signatureBase = _sigbase;
    }

    function info() external view returns(uint, bytes32, uint) {
        return(number, hash, since);
    }

    function sign(bytes calldata signature) external {
        require((block.number - since) < SIGNING_PERIOD, "Signing has ended");

        require(signature.length == 65, "Invalid signature length");
        (bytes32 r, bytes32 s) = abi.decode(signature, (bytes32, bytes32));
        address masternode = ecrecover(signatureBase, uint8(signature[64]), r, s);

        require(signers[masternode] == 0, "Already signed");

        IMasternodeRegistry registry = IMasternodeRegistry(address(mnregistry_proxy.impl()));
        require(registry.isActive(masternode), "Not active MN");

        signature_list.push(signature);
        signers[masternode] = signature_list.length;
    }

    function signature(address masternode) external view returns(bytes memory){
        uint index = signers[masternode];
        require(index != 0, "Not signed yet");
        return signature_list[index - 1];
    }

    function signatures() external view returns(bytes[] memory siglist){
        uint len = signature_list.length;
        siglist = new bytes[](len);
        for (uint i = len; i-- > 0;) {
            siglist[i] = signature_list[i];
        }
    }
}

/**
 * Genesis hardcoded version of CheckpointRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract CheckpointRegistryV1 is
    GovernedContract,
    ICheckpointRegistry
{
    // Data for migration
    //---------------------------------
    StorageCheckpointRegistryV1 public v1storage;
    IGovernedProxy public mnregistry_proxy;
    address public CPP_signer;
    //---------------------------------

    constructor(address _proxy, IGovernedProxy _mnregistry_proxy, address _cpp_signer)
        public GovernedContract(_proxy)
    {
        v1storage = new StorageCheckpointRegistryV1();
        mnregistry_proxy = _mnregistry_proxy;
        CPP_signer = _cpp_signer;
    }

    // IGovernedContract
    //---------------------------------
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // ICheckpointRegistry
    //---------------------------------
    function signatureBase(uint number, bytes32 hash)
        public view
        returns(bytes32 sigbase)
    {
        sigbase = keccak256(
            abi.encodePacked(
                "||Energi Blockchain Checkpoint||",
                number,
                hash
            )
        );
    }

    function propose(uint number, bytes32 hash, bytes calldata signature) external returns(ICheckpoint checkpoint) {
        // Allow to propose by any caller as far as signature is correct.
        // This leaves us possibility of automatic checkpoint creation.
        //require(_callerAddress() == CPP_signer, "Invalid caller");

        bytes32 sigbase = signatureBase(number, hash);
        require(signature.length == 65, "Invalid signature length");
        (bytes32 r, bytes32 s) = abi.decode(signature, (bytes32, bytes32));
        require(ecrecover(sigbase, uint8(signature[64]), r, s) == CPP_signer, "Invalid signer");

        checkpoint = new CheckpointV1(mnregistry_proxy, number, hash, sigbase);
        v1storage.add(checkpoint);

        emit Checkpoint(
            number,
            hash,
            checkpoint
        );
    }

    function checkpoints() external view returns(ICheckpoint[] memory) {
        return v1storage.listCheckpoints();
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
