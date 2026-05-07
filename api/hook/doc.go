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

// Package hook defines API-side contracts for transforming and observing log
// entries during the write pipeline.
//
// Hooks operate after a log entry has been admitted into the runtime pipeline.
// Predicates decide whether an entry should continue through the pipeline;
// hooks may enrich, redact, observe, or react to entries that are already being
// processed.
//
// # Responsibility boundary
//
// This package defines contracts, value types, and function adapters only. It
// does not store hooks, sort hooks, start goroutines, recover panics, retry
// writes, batch events, own shutdown workers, or implement a global registry.
// Concrete hook managers and runtime policy live outside api.
//
// Hooks do not encode entries and do not write to byte sinks directly. Encoding
// belongs to encoder implementations. Byte I/O belongs to core and writer
// implementations.
//
// # Phases
//
// PreWriteHook is the transformation phase. It may return a modified entry,
// return a modified field slice, or veto the write by returning an error.
//
// PostWriteHook is the observation phase after a write attempt. It observes the
// entry, fields, and write result, but cannot modify bytes that have already
// been written.
//
// ErrorHook is the failure reaction phase. It observes write failures and does
// not return another error, avoiding recursive error handling in the API
// contract.
//
// Manager is an optional orchestration contract for runtime components that want
// to expose hook registration and phase dispatch through API types. It is not a
// manager implementation and does not define storage, ordering data structures,
// background workers, panic policy, or shutdown behavior.
//
// # Ownership
//
// Entry values are borrowed from the caller. Field slices are borrowed for the
// duration of the hook call. Hook implementations must not retain entries or
// field slices beyond the call unless they clone the entry and copy the field
// slice first. Field payload ownership remains governed by the field package; a
// copied field slice can still contain caller-owned payload values such as byte
// slices, errors, stringers, or marshalers.
//
// # Concurrency
//
// Hooks are commonly shared by loggers, cores, and runtime managers. Hook
// implementations should be safe for concurrent use unless they explicitly
// document a narrower contract. The function adapters in this package add no
// synchronization around wrapped functions.
package hook
