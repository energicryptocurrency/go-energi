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

import { GlobalConstants } from "./constants.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IBlacklistRegistry, IBlacklistProposal, IProposal } from "./IBlacklistRegistry.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { ITreasury } from "./ITreasury.sol";
import { Gen2Migration } from "./Gen2Migration.sol";
import { GenericProposalV1 } from "./GenericProposalV1.sol";
import { StorageBase }  from "./StorageBase.sol";
import { NonReentrant } from "./NonReentrant.sol";

contract BlacklistProposalV1 is
    GenericProposalV1,
    IBlacklistProposal
{
    constructor(IGovernedProxy _mnregistry_proxy, address payable fee_payer)
        public
        GenericProposalV1(
            _mnregistry_proxy,
            10,
            1 weeks,
            fee_payer
        )
    // solium-disable-next-line no-empty-blocks
    {}

    function isObeyed()
        external view
        returns(bool)
    {
        if (isAccepted()) {
            return true;
        }

        uint accepted = accepted_weight;
        uint rejected = rejected_weight;

        if ((accepted > (rejected*2)) && (accepted > MN_COLLATERAL_MAX)) {
            return true;
        }

        return false;
    }
}

/**
 * Permanent storage of Blacklist Registry V1 data.
 */
contract StorageBlacklistRegistryV1 is
    StorageBase
{
    // NOTE: ABIEncoderV2 is not acceptable at the moment of development!

    struct Info {
        IProposal enforce;
        IProposal revoke;
        uint index;
    }

    mapping(address => Info) public address_info;
    address[] public address_list;

    function setEnforce(address addr, IProposal proposal)
        external
        requireOwner
    {
        Info storage item = address_info[addr];
        assert(address(item.enforce) == address(0));

        item.enforce = proposal;
        item.index = address_list.length;
        address_list.push(addr);
    }

    function setRevoke(address addr, IProposal proposal)
        external
        requireOwner
    {
        Info storage item = address_info[addr];

        assert(address(item.enforce) != address(0));

        item.revoke = proposal;
    }


    function remove(address addr)
        external
        requireOwner
    {
        Info storage item = address_info[addr];
        assert(address(item.enforce) != address(0));

        // Move the last into the gap, NOOP on on match
        address_list[item.index] = address_list[address_list.length - 1];
        address_list.pop();

        delete address_info[addr];
    }

    function addresses()
        external view
        returns(address[] memory result)
    {
        uint len = address_list.length;
        result = new address[](len);

        for (uint i = 0; i < len; ++i) {
            result[i] = address_list[i];
        }
    }
}


/**
 * Genesis hardcoded version of BlacklistRegistry.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract BlacklistRegistryV1 is
    GovernedContract,
    NonReentrant,
    GlobalConstants,
    IBlacklistRegistry
{
    // Data for migration
    //---------------------------------
    StorageBlacklistRegistryV1 public v1storage;
    IGovernedProxy public mnregistry_proxy;
    Gen2Migration public migration;
    ITreasury public compensation_fund;
    //---------------------------------

    constructor(
        address _proxy,
        IGovernedProxy _mnregistry_proxy,
        Gen2Migration _migration,
        ITreasury _compensatin_fund
    )
        public GovernedContract(_proxy) {
        v1storage = new StorageBlacklistRegistryV1();
        mnregistry_proxy = _mnregistry_proxy;
        migration = _migration;
        compensation_fund = _compensatin_fund;
    }

    // IGovernedContract
    //---------------------------------
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // IBlacklistRegistry
    //---------------------------------
    function proposals(address addr)
        external view
        returns(IProposal enforce, IProposal revoke)
    {
        (enforce, revoke,) = v1storage.address_info(addr);
    }

    function propose(address addr)
        external payable
        noReentry
        returns(address)
    {
        require(msg.value == FEE_BLACKLIST_V1, "Invalid fee");

        StorageBlacklistRegistryV1 store = v1storage;
        (IProposal enforce, IProposal revoke,) = store.address_info(addr);

        // Cleanup old
        if (address(enforce) != address(0)) {
            if (address(revoke) != address(0)) {
                // assume enforced
                if (revoke.isAccepted()) {
                    enforce.destroy();
                    revoke.destroy();
                    store.remove(addr);
                } else if (revoke.isFinished()) {
                    revert("Already active (1)");
                }
            } else if (enforce.isFinished() && !enforce.isAccepted()) {
                enforce.collect();
                store.remove(addr);
            } else {
                revert("Already active (2)");
            }
        }

        // Create new
        BlacklistProposalV1 proposal = new BlacklistProposalV1(
            mnregistry_proxy,
            _callerAddress()
        );

        proposal.setFee.value(msg.value)();

        store.setEnforce(addr, proposal);

        emit BlacklistProposal(addr, IProposal(address(proposal)));
    }

    function revokeProposal(address addr)
        external payable
        noReentry
        returns(address)
    {
        require(msg.value == FEE_BLACKLIST_REVOKE_V1, "Invalid fee");

        StorageBlacklistRegistryV1 store = v1storage;
        (IProposal enforce, IProposal revoke,) = store.address_info(addr);

        // Cleanup old
        require(address(enforce) != address(0), "No need (1)");

        if (address(revoke) != address(0)) {
            // assume enforced
            if (!revoke.isFinished()) {
                revert("Already active");
            } else if (!revoke.isAccepted()) {
                revoke.collect();
            }
        } else if (!enforce.isFinished()) {
            revert("Not applicable");
        } else if (!enforce.isAccepted()) {
            revert("No need (2)");
        }

        // Create new
        BlacklistProposalV1 proposal = new BlacklistProposalV1(
            mnregistry_proxy,
            _callerAddress()
        );

        proposal.setFee.value(msg.value)();

        store.setRevoke(addr, proposal);

        emit WhitelistProposal(addr, IProposal(address(proposal)));
    }

    function isBlacklisted(address addr)
        public view
        returns(bool)
    {
        StorageBlacklistRegistryV1 store = v1storage;
        (IProposal enforce, IProposal revoke,) = store.address_info(addr);

        if ((address(revoke) != address(0)) && revoke.isAccepted()) {
            return false;
        }

        if ((address(enforce) != address(0)) && IBlacklistProposal(address(enforce)).isObeyed()) {
            return true;
        }

        return false;
    }

    function collect(address addr)
        external
        noReentry
    {
        StorageBlacklistRegistryV1 store = v1storage;
        (IProposal enforce, IProposal revoke,) = store.address_info(addr);

        require (address(enforce) != address(0), "Nothing to collect (1)");
        require (enforce.isFinished(), "Nothing to collect (2)");

        if (!enforce.isAccepted()) {
            enforce.collect();
            store.remove(addr);
            return;
        }

        require (address(revoke) != address(0), "Nothing to collect (3)");
        require(revoke.isFinished(), "Nothing to collect (4)");

        if (revoke.isAccepted()) {
            enforce.destroy();
            revoke.destroy();
            store.remove(addr);
            return;
        }

        revoke.collect();
        store.setRevoke(addr, IProposal(address(0)));
    }

    function collectMigration(uint item_id, bytes20 owner)
        external
        noReentry
    {
        require(isBlacklisted(address(owner)), "Not blacklisted");
        // TODO: check if valid to collect
        migration.blacklistClaim(item_id, owner);
    }

    function enumerateAll()
        external view
        returns(address[] memory addresses)
    {
        return v1storage.addresses();
    }

    function enumerateBlocked()
        external view
        returns(address[] memory addresses)
    {
        addresses = v1storage.addresses();

        for (uint i = addresses.length; i-- > 0;) {
            if (!isBlacklisted(addresses[i])) {
                addresses[i] = address(0);
            }
        }
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
