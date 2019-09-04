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

interface IProposal {
    function parent() external view returns(address);
    function created_block() external view returns(uint);
    function deadline() external view returns(uint);
    function fee_payer() external view returns(address payable);
    function fee_amount() external view returns(uint);
    function accepted_weight() external view returns(uint);
    function rejected_weight() external view returns(uint);
    function total_weight() external view returns(uint);
    function quorum_weight() external view returns(uint);
    function isFinished() external view returns(bool);
    function isAccepted() external view returns(bool);
    function withdraw() external;
    function destroy() external;
    function collect() external;
    function voteAccept() external;
    function voteReject() external;
    function setFee() external payable;
}

