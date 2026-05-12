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

func TestFromSeverityNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		n    int
		want Level
		ok   bool
	}{
		{name: "trace", n: 1, want: Trace, ok: true},
		{name: "debug", n: 5, want: Debug, ok: true},
		{name: "info", n: severityNumberShift, want: Info, ok: true},
		{name: "notice", n: 10, want: Notice, ok: true},
		{name: "critical", n: 18, want: Critical, ok: true},
		{name: "fatal4", n: 24, want: Level(15), ok: true},
		{name: "zero", n: 0, ok: false},
		{name: "above range", n: 25, ok: false},
		{name: "negative", n: -1, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, ok := FromSeverityNumber(tt.n)
			if ok != tt.ok {
				t.Fatalf("FromSeverityNumber(%d) ok = %v, want %v", tt.n, ok, tt.ok)
			}
			if got != tt.want {
				t.Fatalf("FromSeverityNumber(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestSeverityNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lvl  Level
		want int
	}{
		{lvl: Trace, want: 1},
		{lvl: Debug, want: 5},
		{lvl: Info, want: severityNumberShift},
		{lvl: Notice, want: 10},
		{lvl: Warn, want: 13},
		{lvl: Error, want: 17},
		{lvl: Critical, want: 18},
		{lvl: Fatal, want: 21},
		{lvl: Level(-7), want: 2},
		{lvl: Level(2), want: 11},
		{lvl: Level(10), want: 19},
		{lvl: Level(15), want: 24},
		{lvl: Off, want: 0},
		{lvl: Trace - 1, want: 0},
		{lvl: Level(16), want: 0},
	}

	for _, tt := range tests {
		if got := tt.lvl.SeverityNumber(); got != tt.want {
			t.Fatalf("%v.SeverityNumber() = %d, want %d", tt.lvl, got, tt.want)
		}
	}
}

func TestSeverityText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lvl  Level
		want string
	}{
		{lvl: Trace, want: "TRACE"},
		{lvl: Debug, want: "DEBUG"},
		{lvl: Info, want: "INFO"},
		{lvl: Notice, want: "NOTICE"},
		{lvl: Warn, want: "WARN"},
		{lvl: Error, want: "ERROR"},
		{lvl: Critical, want: "CRITICAL"},
		{lvl: Fatal, want: "FATAL"},
		{lvl: Level(-7), want: "TRACE2"},
		{lvl: Level(-6), want: "TRACE3"},
		{lvl: Level(-5), want: "TRACE4"},
		{lvl: Level(-3), want: "DEBUG2"},
		{lvl: Level(2), want: "INFO3"},
		{lvl: Level(10), want: "ERROR3"},
		{lvl: Level(15), want: "FATAL4"},
		{lvl: Off, want: "OFF"},
		{lvl: Level(16), want: "LEVEL(16)"},
	}

	for _, tt := range tests {
		if got := tt.lvl.SeverityText(); got != tt.want {
			t.Fatalf("%v.SeverityText() = %q, want %q", tt.lvl, got, tt.want)
		}
	}
}
