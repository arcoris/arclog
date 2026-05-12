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
// AppendByte performs no validation or escaping. It is intended for delimiters,
// separators, quotes, and other byte-level fragments that an encoder has
// already decided are correct for its output format.
func (b *Buffer) AppendByte(v byte) {
	b.data = append(b.data, v)
}

// AppendBytes appends v exactly as provided.
//
// The bytes are copied into the buffer by append. Nil and empty slices are
// accepted and leave the contents unchanged. AppendBytes performs no encoding
// validation, escaping, or ownership transfer from v.
func (b *Buffer) AppendBytes(v []byte) {
	b.data = append(b.data, v...)
}

// AppendString appends s as bytes.
//
// AppendString copies the string bytes into the buffer. It does not validate
// UTF-8 and does not perform JSON, console, or protocol-specific escaping.
func (b *Buffer) AppendString(s string) {
	b.data = append(b.data, s...)
}
