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

func TestOrTruthTable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    predicate.Predicate
		want bool
	}{
		{name: "empty", p: predicate.Or(), want: false},
		{name: "all false", p: predicate.Or(predicate.Never(), falsePredicate(), falsePredicate()), want: false},
		{name: "constant true", p: predicate.Or(falsePredicate(), predicate.Always(), falsePredicate()), want: true},
		{name: "dynamic true", p: predicate.Or(falsePredicate(), truePredicate(), falsePredicate()), want: true},
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

func TestOrConstantFoldingDoesNotEvaluateLaterOperands(t *testing.T) {
	t.Parallel()

	p := predicate.Or(predicate.Always(), panicPredicate(t))
	if !p.ShouldLog(predicate.Entry{}, nil) {
		t.Fatal("Or(Always, panic) = false, want true")
	}
}

func TestOrDecisiveConstantConstructionDoesNotAllocate(t *testing.T) {
	dynamic := falsePredicate()
	always := predicate.Always()

	allocs := testing.AllocsPerRun(1000, func() {
		_ = predicate.Or(dynamic, always)
	})
	if allocs != 0 {
		t.Fatalf("allocs per construction = %g, want 0", allocs)
	}
}
