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

func TestAppendInt(t *testing.T) {
	t.Parallel()

	for _, v := range []int{0, 1, -1, math.MaxInt, math.MinInt} {
		var buf Buffer
		buf.AppendInt(v)

		want := string(strconv.AppendInt(nil, int64(v), 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendInt(%d) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendInt8(t *testing.T) {
	t.Parallel()

	for _, v := range []int8{0, 1, -1, math.MinInt8, math.MaxInt8} {
		var buf Buffer
		buf.AppendInt8(v)

		want := string(strconv.AppendInt(nil, int64(v), 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendInt8(%d) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendInt16(t *testing.T) {
	t.Parallel()

	for _, v := range []int16{0, 1, -1, math.MinInt16, math.MaxInt16} {
		var buf Buffer
		buf.AppendInt16(v)

		want := string(strconv.AppendInt(nil, int64(v), 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendInt16(%d) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendInt32(t *testing.T) {
	t.Parallel()

	for _, v := range []int32{0, 1, -1, math.MinInt32, math.MaxInt32} {
		var buf Buffer
		buf.AppendInt32(v)

		want := string(strconv.AppendInt(nil, int64(v), 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendInt32(%d) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendInt64(t *testing.T) {
	t.Parallel()

	for _, v := range []int64{0, 1, -1, math.MinInt64, math.MaxInt64} {
		var buf Buffer
		buf.AppendInt64(v)

		want := string(strconv.AppendInt(nil, v, 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendInt64(%d) = %q, want %q", v, got, want)
		}
	}
}
