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
	"fmt"
	"testing"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder/convert"
)

func TestAddStringerNil(t *testing.T) {
	t.Parallel()

	var value fmt.Stringer
	dst := buffer.New(0)
	enc := expectObjectString(t, "value", "")

	got := convert.AddStringer(dst, enc, "value", value)

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAddStringerTypedNil(t *testing.T) {
	t.Parallel()

	var value *nilStringer
	dst := buffer.New(0)
	enc := expectObjectString(t, "value", "")

	got := convert.AddStringer(dst, enc, "value", value)

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAddStringerNonNil(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := expectObjectString(t, "value", "ok")

	got := convert.AddStringer(dst, enc, "value", stringer("ok"))

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAppendStringerNil(t *testing.T) {
	t.Parallel()

	var value fmt.Stringer
	dst := buffer.New(0)
	enc := expectArrayString(t, "")

	got := convert.AppendStringer(dst, enc, value)

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAppendStringerTypedNil(t *testing.T) {
	t.Parallel()

	var value *nilStringer
	dst := buffer.New(0)
	enc := expectArrayString(t, "")

	got := convert.AppendStringer(dst, enc, value)

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAppendStringerNonNil(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := expectArrayString(t, "ok")

	got := convert.AppendStringer(dst, enc, stringer("ok"))

	enc.requireCalled()
	requireSameBuffer(t, got, dst)
}

func TestAddStringerPropagatesPanic(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := expectObjectString(t, "value", "")

	requirePanic(t, func() {
		_ = convert.AddStringer(dst, enc, "value", panicStringer{})
	})
}

func TestAppendStringerPropagatesPanic(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := expectArrayString(t, "")

	requirePanic(t, func() {
		_ = convert.AppendStringer(dst, enc, panicStringer{})
	})
}

func TestReturnedBufferIsAuthoritativeForStringer(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	returned := buffer.New(0)
	enc := expectReturnedBuffer(t, expectArrayString(t, "ok"), returned)

	got := convert.AppendStringer(dst, enc, stringer("ok"))

	enc.requireCalled()
	requireSameBuffer(t, got, returned)
}

// stringer is a small value implementation used for successful conversion
// tests.
type stringer string

func (s stringer) String() string {
	return string(s)
}

// nilStringer verifies typed-nil fmt.Stringer handling without invoking
// String.
type nilStringer struct{}

func (*nilStringer) String() string {
	return "unreachable"
}

// panicStringer verifies that strict conversion propagates user panics.
type panicStringer struct{}

func (panicStringer) String() string {
	panic("boom")
}
