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

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lvl  Level
		want string
	}{
		{lvl: Trace, want: "trace"},
		{lvl: Debug, want: "debug"},
		{lvl: Info, want: "info"},
		{lvl: Notice, want: "notice"},
		{lvl: Warn, want: "warn"},
		{lvl: Error, want: "error"},
		{lvl: Critical, want: "critical"},
		{lvl: Fatal, want: "fatal"},
		{lvl: Level(-7), want: "trace2"},
		{lvl: Level(-6), want: "trace3"},
		{lvl: Level(-5), want: "trace4"},
		{lvl: Level(-3), want: "debug2"},
		{lvl: Level(2), want: "info3"},
		{lvl: Level(10), want: "error3"},
		{lvl: Level(15), want: "fatal4"},
		{lvl: Off, want: "off"},
		{lvl: Level(16), want: "level(16)"},
	}

	for _, tt := range tests {
		if got := tt.lvl.String(); got != tt.want {
			t.Fatalf("%d.String() = %q, want %q", tt.lvl, got, tt.want)
		}
	}
}
