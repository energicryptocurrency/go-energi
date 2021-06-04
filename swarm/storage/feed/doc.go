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

/*
Package feeds defines Swarm Feeds.

Swarm Feeds allows a user to build an update feed about a particular topic
without resorting to ENS on each update.
The update scheme is built on swarm chunks with chunk keys following
a predictable, versionable pattern.

A Feed is tied to a unique identifier that is deterministically generated out of
the chosen topic.

A Feed is defined as the series of updates of a specific user about a particular topic

Actual data updates are also made in the form of swarm chunks. The keys
of the updates are the hash of a concatenation of properties as follows:

updateAddr = H(Feed, Epoch ID)
where H is the SHA3 hash function
Feed is the combination of Topic and the user address
Epoch ID is a time slot. See the lookup package for more information.

A user looking up a the latest update in a Feed only needs to know the Topic
and the other user's address.

The Feed Update data is:
updatedata = Feed|Epoch|data

The full update data that goes in the chunk payload is:
updatedata|sign(updatedata)

Structure Summary:

Request: Feed Update with signature
	Update: headers + data
		Header: Protocol version and reserved for future use placeholders
		ID: Information about how to locate a specific update
			Feed: Represents a user's series of publications about a specific Topic
				Topic: Item that the updates are about
				User: User who updates the Feed
			Epoch: time slot where the update is stored

*/
package feed
