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
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
)

type testEncoder struct{}

var _ encoder.ObjectEncoder = testEncoder{}
var _ encoder.ArrayEncoder = testEncoder{}

func (testEncoder) AddString(dst *buffer.Buffer, key string, value string) *buffer.Buffer {
	dst.AppendString(key + "=" + value + ";")
	return dst
}

func (testEncoder) AppendString(dst *buffer.Buffer, value string) *buffer.Buffer {
	dst.AppendString(value + ";")
	return dst
}

func (testEncoder) AddBool(dst *buffer.Buffer, key string, value bool) *buffer.Buffer { return dst }
func (testEncoder) AddByteString(dst *buffer.Buffer, key string, value []byte) *buffer.Buffer {
	return dst
}
func (testEncoder) AddComplex128(dst *buffer.Buffer, key string, value complex128) *buffer.Buffer {
	return dst
}
func (testEncoder) AddComplex64(dst *buffer.Buffer, key string, value complex64) *buffer.Buffer {
	return dst
}
func (testEncoder) AddDuration(dst *buffer.Buffer, key string, value time.Duration) *buffer.Buffer {
	return dst
}
func (testEncoder) AddFloat64(dst *buffer.Buffer, key string, value float64) *buffer.Buffer {
	return dst
}
func (testEncoder) AddFloat32(dst *buffer.Buffer, key string, value float32) *buffer.Buffer {
	return dst
}
func (testEncoder) AddInt64(dst *buffer.Buffer, key string, value int64) *buffer.Buffer {
	return dst
}
func (testEncoder) AddInt32(dst *buffer.Buffer, key string, value int32) *buffer.Buffer {
	return dst
}
func (testEncoder) AddInt16(dst *buffer.Buffer, key string, value int16) *buffer.Buffer {
	return dst
}
func (testEncoder) AddInt8(dst *buffer.Buffer, key string, value int8) *buffer.Buffer {
	return dst
}
func (testEncoder) AddInt(dst *buffer.Buffer, key string, value int) *buffer.Buffer { return dst }
func (testEncoder) AddTime(dst *buffer.Buffer, key string, value time.Time) *buffer.Buffer {
	return dst
}
func (testEncoder) AddUint64(dst *buffer.Buffer, key string, value uint64) *buffer.Buffer {
	return dst
}
func (testEncoder) AddUint32(dst *buffer.Buffer, key string, value uint32) *buffer.Buffer {
	return dst
}
func (testEncoder) AddUint16(dst *buffer.Buffer, key string, value uint16) *buffer.Buffer {
	return dst
}
func (testEncoder) AddUint8(dst *buffer.Buffer, key string, value uint8) *buffer.Buffer {
	return dst
}
func (testEncoder) AddUint(dst *buffer.Buffer, key string, value uint) *buffer.Buffer {
	return dst
}
func (testEncoder) AddUintptr(dst *buffer.Buffer, key string, value uintptr) *buffer.Buffer {
	return dst
}
func (testEncoder) AddObject(dst *buffer.Buffer, key string, marshaler encoder.ObjectMarshaler) (*buffer.Buffer, error) {
	return dst, nil
}
func (testEncoder) AddArray(dst *buffer.Buffer, key string, marshaler encoder.ArrayMarshaler) (*buffer.Buffer, error) {
	return dst, nil
}
func (testEncoder) AddReflected(dst *buffer.Buffer, key string, value any) (*buffer.Buffer, error) {
	return dst, nil
}
func (testEncoder) OpenNamespace(dst *buffer.Buffer, key string) *buffer.Buffer {
	return dst
}

func (testEncoder) AppendBool(dst *buffer.Buffer, value bool) *buffer.Buffer { return dst }
func (testEncoder) AppendByteString(dst *buffer.Buffer, value []byte) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendComplex128(dst *buffer.Buffer, value complex128) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendComplex64(dst *buffer.Buffer, value complex64) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendDuration(dst *buffer.Buffer, value time.Duration) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendFloat64(dst *buffer.Buffer, value float64) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendFloat32(dst *buffer.Buffer, value float32) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendInt64(dst *buffer.Buffer, value int64) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendInt32(dst *buffer.Buffer, value int32) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendInt16(dst *buffer.Buffer, value int16) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendInt8(dst *buffer.Buffer, value int8) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendInt(dst *buffer.Buffer, value int) *buffer.Buffer { return dst }
func (testEncoder) AppendTime(dst *buffer.Buffer, value time.Time) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendUint64(dst *buffer.Buffer, value uint64) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendUint32(dst *buffer.Buffer, value uint32) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendUint16(dst *buffer.Buffer, value uint16) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendUint8(dst *buffer.Buffer, value uint8) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendUint(dst *buffer.Buffer, value uint) *buffer.Buffer { return dst }
func (testEncoder) AppendUintptr(dst *buffer.Buffer, value uintptr) *buffer.Buffer {
	return dst
}
func (testEncoder) AppendObject(dst *buffer.Buffer, marshaler encoder.ObjectMarshaler) (*buffer.Buffer, error) {
	return dst, nil
}
func (testEncoder) AppendArray(dst *buffer.Buffer, marshaler encoder.ArrayMarshaler) (*buffer.Buffer, error) {
	return dst, nil
}
func (testEncoder) AppendReflected(dst *buffer.Buffer, value any) (*buffer.Buffer, error) {
	return dst, nil
}
