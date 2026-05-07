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

package hook_test

import (
	"context"
	"errors"
	"testing"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/hook"
)

var _ hook.PreWriteHook = hook.PreWriteFunc(nil)

func TestPreWriteFuncPassesValuesAndReturnsFunctionResult(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("veto")
	inEntry := core.Entry{Message: "before"}
	inFields := []field.Field{field.String("a", "b")}

	pre := hook.PreWriteFunc(func(ctx context.Context, entry core.Entry, fields []field.Field) (core.Entry, []field.Field, error) {
		if ctx == nil {
			t.Fatal("context is nil")
		}
		entry.Message = "after"
		fields = append(fields, field.String("c", "d"))
		return entry, fields, wantErr
	})

	gotEntry, gotFields, err := pre.PreWrite(context.Background(), inEntry, inFields)
	if !errors.Is(err, wantErr) {
		t.Fatalf("PreWrite() error = %v, want %v", err, wantErr)
	}
	if gotEntry.Message != "after" {
		t.Fatalf("Message = %q, want after", gotEntry.Message)
	}
	if len(gotFields) != 2 {
		t.Fatalf("len(fields) = %d, want 2", len(gotFields))
	}
	if len(inFields) != 1 {
		t.Fatalf("input len(fields) = %d, want 1", len(inFields))
	}
}

func TestNilPreWriteFuncReturnsInputsUnchanged(t *testing.T) {
	t.Parallel()

	entry := core.Entry{Message: "entry"}
	fields := []field.Field{field.String("k", "v")}

	var pre hook.PreWriteFunc
	gotEntry, gotFields, err := pre.PreWrite(context.Background(), entry, fields)
	if err != nil {
		t.Fatalf("PreWrite() error = %v", err)
	}
	if gotEntry.Message != entry.Message {
		t.Fatalf("entry = %#v, want %#v", gotEntry, entry)
	}
	if len(gotFields) != len(fields) {
		t.Fatalf("len(fields) = %d, want %d", len(gotFields), len(fields))
	}
	if gotFields[0].Key != fields[0].Key {
		t.Fatalf("field key = %q, want %q", gotFields[0].Key, fields[0].Key)
	}
}

func TestPreWriteFuncDoesNotAllocate(t *testing.T) {
	pre := hook.PreWriteFunc(func(context.Context, core.Entry, []field.Field) (core.Entry, []field.Field, error) {
		return core.Entry{}, nil, nil
	})

	allocs := testing.AllocsPerRun(1000, func() {
		_, _, _ = pre.PreWrite(context.Background(), core.Entry{}, nil)
	})
	if allocs != 0 {
		t.Fatalf("allocs per PreWrite() = %g, want 0", allocs)
	}
}
