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

// AppendInt appends v using strconv.AppendInt with base 10.
//
// The output contains only the decimal integer representation. The method does
// not add separators or any encoding-specific syntax.
func (b *Buffer) AppendInt(v int) {
	b.data = strconv.AppendInt(b.data, int64(v), 10)
}

// AppendInt8 appends v using strconv.AppendInt with base 10.
//
// The value is widened to int64 before formatting. The decimal representation
// matches strconv.AppendInt for the exact int8 value.
func (b *Buffer) AppendInt8(v int8) {
	b.data = strconv.AppendInt(b.data, int64(v), 10)
}

// AppendInt16 appends v using strconv.AppendInt with base 10.
//
// The value is widened to int64 before formatting. The decimal representation
// matches strconv.AppendInt for the exact int16 value.
func (b *Buffer) AppendInt16(v int16) {
	b.data = strconv.AppendInt(b.data, int64(v), 10)
}

// AppendInt32 appends v using strconv.AppendInt with base 10.
//
// The value is widened to int64 before formatting. The decimal representation
// matches strconv.AppendInt for the exact int32 value.
func (b *Buffer) AppendInt32(v int32) {
	b.data = strconv.AppendInt(b.data, int64(v), 10)
}

// AppendInt64 appends v using strconv.AppendInt with base 10.
//
// The decimal representation matches strconv.AppendInt for the exact int64
// value, including math.MinInt64 and math.MaxInt64.
func (b *Buffer) AppendInt64(v int64) {
	b.data = strconv.AppendInt(b.data, v, 10)
}
