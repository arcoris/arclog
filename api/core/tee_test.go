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

package core_test

import (
	"errors"
	"testing"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
)

func TestTeeNoCoresReturnsNoop(t *testing.T) {
	t.Parallel()

	tee := core.Tee(nil, nil)
	if tee.Enabled(level.Info) {
		t.Fatal("empty Tee reported enabled")
	}
}

func TestTeeNoCoresDoesNotAllocate(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() {
		_ = core.Tee(nil, nil)
	})
	if allocs != 0 {
		t.Fatalf("allocs = %g, want 0", allocs)
	}
}

func TestTeeSingleCoreReturnsInput(t *testing.T) {
	t.Parallel()

	c := &recordingCore{enabled: true}
	if got := core.Tee(nil, c, nil); got != c {
		t.Fatal("Tee with one non-nil core did not return input core")
	}
}

func TestTeeSingleCoreDoesNotAllocate(t *testing.T) {
	c := &recordingCore{}

	allocs := testing.AllocsPerRun(1000, func() {
		_ = core.Tee(nil, c, nil)
	})
	if allocs != 0 {
		t.Fatalf("allocs = %g, want 0", allocs)
	}
}

func TestTeeEnabled(t *testing.T) {
	t.Parallel()

	tee := core.Tee(
		&recordingCore{enabled: false},
		&recordingCore{enabled: true},
	)

	if !tee.Enabled(level.Info) {
		t.Fatal("Tee did not report enabled when one core is enabled")
	}
}

func TestTeeSnapshotsCores(t *testing.T) {
	t.Parallel()

	var calls []string
	first := &recordingCore{name: "first", calls: &calls}
	second := &recordingCore{name: "second", calls: &calls}
	replacement := &recordingCore{name: "replacement", calls: &calls}

	cores := []core.Core{first, second}
	tee := core.Tee(cores...)
	cores[0] = replacement

	if err := tee.Write(core.Entry{}, nil); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	want := []string{"first", "second"}
	if len(calls) != len(want) {
		t.Fatalf("calls = %#v, want %#v", calls, want)
	}
	for i := range want {
		if calls[i] != want[i] {
			t.Fatalf("calls = %#v, want %#v", calls, want)
		}
	}
}

func TestTeeWith(t *testing.T) {
	t.Parallel()

	first := &recordingCore{}
	second := &recordingCore{}
	tee := core.Tee(first, second)

	with := tee.With([]field.Field{field.String("k", "v")})
	if with == nil {
		t.Fatal("With returned nil")
	}
	if len(first.withFields) != 1 || len(second.withFields) != 1 {
		t.Fatalf("With did not forward fields")
	}
}

func TestTeeCheck(t *testing.T) {
	t.Parallel()

	tee := core.Tee(
		&recordingCore{enabled: true},
		&recordingCore{enabled: false},
		&recordingCore{enabled: true},
	)

	ce := tee.Check(core.Entry{Level: level.Info}, nil)
	if ce == nil {
		t.Fatal("Check returned nil")
	}
	if ce.Len() != 2 {
		t.Fatalf("Len() = %d, want 2", ce.Len())
	}
}

func TestTeeWriteAttemptsAllCoresAfterError(t *testing.T) {
	t.Parallel()

	writeErr := errors.New("write")
	var calls []string

	tee := core.Tee(
		&recordingCore{name: "first", calls: &calls, writeErr: writeErr},
		&recordingCore{name: "second", calls: &calls},
	)

	if err := tee.Write(core.Entry{}, nil); !errors.Is(err, writeErr) {
		t.Fatalf("Write() error = %v, want writeErr", err)
	}

	want := []string{"first", "second"}
	if len(calls) != len(want) {
		t.Fatalf("calls = %#v, want %#v", calls, want)
	}
	for i := range want {
		if calls[i] != want[i] {
			t.Fatalf("calls = %#v, want %#v", calls, want)
		}
	}
}

func TestTeeWriteAndSyncAggregateErrors(t *testing.T) {
	t.Parallel()

	writeErr := errors.New("write")
	syncErr := errors.New("sync")

	tee := core.Tee(
		&recordingCore{writeErr: writeErr},
		&recordingCore{syncErr: syncErr},
	)

	if err := tee.Write(core.Entry{}, nil); !errors.Is(err, writeErr) {
		t.Fatalf("Write() error = %v, want writeErr", err)
	}
	if err := tee.Sync(); !errors.Is(err, syncErr) {
		t.Fatalf("Sync() error = %v, want syncErr", err)
	}
}
