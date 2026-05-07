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

// andPredicate owns an immutable snapshot of operands for hot-path AND
// evaluation.
//
// The slice must only be created from a caller-independent backing array. The
// constructor enforces that invariant; ShouldLog relies on it and performs no
// defensive copying while evaluating a log entry.
type andPredicate []Predicate

// newAndPredicate returns the narrowest AND representation for operands.
//
// operands must already be caller-independent and must not contain removable
// Always constants. Returning a single operand is safe because there is no
// composite membership to protect.
func newAndPredicate(operands []Predicate) Predicate {
	switch len(operands) {
	case 0:
		return Always()
	case 1:
		return operands[0]
	default:
		return andPredicate(operands)
	}
}

// ShouldLog evaluates operands in order and short-circuits on the first false
// result.
func (p andPredicate) ShouldLog(entry Entry, fields []field.Field) bool {
	for _, predicate := range p {
		if !predicate.ShouldLog(entry, fields) {
			return false
		}
	}
	return true
}
