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
	"fmt"
	"testing"

	"arcoris.dev/arclog/api/level"
)

// TestLevelString verifies the canonical lowercase representation used for
// diagnostics and text marshaling.
func TestLevelString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		lvl  level.Level
		want string
	}{
		{"trace", level.Trace, "trace"},
		{"debug", level.Debug, "debug"},
		{"info", level.Info, "info"},
		{"notice", level.Notice, "notice"},
		{"warn", level.Warn, "warn"},
		{"error", level.Error, "error"},
		{"critical", level.Critical, "critical"},
		{"fatal", level.Fatal, "fatal"},
		{"panic", level.Panic, "panic"},
		{"invalid", level.Invalid, "invalid"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.lvl.String(); got != tt.want {
				t.Fatalf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestLevelStringUnknown verifies the diagnostic fallback for out-of-range
// numeric values.
func TestLevelStringUnknown(t *testing.T) {
	t.Parallel()

	unknown := level.Level(42)
	want := fmt.Sprintf("Level(%d)", int8(unknown))
	if got := unknown.String(); got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}
