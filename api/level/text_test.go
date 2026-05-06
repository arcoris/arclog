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

package level_test

import (
	"encoding"
	"testing"

	"arcoris.dev/arclog/api/level"
)

var (
	_ encoding.TextMarshaler   = level.Level(0)
	_ encoding.TextUnmarshaler = (*level.Level)(nil)
)

// TestLevelMarshalText verifies that every valid severity has a stable textual
// representation.
func TestLevelMarshalText(t *testing.T) {
	t.Parallel()

	valid := []level.Level{
		level.Trace,
		level.Debug,
		level.Info,
		level.Notice,
		level.Warn,
		level.Error,
		level.Critical,
		level.Fatal,
		level.Panic,
	}

	for _, lvl := range valid {
		lvl := lvl
		t.Run(lvl.String(), func(t *testing.T) {
			t.Parallel()

			got, err := lvl.MarshalText()
			if err != nil {
				t.Fatalf("MarshalText(%v) returned error: %v", lvl, err)
			}
			if string(got) != lvl.String() {
				t.Fatalf("MarshalText(%v) = %q, want %q", lvl, string(got), lvl.String())
			}
		})
	}
}

// TestLevelMarshalTextRejectsInvalidValues verifies that sentinel and
// out-of-range levels are not serialized as valid severities.
func TestLevelMarshalTextRejectsInvalidValues(t *testing.T) {
	t.Parallel()

	invalid := []level.Level{level.Invalid, level.Level(42), level.Level(-100)}
	for _, lvl := range invalid {
		lvl := lvl
		t.Run(lvl.String(), func(t *testing.T) {
			t.Parallel()

			if _, err := lvl.MarshalText(); err == nil {
				t.Fatalf("MarshalText(%v) returned nil error", lvl)
			}
		})
	}
}

// TestLevelUnmarshalText verifies canonical parsing through the standard text
// unmarshaling contract.
func TestLevelUnmarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  level.Level
	}{
		{"trace", "trace", level.Trace},
		{"debug-upper", "DEBUG", level.Debug},
		{"info-spaces", " info ", level.Info},
		{"warning-alias", "WARNING", level.Warn},
		{"err-alias", "err", level.Error},
		{"crit-alias", "crit", level.Critical},
		{"fatal", "fatal", level.Fatal},
		{"panic", "panic", level.Panic},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			lvl := level.Trace
			if err := lvl.UnmarshalText([]byte(tt.input)); err != nil {
				t.Fatalf("UnmarshalText(%q) returned error: %v", tt.input, err)
			}
			if lvl != tt.want {
				t.Fatalf("UnmarshalText(%q) produced %v, want %v", tt.input, lvl, tt.want)
			}
		})
	}
}

// TestLevelUnmarshalTextLeavesReceiverUnchangedOnError verifies that failed
// parsing does not partially mutate caller state.
func TestLevelUnmarshalTextLeavesReceiverUnchangedOnError(t *testing.T) {
	t.Parallel()

	lvl := level.Info
	if err := lvl.UnmarshalText([]byte("not-a-level")); err == nil {
		t.Fatalf("UnmarshalText returned nil error")
	}
	if lvl != level.Info {
		t.Fatalf("UnmarshalText changed receiver to %v, want %v", lvl, level.Info)
	}
}

// TestLevelUnmarshalTextRejectsNilReceiver documents the nil-receiver policy.
func TestLevelUnmarshalTextRejectsNilReceiver(t *testing.T) {
	t.Parallel()

	var lvl *level.Level
	if err := lvl.UnmarshalText([]byte("info")); err == nil {
		t.Fatalf("UnmarshalText on nil receiver returned nil error")
	}
}
