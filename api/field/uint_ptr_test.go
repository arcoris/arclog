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

import "testing"

func TestUintPtrConstructors(t *testing.T) {
	t.Parallel()

	u := uint(1)
	u8 := uint8(8)
	u16 := uint16(16)
	u32 := uint32(32)
	u64 := uint64(64)
	tests := []struct {
		name string
		got  Field
		want Field
	}{
		{name: "uint nil", got: UintPtr("k", nil), want: Null("k")},
		{name: "uint", got: UintPtr("k", &u), want: Uint("k", u)},
		{name: "uint8 nil", got: Uint8Ptr("k", nil), want: Null("k")},
		{name: "uint8", got: Uint8Ptr("k", &u8), want: Uint8("k", u8)},
		{name: "uint16 nil", got: Uint16Ptr("k", nil), want: Null("k")},
		{name: "uint16", got: Uint16Ptr("k", &u16), want: Uint16("k", u16)},
		{name: "uint32 nil", got: Uint32Ptr("k", nil), want: Null("k")},
		{name: "uint32", got: Uint32Ptr("k", &u32), want: Uint32("k", u32)},
		{name: "uint64 nil", got: Uint64Ptr("k", nil), want: Null("k")},
		{name: "uint64", got: Uint64Ptr("k", &u64), want: Uint64("k", u64)},
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
