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

package level

// Threshold is a mutable Enabler governed by a minimum enabled level.
//
// A typical Threshold stores the current minimum level and implements Enabled as
// an inclusive comparison against that level. Concrete implementations may use
// atomics, locks, configuration reloads, or administrative endpoints, but this
// API package defines only the stable contract.
//
// Implementations must be safe for concurrent use by multiple goroutines.
type Threshold interface {
	Enabler

	// Level returns the current minimum enabled level.
	//
	// Implementations should return a valid level. Returning Invalid usually
	// indicates a configuration or implementation error unless documented
	// otherwise by the concrete implementation.
	Level() Level

	// SetLevel updates the current minimum enabled level.
	//
	// Callers should pass only valid levels. Implementations must document their
	// invalid-input policy: panic, reject before calling SetLevel, or another
	// explicit behavior. Silent coercion is discouraged because it hides
	// configuration errors.
	SetLevel(Level)
}
