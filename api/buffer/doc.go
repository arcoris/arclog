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

// Package buffer provides a small mutable byte accumulator for encoder hot paths.
//
// Buffer is intentionally narrower than bytes.Buffer. It does not implement
// io.Writer and does not own pooling, synchronization, writer behavior,
// release/free semantics, or encoding policy. Runtime components decide when a
// Buffer is acquired, reused, released, or dropped.
//
// The zero value is ready to use.
//
// A Buffer is not safe for concurrent use.
//
// Bytes returned by Buffer.Bytes are borrowed. Callers must treat them as
// read-only and must not retain them after the next mutation, Reset, Truncate,
// Grow that reallocates, or runtime release by an owner outside this package.
package buffer
