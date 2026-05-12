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

func TestParseValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  Level
	}{
		{input: "trace", want: Trace},
		{input: "trace2", want: Level(-7)},
		{input: "trace3", want: Level(-6)},
		{input: "trace4", want: Level(-5)},
		{input: "DEBUG", want: Debug},
		{input: "debug2", want: Level(-3)},
		{input: "debug3", want: Level(-2)},
		{input: "debug4", want: Level(-1)},
		{input: " InFo ", want: Info},
		{input: "info2", want: Notice},
		{input: "notice", want: Notice},
		{input: "info3", want: Level(2)},
		{input: "info4", want: Level(3)},
		{input: "warn", want: Warn},
		{input: "warning", want: Warn},
		{input: "warn2", want: Level(5)},
		{input: "warn3", want: Level(6)},
		{input: "warn4", want: Level(7)},
		{input: "error", want: Error},
		{input: "error2", want: Critical},
		{input: "critical", want: Critical},
		{input: "crit", want: Critical},
		{input: "error3", want: Level(10)},
		{input: "error4", want: Level(11)},
		{input: "fatal", want: Fatal},
		{input: "fatal2", want: Level(13)},
		{input: "fatal3", want: Level(14)},
		{input: "fatal4", want: Level(15)},
		{input: "off", want: Off},
		{input: "disabled", want: Off},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) returned error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("Parse(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseRejectsInvalid(t *testing.T) {
	t.Parallel()

	for _, input := range []string{"", "panic", "invalid", "unspecified", "none", "whatever", "unknown"} {
		t.Run(input, func(t *testing.T) {
			t.Parallel()

			if got, err := Parse(input); err == nil {
				t.Fatalf("Parse(%q) = %v, nil error; want error", input, got)
			}
		})
	}
}

func TestParseAliases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  Level
	}{
		{input: "info2", want: Notice},
		{input: "notice", want: Notice},
		{input: "error2", want: Critical},
		{input: "critical", want: Critical},
		{input: "crit", want: Critical},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) returned error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("Parse(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseStringRoundTrip(t *testing.T) {
	t.Parallel()

	tests := []Level{
		Trace,
		Debug,
		Info,
		Notice,
		Warn,
		Error,
		Critical,
		Fatal,
		Level(-7),
		Level(2),
		Level(10),
	}

	for _, lvl := range tests {
		t.Run(lvl.String(), func(t *testing.T) {
			t.Parallel()

			got, err := Parse(lvl.String())
			if err != nil {
				t.Fatalf("Parse(%q) returned error: %v", lvl.String(), err)
			}
			if got != lvl {
				t.Fatalf("Parse(%q) = %v, want %v", lvl.String(), got, lvl)
			}
		})
	}
}

func TestMustParse(t *testing.T) {
	t.Parallel()

	if got := MustParse("INFO"); got != Info {
		t.Fatalf("MustParse(INFO) = %v, want %v", got, Info)
	}
}

func TestMustParsePanics(t *testing.T) {
	t.Parallel()

	defer func() {
		if recover() == nil {
			t.Fatal("MustParse did not panic")
		}
	}()

	_ = MustParse("panic")
}
