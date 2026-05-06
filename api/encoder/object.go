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

package encoder

import (
	"time"

	"arcoris.dev/arclog/api/buffer"
)

// ObjectEncoder appends keyed fields to a buffer.
//
// ObjectEncoder is used by ObjectMarshaler implementations. Primitive methods
// return the authoritative buffer after the append so implementations may swap
// storage if their internal policy requires it.
//
// Implementations own concrete formatting: key escaping, separators, namespace
// syntax, object delimiters, and reflected encoding policy are runtime
// concerns. Callers retain ownership of dst and MUST continue with the returned
// buffer from every method.
type ObjectEncoder interface {
	PrimitiveObjectEncoder

	// AddObject appends a nested object field by executing marshaler.
	//
	// The returned error, if any, is the marshaling or runtime encoding failure
	// that prevented a complete append.
	AddObject(dst *buffer.Buffer, key string, marshaler ObjectMarshaler) (*buffer.Buffer, error)

	// AddArray appends a nested array field by executing marshaler.
	//
	// The returned error, if any, is the marshaling or runtime encoding failure
	// that prevented a complete append.
	AddArray(dst *buffer.Buffer, key string, marshaler ArrayMarshaler) (*buffer.Buffer, error)

	// AddReflected appends a field through the implementation's reflection
	// path.
	//
	// Reflection is an implementation-defined fallback path. Callers should
	// prefer primitive or marshaler methods on hot paths when the value kind is
	// known.
	AddReflected(dst *buffer.Buffer, key string, value any) (*buffer.Buffer, error)

	// OpenNamespace appends an object field whose contents will be populated by
	// subsequent Add* calls.
	//
	// The namespace lifetime and concrete delimiter behavior are defined by the
	// runtime encoder. API callers must continue using the returned buffer.
	OpenNamespace(dst *buffer.Buffer, key string) *buffer.Buffer
}

// PrimitiveObjectEncoder contains error-free primitive field append operations.
//
// These methods append already-typed values and do not execute user-provided
// marshalers. They still return the authoritative buffer after the append.
type PrimitiveObjectEncoder interface {
	// AddBool appends value as a boolean field.
	AddBool(dst *buffer.Buffer, key string, value bool) *buffer.Buffer

	// AddByteString appends value as a byte-string field.
	//
	// Implementations decide whether value is copied, escaped, or encoded
	// immediately. Callers should treat value as borrowed unless the concrete
	// encoder documents stronger ownership guarantees.
	AddByteString(dst *buffer.Buffer, key string, value []byte) *buffer.Buffer

	// AddComplex128 appends value as a complex128 field.
	AddComplex128(dst *buffer.Buffer, key string, value complex128) *buffer.Buffer

	// AddComplex64 appends value as a complex64 field.
	AddComplex64(dst *buffer.Buffer, key string, value complex64) *buffer.Buffer

	// AddDuration appends value as a duration field.
	AddDuration(dst *buffer.Buffer, key string, value time.Duration) *buffer.Buffer

	// AddFloat64 appends value as a float64 field.
	AddFloat64(dst *buffer.Buffer, key string, value float64) *buffer.Buffer

	// AddFloat32 appends value as a float32 field.
	AddFloat32(dst *buffer.Buffer, key string, value float32) *buffer.Buffer

	// AddInt64 appends value as an int64 field.
	AddInt64(dst *buffer.Buffer, key string, value int64) *buffer.Buffer

	// AddInt32 appends value as an int32 field.
	AddInt32(dst *buffer.Buffer, key string, value int32) *buffer.Buffer

	// AddInt16 appends value as an int16 field.
	AddInt16(dst *buffer.Buffer, key string, value int16) *buffer.Buffer

	// AddInt8 appends value as an int8 field.
	AddInt8(dst *buffer.Buffer, key string, value int8) *buffer.Buffer

	// AddInt appends value as an int field.
	AddInt(dst *buffer.Buffer, key string, value int) *buffer.Buffer

	// AddString appends value as a string field.
	AddString(dst *buffer.Buffer, key string, value string) *buffer.Buffer

	// AddTime appends value as a time field.
	AddTime(dst *buffer.Buffer, key string, value time.Time) *buffer.Buffer

	// AddUint64 appends value as a uint64 field.
	AddUint64(dst *buffer.Buffer, key string, value uint64) *buffer.Buffer

	// AddUint32 appends value as a uint32 field.
	AddUint32(dst *buffer.Buffer, key string, value uint32) *buffer.Buffer

	// AddUint16 appends value as a uint16 field.
	AddUint16(dst *buffer.Buffer, key string, value uint16) *buffer.Buffer

	// AddUint8 appends value as a uint8 field.
	AddUint8(dst *buffer.Buffer, key string, value uint8) *buffer.Buffer

	// AddUint appends value as a uint field.
	AddUint(dst *buffer.Buffer, key string, value uint) *buffer.Buffer

	// AddUintptr appends value as a uintptr field.
	AddUintptr(dst *buffer.Buffer, key string, value uintptr) *buffer.Buffer
}
