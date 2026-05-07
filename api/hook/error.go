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

// ErrorHook observes a write failure.
//
// ErrorHook does not return an error. Runtime managers should use their own
// internal error reporting policy if an ErrorHook panics or fails indirectly.
// Keeping this contract return-free avoids recursive "error while handling
// error" semantics in the API layer.
type ErrorHook interface {
	// OnError observes a write failure.
	OnError(ctx context.Context, entry core.Entry, fields []field.Field, err error)
}
