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

// WriteResult describes the result observed by PostWriteHook.
//
// The API-level result intentionally contains only the write error. Concrete
// runtime implementations may track byte counts, sink names, retries, or other
// diagnostics internally, but core.Core.Write itself only returns error.
type WriteResult struct {
	// Err is the error returned by the write attempt.
	Err error
}

// Success returns a successful write result.
func Success() WriteResult {
	return WriteResult{}
}

// Failure returns a failed write result.
func Failure(err error) WriteResult {
	return WriteResult{Err: err}
}

// Failed reports whether r represents a failed write attempt.
func (r WriteResult) Failed() bool {
	return r.Err != nil
}

// PostWriteHook observes a log entry after a write attempt.
//
// PostWrite may report an observer error to the runtime manager. It cannot
// change the already-attempted write and should treat entry and fields as
// read-only borrowed values.
type PostWriteHook interface {
	// PostWrite observes entry, fields, and the write result.
	PostWrite(ctx context.Context, entry core.Entry, fields []field.Field, result WriteResult) error
}
