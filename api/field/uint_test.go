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

func TestUintConstructors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		got     Field
		typ     Type
		integer int64
	}{
		{name: "uint zero", got: Uint("k", 0), typ: UintType},
		{name: "uint one", got: Uint("k", 1), typ: UintType, integer: 1},
		{name: "uint8 max", got: Uint8("k", math.MaxUint8), typ: Uint8Type, integer: math.MaxUint8},
		{name: "uint16 max", got: Uint16("k", math.MaxUint16), typ: Uint16Type, integer: math.MaxUint16},
		{name: "uint32 max", got: Uint32("k", math.MaxUint32), typ: Uint32Type, integer: math.MaxUint32},
		{name: "uint64 max", got: Uint64("k", math.MaxUint64), typ: Uint64Type, integer: -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := Field{Key: "k", Type: tt.typ, Integer: tt.integer}
			if !tt.got.Equal(want) {
				t.Fatalf("got %#v, want %#v", tt.got, want)
			}
		})
	}
}
