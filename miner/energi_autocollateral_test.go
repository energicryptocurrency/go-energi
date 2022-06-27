// Copyright 2020 The Energi Core Authors
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

package miner

import (
	"testing"

	"github.com/energicryptocurrency/energi/consensus"
	energi "github.com/energicryptocurrency/energi/energi/consensus"
	"github.com/energicryptocurrency/energi/ethdb"
	"github.com/energicryptocurrency/energi/params"
)

func TestCanAutocollateralizeEnergi(t *testing.T) {
	testCanAutocollateralize(t, energiChainConfig, energi.New(energiChainConfig.Energi, ethdb.NewMemDatabase()))
}

func testCanAutocollateralize(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, 0)
	defer func() {
		w.close()
		_ = b.CleanUp()
	}()
}

func TestDoAutocollateralEnergi(t *testing.T) {
	testDoAutocollateral(t, energiChainConfig, energi.New(energiChainConfig.Energi, ethdb.NewMemDatabase()))
}

func testDoAutocollateral(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, 0)
	defer func() {
		w.close()
		_ = b.CleanUp()
	}()
}