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

// orPredicate owns an immutable snapshot of operands for hot-path OR
// evaluation.
//
// The slice must only be created from a caller-independent backing array. The
// constructor enforces that invariant; ShouldLog relies on it and performs no
// defensive copying while evaluating a log entry.
type orPredicate []Predicate

// newOrPredicate returns the narrowest OR representation for operands.
//
// operands must already be caller-independent and must not contain removable
// Never constants. Returning a single operand is safe because there is no
// composite membership to protect.
func newOrPredicate(operands []Predicate) Predicate {
	switch len(operands) {
	case 0:
		return Never()
	case 1:
		return operands[0]
	default:
		return orPredicate(operands)
	}
}

// ShouldLog evaluates operands in order and short-circuits on the first true
// result.
func (p orPredicate) ShouldLog(entry Entry, fields []field.Field) bool {
	for _, predicate := range p {
		if predicate.ShouldLog(entry, fields) {
			return true
		}
	}
	return false
}
