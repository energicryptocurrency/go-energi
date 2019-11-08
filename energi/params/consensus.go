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

const (
	MaturityPeriod    uint64 = 60 * 60
	AverageTimeBlocks uint64 = 60
	TargetBlockGap    uint64 = 60
	MinBlockGap       uint64 = 30
	MaxFutureGap      uint64 = 3
	TargetPeriodGap   uint64 = AverageTimeBlocks * TargetBlockGap

	// DoS protection
	OldForkPeriod uint64 = 15 * 60
	StakeThrottle uint64 = 60

	UnlimitedGas uint64 = (1 << 40)

	// MaxCheckpointVoteBlockAge defines the period in blocks count from the time
	// the checkpoint signer account proposes a checkpoint in which its voting
	// is permitted.
	MaxCheckpointVoteBlockAge = 1440
)
