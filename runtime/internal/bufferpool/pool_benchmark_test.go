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

var (
	benchmarkBytesSink  []byte
	benchmarkBufferSink *buffer.Buffer
)

func BenchmarkPoolGetPut(b *testing.B) {
	p := New()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		benchmarkBufferSink = buf
		p.Put(buf)
	}
}

func BenchmarkPoolGetAppendSmallPut(b *testing.B) {
	p := New()
	const smallRecord = `{"level":"info","msg":"ok"}`

	// This benchmark intentionally includes the cost of appending a small
	// record that stays within the pool's default initial capacity.
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		buf.AppendString(smallRecord)
		benchmarkBytesSink = buf.Bytes()
		p.Put(buf)
	}
}

func BenchmarkPoolGetAppendWithinRetentionPut(b *testing.B) {
	p := NewWithOptions(Options{
		InitialCapacity:     1024,
		MaxRetainedCapacity: 4 * 1024,
	})
	payload := make([]byte, 3*1024)
	warm := p.Get()
	warm.AppendBytes(payload)
	p.Put(warm)

	// This benchmark intentionally includes payload copy cost, but warms the
	// pool first so the timed loop measures the steady retained path.
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		buf.AppendBytes(payload)
		benchmarkBytesSink = buf.Bytes()
		p.Put(buf)
	}
}

func BenchmarkPoolGetPutRetainedGrownBuffer(b *testing.B) {
	p := NewWithOptions(Options{
		InitialCapacity:     1024,
		MaxRetainedCapacity: 4 * 1024,
	})
	payload := make([]byte, 3*1024)
	warm := p.Get()
	warm.AppendBytes(payload)
	p.Put(warm)

	// This isolates Get/Put overhead when the pool already retains a grown
	// buffer that is still within the configured retention ceiling.
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		benchmarkBufferSink = buf
		p.Put(buf)
	}
}

func BenchmarkPoolOversizedLifecycleDrop(b *testing.B) {
	p := NewWithOptions(Options{
		InitialCapacity:     256,
		MaxRetainedCapacity: 1024,
	})
	payload := make([]byte, 2048)

	// This benchmark includes buffer growth because the pool intentionally drops
	// buffers above MaxRetainedCapacity instead of retaining them for the next
	// iteration. It documents the cost of repeated oversized records, not the
	// normal logging hot path.
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		buf.AppendBytes(payload)
		benchmarkBytesSink = buf.Bytes()
		p.Put(buf)
	}
}

func BenchmarkPoolZeroValueFallback(b *testing.B) {
	var p Pool

	// The zero-value Pool is safe but non-pooling. This benchmark is not the
	// runtime hot path; it documents fallback cost only.
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		benchmarkBufferSink = buf
		p.Put(buf)
	}
}

func BenchmarkSyncPoolBaselineGetPut(b *testing.B) {
	sp := sync.Pool{
		New: func() any {
			return buffer.New(DefaultInitialCapacity)
		},
	}

	// This is a raw lower-bound comparison for sync.Pool. It is not
	// feature-equivalent to Pool because it does not include arclog's retention
	// predicate, oversized-drop policy, zero-value fallback, or ownership docs.
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := sp.Get().(*buffer.Buffer)
		buf.Reset()
		benchmarkBufferSink = buf
		sp.Put(buf)
	}
}
