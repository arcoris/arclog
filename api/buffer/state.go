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

import "slices"

// Len returns the number of accumulated bytes.
//
// Len is the length of the slice returned by Bytes. It does not report spare
// capacity and does not imply anything about future growth behavior.
func (b *Buffer) Len() int {
	return len(b.data)
}

// Cap returns the capacity of the underlying byte slice.
//
// Cap is exposed for allocation-aware callers and tests. Callers must not rely
// on a precise growth strategy beyond the guarantee that Grow reserves room for
// the requested additional bytes.
func (b *Buffer) Cap() int {
	return cap(b.data)
}

// Bytes returns the accumulated bytes.
//
// The returned slice aliases the buffer's internal storage. The bytes are
// borrowed: callers must treat them as read-only and must not retain them after
// the next mutation, Reset, Truncate, Grow that reallocates, or runtime release
// by an owner outside this package.
func (b *Buffer) Bytes() []byte {
	return b.data
}

// Reset clears the accumulated bytes and retains capacity for reuse.
//
// Reset does not zero the old contents, release memory, return the buffer to a
// pool, or apply any retention limit. Those policies belong to the current
// owner outside this package.
func (b *Buffer) Reset() {
	b.data = b.data[:0]
}

// Grow ensures that the buffer has room for n additional bytes without another
// allocation.
//
// Grow preserves the current contents and length. If enough spare capacity is
// already available, Grow does not allocate. If the backing array grows, any
// previously borrowed slice returned by Bytes must be considered stale.
//
// Grow panics with "buffer: negative grow" if n is negative.
func (b *Buffer) Grow(n int) {
	if n < 0 {
		panic("buffer: negative grow")
	}

	b.data = slices.Grow(b.data, n)
}

// Truncate shortens the buffer to n bytes without changing capacity.
//
// Truncate keeps the first n bytes and discards bytes after n from the logical
// contents. It never changes capacity and it does not zero the discarded tail.
//
// Truncate panics with "buffer: truncate out of range" if n is negative or
// greater than Len.
func (b *Buffer) Truncate(n int) {
	if n < 0 || n > len(b.data) {
		panic("buffer: truncate out of range")
	}

	b.data = b.data[:n]
}
