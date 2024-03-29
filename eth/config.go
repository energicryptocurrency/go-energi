// Copyright 2018 The Energi Core Authors
// Copyright 2017 The go-ethereum Authors
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

package eth

import (
	"math/big"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/energicryptocurrency/go-energi/common"
	"github.com/energicryptocurrency/go-energi/common/hexutil"
	"github.com/energicryptocurrency/go-energi/consensus/ethash"
	"github.com/energicryptocurrency/go-energi/core"
	"github.com/energicryptocurrency/go-energi/eth/downloader"
	"github.com/energicryptocurrency/go-energi/eth/gasprice"
	"github.com/energicryptocurrency/go-energi/params"
)

// DefaultConfig contains default settings for use on the Ethereum main net.
var DefaultConfig = Config{
	// TODO: special PoS-aware fast sync
	SyncMode: downloader.FullSync,
	Ethash: ethash.Config{
		CacheDir:       "ethash",
		CachesInMem:    2,
		CachesOnDisk:   3,
		DatasetsInMem:  1,
		DatasetsOnDisk: 2,
	},
	NetworkId:      39797,
	LightPeers:     100,
	DatabaseCache:  512,
	TrieCleanCache: 256,
	TrieDirtyCache: 256,
	TrieTimeout:    60 * time.Minute,
	TrieRapidTime:  10 * time.Second,
	MinerGasFloor:  40000000,
	MinerGasCeil:   80000000,
	MinerGasPrice:  big.NewInt(1 * params.GWei),
	MinerRecommit:  3 * time.Second,
	MinerNonceCap:  0,

	MinerAutoCompound: 1,

	TxPool: core.DefaultTxPoolConfig,
	GPO: gasprice.Config{
		Blocks:     20,
		Percentile: 60,
	},
}

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		if user, err := user.Current(); err == nil {
			home = user.HomeDir
		}
	}
	if runtime.GOOS == "windows" {
		DefaultConfig.Ethash.DatasetDir = filepath.Join(home, "AppData", "Ethash")
	} else {
		DefaultConfig.Ethash.DatasetDir = filepath.Join(home, ".ethash")
	}
}

//go:generate gencodec -type Config -field-override configMarshaling -formats toml -out gen_config.go

type DPoSMap map[common.Address]common.Address

type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Ethereum main net block is used.
	Genesis *core.Genesis `toml:",omitempty"`

	// Protocol options
	NetworkId uint64 // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode
	NoPruning bool

	// Whitelist of required block number -> hash values to accept
	Whitelist map[uint64]common.Hash `toml:"-"`

	// Light client options
	LightServ  int `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightPeers int `toml:",omitempty"` // Maximum number of LES client peers

	// Database options
	SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	TrieCleanCache     int
	TrieDirtyCache     int
	TrieTimeout        time.Duration
	TrieRapidTime      time.Duration

	// Mining-related options
	Etherbase      common.Address `toml:",omitempty"`
	MinerNotify    []string       `toml:",omitempty"`
	MinerExtraData []byte         `toml:",omitempty"`
	MinerGasFloor  uint64
	MinerGasCeil   uint64
	MinerGasPrice  *big.Int
	MinerRecommit  time.Duration
	MinerNoverify  bool

	MinerDPoS      DPoSMap `toml:",omitempty"`
	MinerMigration string  `toml:",omitempty"`
	MinerNonceCap  uint64  `toml:"-"`

	MinerAutoCompound uint64 `toml:",omitempty"`

	PublicService bool `toml:",omitempty"`

	// Ethash options
	Ethash ethash.Config

	// Transaction pool options
	TxPool core.TxPoolConfig

	// Gas Price Oracle options
	GPO gasprice.Config

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// Miscellaneous options
	DocRoot string `toml:"-"`

	// Type of the EWASM interpreter ("" for default)
	EWASMInterpreter string

	// Type of the EVM interpreter ("" for default)
	EVMInterpreter string

	// Constantinople block override (TODO: remove after the fork)
	ConstantinopleOverride *big.Int

	// RPCGasCap is the global gas cap for eth-call variants.
	RPCGasCap *big.Int `toml:",omitempty"`
}

type configMarshaling struct {
	MinerExtraData hexutil.Bytes
}
