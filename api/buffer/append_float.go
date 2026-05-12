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

import "strconv"

// AppendFloat32 appends v using strconv.AppendFloat with format 'g',
// precision -1, and bitSize 32.
//
// The value is converted to float64 for the strconv call, while bitSize 32
// keeps the representation compatible with float32 formatting. Special values
// such as NaN and infinities follow strconv.AppendFloat.
func (b *Buffer) AppendFloat32(v float32) {
	b.data = strconv.AppendFloat(b.data, float64(v), 'g', -1, 32)
}

// AppendFloat64 appends v using strconv.AppendFloat with format 'g',
// precision -1, and bitSize 64.
//
// The representation is the compact strconv form for the exact float64 value.
// Special values such as NaN and infinities follow strconv.AppendFloat.
func (b *Buffer) AppendFloat64(v float64) {
	b.data = strconv.AppendFloat(b.data, v, 'g', -1, 64)
}
