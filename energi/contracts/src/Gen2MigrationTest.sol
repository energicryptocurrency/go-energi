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

import { Gen2Migration } from "./Gen2Migration.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { ITreasury } from "./ITreasury.sol";
import { BlacklistRegistryV1 } from "./BlacklistRegistryV1.sol";

contract MockGen2Migration is Gen2Migration {
    constructor(IGovernedProxy _blacklist_proxy, uint _chain_id, address _signer) public
        Gen2Migration(_blacklist_proxy, _chain_id, _signer)
    // solium-disable-next-line no-empty-blocks
    {
    }

    function setCoins(
        bytes20[] calldata _owners,
        uint[] calldata _amounts,
        bytes20[] calldata _blacklist
    ) external payable {
        coins.length = 0;
        address orig_signer = signerAddress;
        signerAddress = address(this);
        this.setSnapshot(_owners, _amounts, _blacklist);
        signerAddress = orig_signer;
    }
}

contract MockGen2MigrationBlacklist is BlacklistRegistryV1 {
    mapping(address => bool) is_blacklisted;

    constructor(address _proxy, IGovernedProxy _mnregistry_proxy, Gen2Migration _migration, ITreasury _fund)
        public
        BlacklistRegistryV1(_proxy, _mnregistry_proxy, _migration, _fund)
    // solium-disable-next-line no-empty-blocks
    {}

    function setBlacklisted(address addr, bool is_bl) external {
        is_blacklisted[addr] = is_bl;
    }

    function isBlacklisted(address addr) public view returns(bool) {
        return is_blacklisted[addr];
    }

    function drainMigration(uint item_id, bytes20 owner) external {
        migration.blacklistClaim(item_id, owner);
        _onDrain(address(owner));
    }
}
