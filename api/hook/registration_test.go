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

package hook_test

import (
	"testing"

	"arcoris.dev/arclog/api/hook"
)

func TestRegistrationFunc(t *testing.T) {
	t.Parallel()

	called := false
	registration := hook.RegistrationFunc(func() bool {
		called = true
		return true
	})

	var _ hook.Registration = registration

	if !registration.Remove() {
		t.Fatal("Remove() = false, want true")
	}
	if !called {
		t.Fatal("underlying function was not called")
	}
}

func TestNilRegistrationFunc(t *testing.T) {
	t.Parallel()

	var registration hook.RegistrationFunc
	if registration.Remove() {
		t.Fatal("nil RegistrationFunc returned true")
	}
}
