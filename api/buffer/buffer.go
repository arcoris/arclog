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

// Buffer is a growable byte buffer for encoded log records and other serialized
// payloads produced by the logging pipeline.
//
// The zero value is ready to use. Buffer owns only mutable byte accumulation and
// logical reset semantics; pooling, object lifecycle, and retention policy
// belong to runtime packages that depend on this API package.
//
// Buffer is not safe for concurrent use. A caller that shares a Buffer between
// goroutines MUST synchronize all method calls and all access to slices returned
// by Bytes.
type Buffer struct {
	// data stores the current logical contents of the buffer. len(data) is the
	// number of valid bytes and cap(data) is reusable storage.
	data []byte
}

// New creates a Buffer with the provided initial capacity.
//
// If capacity is negative, it is treated as 0. The returned buffer is not
// attached to a pool; runtime pooling packages should allocate or reset buffers
// explicitly around this concrete type.
func New(capacity int) *Buffer {
	return &Buffer{
		data: make([]byte, 0, normalizeCapacity(capacity)),
	}
}

// normalizeCapacity keeps public constructor behavior explicit without exposing
// a runtime allocation policy knob.
func normalizeCapacity(capacity int) int {
	if capacity < 0 {
		return 0
	}
	return capacity
}
