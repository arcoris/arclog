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
// The function receives exactly the values supplied to PreWrite and controls the
// returned entry, returned field slice, and veto error. The adapter does not
// clone Entry, copy fields, recover panics, add synchronization, or inspect the
// context.
//
// A nil PreWriteFunc is a no-op hook. It returns the original entry and field
// slice unchanged with a nil error, which makes optional pre-write hooks easy to
// wire without nil-interface checks.
type PreWriteFunc func(context.Context, core.Entry, []field.Field) (core.Entry, []field.Field, error)

// PreWrite calls f or returns the input unchanged when f is nil.
func (f PreWriteFunc) PreWrite(ctx context.Context, entry core.Entry, fields []field.Field) (core.Entry, []field.Field, error) {
	if f == nil {
		return entry, fields, nil
	}

	return f(ctx, entry, fields)
}
