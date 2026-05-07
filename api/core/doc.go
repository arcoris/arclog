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

// Package core defines the central write-pipeline contracts for arclog.
//
// The package is the API boundary between logger orchestration and concrete
// output implementations. It defines the metadata carried by a log entry, the
// Core interface that accepts entries, CheckedEntry as the result of the check
// phase, and small pure core combinators such as Tee and Noop.
//
// # Responsibility boundary
//
// Core contracts operate on Entry and []field.Field. They do not define JSON,
// console, or binary encoding; they do not own writer adapters; they do not run
// hooks or predicates; they do not capture callers or stack traces; and they do
// not implement logger facade methods such as Info or Error. Those concerns
// belong to runtime packages that depend on this API package.
//
// # Check and write lifecycle
//
// A Core first participates in the check phase. If it will write an entry, it
// adds itself to a CheckedEntry. If it will not write, it returns the CheckedEntry
// unchanged. The write phase must not repeat check logic: when Core.Write is
// called, the core is expected to write the entry or return an error.
//
// # Pure primitives
//
// This package includes only primitives that are independent from concrete
// encoders, buffers, writers, and runtime configuration. Noop is a no-op Core.
// Tee composes multiple Core values. I/O cores, sampled cores, async cores,
// hook managers, and encoder-backed cores belong outside this API package.
//
// # Ownership
//
// Entry values are cheap transport values. Field slices passed to Core methods
// must be treated as read-only unless the implementation owns the complete
// pipeline and documents a stronger contract. Implementations that retain fields
// after a call returns must copy the slice.
package core
