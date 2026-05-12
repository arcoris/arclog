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

// Package sync defines the Sink contract for encoded log payloads.
//
// A Sink receives bytes that have already been encoded by an arclog runtime
// encoder. The byte slice passed to Sink.Write is borrowed and may alias a
// pooled runtime buffer. Implementations must consume the slice before returning
// and must not retain, publish, or mutate it unless they first make their own
// copy. Asynchronous or queueing sinks must copy the slice before returning from
// Write.
//
// # Sync semantics
//
// Sink.Sync and Syncer.Sync flush or synchronize implementation-defined state.
// Sync may flush buffered bytes, forward to fsync, flush an exporter client, or
// do nothing when there is no buffered state to flush.
//
// # Resource ownership
//
// Sink does not define resource ownership. Closing files, network connections,
// or exporter clients remains the responsibility of the code that owns those
// resources.
//
// # Non-goals
//
// Concrete wrappers such as locking, multi-sink fanout, discard, and add-sync
// belong to runtime packages, not api/sync.
package sync
