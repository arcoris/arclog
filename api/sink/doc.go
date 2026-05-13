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

// Package sink defines contracts for consuming already encoded log payloads.
//
// A Sink receives bytes after a runtime encoder has finished turning a log
// record into its physical representation. That representation may be JSON,
// OTLP, a binary stream, or some later runtime-owned format; this package does
// not define or inspect it. The sink boundary is deliberately after encoding:
// sinks move bytes to files, sockets, exporters, queues, or test collectors,
// while encoders own field interpretation and serialization.
//
// The byte slice passed to Sink.Write is borrowed and may alias a pooled
// runtime buffer. Implementations must consume the slice before returning and
// must not retain, publish, or mutate it unless they first make their own copy.
// Asynchronous, batching, or queueing sinks must copy the slice before
// returning from Write.
//
// # Sync semantics
//
// Sink.Sync and Syncer.Sync flush or synchronize implementation-defined state.
// Sync may flush buffered bytes, forward to fsync, flush an exporter client,
// commit an in-memory test buffer, or do nothing when there is no state to
// flush. Durability is not implied by the interface; it belongs to the concrete
// implementation.
//
// # Resource ownership
//
// Sink does not define resource ownership. Closing files, network connections,
// exporter clients, and worker goroutines remains the responsibility of the
// component that created and owns those resources.
//
// # Non-goals
//
// Concrete wrappers such as locking, fanout, discard, retry, sampling, and
// add-sync belong to runtime packages, not api/sink.
package sink
