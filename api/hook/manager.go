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
// Manager is an interface only. This package does not define storage,
// synchronization, ordering implementation, asynchronous delivery, panic
// recovery, retry behavior, or shutdown workers.
type Manager interface {
	// AddPreWrite registers a pre-write hook.
	AddPreWrite(PreWriteHook, Priority) Registration

	// AddPostWrite registers a post-write hook.
	AddPostWrite(PostWriteHook, Priority) Registration

	// AddError registers a write-error hook.
	AddError(ErrorHook, Priority) Registration

	// FirePreWrite runs the pre-write phase.
	FirePreWrite(context.Context, core.Entry, []field.Field) (core.Entry, []field.Field, error)

	// FirePostWrite runs the post-write phase.
	FirePostWrite(context.Context, core.Entry, []field.Field, WriteResult) error

	// FireError runs the write-error phase.
	FireError(context.Context, core.Entry, []field.Field, error)

	// Stop releases manager resources.
	//
	// Managers without background resources may return nil immediately.
	Stop(context.Context) error
}
