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

// PostWriteFunc adapts a function to PostWriteHook.
//
// The adapter passes through the borrowed entry, borrowed field slice, and
// WriteResult without modification. It does not recover panics, copy values, add
// synchronization, or reinterpret write errors.
//
// A nil PostWriteFunc is a no-op observer and returns nil.
type PostWriteFunc func(context.Context, core.Entry, []field.Field, WriteResult) error

// PostWrite calls f or returns nil when f is nil.
func (f PostWriteFunc) PostWrite(ctx context.Context, entry core.Entry, fields []field.Field, result WriteResult) error {
	if f == nil {
		return nil
	}

	return f(ctx, entry, fields, result)
}
