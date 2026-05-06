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

package encoders

import (
	"fmt"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
)

// AddError appends an error value as a string field.
//
// Nil and typed-nil errors are encoded as an empty string. This keeps the
// helper total and prevents logging paths from panicking on nil error values.
func AddError(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, err error) *buffer.Buffer {
	if IsNil(err) {
		return enc.AddString(dst, key, "")
	}

	return enc.AddString(dst, key, err.Error())
}

// AppendError appends an error value as a string array element.
func AppendError(dst *buffer.Buffer, enc encoder.ArrayEncoder, err error) *buffer.Buffer {
	if IsNil(err) {
		return enc.AppendString(dst, "")
	}

	return enc.AppendString(dst, err.Error())
}

// AddErrorSafe appends an error field and converts panic from Error into a
// diagnostic string.
//
// Error implementations should not panic, but logging must not let a broken
// error value crash the log path unless the caller explicitly chooses that
// policy.
func AddErrorSafe(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, err error) (out *buffer.Buffer) {
	defer func() {
		if recovered := recover(); recovered != nil {
			out = enc.AddString(dst, key, fmt.Sprintf("PANIC=%v", recovered))
		}
	}()

	return AddError(dst, enc, key, err)
}
