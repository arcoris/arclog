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

package entrykey_test

import (
	"testing"

	"arcoris.dev/arclog/api/entrykey"
)

func TestCanonicalKeyValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		key  entrykey.Key
		want string
	}{
		{key: entrykey.Time, want: "time"},
		{key: entrykey.Level, want: "level"},
		{key: entrykey.Logger, want: "logger"},
		{key: entrykey.Message, want: "message"},
		{key: entrykey.Caller, want: "caller"},
		{key: entrykey.Function, want: "function"},
		{key: entrykey.Stacktrace, want: "stacktrace"},
		{key: entrykey.Error, want: "error"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			if got := tt.key.String(); got != tt.want {
				t.Fatalf("key = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMetadata(t *testing.T) {
	t.Parallel()

	got := entrykey.Metadata()
	want := []entrykey.Key{
		entrykey.Time,
		entrykey.Level,
		entrykey.Logger,
		entrykey.Message,
		entrykey.Caller,
		entrykey.Function,
		entrykey.Stacktrace,
	}

	assertKeys(t, got, want)
}

func TestKnown(t *testing.T) {
	t.Parallel()

	got := entrykey.Known()
	want := []entrykey.Key{
		entrykey.Time,
		entrykey.Level,
		entrykey.Logger,
		entrykey.Message,
		entrykey.Caller,
		entrykey.Function,
		entrykey.Stacktrace,
		entrykey.Error,
	}

	assertKeys(t, got, want)
}

func TestReturnedSlicesAreCopies(t *testing.T) {
	t.Parallel()

	metadata := entrykey.Metadata()
	metadata[0] = "mutated"

	if got := entrykey.Metadata()[0]; got != entrykey.Time {
		t.Fatalf("Metadata()[0] = %q, want %q after caller mutation", got, entrykey.Time)
	}

	known := entrykey.Known()
	known[0] = "mutated"

	if got := entrykey.Known()[0]; got != entrykey.Time {
		t.Fatalf("Known()[0] = %q, want %q after caller mutation", got, entrykey.Time)
	}
}

func TestKnownKeysAreUnique(t *testing.T) {
	t.Parallel()

	seen := make(map[entrykey.Key]struct{})
	for _, key := range entrykey.Known() {
		if key.IsZero() {
			t.Fatal("Known contains the zero key")
		}
		if _, ok := seen[key]; ok {
			t.Fatalf("Known contains duplicate key %q", key)
		}
		seen[key] = struct{}{}
	}
}

func assertKeys(t *testing.T, got, want []entrykey.Key) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d; got %#v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("key[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
