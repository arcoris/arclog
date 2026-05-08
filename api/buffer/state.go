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

// Bytes returns the current contents of the buffer.
//
// The returned slice aliases the buffer's internal storage. Callers MUST treat
// it as read-only, MUST NOT mutate it, and MUST NOT retain it after the next
// mutation or Reset.
func (b *Buffer) Bytes() []byte {
	return b.data
}

// String returns the current contents of the buffer as a string.
//
// The returned string is a snapshot of the current bytes. Callers should assume
// that the conversion may allocate and should prefer Bytes when they can consume
// a borrowed byte slice safely.
func (b *Buffer) String() string {
	return string(b.data)
}

// Len returns the number of bytes currently stored in the buffer.
func (b *Buffer) Len() int {
	return len(b.data)
}

// Cap returns the capacity of the buffer's internal byte slice.
//
// Cap is exposed for diagnostics and allocation-sensitive tests. Callers MUST
// NOT rely on a specific growth strategy.
func (b *Buffer) Cap() int {
	return cap(b.data)
}

// Reset clears the current contents while retaining reusable capacity.
//
// Reset does not release memory and does not return the buffer to any pool.
// Pooling and retention policy belong to runtime packages.
func (b *Buffer) Reset() {
	b.data = b.data[:0]
}
