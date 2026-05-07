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

	"arcoris.dev/arclog/api/predicate"
)

func TestAndTruthTable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    predicate.Predicate
		want bool
	}{
		{name: "empty", p: predicate.And(), want: true},
		{name: "all true", p: predicate.And(predicate.Always(), truePredicate(), truePredicate()), want: true},
		{name: "constant false", p: predicate.And(truePredicate(), predicate.Never(), truePredicate()), want: false},
		{name: "dynamic false", p: predicate.And(truePredicate(), falsePredicate(), truePredicate()), want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.p.ShouldLog(predicate.Entry{}, nil); got != tt.want {
				t.Fatalf("ShouldLog() = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestAndConstantFoldingDoesNotEvaluateLaterOperands(t *testing.T) {
	t.Parallel()

	p := predicate.And(predicate.Never(), panicPredicate(t))
	if p.ShouldLog(predicate.Entry{}, nil) {
		t.Fatal("And(Never, panic) = true, want false")
	}
}

func TestAndDecisiveConstantConstructionDoesNotAllocate(t *testing.T) {
	dynamic := truePredicate()
	never := predicate.Never()

	allocs := testing.AllocsPerRun(1000, func() {
		_ = predicate.And(dynamic, never)
	})
	if allocs != 0 {
		t.Fatalf("allocs per construction = %g, want 0", allocs)
	}
}
