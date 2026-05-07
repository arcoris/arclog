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
	"testing"
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
	"arcoris.dev/arclog/api/field"
)

type typedNilObject struct{}

func (*typedNilObject) MarshalLogObject(dst *buffer.Buffer, enc encoder.ObjectEncoder) (*buffer.Buffer, error) {
	return dst, nil
}

func TestAnyChoosesTypedConstructor(t *testing.T) {
	t.Parallel()

	var nilObject *typedNilObject
	fields := field.Fields{field.String("x", "y")}
	tests := []struct {
		name string
		got  field.Field
		want field.Type
	}{
		{name: "object", got: field.Any("v", encoder.ObjectMarshalerFunc(func(dst *buffer.Buffer, enc encoder.ObjectEncoder) (*buffer.Buffer, error) {
			return dst, nil
		})), want: field.ObjectMarshalerType},
		{name: "typed nil object", got: field.Any("v", nilObject), want: field.ReflectType},
		{name: "fields", got: field.Any("v", fields), want: field.ObjectMarshalerType},
		{name: "field slice", got: field.Any("v", []field.Field{field.Int("x", 1)}), want: field.ObjectMarshalerType},
		{name: "string", got: field.Any("v", "x"), want: field.StringType},
		{name: "int", got: field.Any("v", 1), want: field.IntType},
		{name: "bytes", got: field.Any("v", []byte("x")), want: field.ByteStringType},
		{name: "duration", got: field.Any("v", time.Second), want: field.DurationType},
		{name: "nil", got: field.Any("v", nil), want: field.ReflectType},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.got.Type != tt.want {
				t.Fatalf("Type=%s want %s", tt.got.Type, tt.want)
			}
		})
	}
}
