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
	"bytes"
	"testing"
)

var (
	benchSinkBytes []byte
	benchSinkInt   int
)

func BenchmarkAppendByte(b *testing.B) {
	var buf Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendByte('x')
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkAppendBytesSmall(b *testing.B) {
	payload := []byte("hello")
	var buf Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendBytes(payload)
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkAppendBytesLarge(b *testing.B) {
	payload := bytes.Repeat([]byte("a"), 4096)
	buf := New(len(payload))
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendBytes(payload)
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkAppendStringSmall(b *testing.B) {
	payload := "hello"
	var buf Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendString(payload)
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkAppendStringLarge(b *testing.B) {
	payload := string(bytes.Repeat([]byte("a"), 4096))
	buf := New(len(payload))
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendString(payload)
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkAppendBool(b *testing.B) {
	var buf Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendBool(true)
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkAppendInt64(b *testing.B) {
	var buf Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendInt64(9223372036854775807)
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkAppendUint64(b *testing.B) {
	var buf Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendUint64(18446744073709551615)
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkAppendFloat64(b *testing.B) {
	var buf Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendFloat64(3.141592653589793)
	}

	benchSinkBytes = buf.Bytes()
}

func BenchmarkGrowEnoughCapacity(b *testing.B) {
	buf := New(128)
	buf.AppendString("hello")
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendString("hello")
		buf.Grow(16)
	}

	benchSinkInt = buf.Cap()
}

func BenchmarkGrowNeedsAllocation(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf := New(4)
		buf.AppendString("abcd")
		buf.Grow(64)
		benchSinkInt = buf.Cap()
	}
}

func BenchmarkReset(b *testing.B) {
	buf := New(128)
	payload := bytes.Repeat([]byte("a"), 64)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.AppendBytes(payload)
		buf.Reset()
	}

	benchSinkInt = buf.Len()
}

func BenchmarkTruncate(b *testing.B) {
	buf := New(128)
	payload := bytes.Repeat([]byte("a"), 64)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		buf.AppendBytes(payload)
		buf.Truncate(16)
	}

	benchSinkInt = buf.Len()
}

func BenchmarkAppendStringBaselineRawAppend(b *testing.B) {
	payload := "hello"
	var data []byte
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		data = data[:0]
		data = append(data, payload...)
	}

	benchSinkBytes = data
}

func BenchmarkAppendStringBaselineBytesBuffer(b *testing.B) {
	payload := "hello"
	var buf bytes.Buffer
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf.Reset()
		_, _ = buf.WriteString(payload)
	}

	benchSinkBytes = buf.Bytes()
}
