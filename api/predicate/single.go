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

// normalizeSingle returns the narrowest representation for one operand.
//
// AND, OR, and XOR all have the same truth table when there is exactly one
// operand. Constants are returned as their concrete unexported type so later
// composition can fold them without evaluating through the Predicate interface.
func normalizeSingle(p Predicate) Predicate {
	if constant, ok := p.(constantPredicate); ok {
		return constant
	}
	return p
}
