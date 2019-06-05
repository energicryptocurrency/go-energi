// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package metrics

import (
	"fmt"
	"testing"
)

func BenchmarkGuage(b *testing.B) {
	g := NewGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Update(int64(i))
	}
}

func TestGauge(t *testing.T) {
	g := NewGauge()
	g.Update(int64(47))
	if v := g.Value(); 47 != v {
		t.Errorf("g.Value(): 47 != %v\n", v)
	}
}

func TestGaugeSnapshot(t *testing.T) {
	g := NewGauge()
	g.Update(int64(47))
	snapshot := g.Snapshot()
	g.Update(int64(0))
	if v := snapshot.Value(); 47 != v {
		t.Errorf("g.Value(): 47 != %v\n", v)
	}
}

func TestGetOrRegisterGauge(t *testing.T) {
	r := NewRegistry()
	NewRegisteredGauge("foo", r).Update(47)
	if g := GetOrRegisterGauge("foo", r); 47 != g.Value() {
		t.Fatal(g)
	}
}

func TestFunctionalGauge(t *testing.T) {
	var counter int64
	fg := NewFunctionalGauge(func() int64 {
		counter++
		return counter
	})
	fg.Value()
	fg.Value()
	if counter != 2 {
		t.Error("counter != 2")
	}
}

func TestGetOrRegisterFunctionalGauge(t *testing.T) {
	r := NewRegistry()
	NewRegisteredFunctionalGauge("foo", r, func() int64 { return 47 })
	if g := GetOrRegisterGauge("foo", r); 47 != g.Value() {
		t.Fatal(g)
	}
}

func ExampleGetOrRegisterGauge() {
	m := "server.bytes_sent"
	g := GetOrRegisterGauge(m, nil)
	g.Update(47)
	fmt.Println(g.Value()) // Output: 47
}
