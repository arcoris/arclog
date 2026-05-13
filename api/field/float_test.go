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

func TestFloatConstructors(t *testing.T) {
	t.Parallel()

	float32Values := []float32{0, -1.5, 3.25, float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())}
	for _, value := range float32Values {
		got := Float32("f32", value)
		if got.Type != Float32Type || uint32(got.Integer) != math.Float32bits(value) {
			t.Fatalf("Float32(%v) = %#v", value, got)
		}
	}

	float64Values := []float64{0, -1.5, 3.25, math.Inf(1), math.Inf(-1), math.NaN()}
	for _, value := range float64Values {
		got := Float64("f64", value)
		if got.Type != Float64Type || uint64(got.Integer) != math.Float64bits(value) {
			t.Fatalf("Float64(%v) = %#v", value, got)
		}
	}
}
