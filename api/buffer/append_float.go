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

// AppendFloat64 appends v using compact 64-bit floating-point formatting.
//
// The method uses strconv.AppendFloat with format 'g', precision -1, and
// bitSize 64. This preserves round-trip precision while keeping common values
// compact.
func (b *Buffer) AppendFloat64(v float64) {
	b.data = strconv.AppendFloat(b.data, v, 'g', -1, 64)
}

// AppendFloat32 appends v using compact 32-bit floating-point formatting.
//
// The value is widened to float64 for strconv.AppendFloat, but bitSize remains
// 32 so the output reflects float32 precision.
func (b *Buffer) AppendFloat32(v float32) {
	b.data = strconv.AppendFloat(b.data, float64(v), 'g', -1, 32)
}
