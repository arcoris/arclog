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

package bufferpool

import (
	objectpool "arcoris.dev/pool"

	"arcoris.dev/arclog/api/buffer"
)

// Pool reuses Buffer values for runtime encoding paths.
//
// Pool is safe for concurrent Get and Put calls. A Buffer returned by Get is
// still a single-owner mutable object; Pool does not make that Buffer safe for
// concurrent mutation.
//
// The zero value is safe but non-pooling: Get returns a standalone default-size
// buffer and Put discards its argument. Runtime code that wants reuse should
// construct pools with New or NewWithOptions so the shared backend and
// normalized retention policy are installed together.
type Pool struct {
	// backend owns construction, storage, reset-on-return, and reuse admission.
	//
	// Pool keeps this field unexported so runtime code depends on the small
	// bufferpool contract instead of the generic object-pool backend.
	backend *objectpool.Pool[*buffer.Buffer]

	// initialCapacity is applied only when backend creates a new buffer.
	//
	// It is stored to make the construction policy explicit inside the pool and
	// to keep normalized Options attached to the runtime object that uses them.
	initialCapacity int

	// maxRetainedCapacity is the capacity ceiling used by shouldRetain.
	//
	// Buffers above this capacity are discarded on Put by the backend reuse
	// predicate instead of being retained for future log records.
	maxRetainedCapacity int
}

// New creates a Pool with default options.
//
// The returned pool uses DefaultInitialCapacity for new buffers and
// DefaultMaxRetainedCapacity as the retention ceiling.
func New() *Pool {
	return NewWithOptions(Options{})
}

// NewWithOptions creates a Pool with options.
//
// Options are normalized and copied during construction. The returned Pool is
// ready for concurrent Get and Put calls.
func NewWithOptions(options Options) *Pool {
	options = normalizeOptions(options)

	p := &Pool{
		initialCapacity:     options.InitialCapacity,
		maxRetainedCapacity: options.MaxRetainedCapacity,
	}

	p.backend = objectpool.New(objectpool.Options[*buffer.Buffer]{
		New: func() *buffer.Buffer {
			return buffer.New(options.InitialCapacity)
		},
		Reset: resetBufferForReuse,
		Reuse: func(buf *buffer.Buffer) bool {
			return shouldRetain(buf, options.MaxRetainedCapacity)
		},
		OnDrop: dropBuffer,
	})

	return p
}

// Get obtains an empty Buffer from the pool.
//
// The caller owns the returned buffer until passing it to Put. Callers MUST NOT
// assume a specific backing capacity beyond the configured initial allocation
// policy for newly constructed buffers.
func (p *Pool) Get() *buffer.Buffer {
	if p == nil || p.backend == nil {
		return buffer.New(DefaultInitialCapacity)
	}

	return p.backend.Get()
}

// Put releases buf back to the pool.
//
// Put is a no-op for nil buffers. For non-nil buffers, caller ownership ends
// when Put is called; callers MUST NOT use buf or any slices previously returned
// by buf.Bytes after Put returns.
func (p *Pool) Put(buf *buffer.Buffer) {
	if buf == nil {
		return
	}
	if p == nil || p.backend == nil {
		return
	}

	p.backend.Put(buf)
}

// resetBufferForReuse prepares an accepted buffer for future owners.
//
// arcoris.dev/pool calls Reset only after Reuse accepts the value. Keeping this
// logic separate makes that lifecycle order visible: accepted buffers keep their
// backing array, but their logical contents and borrowed Bytes views end here.
func resetBufferForReuse(buf *buffer.Buffer) {
	if buf != nil {
		buf.Reset()
	}
}

// dropBuffer clears a buffer rejected by the retention policy.
//
// Rejected buffers are not stored by the backend, so this function cannot and
// does not try to make their backing arrays reusable. It still resets logical
// contents before the object is discarded so callers cannot accidentally observe
// stale payload bytes through a value they no longer own.
func dropBuffer(buf *buffer.Buffer) {
	if buf != nil {
		buf.Reset()
	}
}

// shouldRetain reports whether buf is small enough to return to the backend
// pool.
//
// Nil buffers are rejected defensively. Valid callers normally use Put, which
// handles nil before calling the backend. arcoris.dev/pool evaluates this
// predicate before Reset, so it must inspect the current capacity directly.
func shouldRetain(buf *buffer.Buffer, maxRetainedCapacity int) bool {
	return buf != nil && buf.Cap() <= maxRetainedCapacity
}
