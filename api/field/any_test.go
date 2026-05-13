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
	"errors"
	"testing"
	"time"
)

type anyStringer struct{}

func (*anyStringer) String() string { return "value" }

type anyError struct{}

func (*anyError) Error() string { return "typed nil" }

func TestAny(t *testing.T) {
	t.Parallel()

	errBoom := errors.New("boom")
	dur := time.Second
	f64 := 1.25
	f32 := float32(1.5)
	i := 10
	i64 := int64(64)
	i32 := int32(32)
	i16 := int16(16)
	i8 := int8(8)
	s := "value"
	ts := time.Unix(123, 456)
	u := uint(10)
	u64 := uint64(64)
	u32 := uint32(32)
	u16 := uint16(16)
	u8 := uint8(8)
	b := true
	str := &anyStringer{}
	var nilErr error
	var typedNilErr error = (*anyError)(nil)

	tests := []struct {
		name  string
		value any
		want  Field
	}{
		{name: "nil", value: nil, want: Null("k")},
		{name: "bool", value: true, want: Bool("k", true)},
		{name: "bool ptr nil", value: (*bool)(nil), want: Null("k")},
		{name: "bool ptr", value: &b, want: Bool("k", true)},
		{name: "bytes", value: []byte("abc"), want: Bytes("k", []byte("abc"))},
		{name: "duration", value: dur, want: Duration("k", dur)},
		{name: "duration ptr nil", value: (*time.Duration)(nil), want: Null("k")},
		{name: "duration ptr", value: &dur, want: Duration("k", dur)},
		{name: "float64", value: f64, want: Float64("k", f64)},
		{name: "float64 ptr nil", value: (*float64)(nil), want: Null("k")},
		{name: "float64 ptr", value: &f64, want: Float64("k", f64)},
		{name: "float32", value: f32, want: Float32("k", f32)},
		{name: "float32 ptr nil", value: (*float32)(nil), want: Null("k")},
		{name: "float32 ptr", value: &f32, want: Float32("k", f32)},
		{name: "int", value: i, want: Int("k", i)},
		{name: "int ptr nil", value: (*int)(nil), want: Null("k")},
		{name: "int ptr", value: &i, want: Int("k", i)},
		{name: "int64", value: i64, want: Int64("k", i64)},
		{name: "int64 ptr nil", value: (*int64)(nil), want: Null("k")},
		{name: "int64 ptr", value: &i64, want: Int64("k", i64)},
		{name: "int32", value: i32, want: Int32("k", i32)},
		{name: "int32 ptr nil", value: (*int32)(nil), want: Null("k")},
		{name: "int32 ptr", value: &i32, want: Int32("k", i32)},
		{name: "int16", value: i16, want: Int16("k", i16)},
		{name: "int16 ptr nil", value: (*int16)(nil), want: Null("k")},
		{name: "int16 ptr", value: &i16, want: Int16("k", i16)},
		{name: "int8", value: i8, want: Int8("k", i8)},
		{name: "int8 ptr nil", value: (*int8)(nil), want: Null("k")},
		{name: "int8 ptr", value: &i8, want: Int8("k", i8)},
		{name: "string", value: s, want: String("k", s)},
		{name: "string ptr nil", value: (*string)(nil), want: Null("k")},
		{name: "string ptr", value: &s, want: String("k", s)},
		{name: "time", value: ts, want: Time("k", ts)},
		{name: "time ptr nil", value: (*time.Time)(nil), want: Null("k")},
		{name: "time ptr", value: &ts, want: Time("k", ts)},
		{name: "uint", value: u, want: Uint("k", u)},
		{name: "uint ptr nil", value: (*uint)(nil), want: Null("k")},
		{name: "uint ptr", value: &u, want: Uint("k", u)},
		{name: "uint64", value: u64, want: Uint64("k", u64)},
		{name: "uint64 ptr nil", value: (*uint64)(nil), want: Null("k")},
		{name: "uint64 ptr", value: &u64, want: Uint64("k", u64)},
		{name: "uint32", value: u32, want: Uint32("k", u32)},
		{name: "uint32 ptr nil", value: (*uint32)(nil), want: Null("k")},
		{name: "uint32 ptr", value: &u32, want: Uint32("k", u32)},
		{name: "uint16", value: u16, want: Uint16("k", u16)},
		{name: "uint16 ptr nil", value: (*uint16)(nil), want: Null("k")},
		{name: "uint16 ptr", value: &u16, want: Uint16("k", u16)},
		{name: "uint8", value: u8, want: Uint8("k", u8)},
		{name: "uint8 ptr nil", value: (*uint8)(nil), want: Null("k")},
		{name: "uint8 ptr", value: &u8, want: Uint8("k", u8)},
		{name: "nil error interface", value: nilErr, want: Null("k")},
		{name: "error", value: errBoom, want: NamedError("k", errBoom)},
		{name: "typed nil error", value: typedNilErr, want: Skip()},
		{name: "stringer", value: str, want: Stringer("k", str)},
		{name: "fallback", value: struct{ A int }{A: 1}, want: Reflect("k", struct{ A int }{A: 1})},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Any("k", tt.value)
			if !got.Equal(tt.want) {
				t.Fatalf("Any(%T) = %#v, want %#v", tt.value, got, tt.want)
			}
		})
	}
}
