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

package predicate_test

import (
	"testing"

	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
	"arcoris.dev/arclog/api/predicate"
)

func BenchmarkAlwaysEvaluation(b *testing.B) {
	benchmarkEvaluation(b, predicate.Always())
}

func BenchmarkAndEvaluation(b *testing.B) {
	benchmarkEvaluation(b, predicate.And(truePredicate(), truePredicate(), truePredicate()))
}

func BenchmarkOrEvaluation(b *testing.B) {
	benchmarkEvaluation(b, predicate.Or(falsePredicate(), falsePredicate(), truePredicate()))
}

func BenchmarkXorEvaluation(b *testing.B) {
	benchmarkEvaluation(b, predicate.Xor(falsePredicate(), truePredicate(), falsePredicate()))
}

func BenchmarkAndConstruction(b *testing.B) {
	p1 := truePredicate()
	p2 := truePredicate()
	p3 := truePredicate()

	b.ReportAllocs()
	for b.Loop() {
		_ = predicate.And(p1, p2, p3)
	}
}

func BenchmarkConstantFoldedConstruction(b *testing.B) {
	always := predicate.Always()
	never := predicate.Never()

	b.ReportAllocs()
	for b.Loop() {
		_ = predicate.And(always, always, never)
	}
}

func BenchmarkDecisiveConstantConstruction(b *testing.B) {
	p := truePredicate()
	never := predicate.Never()

	b.ReportAllocs()
	for b.Loop() {
		_ = predicate.And(p, never)
	}
}

func benchmarkEvaluation(b *testing.B, p predicate.Predicate) {
	entry := predicate.Entry{Level: level.Info, LoggerName: "api"}
	fields := []field.Field{field.String("service", "auth")}

	b.Helper()
	b.ReportAllocs()
	for b.Loop() {
		_ = p.ShouldLog(entry, fields)
	}
}
