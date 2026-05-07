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

package hook

import (
	"context"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
)

// PreWriteFunc adapts a function to PreWriteHook.
//
// A nil PreWriteFunc is a no-op hook that returns entry and fields unchanged.
type PreWriteFunc func(context.Context, core.Entry, []field.Field) (core.Entry, []field.Field, error)

// PreWrite calls f.
func (f PreWriteFunc) PreWrite(ctx context.Context, entry core.Entry, fields []field.Field) (core.Entry, []field.Field, error) {
	if f == nil {
		return entry, fields, nil
	}

	return f(ctx, entry, fields)
}

// PostWriteFunc adapts a function to PostWriteHook.
//
// A nil PostWriteFunc is a no-op hook.
type PostWriteFunc func(context.Context, core.Entry, []field.Field, WriteResult) error

// PostWrite calls f.
func (f PostWriteFunc) PostWrite(ctx context.Context, entry core.Entry, fields []field.Field, result WriteResult) error {
	if f == nil {
		return nil
	}

	return f(ctx, entry, fields, result)
}

// ErrorFunc adapts a function to ErrorHook.
//
// A nil ErrorFunc is a no-op hook.
type ErrorFunc func(context.Context, core.Entry, []field.Field, error)

// OnError calls f.
func (f ErrorFunc) OnError(ctx context.Context, entry core.Entry, fields []field.Field, err error) {
	if f == nil {
		return
	}

	f(ctx, entry, fields, err)
}
