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

// Buffer is a small mutable byte accumulator for encoder hot paths.
//
// The zero value is ready to use. Buffer owns its byte slice while it is in
// use, but it does not own pooling or lifecycle management outside this
// package.
//
// Methods mutate the receiver in place and return no fluent value. Callers that
// need chaining should provide it at a higher layer, such as an encoder session
// or event builder.
//
// Buffer is not safe for concurrent use.
type Buffer struct {
	// data is the only mutable storage owned by Buffer.
	//
	// It is deliberately unexported so callers can observe contents through
	// Bytes without controlling capacity, growth, truncation, or reuse policy.
	data []byte
}

// New returns an empty buffer with at least the requested capacity.
//
// The returned buffer is not attached to any pool and has length zero. Runtime
// packages that pool buffers own any acquire, release, reuse, drop, and
// retention decisions around this concrete type.
//
// New panics with "buffer: negative capacity" if capacity is negative.
func New(capacity int) *Buffer {
	if capacity < 0 {
		panic("buffer: negative capacity")
	}

	return &Buffer{data: make([]byte, 0, capacity)}
}
