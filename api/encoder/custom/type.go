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

package custom

import (
	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
)

// TypeEncoder encodes a value that is not represented by one of the built-in
// primitive field kinds.
//
// TypeEncoder is a low-level extension point for concrete encoders. The value
// is passed as any because the type selection policy belongs to the caller or
// registry that invoked this encoder.
//
// Implementations MUST return the buffer that should be used after encoding and
// MUST NOT release dst. Returned errors propagate to the caller's field
// dispatch or runtime encoder path.
type TypeEncoder interface {
	// EncodeType appends value under key using enc's object-field contract.
	EncodeType(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, value any) (*buffer.Buffer, error)
}

// TypeEncoderFunc adapts a function to TypeEncoder.
//
// A nil TypeEncoderFunc is invalid and will panic when EncodeType is called.
type TypeEncoderFunc func(*buffer.Buffer, encoder.ObjectEncoder, string, any) (*buffer.Buffer, error)

// EncodeType calls f(dst, enc, key, value).
func (f TypeEncoderFunc) EncodeType(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, value any) (*buffer.Buffer, error) {
	return f(dst, enc, key, value)
}
