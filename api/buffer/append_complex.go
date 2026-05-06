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
)

// AppendComplex128 appends v in the form "<real><+|-><imag>i".
//
// Real and imaginary parts are formatted with strconv.AppendFloat using 'g',
// precision -1, and bitSize 64. A plus sign is emitted only when the imaginary
// component is not negative, including negative zero handling through
// math.Signbit.
func (b *Buffer) AppendComplex128(v complex128) {
	appendComplex(b, real(v), imag(v), 64)
}

// AppendComplex64 appends v in the form "<real><+|-><imag>i".
//
// Real and imaginary parts are formatted with 32-bit precision.
func (b *Buffer) AppendComplex64(v complex64) {
	appendComplex(b, float64(real(v)), float64(imag(v)), 32)
}

// appendComplex centralizes complex-number formatting so the complex64 and
// complex128 paths keep identical sign handling.
func appendComplex(b *Buffer, realPart, imagPart float64, bitSize int) {
	b.data = strconv.AppendFloat(b.data, realPart, 'g', -1, bitSize)
	if !math.Signbit(imagPart) {
		b.data = append(b.data, '+')
	}
	b.data = strconv.AppendFloat(b.data, imagPart, 'g', -1, bitSize)
	b.data = append(b.data, 'i')
}
