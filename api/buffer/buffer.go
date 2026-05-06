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

// Buffer is a reusable, growable byte buffer for encoded log records and other
// serialized payloads produced by the logging pipeline.
//
// The zero value is ready to use. A zero-value Buffer is not attached to any
// Pool; calling Free on such a buffer is a no-op. Buffers obtained from a Pool
// remember their originating pool and Free returns them to that pool.
//
// Buffer is not safe for concurrent use. A caller that shares a Buffer between
// goroutines MUST synchronize all method calls and all access to slices returned
// by Bytes.
type Buffer struct {
	// data stores the current logical contents of the buffer. len(data) is the
	// number of valid bytes and cap(data) is reusable storage.
	data []byte

	// pool is the originating Pool. It is set by Pool.Get and NewWithPool so
	// that Free can return the buffer to the correct pool.
	pool Pool
}

var _ Interface = (*Buffer)(nil)

// New creates a standalone Buffer with the provided initial capacity.
//
// If capacity is negative, it is treated as 0. The returned buffer is not
// attached to a Pool; calling Free on it is a no-op.
func New(capacity int) *Buffer {
	return NewWithPool(Pool{}, capacity)
}

// NewWithPool creates a Buffer with the provided initial capacity and
// originating pool.
//
// This constructor is intended for pooling infrastructure. Most callers SHOULD
// acquire buffers with Pool.Get. If capacity is negative, it is treated as 0.
func NewWithPool(pool Pool, capacity int) *Buffer {
	return &Buffer{
		data: make([]byte, 0, normalizeCapacity(capacity)),
		pool: pool,
	}
}
