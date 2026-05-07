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

import "arcoris.dev/arclog/api/caller"

// Stack is a stack trace represented as an ordered sequence of call-site frames.
//
// The zero value is an empty stack. An empty Stack is safe to pass through the
// logging pipeline and normally means that stack capture was disabled,
// unavailable, or not requested for the entry.
//
// Stack is immutable by convention, not by enforcement. The frame slice returned
// by Frames aliases the slice retained by the Stack value. Code that constructs
// a Stack with New must not mutate the supplied slice while the Stack may still
// be observed by encoders, hooks, cores, or asynchronous sinks.
type Stack struct {
	frames []caller.Caller
}

// New creates a Stack backed by frames.
//
// New retains frames without copying it. This is intentional for
// allocation-sensitive capture paths. The caller remains responsible for
// keeping the slice stable for as long as the Stack may be observed. Use Clone
// when the input slice is temporary, reused, or may be mutated by the caller.
func New(frames []caller.Caller) Stack {
	if len(frames) == 0 {
		return Stack{}
	}

	return Stack{frames: frames}
}

// Clone creates a Stack backed by a newly allocated copy of frames.
//
// Clone is useful when the caller cannot guarantee that the input slice will
// remain unchanged for the lifetime of the resulting Stack. This is the safer
// construction path for asynchronous hooks, deferred sinks, retained entries,
// or tests that intentionally mutate the original slice after construction.
func Clone(frames []caller.Caller) Stack {
	if len(frames) == 0 {
		return Stack{}
	}

	copied := make([]caller.Caller, len(frames))
	copy(copied, frames)
	return Stack{frames: copied}
}

// Len reports the number of frames in s.
func (s Stack) Len() int {
	return len(s.frames)
}

// IsEmpty reports whether s contains no frames.
func (s Stack) IsEmpty() bool {
	return len(s.frames) == 0
}

// Frames returns the ordered frames retained by s.
//
// The returned slice aliases s. Callers must treat it as read-only unless they
// own the Stack and can prove that no encoder, hook, core, asynchronous sink, or
// retained entry can observe the mutation.
func (s Stack) Frames() []caller.Caller {
	return s.frames
}

// Frame returns the frame at index i.
//
// The second result is false when i is outside the range [0, Len()).
func (s Stack) Frame(i int) (caller.Caller, bool) {
	if i < 0 || i >= len(s.frames) {
		return caller.Caller{}, false
	}

	return s.frames[i], true
}
