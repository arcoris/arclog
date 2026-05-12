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

// AppendUint appends v using strconv.AppendUint with base 10.
//
// The output contains only the decimal integer representation. The method does
// not add separators or any encoding-specific syntax.
func (b *Buffer) AppendUint(v uint) {
	b.data = strconv.AppendUint(b.data, uint64(v), 10)
}

// AppendUint8 appends v using strconv.AppendUint with base 10.
//
// The value is widened to uint64 before formatting. The decimal representation
// matches strconv.AppendUint for the exact uint8 value.
func (b *Buffer) AppendUint8(v uint8) {
	b.data = strconv.AppendUint(b.data, uint64(v), 10)
}

// AppendUint16 appends v using strconv.AppendUint with base 10.
//
// The value is widened to uint64 before formatting. The decimal representation
// matches strconv.AppendUint for the exact uint16 value.
func (b *Buffer) AppendUint16(v uint16) {
	b.data = strconv.AppendUint(b.data, uint64(v), 10)
}

// AppendUint32 appends v using strconv.AppendUint with base 10.
//
// The value is widened to uint64 before formatting. The decimal representation
// matches strconv.AppendUint for the exact uint32 value.
func (b *Buffer) AppendUint32(v uint32) {
	b.data = strconv.AppendUint(b.data, uint64(v), 10)
}

// AppendUint64 appends v using strconv.AppendUint with base 10.
//
// The decimal representation matches strconv.AppendUint for the exact uint64
// value, including math.MaxUint64.
func (b *Buffer) AppendUint64(v uint64) {
	b.data = strconv.AppendUint(b.data, v, 10)
}
