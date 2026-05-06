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
type ArrayEncoder interface {
	PrimitiveArrayEncoder

	// AppendObject appends a nested object element.
	AppendObject(dst *buffer.Buffer, marshaler ObjectMarshaler) (*buffer.Buffer, error)

	// AppendArray appends a nested array element.
	AppendArray(dst *buffer.Buffer, marshaler ArrayMarshaler) (*buffer.Buffer, error)

	// AppendReflected appends an element through the implementation's reflection
	// path.
	AppendReflected(dst *buffer.Buffer, value any) (*buffer.Buffer, error)
}

// PrimitiveArrayEncoder contains primitive array-element append operations.
type PrimitiveArrayEncoder interface {
	AppendBool(dst *buffer.Buffer, value bool) *buffer.Buffer
	AppendByteString(dst *buffer.Buffer, value []byte) *buffer.Buffer
	AppendComplex128(dst *buffer.Buffer, value complex128) *buffer.Buffer
	AppendComplex64(dst *buffer.Buffer, value complex64) *buffer.Buffer
	AppendDuration(dst *buffer.Buffer, value time.Duration) *buffer.Buffer
	AppendFloat64(dst *buffer.Buffer, value float64) *buffer.Buffer
	AppendFloat32(dst *buffer.Buffer, value float32) *buffer.Buffer
	AppendInt64(dst *buffer.Buffer, value int64) *buffer.Buffer
	AppendInt32(dst *buffer.Buffer, value int32) *buffer.Buffer
	AppendInt16(dst *buffer.Buffer, value int16) *buffer.Buffer
	AppendInt8(dst *buffer.Buffer, value int8) *buffer.Buffer
	AppendInt(dst *buffer.Buffer, value int) *buffer.Buffer
	AppendString(dst *buffer.Buffer, value string) *buffer.Buffer
	AppendTime(dst *buffer.Buffer, value time.Time) *buffer.Buffer
	AppendUint64(dst *buffer.Buffer, value uint64) *buffer.Buffer
	AppendUint32(dst *buffer.Buffer, value uint32) *buffer.Buffer
	AppendUint16(dst *buffer.Buffer, value uint16) *buffer.Buffer
	AppendUint8(dst *buffer.Buffer, value uint8) *buffer.Buffer
	AppendUint(dst *buffer.Buffer, value uint) *buffer.Buffer
	AppendUintptr(dst *buffer.Buffer, value uintptr) *buffer.Buffer
}
