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

package field

import (
	"errors"
	"testing"
)

type typedNilError struct{}

func (*typedNilError) Error() string { return "typed nil" }

type countingError struct {
	calls int
}

func (e *countingError) Error() string {
	e.calls++
	return "counting"
}

func TestErrorConstructors(t *testing.T) {
	t.Parallel()

	if got := Error(nil); !got.IsSkip() {
		t.Fatalf("Error(nil) = %#v", got)
	}
	if got := NamedError("named", nil); !got.IsSkip() {
		t.Fatalf("NamedError(nil) = %#v", got)
	}

	err := errors.New("boom")
	got := Error(err)
	if got.Key != "error" || got.Type != ErrorType || got.Interface != err {
		t.Fatalf("Error(err) = %#v", got)
	}

	named := NamedError("failure", err)
	if named.Key != "failure" || named.Type != ErrorType || named.Interface != err {
		t.Fatalf("NamedError(err) = %#v", named)
	}

	var typedNil error = (*typedNilError)(nil)
	if got := NamedError("failure", typedNil); !got.IsSkip() {
		t.Fatalf("typed nil error = %#v", got)
	}

	counting := &countingError{}
	_ = Error(counting)
	if counting.calls != 0 {
		t.Fatal("Error constructor must not call Error")
	}
}
