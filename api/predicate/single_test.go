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

func TestSingleOperandCompositionsBehaveLikeOperand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    predicate.Predicate
		want bool
	}{
		{name: "and true", p: predicate.And(truePredicate()), want: true},
		{name: "and false", p: predicate.And(falsePredicate()), want: false},
		{name: "or true", p: predicate.Or(truePredicate()), want: true},
		{name: "or false", p: predicate.Or(falsePredicate()), want: false},
		{name: "xor true", p: predicate.Xor(truePredicate()), want: true},
		{name: "xor false", p: predicate.Xor(falsePredicate()), want: false},
		{name: "and always", p: predicate.And(predicate.Always()), want: true},
		{name: "or never", p: predicate.Or(predicate.Never()), want: false},
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
