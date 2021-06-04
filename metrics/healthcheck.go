// Copyright 2018 The Energi Core Authors
// Copyright 2015 The go-ethereum Authors
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

// Healthchecks hold an error value describing an arbitrary up/down status.
type Healthcheck interface {
	Check()
	Error() error
	Healthy()
	Unhealthy(error)
}

// NewHealthcheck constructs a new Healthcheck which will use the given
// function to update its status.
func NewHealthcheck(f func(Healthcheck)) Healthcheck {
	if !Enabled {
		return NilHealthcheck{}
	}
	return &StandardHealthcheck{nil, f}
}

// NilHealthcheck is a no-op.
type NilHealthcheck struct{}

// Check is a no-op.
func (NilHealthcheck) Check() {}

// Error is a no-op.
func (NilHealthcheck) Error() error { return nil }

// Healthy is a no-op.
func (NilHealthcheck) Healthy() {}

// Unhealthy is a no-op.
func (NilHealthcheck) Unhealthy(error) {}

// StandardHealthcheck is the standard implementation of a Healthcheck and
// stores the status and a function to call to update the status.
type StandardHealthcheck struct {
	err error
	f   func(Healthcheck)
}

// Check runs the healthcheck function to update the healthcheck's status.
func (h *StandardHealthcheck) Check() {
	h.f(h)
}

// Error returns the healthcheck's status, which will be nil if it is healthy.
func (h *StandardHealthcheck) Error() error {
	return h.err
}

// Healthy marks the healthcheck as healthy.
func (h *StandardHealthcheck) Healthy() {
	h.err = nil
}

// Unhealthy marks the healthcheck as unhealthy.  The error is stored and
// may be retrieved by the Error method.
func (h *StandardHealthcheck) Unhealthy(err error) {
	h.err = err
}
