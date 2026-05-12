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

func TestZeroValueLevelIsInfo(t *testing.T) {
	t.Parallel()

	var lvl Level
	if lvl != Info {
		t.Fatalf("zero value = %v, want Info", lvl)
	}
}

func TestNamedLevelNumericValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		lvl  Level
		want int8
	}{
		{name: "trace", lvl: Trace, want: -8},
		{name: "debug", lvl: Debug, want: -4},
		{name: "info", lvl: Info, want: 0},
		{name: "notice", lvl: Notice, want: 1},
		{name: "warn", lvl: Warn, want: 4},
		{name: "error", lvl: Error, want: 8},
		{name: "critical", lvl: Critical, want: int8(Critical)},
		{name: "fatal", lvl: Fatal, want: 12},
		{name: "off", lvl: Off, want: 127},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := int8(tt.lvl); got != tt.want {
				t.Fatalf("int8(%s) = %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}

func TestIsSeverity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		lvl  Level
		want bool
	}{
		{name: "trace", lvl: Trace, want: true},
		{name: "debug", lvl: Debug, want: true},
		{name: "info", lvl: Info, want: true},
		{name: "notice", lvl: Notice, want: true},
		{name: "warn", lvl: Warn, want: true},
		{name: "error", lvl: Error, want: true},
		{name: "critical", lvl: Critical, want: true},
		{name: "fatal", lvl: Fatal, want: true},
		{name: "trace2", lvl: Level(-7), want: true},
		{name: "fatal4", lvl: Level(15), want: true},
		{name: "off", lvl: Off, want: false},
		{name: "below range", lvl: Trace - 1, want: false},
		{name: "above range", lvl: Level(16), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.lvl.IsSeverity(); got != tt.want {
				t.Fatalf("%v.IsSeverity() = %v, want %v", tt.lvl, got, tt.want)
			}
		})
	}
}

func TestIsNamed(t *testing.T) {
	t.Parallel()

	named := []Level{Trace, Debug, Info, Notice, Warn, Error, Critical, Fatal}
	for _, lvl := range named {
		if !lvl.IsNamed() {
			t.Fatalf("%v.IsNamed() = false, want true", lvl)
		}
	}

	notNamed := []Level{Level(-7), Level(2), Level(10), Level(15), Off, Trace - 1, Level(16)}
	for _, lvl := range notNamed {
		if lvl.IsNamed() {
			t.Fatalf("%v.IsNamed() = true, want false", lvl)
		}
	}
}

func TestIsThreshold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		lvl  Level
		want bool
	}{
		{name: "severity", lvl: Info, want: true},
		{name: "custom severity", lvl: Level(15), want: true},
		{name: "off", lvl: Off, want: true},
		{name: "below range", lvl: Trace - 1, want: false},
		{name: "above range", lvl: Level(16), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.lvl.IsThreshold(); got != tt.want {
				t.Fatalf("%v.IsThreshold() = %v, want %v", tt.lvl, got, tt.want)
			}
		})
	}
}
