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

package testutils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/params"
	"github.com/shengdoushi/base58"
)

type snapshotItem struct {
	Owner  string   `json:"owner"`
	Amount *big.Int `json:"amount"`
	Atype  string   `json:"type"`
}

type snapshot struct {
	Txouts    []snapshotItem `json:"snapshot_utxos"`
	Blacklist []string       `json:"snapshot_blacklist"`
	Hash      string         `json:"snapshot_hash"`
}

type TestGen2Migration struct {
	tempFile *os.File
}

// NewTestGen2Migration returns a testGen2Migration instance
func NewTestGen2Migration() *TestGen2Migration {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "node-simulation-")
	if err != nil {
		panic(fmt.Errorf("Cannot create temporary file: %v", err))
	}
	return &TestGen2Migration{tmpFile}
}

// CleanUp deletes the temporary file created. This should be called once the
// test is complete.
func (tg *TestGen2Migration) CleanUp() error {
	if tg == nil {
		return nil
	}
	return os.Remove(tg.tempFile.Name())
}

// TempFileName returns the temp migrations file path.
func (tg *TestGen2Migration) TempFileName() string {
	if tg == nil {
		return ""
	}
	return tg.tempFile.Name()
}

// PrepareTestGen2Migration creates a gen2 migration temp file.
func (tg *TestGen2Migration) PrepareTestGen2Migration(chainID uint64) error {
	if tg == nil {
		return nil
	}
	prefix := byte(33)
	if chainID == 49797 {
		prefix = byte(127)
	} else if chainID == 59797 {
		prefix = byte(129)
	} else if chainID != 39797 {
		log.Warn("unknown chain ID found: %d", chainID)
	}

	res := make([]byte, 20)
	_, err := rand.Read(res)
	if err != nil {
		return err
	}

	owner := make([]byte, 25)
	owner[0] = prefix
	copy(owner[1:], res[:])
	ownerhash := sha256.Sum256(owner[:21])
	copy(owner[21:], ownerhash[:4])

	items := int(params.MinGasLimit / 100000)
	snapshotItems := make([]snapshotItem, 0, items)
	for i := 0; i < items; i++ {
		snapshotItems = append(snapshotItems, snapshotItem{
			Owner:  base58.Encode(owner, base58.BitcoinAlphabet),
			Amount: big.NewInt(1000),
			Atype:  "pubkeyhash",
		})
	}

	migrations := snapshot{
		Txouts:    snapshotItems,
		Blacklist: []string{"tWFyUdwGxEkcam2aikVsDMPDpvMNKfP2XV"},
		Hash:      "778d7a438e3b86e0e754c4e46af802f852eb7c051d268c8599aa17c0cb9ce819",
	}

	data, err := json.Marshal(migrations)
	if err != nil {
		return err
	}

	_, err = tg.tempFile.Write(data)
	return err
}
