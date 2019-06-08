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
import { IMasternodeToken } from "./IMasternodeToken.sol";
import { StorageBase }  from "./StorageBase.sol";

/**
 * Permanent storage of Masternode Token V1 data.
 */
contract StorageMasternodeTokenV1 is
    StorageBase
{
    struct Balance {
        uint256 amount;
        uint256 last_change;
    }
    mapping(address => Balance) public balances;

    // NOTE: ABIEncoderV2 is not acceptable at the moment of development!
    function balanceOnly(address _account)
        external view
        returns(uint256 amount)
    {
        return balances[_account].amount;
    }

    function setBalance(address _account, uint256 _amount, uint256 _last_change)
        external
        requireOwner
    {
        // NOTE: DO NOT process last_change as part of storage logic!
        Balance storage item = balances[_account];
        item.amount = _amount;
        item.last_change = _last_change;
    }
}

/**
 * MN-1: Genesis hardcoded version of MasternodeToken.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract MasternodeTokenV1 is
    GlobalConstants,
    GovernedContract,
    IMasternodeToken
{
    // Data for migration
    //---------------------------------
    StorageMasternodeTokenV1 public v1storage;
    //---------------------------------

    constructor(address _proxy) public GovernedContract(_proxy) {
        v1storage = new StorageMasternodeTokenV1();

        // ERC20
        emit Transfer(address(0), address(0), 0);
    }

    // IGovernedContract
    //---------------------------------
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
    }

    // MasternodeTokenV1
    //---------------------------------

    function totalSupply() external view returns (uint256) {
        return address(this).balance;
    }

    function name() external view returns (string memory) {
        return "Masternode Collateral";
    }

    function symbol() external view returns (string memory) {
        return "MNGR";
    }

    function decimals() external view returns (uint8) {
        return 18 + 4;
    }

    function balanceOf(address _owner) external view returns (uint256) {
        return v1storage.balanceOnly(_owner);
    }

    function transfer(address, uint256) external returns (bool) {
        revert("Not allowed");
    }

    function transferFrom(address, address, uint256) external returns (bool) {
        revert("Not allowed");
    }

    function approve(address, uint256) external returns (bool) {
        revert("Not allowed");
    }

    function allowance(address, address) external view returns (uint256) {
        return 0;
    }

    // Energi
    //---------------------------------

    // solium-disable security/no-block-members

    function balanceInfo(address _tokenOwner)
        external view
        returns (uint256 balance, uint256 age)
    {
        (balance, age) = v1storage.balances(_tokenOwner);

        assert(block.timestamp >= age);

        age = block.timestamp - age;
    }

    function withdrawCollateral(uint256 _amount) external {
        // Retrieve
        address payable tokenOwner = _ownerAddress();
        uint256 balance = v1storage.balanceOnly(tokenOwner);

        // Process
        if (balance < _amount) {
            revert("Not enough");
        }

        balance -= _amount;
        _validateBalance(balance);

        // Store
        v1storage.setBalance(tokenOwner, balance, block.timestamp);

        // Events
        emit Transfer(tokenOwner, address(0), _amount);

        // TODO: we may need to allow more gas here for shared masternode contracts!
        tokenOwner.transfer(_amount);
    }

    function depositCollateral() external payable {
        // Retrieve
        address payable tokenOwner = _ownerAddress();
        uint256 balance = v1storage.balanceOnly(tokenOwner);

        // Process
        balance += msg.value;
        _validateBalance(balance);

        // Store
        v1storage.setBalance(tokenOwner, balance, block.timestamp);

        // Events
        emit Transfer(address(0), tokenOwner, msg.value);
    }

    function _validateBalance(uint256 _amount) internal pure {
        // NOTE: "Too small" check makes no sense as it would be just zero.

        if (_amount > MN_COLLATERAL_MAX) {
            revert("Too much");
        }

        if ((_amount % MN_COLLATERAL_MIN) != 0) {
            revert("Not a multiple");
        }
    }

    function _ownerAddress()
        internal view
        returns (address payable)
    {
        if (msg.sender == proxy) {
            // This is guarantee of the GovernedProxy
            // solium-disable-next-line security/no-tx-origin
            return tx.origin;
        } else {
            return msg.sender;
        }
    }

    // Safety
    //---
    function () external payable {
        revert("Not supported");
    }
}
