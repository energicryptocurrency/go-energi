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

import { GlobalConstants } from "./constants.sol";
import { IDelegatedPoS } from "./IDelegatedPoS.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IBlacklistRegistry } from "./IBlacklistRegistry.sol";

/**
 * Genesis hardcoded version of Gen 2 Migration
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract Gen2Migration is
	GlobalConstants,
	IDelegatedPoS
{
    struct UnspentCoins {
        bytes20 owner; // Gen 2 P2PKH
        uint amount;
    }

    event Migrated(
        uint item_id,
        address destination,
        uint amount
    );

    IGovernedProxy public blacklist_proxy;
    uint public chain_id;
    address public signerAddress; // IDelegatedPoS
    uint public totalAmount;
    UnspentCoins[] public coins;
    mapping(bytes20 => bool) hard_blacklist;

    // NOTE: this c-tor is used during testing
    constructor(IGovernedProxy _blacklist_proxy, uint _chain_id, address _signer) public {
        blacklist_proxy = _blacklist_proxy;
        chain_id = _chain_id;
        signerAddress = _signer;
    }

    function setSnapshot(
        bytes20[] calldata _owners,
        uint[] calldata _amounts,
        bytes20[] calldata _blacklist
    ) external {
        require(coins.length == 0, "Already set");
        require(msg.sender == signerAddress, "Invalid sender");
        require(_owners.length == _amounts.length, "match length");
        require(_owners.length > 0, "has data");

        coins.length = _owners.length;
        uint total;

        for (uint i = _owners.length; i-- > 0;) {
            coins[i].owner = _owners[i];
            coins[i].amount = _amounts[i];
            total += _amounts[i];
        }

        totalAmount = total;
        // NOTE: there is a special consensus procedure to setup account balance based on
        //       totalAmount().

        for (uint i = _blacklist.length; i-- > 0;) {
            hard_blacklist[_blacklist[i]] = true;
        }
    }

    function itemCount() external view returns(uint) {
        return coins.length;
    }

    function hashToSign(address payable _destination)
        public view
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _destination,
                "||Energi Gen 2 migration claim||",
                chain_id
            )
        );
    }

    function verifyClaim(uint _item_id, address payable _destination, uint8 sig_v, bytes32 sig_r, bytes32 sig_s)
        public view
        returns(uint amount)
    {
        // Check ID
        require(_item_id < coins.length, "Invalid ID");

        // Recover owner
        bytes32 hash = hashToSign(_destination);
        bytes20 owner = bytes20(ecrecover(hash, sig_v, sig_r, sig_s));

        // Validate Owner
        require(coins[_item_id].owner == owner, "Invalid signature");

        // Check if blacklisted
        IBlacklistRegistry blacklist = IBlacklistRegistry(address(blacklist_proxy.impl()));
        require(!blacklist.isBlacklisted(address(owner)), "Owner is blacklisted");
        require(!hard_blacklist[owner], "Owner is hard blacklisted");

        // Validate amount
        amount = coins[_item_id].amount;
    }

    function claim(uint _item_id, address payable _destination, uint8 sig_v, bytes32 sig_r, bytes32 sig_s)
        external
    {
        uint amount = verifyClaim(_item_id, _destination, sig_v, sig_r, sig_s);
        require(amount > 0, "Already spent");

        // Spend
        coins[_item_id].amount = 0;

        emit Migrated(
            _item_id,
            _destination,
            amount
        );

        _destination.transfer(amount);
    }

    // SECURITY: emergency drain procedure has to be implemented as blacklist
    //           followed by consensus-level drain to Blacklist registry.

    function blacklistClaim(uint _item_id, bytes20 _owner) external {
        require(_item_id < coins.length, "Invalid ID");

        uint amount = coins[_item_id].amount;
        require(amount > 0, "Already spent");

        require(coins[_item_id].owner == _owner, "Invalid Owner");

        IBlacklistRegistry blacklist = IBlacklistRegistry(address(blacklist_proxy.impl()));
        require(msg.sender == address(blacklist), "Not blacklist registry");

        // Spend
        coins[_item_id].amount = 0;

        blacklist.compensation_fund().contribute.value(amount)();
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
