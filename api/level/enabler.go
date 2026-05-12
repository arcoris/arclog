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

// Enabler decides whether a record severity is enabled.
//
// Enabler is the minimal hot-path contract for severity filtering. It should be
// cheap to call and should not perform I/O, allocation-heavy work, sampling,
// encoding, or mutation. Runtime packages can provide atomic or static
// implementations on top of this API.
type Enabler interface {
	// Enabled reports whether records at lvl should pass this severity filter.
	Enabled(lvl Level) bool
}
