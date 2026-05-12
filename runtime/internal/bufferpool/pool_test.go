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

package bufferpool

import (
	"sync"
	"testing"

	"arcoris.dev/arclog/api/buffer"
)

func TestNewReturnsPool(t *testing.T) {
	t.Parallel()

	p := New()
	if p == nil {
		t.Fatal("New() returned nil")
	}
	if p.backend == nil {
		t.Fatal("New() returned Pool with nil backend")
	}
	if p.initialCapacity != DefaultInitialCapacity {
		t.Fatalf("initialCapacity = %d, want %d", p.initialCapacity, DefaultInitialCapacity)
	}
	if p.maxRetainedCapacity != DefaultMaxRetainedCapacity {
		t.Fatalf("maxRetainedCapacity = %d, want %d", p.maxRetainedCapacity, DefaultMaxRetainedCapacity)
	}
}

func TestNewWithOptionsCopiesNormalizedOptions(t *testing.T) {
	t.Parallel()

	options := Options{
		InitialCapacity:     128,
		MaxRetainedCapacity: 1024,
	}
	p := NewWithOptions(options)

	options.InitialCapacity = 1
	options.MaxRetainedCapacity = 1

	if p.initialCapacity != 128 {
		t.Fatalf("initialCapacity = %d, want 128", p.initialCapacity)
	}
	if p.maxRetainedCapacity != 1024 {
		t.Fatalf("maxRetainedCapacity = %d, want 1024", p.maxRetainedCapacity)
	}
}

func TestGetReturnsEmptyBuffer(t *testing.T) {
	t.Parallel()

	p := NewWithOptions(Options{InitialCapacity: 256})
	buf := p.Get()
	if buf == nil {
		t.Fatal("Get() returned nil")
	}
	defer p.Put(buf)

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len() = %d, want 0", got)
	}
	if got := len(buf.Bytes()); got != 0 {
		t.Fatalf("len(Bytes()) = %d, want 0", got)
	}
	if got := buf.Cap(); got < 256 {
		t.Fatalf("Cap() = %d, want >= 256", got)
	}
}

func TestGetReusesNormalCapacityAfterPut(t *testing.T) {
	t.Parallel()

	p := NewWithOptions(Options{
		InitialCapacity:     32,
		MaxRetainedCapacity: 256,
	})

	buf := p.Get()
	buf.AppendString("payload")
	retainedCapacity := buf.Cap()
	p.Put(buf)

	next := p.Get()
	defer p.Put(next)

	if got := next.Len(); got != 0 {
		t.Fatalf("Len() after normal Put/Get = %d, want 0", got)
	}
	if got := string(next.Bytes()); got != "" {
		t.Fatalf("Bytes() after normal Put/Get = %q, want empty", got)
	}
	if got := next.Cap(); got != retainedCapacity {
		t.Fatalf("Cap() after normal Put/Get = %d, want retained capacity %d", got, retainedCapacity)
	}
}

func TestZeroValuePoolIsNonPooling(t *testing.T) {
	t.Parallel()

	var p Pool

	buf := p.Get()
	if buf == nil {
		t.Fatal("zero-value Pool.Get() returned nil")
	}
	if got := buf.Cap(); got < DefaultInitialCapacity {
		t.Fatalf("zero-value Pool.Get() Cap() = %d, want >= %d", got, DefaultInitialCapacity)
	}

	buf.AppendString("discarded")
	p.Put(buf)

	next := p.Get()
	if next == nil {
		t.Fatal("second zero-value Pool.Get() returned nil")
	}
	if next == buf {
		t.Fatal("zero-value Pool unexpectedly retained a buffer")
	}
	if got := next.Len(); got != 0 {
		t.Fatalf("second zero-value Pool.Get() Len() = %d, want 0", got)
	}
}

func TestNilPoolIsNonPooling(t *testing.T) {
	t.Parallel()

	var p *Pool

	buf := p.Get()
	if buf == nil {
		t.Fatal("nil Pool.Get() returned nil")
	}

	p.Put(buf)
	p.Put(nil)
}

func TestPutNilIsNoop(t *testing.T) {
	t.Parallel()

	p := New()
	p.Put(nil)
}

func TestPutEndsOwnershipAndNextGetIsEmpty(t *testing.T) {
	t.Parallel()

	p := New()

	buf := p.Get()
	buf.AppendString("payload")
	p.Put(buf)

	next := p.Get()
	defer p.Put(next)

	if got := next.Len(); got != 0 {
		t.Fatalf("Len() after Put/Get = %d, want 0", got)
	}
	if got := string(next.Bytes()); got != "" {
		t.Fatalf("Bytes() after Put/Get = %q, want empty", got)
	}
}

func TestPutDropsOversizedBuffers(t *testing.T) {
	t.Parallel()

	p := NewWithOptions(Options{
		InitialCapacity:     16,
		MaxRetainedCapacity: 1024,
	})

	buf := p.Get()
	growBeyondCapacity(buf, 1024)
	oversizedCapacity := buf.Cap()
	p.Put(buf)

	next := p.Get()
	defer p.Put(next)

	if got := next.Cap(); got >= oversizedCapacity {
		t.Fatalf("Cap() after oversized Put/Get = %d, want < oversized capacity %d", got, oversizedCapacity)
	}
	if got := next.Len(); got != 0 {
		t.Fatalf("Len() after oversized Put/Get = %d, want 0", got)
	}
}

func TestOversizedBufferDoesNotRetainCapacity(t *testing.T) {
	t.Parallel()

	p := NewWithOptions(Options{
		InitialCapacity:     64,
		MaxRetainedCapacity: 256,
	})

	buf := p.Get()
	growBeyondCapacity(buf, 256)
	oversizedCapacity := buf.Cap()
	p.Put(buf)

	next := p.Get()
	defer p.Put(next)

	if got := next.Cap(); got > 256 {
		t.Fatalf("Cap() after oversized Put/Get = %d, want <= 256", got)
	}
	if got := next.Cap(); got >= oversizedCapacity {
		t.Fatalf("Cap() after oversized Put/Get = %d, want < dropped oversized capacity %d", got, oversizedCapacity)
	}
}

func TestResetBufferForReuseClearsContents(t *testing.T) {
	t.Parallel()

	buf := buffer.New(16)
	buf.AppendString("payload")

	resetBufferForReuse(buf)

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len() after resetBufferForReuse = %d, want 0", got)
	}
	if got := string(buf.Bytes()); got != "" {
		t.Fatalf("Bytes() after resetBufferForReuse = %q, want empty", got)
	}
	if got := buf.Cap(); got < 16 {
		t.Fatalf("Cap() after resetBufferForReuse = %d, want >= 16", got)
	}
}

func TestDropBufferClearsRejectedBufferContents(t *testing.T) {
	t.Parallel()

	buf := buffer.New(16)
	buf.AppendString("payload")

	dropBuffer(buf)

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len() after dropBuffer = %d, want 0", got)
	}
	if got := string(buf.Bytes()); got != "" {
		t.Fatalf("Bytes() after dropBuffer = %q, want empty", got)
	}
}

func TestShouldRetain(t *testing.T) {
	t.Parallel()

	buf := buffer.New(16)
	if !shouldRetain(buf, 16) {
		t.Fatal("shouldRetain(buffer.New(16), 16) = false, want true")
	}
	if shouldRetain(buf, 15) {
		t.Fatal("shouldRetain(buffer.New(16), 15) = true, want false")
	}
	if shouldRetain(nil, 16) {
		t.Fatal("shouldRetain(nil, 16) = true, want false")
	}
}

func TestConcurrentGetPut(t *testing.T) {
	p := New()

	const (
		goroutines = 16
		iterations = 1000
	)

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		i := i
		go func() {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				buf := p.Get()
				if buf == nil {
					t.Errorf("Get() returned nil in goroutine %d", i)
					return
				}

				buf.AppendInt(i)
				buf.AppendByte(':')
				buf.AppendInt(j)
				_ = buf.Bytes()

				p.Put(buf)
			}
		}()
	}

	wg.Wait()
}

func growBeyondCapacity(buf *buffer.Buffer, maxCapacity int) {
	chunk := make([]byte, maxCapacity+1)
	for buf.Cap() <= maxCapacity {
		buf.AppendBytes(chunk)
	}
}
