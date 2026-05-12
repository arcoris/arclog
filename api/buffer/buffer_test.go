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

package buffer

import (
	"testing"
)

func TestZeroValueBufferIsUsable(t *testing.T) {
	t.Parallel()

	var buf Buffer
	buf.AppendString("x")

	if got, want := string(buf.Bytes()), "x"; got != want {
		t.Fatalf("Bytes = %q, want %q", got, want)
	}
}

func TestNewZeroCapacity(t *testing.T) {
	t.Parallel()

	buf := New(0)

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len = %d, want 0", got)
	}
	if got := buf.Cap(); got != 0 {
		t.Fatalf("Cap = %d, want 0", got)
	}
}

func TestNewPositiveCapacity(t *testing.T) {
	t.Parallel()

	buf := New(32)

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len = %d, want 0", got)
	}
	if got := buf.Cap(); got < 32 {
		t.Fatalf("Cap = %d, want >= 32", got)
	}
}

func TestNewNegativeCapacityPanics(t *testing.T) {
	t.Parallel()

	assertPanicMessage(t, "buffer: negative capacity", func() {
		_ = New(-1)
	})
}

func assertPanicMessage(t *testing.T, want string, fn func()) {
	t.Helper()

	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatalf("panic = nil, want %q", want)
		}

		got, ok := recovered.(string)
		if !ok {
			t.Fatalf("panic type = %T, want string", recovered)
		}
		if got != want {
			t.Fatalf("panic = %q, want %q", got, want)
		}
	}()

	fn()
}
