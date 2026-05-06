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
	"errors"
	"testing"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder/encoders"
)

func TestAddError(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got := encoders.AddError(dst, testEncoder{}, "error", errors.New("failed"))

	if got.String() != "error=failed;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "error=failed;")
	}
}

func TestAddErrorNil(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got := encoders.AddError(dst, testEncoder{}, "error", nil)

	if got.String() != "error=;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "error=;")
	}
}

func TestAddErrorTypedNil(t *testing.T) {
	t.Parallel()

	var err *nilError
	dst := buffer.New(0)
	got := encoders.AddError(dst, testEncoder{}, "error", err)

	if got.String() != "error=;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "error=;")
	}
}

func TestAppendError(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got := encoders.AppendError(dst, testEncoder{}, errors.New("failed"))

	if got.String() != "failed;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "failed;")
	}
}

func TestAddErrorSafeRecoversPanic(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got := encoders.AddErrorSafe(dst, testEncoder{}, "error", panicError{})

	if got.String() != "error=PANIC=boom;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "error=PANIC=boom;")
	}
}

// panicError verifies the safe helper's recovery path.
type panicError struct{}

func (panicError) Error() string {
	panic("boom")
}

// nilError verifies typed-nil error handling without invoking Error.
type nilError struct{}

func (*nilError) Error() string {
	return "unreachable"
}
