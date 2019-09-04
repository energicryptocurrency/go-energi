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

import { IProposal } from "./IProposal.sol";
import { ITreasury } from "./ITreasury.sol";

/**
 * Interface of BlacklistProposal
 */
contract IBlacklistProposal is IProposal {
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
        IBlacklistProposal proposal
    );
    event WhitelistProposal(
        address indexed target,
        IBlacklistProposal proposal
    );
    event DrainProposal(
        address indexed target,
        IBlacklistProposal proposal
    );

    function compensation_fund() external view returns(ITreasury);
    function EBI_signer() external view returns(address);

    function proposals(address) external view returns(
        IBlacklistProposal enforce,
        IBlacklistProposal revoke,
        IBlacklistProposal drain);
    function propose(address) external payable returns(IBlacklistProposal);
    function proposeRevoke(address) external payable returns(IBlacklistProposal);
    function proposeDrain(address) external payable returns(IBlacklistProposal);
    function isBlacklisted(address) external view returns(bool);
    function isDrainable(address) external view returns(bool);
    function collect(address) external;
    function drainMigration(uint item_id, bytes20 owner) external;
    function enumerateAll() external view returns(address[] memory addresses);
    function enumerateBlocked() external view returns(address[] memory addresses);
    function enumerateDrainable() external view returns(address[] memory addresses);
    function onDrain(address) external;
}
