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
	"runtime/debug"
	"sync"
	"testing"

	"arcoris.dev/arclog/api/buffer"
)

func TestPoolGetProvidesEmptyBuffer(t *testing.T) {
	pool := buffer.NewPool()

	buf := pool.Get()
	if buf == nil {
		t.Fatal("Get returned nil")
	}
	defer pool.Put(buf)

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len after Get = %d, want 0", got)
	}
	if got := len(buf.Bytes()); got != 0 {
		t.Fatalf("len(Bytes) after Get = %d, want 0", got)
	}
	if got := buf.Cap(); got < buffer.Size {
		t.Fatalf("Cap after Get = %d, want >= %d", got, buffer.Size)
	}
}

func TestPoolGetResetsBufferAfterPut(t *testing.T) {
	pool := buffer.NewPool()

	buf := pool.Get()
	buf.AppendString("hello")
	pool.Put(buf)

	buf = pool.Get()
	defer pool.Put(buf)

	if got := buf.Len(); got != 0 {
		t.Fatalf("Len after Put/Get = %d, want 0", got)
	}
	if got := string(buf.Bytes()); got != "" {
		t.Fatalf("Bytes after Put/Get = %q, want empty", got)
	}
}

func TestPoolReusesBuffersWithGCDisabled(t *testing.T) {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	pool := buffer.NewPool()

	buf1 := pool.Get()
	pool.Put(buf1)

	buf2 := pool.Get()
	defer pool.Put(buf2)

	if buf1 != buf2 {
		t.Fatalf("pool did not reuse buffer: got %p, want %p", buf2, buf1)
	}
}

func TestPoolPreservesReasonableCapacity(t *testing.T) {
	pool := buffer.NewPool()

	buf := pool.Get()
	for i := 0; i < 1000; i++ {
		buf.AppendByte('x')
	}
	capacity := buf.Cap()
	if capacity < 1000 {
		t.Fatalf("Cap after growth = %d, want >= 1000", capacity)
	}
	pool.Put(buf)

	buf = pool.Get()
	defer pool.Put(buf)

	if got := buf.Cap(); got < capacity {
		t.Fatalf("Cap after Put/Get = %d, want >= %d", got, capacity)
	}
	if got := buf.Len(); got != 0 {
		t.Fatalf("Len after Put/Get = %d, want 0", got)
	}
}

func TestPoolDiscardsOversizedBuffers(t *testing.T) {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	pool := buffer.NewPool()

	buf := pool.Get()
	for buf.Cap() <= buffer.MaxRetainedSize {
		buf.AppendBytes(make([]byte, 4096))
	}
	oversizedCapacity := buf.Cap()
	pool.Put(buf)

	buf = pool.Get()
	defer pool.Put(buf)

	if got := buf.Cap(); got >= oversizedCapacity {
		t.Fatalf("Cap after oversized Put/Get = %d, want < oversized capacity %d", got, oversizedCapacity)
	}
}

func TestPoolPutNilIsNoop(t *testing.T) {
	pool := buffer.NewPool()
	pool.Put(nil)
}

func TestZeroValuePoolIsSafeButDoesNotPool(t *testing.T) {
	var pool buffer.Pool

	buf := pool.Get()
	if buf == nil {
		t.Fatal("zero-value Pool.Get returned nil")
	}
	buf.AppendString("hello")
	pool.Put(buf)

	buf2 := pool.Get()
	if buf2 == nil {
		t.Fatal("second zero-value Pool.Get returned nil")
	}
	if buf == buf2 {
		t.Fatalf("zero-value Pool unexpectedly reused buffer %p", buf)
	}
}

func TestPoolWithCapacity(t *testing.T) {
	const wantCapacity = 2048

	pool := buffer.NewPoolWithCapacity(wantCapacity)
	buf := pool.Get()
	defer pool.Put(buf)

	if got := buf.Cap(); got < wantCapacity {
		t.Fatalf("Cap = %d, want >= %d", got, wantCapacity)
	}
	if got := buf.Len(); got != 0 {
		t.Fatalf("Len = %d, want 0", got)
	}
}

func TestPoolWithNegativeCapacity(t *testing.T) {
	pool := buffer.NewPoolWithCapacity(-1)
	buf := pool.Get()
	defer pool.Put(buf)

	if got := buf.Cap(); got != 0 {
		t.Fatalf("Cap for negative capacity pool = %d, want 0", got)
	}
}

func TestPoolConcurrentGetPut(t *testing.T) {
	pool := buffer.NewPool()

	const (
		goroutines = 16
		iterations = 1000
	)

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				buf := pool.Get()
				if buf == nil {
					t.Errorf("Get returned nil in goroutine %d", id)
					return
				}
				buf.AppendInt(id)
				buf.AppendByte(':')
				buf.AppendInt(j)
				_ = buf.Bytes()
				pool.Put(buf)
			}
		}(i)
	}

	wg.Wait()

	buf := pool.Get()
	defer pool.Put(buf)
	if got := buf.Len(); got != 0 {
		t.Fatalf("Len after concurrent use = %d, want 0", got)
	}
}
