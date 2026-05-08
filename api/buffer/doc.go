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

// Package buffer defines the low-level byte accumulation type used by ARCORIS
// logging encoders and logging cores.
//
// The package is intentionally small and allocation-aware. It provides a
// concrete *Buffer implementation backed by a growable byte slice, plus a Pool
// adapter that delegates reusable-object lifecycle to arcoris.dev/pool. The
// concrete type is intentional: the API avoids publishing broad buffer
// interfaces so future implementations do not freeze more method surface than
// encoder contracts actually need.
// Encoders use Buffer as scratch space while constructing serialized records
// such as JSON log lines.
//
// # Ownership model
//
// A Buffer has a single owner at any point in time. Code that obtains a buffer
// from a Pool owns that buffer until it returns the buffer with Free or Put.
// After release, the previous owner MUST NOT read from the buffer, write to the
// buffer, or retain slices returned by Bytes.
//
// Bytes returns a borrowed view of the internal byte slice. The view is valid
// only until the next write that may grow the buffer, Reset, Free, or Put. The
// caller MUST treat the returned slice as read-only.
//
// # Concurrency model
//
// Buffer values are not safe for concurrent use. Callers that share a buffer
// between goroutines MUST provide external synchronization around all buffer
// operations and all borrowed slices returned by Bytes.
//
// Pool values created with NewPool or NewPoolWithCapacity are safe for
// concurrent Get and Put calls. Pool uses arcoris.dev/pool for construction,
// reset-on-return, reuse admission, and backend storage. A Pool never makes an
// individual Buffer safe for concurrent mutation; it only coordinates reuse of
// different Buffer instances.
//
// # API boundary
//
// This package is part of the stable arclog API module because custom encoders,
// custom cores, and custom field marshalers need a shared byte accumulation
// type. Higher-level encoding policy, JSON escaping, field dispatch, writer
// synchronization, and size-bucketed pooling strategies belong outside this
// package.
//
// # Non-goals
//
// Buffer is not a drop-in replacement for bytes.Buffer. It intentionally omits
// APIs that are not needed by the logging pipeline and it exposes stricter
// lifetime rules because its values are commonly pooled.
package buffer
