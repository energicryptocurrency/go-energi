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

package core

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"testing"

	"github.com/energicryptocurrency/energi/consensus/ethash"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/crypto"

	// "github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/params"

	"github.com/stretchr/testify/assert"
)

func TestCheckpoints(t *testing.T) {
	t.Parallel()
	// log.Root().SetHandler(log.StdoutHandler)

	engine := ethash.NewFaker()
	db, chain, err := newCanonical(engine, 10, true)
	if err != nil {
		t.Fatalf("failed to create pristine chain: %v", err)
	}
	defer chain.Stop()

	fpn := uint64(3)
	fp := chain.GetBlockByNumber(fpn)
	first_fork := chain.GetHeaderByNumber(fpn + 1).Hash()

	blocks := makeBlockChain(fp, 2, engine, db, canonicalSeed+1)
	_, err = chain.InsertChain(blocks)
	assert.Empty(t, err)
	second_fork := blocks[0].Hash()
	assert.NotEqual(t, first_fork, second_fork)
	assert.Equal(t, chain.checkpoints.latest, uint64(0))

	// Orig long chain
	curr_fork := chain.GetHeaderByNumber(fpn + 1).Hash()
	assert.Equal(t, first_fork, curr_fork)

	// log.Trace("Forced fork via checkpoint")
	err = chain.AddCheckpoint(
		Checkpoint{
			Number: fpn + 1,
			Hash:   second_fork,
		},
		[]CheckpointSignature{},
		true,
	)
	assert.Empty(t, err)

	curr_fork = chain.GetHeaderByNumber(fpn + 1).Hash()
	assert.Equal(t, second_fork, curr_fork)
	assert.Equal(t, chain.checkpoints.latest, fpn+1)

	// log.Trace("Unknown fork")
	unknown_fork := second_fork
	unknown_fork[0] = ^unknown_fork[0]

	err = chain.AddCheckpoint(
		Checkpoint{
			Number: fpn + 1,
			Hash:   unknown_fork,
		},
		[]CheckpointSignature{},
		true,
	)
	assert.Empty(t, err)

	assert.Empty(t, chain.GetHeaderByNumber(fpn+1))
	curr_fork = chain.CurrentHeader().Hash()
	assert.Equal(t, fp.Hash(), curr_fork)

	// log.Trace("Setup fake signer")
	signer, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)

	cfg := *chain.chainConfig
	chain.chainConfig = &cfg
	cfg.Energi = &params.EnergiConfig{
		CPPSigner: crypto.PubkeyToAddress(signer.PublicKey),
	}

	// log.Trace("Missing signature (remote)")
	err = chain.AddCheckpoint(
		Checkpoint{
			Number: fpn + 1,
			Hash:   curr_fork,
		},
		[]CheckpointSignature{},
		false,
	)
	assert.Equal(t, errors.New("missing checkpoint signatures"), err)

	// log.Trace("Invalid signature (remote)")
	sig, _ := crypto.Sign(curr_fork.Bytes(), signer)
	err = chain.AddCheckpoint(
		Checkpoint{
			Number: fpn + 1,
			Hash:   curr_fork,
		},
		[]CheckpointSignature{CheckpointSignature(sig)},
		false,
	)
	assert.Equal(t, errors.New("invalid CPP signature"), err)

	// log.Trace("Failed at checkpoint")
	blocks = makeBlockChain(fp, 2, engine, db, canonicalSeed+2)
	third_fork := blocks[0].Hash()
	third_second := blocks[1].Hash()

	_, err = chain.InsertHeaderChain([]*types.Header{
		blocks[0].Header(),
		blocks[1].Header(),
	}, 1)

	assert.Equal(t, ErrCheckpointMismatch, err)

	_, err = chain.InsertChain(blocks)
	assert.Equal(t, ErrCheckpointMismatch, err)

	// log.Trace("Valid remote checkpoint")
	cp := Checkpoint{
		Number: fpn + 1,
		Hash:   third_fork,
	}
	sig, _ = crypto.Sign(chain.checkpoints.hashToSign(&cp), signer)
	err = chain.AddCheckpoint(
		cp,
		[]CheckpointSignature{CheckpointSignature(sig)},
		false,
	)
	assert.Empty(t, err)
	assert.Equal(t, chain.checkpoints.latest, fpn+1)

	_, err = chain.InsertChain(blocks)
	assert.Empty(t, err)

	curr_fork = chain.GetHeaderByNumber(fpn + 1).Hash()
	assert.Equal(t, third_fork, curr_fork)

	// log.Trace("Valid remote checkpoint (second)")
	cp = Checkpoint{
		Number: fpn + 2,
		Hash:   third_second,
	}
	sig, _ = crypto.Sign(chain.checkpoints.hashToSign(&cp), signer)
	err = chain.AddCheckpoint(
		cp,
		[]CheckpointSignature{CheckpointSignature(sig)},
		false,
	)
	assert.Empty(t, err)
	assert.Equal(t, chain.checkpoints.latest, fpn+2)

	// log.Trace("Valid remote checkpoint (future)")
	cp = Checkpoint{
		Number: fpn + 10,
		Hash:   third_second,
	}
	sig, _ = crypto.Sign(chain.checkpoints.hashToSign(&cp), signer)
	err = chain.AddCheckpoint(
		cp,
		[]CheckpointSignature{CheckpointSignature(sig)},
		false,
	)
	assert.Empty(t, err)
	assert.Equal(t, chain.checkpoints.latest, fpn+2)
}
