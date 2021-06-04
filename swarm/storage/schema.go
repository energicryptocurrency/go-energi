// Copyright 2018 The Energi Core Authors
// Copyright 2018 The go-ethereum Authors
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

package storage

// The DB schema we want to use. The actual/current DB schema might differ
// until migrations are run.
const CurrentDbSchema = DbSchemaHalloween

// There was a time when we had no schema at all.
const DbSchemaNone = ""

// "purity" is the first formal schema of LevelDB we release together with Swarm 0.3.5
const DbSchemaPurity = "purity"

// "halloween" is here because we had a screw in the garbage collector index.
// Because of that we had to rebuild the GC index to get rid of erroneous
// entries and that takes a long time. This schema is used for bookkeeping,
// so rebuild index will run just once.
const DbSchemaHalloween = "halloween"
