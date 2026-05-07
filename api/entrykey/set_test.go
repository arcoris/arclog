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

func TestDefaultSet(t *testing.T) {
	t.Parallel()

	set := entrykey.DefaultSet()
	if !set.IsDefault() {
		t.Fatal("DefaultSet().IsDefault() = false")
	}
	if set.IsZero() {
		t.Fatal("DefaultSet().IsZero() = true")
	}

	assertKeys(t, set.Metadata(), entrykey.Metadata())
	assertKeys(t, set.Known(), entrykey.Known())
}

func TestZeroSet(t *testing.T) {
	t.Parallel()

	var set entrykey.Set
	if !set.IsZero() {
		t.Fatal("zero Set is not zero")
	}
	if set.IsDefault() {
		t.Fatal("zero Set reported as default")
	}

	wantMetadata := []entrykey.Key{"", "", "", "", "", "", ""}
	wantKnown := []entrykey.Key{"", "", "", "", "", "", "", ""}

	assertKeys(t, set.Metadata(), wantMetadata)
	assertKeys(t, set.Known(), wantKnown)
}

func TestCustomSetPreservesEmptyKeys(t *testing.T) {
	t.Parallel()

	set := entrykey.DefaultSet()
	set.Stacktrace = ""
	set.Function = "func"

	got := set.Metadata()
	want := []entrykey.Key{
		entrykey.Time,
		entrykey.Level,
		entrykey.Logger,
		entrykey.Message,
		entrykey.Caller,
		"func",
		"",
	}

	assertKeys(t, got, want)
}
