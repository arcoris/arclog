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

// Always returns a Predicate that accepts every entry.
//
// Always is useful as an explicit default when a runtime component wants to
// represent "no predicate configured" without using nil. It ignores Entry and
// fields, performs no allocation during evaluation, and is safe for concurrent
// use.
func Always() Predicate { return constantPredicate(true) }

// Never returns a Predicate that suppresses every entry.
//
// Never is useful for disabled hooks, sinks, or routes that still need a
// concrete Predicate value. It ignores Entry and fields, performs no allocation
// during evaluation, and is safe for concurrent use.
func Never() Predicate { return constantPredicate(false) }

// constantPredicate avoids closures for the two identity predicates and lets
// composition helpers fold Always and Never at construction time.
//
// The type is intentionally unexported. External packages should treat Always
// and Never as ordinary Predicate values rather than depend on their concrete
// representation.
type constantPredicate bool

// ShouldLog returns the constant boolean value represented by p.
//
// It intentionally ignores its arguments and therefore does not inspect or
// retain the borrowed field slice.
func (p constantPredicate) ShouldLog(Entry, []field.Field) bool {
	return bool(p)
}
