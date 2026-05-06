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

package buffer_test

import (
	"math"
	"testing"
)

func TestBufferAppendComplex128(t *testing.T) {
	tests := []struct {
		name string
		val  complex128
		want string
	}{
		{name: "positive imaginary", val: 1 + 2i, want: "1+2i"},
		{name: "negative imaginary", val: 1 - 2i, want: "1-2i"},
		{name: "negative zero imaginary", val: complex(1, math.Copysign(0, -1)), want: "1-0i"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := newBuffer()
			buf.AppendComplex128(tt.val)

			if got := string(buf.Bytes()); got != tt.want {
				t.Fatalf("AppendComplex128(%v) = %q, want %q", tt.val, got, tt.want)
			}
		})
	}
}

func TestBufferAppendComplex64(t *testing.T) {
	tests := []struct {
		name string
		val  complex64
		want string
	}{
		{name: "positive imaginary", val: complex64(1 + 2i), want: "1+2i"},
		{name: "negative imaginary", val: complex64(1 - 2i), want: "1-2i"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := newBuffer()
			buf.AppendComplex64(tt.val)

			if got := string(buf.Bytes()); got != tt.want {
				t.Fatalf("AppendComplex64(%v) = %q, want %q", tt.val, got, tt.want)
			}
		})
	}
}
