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

package buffer_test

import (
	"testing"

	"arcoris.dev/arclog/api/buffer"
)

func TestBufferStateAndReset(t *testing.T) {
	buf := newBuffer()

	if got := buf.Len(); got != 0 {
		t.Fatalf("initial Len = %d, want 0", got)
	}
	if got := buf.Cap(); got != 0 {
		t.Fatalf("initial Cap = %d, want 0", got)
	}

	buf.AppendString("hello")
	if got, want := buf.Len(), len("hello"); got != want {
		t.Fatalf("Len after AppendString = %d, want %d", got, want)
	}
	if got, want := string(buf.Bytes()), "hello"; got != want {
		t.Fatalf("Bytes after AppendString = %q, want %q", got, want)
	}

	capacity := buf.Cap()
	if capacity < buf.Len() {
		t.Fatalf("Cap = %d, want >= Len %d", capacity, buf.Len())
	}

	buf.Reset()
	if got := buf.Len(); got != 0 {
		t.Fatalf("Len after Reset = %d, want 0", got)
	}
	if got := buf.Cap(); got != capacity {
		t.Fatalf("Cap after Reset = %d, want %d", got, capacity)
	}
	if got := string(buf.Bytes()); got != "" {
		t.Fatalf("Bytes after Reset = %q, want empty", got)
	}
}

func TestBufferNewAndString(t *testing.T) {
	buf := buffer.New(8)

	if got := buf.Cap(); got < 8 {
		t.Fatalf("New cap = %d, want >= 8", got)
	}

	buf.AppendString("hello")
	if got := buf.String(); got != "hello" {
		t.Fatalf("String = %q, want %q", got, "hello")
	}
}

func TestNewNegativeCapacity(t *testing.T) {
	t.Parallel()

	buf := buffer.New(-1)

	if got := buf.Cap(); got != 0 {
		t.Fatalf("Cap = %d, want 0", got)
	}
}

func TestZeroValueBuffer(t *testing.T) {
	t.Parallel()

	var buf buffer.Buffer
	if got := buf.Len(); got != 0 {
		t.Fatalf("Len = %d, want 0", got)
	}
	if got := buf.String(); got != "" {
		t.Fatalf("String = %q, want empty", got)
	}

	buf.AppendString("ready")
	if got := buf.String(); got != "ready" {
		t.Fatalf("String after AppendString = %q, want %q", got, "ready")
	}
}

func TestBytesAliasing(t *testing.T) {
	t.Parallel()

	buf := buffer.New(0)
	buf.AppendString("abc")

	view := buf.Bytes()
	buf.AppendString("d")

	if got, want := string(view), "abc"; got != want {
		t.Fatalf("borrowed view = %q, want %q", got, want)
	}
	if got, want := string(buf.Bytes()), "abcd"; got != want {
		t.Fatalf("buffer bytes = %q, want %q", got, want)
	}
}
