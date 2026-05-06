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

// Stringer is the minimal fmt.Stringer-compatible contract used by helpers.
//
// The interface is defined locally so callers can use these helpers without
// importing fmt at call sites.
type Stringer interface {
	// String returns the textual representation to encode.
	String() string
}

// AddStringer appends a fmt.Stringer-compatible value as a string field.
//
// Nil and typed-nil values are encoded as an empty string. Non-nil values are
// converted by calling String directly; panics from String propagate to the
// caller.
func AddStringer(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, value Stringer) *buffer.Buffer {
	if IsNil(value) {
		return enc.AddString(dst, key, "")
	}

	return enc.AddString(dst, key, value.String())
}

// AppendStringer appends a fmt.Stringer-compatible value as a string element.
//
// It follows the same nil and panic behavior as AddStringer.
func AppendStringer(dst *buffer.Buffer, enc encoder.ArrayEncoder, value Stringer) *buffer.Buffer {
	if IsNil(value) {
		return enc.AppendString(dst, "")
	}

	return enc.AppendString(dst, value.String())
}

// AddStringerSafe appends a stringer field and converts panic from String into
// a diagnostic string.
//
// The recovered value is encoded as "PANIC=<value>".
func AddStringerSafe(dst *buffer.Buffer, enc encoder.ObjectEncoder, key string, value Stringer) (out *buffer.Buffer) {
	defer func() {
		if recovered := recover(); recovered != nil {
			out = enc.AddString(dst, key, fmt.Sprintf("PANIC=%v", recovered))
		}
	}()

	return AddStringer(dst, enc, key, value)
}
