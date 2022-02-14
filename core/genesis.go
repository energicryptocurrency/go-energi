// Copyright 2018 The Energi Core Authors
// Copyright 2014 The go-ethereum Authors
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
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/common/hexutil"
	"github.com/energicryptocurrency/energi/common/math"
	"github.com/energicryptocurrency/energi/core/rawdb"
	"github.com/energicryptocurrency/energi/core/state"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/core/vm"
	energi_abi "github.com/energicryptocurrency/energi/energi/abi"
	energi_params "github.com/energicryptocurrency/energi/energi/params"
	"github.com/energicryptocurrency/energi/ethdb"
	"github.com/energicryptocurrency/energi/log"
	"github.com/energicryptocurrency/energi/params"
	"github.com/energicryptocurrency/energi/rlp"
)

//go:generate gencodec -type Genesis -field-override genesisSpecMarshaling -out gen_genesis.go
//go:generate gencodec -type GenesisAccount -field-override genesisAccountMarshaling -out gen_genesis_account.go
//go:generate gencodec -type GenesisXfer -field-override genesisXferMarshaling -out gen_genesis_xfer.go
var errGenesisNoConfig = errors.New("genesis has no chain configuration")

// Genesis specifies the header fields, state of a genesis block. It also defines hard
// fork switch-over blocks through the chain configuration.
type Genesis struct {
	Config     *params.ChainConfig `json:"config"`
	Nonce      uint64              `json:"nonce"`
	Timestamp  uint64              `json:"timestamp"`
	ExtraData  []byte              `json:"extraData"`
	GasLimit   uint64              `json:"gasLimit"   gencodec:"required"`
	Difficulty *big.Int            `json:"difficulty" gencodec:"required"`
	Mixhash    common.Hash         `json:"mixHash"`
	Coinbase   common.Address      `json:"coinbase"`
	Alloc      GenesisAlloc        `json:"alloc"      gencodec:"required"`
	Xfers      GenesisXfers        `json:"xfers"`

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number     uint64      `json:"number"`
	GasUsed    uint64      `json:"gasUsed"`
	ParentHash common.Hash `json:"parentHash"`
}

// GenesisAlloc specifies the initial state that is part of the genesis block.
type GenesisAlloc map[common.Address]GenesisAccount

func (ga *GenesisAlloc) UnmarshalJSON(data []byte) error {
	m := make(map[common.UnprefixedAddress]GenesisAccount)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	*ga = make(GenesisAlloc)
	for addr, a := range m {
		(*ga)[common.Address(addr)] = a
	}
	return nil
}

// GenesisAccount is an account in the state of the genesis block.
type GenesisAccount struct {
	Code       []byte                      `json:"code,omitempty"`
	Storage    map[common.Hash]common.Hash `json:"storage,omitempty"`
	Balance    *big.Int                    `json:"balance" gencodec:"required"`
	Nonce      uint64                      `json:"nonce,omitempty"`
	PrivateKey []byte                      `json:"secretKey,omitempty"` // for tests
}

type GenesisXfers []GenesisXfer

type GenesisXfer struct {
	Addr  common.Address `json:"addr" gencodec:"required"`
	Code  []byte         `json:"code" gencodec:"required"`
	Value *big.Int       `json:"value,omitempty"`
}

// field type overrides for gencodec
type genesisSpecMarshaling struct {
	Nonce      math.HexOrDecimal64
	Timestamp  math.HexOrDecimal64
	ExtraData  hexutil.Bytes
	GasLimit   math.HexOrDecimal64
	GasUsed    math.HexOrDecimal64
	Number     math.HexOrDecimal64
	Difficulty *math.HexOrDecimal256
	Alloc      map[common.UnprefixedAddress]GenesisAccount
}

type genesisAccountMarshaling struct {
	Code       hexutil.Bytes
	Balance    *math.HexOrDecimal256
	Nonce      math.HexOrDecimal64
	Storage    map[storageJSON]storageJSON
	PrivateKey hexutil.Bytes
}

type genesisXferMarshaling struct {
	Addr  common.UnprefixedAddress
	Code  hexutil.Bytes
	Value *math.HexOrDecimal256
}

// storageJSON represents a 256 bit byte array, but allows less than 256 bits when
// unmarshaling from hex.
type storageJSON common.Hash

func (h *storageJSON) UnmarshalText(text []byte) error {
	text = bytes.TrimPrefix(text, []byte("0x"))
	if len(text) > 64 {
		return fmt.Errorf("too many hex characters in storage key/value %q", text)
	}
	offset := len(h) - len(text)/2 // pad on the left
	if _, err := hex.Decode(h[offset:], text); err != nil {
		fmt.Println(err)
		return fmt.Errorf("invalid hex storage key/value %q", text)
	}
	return nil
}

func (h storageJSON) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

// GenesisMismatchError is raised when trying to overwrite an existing
// genesis block with an incompatible one.
type GenesisMismatchError struct {
	Stored, New common.Hash
}

func (e *GenesisMismatchError) Error() string {
	return fmt.Sprintf("database already contains an incompatible genesis block (have %x, new %x)", e.Stored[:8], e.New[:8])
}

// SetupGenesisBlock writes or updates the genesis block in db.
// The block that will be used is:
//
//                          genesis == nil       genesis != nil
//                       +------------------------------------------
//     db has no genesis |  main-net default  |  genesis
//     db has genesis    |  from DB           |  genesis (if compatible)
//
// The stored chain configuration will be updated if it is compatible (i.e. does not
// specify a fork block below the local head block). In case of a conflict, the
// error is a *params.ConfigCompatError and the new, unwritten config is returned.
//
// The returned chain configuration is never nil.
func SetupGenesisBlock(db ethdb.Database, genesis *Genesis) (*params.ChainConfig, common.Hash, error) {
	return SetupGenesisBlockWithOverride(db, genesis, nil)
}
func SetupGenesisBlockWithOverride(db ethdb.Database, genesis *Genesis, constantinopleOverride *big.Int) (*params.ChainConfig, common.Hash, error) {
	if genesis != nil && genesis.Config == nil {
		return params.AllEthashProtocolChanges, common.Hash{}, errGenesisNoConfig
	}
	// Just commit the new block if there is no stored genesis block.
	stored := rawdb.ReadCanonicalHash(db, 0)
	if (stored == common.Hash{}) {
		if genesis == nil {
			log.Info("Writing default main-net genesis block")
			genesis = DefaultGenesisBlock()
		} else {
			log.Info("Writing custom genesis block")
		}
		block, err := genesis.Commit(db)
		return genesis.Config, block.Hash(), err
	}

	// Check whether the genesis block is already written.
	if genesis != nil {
		hash := genesis.ToBlock(nil).Hash()
		if hash != stored {
			return genesis.Config, hash, &GenesisMismatchError{stored, hash}
		}
	}

	// Get the existing chain configuration.
	newcfg := genesis.configOrDefault(stored)
	if constantinopleOverride != nil {
		newcfg.ConstantinopleBlock = constantinopleOverride
		newcfg.PetersburgBlock = constantinopleOverride
	}
	storedcfg := rawdb.ReadChainConfig(db, stored)
	if storedcfg == nil {
		log.Warn("Found genesis block without chain config")
		rawdb.WriteChainConfig(db, stored, newcfg)
		return newcfg, stored, nil
	}
	// Special case: don't change the existing config of a non-mainnet chain if no new
	// config is supplied. These chains would get AllProtocolChanges (and a compat error)
	// if we just continued here.
	if genesis == nil && stored != params.MainnetGenesisHash {
		return storedcfg, stored, nil
	}

	// Check config compatibility and write the config. Compatibility errors
	// are returned to the caller unless we're already at block zero.
	height := rawdb.ReadHeaderNumber(db, rawdb.ReadHeadHeaderHash(db))
	if height == nil {
		return newcfg, stored, fmt.Errorf("missing block number for head header hash")
	}
	compatErr := storedcfg.CheckCompatible(newcfg, *height)
	if compatErr != nil && *height != 0 && compatErr.RewindTo != 0 {
		return newcfg, stored, compatErr
	}
	rawdb.WriteChainConfig(db, stored, newcfg)
	return newcfg, stored, nil
}

func (g *Genesis) configOrDefault(ghash common.Hash) *params.ChainConfig {
	switch {
	case g != nil:
		return g.Config
	case ghash == params.MainnetGenesisHash:
		return params.EnergiMainnetChainConfig
	case ghash == params.TestnetGenesisHash:
		return params.EnergiTestnetChainConfig
	default:
		// Returned when genesis is nil and matching ghash is not found.
		return params.TestChainConfig
	}
}

// ToBlock creates the genesis block and writes state of a genesis specification
// to the given database (or discards it if nil).
func (g *Genesis) ToBlock(db ethdb.Database) *types.Block {
	if db == nil {
		db = ethdb.NewMemDatabase()
	}
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db))
	for addr, account := range g.Alloc {
		statedb.AddBalance(addr, account.Balance)
		statedb.SetCode(addr, account.Code)
		statedb.SetNonce(addr, account.Nonce)
		for key, value := range account.Storage {
			statedb.SetState(addr, key, value)
		}
	}
	root := statedb.IntermediateRoot(false)
	head := &types.Header{
		Number:     new(big.Int).SetUint64(g.Number),
		Nonce:      types.EncodeNonce(g.Nonce),
		Time:       g.Timestamp,
		ParentHash: g.ParentHash,
		Extra:      g.ExtraData,
		GasLimit:   g.GasLimit,
		GasUsed:    g.GasUsed,
		Difficulty: g.Difficulty,
		MixDigest:  g.Mixhash,
		Coinbase:   g.Coinbase,
		Root:       root,
	}
	if g.GasLimit == 0 {
		head.GasLimit = params.GenesisGasLimit
	}
	if g.Difficulty == nil {
		head.Difficulty = params.GenesisDifficulty
	}

	// Process transactions, but do not really record them in Genesis
	//---
	if config := g.Config; config != nil {
		debug := false
		author := energi_params.Energi_TreasuryV1
		gasLimit := uint64(100000000)
		gp := new(GasPool)

		systemFaucet := energi_params.Energi_SystemFaucet
		statedb.SetBalance(systemFaucet, math.MaxBig256)

		vmcfg := vm.Config{}

		if debug {
			vmcfg = vm.Config{
				Debug: true,
				Tracer: vm.NewStructLogger(&vm.LogConfig{
					Debug: true,
				}),
			}
		}

		for i, tx := range g.Xfers {
			gp.AddGas(gasLimit)
			val := tx.Value

			if val == nil {
				val = common.Big1
			}

			msg := types.NewMessage(
				systemFaucet,
				&tx.Addr,
				uint64(i),
				common.Big0,
				gasLimit,
				common.Big0,
				tx.Code,
				false,
			)
			ctx := NewEVMContext(msg, head, nil, &author)
			ctx.GasLimit = gasLimit
			evm := vm.NewEVM(ctx, statedb, g.Config, vmcfg)
			sttrans := NewStateTransition(evm, msg, gp)
			sttrans.inSetup = true
			_, _, _, err := sttrans.TransitionDb()
			if err != nil {
				panic(fmt.Errorf("invalid transaction: %v", err))
			}

			if statedb.GetCodeSize(tx.Addr) == 0 {
				panic(fmt.Errorf("Failed to create a contract%v", tx.Addr))
			}

			statedb.AddBalance(tx.Addr, val)
		}

		if debug {
			vm.WriteTrace(os.Stderr, vmcfg.Tracer.(*vm.StructLogger).StructLogs())
			vm.WriteLogs(os.Stderr, statedb.Logs())
		}

		statedb.SetBalance(systemFaucet, big.NewInt(0))
		statedb.SetBalance(author, big.NewInt(0))
		root = statedb.IntermediateRoot(false)
		head.Root = root
	}
	//---

	statedb.Commit(false)
	statedb.Database().TrieDB().Commit(root, true)

	return types.NewBlock(head, nil, nil, nil)
}

// Commit writes the block and state of a genesis specification to the database.
// The block is committed as the canonical head block.
func (g *Genesis) Commit(db ethdb.Database) (*types.Block, error) {
	block := g.ToBlock(db)
	if block.Number().Sign() != 0 {
		return nil, fmt.Errorf("can't commit genesis block with number > 0")
	}
	rawdb.WriteTd(db, block.Hash(), block.NumberU64(), g.Difficulty)
	rawdb.WriteBlock(db, block)
	rawdb.WriteReceipts(db, block.Hash(), block.NumberU64(), nil)
	rawdb.WriteCanonicalHash(db, block.Hash(), block.NumberU64())
	rawdb.WriteHeadBlockHash(db, block.Hash())
	rawdb.WriteHeadHeaderHash(db, block.Hash())

	config := g.Config
	if config == nil {
		config = params.AllEthashProtocolChanges
	}
	rawdb.WriteChainConfig(db, block.Hash(), config)
	return block, nil
}

// MustCommit writes the genesis block and state to db, panicking on error.
// The block is committed as the canonical head block.
func (g *Genesis) MustCommit(db ethdb.Database) *types.Block {
	block, err := g.Commit(db)
	if err != nil {
		panic(err)
	}
	return block
}

// GenesisBlockForTesting creates and writes a block in which addr has the given wei balance.
func GenesisBlockForTesting(db ethdb.Database, addr common.Address, balance *big.Int) *types.Block {
	g := Genesis{Alloc: GenesisAlloc{addr: {Balance: balance}}}
	return g.MustCommit(db)
}

// DefaultGenesisBlock returns the Ethereum main net genesis block.
func DefaultGenesisBlock() *Genesis {
	return &Genesis{
		Config:     params.MainnetChainConfig,
		Nonce:      66,
		ExtraData:  hexutil.MustDecode("0x11bbe8db4e347b4e8c937c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa"),
		GasLimit:   0,
		Difficulty: big.NewInt(17179869184),
		Alloc:      decodePrealloc(mainnetAllocData),
	}
}

// DefaultTestnetGenesisBlock returns the Ropsten network genesis block.
func DefaultTestnetGenesisBlock() *Genesis {
	return &Genesis{
		Config:     params.TestnetChainConfig,
		Nonce:      66,
		ExtraData:  hexutil.MustDecode("0x3535353535353535353535353535353535353535353535353535353535353535"),
		GasLimit:   0,
		Difficulty: big.NewInt(1048576),
		Alloc:      decodePrealloc(testnetAllocData),
	}
}

func DefaultEnergiMainnetGenesisBlock() *Genesis {
	return &Genesis{
		Config:     params.EnergiMainnetChainConfig,
		Coinbase:   energi_params.Energi_Treasury,
		Nonce:      0,
		Timestamp:  1583852648,
		ExtraData:  []byte{},
		GasLimit:   8000000,
		Difficulty: big.NewInt(0xFFFF),
		Alloc:      DefaultPrealloc(),
		Xfers:      DeployEnergiGovernance(params.EnergiMainnetChainConfig),
	}
}

func DefaultEnergiTestnetGenesisBlock() *Genesis {
	// Gen2 e3d97b4c8ce67242fbbc857ee64b49f3ce32b02df81b45359d3cd0b03c7b53ee
	// Time: 1561134790
	// Number: 367350
	return &Genesis{
		Config:     params.EnergiTestnetChainConfig,
		Coinbase:   energi_params.Energi_Treasury,
		Timestamp:  1561134790,
		Nonce:      0,
		ExtraData:  []byte{},
		GasLimit:   8000000,
		Difficulty: big.NewInt(0xFFFF),
		Alloc:      DefaultPrealloc(),
		Xfers:      DeployEnergiGovernance(params.EnergiTestnetChainConfig),
	}
}

// DeveloperEnergiGenesisBlock scans the custom genesis block from the provided
// file path.
func DeveloperEnergiGenesisBlock(customGenesisPath string) (*Genesis, error) {
	file, err := os.Open(customGenesisPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("Failed to read genesis file: %v", err)
	}

	genesis := new(Genesis)
	if err := json.NewDecoder(file).Decode(genesis); err != nil {
		return nil, fmt.Errorf("invalid genesis file: %v", err)
	}
	genesis.Xfers = append(genesis.Xfers, DeployEnergiGovernance(genesis.Config)...)
	for addr, account := range DefaultPrealloc() {
		genesis.Alloc[addr] = account
	}
	return genesis, nil
}

// DeveloperGenesisBlock returns the 'geth --dev' genesis block. Note, this must
// be seeded with the
func DeveloperGenesisBlock(period uint64, faucet common.Address) *Genesis {
	// Override the default period to the user requested one
	config := *params.AllCliqueProtocolChanges
	config.Clique.Period = period

	// Assemble and return the genesis with the precompiles and faucet pre-funded
	return &Genesis{
		Config:     &config,
		ExtraData:  append(append(make([]byte, 32), faucet[:]...), make([]byte, 65)...),
		GasLimit:   0,
		Difficulty: big.NewInt(1),
		Alloc: map[common.Address]GenesisAccount{
			common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
			common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
			common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
			common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
			common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
			common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
			common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
			common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
			faucet:                           {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
		},
	}
}

func DefaultPrealloc() GenesisAlloc {
	return GenesisAlloc{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
	}
}

func decodePrealloc(data string) GenesisAlloc {
	if len(data) == 0 {
		return DefaultPrealloc()
	}

	var p []struct{ Addr, Balance *big.Int }
	if err := rlp.NewStream(strings.NewReader(data), 0).Decode(&p); err != nil {
		panic(err)
	}
	ga := make(GenesisAlloc, len(p))
	for _, account := range p {
		ga[common.BigToAddress(account.Addr)] = GenesisAccount{Balance: account.Balance}
	}
	return ga
}

//=====================================
func deployEnergiContract(
	xfers *GenesisXfers,
	dst common.Address,
	value *big.Int,
	abi_json string, hex_code string,
	params ...interface{},
) {
	parsed_abi, err := abi.JSON(strings.NewReader(abi_json))
	if err != nil {
		panic(fmt.Errorf("invalid JSON: %v", err))
	}

	input, err := parsed_abi.Pack("", params...)
	if err != nil {
		panic(fmt.Errorf("invalid ABI: %v", err))
	}

	code := append(common.FromHex(hex_code), input...)
	*xfers = append(*xfers, GenesisXfer{
		Addr:  dst,
		Code:  code,
		Value: value,
	})
}

func DeployEnergiGovernance(config *params.ChainConfig) GenesisXfers {
	xfers := make(GenesisXfers, 0, 16)

	if config == nil {
		return xfers
	}

	//---
	ver := 2

	// Hardcoded Governance V1
	deployEnergiContract(
		&xfers,
		energi_params.Energi_BlockRewardV1,
		nil,
		energi_abi.BlockRewardV1ABI,
		energi_abi.BlockRewardV1Bin,
		energi_params.Energi_BlockReward,
		[4]common.Address{
			energi_params.Energi_StakerReward,
			energi_params.Energi_BackboneReward,
			energi_params.Energi_Treasury,
			energi_params.Energi_MasternodeRegistry,
		},
	)
	deployEnergiContract(
		&xfers,
		energi_params.Energi_TreasuryV1,
		nil,
		energi_abi.TreasuryV1ABI,
		energi_abi.TreasuryV1Bin,
		energi_params.Energi_Treasury,
		energi_params.Energi_MasternodeRegistry,
		config.SuperblockCycle,
	)
	deployEnergiContract(
		&xfers,
		energi_params.Energi_StakerRewardV1,
		nil,
		energi_abi.StakerRewardV1ABI,
		energi_abi.StakerRewardV1Bin,
		energi_params.Energi_StakerReward,
	)
	deployEnergiContract(
		&xfers,
		energi_params.Energi_BackboneRewardV1,
		nil,
		energi_abi.BackboneRewardV1ABI,
		energi_abi.BackboneRewardV1Bin,
		energi_params.Energi_BackboneReward,
		config.Energi.BackboneAddress,
	)
	if ver > 1 {
		deployEnergiContract(
			&xfers,
			energi_params.Energi_SporkRegistryV1,
			nil,
			energi_abi.SporkRegistryV2ABI,
			energi_abi.SporkRegistryV2Bin,
			energi_params.Energi_SporkRegistry,
			energi_params.Energi_MasternodeRegistry,
			config.Energi.CPPSigner,
		)
		deployEnergiContract(
			&xfers,
			energi_params.Energi_CheckpointRegistryV1,
			nil,
			energi_abi.CheckpointRegistryV2ABI,
			energi_abi.CheckpointRegistryV2Bin,
			energi_params.Energi_CheckpointRegistry,
			energi_params.Energi_MasternodeRegistry,
			config.Energi.CPPSigner,
		)
		deployEnergiContract(
			&xfers,
			energi_params.Energi_MasternodeTokenV1,
			nil,
			energi_abi.MasternodeTokenV2ABI,
			energi_abi.MasternodeTokenV2Bin,
			energi_params.Energi_MasternodeToken,
			energi_params.Energi_MasternodeRegistry,
		)
		deployEnergiContract(
			&xfers,
			energi_params.Energi_MasternodeRegistryV1,
			nil,
			energi_abi.MasternodeRegistryV2ABI,
			energi_abi.MasternodeRegistryV2Bin,
			energi_params.Energi_MasternodeRegistry,
			energi_params.Energi_MasternodeToken,
			energi_params.Energi_Treasury,
			[5]*big.Int{
				config.MNRequireValidation,
				config.MNValidationPeriod,
				config.MNCleanupPeriod,
				config.MNEverCollateral,
				config.MNRewardsPerBlock,
			},
		)
	} else {
		panic("Energi Genesis V1 is not supported anymore")
	}
	deployEnergiContract(
		&xfers,
		energi_params.Energi_BlacklistRegistryV1,
		nil,
		energi_abi.BlacklistRegistryV1ABI,
		energi_abi.BlacklistRegistryV1Bin,
		energi_params.Energi_BlacklistRegistry,
		energi_params.Energi_MasternodeRegistry,
		energi_params.Energi_MigrationContract,
		energi_params.Energi_CompensationFundV1,
		config.Energi.EBISigner,
	)
	deployEnergiContract(
		&xfers,
		energi_params.Energi_CompensationFundV1,
		nil,
		energi_abi.TreasuryV1ABI,
		energi_abi.TreasuryV1Bin,
		energi_params.Energi_BlacklistRegistryV1,
		energi_params.Energi_MasternodeRegistry,
		common.Big1, // unused
	)
	deployEnergiContract(
		&xfers,
		energi_params.Energi_MigrationContract,
		new(big.Int).Mul(big.NewInt(0xFFFF), big.NewInt(1e18)),
		energi_abi.Gen2MigrationABI,
		energi_abi.Gen2MigrationBin,
		energi_params.Energi_BlacklistRegistry,
		config.ChainID,
		config.Energi.MigrationSigner,
	)
	deployEnergiContract(
		&xfers,
		energi_params.Energi_Blacklist,
		nil,
		energi_abi.DummyAccountABI,
		energi_abi.DummyAccountBin,
	)
	deployEnergiContract(
		&xfers,
		energi_params.Energi_Whitelist,
		nil,
		energi_abi.DummyAccountABI,
		energi_abi.DummyAccountBin,
	)
	deployEnergiContract(
		&xfers,
		energi_params.Energi_MasternodeList,
		nil,
		energi_abi.DummyAccountABI,
		energi_abi.DummyAccountBin,
	)

	// Proxy List
	proxies := map[common.Address]common.Address{
		energi_params.Energi_BlockReward:        energi_params.Energi_BlockRewardV1,
		energi_params.Energi_Treasury:           energi_params.Energi_TreasuryV1,
		energi_params.Energi_MasternodeRegistry: energi_params.Energi_MasternodeRegistryV1,
		energi_params.Energi_StakerReward:       energi_params.Energi_StakerRewardV1,
		energi_params.Energi_BackboneReward:     energi_params.Energi_BackboneRewardV1,
		energi_params.Energi_SporkRegistry:      energi_params.Energi_SporkRegistryV1,
		energi_params.Energi_CheckpointRegistry: energi_params.Energi_CheckpointRegistryV1,
		energi_params.Energi_BlacklistRegistry:  energi_params.Energi_BlacklistRegistryV1,
		energi_params.Energi_MasternodeToken:    energi_params.Energi_MasternodeTokenV1,
	}

	// mainnet and testnet were deployed without the HF registry in the genesis block
	// therefore we only deploy when on some other network (devnet / simnet)
	if (config != params.EnergiMainnetChainConfig) && (config != params.EnergiTestnetChainConfig) {
		deployEnergiContract(
			&xfers,
			energi_params.Energi_HardforkRegistryV1,
			nil,
			energi_abi.HardforkRegistryV1ABI,
			energi_abi.HardforkRegistryV1Bin,
			energi_params.Energi_HardforkRegistry,
			config.Energi.HFSigner,
			config.HFFinalizationPeriod,
		)
		proxies[energi_params.Energi_HardforkRegistry] = energi_params.Energi_HardforkRegistryV1
	}

	for k, v := range proxies {
		deployEnergiContract(
			&xfers,
			k,
			nil,
			energi_abi.GovernedProxyABI,
			energi_abi.GovernedProxyBin,
			v,
			energi_params.Energi_SporkRegistry,
		)
	}

	return xfers
}
