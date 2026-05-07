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
	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
	"arcoris.dev/arclog/api/internal/nilx"
)

// AddError appends err as a string field.
//
// Nil and typed-nil errors are encoded as an empty string. Non-nil errors are
// converted by calling err.Error directly; panics from Error propagate to the
// caller. AddError does not recover, wrap, classify, or format panics because
// those choices belong to runtime policy.
//
// AddError returns the buffer returned by enc.AddString. Callers must continue
// with that returned buffer rather than assuming dst is still authoritative.
func AddError(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, err error) *buffer.Buffer {
	if nilx.IsNil(err) {
		return enc.AddString(dst, key, "")
	}

	return enc.AddString(dst, key, err.Error())
}

// AppendError appends err as a string array element.
//
// AppendError follows the same nil, typed-nil, strict panic propagation, and
// returned-buffer semantics as AddError.
func AppendError(dst *buffer.Buffer, enc encoder.ArrayEncoder, err error) *buffer.Buffer {
	if nilx.IsNil(err) {
		return enc.AppendString(dst, "")
	}

	return enc.AppendString(dst, err.Error())
}
