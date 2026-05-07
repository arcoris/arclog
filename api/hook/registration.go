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

package hook

// Registration represents a hook registration returned by a Manager.
//
// Registration removes a specific registration, not necessarily every
// registration of the same hook value. This avoids requiring every hook to have
// a globally unique name.
type Registration interface {
	// Remove unregisters the associated hook.
	//
	// Remove returns true when it removed an active registration and false when
	// the registration had already been removed or was never active.
	Remove() bool
}

// RegistrationFunc adapts a function to Registration.
//
// A nil RegistrationFunc is a no-op registration whose Remove method returns
// false.
type RegistrationFunc func() bool

// Remove calls f.
func (f RegistrationFunc) Remove() bool {
	if f == nil {
		return false
	}

	return f()
}
