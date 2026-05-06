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

// Enabler decides whether a level is enabled.
//
// Enabler is the minimal stable filtering contract used by arclog cores and
// loggers. Implementations usually model an inclusive severity threshold: the
// threshold itself is enabled, and every more severe level is enabled too.
//
// Enabled must be safe for concurrent use and cheap enough to run on every log
// attempt. It should not perform sampling, rate limiting, content inspection,
// I/O, allocation-heavy work, or mutation. Those concerns belong to higher-level
// predicates, cores, or runtime components.
type Enabler interface {
	// Enabled reports whether records at lvl should pass this severity filter.
	//
	// Implementations should return false for Invalid and out-of-range levels
	// unless they explicitly document a different compatibility policy.
	Enabled(lvl Level) bool
}
