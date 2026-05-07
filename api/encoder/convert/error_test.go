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

package convert_test

import (
	"errors"
	"testing"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder/convert"
)

func TestAddErrorNil(t *testing.T) {
	t.Parallel()

	var err error
	dst := buffer.New(0)
	enc := expectObjectString(t, "error", "")

	got := convert.AddError(dst, enc, "error", err)

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAddErrorTypedNil(t *testing.T) {
	t.Parallel()

	var err *nilError
	dst := buffer.New(0)
	enc := expectObjectString(t, "error", "")

	got := convert.AddError(dst, enc, "error", err)

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAddErrorNonNil(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := expectObjectString(t, "error", "failed")

	got := convert.AddError(dst, enc, "error", errors.New("failed"))

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAppendErrorNil(t *testing.T) {
	t.Parallel()

	var err error
	dst := buffer.New(0)
	enc := expectArrayString(t, "")

	got := convert.AppendError(dst, enc, err)

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAppendErrorTypedNil(t *testing.T) {
	t.Parallel()

	var err *nilError
	dst := buffer.New(0)
	enc := expectArrayString(t, "")

	got := convert.AppendError(dst, enc, err)

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAppendErrorNonNil(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := expectArrayString(t, "failed")

	got := convert.AppendError(dst, enc, errors.New("failed"))

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAddErrorPropagatesPanic(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := expectObjectString(t, "error", "")

	requirePanic(t, func() {
		_ = convert.AddError(dst, enc, "error", panicError{})
	})
}

func TestAppendErrorPropagatesPanic(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := expectArrayString(t, "")

	requirePanic(t, func() {
		_ = convert.AppendError(dst, enc, panicError{})
	})
}

func TestReturnedBufferIsAuthoritative(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	returned := buffer.New(0)
	enc := expectReturnedBuffer(t, expectObjectString(t, "error", "failed"), returned)

	got := convert.AddError(dst, enc, "error", errors.New("failed"))

	enc.requireCalled()
	requireSameBuffer(t, got, returned)
}

// nilError verifies typed-nil error handling without invoking Error.
type nilError struct{}

func (*nilError) Error() string {
	return "unreachable"
}

// panicError verifies that strict conversion keeps user panics visible to the
// caller instead of replacing them with API-level diagnostic strings.
type panicError struct{}

func (panicError) Error() string {
	panic("boom")
}
