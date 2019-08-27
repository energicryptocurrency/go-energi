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
pragma solidity 0.5.10;
//pragma experimental SMTChecker;
pragma experimental ABIEncoderV2;

/**
 * Genesis version of Checkpoint interface.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
interface ICheckpoint {
    function info() external view returns(uint number, bytes32 hash, uint since);
    function sign(bytes calldata signature) external;
    function signatures() external view returns(bytes[] memory siglist);
    function signature(address masternode) external view returns(bytes memory);
    function signatureBase() external view returns(bytes32 sigbase);
}
