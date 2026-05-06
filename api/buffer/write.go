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

// Write appends p to the buffer and reports a full write.
//
// Write implements io.Writer. It always returns len(p), nil unless the call
// panics because the underlying slice cannot grow.
func (b *Buffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

// WriteByte appends v to the buffer and reports success.
//
// WriteByte implements io.ByteWriter. It always returns nil unless the call
// panics because the underlying slice cannot grow.
func (b *Buffer) WriteByte(v byte) error {
	b.AppendByte(v)
	return nil
}

// WriteString appends s to the buffer and reports a full write.
//
// WriteString implements io.StringWriter. It always returns len(s), nil unless
// the call panics because the underlying slice cannot grow.
func (b *Buffer) WriteString(s string) (int, error) {
	b.AppendString(s)
	return len(s), nil
}
