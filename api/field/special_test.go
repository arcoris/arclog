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
	"fmt"
	"testing"
)

type testStringer struct {
	calls int
}

func (s *testStringer) String() string {
	s.calls++
	return "value"
}

type typedNilError struct{}

func (*typedNilError) Error() string { return "typed nil" }

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
}

func TestStringerConstructor(t *testing.T) {
	t.Parallel()

	var nilStringer fmt.Stringer = (*testStringer)(nil)
	if got := Stringer("name", nilStringer); !got.Equal(Null("name")) {
		t.Fatalf("Stringer(nil) = %#v", got)
	}

	value := &testStringer{}
	got := Stringer("name", value)
	if got.Key != "name" || got.Type != StringerType || got.Interface != value {
		t.Fatalf("Stringer(value) = %#v", got)
	}
	if value.calls != 0 {
		t.Fatal("Stringer constructor must not call String")
	}
}

func TestReflectConstructor(t *testing.T) {
	t.Parallel()

	if got := Reflect("key", nil); !got.Equal(Null("key")) {
		t.Fatalf("Reflect(nil) = %#v", got)
	}

	value := struct{ Name string }{Name: "arc"}
	got := Reflect("key", value)
	if got.Key != "key" || got.Type != ReflectType {
		t.Fatalf("Reflect(value) = %#v", got)
	}
	if got.Interface != value {
		t.Fatalf("Reflect(value).Interface = %#v", got.Interface)
	}
}

func TestNamespaceConstructor(t *testing.T) {
	t.Parallel()

	got := Namespace("ns")
	want := Field{Key: "ns", Type: NamespaceType}
	if !got.Equal(want) {
		t.Fatalf("Namespace() = %#v, want %#v", got, want)
	}
}
