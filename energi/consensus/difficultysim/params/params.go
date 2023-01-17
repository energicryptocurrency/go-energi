package params

const (
	// MaturityPeriodAsgard 30 min PoS cooldown for staking
	MaturityPeriodAsgard uint64 = 30 * 60
	// MaturityPeriod for diffV1
	MaturityPeriod uint64 = 60 * 60 // 1 hour PoS cooldown
	// TargetBlockGap targeted time interval between blocks
	TargetBlockGap uint64 = 60
	// MinBlockGap required time to pass after each block to start mining
	MinBlockGap uint64 = 30
	// BlockTimeEMAPeriod EMA calculation window size
	BlockTimeEMAPeriod uint64 = 360
	// MaxFutureGap block can be mined MaxFutureGap seconds into future
	MaxFutureGap uint64 = 3
	// InitialDifficulty starting difficulty
	InitialDifficulty uint64 = 210000
	// StakeReward Stake reward for each block
	StakeReward uint64 = 1
	// AveragingWindow average block window
	AveragingWindow uint64 = 60 // 60 blocks
	TargetPeriodGap uint64 = AveragingWindow * TargetBlockGap

	// AsgardIsActive determines the difficulty algorithm that is used for minig
	AsgardIsActive bool = true

	// BalanaPOSIsActive changes nonce selection algorithm
	BananaPOSIsActive bool = true
	BananaStakeCheckDepth uint64 = 0

	// hardfork that activates 15s block time
	BananaBlockTimeIsActive bool = true
	BananaTargetBlockGap uint64 = 15
	BananaMinBlockGap uint64 = 7

	// upon activating BananaDifficultyAdjustment hardfork new gain value will be used
	BananaDifficultyAdjustment bool = true
	GainBanana int64 = 20000

	// Parameters for difficulty calculations
	Gain           int64 = 50000
	IntegralTime   int64 = 720
	DerivativeTime int64 = 60
)
