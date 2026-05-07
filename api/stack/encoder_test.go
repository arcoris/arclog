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

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/caller"
	"arcoris.dev/arclog/api/stack"
)

func TestEncoder(t *testing.T) {
	t.Parallel()

	enc := stack.Encoder(func(dst *buffer.Buffer, s stack.Stack) *buffer.Buffer {
		for i, frame := range s.Frames() {
			if i > 0 {
				dst.AppendByte('\n')
			}
			dst.AppendString(frame.File)
			dst.AppendByte(':')
			dst.AppendInt(frame.Line)
		}
		return dst
	})

	dst := buffer.New(0)
	got := enc(dst, stack.New([]caller.Caller{
		{Defined: true, File: "a.go", Line: 1},
		{Defined: true, File: "b.go", Line: 2},
	}))

	if got != dst {
		t.Fatal("Encoder returned a different buffer")
	}
	if got.String() != "a.go:1\nb.go:2" {
		t.Fatalf("encoded stack = %q, want %q", got.String(), "a.go:1\nb.go:2")
	}
}
