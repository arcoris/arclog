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

func TestCapturerFuncCaptureStack(t *testing.T) {
	t.Parallel()

	capturer := stack.CapturerFunc(func(skip int) stack.Stack {
		if skip != 2 {
			t.Fatalf("skip = %d, want 2", skip)
		}

		return stack.New([]caller.Caller{
			{Defined: true, File: "entry.go", Line: 42, Function: "write"},
		})
	})

	var _ stack.Capturer = capturer

	got := capturer.CaptureStack(2)
	if got.Len() != 1 {
		t.Fatalf("Len() = %d, want 1", got.Len())
	}
}

func TestNilCapturerFuncPanics(t *testing.T) {
	t.Parallel()

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("CaptureStack() did not panic for nil CapturerFunc")
		}
	}()

	var capturer stack.CapturerFunc
	_ = capturer.CaptureStack(0)
}
