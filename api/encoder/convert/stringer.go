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

package convert

import (
	"fmt"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
	"arcoris.dev/arclog/api/internal/nilx"
)

// AddStringer appends value as a string field.
//
// Nil and typed-nil values are encoded as an empty string. Non-nil values are
// converted by calling value.String directly; panics from String propagate to
// the caller. AddStringer does not recover, wrap, or replace panics because safe
// conversion is a runtime policy.
//
// AddStringer returns the buffer returned by enc.AddString. Callers must
// continue with that returned buffer rather than assuming dst is still
// authoritative.
func AddStringer(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, value fmt.Stringer) *buffer.Buffer {
	if nilx.IsNil(value) {
		return enc.AddString(dst, key, "")
	}

	return enc.AddString(dst, key, value.String())
}

// AppendStringer appends value as a string array element.
//
// AppendStringer follows the same nil, typed-nil, strict panic propagation, and
// returned-buffer semantics as AddStringer.
func AppendStringer(dst *buffer.Buffer, enc encoder.ArrayEncoder, value fmt.Stringer) *buffer.Buffer {
	if nilx.IsNil(value) {
		return enc.AppendString(dst, "")
	}

	return enc.AppendString(dst, value.String())
}
