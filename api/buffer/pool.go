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

import objectpool "arcoris.dev/pool"

// Pool reuses Buffer values between encoded log records.
//
// Pool is the buffer package's domain-specific adapter over arcoris.dev/pool.
// The generic pool package owns the reusable-object lifecycle: construction,
// reuse admission, reset-on-return, and backend storage. This package adds only
// Buffer-specific policy:
//
//   - freshly constructed buffers start with the configured initial capacity;
//   - returned buffers whose backing array grew beyond the retention limit are
//     dropped instead of being retained;
//   - retained buffers are reset before the next caller can acquire them;
//   - buffers remember their originating Pool so Buffer.Free can delegate to
//     the correct return path.
//
// A Pool created by NewPool or NewPoolWithCapacity is safe for concurrent Get
// and Put calls. Individual Buffer values remain single-owner mutable objects:
// callers MUST NOT use a buffer after Put or Free returns.
//
// The zero value of Pool is intentionally safe but non-pooling. A zero-value
// Pool.Get returns a standalone buffer and Pool.Put discards it. This wrapper
// policy keeps accidental zero-value use from panicking while preserving the
// stricter construction contract of the shared arcoris.dev/pool package for
// explicitly configured pools.
type Pool struct {
	// pool is nil for the safe non-pooling zero value.
	pool *objectpool.Pool[*Buffer]

	// initialCapacity is the capacity used only for newly allocated buffers.
	initialCapacity int
	// maxRetainedCapacity is the capacity ceiling checked on the return path.
	maxRetainedCapacity int
}

// NewPool creates a Pool with the default initial buffer capacity.
func NewPool() Pool {
	return NewPoolWithCapacity(Size)
}

// NewPoolWithCapacity creates a Pool whose newly allocated buffers start with
// at least capacity bytes of backing storage.
//
// A negative capacity is treated as 0. The value controls only the initial
// backing capacity of freshly allocated buffers; buffers may grow beyond it
// while in use. Oversized grown buffers are filtered by MaxRetainedSize on the
// return path.
func NewPoolWithCapacity(capacity int) Pool {
	initialCapacity := normalizeCapacity(capacity)

	p := Pool{
		initialCapacity:     initialCapacity,
		maxRetainedCapacity: MaxRetainedSize,
	}

	p.pool = objectpool.New(objectpool.Options[*Buffer]{
		New: func() *Buffer {
			return &Buffer{
				data: make([]byte, 0, initialCapacity),
				pool: p,
			}
		},
		Reset: func(buf *Buffer) {
			buf.Reset()
			buf.pool = p
		},
		Reuse: func(buf *Buffer) bool {
			return buf != nil && cap(buf.data) <= retainedCapacityLimit(p.maxRetainedCapacity)
		},
	})

	return p
}

// Get obtains an empty Buffer from the pool.
//
// The returned buffer is owned by the caller until it is released with Free or
// Put. Buffers acquired from a configured Pool come from arcoris.dev/pool and
// are already reset on the previous Put path; Get still normalizes the buffer's
// origin to this Pool so Free remains correct after value-copying the Pool.
func (p Pool) Get() *Buffer {
	if p.pool == nil {
		return NewWithPool(p, 0)
	}

	buf := p.pool.Get()
	buf.pool = p
	return buf
}

// Put releases buf back to the pool.
//
// Put is a no-op for nil buffers and for zero-value pools. For configured
// pools, Put delegates the return path to arcoris.dev/pool: the shared pool
// evaluates the reuse policy, resets accepted buffers, and stores them in the
// backend for possible reuse. Buffers rejected by the retention policy are
// discarded.
func (p Pool) Put(buf *Buffer) {
	if buf == nil || p.pool == nil {
		return
	}

	p.pool.Put(buf)
}

// normalizeCapacity keeps all public constructors aligned on negative-capacity
// handling without exporting a policy knob.
func normalizeCapacity(capacity int) int {
	if capacity < 0 {
		return 0
	}
	return capacity
}

// retainedCapacityLimit preserves the package default when an internal Pool
// value carries an unset or invalid retention limit.
func retainedCapacityLimit(limit int) int {
	if limit <= 0 {
		return MaxRetainedSize
	}
	return limit
}
