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

// EnablerFunc adapts a function to the Enabler interface.
//
// It is useful for tests and small adapters. A nil EnablerFunc disables every
// level instead of panicking, which makes optional filters safe to call.
type EnablerFunc func(Level) bool

// Enabled implements Enabler.
//
// If f is nil, Enabled returns false.
func (f EnablerFunc) Enabled(lvl Level) bool {
	if f == nil {
		return false
	}
	return f(lvl)
}
