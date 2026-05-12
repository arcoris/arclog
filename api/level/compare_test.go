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

import "testing"

func TestNamedLevelOrdering(t *testing.T) {
	t.Parallel()

	ordered := []Level{Trace, Debug, Info, Notice, Warn, Error, Critical, Fatal}
	for i := 1; i < len(ordered); i++ {
		if !ordered[i].Higher(ordered[i-1]) {
			t.Fatalf("%v should be higher than %v", ordered[i], ordered[i-1])
		}
		if !ordered[i-1].Lower(ordered[i]) {
			t.Fatalf("%v should be lower than %v", ordered[i-1], ordered[i])
		}
	}
}

func TestEnabled(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		record    Level
		threshold Level
		want      bool
	}{
		{name: "debug below info", record: Debug, threshold: Info, want: false},
		{name: "info at info", record: Info, threshold: Info, want: true},
		{name: "notice above info", record: Notice, threshold: Info, want: true},
		{name: "warn above info", record: Warn, threshold: Info, want: true},
		{name: "error above warn", record: Error, threshold: Warn, want: true},
		{name: "error below critical", record: Error, threshold: Critical, want: false},
		{name: "fatal at fatal", record: Fatal, threshold: Fatal, want: true},
		{name: "off threshold disables info", record: Info, threshold: Off, want: false},
		{name: "off threshold disables fatal", record: Fatal, threshold: Off, want: false},
		{name: "off is not record", record: Off, threshold: Info, want: false},
		{name: "invalid record", record: Level(16), threshold: Info, want: false},
		{name: "invalid threshold", record: Info, threshold: Level(16), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.record.Enabled(tt.threshold); got != tt.want {
				t.Fatalf("%v.Enabled(%v) = %v, want %v", tt.record, tt.threshold, got, tt.want)
			}
		})
	}
}

func TestHigherAndLowerRejectNonSeverities(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		left  Level
		right Level
	}{
		{name: "off left", left: Off, right: Info},
		{name: "off right", left: Error, right: Off},
		{name: "invalid left", left: Level(16), right: Info},
		{name: "invalid right", left: Error, right: Level(16)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.left.Higher(tt.right) {
				t.Fatalf("%v.Higher(%v) = true, want false", tt.left, tt.right)
			}
			if tt.left.Lower(tt.right) {
				t.Fatalf("%v.Lower(%v) = true, want false", tt.left, tt.right)
			}
		})
	}
}
