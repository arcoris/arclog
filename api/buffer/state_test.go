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

func TestLenAndCap(t *testing.T) {
	t.Parallel()

	buf := New(4)
	buf.AppendString("hello")

	if got, want := buf.Len(), 5; got != want {
		t.Fatalf("Len = %d, want %d", got, want)
	}
	if got := buf.Cap(); got < 5 {
		t.Fatalf("Cap = %d, want >= 5", got)
	}
}

func TestBytesReturnsCurrentContents(t *testing.T) {
	t.Parallel()

	var buf Buffer
	buf.AppendString("hello")
	buf.AppendByte(' ')
	buf.AppendBytes([]byte("world"))

	if got, want := string(buf.Bytes()), "hello world"; got != want {
		t.Fatalf("Bytes = %q, want %q", got, want)
	}
}

func TestResetClearsLengthAndRetainsCapacity(t *testing.T) {
	t.Parallel()

	buf := New(8)
	buf.AppendString("payload")

	capacity := buf.Cap()
	buf.Reset()

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len after Reset = %d, want 0", got)
	}
	if got := buf.Cap(); got != capacity {
		t.Fatalf("Cap after Reset = %d, want %d", got, capacity)
	}
	if got := len(buf.Bytes()); got != 0 {
		t.Fatalf("len(Bytes()) after Reset = %d, want 0", got)
	}
}

func TestGrowEnoughCapacityDoesNotChangeContents(t *testing.T) {
	t.Parallel()

	buf := New(16)
	buf.AppendString("hello")

	capacity := buf.Cap()
	buf.Grow(3)

	if got, want := string(buf.Bytes()), "hello"; got != want {
		t.Fatalf("Bytes after Grow = %q, want %q", got, want)
	}
	if got := buf.Cap(); got != capacity {
		t.Fatalf("Cap after Grow = %d, want %d", got, capacity)
	}
}

func TestGrowNeedsAllocationPreservesContentsAndIncreasesSpareCapacity(t *testing.T) {
	t.Parallel()

	buf := New(4)
	buf.AppendString("abcd")

	buf.Grow(10)

	if got, want := string(buf.Bytes()), "abcd"; got != want {
		t.Fatalf("Bytes after Grow = %q, want %q", got, want)
	}
	if got := buf.Cap() - buf.Len(); got < 10 {
		t.Fatalf("spare capacity after Grow = %d, want >= 10", got)
	}
}

func TestGrowNegativePanics(t *testing.T) {
	t.Parallel()

	var buf Buffer
	assertPanicMessage(t, "buffer: negative grow", func() {
		buf.Grow(-1)
	})
}

func TestTruncateShortensContents(t *testing.T) {
	t.Parallel()

	var buf Buffer
	buf.AppendString("hello")

	capacity := buf.Cap()
	buf.Truncate(3)

	if got, want := string(buf.Bytes()), "hel"; got != want {
		t.Fatalf("Bytes after Truncate = %q, want %q", got, want)
	}
	if got := buf.Cap(); got != capacity {
		t.Fatalf("Cap after Truncate = %d, want %d", got, capacity)
	}
}

func TestTruncateToZero(t *testing.T) {
	t.Parallel()

	var buf Buffer
	buf.AppendString("hello")

	buf.Truncate(0)

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len after Truncate(0) = %d, want 0", got)
	}
}

func TestTruncateToLen(t *testing.T) {
	t.Parallel()

	var buf Buffer
	buf.AppendString("hello")

	capacity := buf.Cap()
	buf.Truncate(buf.Len())

	if got, want := string(buf.Bytes()), "hello"; got != want {
		t.Fatalf("Bytes after Truncate(Len) = %q, want %q", got, want)
	}
	if got := buf.Cap(); got != capacity {
		t.Fatalf("Cap after Truncate(Len) = %d, want %d", got, capacity)
	}
}

func TestTruncateNegativePanics(t *testing.T) {
	t.Parallel()

	var buf Buffer
	assertPanicMessage(t, "buffer: truncate out of range", func() {
		buf.Truncate(-1)
	})
}

func TestTruncateBeyondLenPanics(t *testing.T) {
	t.Parallel()

	var buf Buffer
	buf.AppendString("abc")

	assertPanicMessage(t, "buffer: truncate out of range", func() {
		buf.Truncate(4)
	})
}
