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

// Package writer defines API-side sink contracts for encoded log records.
//
// The package is deliberately small. It describes where already-encoded bytes
// can be written and how buffered sink state can be synchronized; it does not
// encode log entries, manage log levels, dispatch fields, or own writer
// lifecycle.
//
// # Responsibility boundary
//
// WriteSyncer receives bytes after the encoding layer has already produced a
// complete record or record fragment. The writer package does not define record
// framing, newline policy, JSON validity, field ordering, retry behavior,
// buffering strategy, file rotation, locking, fan-out, sampling, or backpressure.
// Those concerns belong to runtime packages or application-specific sinks that
// implement these contracts.
//
// # Sync semantics
//
// Syncer is intentionally weaker than a universal fsync guarantee. A concrete
// implementation defines what synchronization means for its sink: flushing a
// buffered writer, calling fsync on a file, flushing a network client, or doing
// nothing when there is no buffered state to flush. Implementations SHOULD make
// Sync safe to call multiple times.
//
// # Concurrency
//
// The interfaces in this package do not imply concurrency safety. A concrete
// sink implementation MUST document whether Write and Sync may be called
// concurrently. Runtime wrappers such as locked writers or multi-writers should
// live outside this API package.
//
// # Non-goals
//
// This package intentionally does not provide helper implementations such as
// AddSync, Lock, MultiWriteSyncer, or NopSyncer. Those helpers are useful, but
// they are runtime/facade conveniences rather than stable API contracts.
package writer
