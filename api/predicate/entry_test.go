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

package predicate_test

import (
	"testing"

	"arcoris.dev/arclog/api/caller"
	"arcoris.dev/arclog/api/level"
	"arcoris.dev/arclog/api/predicate"
)

func TestEntryZeroValue(t *testing.T) {
	t.Parallel()

	var entry predicate.Entry

	if entry.Level != level.Debug {
		t.Fatalf("Level = %s, want zero-value level %s", entry.Level, level.Debug)
	}
	if entry.Logger != "" {
		t.Fatalf("Logger = %q, want empty string", entry.Logger)
	}
	if entry.Message != "" {
		t.Fatalf("Message = %q, want empty string", entry.Message)
	}
	if entry.Caller.Defined {
		t.Fatal("Caller.Defined = true, want false")
	}
}

func TestEntryCarriesAPILayerMetadata(t *testing.T) {
	t.Parallel()

	want := predicate.Entry{
		Level:   level.Warn,
		Logger:  "api",
		Message: "connected",
		Caller: caller.Caller{
			Defined:  true,
			File:     "logger.go",
			Line:     42,
			Function: "log",
		},
	}

	got := want
	if got != want {
		t.Fatalf("Entry = %#v, want %#v", got, want)
	}
}
