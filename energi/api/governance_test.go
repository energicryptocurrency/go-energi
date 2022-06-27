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

package api

import "testing"

func TestIsStrictUUID4(t *testing.T) {
	cases := []struct {
		input string;
		output bool;
	}{
		{ "8c65a180-f602-11ec-b939-0242ac120002", false },
		{ "c5f5fe2c-f602-11ec-b939-0242ac120002", false },
		{ "3d7133f7-49b9-4991-bb93-1eb660789ef1", true },
		{ "adcdac65-875a-41d2-b02d-d0cce7e2bcfd", true },
	}

	for _, tc := range cases {
		result, err := isStrictUUID4(tc.input)
		if err != nil {
			t.Error(err)
		} else if result != tc.output {
			t.Errorf("Test: %s. Expected: %v. Got: %v", tc.input, tc.output, result)
		}
	}
}