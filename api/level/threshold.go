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

// Threshold is an Enabler with an inspectable and mutable minimum level.
//
// Mutation is an optional runtime control surface. Hot paths that only need to
// check whether a record is enabled should depend on Enabler instead of
// Threshold so they do not take a dependency on configuration mutation.
type Threshold interface {
	Enabler

	// Level returns the current threshold level.
	//
	// Implementations should return either a valid severity threshold or Off.
	Level() Level

	// SetLevel updates the current threshold level.
	//
	// Implementations must document whether invalid thresholds panic, return an
	// error elsewhere, or are rejected by a companion method.
	SetLevel(Level)
}
