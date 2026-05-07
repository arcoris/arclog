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

// ErrorHook observes a write failure after the runtime write path reports it.
//
// ErrorHook does not return an error. Runtime managers should use their own
// internal error reporting policy if an ErrorHook panics or fails indirectly.
// Keeping this contract return-free avoids recursive "error while handling
// error" semantics in the API layer.
//
// Entry and fields are borrowed for the duration of the call. Implementations
// that retain either value must clone Entry when stack ownership is uncertain
// and must copy the field slice.
type ErrorHook interface {
	// OnError observes a write failure.
	//
	// err may be nil only if a runtime manager chooses to signal an unusual
	// condition without a concrete write error. Implementations that require a
	// concrete error should check err before using it.
	OnError(ctx context.Context, entry core.Entry, fields []field.Field, err error)
}
