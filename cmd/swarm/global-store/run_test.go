// Copyright 2019 The Energi Core Authors
// Copyright 2018 The go-ethereum Authors
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

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/docker/pkg/reexec"
	"energi.world/core/gen3/internal/cmdtest"
)

func init() {
	reexec.Register("swarm-global-store", func() {
		if err := newApp().Run(os.Args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	})
}

func runGlobalStore(t *testing.T, args ...string) *cmdtest.TestCmd {
	tt := cmdtest.NewTestCmd(t, nil)
	tt.Run("swarm-global-store", args...)
	return tt
}

func TestMain(m *testing.M) {
	if reexec.Init() {
		return
	}
	os.Exit(m.Run())
}
