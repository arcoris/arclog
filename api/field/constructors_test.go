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
	"errors"
	"math"
	"testing"
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
	"arcoris.dev/arclog/api/field"
)

func TestPrimitiveConstructors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		got, want field.Field
	}{
		{"bool true", field.Bool("ok", true), field.Field{Key: "ok", Type: field.BoolType, Integer: 1}},
		{"string", field.String("name", "arcoris"), field.Field{Key: "name", Type: field.StringType, String: "arcoris"}},
		{"int", field.Int("n", 10), field.Field{Key: "n", Type: field.IntType, Integer: 10}},
		{"duration", field.Duration("d", 5*time.Second), field.Field{Key: "d", Type: field.DurationType, Integer: int64(5 * time.Second)}},
		{"float64", field.Float64("f", 1.25), field.Field{Key: "f", Type: field.Float64Type, Integer: int64(math.Float64bits(1.25))}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if !tt.got.Equal(tt.want) {
				t.Fatalf("got %#v want %#v", tt.got, tt.want)
			}
		})
	}
}

type typedNilArray struct{}

func (*typedNilArray) MarshalLogArray(dst *buffer.Buffer, enc encoder.ArrayEncoder) (*buffer.Buffer, error) {
	return dst, nil
}

type typedNilError struct{}

func (*typedNilError) Error() string { return "typed nil" }

func TestNilAwareConstructors(t *testing.T) {
	t.Parallel()

	var object *typedNilObject
	var array *typedNilArray
	var stringer *testStringer
	var err *typedNilError

	tests := []struct {
		name string
		got  field.Field
		want field.Field
	}{
		{name: "object", got: field.Object("v", object), want: field.Nil("v")},
		{name: "array", got: field.Array("v", array), want: field.Nil("v")},
		{name: "inline", got: field.Inline(object), want: field.Skip()},
		{name: "stringer", got: field.Stringer("v", stringer), want: field.Nil("v")},
		{name: "error", got: field.Error("v", err), want: field.Nil("v")},
		{name: "nil error interface", got: field.Error("v", nil), want: field.Nil("v")},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if !tt.got.Equal(tt.want) {
				t.Fatalf("got %#v want %#v", tt.got, tt.want)
			}
		})
	}
}

func TestErrorConstructorKeepsNonNilError(t *testing.T) {
	t.Parallel()

	err := errors.New("failed")
	got := field.Error("err", err)
	want := field.Field{Key: "err", Type: field.ErrorType, Interface: err}
	if !got.Equal(want) {
		t.Fatalf("got %#v want %#v", got, want)
	}
}

func TestByteStringRetainsSlice(t *testing.T) {
	t.Parallel()
	value := []byte("abc")
	f := field.ByteString("bytes", value)
	value[0] = 'z'
	if got := string(f.Interface.([]byte)); got != "zbc" {
		t.Fatalf("got %q", got)
	}
}
func TestTimeConstructor(t *testing.T) {
	t.Parallel()
	loc := time.FixedZone("T", 3600)
	value := time.Date(2026, 5, 7, 10, 30, 0, 123, loc)
	f := field.Time("time", value)
	if f.Type != field.TimeType {
		t.Fatalf("Type=%s", f.Type)
	}
	if f.Integer != value.UnixNano() {
		t.Fatalf("Integer=%d", f.Integer)
	}
	if f.Interface != loc {
		t.Fatalf("location mismatch")
	}
}
