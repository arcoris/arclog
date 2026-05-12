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

func BenchmarkGetPut(b *testing.B) {
	p := New()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		benchmarkBufferSink = buf
		p.Put(buf)
	}
}

func BenchmarkGetPutSmallRecord(b *testing.B) {
	p := New()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		buf.AppendString(`{"level":"info","msg":"ok"}`)
		benchmarkBytesSink = buf.Bytes()
		p.Put(buf)
	}
}

func BenchmarkGetPutWithinRetention(b *testing.B) {
	p := NewWithOptions(Options{
		InitialCapacity:     1024,
		MaxRetainedCapacity: 4 * 1024,
	})
	payload := make([]byte, 3*1024)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		buf.AppendBytes(payload)
		benchmarkBytesSink = buf.Bytes()
		p.Put(buf)
	}
}

func BenchmarkPutOversizedDrop(b *testing.B) {
	p := NewWithOptions(Options{
		InitialCapacity:     256,
		MaxRetainedCapacity: 1024,
	})
	payload := make([]byte, 2048)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := p.Get()
		buf.AppendBytes(payload)
		benchmarkBytesSink = buf.Bytes()
		p.Put(buf)
	}
}

func BenchmarkZeroValuePoolGetPut(b *testing.B) {
	var p Pool

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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := sp.Get().(*buffer.Buffer)
		buf.Reset()
		benchmarkBufferSink = buf
		sp.Put(buf)
	}
}
