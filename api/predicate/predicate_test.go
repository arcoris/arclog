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
	"arcoris.dev/arclog/api/predicate"
)

var _ predicate.Predicate = predicate.Func(nil)

func truePredicate() predicate.Predicate {
	return boolPredicate(true)
}

func falsePredicate() predicate.Predicate {
	return boolPredicate(false)
}

func boolPredicate(value bool) predicate.Predicate {
	return predicate.Func(func(predicate.Entry, []field.Field) bool {
		return value
	})
}

func panicPredicate(t *testing.T) predicate.Predicate {
	t.Helper()

	return predicate.Func(func(predicate.Entry, []field.Field) bool {
		t.Fatal("predicate should not have been evaluated")
		return false
	})
}

func mustPanic(t *testing.T, fn func()) {
	t.Helper()

	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()

	fn()
}
