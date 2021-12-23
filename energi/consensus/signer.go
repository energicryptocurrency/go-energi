// Copyright 2019 The Energi Core Authors
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

package consensus

import (
	"errors"
	"math/big"

	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
)

// NOTE: it MUST NOT for untrusted transactions
type ConsensusSigner struct{}

func NewConsensusSigner() *ConsensusSigner {
	return &ConsensusSigner{}
}

func (cs ConsensusSigner) Equal(s2 types.Signer) bool {
	_, ok := s2.(ConsensusSigner)
	return ok
}
func (cs ConsensusSigner) SignatureValues(tx *types.Transaction, sig []byte) (r, s, v *big.Int, err error) {
	return nil, nil, nil, errors.New("Not Supported")
}
func (cs ConsensusSigner) Hash(tx *types.Transaction) common.Hash {
	return common.Hash{}
}
func (cs ConsensusSigner) Sender(tx *types.Transaction) (common.Address, error) {
	if !tx.IsConsensus() {
		return common.Address{}, errors.New("Not Consensus Tx")
	}
	return tx.ConsensusSender(), nil
}
