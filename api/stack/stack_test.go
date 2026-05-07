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

package stack_test

import (
	"testing"

	"arcoris.dev/arclog/api/caller"
	"arcoris.dev/arclog/api/stack"
)

func TestZeroValueStack(t *testing.T) {
	t.Parallel()

	var s stack.Stack

	if !s.IsEmpty() {
		t.Fatal("zero-value Stack is not empty")
	}
	if s.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", s.Len())
	}
	if got := s.Frames(); got != nil {
		t.Fatalf("Frames() = %#v, want nil", got)
	}
}

func TestNewRetainsFrames(t *testing.T) {
	t.Parallel()

	frames := []caller.Caller{
		{Defined: true, File: "logger.go", Line: 10, Function: "log"},
	}
	s := stack.New(frames)

	frames[0].Line = 11

	got, ok := s.Frame(0)
	if !ok {
		t.Fatal("Frame(0) returned ok=false")
	}
	if got.Line != 11 {
		t.Fatalf("Frame(0).Line = %d, want retained slice mutation to be visible", got.Line)
	}
}

func TestCloneCopiesFrames(t *testing.T) {
	t.Parallel()

	frames := []caller.Caller{
		{Defined: true, File: "logger.go", Line: 10, Function: "log"},
	}
	s := stack.Clone(frames)

	frames[0].Line = 11

	got, ok := s.Frame(0)
	if !ok {
		t.Fatal("Frame(0) returned ok=false")
	}
	if got.Line != 10 {
		t.Fatalf("Frame(0).Line = %d, want copied value 10", got.Line)
	}
}

func TestFrameBounds(t *testing.T) {
	t.Parallel()

	s := stack.New([]caller.Caller{
		{Defined: true, File: "a.go", Line: 1},
	})

	if _, ok := s.Frame(-1); ok {
		t.Fatal("Frame(-1) returned ok=true")
	}
	if _, ok := s.Frame(1); ok {
		t.Fatal("Frame(1) returned ok=true")
	}
}

func TestFramesReturnsRetainedSlice(t *testing.T) {
	t.Parallel()

	frames := []caller.Caller{
		{Defined: true, File: "a.go", Line: 1},
		{Defined: true, File: "b.go", Line: 2},
	}
	s := stack.New(frames)

	got := s.Frames()
	if len(got) != len(frames) {
		t.Fatalf("len(Frames()) = %d, want %d", len(got), len(frames))
	}
	got[1].Line = 22

	frame, ok := s.Frame(1)
	if !ok {
		t.Fatal("Frame(1) returned ok=false")
	}
	if frame.Line != 22 {
		t.Fatalf("Frame(1).Line = %d, want mutation through retained slice to be visible", frame.Line)
	}
}
