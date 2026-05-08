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
// The package provides the concrete *Buffer type shared by encoder-facing API
// contracts. The concrete type is intentional: the API avoids publishing broad
// buffer interfaces so future implementations do not freeze more method surface
// than encoder contracts actually need.
//
// # Ownership model
//
// A Buffer has a single owner at any point in time. Methods mutate the buffer's
// internal byte slice directly, and Bytes returns a borrowed view of that slice.
// The view is valid only until the next write that may grow the buffer or until
// Reset. Callers MUST treat the returned slice as read-only and must not retain
// it across later buffer mutation.
//
// Reset clears the logical contents while retaining capacity for reuse by the
// current owner. Reset does not release memory, return the buffer to a pool, or
// apply any retention policy.
//
// # Concurrency model
//
// Buffer values are not safe for concurrent use. Callers that share a buffer
// between goroutines MUST provide external synchronization around all buffer
// operations and all borrowed slices returned by Bytes.
//
// # API boundary
//
// This package is part of the stable arclog API module because custom encoders,
// custom cores, and custom field marshalers need a shared byte accumulation
// type. Pooling, object lifecycle, capacity retention limits, runtime default
// capacities, JSON escaping, field dispatch, writer synchronization, and
// concrete encoder policy belong outside this package.
//
// Runtime packages that want pooling should build it around *Buffer explicitly,
// for example by acquiring a buffer from a runtime-owned pool, using Reset on
// return, and applying retention policy in that runtime package.
//
// # Non-goals
//
// Buffer is not a drop-in replacement for bytes.Buffer. It intentionally omits
// APIs that are not needed by the logging pipeline and keeps lifecycle policy
// outside the API layer.
package buffer
