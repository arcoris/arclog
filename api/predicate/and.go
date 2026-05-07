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

// And returns a Predicate that accepts an entry only when every operand accepts
// it.
//
// Operands are evaluated in the order provided. Evaluation stops at the first
// false result, so operands after a rejecting predicate are not called. With no
// operands, And returns Always. With one non-constant operand, And returns that
// operand directly.
//
// Construction folds constants before creating a composite predicate. Any Never
// operand makes the whole result Never, and Always operands are removed. When a
// composite predicate is still needed, the remaining operands are copied into a
// caller-independent slice so later mutation of the caller's variadic source
// cannot change evaluation.
//
// Nil operands are invalid. And intentionally does not scan for nil values
// beyond the normal constant-folding pass; a nil operand panics if evaluation
// reaches it.
func And(predicates ...Predicate) Predicate {
	switch len(predicates) {
	case 0:
		return Always()
	case 1:
		return normalizeSingle(predicates[0])
	}

	nonConstants := 0
	for _, p := range predicates {
		constant, ok := p.(constantPredicate)
		if !ok {
			nonConstants++
			continue
		}
		if !bool(constant) {
			return Never()
		}
	}

	if nonConstants == 0 {
		return Always()
	}

	operands := make([]Predicate, 0, nonConstants)
	for _, p := range predicates {
		if _, ok := p.(constantPredicate); ok {
			continue
		}
		operands = append(operands, p)
	}

	return newAndPredicate(operands)
}
