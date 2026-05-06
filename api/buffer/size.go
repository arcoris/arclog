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

const (
	// Size is the default initial capacity, in bytes, used by NewPool.
	//
	// Size is a starting capacity, not a hard limit. A Buffer may grow beyond
	// this value when callers append larger payloads. Code that requires a
	// different initial size SHOULD use NewPoolWithCapacity.
	Size = 1024

	// MaxRetainedSize is the largest buffer capacity, in bytes, that the default
	// Pool retains after Put or Free.
	//
	// Buffers larger than this value are discarded instead of being retained by
	// the arcoris.dev/pool-backed pool. This prevents a single unusually large
	// log record from pinning a large backing array for the lifetime of a
	// long-running process.
	MaxRetainedSize = 64 * 1024
)
