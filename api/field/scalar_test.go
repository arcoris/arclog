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

package field

import (
	"math"
	"testing"
	"time"
)

func TestBasicConstructors(t *testing.T) {
	t.Parallel()

	t.Run("skip", func(t *testing.T) {
		t.Parallel()
		if !Skip().IsSkip() {
			t.Fatal("Skip() must return a skip field")
		}
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		got := Null("key")
		want := Field{Key: "key", Type: NullType}
		if !got.Equal(want) {
			t.Fatalf("Null() = %#v, want %#v", got, want)
		}
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		if got := Bool("flag", true); !got.Equal(Field{Key: "flag", Type: BoolType, Integer: 1}) {
			t.Fatalf("Bool(true) = %#v", got)
		}
		if got := Bool("flag", false); !got.Equal(Field{Key: "flag", Type: BoolType}) {
			t.Fatalf("Bool(false) = %#v", got)
		}
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		if got := String("name", ""); !got.Equal(Field{Key: "name", Type: StringType}) {
			t.Fatalf("String(empty) = %#v", got)
		}
		if got := String("name", "value"); !got.Equal(Field{Key: "name", Type: StringType, String: "value"}) {
			t.Fatalf("String(value) = %#v", got)
		}
	})

	t.Run("bytes", func(t *testing.T) {
		t.Parallel()
		var nilBytes []byte
		gotNil := Bytes("data", nilBytes)
		if gotNil.Type != BytesType {
			t.Fatalf("Bytes(nil) type = %v", gotNil.Type)
		}
		if gotNil.Interface != nil {
			if _, ok := gotNil.Interface.([]byte); !ok {
				t.Fatalf("Bytes(nil) interface = %#v", gotNil.Interface)
			}
		}

		payload := []byte("abc")
		got := Bytes("data", payload)
		stored, ok := got.Interface.([]byte)
		if !ok {
			t.Fatalf("Bytes() interface type = %T", got.Interface)
		}
		if &stored[0] != &payload[0] {
			t.Fatal("Bytes() must not copy")
		}
	})

	t.Run("duration", func(t *testing.T) {
		t.Parallel()
		tests := []time.Duration{0, -time.Second, 3 * time.Second}
		for _, value := range tests {
			got := Duration("dur", value)
			want := Field{Key: "dur", Type: DurationType, Integer: int64(value)}
			if !got.Equal(want) {
				t.Fatalf("Duration(%v) = %#v, want %#v", value, got, want)
			}
		}
	})

	t.Run("time", func(t *testing.T) {
		t.Parallel()

		zero := time.Time{}
		gotZero := Time("ts", zero)
		if gotZero.Type != TimeFullType {
			t.Fatalf("zero time type = %v, want %v", gotZero.Type, TimeFullType)
		}
		storedZero, ok := gotZero.Interface.(time.Time)
		if !ok || !storedZero.Equal(zero) {
			t.Fatalf("zero time value = %#v", gotZero.Interface)
		}

		loc := time.FixedZone("UTC+3", 3*60*60)
		normal := time.Date(2026, 5, 13, 10, 11, 12, 13, loc)
		gotNormal := Time("ts", normal)
		if gotNormal.Type != TimeType {
			t.Fatalf("normal time type = %v, want %v", gotNormal.Type, TimeType)
		}
		if gotNormal.Integer != normal.UnixNano() {
			t.Fatalf("normal time integer = %d, want %d", gotNormal.Integer, normal.UnixNano())
		}
		if gotNormal.Interface != normal.Location() {
			t.Fatal("normal time must preserve location")
		}

		outOfRange := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)
		gotFull := Time("ts", outOfRange)
		if gotFull.Type != TimeFullType {
			t.Fatalf("out-of-range time type = %v, want %v", gotFull.Type, TimeFullType)
		}
		stored, ok := gotFull.Interface.(time.Time)
		if !ok || !stored.Equal(outOfRange) {
			t.Fatalf("out-of-range time = %#v", gotFull.Interface)
		}
	})
}

func TestSignedIntegerConstructors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		got  Field
		want Field
	}{
		{name: "int zero", got: Int("k", 0), want: Field{Key: "k", Type: IntType}},
		{name: "int positive", got: Int("k", 42), want: Field{Key: "k", Type: IntType, Integer: 42}},
		{name: "int negative", got: Int("k", -42), want: Field{Key: "k", Type: IntType, Integer: -42}},
		{name: "int8 min", got: Int8("k", math.MinInt8), want: Field{Key: "k", Type: Int8Type, Integer: math.MinInt8}},
		{name: "int8 max", got: Int8("k", math.MaxInt8), want: Field{Key: "k", Type: Int8Type, Integer: math.MaxInt8}},
		{name: "int16 min", got: Int16("k", math.MinInt16), want: Field{Key: "k", Type: Int16Type, Integer: math.MinInt16}},
		{name: "int16 max", got: Int16("k", math.MaxInt16), want: Field{Key: "k", Type: Int16Type, Integer: math.MaxInt16}},
		{name: "int32 min", got: Int32("k", math.MinInt32), want: Field{Key: "k", Type: Int32Type, Integer: math.MinInt32}},
		{name: "int32 max", got: Int32("k", math.MaxInt32), want: Field{Key: "k", Type: Int32Type, Integer: math.MaxInt32}},
		{name: "int64 min", got: Int64("k", math.MinInt64), want: Field{Key: "k", Type: Int64Type, Integer: math.MinInt64}},
		{name: "int64 max", got: Int64("k", math.MaxInt64), want: Field{Key: "k", Type: Int64Type, Integer: math.MaxInt64}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if !tt.got.Equal(tt.want) {
				t.Fatalf("got %#v, want %#v", tt.got, tt.want)
			}
		})
	}
}

func TestUnsignedIntegerConstructors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		typ     Type
		integer int64
		field   Field
	}{
		{name: "uint zero", typ: UintType, integer: 0, field: Uint("k", 0)},
		{name: "uint one", typ: UintType, integer: 1, field: Uint("k", 1)},
		{name: "uint8 max", typ: Uint8Type, integer: math.MaxUint8, field: Uint8("k", math.MaxUint8)},
		{name: "uint16 max", typ: Uint16Type, integer: math.MaxUint16, field: Uint16("k", math.MaxUint16)},
		{name: "uint32 max", typ: Uint32Type, integer: math.MaxUint32, field: Uint32("k", math.MaxUint32)},
		{name: "uint64 max", typ: Uint64Type, integer: -1, field: Uint64("k", math.MaxUint64)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := Field{Key: "k", Type: tt.typ, Integer: tt.integer}
			if !tt.field.Equal(want) {
				t.Fatalf("got %#v, want %#v", tt.field, want)
			}
		})
	}
}

func TestFloatConstructors(t *testing.T) {
	t.Parallel()

	float32Values := []float32{0, -1.5, 3.25, float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())}
	for _, value := range float32Values {
		got := Float32("f32", value)
		decoded := math.Float32frombits(uint32(got.Integer))
		if math.IsNaN(float64(value)) {
			if !math.IsNaN(float64(decoded)) {
				t.Fatalf("Float32(%v) decoded to %v", value, decoded)
			}
			continue
		}
		if decoded != value {
			t.Fatalf("Float32(%v) decoded to %v", value, decoded)
		}
	}

	float64Values := []float64{0, -1.5, 3.25, math.Inf(1), math.Inf(-1), math.NaN()}
	for _, value := range float64Values {
		got := Float64("f64", value)
		decoded := math.Float64frombits(uint64(got.Integer))
		if math.IsNaN(value) {
			if !math.IsNaN(decoded) {
				t.Fatalf("Float64(%v) decoded to %v", value, decoded)
			}
			continue
		}
		if decoded != value {
			t.Fatalf("Float64(%v) decoded to %v", value, decoded)
		}
	}
}
