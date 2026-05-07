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

func TestKeyString(t *testing.T) {
	t.Parallel()

	if got := entrykey.Message.String(); got != "message" {
		t.Fatalf("String() = %q, want %q", got, "message")
	}
}

func TestKeyIsZero(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		key  entrykey.Key
		want bool
	}{
		{name: "zero", key: "", want: true},
		{name: "known", key: entrykey.Level, want: false},
		{name: "unknown", key: "custom", want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.key.IsZero(); got != tt.want {
				t.Fatalf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyIsKnown(t *testing.T) {
	t.Parallel()

	for _, key := range entrykey.Known() {
		key := key
		t.Run(key.String(), func(t *testing.T) {
			t.Parallel()

			if !key.IsKnown() {
				t.Fatalf("%q is not known", key)
			}
		})
	}

	unknown := []entrykey.Key{"", "custom", "http.method", "db.statement", "trace_id"}
	for _, key := range unknown {
		key := key
		t.Run("unknown/"+key.String(), func(t *testing.T) {
			t.Parallel()

			if key.IsKnown() {
				t.Fatalf("%q reported as known", key)
			}
		})
	}
}
