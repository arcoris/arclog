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

// Not returns a Predicate that accepts an entry exactly when p suppresses it.
//
// If p is Always or Never, Not returns the opposite constant predicate without
// allocating a wrapper. For non-constant predicates, Not stores p as-is; it does
// not copy or synchronize any mutable state owned by p.
//
// A nil predicate is invalid. Not does not panic during construction for nil;
// the returned predicate panics if evaluated.
func Not(p Predicate) Predicate {
	if constant, ok := p.(constantPredicate); ok {
		return constantPredicate(!bool(constant))
	}
	return newNotPredicate(p)
}
