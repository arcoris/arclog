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

// Clock reports timestamps for log entries.
//
// Clock is a timestamp-injection contract, not a timer, scheduler, or runtime
// clock-selection policy. Implementations provide only Now. They do not own
// sleeping, ticking, deadlines, cancellation, background execution, timestamp
// formatting, or monotonic-clock normalization.
//
// Implementations should be safe for concurrent use unless they document a
// narrower contract, because a single Clock may be shared by multiple loggers,
// cores, or tests. A nil Clock interface is invalid; runtime configuration that
// needs an optional clock should handle nil before calling Now.
type Clock interface {
	// Now returns the timestamp selected by the implementation.
	//
	// The returned value may be zero, may carry a monotonic component, and may use
	// any location. The clock package preserves the value exactly as returned by
	// the implementation and assigns no encoding or omission semantics to it.
	Now() time.Time
}
