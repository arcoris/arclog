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
type ObjectEncoder interface {
	PrimitiveObjectEncoder

	// AddObject appends a nested object field.
	AddObject(dst *buffer.Buffer, key string, marshaler ObjectMarshaler) (*buffer.Buffer, error)

	// AddArray appends a nested array field.
	AddArray(dst *buffer.Buffer, key string, marshaler ArrayMarshaler) (*buffer.Buffer, error)

	// AddReflected appends a field through the implementation's reflection
	// path.
	AddReflected(dst *buffer.Buffer, key string, value any) (*buffer.Buffer, error)

	// OpenNamespace appends an object field whose contents will be populated by
	// subsequent Add* calls.
	OpenNamespace(dst *buffer.Buffer, key string) *buffer.Buffer
}

// PrimitiveObjectEncoder contains primitive field append operations.
type PrimitiveObjectEncoder interface {
	AddBool(dst *buffer.Buffer, key string, value bool) *buffer.Buffer
	AddByteString(dst *buffer.Buffer, key string, value []byte) *buffer.Buffer
	AddComplex128(dst *buffer.Buffer, key string, value complex128) *buffer.Buffer
	AddComplex64(dst *buffer.Buffer, key string, value complex64) *buffer.Buffer
	AddDuration(dst *buffer.Buffer, key string, value time.Duration) *buffer.Buffer
	AddFloat64(dst *buffer.Buffer, key string, value float64) *buffer.Buffer
	AddFloat32(dst *buffer.Buffer, key string, value float32) *buffer.Buffer
	AddInt64(dst *buffer.Buffer, key string, value int64) *buffer.Buffer
	AddInt32(dst *buffer.Buffer, key string, value int32) *buffer.Buffer
	AddInt16(dst *buffer.Buffer, key string, value int16) *buffer.Buffer
	AddInt8(dst *buffer.Buffer, key string, value int8) *buffer.Buffer
	AddInt(dst *buffer.Buffer, key string, value int) *buffer.Buffer
	AddString(dst *buffer.Buffer, key string, value string) *buffer.Buffer
	AddTime(dst *buffer.Buffer, key string, value time.Time) *buffer.Buffer
	AddUint64(dst *buffer.Buffer, key string, value uint64) *buffer.Buffer
	AddUint32(dst *buffer.Buffer, key string, value uint32) *buffer.Buffer
	AddUint16(dst *buffer.Buffer, key string, value uint16) *buffer.Buffer
	AddUint8(dst *buffer.Buffer, key string, value uint8) *buffer.Buffer
	AddUint(dst *buffer.Buffer, key string, value uint) *buffer.Buffer
	AddUintptr(dst *buffer.Buffer, key string, value uintptr) *buffer.Buffer
}
