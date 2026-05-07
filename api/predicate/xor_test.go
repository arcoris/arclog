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

func TestXorTruthTable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    predicate.Predicate
		want bool
	}{
		{name: "empty", p: predicate.Xor(), want: false},
		{name: "single true", p: predicate.Xor(truePredicate()), want: true},
		{name: "single false", p: predicate.Xor(falsePredicate()), want: false},
		{name: "one true", p: predicate.Xor(falsePredicate(), truePredicate(), falsePredicate()), want: true},
		{name: "zero true", p: predicate.Xor(falsePredicate(), falsePredicate()), want: false},
		{name: "two true", p: predicate.Xor(truePredicate(), truePredicate()), want: false},
		{name: "constant true with all dynamic false", p: predicate.Xor(predicate.Always(), falsePredicate(), falsePredicate()), want: true},
		{name: "constant true with dynamic true", p: predicate.Xor(predicate.Always(), truePredicate()), want: false},
		{name: "two constant true", p: predicate.Xor(predicate.Always(), predicate.Always(), truePredicate()), want: false},
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

func TestXorDecisiveConstantConstructionDoesNotAllocate(t *testing.T) {
	always := predicate.Always()
	dynamic := truePredicate()

	allocs := testing.AllocsPerRun(1000, func() {
		_ = predicate.Xor(always, always, dynamic)
	})
	if allocs != 0 {
		t.Fatalf("allocs per construction = %g, want 0", allocs)
	}
}
