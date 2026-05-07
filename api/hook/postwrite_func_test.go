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

var _ hook.PostWriteHook = hook.PostWriteFunc(nil)

func TestPostWriteFuncPassesValuesAndReturnsFunctionError(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("observer failed")
	writeErr := errors.New("write failed")
	entry := core.Entry{Message: "entry"}
	fields := []field.Field{field.String("k", "v")}
	called := false

	post := hook.PostWriteFunc(func(ctx context.Context, gotEntry core.Entry, gotFields []field.Field, result hook.WriteResult) error {
		called = true
		if ctx == nil {
			t.Fatal("context is nil")
		}
		if gotEntry.Message != entry.Message {
			t.Fatalf("entry = %#v, want %#v", gotEntry, entry)
		}
		if len(gotFields) != len(fields) || gotFields[0].Key != fields[0].Key {
			t.Fatalf("fields = %#v, want %#v", gotFields, fields)
		}
		if !errors.Is(result.Err, writeErr) {
			t.Fatalf("result.Err = %v, want %v", result.Err, writeErr)
		}
		return wantErr
	})

	err := post.PostWrite(context.Background(), entry, fields, hook.Failure(writeErr))
	if !errors.Is(err, wantErr) {
		t.Fatalf("PostWrite() error = %v, want %v", err, wantErr)
	}
	if !called {
		t.Fatal("underlying function was not called")
	}
}

func TestNilPostWriteFuncReturnsNil(t *testing.T) {
	t.Parallel()

	var post hook.PostWriteFunc
	if err := post.PostWrite(context.Background(), core.Entry{}, nil, hook.Success()); err != nil {
		t.Fatalf("PostWrite() error = %v", err)
	}
}

func TestPostWriteFuncDoesNotAllocate(t *testing.T) {
	post := hook.PostWriteFunc(func(context.Context, core.Entry, []field.Field, hook.WriteResult) error {
		return nil
	})

	allocs := testing.AllocsPerRun(1000, func() {
		_ = post.PostWrite(context.Background(), core.Entry{}, nil, hook.Success())
	})
	if allocs != 0 {
		t.Fatalf("allocs per PostWrite() = %g, want 0", allocs)
	}
}
