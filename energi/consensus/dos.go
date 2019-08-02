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
	"time"

	eth_consensus "github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"

	energi_params "energi.world/core/gen3/energi/params"
)

const (
	oldForkPeriod = time.Duration(15) * time.Minute
)

func (e *Energi) checkDoS(
	chain ChainReader,
	header *types.Header,
	parent *types.Header,
) error {
	old_fork_threshold := e.now() - energi_params.OldForkPeriod

	// POS-8: allow old fork only if current head is not fresh enough
	if parent.Time < old_fork_threshold {
		current := chain.CurrentHeader()

		if current.Time > old_fork_threshold {
			return eth_consensus.ErrDoSThrottle
		}
	}

	return nil
}
