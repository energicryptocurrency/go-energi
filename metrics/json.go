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

package metrics

import (
	"encoding/json"
	"io"
	"time"
)

// MarshalJSON returns a byte slice containing a JSON representation of all
// the metrics in the Registry.
func (r *StandardRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.GetAll())
}

// WriteJSON writes metrics from the given registry  periodically to the
// specified io.Writer as JSON.
func WriteJSON(r Registry, d time.Duration, w io.Writer) {
	for range time.Tick(d) {
		WriteJSONOnce(r, w)
	}
}

// WriteJSONOnce writes metrics from the given registry to the specified
// io.Writer as JSON.
func WriteJSONOnce(r Registry, w io.Writer) {
	json.NewEncoder(w).Encode(r)
}

func (p *PrefixedRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.GetAll())
}
