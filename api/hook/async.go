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
	"arcoris.dev/arclog/api/internal/nilx"
)

// AsyncHook marks a hook that may be executed asynchronously by a runtime
// manager.
//
// AsyncHook is a capability declaration only. The API package does not start
// goroutines, allocate queues, copy entries, or provide delivery guarantees.
// Runtime managers decide whether to honor this preference.
//
// Hooks that may run asynchronously must assume that the caller-owned entry and
// field slice stop being valid after the dispatch call unless the runtime
// manager explicitly documents that it cloned them. Implementations should treat
// Async as a stable capability declaration, not as per-entry dynamic policy.
type AsyncHook interface {
	// Async reports whether the hook may be executed asynchronously.
	Async() bool
}

// AllowsAsync reports whether hook declares asynchronous execution support.
//
// Nil values, typed nil values, and hooks that do not implement AsyncHook are
// treated as synchronous-only. AllowsAsync does not recover panics from Async;
// the Async method is part of the hook's own contract.
func AllowsAsync(h any) bool {
	if nilx.IsNil(h) {
		return false
	}

	async, ok := h.(AsyncHook)
	return ok && async.Async()
}
