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

// PostWriteHook observes a log entry after a write attempt.
//
// PostWrite runs after a core write attempt has completed. It may observe
// success or failure and may return an observer error to the runtime manager,
// but it cannot change bytes that were already written or retry the write
// through this API contract.
//
// Entry and fields are borrowed for the duration of the call. Implementations
// that retain either value must clone Entry when stack ownership is uncertain
// and must copy the field slice.
type PostWriteHook interface {
	// PostWrite observes entry, fields, and result.
	//
	// The context is supplied by the runtime pipeline. PostWriteHook
	// implementations should respect cancellation for expensive observer work,
	// but this package does not enforce cancellation, deadlines, retry behavior,
	// or error aggregation.
	PostWrite(ctx context.Context, entry core.Entry, fields []field.Field, result WriteResult) error
}
