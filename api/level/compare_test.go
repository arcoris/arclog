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

package level

import (
	"fmt"
	"testing"
)

// TestLevelEnabled verifies the inclusive threshold rule used by arclog.
func TestLevelEnabled(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lvl       Level
		threshold Level
		want      bool
	}{
		{Trace, Trace, true},
		{Debug, Debug, true},
		{Info, Info, true},
		{Notice, Notice, true},
		{Warn, Warn, true},
		{Error, Error, true},
		{Critical, Critical, true},
		{Fatal, Fatal, true},
		{Panic, Panic, true},
		{Error, Info, true},
		{Warn, Info, true},
		{Critical, Error, true},
		{Fatal, Critical, true},
		{Panic, Fatal, true},
		{Debug, Info, false},
		{Trace, Debug, false},
		{Info, Warn, false},
		{Warn, Error, false},
		{Error, Fatal, false},
		{Invalid, Info, false},
		{Error, Invalid, false},
		{Level(42), Info, false},
		{Error, Level(42), false},
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
		lvl   Level
		other Level
		want  bool
	}{
		{Info, Info, false},
		{Error, Warn, true},
		{Warn, Info, true},
		{Fatal, Error, true},
		{Panic, Fatal, true},
		{Trace, Debug, false},
		{Debug, Info, false},
		{Invalid, Info, false},
		{Error, Invalid, false},
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
		lvl   Level
		other Level
		want  bool
	}{
		{Info, Info, false},
		{Trace, Debug, true},
		{Debug, Info, true},
		{Info, Warn, true},
		{Error, Fatal, true},
		{Fatal, Error, false},
		{Panic, Fatal, false},
		{Invalid, Info, false},
		{Error, Invalid, false},
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
