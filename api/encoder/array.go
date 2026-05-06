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

// ArrayEncoder appends array elements to a buffer.
//
// ArrayEncoder is used by ArrayMarshaler implementations. It mirrors the object
// encoder contract but does not accept field keys because array elements are
// positional.
//
// Unless a concrete implementation documents stronger guarantees, ArrayEncoder
// values are mutable builders and are not safe for concurrent use.
// Implementations own concrete formatting and callers MUST continue with the
// returned buffer from every method.
type ArrayEncoder interface {
	PrimitiveArrayEncoder

	// AppendObject appends a nested object element by executing marshaler.
	AppendObject(dst *buffer.Buffer, marshaler ObjectMarshaler) (*buffer.Buffer, error)

	// AppendArray appends a nested array element by executing marshaler.
	AppendArray(dst *buffer.Buffer, marshaler ArrayMarshaler) (*buffer.Buffer, error)

	// AppendReflected appends an element through the implementation's reflection
	// path.
	//
	// Reflection is an implementation-defined fallback path and may return an
	// error for unsupported values.
	AppendReflected(dst *buffer.Buffer, value any) (*buffer.Buffer, error)
}

// PrimitiveArrayEncoder contains error-free primitive array-element append
// operations.
//
// These methods append already-typed values and do not execute user-provided
// marshalers. They still return the authoritative buffer after the append.
type PrimitiveArrayEncoder interface {
	// AppendBool appends value as a boolean element.
	AppendBool(dst *buffer.Buffer, value bool) *buffer.Buffer

	// AppendByteString appends value as a byte-string element.
	//
	// Implementations decide whether value is copied, escaped, or encoded
	// immediately. Callers should treat value as borrowed unless the concrete
	// encoder documents stronger ownership guarantees.
	AppendByteString(dst *buffer.Buffer, value []byte) *buffer.Buffer

	// AppendComplex128 appends value as a complex128 element.
	AppendComplex128(dst *buffer.Buffer, value complex128) *buffer.Buffer

	// AppendComplex64 appends value as a complex64 element.
	AppendComplex64(dst *buffer.Buffer, value complex64) *buffer.Buffer

	// AppendDuration appends value as a duration element.
	AppendDuration(dst *buffer.Buffer, value time.Duration) *buffer.Buffer

	// AppendFloat64 appends value as a float64 element.
	AppendFloat64(dst *buffer.Buffer, value float64) *buffer.Buffer

	// AppendFloat32 appends value as a float32 element.
	AppendFloat32(dst *buffer.Buffer, value float32) *buffer.Buffer

	// AppendInt64 appends value as an int64 element.
	AppendInt64(dst *buffer.Buffer, value int64) *buffer.Buffer

	// AppendInt32 appends value as an int32 element.
	AppendInt32(dst *buffer.Buffer, value int32) *buffer.Buffer

	// AppendInt16 appends value as an int16 element.
	AppendInt16(dst *buffer.Buffer, value int16) *buffer.Buffer

	// AppendInt8 appends value as an int8 element.
	AppendInt8(dst *buffer.Buffer, value int8) *buffer.Buffer

	// AppendInt appends value as an int element.
	AppendInt(dst *buffer.Buffer, value int) *buffer.Buffer

	// AppendString appends value as a string element.
	AppendString(dst *buffer.Buffer, value string) *buffer.Buffer

	// AppendTime appends value as a time element.
	AppendTime(dst *buffer.Buffer, value time.Time) *buffer.Buffer

	// AppendUint64 appends value as a uint64 element.
	AppendUint64(dst *buffer.Buffer, value uint64) *buffer.Buffer

	// AppendUint32 appends value as a uint32 element.
	AppendUint32(dst *buffer.Buffer, value uint32) *buffer.Buffer

	// AppendUint16 appends value as a uint16 element.
	AppendUint16(dst *buffer.Buffer, value uint16) *buffer.Buffer

	// AppendUint8 appends value as a uint8 element.
	AppendUint8(dst *buffer.Buffer, value uint8) *buffer.Buffer

	// AppendUint appends value as a uint element.
	AppendUint(dst *buffer.Buffer, value uint) *buffer.Buffer

	// AppendUintptr appends value as a uintptr element.
	AppendUintptr(dst *buffer.Buffer, value uintptr) *buffer.Buffer
}
