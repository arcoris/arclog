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

// Func adapts a function to Predicate.
//
// A Func must obey the same ownership, concurrency, and side-effect contract as
// any other Predicate. A nil Func is invalid and will panic when ShouldLog is
// called, matching ordinary nil function-call behavior.
//
// Func is intended for small inline policies and tests. Long-lived runtime
// implementations may prefer named types when they need explicit state,
// synchronization, or documentation for a non-trivial policy.
type Func func(entry Entry, fields []field.Field) bool

// ShouldLog calls f(entry, fields) and returns its decision unchanged.
//
// The adapter does not recover panics, copy fields, or add synchronization. The
// wrapped function owns those choices as part of the Predicate contract.
func (f Func) ShouldLog(entry Entry, fields []field.Field) bool {
	return f(entry, fields)
}
