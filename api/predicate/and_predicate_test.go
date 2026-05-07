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
	"sync"
	"sync/atomic"
	"testing"

	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
	"arcoris.dev/arclog/api/predicate"
)

func TestAndPredicateShortCircuitsOnFalse(t *testing.T) {
	t.Parallel()

	p := predicate.And(falsePredicate(), panicPredicate(t))
	if p.ShouldLog(predicate.Entry{}, nil) {
		t.Fatal("And(false, panic) = true, want false")
	}
}

func TestAndPredicateSnapshotsOperands(t *testing.T) {
	t.Parallel()

	operands := []predicate.Predicate{falsePredicate(), truePredicate()}
	p := predicate.And(operands...)
	operands[0] = truePredicate()
	operands[1] = truePredicate()

	if p.ShouldLog(predicate.Entry{}, nil) {
		t.Fatal("And result changed after caller-owned slice mutation")
	}
}

func TestAndPredicateSupportsConcurrentEvaluation(t *testing.T) {
	t.Parallel()

	var calls atomic.Int64
	var failures atomic.Int64
	p := predicate.And(
		truePredicate(),
		predicate.Func(func(predicate.Entry, []field.Field) bool {
			calls.Add(1)
			return true
		}),
		truePredicate(),
	)

	const goroutines = 16
	const iterations = 128

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for range goroutines {
		go func() {
			defer wg.Done()
			for range iterations {
				if !p.ShouldLog(predicate.Entry{}, nil) {
					failures.Add(1)
				}
			}
		}()
	}
	wg.Wait()

	if failures.Load() != 0 {
		t.Fatalf("failures = %d, want 0", failures.Load())
	}
	if got, want := calls.Load(), int64(goroutines*iterations); got != want {
		t.Fatalf("calls = %d, want %d", got, want)
	}
}

func TestAndPredicateEvaluationDoesNotAllocate(t *testing.T) {
	p := predicate.And(truePredicate(), truePredicate())
	entry := predicate.Entry{Level: level.Info, Logger: "api"}
	fields := []field.Field{field.String("service", "auth")}

	allocs := testing.AllocsPerRun(1000, func() {
		_ = p.ShouldLog(entry, fields)
	})
	if allocs != 0 {
		t.Fatalf("allocs per evaluation = %g, want 0", allocs)
	}
}
