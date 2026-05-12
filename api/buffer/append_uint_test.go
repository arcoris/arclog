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

func TestAppendUint(t *testing.T) {
	t.Parallel()

	for _, v := range []uint{0, 1, math.MaxUint} {
		var buf Buffer
		buf.AppendUint(v)

		want := string(strconv.AppendUint(nil, uint64(v), 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendUint(%d) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendUint8(t *testing.T) {
	t.Parallel()

	for _, v := range []uint8{0, 1, math.MaxUint8} {
		var buf Buffer
		buf.AppendUint8(v)

		want := string(strconv.AppendUint(nil, uint64(v), 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendUint8(%d) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendUint16(t *testing.T) {
	t.Parallel()

	for _, v := range []uint16{0, 1, math.MaxUint16} {
		var buf Buffer
		buf.AppendUint16(v)

		want := string(strconv.AppendUint(nil, uint64(v), 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendUint16(%d) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendUint32(t *testing.T) {
	t.Parallel()

	for _, v := range []uint32{0, 1, math.MaxUint32} {
		var buf Buffer
		buf.AppendUint32(v)

		want := string(strconv.AppendUint(nil, uint64(v), 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendUint32(%d) = %q, want %q", v, got, want)
		}
	}
}

func TestAppendUint64(t *testing.T) {
	t.Parallel()

	for _, v := range []uint64{0, 1, math.MaxUint64} {
		var buf Buffer
		buf.AppendUint64(v)

		want := string(strconv.AppendUint(nil, v, 10))
		if got := string(buf.Bytes()); got != want {
			t.Fatalf("AppendUint64(%d) = %q, want %q", v, got, want)
		}
	}
}
