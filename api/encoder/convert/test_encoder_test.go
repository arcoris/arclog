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
	"testing"
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
)

const (
	objectStringCall = "AddString"
	arrayStringCall  = "AppendString"
)

// testEncoder implements the object and array contracts narrowly enough to
// prove that conversion functions call the string-oriented encoder methods and
// preserve the returned-buffer contract. Any non-string method is treated as a
// test failure because convert must not route values through structured,
// reflection, or primitive alternatives.
type testEncoder struct {
	t testing.TB

	wantCall  string
	wantKey   string
	wantValue string

	returned *buffer.Buffer
	called   bool
}

var _ encoder.ObjectEncoder = (*testEncoder)(nil)
var _ encoder.ArrayEncoder = (*testEncoder)(nil)

func expectObjectString(t testing.TB, key string, value string) *testEncoder {
	t.Helper()

	return &testEncoder{
		t:         t,
		wantCall:  objectStringCall,
		wantKey:   key,
		wantValue: value,
	}
}

func expectArrayString(t testing.TB, value string) *testEncoder {
	t.Helper()

	return &testEncoder{
		t:         t,
		wantCall:  arrayStringCall,
		wantValue: value,
	}
}

func expectReturnedBuffer(t testing.TB, enc *testEncoder, returned *buffer.Buffer) *testEncoder {
	t.Helper()

	enc.returned = returned
	return enc
}

func (e *testEncoder) requireCalled() {
	e.t.Helper()

	if !e.called {
		e.t.Fatal("encoder string method was not called")
	}
}

func (e *testEncoder) authoritative(dst *buffer.Buffer) *buffer.Buffer {
	if e.returned != nil {
		return e.returned
	}

	return dst
}

func (e *testEncoder) AddString(dst *buffer.Buffer, key string, value string) *buffer.Buffer {
	e.t.Helper()

	e.called = true
	if e.wantCall != objectStringCall {
		e.t.Fatalf("AddString called, want %s", e.wantCall)
	}
	if key != e.wantKey {
		e.t.Fatalf("AddString key = %q, want %q", key, e.wantKey)
	}
	if value != e.wantValue {
		e.t.Fatalf("AddString value = %q, want %q", value, e.wantValue)
	}

	return e.authoritative(dst)
}

func (e *testEncoder) AppendString(dst *buffer.Buffer, value string) *buffer.Buffer {
	e.t.Helper()

	e.called = true
	if e.wantCall != arrayStringCall {
		e.t.Fatalf("AppendString called, want %s", e.wantCall)
	}
	if value != e.wantValue {
		e.t.Fatalf("AppendString value = %q, want %q", value, e.wantValue)
	}

	return e.authoritative(dst)
}

func (e *testEncoder) unexpected(method string) {
	e.t.Helper()
	e.t.Fatalf("unexpected encoder method %s", method)
}

func (e *testEncoder) AddBool(dst *buffer.Buffer, key string, value bool) *buffer.Buffer {
	e.unexpected("AddBool")
	return dst
}

func (e *testEncoder) AddByteString(dst *buffer.Buffer, key string, value []byte) *buffer.Buffer {
	e.unexpected("AddByteString")
	return dst
}

func (e *testEncoder) AddComplex128(dst *buffer.Buffer, key string, value complex128) *buffer.Buffer {
	e.unexpected("AddComplex128")
	return dst
}

func (e *testEncoder) AddComplex64(dst *buffer.Buffer, key string, value complex64) *buffer.Buffer {
	e.unexpected("AddComplex64")
	return dst
}

func (e *testEncoder) AddDuration(dst *buffer.Buffer, key string, value time.Duration) *buffer.Buffer {
	e.unexpected("AddDuration")
	return dst
}

func (e *testEncoder) AddFloat64(dst *buffer.Buffer, key string, value float64) *buffer.Buffer {
	e.unexpected("AddFloat64")
	return dst
}

func (e *testEncoder) AddFloat32(dst *buffer.Buffer, key string, value float32) *buffer.Buffer {
	e.unexpected("AddFloat32")
	return dst
}

func (e *testEncoder) AddInt64(dst *buffer.Buffer, key string, value int64) *buffer.Buffer {
	e.unexpected("AddInt64")
	return dst
}

func (e *testEncoder) AddInt32(dst *buffer.Buffer, key string, value int32) *buffer.Buffer {
	e.unexpected("AddInt32")
	return dst
}

func (e *testEncoder) AddInt16(dst *buffer.Buffer, key string, value int16) *buffer.Buffer {
	e.unexpected("AddInt16")
	return dst
}

func (e *testEncoder) AddInt8(dst *buffer.Buffer, key string, value int8) *buffer.Buffer {
	e.unexpected("AddInt8")
	return dst
}

func (e *testEncoder) AddInt(dst *buffer.Buffer, key string, value int) *buffer.Buffer {
	e.unexpected("AddInt")
	return dst
}

func (e *testEncoder) AddTime(dst *buffer.Buffer, key string, value time.Time) *buffer.Buffer {
	e.unexpected("AddTime")
	return dst
}

func (e *testEncoder) AddUint64(dst *buffer.Buffer, key string, value uint64) *buffer.Buffer {
	e.unexpected("AddUint64")
	return dst
}

func (e *testEncoder) AddUint32(dst *buffer.Buffer, key string, value uint32) *buffer.Buffer {
	e.unexpected("AddUint32")
	return dst
}

func (e *testEncoder) AddUint16(dst *buffer.Buffer, key string, value uint16) *buffer.Buffer {
	e.unexpected("AddUint16")
	return dst
}

func (e *testEncoder) AddUint8(dst *buffer.Buffer, key string, value uint8) *buffer.Buffer {
	e.unexpected("AddUint8")
	return dst
}

func (e *testEncoder) AddUint(dst *buffer.Buffer, key string, value uint) *buffer.Buffer {
	e.unexpected("AddUint")
	return dst
}

func (e *testEncoder) AddUintptr(dst *buffer.Buffer, key string, value uintptr) *buffer.Buffer {
	e.unexpected("AddUintptr")
	return dst
}

func (e *testEncoder) AddObject(dst *buffer.Buffer, key string, marshaler encoder.ObjectMarshaler) (*buffer.Buffer, error) {
	e.unexpected("AddObject")
	return dst, nil
}

func (e *testEncoder) AddArray(dst *buffer.Buffer, key string, marshaler encoder.ArrayMarshaler) (*buffer.Buffer, error) {
	e.unexpected("AddArray")
	return dst, nil
}

func (e *testEncoder) AddReflected(dst *buffer.Buffer, key string, value any) (*buffer.Buffer, error) {
	e.unexpected("AddReflected")
	return dst, nil
}

func (e *testEncoder) OpenNamespace(dst *buffer.Buffer, key string) *buffer.Buffer {
	e.unexpected("OpenNamespace")
	return dst
}

func (e *testEncoder) AppendBool(dst *buffer.Buffer, value bool) *buffer.Buffer {
	e.unexpected("AppendBool")
	return dst
}

func (e *testEncoder) AppendByteString(dst *buffer.Buffer, value []byte) *buffer.Buffer {
	e.unexpected("AppendByteString")
	return dst
}

func (e *testEncoder) AppendComplex128(dst *buffer.Buffer, value complex128) *buffer.Buffer {
	e.unexpected("AppendComplex128")
	return dst
}

func (e *testEncoder) AppendComplex64(dst *buffer.Buffer, value complex64) *buffer.Buffer {
	e.unexpected("AppendComplex64")
	return dst
}

func (e *testEncoder) AppendDuration(dst *buffer.Buffer, value time.Duration) *buffer.Buffer {
	e.unexpected("AppendDuration")
	return dst
}

func (e *testEncoder) AppendFloat64(dst *buffer.Buffer, value float64) *buffer.Buffer {
	e.unexpected("AppendFloat64")
	return dst
}

func (e *testEncoder) AppendFloat32(dst *buffer.Buffer, value float32) *buffer.Buffer {
	e.unexpected("AppendFloat32")
	return dst
}

func (e *testEncoder) AppendInt64(dst *buffer.Buffer, value int64) *buffer.Buffer {
	e.unexpected("AppendInt64")
	return dst
}

func (e *testEncoder) AppendInt32(dst *buffer.Buffer, value int32) *buffer.Buffer {
	e.unexpected("AppendInt32")
	return dst
}

func (e *testEncoder) AppendInt16(dst *buffer.Buffer, value int16) *buffer.Buffer {
	e.unexpected("AppendInt16")
	return dst
}

func (e *testEncoder) AppendInt8(dst *buffer.Buffer, value int8) *buffer.Buffer {
	e.unexpected("AppendInt8")
	return dst
}

func (e *testEncoder) AppendInt(dst *buffer.Buffer, value int) *buffer.Buffer {
	e.unexpected("AppendInt")
	return dst
}

func (e *testEncoder) AppendTime(dst *buffer.Buffer, value time.Time) *buffer.Buffer {
	e.unexpected("AppendTime")
	return dst
}

func (e *testEncoder) AppendUint64(dst *buffer.Buffer, value uint64) *buffer.Buffer {
	e.unexpected("AppendUint64")
	return dst
}

func (e *testEncoder) AppendUint32(dst *buffer.Buffer, value uint32) *buffer.Buffer {
	e.unexpected("AppendUint32")
	return dst
}

func (e *testEncoder) AppendUint16(dst *buffer.Buffer, value uint16) *buffer.Buffer {
	e.unexpected("AppendUint16")
	return dst
}

func (e *testEncoder) AppendUint8(dst *buffer.Buffer, value uint8) *buffer.Buffer {
	e.unexpected("AppendUint8")
	return dst
}

func (e *testEncoder) AppendUint(dst *buffer.Buffer, value uint) *buffer.Buffer {
	e.unexpected("AppendUint")
	return dst
}

func (e *testEncoder) AppendUintptr(dst *buffer.Buffer, value uintptr) *buffer.Buffer {
	e.unexpected("AppendUintptr")
	return dst
}

func (e *testEncoder) AppendObject(dst *buffer.Buffer, marshaler encoder.ObjectMarshaler) (*buffer.Buffer, error) {
	e.unexpected("AppendObject")
	return dst, nil
}

func (e *testEncoder) AppendArray(dst *buffer.Buffer, marshaler encoder.ArrayMarshaler) (*buffer.Buffer, error) {
	e.unexpected("AppendArray")
	return dst, nil
}

func (e *testEncoder) AppendReflected(dst *buffer.Buffer, value any) (*buffer.Buffer, error) {
	e.unexpected("AppendReflected")
	return dst, nil
}

func requireSameBuffer(t testing.TB, got *buffer.Buffer, want *buffer.Buffer) {
	t.Helper()

	if got != want {
		t.Fatalf("returned buffer = %p, want %p", got, want)
	}
}

func requirePanic(t testing.TB, fn func()) {
	t.Helper()

	defer func() {
		if recover() == nil {
			t.Fatal("function did not panic")
		}
	}()

	fn()
}
