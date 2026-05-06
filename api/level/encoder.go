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

package level

import "arcoris.dev/arclog/api/buffer"

// Encoder appends the representation of a Level to dst and returns the buffer
// that should be used for subsequent writes.
//
// Implementations should append to dst when possible. They may return a
// different buffer only if their own documented policy requires replacing the
// destination. The returned buffer is authoritative; callers must continue with
// that value after the encoder call.
//
// Encoder implementations must not release dst or the returned buffer. Buffer
// ownership always remains with the caller.
type Encoder func(dst *buffer.Buffer, lvl Level) *buffer.Buffer
