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

// AppendByte appends v as a single byte.
//
// This helper is intended for delimiters, quotes, separators, and other
// one-byte fragments where constructing a temporary slice would be unnecessary.
func (b *Buffer) AppendByte(v byte) {
	b.data = append(b.data, v)
}

// AppendBytes appends v exactly as provided.
//
// No escaping, validation, or interpretation is performed. Callers MUST only
// pass bytes that are already valid for the target encoding.
func (b *Buffer) AppendBytes(v []byte) {
	b.data = append(b.data, v...)
}

// AppendString appends s as bytes.
//
// The bytes of s are copied into the buffer. The method does not perform any
// escaping; callers that need JSON, console, or protocol-specific escaping MUST
// do that before or while appending.
func (b *Buffer) AppendString(s string) {
	b.data = append(b.data, s...)
}
