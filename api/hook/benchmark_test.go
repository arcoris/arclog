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
	"testing"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/hook"
)

func BenchmarkAllowsAsync(b *testing.B) {
	h := asyncDeclaration(true)

	b.ReportAllocs()
	for b.Loop() {
		_ = hook.AllowsAsync(h)
	}
}

func BenchmarkPreWriteFunc(b *testing.B) {
	h := hook.PreWriteFunc(func(_ context.Context, entry core.Entry, fields []field.Field) (core.Entry, []field.Field, error) {
		return entry, fields, nil
	})

	b.ReportAllocs()
	for b.Loop() {
		_, _, _ = h.PreWrite(context.Background(), core.Entry{}, nil)
	}
}

func BenchmarkPostWriteFunc(b *testing.B) {
	h := hook.PostWriteFunc(func(context.Context, core.Entry, []field.Field, hook.WriteResult) error {
		return nil
	})

	b.ReportAllocs()
	for b.Loop() {
		_ = h.PostWrite(context.Background(), core.Entry{}, nil, hook.Success())
	}
}

func BenchmarkErrorFunc(b *testing.B) {
	h := hook.ErrorFunc(func(context.Context, core.Entry, []field.Field, error) {})

	b.ReportAllocs()
	for b.Loop() {
		h.OnError(context.Background(), core.Entry{}, nil, nil)
	}
}

func BenchmarkRegistrationFunc(b *testing.B) {
	registration := hook.RegistrationFunc(func() bool {
		return true
	})

	b.ReportAllocs()
	for b.Loop() {
		_ = registration.Remove()
	}
}
