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

package predicate

import "arcoris.dev/arclog/api/field"

// notPredicate stores the wrapped predicate for hot-path negation.
//
// The wrapped Predicate is intentionally not copied. Predicate implementations
// are interfaces, and their concurrency and lifetime guarantees belong to the
// concrete value supplied to Not.
type notPredicate struct {
	predicate Predicate
}

// newNotPredicate returns the concrete negating wrapper for p.
//
// p may be nil; that invalid state is preserved so evaluation panics at the same
// point as a nil Predicate operand in the other composition forms.
func newNotPredicate(p Predicate) Predicate {
	return notPredicate{predicate: p}
}

// ShouldLog returns the logical negation of the wrapped predicate.
func (p notPredicate) ShouldLog(entry Entry, fields []field.Field) bool {
	return !p.predicate.ShouldLog(entry, fields)
}
