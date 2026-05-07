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

	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
	"arcoris.dev/arclog/api/predicate"
)

func TestFuncAdapterPassesEntryAndFields(t *testing.T) {
	t.Parallel()

	entry := predicate.Entry{
		Level:   level.Warn,
		Logger:  "api",
		Message: "connected",
	}
	fields := []field.Field{field.String("service", "auth")}

	var (
		called    bool
		gotEntry  predicate.Entry
		gotFields []field.Field
	)

	fn := predicate.Func(func(entry predicate.Entry, fields []field.Field) bool {
		called = true
		gotEntry = entry
		gotFields = fields
		return true
	})

	if !fn.ShouldLog(entry, fields) {
		t.Fatal("Func.ShouldLog() = false, want true")
	}
	if !called {
		t.Fatal("underlying function was not called")
	}
	if gotEntry != entry {
		t.Fatalf("entry = %#v, want %#v", gotEntry, entry)
	}
	if len(gotFields) != len(fields) || !gotFields[0].Equal(fields[0]) {
		t.Fatalf("fields = %#v, want %#v", gotFields, fields)
	}
}

func TestNilFuncPanics(t *testing.T) {
	t.Parallel()

	var fn predicate.Func
	mustPanic(t, func() {
		_ = fn.ShouldLog(predicate.Entry{}, nil)
	})
}
