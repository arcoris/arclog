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

const (
	// DefaultInitialCapacity is the initial capacity used for newly allocated
	// buffers when Options.InitialCapacity is not positive.
	//
	// The value is a starting point, not a limit. Buffers may grow beyond this
	// capacity while owned by a caller.
	DefaultInitialCapacity = 1024

	// DefaultMaxRetainedCapacity is the largest buffer capacity retained by a Pool
	// when Options.MaxRetainedCapacity is not positive.
	//
	// Buffers whose capacity is greater than the effective retained capacity are
	// discarded on Put instead of being stored for reuse.
	DefaultMaxRetainedCapacity = 64 * 1024
)

// Options configures Pool construction.
//
// Options are copied and normalized by NewWithOptions. Mutating an Options value
// after Pool construction has no effect. Non-positive values request package
// defaults; callers should not use negative values as sentinels for "disable
// pooling" because this package always returns usable buffers.
type Options struct {
	// InitialCapacity is the capacity used for newly allocated buffers.
	//
	// Values less than or equal to zero are replaced with
	// DefaultInitialCapacity.
	InitialCapacity int

	// MaxRetainedCapacity is the largest buffer capacity retained on Put.
	//
	// Values less than or equal to zero are replaced with
	// DefaultMaxRetainedCapacity. Values smaller than the effective
	// InitialCapacity are raised to the effective InitialCapacity so a pool can
	// retain the buffers it creates by default.
	MaxRetainedCapacity int
}

// normalizeOptions applies package defaults and keeps Pool construction
// independent from caller-owned Options values.
//
// The retention ceiling is never allowed to fall below InitialCapacity. Without
// that invariant, a pool configured with both capacities could allocate buffers
// that it immediately refuses to retain on the first Put.
func normalizeOptions(options Options) Options {
	options.InitialCapacity = normalizePositive(
		options.InitialCapacity,
		DefaultInitialCapacity,
	)
	options.MaxRetainedCapacity = normalizePositive(
		options.MaxRetainedCapacity,
		DefaultMaxRetainedCapacity,
	)

	if options.MaxRetainedCapacity < options.InitialCapacity {
		options.MaxRetainedCapacity = options.InitialCapacity
	}

	return options
}

// normalizePositive returns value when it is positive and fallback otherwise.
func normalizePositive(value, fallback int) int {
	if value > 0 {
		return value
	}
	return fallback
}
