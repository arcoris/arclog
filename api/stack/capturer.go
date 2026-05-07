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

package stack

// Capturer captures stack trace metadata for the current execution context.
//
// Capturer is an API contract, not a runtime stack-walking implementation.
// Concrete runtime packages may implement it with runtime.Callers,
// preallocated frame buffers, platform-specific symbolization, deterministic
// test doubles, or another capture strategy.
//
// Unless a concrete implementation documents otherwise, Capturer values should
// be safe for concurrent use because loggers typically capture stack traces from
// many goroutines.
type Capturer interface {
	// CaptureStack returns a stack trace for the current execution context.
	//
	// The skip parameter follows the conventional runtime.Callers-style idea:
	// skip == 0 refers to the frame of CaptureStack itself, skip == 1 refers to
	// its caller, and higher values walk further up the stack. Implementations
	// may adjust skip internally to hide wrapper frames, but they should document
	// that policy.
	CaptureStack(skip int) Stack
}

// CapturerFunc adapts a function to Capturer.
//
// A nil CapturerFunc is invalid and will panic when CaptureStack is called.
type CapturerFunc func(skip int) Stack

// CaptureStack calls f(skip).
func (f CapturerFunc) CaptureStack(skip int) Stack {
	return f(skip)
}
