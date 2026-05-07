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

package clock

import "time"

// Func adapts a function to Clock.
//
// Func is the smallest adapter for tests, integration glue, and runtime
// configuration that already has a timestamp-producing function. The adapter
// does not recover panics, synchronize access, cache values, strip monotonic
// readings, or replace zero timestamps.
//
// A nil Func is invalid and panics when Now is called, matching ordinary nil
// function-call behavior. Use an explicit Clock implementation or a non-nil Func
// that returns time.Time{} when a zero timestamp is desired.
type Func func() time.Time

// Now calls f and returns its timestamp unchanged.
func (f Func) Now() time.Time {
	return f()
}
