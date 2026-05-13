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

func TestFloatPtrConstructors(t *testing.T) {
	t.Parallel()

	f32 := float32(1.5)
	f64 := float64(1.25)
	tests := []struct {
		name string
		got  Field
		want Field
	}{
		{name: "float32 nil", got: Float32Ptr("k", nil), want: Null("k")},
		{name: "float32", got: Float32Ptr("k", &f32), want: Float32("k", f32)},
		{name: "float64 nil", got: Float64Ptr("k", nil), want: Null("k")},
		{name: "float64", got: Float64Ptr("k", &f64), want: Float64("k", f64)},
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
