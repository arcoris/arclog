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

package buffer

import (
	"math"
	"strconv"
	"testing"
)

func TestAppendFloat32(t *testing.T) {
	t.Parallel()

	values := []float32{
		0,
		float32(math.Copysign(0, -1)),
		1.25,
		-3.5,
		float32(math.Inf(1)),
		float32(math.Inf(-1)),
		float32(math.NaN()),
	}

	for _, v := range values {
		var buf Buffer
		buf.AppendFloat32(v)

		want := string(strconv.AppendFloat(nil, float64(v), 'g', -1, 32))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendFloat32(%v) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendFloat64(t *testing.T) {
	t.Parallel()

	values := []float64{
		0,
		math.Copysign(0, -1),
		1.25,
		-3.5,
		math.Inf(1),
		math.Inf(-1),
		math.NaN(),
	}

	for _, v := range values {
		var buf Buffer
		buf.AppendFloat64(v)

		want := string(strconv.AppendFloat(nil, v, 'g', -1, 64))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendFloat64(%v) = %q, want %q", v, got, want)
		}
	}
}
