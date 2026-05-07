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

package stack

import "arcoris.dev/arclog/api/buffer"

// Encoder appends a representation of a Stack to dst and returns the buffer
// that should be used for subsequent writes.
//
// Encoder is a value-specific append contract. It is not a full log-entry
// encoder and it does not define JSON, console, or binary stack formatting.
// Runtime encoders choose the concrete representation: multiline text,
// structured frame arrays, function-only frames, file:line pairs, or another
// format.
//
// Implementations should append to dst when possible. They may return a
// different buffer only if their own documented buffering policy requires it.
// The returned buffer is authoritative; callers must continue with that value
// after the encoder call.
//
// Encoder implementations must not call Free on dst or on the returned buffer.
// Buffer lifetime remains the caller's responsibility.
type Encoder func(dst *buffer.Buffer, stack Stack) *buffer.Buffer
