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

// Package clock defines API-side timestamp source contracts for arclog.
//
// The package is intentionally tiny. It lets runtime loggers, tests, and
// integrations share a timestamp source contract without coupling the API layer
// to time.Now, a fake clock implementation, or a concrete runtime configuration
// policy.
//
// # Responsibility boundary
//
// Clock values answer only one question: what timestamp should be attached now?
// They do not decide when a logger should request a timestamp, whether a
// timestamp should be omitted, how monotonic clock readings are interpreted, or
// how times are encoded. Timestamp formatting belongs to encoders. Clock
// selection belongs to runtime configuration or the user-facing facade.
//
// This package does not provide timers, tickers, sleeping, deadlines,
// scheduling, backoff, or fake-clock orchestration. Those concerns belong to
// runtime or test-support packages that depend on this API package.
//
// # Values and ownership
//
// time.Time is returned by value, so callers may retain a timestamp without
// coordinating with the Clock implementation. The clock package does not strip
// monotonic readings or force locations. Implementations may return a zero time
// when that is their documented policy; this package does not assign omission
// semantics to zero timestamps.
//
// # Concurrency
//
// Clock implementations are commonly shared by loggers and cores.
// Implementations should be safe for concurrent use unless they explicitly
// document a narrower contract. The Func adapter adds no synchronization around
// the wrapped function.
package clock
