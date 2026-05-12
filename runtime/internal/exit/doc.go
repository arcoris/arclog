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

// Package exit isolates process termination behind a small test seam.
//
// Production code should call With instead of os.Exit directly when process
// termination is part of logging behavior. Tests can install a Recorder with
// Stub or WithStub to observe exit attempts without terminating the test
// process.
//
// The package intentionally follows the same broad idea as the small exit seam
// used by established logging projects: process termination is process-global
// and hard to test, so production behavior is represented by one replaceable
// function. ARCORIS keeps that idea narrow and internal instead of exposing
// fatal behavior through API/core contracts.
//
// # Responsibility boundary
//
// This package owns only process-termination indirection. It does not decide:
//
//   - when a fatal log entry should terminate the process;
//   - whether a logger should call Sync before termination;
//   - whether panic-level entries should panic, exit, or both;
//   - how terminal actions are configured;
//   - how write errors affect fatal behavior.
//
// Those policies belong to the logger runtime or to the user-facing facade.
// API packages must not import this package.
//
// # Global state
//
// Stub and WithStub temporarily replace package-global process termination
// behavior. Tests that use them must not run concurrently with other tests that
// depend on this package's exit function.
//
// Recorder state is protected by a mutex so exit attempts from goroutines can
// be observed without data races. The package-global replacement itself is still
// a process-wide test seam and should be treated as exclusive test state.
//
// # Exit codes
//
// Success and Failure provide the only conventional codes owned by this package.
// More specific terminal policy, such as panic handling, should be introduced
// only when the runtime logger actually needs it.
//
// This package is internal to arclog and is not part of the public API.
package exit
