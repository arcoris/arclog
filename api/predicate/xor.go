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

// Xor returns a Predicate that accepts an entry when exactly one operand accepts
// it.
//
// Operands are evaluated in order. Evaluation stops as soon as a second true
// result is observed, because the final result is then known to be false. With
// no operands, Xor returns Never. With one non-constant operand, Xor returns
// that operand directly.
//
// Construction folds constants before creating a composite predicate. False
// constants are removed. Two or more true constants make the result Never. One
// true constant inverts the OR of the remaining operands, because exactly one
// operand can be true only when every non-constant operand is false.
//
// Nil operands are invalid. Xor intentionally does not scan for nil values
// beyond the normal constant-folding pass; a nil operand panics if evaluation
// reaches it.
func Xor(predicates ...Predicate) Predicate {
	switch len(predicates) {
	case 0:
		return Never()
	case 1:
		return normalizeSingle(predicates[0])
	}

	trueConstants := 0
	nonConstants := 0
	for _, p := range predicates {
		constant, ok := p.(constantPredicate)
		if !ok {
			nonConstants++
			continue
		}
		if !bool(constant) {
			continue
		}
		trueConstants++
		if trueConstants == 2 {
			return Never()
		}
	}

	if nonConstants == 0 {
		if trueConstants == 1 {
			return Always()
		}
		return Never()
	}

	operands := make([]Predicate, 0, nonConstants)
	for _, p := range predicates {
		if _, ok := p.(constantPredicate); ok {
			continue
		}
		operands = append(operands, p)
	}

	if trueConstants == 1 {
		return Not(Or(operands...))
	}

	return newXorPredicate(operands)
}
