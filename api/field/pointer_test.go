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

	"arcoris.dev/arclog/api/field"
)

func TestPointerConstructors(t *testing.T) {
	t.Parallel()

	boolValue := true
	durationValue := 2 * time.Second
	intValue := 42
	stringValue := "value"
	timeValue := time.Date(2026, 5, 7, 12, 0, 0, 0, time.UTC)
	uintValue := uint(9)

	tests := []struct {
		name string
		got  field.Field
		want field.Field
	}{
		{name: "bool", got: field.BoolPtr("v", &boolValue), want: field.Bool("v", boolValue)},
		{name: "duration", got: field.DurationPtr("v", &durationValue), want: field.Duration("v", durationValue)},
		{name: "int", got: field.IntPtr("v", &intValue), want: field.Int("v", intValue)},
		{name: "string", got: field.StringPtr("v", &stringValue), want: field.String("v", stringValue)},
		{name: "time", got: field.TimePtr("v", &timeValue), want: field.Time("v", timeValue)},
		{name: "uint", got: field.UintPtr("v", &uintValue), want: field.Uint("v", uintValue)},
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

func TestNilPointerConstructorReturnsNilField(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		got  field.Field
	}{
		{name: "bool", got: field.BoolPtr("v", nil)},
		{name: "complex128", got: field.Complex128Ptr("v", nil)},
		{name: "duration", got: field.DurationPtr("v", nil)},
		{name: "float64", got: field.Float64Ptr("v", nil)},
		{name: "int", got: field.IntPtr("v", nil)},
		{name: "string", got: field.StringPtr("v", nil)},
		{name: "time", got: field.TimePtr("v", nil)},
		{name: "uint", got: field.UintPtr("v", nil)},
		{name: "uintptr", got: field.UintptrPtr("v", nil)},
	}

	want := field.Nil("v")
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
