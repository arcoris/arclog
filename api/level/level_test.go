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
	"testing"

	"arcoris.dev/arclog/api/level"
)

// TestLevelOrdering documents the severity ordering that filtering code relies
// on throughout arclog.
func TestLevelOrdering(t *testing.T) {
	t.Parallel()

	ordered := []level.Level{
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

	for i := 1; i < len(ordered); i++ {
		if !(ordered[i-1] < ordered[i]) {
			t.Fatalf("level order broken at index %d: %v should be lower than %v", i, ordered[i-1], ordered[i])
		}
	}

	if !(level.Invalid > level.Panic) {
		t.Fatalf("Invalid = %d, want greater than Panic = %d", level.Invalid, level.Panic)
	}
}

// TestLevelIsValid verifies that only real log-entry severities are valid.
func TestLevelIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		lvl  level.Level
		want bool
	}{
		{"trace", level.Trace, true},
		{"debug", level.Debug, true},
		{"info", level.Info, true},
		{"notice", level.Notice, true},
		{"warn", level.Warn, true},
		{"error", level.Error, true},
		{"critical", level.Critical, true},
		{"fatal", level.Fatal, true},
		{"panic", level.Panic, true},
		{"invalid", level.Invalid, false},
		{"below-range", level.Level(-100), false},
		{"above-range", level.Level(100), false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.lvl.IsValid(); got != tt.want {
				t.Fatalf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
