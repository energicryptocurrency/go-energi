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
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IMasternodeToken } from "./IMasternodeToken.sol";
import { IMasternodeRegistry } from "./IMasternodeRegistry.sol";
import { NonReentrant } from "./NonReentrant.sol";
import { StorageBase }  from "./StorageBase.sol";

/**
 * Permanent storage of Masternode Token V1 data.
 */
contract StorageMasternodeTokenV1 is
    StorageBase
{
    struct Balance {
        uint256 amount;
        uint256 last_block;
    }
    mapping(address => Balance) public balances;

    // NOTE: ABIEncoderV2 is not acceptable at the moment of development!
    function balanceOnly(address _account)
        external view
        returns(uint256 amount)
    {
        return balances[_account].amount;
    }

    function setBalance(address _account, uint256 _amount, uint256 _last_block)
        external
        requireOwner
    {
        // NOTE: DO NOT process last_block as part of storage logic!
        Balance storage item = balances[_account];
        item.amount = _amount;
        item.last_block = _last_block;
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
    IMasternodeToken,
    NonReentrant
{
    // Data for migration
    //---------------------------------
    StorageMasternodeTokenV1 public v1storage;
    IGovernedProxy public registry_proxy;
    //---------------------------------

    constructor(address _proxy, IGovernedProxy _registry_proxy)
        public
        GovernedContract(_proxy)
    {
        v1storage = new StorageMasternodeTokenV1();
        registry_proxy = _registry_proxy;

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
        returns (uint256 balance, uint256 last_block)
    {
        (balance, last_block) = v1storage.balances(_tokenOwner);
    }

    function withdrawCollateral(uint256 _amount) external noReentry {
        // Retrieve
        address payable tokenOwner = _callerAddress();
        uint256 balance = v1storage.balanceOnly(tokenOwner);

        // Process
        if (balance < _amount) {
            revert("Not enough");
        }

        balance -= _amount;
        _validateBalance(balance);

        // Store
        v1storage.setBalance(tokenOwner, balance, block.number);

        // Events
        emit Transfer(tokenOwner, address(0), _amount);

        // Notify the registry
        IMasternodeRegistry(address(registry_proxy.impl())).onCollateralUpdate(tokenOwner);

        // TODO: we may need to allow more gas here for shared masternode contracts!
        tokenOwner.transfer(_amount);
    }

    function depositCollateral() external payable noReentry {
        // Retrieve
        address payable tokenOwner = _callerAddress();
        uint256 balance = v1storage.balanceOnly(tokenOwner);

        // Process
        balance += msg.value;
        _validateBalance(balance);

        // Store
        v1storage.setBalance(tokenOwner, balance, block.number);

        // Events
        emit Transfer(address(0), tokenOwner, msg.value);

        // Notify the registry
        IMasternodeRegistry(address(registry_proxy.impl())).onCollateralUpdate(tokenOwner);
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

    // Safety
    //---
    function () external payable {
        revert("Not supported");
    }
}
