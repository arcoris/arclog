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

// Predicate decides whether a log entry should continue through the logging
// pipeline.
//
// ShouldLog returns true when the entry is accepted and false when it is
// suppressed. Implementations must not mutate entry or fields, and must not
// retain fields beyond the call unless they document and enforce a stronger
// ownership contract. Because predicates are usually shared by loggers, cores,
// hooks, or writer routes, implementations must be safe for concurrent calls
// unless the concrete type explicitly documents a narrower contract.
//
// A nil Predicate is invalid. Use Always to represent an absent predicate and
// Never to represent an intentionally disabled route.
type Predicate interface {
	// ShouldLog evaluates entry and fields.
	//
	// The fields slice is borrowed. Implementations may inspect Field values but
	// must not mutate the slice, mutate retained payloads such as []byte values,
	// or retain references to the slice for later use. Expensive work should be
	// avoided because predicates may run before every encode or write attempt.
	ShouldLog(entry Entry, fields []field.Field) bool
}
