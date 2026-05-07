/*
   Copyright 2026 The ARCORIS Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package clock_test

import (
	"testing"
	"time"

	"arcoris.dev/arclog/api/clock"
)

var _ clock.Clock = clock.Func(nil)

func TestFuncNowReturnsFunctionTimestamp(t *testing.T) {
	t.Parallel()

	want := time.Date(2026, 5, 8, 12, 30, 0, 0, time.UTC)
	c := clock.Func(func() time.Time {
		return want
	})

	if got := c.Now(); got != want {
		t.Fatalf("Now() = %#v, want %#v", got, want)
	}
}

func TestFuncNowCallsFunctionEachTime(t *testing.T) {
	t.Parallel()

	var calls int
	c := clock.Func(func() time.Time {
		calls++
		return time.Unix(int64(calls), int64(calls)).UTC()
	})

	first := c.Now()
	second := c.Now()

	if calls != 2 {
		t.Fatalf("calls = %d, want 2", calls)
	}
	if first == second {
		t.Fatalf("successive timestamps are equal: %#v", first)
	}
}

func TestFuncNowPreservesMonotonicValue(t *testing.T) {
	t.Parallel()

	want := time.Now()
	c := clock.Func(func() time.Time {
		return want
	})

	if got := c.Now(); got != want {
		t.Fatalf("Now() = %#v, want exact time value %#v", got, want)
	}
}

func TestFuncNowAllowsZeroTimestamp(t *testing.T) {
	t.Parallel()

	c := clock.Func(func() time.Time {
		return time.Time{}
	})

	if got := c.Now(); !got.IsZero() {
		t.Fatalf("Now().IsZero() = false for %#v", got)
	}
}

func TestNilFuncPanics(t *testing.T) {
	t.Parallel()

	mustPanic(t, func() {
		var c clock.Func
		_ = c.Now()
	})
}

func TestFuncNowDoesNotAllocate(t *testing.T) {
	want := time.Date(2026, 5, 8, 12, 30, 0, 0, time.UTC)
	c := clock.Func(func() time.Time {
		return want
	})

	allocs := testing.AllocsPerRun(1000, func() {
		_ = c.Now()
	})
	if allocs != 0 {
		t.Fatalf("allocs per Now() = %g, want 0", allocs)
	}
}

func BenchmarkFuncNow(b *testing.B) {
	want := time.Date(2026, 5, 8, 12, 30, 0, 0, time.UTC)
	c := clock.Func(func() time.Time {
		return want
	})

	b.ReportAllocs()
	for b.Loop() {
		_ = c.Now()
	}
}

func mustPanic(t *testing.T, fn func()) {
	t.Helper()

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("function did not panic")
		}
	}()

	fn()
}
