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
	"runtime"
	"testing"
	"time"
)

func BenchmarkRuntimeMemStats(b *testing.B) {
	r := NewRegistry()
	RegisterRuntimeMemStats(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CaptureRuntimeMemStatsOnce(r)
	}
}

func TestRuntimeMemStats(t *testing.T) {
	r := NewRegistry()
	RegisterRuntimeMemStats(r)
	CaptureRuntimeMemStatsOnce(r)
	zero := runtimeMetrics.MemStats.PauseNs.Count() // Get a "zero" since GC may have run before these tests.
	runtime.GC()
	CaptureRuntimeMemStatsOnce(r)
	if count := runtimeMetrics.MemStats.PauseNs.Count(); 1 != count-zero {
		t.Fatal(count - zero)
	}
	runtime.GC()
	runtime.GC()
	CaptureRuntimeMemStatsOnce(r)
	if count := runtimeMetrics.MemStats.PauseNs.Count(); 3 != count-zero {
		t.Fatal(count - zero)
	}
	for i := 0; i < 256; i++ {
		runtime.GC()
	}
	CaptureRuntimeMemStatsOnce(r)
	if count := runtimeMetrics.MemStats.PauseNs.Count(); 259 != count-zero {
		t.Fatal(count - zero)
	}
	for i := 0; i < 257; i++ {
		runtime.GC()
	}
	CaptureRuntimeMemStatsOnce(r)
	if count := runtimeMetrics.MemStats.PauseNs.Count(); 515 != count-zero { // We lost one because there were too many GCs between captures.
		t.Fatal(count - zero)
	}
}

func TestRuntimeMemStatsNumThread(t *testing.T) {
	r := NewRegistry()
	RegisterRuntimeMemStats(r)
	CaptureRuntimeMemStatsOnce(r)

	if value := runtimeMetrics.NumThread.Value(); value < 1 {
		t.Fatalf("got NumThread: %d, wanted at least 1", value)
	}
}

func TestRuntimeMemStatsBlocking(t *testing.T) {
	if g := runtime.GOMAXPROCS(0); g < 2 {
		t.Skipf("skipping TestRuntimeMemStatsBlocking with GOMAXPROCS=%d\n", g)
	}
	ch := make(chan int)
	go testRuntimeMemStatsBlocking(ch)
	var memStats runtime.MemStats
	t0 := time.Now()
	runtime.ReadMemStats(&memStats)
	t1 := time.Now()
	t.Log("i++ during runtime.ReadMemStats:", <-ch)
	go testRuntimeMemStatsBlocking(ch)
	d := t1.Sub(t0)
	t.Log(d)
	time.Sleep(d)
	t.Log("i++ during time.Sleep:", <-ch)
}

func testRuntimeMemStatsBlocking(ch chan int) {
	i := 0
	for {
		select {
		case ch <- i:
			return
		default:
			i++
		}
	}
}
