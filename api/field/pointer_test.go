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
	"testing"
	"time"
)

func TestPointerConstructors(t *testing.T) {
	t.Parallel()

	boolValue := true
	durationValue := 2 * time.Second
	float64Value := 1.25
	float32Value := float32(1.5)
	intValue := 42
	int64Value := int64(43)
	int32Value := int32(44)
	int16Value := int16(45)
	int8Value := int8(46)
	stringValue := "value"
	timeValue := time.Date(2026, 5, 7, 12, 0, 0, 0, time.UTC)
	uintValue := uint(9)
	uint64Value := uint64(10)
	uint32Value := uint32(11)
	uint16Value := uint16(12)
	uint8Value := uint8(13)

	tests := []struct {
		name string
		got  Field
		want Field
	}{
		{name: "bool", got: BoolPtr("v", &boolValue), want: Bool("v", boolValue)},
		{name: "duration", got: DurationPtr("v", &durationValue), want: Duration("v", durationValue)},
		{name: "float64", got: Float64Ptr("v", &float64Value), want: Float64("v", float64Value)},
		{name: "float32", got: Float32Ptr("v", &float32Value), want: Float32("v", float32Value)},
		{name: "int", got: IntPtr("v", &intValue), want: Int("v", intValue)},
		{name: "int64", got: Int64Ptr("v", &int64Value), want: Int64("v", int64Value)},
		{name: "int32", got: Int32Ptr("v", &int32Value), want: Int32("v", int32Value)},
		{name: "int16", got: Int16Ptr("v", &int16Value), want: Int16("v", int16Value)},
		{name: "int8", got: Int8Ptr("v", &int8Value), want: Int8("v", int8Value)},
		{name: "string", got: StringPtr("v", &stringValue), want: String("v", stringValue)},
		{name: "time", got: TimePtr("v", &timeValue), want: Time("v", timeValue)},
		{name: "uint", got: UintPtr("v", &uintValue), want: Uint("v", uintValue)},
		{name: "uint64", got: Uint64Ptr("v", &uint64Value), want: Uint64("v", uint64Value)},
		{name: "uint32", got: Uint32Ptr("v", &uint32Value), want: Uint32("v", uint32Value)},
		{name: "uint16", got: Uint16Ptr("v", &uint16Value), want: Uint16("v", uint16Value)},
		{name: "uint8", got: Uint8Ptr("v", &uint8Value), want: Uint8("v", uint8Value)},
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

func TestNilPointerConstructorsReturnNull(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		got  Field
	}{
		{name: "bool", got: BoolPtr("v", nil)},
		{name: "duration", got: DurationPtr("v", nil)},
		{name: "float64", got: Float64Ptr("v", nil)},
		{name: "float32", got: Float32Ptr("v", nil)},
		{name: "int", got: IntPtr("v", nil)},
		{name: "int64", got: Int64Ptr("v", nil)},
		{name: "int32", got: Int32Ptr("v", nil)},
		{name: "int16", got: Int16Ptr("v", nil)},
		{name: "int8", got: Int8Ptr("v", nil)},
		{name: "string", got: StringPtr("v", nil)},
		{name: "time", got: TimePtr("v", nil)},
		{name: "uint", got: UintPtr("v", nil)},
		{name: "uint64", got: Uint64Ptr("v", nil)},
		{name: "uint32", got: Uint32Ptr("v", nil)},
		{name: "uint16", got: Uint16Ptr("v", nil)},
		{name: "uint8", got: Uint8Ptr("v", nil)},
	}

	want := Null("v")
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if !tt.got.Equal(want) {
				t.Fatalf("got %#v want %#v", tt.got, want)
			}
		})
	}
}
