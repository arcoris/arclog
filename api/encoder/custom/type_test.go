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

package custom_test

import (
	"testing"
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
	"arcoris.dev/arclog/api/encoder/custom"
)

func TestTypeEncoderFunc(t *testing.T) {
	t.Parallel()

	called := false
	fn := custom.TypeEncoderFunc(func(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, value any) (*buffer.Buffer, error) {
		called = true
		return enc.AddString(dst, key, value.(string)), nil
	})

	dst := buffer.New(0)
	got, err := fn.EncodeType(dst, testObjectEncoder{}, "name", "arcoris")
	if err != nil {
		t.Fatalf("EncodeType() error = %v", err)
	}
	if !called {
		t.Fatal("EncodeType() did not call adapter function")
	}
	if got.String() != "name=arcoris;" {
		t.Fatalf("buffer = %q, want %q", got.String(), "name=arcoris;")
	}
}

// testObjectEncoder records string fields and stubs the remaining object
// methods required by the public ObjectEncoder interface.
type testObjectEncoder struct{}

func (testObjectEncoder) AddString(dst *buffer.Buffer, key string, value string) *buffer.Buffer {
	dst.AppendString(key + "=" + value + ";")
	return dst
}

func (testObjectEncoder) AddBool(dst *buffer.Buffer, key string, value bool) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddByteString(dst *buffer.Buffer, key string, value []byte) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddComplex128(dst *buffer.Buffer, key string, value complex128) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddComplex64(dst *buffer.Buffer, key string, value complex64) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddDuration(dst *buffer.Buffer, key string, value time.Duration) *buffer.Buffer {
	return dst
}

func (testObjectEncoder) AddFloat64(dst *buffer.Buffer, key string, value float64) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddFloat32(dst *buffer.Buffer, key string, value float32) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddInt64(dst *buffer.Buffer, key string, value int64) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddInt32(dst *buffer.Buffer, key string, value int32) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddInt16(dst *buffer.Buffer, key string, value int16) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddInt8(dst *buffer.Buffer, key string, value int8) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddInt(dst *buffer.Buffer, key string, value int) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddTime(dst *buffer.Buffer, key string, value time.Time) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddUint64(dst *buffer.Buffer, key string, value uint64) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddUint32(dst *buffer.Buffer, key string, value uint32) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddUint16(dst *buffer.Buffer, key string, value uint16) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddUint8(dst *buffer.Buffer, key string, value uint8) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddUint(dst *buffer.Buffer, key string, value uint) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddUintptr(dst *buffer.Buffer, key string, value uintptr) *buffer.Buffer {
	return dst
}
func (testObjectEncoder) AddObject(dst *buffer.Buffer, key string, marshaler encoder.ObjectMarshaler) (*buffer.Buffer, error) {
	return dst, nil
}
func (testObjectEncoder) AddArray(dst *buffer.Buffer, key string, marshaler encoder.ArrayMarshaler) (*buffer.Buffer, error) {
	return dst, nil
}
func (testObjectEncoder) AddReflected(dst *buffer.Buffer, key string, value any) (*buffer.Buffer, error) {
	return dst, nil
}
func (testObjectEncoder) OpenNamespace(dst *buffer.Buffer, key string) *buffer.Buffer {
	return dst
}
