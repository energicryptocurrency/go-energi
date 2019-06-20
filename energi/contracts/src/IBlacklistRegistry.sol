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
pragma solidity 0.5.9;
//pragma experimental SMTChecker;

import { IProposal } from "./IProposal.sol";

/**
 * Interface of BlacklistProposal
 */
interface IBlacklistProposal {
    function isObeyed() external view returns(bool);
}

/**
 * Genesis version of BlacklistRegistry interface.
 *
 * Base Consensus interface for blocking outgoing transactions from
 * blacklisted accounts.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
interface IBlacklistRegistry {
    event BlacklistProposal(
        address indexed target,
        IProposal proposal
    );
    event WhitelistProposal(
        address indexed target,
        IProposal proposal
    );

    function proposals(address) external view returns(IProposal enforce, IProposal revoke);
    function propose(address) external payable returns(address);
    function revokeProposal(address) external payable returns(address);
    function isBlacklisted(address) external view returns(bool);
    function collect(address) external;
}
