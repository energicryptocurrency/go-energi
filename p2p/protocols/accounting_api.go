// Copyright 2018 The Energi Core Authors
// Copyright 2018 The go-ethereum Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

package protocols

import (
	"errors"
)

// Textual version number of accounting API
const AccountingVersion = "1.0"

var errNoAccountingMetrics = errors.New("accounting metrics not enabled")

// AccountingApi provides an API to access account related information
type AccountingApi struct {
	metrics *AccountingMetrics
}

// NewAccountingApi creates a new AccountingApi
// m will be used to check if accounting metrics are enabled
func NewAccountingApi(m *AccountingMetrics) *AccountingApi {
	return &AccountingApi{m}
}

// Balance returns local node balance (units credited - units debited)
func (self *AccountingApi) Balance() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	balance := mBalanceCredit.Count() - mBalanceDebit.Count()
	return balance, nil
}

// BalanceCredit returns total amount of units credited by local node
func (self *AccountingApi) BalanceCredit() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mBalanceCredit.Count(), nil
}

// BalanceCredit returns total amount of units debited by local node
func (self *AccountingApi) BalanceDebit() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mBalanceDebit.Count(), nil
}

// BytesCredit returns total amount of bytes credited by local node
func (self *AccountingApi) BytesCredit() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mBytesCredit.Count(), nil
}

// BalanceCredit returns total amount of bytes debited by local node
func (self *AccountingApi) BytesDebit() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mBytesDebit.Count(), nil
}

// MsgCredit returns total amount of messages credited by local node
func (self *AccountingApi) MsgCredit() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mMsgCredit.Count(), nil
}

// MsgDebit returns total amount of messages debited by local node
func (self *AccountingApi) MsgDebit() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mMsgDebit.Count(), nil
}

// PeerDrops returns number of times when local node had to drop remote peers
func (self *AccountingApi) PeerDrops() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mPeerDrops.Count(), nil
}

// SelfDrops returns number of times when local node was overdrafted and dropped
func (self *AccountingApi) SelfDrops() (int64, error) {
	if self.metrics == nil {
		return 0, errNoAccountingMetrics
	}
	return mSelfDrops.Count(), nil
}
