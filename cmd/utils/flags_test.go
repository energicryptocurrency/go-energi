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

// Package utils contains internal helper functions for go-ethereum commands.
package utils

import (
	"flag"
	"reflect"
	"strconv"
	"testing"

	"github.com/energicryptocurrency/energi/p2p"

	cli "gopkg.in/urfave/cli.v1"
)

func TestTestnetPort(t *testing.T) {
	// create flags
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("testnet", "1", "doc")
	globalSet.Int64("port", 0, "doc")

	// set flags
	globalCtx := cli.NewContext(nil, globalSet, nil)
	_ = globalCtx.Set("testnet", "23")

	cfg := &p2p.Config{}

	// check port remains 49797
	_ = globalCtx.Set("port", "49797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(testnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(testnetDefaultPort))
	}

	// check port 39797 isn't allowed and port is default
	_ = globalCtx.Set("port", "39797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(testnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(testnetDefaultPort))
	}

	// check port 59797 isn't allowed and port is default
	_ = globalCtx.Set("port", "59797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(testnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(testnetDefaultPort))
	}

	// check any other port can be set
	_ = globalCtx.Set("port", "123")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(123) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(testnetDefaultPort))
	}
}

func TestSimnetPort(t *testing.T) {
	// create flags
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("simnet", "1", "doc")
	globalSet.Int64("port", 0, "doc")

	// set flags
	globalCtx := cli.NewContext(nil, globalSet, nil)
	_ = globalCtx.Set("simnet", "23")

	cfg := &p2p.Config{}

	// check port remains 49797
	_ = globalCtx.Set("port", "49797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(simnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(simnetDefaultPort))
	}

	// check port 39797 isn't allowed and port is default
	_ = globalCtx.Set("port", "39797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(simnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(simnetDefaultPort))
	}

	// check port 59797 isn't allowed and port is default
	_ = globalCtx.Set("port", "59797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(simnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(simnetDefaultPort))
	}

	// check any other port can be set
	_ = globalCtx.Set("port", "123")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(123) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(123))
	}
}

func TestMainnetPort(t *testing.T) {
	// create flags
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("mainnet", "1", "doc")
	globalSet.Int64("port", 0, "doc")

	// set flags
	globalCtx := cli.NewContext(nil, globalSet, nil)
	_ = globalCtx.Set("mainnet", "23")

	cfg := &p2p.Config{}

	// check port remains 49797
	_ = globalCtx.Set("port", "49797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(mainnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(mainnetDefaultPort))
	}

	// check port 39797 isn't allowed and port is default
	_ = globalCtx.Set("port", "59797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(mainnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(mainnetDefaultPort))
	}

	// check port 59797 isn't allowed and port is default
	_ = globalCtx.Set("port", "39797")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(mainnetDefaultPort) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(mainnetDefaultPort))
	}

	// check any other port can be set
	_ = globalCtx.Set("port", "123")
	setListenAddress(globalCtx, cfg)
	if cfg.ListenAddr[1:] != strconv.Itoa(123) {
		t.Errorf("got %s, want %s", cfg.ListenAddr, strconv.Itoa(123))
	}
}

func Test_SplitTagsFlag(t *testing.T) {
	tests := []struct {
		name string
		args string
		want map[string]string
	}{
		{
			"2 tags case",
			"host=localhost,bzzkey=123",
			map[string]string{
				"host":   "localhost",
				"bzzkey": "123",
			},
		},
		{
			"1 tag case",
			"host=localhost123",
			map[string]string{
				"host": "localhost123",
			},
		},
		{
			"empty case",
			"",
			map[string]string{},
		},
		{
			"garbage",
			"smth=smthelse=123",
			map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitTagsFlag(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitTagsFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
