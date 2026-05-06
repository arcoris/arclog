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

package encoders_test

import (
	"testing"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder/encoders"
)

func TestAddStringer(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got := encoders.AddStringer(dst, testEncoder{}, "value", stringer("ok"))

	if got.String() != "value=ok;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "value=ok;")
	}
}

func TestAddStringerNil(t *testing.T) {
	t.Parallel()

	var value *nilStringer
	dst := buffer.New(0)
	got := encoders.AddStringer(dst, testEncoder{}, "value", value)

	if got.String() != "value=;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "value=;")
	}
}

func TestAddStringerSafeRecoversPanic(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got := encoders.AddStringerSafe(dst, testEncoder{}, "value", panicStringer{})

	if got.String() != "value=PANIC=boom;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "value=PANIC=boom;")
	}
}

type stringer string

func (s stringer) String() string {
	return string(s)
}

type nilStringer struct{}

func (*nilStringer) String() string {
	return "unreachable"
}

type panicStringer struct{}

func (panicStringer) String() string {
	panic("boom")
}
