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

package params

type ctxKey string

const (
	MaturityPeriod  uint64 = 60 * 60
	AveragingWindow uint64 = 60 // 60 blocks
	TargetBlockGap  uint64 = 60 // 60 seconds after the target is max diff
	MinBlockGap     uint64 = 30 // 30 seconds minimum before new timestamp
	MaxFutureGap    uint64 = 3  // only accept blocks this many seconds ahead
	// (subjective time)
	TargetPeriodGap uint64 = AveragingWindow * TargetBlockGap

	// DoS protection
	OldForkPeriod uint64 = 15 * 60
	StakeThrottle uint64 = 60

	UnlimitedGas uint64 = 1 << 40

	MasternodeCallGas uint64 = 1000000

	// TargetWindow defines the number of blocks within the AveragingWindow that
	// will be averaged when calculating the simple moving average.
	TargetWindow uint64 = 5

	// MaxCheckpointVoteBlockAge defines the period in blocks count from the time
	// the checkpoint signer account proposes a checkpoint in which its voting
	// is permitted.
	MaxCheckpointVoteBlockAge = 1440

	// GeneralProxyCtxKey is used to pass the governed proxy address hash to
	// the filter logs interface.
	GeneralProxyCtxKey = ctxKey("governedProxyAddressHash")

	// NB: Time difference between the block target and new block time is always
	// calulated as (blockTargetTime - newBlockTime).

	// MaxTimeDifferenceDrop defines the maximum time difference that can be used
	// to calculate the difficulty drop when the newly created block is found long
	// after block target time. This
	MaxTimeDifferenceDrop = -30

	// DifficultyChangeBase defines the base constant that is used generate the
	// difficulty multiplier. i.e.
	// Time difference :   Difficulty Multiplier
	//  +23            =>  1.0001^23 = 1.00230253177
	//  -23            =>  1.0001^-23 = 0.9977027577

	// Maximum positive time difference will never exceed 60 (TargetBlockGap)
	// and Minimum negative time difference will never exceed -30 (MaxTimeDifferenceDrop)
	// Therefore:

	// Max difficulty rise multiplier => 1.0001^60 = 1.00601773427 (increases to 100.6%)
	// Min difficulty drop multiplier => 1.0001^-30 = 0.9977027577 (drops to 99.77%)
	DifficultyChangeBase = 1.0001

	// DiffV2MigrationStakerTimeDelay roughly get 2x difficulty decrease
	DiffV2MigrationStakerTimeDelay uint64 = 15

	// DiffV2MigrationStakerBlockDelay defines block past which the migration staker
	// difficulty decrease is no longer enforced.
	DiffV2MigrationStakerBlockDelay uint64 = 10

	// DiffV2MigrationStakerTarget defines the maximum difficulty that the parent block
	// to what the migration staker is mining results to time decrease.
	DiffV2MigrationStakerTarget uint64 = 0xFFFF
)
