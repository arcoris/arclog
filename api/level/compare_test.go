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

// TestLevelEnabled verifies the inclusive threshold rule used by arclog.
func TestLevelEnabled(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lvl       level.Level
		threshold level.Level
		want      bool
	}{
		{level.Trace, level.Trace, true},
		{level.Debug, level.Debug, true},
		{level.Info, level.Info, true},
		{level.Notice, level.Notice, true},
		{level.Warn, level.Warn, true},
		{level.Error, level.Error, true},
		{level.Critical, level.Critical, true},
		{level.Fatal, level.Fatal, true},
		{level.Panic, level.Panic, true},
		{level.Error, level.Info, true},
		{level.Warn, level.Info, true},
		{level.Critical, level.Error, true},
		{level.Fatal, level.Critical, true},
		{level.Panic, level.Fatal, true},
		{level.Debug, level.Info, false},
		{level.Trace, level.Debug, false},
		{level.Info, level.Warn, false},
		{level.Warn, level.Error, false},
		{level.Error, level.Fatal, false},
		{level.Invalid, level.Info, false},
		{level.Error, level.Invalid, false},
		{level.Level(42), level.Info, false},
		{level.Error, level.Level(42), false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%s/%s", tt.lvl, tt.threshold), func(t *testing.T) {
			t.Parallel()

			if got := tt.lvl.Enabled(tt.threshold); got != tt.want {
				t.Fatalf("Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestLevelHigher verifies valid greater-than severity comparison.
func TestLevelHigher(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lvl   level.Level
		other level.Level
		want  bool
	}{
		{level.Info, level.Info, false},
		{level.Error, level.Warn, true},
		{level.Warn, level.Info, true},
		{level.Fatal, level.Error, true},
		{level.Panic, level.Fatal, true},
		{level.Trace, level.Debug, false},
		{level.Debug, level.Info, false},
		{level.Invalid, level.Info, false},
		{level.Error, level.Invalid, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%s/%s", tt.lvl, tt.other), func(t *testing.T) {
			t.Parallel()

			if got := tt.lvl.Higher(tt.other); got != tt.want {
				t.Fatalf("Higher() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestLevelLower verifies valid less-than severity comparison.
func TestLevelLower(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lvl   level.Level
		other level.Level
		want  bool
	}{
		{level.Info, level.Info, false},
		{level.Trace, level.Debug, true},
		{level.Debug, level.Info, true},
		{level.Info, level.Warn, true},
		{level.Error, level.Fatal, true},
		{level.Fatal, level.Error, false},
		{level.Panic, level.Fatal, false},
		{level.Invalid, level.Info, false},
		{level.Error, level.Invalid, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%s/%s", tt.lvl, tt.other), func(t *testing.T) {
			t.Parallel()

			if got := tt.lvl.Lower(tt.other); got != tt.want {
				t.Fatalf("Lower() = %v, want %v", got, tt.want)
			}
		})
	}
}
