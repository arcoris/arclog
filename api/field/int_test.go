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
)

func TestIntConstructors(t *testing.T) {
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
