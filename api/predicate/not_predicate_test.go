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

func TestNotPredicateNilPanicsWhenEvaluated(t *testing.T) {
	t.Parallel()

	var nilPredicate predicate.Predicate
	mustPanic(t, func() {
		_ = predicate.Not(nilPredicate).ShouldLog(predicate.Entry{}, nil)
	})
}

func TestNotPredicateEvaluationDoesNotAllocate(t *testing.T) {
	p := predicate.Not(falsePredicate())
	entry := predicate.Entry{Level: level.Info, LoggerName: "api"}
	fields := []field.Field{field.String("service", "auth")}

	allocs := testing.AllocsPerRun(1000, func() {
		_ = p.ShouldLog(entry, fields)
	})
	if allocs != 0 {
		t.Fatalf("allocs per evaluation = %g, want 0", allocs)
	}
}
