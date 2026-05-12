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

// Package bufferpool provides the runtime buffer pool used by encoder sessions.
//
// The package owns reuse and retention policy for api/buffer.Buffer values.
// api/buffer.Buffer is only a raw byte accumulator; it does not own pooling,
// sink writing, encoding policy, or session lifecycle. Runtime sessions acquire
// buffers from Pool, write encoded bytes into them, pass the borrowed bytes to
// a writer sink, and return the buffer after the sink call returns.
//
// A buffer returned by Pool.Get is owned by the caller until Pool.Put. Calling
// Pool.Put ends that ownership and invalidates the buffer and all byte slices
// previously returned by Buffer.Bytes.
//
// Pools created with New or NewWithOptions are safe for concurrent Get and Put
// calls. The zero value is safe but non-pooling, and exists only as a fallback
// for incomplete internal wiring rather than as the recommended hot-path
// configuration. Individual Buffer values remain single-owner mutable objects
// and are not safe for concurrent mutation.
//
// Pool retains buffers whose capacity is at or below MaxRetainedCapacity.
// Oversized buffers are reset and discarded instead of being returned to the
// backend pool so unusually large records do not pin large backing arrays in
// steady-state runtime reuse.
//
// This package does not own JSON encoding, sink behavior, asynchronous
// retention, exporter queues, or session lifecycle semantics. It is an
// internal runtime detail and not a public API.
package bufferpool
