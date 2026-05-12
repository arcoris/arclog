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

package nilx_test

import (
	"testing"

	"arcoris.dev/arclog/runtime/internal/nilx"
)

func TestIsNilNilInterface(t *testing.T) {
	t.Parallel()

	if !nilx.IsNil(nil) {
		t.Fatal("IsNil(nil) = false, want true")
	}
}

func TestIsNilTypedNilPointer(t *testing.T) {
	t.Parallel()

	var value *int
	if !nilx.IsNil(value) {
		t.Fatal("IsNil(typed nil pointer) = false, want true")
	}
}

func TestIsNilTypedNilSlice(t *testing.T) {
	t.Parallel()

	var value []string
	if !nilx.IsNil(value) {
		t.Fatal("IsNil(typed nil slice) = false, want true")
	}
}

func TestIsNilTypedNilMap(t *testing.T) {
	t.Parallel()

	var value map[string]int
	if !nilx.IsNil(value) {
		t.Fatal("IsNil(typed nil map) = false, want true")
	}
}

func TestIsNilTypedNilChan(t *testing.T) {
	t.Parallel()

	var value chan int
	if !nilx.IsNil(value) {
		t.Fatal("IsNil(typed nil chan) = false, want true")
	}
}

func TestIsNilTypedNilFunc(t *testing.T) {
	t.Parallel()

	var value func()
	if !nilx.IsNil(value) {
		t.Fatal("IsNil(typed nil func) = false, want true")
	}
}

func TestIsNilTypedNilInterface(t *testing.T) {
	t.Parallel()

	var value interface{ Method() }
	if !nilx.IsNil(value) {
		t.Fatal("IsNil(typed nil interface) = false, want true")
	}
}

func TestIsNilNonNilPointer(t *testing.T) {
	t.Parallel()

	value := 1
	if nilx.IsNil(&value) {
		t.Fatal("IsNil(non-nil pointer) = true, want false")
	}
}

func TestIsNilNonNilSlice(t *testing.T) {
	t.Parallel()

	if nilx.IsNil([]string{}) {
		t.Fatal("IsNil(non-nil slice) = true, want false")
	}
}

func TestIsNilScalar(t *testing.T) {
	t.Parallel()

	if nilx.IsNil(0) {
		t.Fatal("IsNil(scalar) = true, want false")
	}
}

func TestIsNilStruct(t *testing.T) {
	t.Parallel()

	if nilx.IsNil(struct{}{}) {
		t.Fatal("IsNil(struct) = true, want false")
	}
}

func TestIsNilArray(t *testing.T) {
	t.Parallel()

	if nilx.IsNil([1]int{}) {
		t.Fatal("IsNil(array) = true, want false")
	}
}

type methodValue struct{}

func (*methodValue) Method() {
	panic("IsNil must not call methods")
}

func TestIsNilDoesNotCallMethods(t *testing.T) {
	t.Parallel()

	var value *methodValue
	if !nilx.IsNil(value) {
		t.Fatal("IsNil(typed nil method value) = false, want true")
	}
}
