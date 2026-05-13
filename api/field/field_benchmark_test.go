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

package field

import (
	"errors"
	"testing"
	"time"
)

var benchmarkFieldSink Field

func BenchmarkStringField(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = String("name", "arcoris")
	}
}

func BenchmarkBoolField(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Bool("ok", true)
	}
}

func BenchmarkInt64Field(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Int64("id", 42)
	}
}

func BenchmarkUint64Field(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Uint64("id", 42)
	}
}

func BenchmarkFloat64Field(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Float64("ratio", 3.14)
	}
}

func BenchmarkDurationField(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Duration("dur", time.Second)
	}
}

func BenchmarkTimeField(b *testing.B) {
	ts := time.Date(2026, 5, 13, 12, 0, 0, 0, time.UTC)
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Time("ts", ts)
	}
}

func BenchmarkBytesField(b *testing.B) {
	value := []byte("payload")
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Bytes("data", value)
	}
}

func BenchmarkErrorField(b *testing.B) {
	err := errors.New("boom")
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Error(err)
	}
}

func BenchmarkAnyStringField(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Any("name", "arcoris")
	}
}

func BenchmarkAnyIntField(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Any("id", 42)
	}
}

func BenchmarkAnyBytesField(b *testing.B) {
	value := []byte("payload")
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Any("data", value)
	}
}

func BenchmarkAnyReflectFallbackField(b *testing.B) {
	value := struct {
		Name string
		ID   int
	}{Name: "arc", ID: 1}

	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Any("value", value)
	}
}

func BenchmarkReflectField(b *testing.B) {
	value := struct {
		Name string
		ID   int
	}{Name: "arc", ID: 1}

	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Reflect("value", value)
	}
}

func BenchmarkNamespaceField(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		benchmarkFieldSink = Namespace("ctx")
	}
}
