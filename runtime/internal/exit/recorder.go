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

package exit

import "sync"

// Recorder records exit attempts while installed by Stub.
//
// Recorder is safe for concurrent observation and for exit attempts from
// goroutines. Installing and restoring a recorder still mutates package-global
// exit state, so tests using Stub or WithStub must not run in parallel with
// other tests that depend on this package's exit function.
type Recorder struct {
	// mu protects all mutable recorder state below.
	//
	// The recorder can be called by With from goroutines while the test goroutine
	// reads Exited, Code, Calls, or Restored. Keeping this state private and
	// accessor-based prevents data races that would be possible with exported
	// fields.
	mu sync.Mutex

	// previous is the exit function that was active before this recorder was
	// installed.
	//
	// Unstub restores previous exactly once. It is nil only for a malformed
	// Recorder value constructed outside this package, which callers cannot do
	// directly because the fields are unexported.
	previous func(int)

	// exited records whether at least one exit attempt was observed.
	exited bool

	// code records the most recent code passed to With.
	//
	// Code intentionally stores the latest value rather than the first one because
	// repeated exit attempts indicate a broken terminal path and the last attempt
	// is usually the most useful diagnostic state.
	code int

	// calls records how many exit attempts were observed.
	//
	// This is useful for tests that need to assert that a fatal path terminates
	// exactly once.
	calls int

	// restored records whether Unstub has already restored previous.
	//
	// The flag makes Unstub idempotent and prevents a second Unstub call from
	// replacing a newer recorder or production exit function with stale state.
	restored bool
}

// Stub installs a Recorder and returns it.
//
// The returned Recorder remains installed until Unstub is called. Stub replaces
// package-global state and therefore must not be used concurrently with another
// Stub or WithStub call. Use t.Cleanup or defer to restore the previous exit
// function even when a test fails.
func Stub() *Recorder {
	stateMu.Lock()
	defer stateMu.Unlock()

	recorder := &Recorder{
		previous: exitFn,
	}
	exitFn = recorder.record
	return recorder
}

// WithStub runs f with process termination recorded instead of executed.
//
// WithStub restores the previous exit function before returning. If f panics,
// the previous exit function is restored before the panic continues. A nil f is
// treated as a no-op.
//
// WithStub is convenient for tests that need a short scoped replacement. Tests
// that need to inspect state before restoration should use Stub and call Unstub
// manually.
func WithStub(f func()) *Recorder {
	recorder := Stub()
	defer recorder.Unstub()

	if f != nil {
		f()
	}

	return recorder
}

// Unstub restores the exit function that was active before Stub installed r.
//
// Unstub is idempotent. Calling Unstub on a nil Recorder is a no-op. If another
// recorder was installed after r without first restoring r, Unstub will restore
// r's previous function and therefore overwrite that newer installation; tests
// must serialize use of this package-level seam.
func (r *Recorder) Unstub() {
	if r == nil {
		return
	}

	stateMu.Lock()
	defer stateMu.Unlock()

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.restored {
		return
	}

	exitFn = r.previous
	r.restored = true
}

// Exited reports whether With was called while r was installed.
func (r *Recorder) Exited() bool {
	if r == nil {
		return false
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	return r.exited
}

// Code returns the most recent exit code recorded by r.
//
// Code returns 0 when r is nil or no exit attempt has been recorded.
func (r *Recorder) Code() int {
	if r == nil {
		return 0
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	return r.code
}

// Calls reports how many exit attempts r recorded.
//
// Calls returns 0 for a nil Recorder.
func (r *Recorder) Calls() int {
	if r == nil {
		return 0
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	return r.calls
}

// Restored reports whether r has restored the previous exit function.
//
// Restored is primarily useful for tests that assert cleanup behavior.
func (r *Recorder) Restored() bool {
	if r == nil {
		return false
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	return r.restored
}

// record stores one exit attempt.
//
// record is installed as the package exit function while a Recorder is active.
// It deliberately does not call os.Exit. The method is private because only
// Stub should install it.
func (r *Recorder) record(code int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.exited = true
	r.code = code
	r.calls++
}
