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
	"testing"
)

// TestParseValidLevels verifies canonical names, aliases, case-insensitivity,
// and surrounding whitespace handling.
func TestParseValidLevels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  Level
	}{
		{"trace", "trace", Trace},
		{"trace-upper", "TRACE", Trace},
		{"debug", "debug", Debug},
		{"info", "info", Info},
		{"information", "Information", Info},
		{"informational", "informational", Info},
		{"notice", "notice", Notice},
		{"warn", "warn", Warn},
		{"warning", "WARNING", Warn},
		{"error", "error", Error},
		{"err", "ERR", Error},
		{"critical", "critical", Critical},
		{"crit", "CRIT", Critical},
		{"fatal", "fatal", Fatal},
		{"panic", "panic", Panic},
		{"spaces", "  info  ", Info},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
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

// TestParseInvalidLevels verifies that unknown names are rejected rather than
// silently coerced to a default threshold.
func TestParseInvalidLevels(t *testing.T) {
	t.Parallel()

	inputs := []string{"", "noop", "verbose", " infoo ", "warnn", "panic!"}
	for _, input := range inputs {
		input := input
		t.Run(input, func(t *testing.T) {
			t.Parallel()

			got, err := Parse(input)
			if err == nil {
				t.Fatalf("Parse(%q) returned nil error", input)
			}
			if got != Invalid {
				t.Fatalf("Parse(%q) = %v, want Invalid", input, got)
			}
		})
	}
}

// TestMustParseSucceeds verifies the intended static-initialization path.
func TestMustParseSucceeds(t *testing.T) {
	t.Parallel()

	if got := MustParse("INFO"); got != Info {
		t.Fatalf("MustParse(%q) = %v, want %v", "INFO", got, Info)
	}
}

// TestMustParsePanics verifies that MustParse is not a user-input API.
func TestMustParsePanics(t *testing.T) {
	t.Parallel()

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatalf("MustParse did not panic")
		}
	}()

	_ = MustParse("not-a-level")
}
