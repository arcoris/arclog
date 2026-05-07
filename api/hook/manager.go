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

// Manager is the API-side contract for registering and firing hooks.
//
// Manager is an interface only. It defines the shape that runtime hook
// orchestrators may expose to logger, core, or integration code. This package
// does not define storage, synchronization, priority sorting, asynchronous
// delivery, panic recovery, retry behavior, batching, or shutdown workers.
//
// Implementations should be safe for concurrent use unless they document a
// narrower contract. Registered hooks receive borrowed Entry and field slices
// according to the phase-specific contracts.
type Manager interface {
	// AddPreWrite registers a pre-write hook.
	//
	// The returned Registration removes this specific registration. Passing nil
	// is implementation-defined; runtimes may reject it, ignore it, or register a
	// no-op.
	AddPreWrite(PreWriteHook, Priority) Registration

	// AddPostWrite registers a post-write hook.
	//
	// The returned Registration removes this specific registration. Priority is a
	// relative ordering hint for hooks in the post-write phase.
	AddPostWrite(PostWriteHook, Priority) Registration

	// AddError registers a write-error hook.
	//
	// The returned Registration removes this specific registration. Error hooks
	// observe failures; they do not replace the original write error.
	AddError(ErrorHook, Priority) Registration

	// FirePreWrite runs the pre-write phase.
	//
	// Implementations should return the authoritative entry and field slice that
	// later phases must continue with. A non-nil error represents the runtime's
	// chosen pre-write veto or transformation failure.
	FirePreWrite(context.Context, core.Entry, []field.Field) (core.Entry, []field.Field, error)

	// FirePostWrite runs the post-write phase.
	//
	// Implementations should return an observer error only for hook execution
	// failure. The WriteResult argument remains the write outcome being observed.
	FirePostWrite(context.Context, core.Entry, []field.Field, WriteResult) error

	// FireError runs the write-error phase.
	//
	// FireError intentionally has no return value to avoid recursive error
	// handling in the public contract.
	FireError(context.Context, core.Entry, []field.Field, error)

	// Stop releases manager resources.
	//
	// Managers without background resources may return nil immediately.
	Stop(context.Context) error
}
