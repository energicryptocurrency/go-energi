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
//pragma experimental SMTChecker;
pragma experimental ABIEncoderV2;

import { ICheckpoint } from "./ICheckpoint.sol";

/**
 * Genesis version of CheckpointRegistry interface.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
interface ICheckpointRegistry {
    event Checkpoint(
        uint indexed number,
        bytes32 hash,
        ICheckpoint checkpoint
    );

    function CPP_signer() external view returns(address);

    function propose(uint number, bytes32 hash, bytes calldata signature) external returns(ICheckpoint);
    function checkpoints() external view returns(ICheckpoint[] memory);
    function signatureBase(uint number, bytes32 hash) external view returns(bytes32 sigbase);
    function sign(ICheckpoint checkpoint, bytes calldata signature) external;
}
