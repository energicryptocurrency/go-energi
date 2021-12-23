// Copyright 2018 The Energi Core Authors
// Copyright 2016 The go-ethereum Authors
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

package params

import (
	"fmt"
	"math/big"

	"github.com/energicryptocurrency/energi/common"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0xd8e2a3b0ad08f8eaabaf653d25b7d9beee2911101010a2cd4f6692a9a1dc228a")
	TestnetGenesisHash = common.HexToHash("0x93b3a19ffad91326cd638f15930fdea1268c24d599b50b0e19299209c70c0e4e")

	MainnetMigrationSigner = common.HexToAddress("0xac34a2555de08384cd7960f35d3ab048fcf9f83a")
	TestnetMigrationSigner = common.HexToAddress("0xb1372ea07f6a92bc86fd5f8cdf468528f79f87ca")

	MainnetEBISigner = common.HexToAddress("0x01C3bB0b90C07b89CA38F56Bd9f3E8C160fec4dC")
	TestnetEBISigner = common.HexToAddress("0x25bbaaaf27ab1966c3ab9faf31277a1db7601f3f")

	MainnetCPPSigner = common.HexToAddress("0xBD1C57eACcfD1519E342F870C1c551983F839479")
	TestnetCPPSigner = common.HexToAddress("0xb1372ea07f6a92bc86fd5f8cdf468528f79f87ca")

	MainnetHFSigner = common.HexToAddress("0x44D16E845ec2d2D6A99a10fe44EE99DA0541CF31")
	TestnetHFSigner = common.HexToAddress("0x5b00118464fa6e73f9c2a4ea44e1cbfa9f5b83c6")

	MainnetBackbone = common.HexToAddress("0x79C7CF016E53e5C47906c2daF6De2aA00AAcdB1e")
	TestnetBackbone = common.HexToAddress("0x5143c57fcde025f05a19d0de9a7dac852e553624")

	MainnetHFProxy = common.HexToAddress("0xe2616f793916A0BD9C66939c08c94693d483df03")
	TestnetHFProxy = common.HexToAddress("0x886c71F1Af530478204dD12fB0BA34A46899C16D")
)

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	//MainnetGenesisHash: MainnetTrustedCheckpoint,
	//TestnetGenesisHash: TestnetTrustedCheckpoint,
}

// Seeing as there are 526,000 blocks per year, and there is a 12M NRG annual emission
// masternodes get 40% of all coins or 4.8M / 526,000 ~ 9.14 (per block except the super block)
// miners get 10% of all coins or 1.2M / 526,000 ~ 2.28 (per block including the super block)
// backbone gets 10% of all coins or 1.2M / 526,000 ~ 2.28 (per block including the super block)
// which adds up to 13.7 as block subsidy (block reward)
// Therefore with the total annual emission of ~12M per year:
// => 10% goes to energi backbone.
// => 10% goes to miners.
// => 40% goes to masternodes.
// => Remaining 40% goes to the Treasury. (~184000 NRG, super block cycle must be 20160 blocks)
// The above coins supply remain true even in the superblock interval of 20160 blocks.

var (
	EnergiMainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(39797),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		Energi: &EnergiConfig{
			BackboneAddress:              MainnetBackbone,
			MigrationSigner:              MainnetMigrationSigner,
			EBISigner:                    MainnetEBISigner,
			CPPSigner:                    MainnetCPPSigner,
			HFSigner:                     MainnetHFSigner,
			HardforkRegistryProxyAddress: MainnetHFProxy,
		},
		SuperblockCycle:     big.NewInt(60 * 24 * 14), // A super block happens at the end of every 20160 block (Approx. 14 days)
		MNRequireValidation: big.NewInt(10),
		MNValidationPeriod:  big.NewInt(5),
		MNCleanupPeriod:     big.NewInt(60 * 60 * 24 * 14), // Inactive MN denounced after 1209600 sec (14 days/ 2 weeks)
		MNEverCollateral:    new(big.Int).Mul(big.NewInt(3000000), big.NewInt(Ether)),
		MNRewardsPerBlock:   big.NewInt(10), // MN with the minimum collateral amount gets a block reward of (9.14/10) 0.914 NRG.
		HFFinalizationPeriod: big.NewInt(30), // The hardfork should be finalized in 30 blocks.
	}

	EnergiTestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(49797),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		Energi: &EnergiConfig{
			BackboneAddress:              TestnetBackbone,
			MigrationSigner:              TestnetMigrationSigner,
			EBISigner:                    TestnetEBISigner,
			CPPSigner:                    TestnetCPPSigner,
			HFSigner:                     TestnetHFSigner,
			HardforkRegistryProxyAddress: TestnetHFProxy,
		},
		SuperblockCycle:     big.NewInt(60 * 24),
		MNRequireValidation: big.NewInt(5),
		MNValidationPeriod:  big.NewInt(5),
		MNCleanupPeriod:     big.NewInt(60 * 60 * 3),
		MNEverCollateral:    new(big.Int).Mul(big.NewInt(30000), big.NewInt(Ether)),
		MNRewardsPerBlock:   big.NewInt(10),
		HFFinalizationPeriod: big.NewInt(10), // The hardfork should be finalized in 10 blocks.
	}

	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(39797),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		Ethash:              new(EthashConfig),
		Energi: &EnergiConfig{
			HardforkRegistryProxyAddress: MainnetHFProxy,
		},
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "mainnet",
		SectionIndex: 1,
		SectionHead:  common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		CHTRoot:      common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BloomRoot:    common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(49797),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		Ethash:              new(EthashConfig),
		Energi: &EnergiConfig{
			HardforkRegistryProxyAddress: TestnetHFProxy,
		},
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "testnet",
		SectionIndex: 1,
		SectionHead:  common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		CHTRoot:      common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BloomRoot:    common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, &EnergiConfig{ HardforkRegistryProxyAddress: common.BigToAddress(big.NewInt(0x30D)) }, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, &CliqueConfig{Period: 0, Epoch: 30000}, &EnergiConfig{ HardforkRegistryProxyAddress: common.BigToAddress(big.NewInt(0x30D)) }, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, &EnergiConfig{ HardforkRegistryProxyAddress: common.BigToAddress(big.NewInt(0x30D)) }, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	Name         string      `json:"-"`
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	EWASMBlock          *big.Int `json:"ewasmBlock,omitempty"`          // EWASM switch block (nil = no fork, 0 = already activated)

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	Energi *EnergiConfig `json:"energi,omitempty"`

	// This parameters are not part of hardcoded consensus!

	// SuperBlockCycle defines a period whose end results into a super block.
	// A superblock sends a lump sum reward to the treasury. During a superblock
	// no masternode reward that is sent out.
	SuperblockCycle *big.Int `json:"superblockCycle"`
	// MNRequireValidation defines the minimum number of active announced MN that
	// disable the invalidations check when checking if one of the active announced
	// MN should be rewarded.
	MNRequireValidation *big.Int `json:"mnRequireValidation"`
	// MNValidationPeriod maximum amount of time in seconds after which a MN can
	// possibly be invalidated.
	MNValidationPeriod *big.Int `json:"mnValidationPeriod"`
	// MNCleanupPeriod defines the waiting time after which inactive MN are
	// automatically denounced & dropped from the list of MNs expecting a reward.
	MNCleanupPeriod *big.Int `json:"mnCleanupPeriod"`
	// MNEverCollateral minimum MN collateral shown to the users on stats endpoint.
	MNEverCollateral *big.Int `json:"mnEverCollateral"`
	// MNRewardsPerBlock defines the fraction of the total MN reward per block share
	// payable to the MN holding the minimum amount of collateral.
	MNRewardsPerBlock *big.Int `json:"mnRewardsPerBlock"`
	// HFFinalizationPeriod is the number of blocks after the hardfork block,
	// within which a given hardfork must be finalized and made immutable or
	// rendered invalid and editable.
	HFFinalizationPeriod *big.Int `json:"hfFinalizationPeriod"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// EnergiConfig is the consensus engine config for proof-of-stake based sealing.
type EnergiConfig struct {
	BackboneAddress               common.Address `json:"backboneAddress"`
	MigrationSigner               common.Address `json:"migrationSigner"`
	EBISigner                     common.Address `json:"ebiSigner"`
	CPPSigner                     common.Address `json:"cppSigner"`
	HFSigner                      common.Address `json:"hfSigner"`
	HardforkRegistryProxyAddress  common.Address `json:"hfProxyAddress"`
}

// String implements the stringer interface, returning the consensus engine details.
func (*EnergiConfig) String() string {
	return "energi"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Energi != nil:
		engine = c.Energi
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v  ConstantinopleFix: %v Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsConstantinople(num):
		return GasTableConstantinople
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		return newCompatError("ConstantinopleFix fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                     *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158   bool
	IsByzantium, IsConstantinople, IsPetersburg bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
	}
}
