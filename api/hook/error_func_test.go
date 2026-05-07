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

var _ hook.ErrorHook = hook.ErrorFunc(nil)

func TestErrorFuncPassesValues(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("write failed")
	entry := core.Entry{Message: "entry"}
	fields := []field.Field{field.String("k", "v")}
	called := false

	errorHook := hook.ErrorFunc(func(ctx context.Context, gotEntry core.Entry, gotFields []field.Field, err error) {
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
		if !errors.Is(err, wantErr) {
			t.Fatalf("err = %v, want %v", err, wantErr)
		}
	})

	errorHook.OnError(context.Background(), entry, fields, wantErr)
	if !called {
		t.Fatal("underlying function was not called")
	}
}

func TestNilErrorFuncIsNoop(t *testing.T) {
	t.Parallel()

	var errorHook hook.ErrorFunc
	errorHook.OnError(context.Background(), core.Entry{}, nil, errors.New("ignored"))
}

func TestErrorFuncDoesNotAllocate(t *testing.T) {
	errorHook := hook.ErrorFunc(func(context.Context, core.Entry, []field.Field, error) {})

	allocs := testing.AllocsPerRun(1000, func() {
		errorHook.OnError(context.Background(), core.Entry{}, nil, nil)
	})
	if allocs != 0 {
		t.Fatalf("allocs per OnError() = %g, want 0", allocs)
	}
}
