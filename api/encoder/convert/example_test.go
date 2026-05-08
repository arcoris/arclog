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
	"fmt"
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
	"arcoris.dev/arclog/api/encoder/convert"
)

func ExampleAddError() {
	dst := buffer.New(0)
	dst = convert.AddError(dst, exampleObjectEncoder{}, "error", errors.New("failed"))

	fmt.Println(dst.String())

	// Output:
	// error=failed
}

// exampleObjectEncoder is deliberately minimal: convert.AddError only requires
// AddString, but the interface shape keeps examples honest about the API
// boundary.
type exampleObjectEncoder struct{}

var _ encoder.ObjectEncoder = exampleObjectEncoder{}

func (exampleObjectEncoder) AddBool(dst *buffer.Buffer, _ string, _ bool) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddByteString(dst *buffer.Buffer, _ string, _ []byte) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddComplex128(dst *buffer.Buffer, _ string, _ complex128) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddComplex64(dst *buffer.Buffer, _ string, _ complex64) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddDuration(dst *buffer.Buffer, _ string, _ time.Duration) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddFloat64(dst *buffer.Buffer, _ string, _ float64) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddFloat32(dst *buffer.Buffer, _ string, _ float32) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddInt64(dst *buffer.Buffer, _ string, _ int64) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddInt32(dst *buffer.Buffer, _ string, _ int32) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddInt16(dst *buffer.Buffer, _ string, _ int16) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddInt8(dst *buffer.Buffer, _ string, _ int8) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddInt(dst *buffer.Buffer, _ string, _ int) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddString(dst *buffer.Buffer, key string, value string) *buffer.Buffer {
	dst.AppendString(key)
	dst.AppendByte('=')
	dst.AppendString(value)
	return dst
}

func (exampleObjectEncoder) AddTime(dst *buffer.Buffer, _ string, _ time.Time) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddUint64(dst *buffer.Buffer, _ string, _ uint64) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddUint32(dst *buffer.Buffer, _ string, _ uint32) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddUint16(dst *buffer.Buffer, _ string, _ uint16) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddUint8(dst *buffer.Buffer, _ string, _ uint8) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddUint(dst *buffer.Buffer, _ string, _ uint) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddUintptr(dst *buffer.Buffer, _ string, _ uintptr) *buffer.Buffer {
	return dst
}

func (exampleObjectEncoder) AddObject(dst *buffer.Buffer, _ string, _ encoder.ObjectMarshaler) (*buffer.Buffer, error) {
	return dst, nil
}

func (exampleObjectEncoder) AddArray(dst *buffer.Buffer, _ string, _ encoder.ArrayMarshaler) (*buffer.Buffer, error) {
	return dst, nil
}

func (exampleObjectEncoder) AddReflected(dst *buffer.Buffer, _ string, _ any) (*buffer.Buffer, error) {
	return dst, nil
}

func (exampleObjectEncoder) OpenNamespace(dst *buffer.Buffer, _ string) *buffer.Buffer {
	return dst
}
