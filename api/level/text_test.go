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

func TestMarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lvl  Level
		want string
	}{
		{lvl: Trace, want: "trace"},
		{lvl: Level(-7), want: "trace2"},
		{lvl: Notice, want: "notice"},
		{lvl: Critical, want: "critical"},
		{lvl: Level(2), want: "info3"},
		{lvl: Level(10), want: "error3"},
		{lvl: Off, want: "off"},
	}

	for _, tt := range tests {
		got, err := tt.lvl.MarshalText()
		if err != nil {
			t.Fatalf("%v.MarshalText() returned error: %v", tt.lvl, err)
		}
		if string(got) != tt.want {
			t.Fatalf("%v.MarshalText() = %q, want %q", tt.lvl, string(got), tt.want)
		}
	}
}

func TestMarshalTextErrorsOnInvalid(t *testing.T) {
	t.Parallel()

	for _, lvl := range []Level{Trace - 1, Level(16), Level(126)} {
		if got, err := lvl.MarshalText(); err == nil {
			t.Fatalf("%v.MarshalText() = %q, nil error; want error", lvl, string(got))
		}
	}
}

func TestUnmarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  Level
	}{
		{input: "warn2", want: Level(5)},
		{input: "info2", want: Notice},
		{input: "error2", want: Critical},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			var lvl Level
			if err := lvl.UnmarshalText([]byte(tt.input)); err != nil {
				t.Fatalf("UnmarshalText returned error: %v", err)
			}
			if lvl != tt.want {
				t.Fatalf("level after UnmarshalText = %v, want %v", lvl, tt.want)
			}
		})
	}
}

func TestUnmarshalTextLeavesReceiverUnchangedOnError(t *testing.T) {
	t.Parallel()

	lvl := Warn
	if err := lvl.UnmarshalText([]byte("panic")); err == nil {
		t.Fatal("UnmarshalText returned nil error, want error")
	}
	if lvl != Warn {
		t.Fatalf("level after failed UnmarshalText = %v, want %v", lvl, Warn)
	}
}

func TestUnmarshalTextNilReceiver(t *testing.T) {
	t.Parallel()

	var lvl *Level
	if err := lvl.UnmarshalText([]byte("info")); err == nil {
		t.Fatal("nil receiver UnmarshalText returned nil error, want error")
	}
}
