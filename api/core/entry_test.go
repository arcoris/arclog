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

package core_test

import (
	"testing"
	"time"

	"arcoris.dev/arclog/api/caller"
	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/level"
	"arcoris.dev/arclog/api/stack"
)

func TestZeroEntry(t *testing.T) {
	t.Parallel()

	var entry core.Entry
	if !entry.IsZero() {
		t.Fatal("zero Entry is not zero")
	}
}

func TestEntryIsZeroReportsNonZeroFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		entry core.Entry
	}{
		{name: "time", entry: core.Entry{Time: time.Unix(1, 0)}},
		{name: "level", entry: core.Entry{Level: level.Trace}},
		{name: "logger", entry: core.Entry{LoggerName: "api"}},
		{name: "message", entry: core.Entry{Message: "hello"}},
		{name: "caller", entry: core.Entry{Caller: caller.Caller{Defined: true}}},
		{name: "stack", entry: core.Entry{Stack: stack.New([]caller.Caller{{Defined: true, File: "a.go", Line: 1}})}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.entry.IsZero() {
				t.Fatal("IsZero() = true, want false")
			}
		})
	}
}

func TestEntryCloneCopiesStackFrames(t *testing.T) {
	t.Parallel()

	entry := core.Entry{
		Stack: stack.New([]caller.Caller{
			{Defined: true, File: "a.go", Line: 1},
		}),
	}

	cloned := entry.Clone()
	clonedFrames := cloned.Stack.Frames()
	clonedFrames[0].Line = 2

	frame, ok := entry.Stack.Frame(0)
	if !ok {
		t.Fatal("original frame missing")
	}
	if frame.Line != 1 {
		t.Fatalf("original line = %d, want 1", frame.Line)
	}
}

func TestEntryCloneKeepsValueFields(t *testing.T) {
	t.Parallel()

	entry := core.Entry{
		Time:       time.Unix(10, 20),
		Level:      level.Error,
		LoggerName: "api",
		Message:    "failed",
		Caller:     caller.Caller{Defined: true, File: "a.go", Line: 1},
	}

	if cloned := entry.Clone(); !entriesEqual(cloned, entry) {
		t.Fatalf("Clone() = %#v, want %#v", cloned, entry)
	}
}
