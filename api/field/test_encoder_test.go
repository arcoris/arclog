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

package field_test

import (
	"fmt"
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
)

type recordingEncoder struct{}

var _ encoder.ObjectEncoder = recordingEncoder{}
var _ encoder.ArrayEncoder = recordingEncoder{}

func (recordingEncoder) AddBool(dst *buffer.Buffer, key string, value bool) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%t;", key, value))
	return dst
}
func (recordingEncoder) AddByteString(dst *buffer.Buffer, key string, value []byte) *buffer.Buffer {
	dst.AppendString(key + "=" + string(value) + ";")
	return dst
}
func (recordingEncoder) AddComplex128(dst *buffer.Buffer, key string, value complex128) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%v;", key, value))
	return dst
}
func (recordingEncoder) AddComplex64(dst *buffer.Buffer, key string, value complex64) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%v;", key, value))
	return dst
}
func (recordingEncoder) AddDuration(dst *buffer.Buffer, key string, value time.Duration) *buffer.Buffer {
	dst.AppendString(key + "=" + value.String() + ";")
	return dst
}
func (recordingEncoder) AddFloat64(dst *buffer.Buffer, key string, value float64) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%g;", key, value))
	return dst
}
func (recordingEncoder) AddFloat32(dst *buffer.Buffer, key string, value float32) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%g;", key, value))
	return dst
}
func (recordingEncoder) AddInt64(dst *buffer.Buffer, key string, value int64) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddInt32(dst *buffer.Buffer, key string, value int32) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddInt16(dst *buffer.Buffer, key string, value int16) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddInt8(dst *buffer.Buffer, key string, value int8) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddInt(dst *buffer.Buffer, key string, value int) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddString(dst *buffer.Buffer, key string, value string) *buffer.Buffer {
	dst.AppendString(key + "=" + value + ";")
	return dst
}
func (recordingEncoder) AddTime(dst *buffer.Buffer, key string, value time.Time) *buffer.Buffer {
	dst.AppendString(key + "=" + value.UTC().Format(time.RFC3339Nano) + ";")
	return dst
}
func (recordingEncoder) AddUint64(dst *buffer.Buffer, key string, value uint64) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddUint32(dst *buffer.Buffer, key string, value uint32) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddUint16(dst *buffer.Buffer, key string, value uint16) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddUint8(dst *buffer.Buffer, key string, value uint8) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddUint(dst *buffer.Buffer, key string, value uint) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddUintptr(dst *buffer.Buffer, key string, value uintptr) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%s=%d;", key, value))
	return dst
}
func (recordingEncoder) AddObject(dst *buffer.Buffer, key string, marshaler encoder.ObjectMarshaler) (*buffer.Buffer, error) {
	dst.AppendString(key + "={")
	var err error
	dst, err = marshaler.MarshalLogObject(dst, recordingEncoder{})
	if err != nil {
		return dst, err
	}
	dst.AppendString("};")
	return dst, nil
}
func (recordingEncoder) AddArray(dst *buffer.Buffer, key string, marshaler encoder.ArrayMarshaler) (*buffer.Buffer, error) {
	dst.AppendString(key + "=[")
	var err error
	dst, err = marshaler.MarshalLogArray(dst, recordingEncoder{})
	if err != nil {
		return dst, err
	}
	dst.AppendString("];")
	return dst, nil
}
func (recordingEncoder) AddReflected(dst *buffer.Buffer, key string, value any) (*buffer.Buffer, error) {
	dst.AppendString(fmt.Sprintf("%s=%v;", key, value))
	return dst, nil
}
func (recordingEncoder) OpenNamespace(dst *buffer.Buffer, key string) *buffer.Buffer {
	dst.AppendString(key + ".")
	return dst
}
func (recordingEncoder) AppendBool(dst *buffer.Buffer, value bool) *buffer.Buffer         { return dst }
func (recordingEncoder) AppendByteString(dst *buffer.Buffer, value []byte) *buffer.Buffer { return dst }
func (recordingEncoder) AppendComplex128(dst *buffer.Buffer, value complex128) *buffer.Buffer {
	return dst
}
func (recordingEncoder) AppendComplex64(dst *buffer.Buffer, value complex64) *buffer.Buffer {
	return dst
}
func (recordingEncoder) AppendDuration(dst *buffer.Buffer, value time.Duration) *buffer.Buffer {
	return dst
}
func (recordingEncoder) AppendFloat64(dst *buffer.Buffer, value float64) *buffer.Buffer { return dst }
func (recordingEncoder) AppendFloat32(dst *buffer.Buffer, value float32) *buffer.Buffer { return dst }
func (recordingEncoder) AppendInt64(dst *buffer.Buffer, value int64) *buffer.Buffer     { return dst }
func (recordingEncoder) AppendInt32(dst *buffer.Buffer, value int32) *buffer.Buffer     { return dst }
func (recordingEncoder) AppendInt16(dst *buffer.Buffer, value int16) *buffer.Buffer     { return dst }
func (recordingEncoder) AppendInt8(dst *buffer.Buffer, value int8) *buffer.Buffer       { return dst }
func (recordingEncoder) AppendInt(dst *buffer.Buffer, value int) *buffer.Buffer {
	dst.AppendString(fmt.Sprintf("%d;", value))
	return dst
}
func (recordingEncoder) AppendString(dst *buffer.Buffer, value string) *buffer.Buffer {
	dst.AppendString(value + ";")
	return dst
}
func (recordingEncoder) AppendTime(dst *buffer.Buffer, value time.Time) *buffer.Buffer  { return dst }
func (recordingEncoder) AppendUint64(dst *buffer.Buffer, value uint64) *buffer.Buffer   { return dst }
func (recordingEncoder) AppendUint32(dst *buffer.Buffer, value uint32) *buffer.Buffer   { return dst }
func (recordingEncoder) AppendUint16(dst *buffer.Buffer, value uint16) *buffer.Buffer   { return dst }
func (recordingEncoder) AppendUint8(dst *buffer.Buffer, value uint8) *buffer.Buffer     { return dst }
func (recordingEncoder) AppendUint(dst *buffer.Buffer, value uint) *buffer.Buffer       { return dst }
func (recordingEncoder) AppendUintptr(dst *buffer.Buffer, value uintptr) *buffer.Buffer { return dst }
func (recordingEncoder) AppendObject(dst *buffer.Buffer, marshaler encoder.ObjectMarshaler) (*buffer.Buffer, error) {
	return marshaler.MarshalLogObject(dst, recordingEncoder{})
}
func (recordingEncoder) AppendArray(dst *buffer.Buffer, marshaler encoder.ArrayMarshaler) (*buffer.Buffer, error) {
	return marshaler.MarshalLogArray(dst, recordingEncoder{})
}
func (recordingEncoder) AppendReflected(dst *buffer.Buffer, value any) (*buffer.Buffer, error) {
	dst.AppendString(fmt.Sprintf("%v;", value))
	return dst, nil
}
