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

package encoder

import "arcoris.dev/arclog/api/buffer"

// ObjectMarshaler appends object fields through an ObjectEncoder.
//
// MarshalLogObject receives caller-owned buffer storage and must return the
// buffer that should be used after marshaling. Implementations must not release
// dst. Any returned error is propagated by the concrete encoder or field
// dispatch path that invoked the marshaler.
type ObjectMarshaler interface {
	// MarshalLogObject appends this value's object fields to dst through enc.
	MarshalLogObject(dst *buffer.Buffer, enc ObjectEncoder) (*buffer.Buffer, error)
}

// ObjectMarshalerFunc adapts a function to ObjectMarshaler.
//
// A nil ObjectMarshalerFunc is invalid and will panic when MarshalLogObject is
// called.
type ObjectMarshalerFunc func(*buffer.Buffer, ObjectEncoder) (*buffer.Buffer, error)

// MarshalLogObject calls f(dst, enc).
func (f ObjectMarshalerFunc) MarshalLogObject(dst *buffer.Buffer, enc ObjectEncoder) (*buffer.Buffer, error) {
	return f(dst, enc)
}

// ArrayMarshaler appends array elements through an ArrayEncoder.
//
// MarshalLogArray follows the same buffer ownership and error propagation rules
// as ObjectMarshaler, but emits positional array elements instead of keyed
// fields.
type ArrayMarshaler interface {
	// MarshalLogArray appends this value's array elements to dst through enc.
	MarshalLogArray(dst *buffer.Buffer, enc ArrayEncoder) (*buffer.Buffer, error)
}

// ArrayMarshalerFunc adapts a function to ArrayMarshaler.
//
// A nil ArrayMarshalerFunc is invalid and will panic when MarshalLogArray is
// called.
type ArrayMarshalerFunc func(*buffer.Buffer, ArrayEncoder) (*buffer.Buffer, error)

// MarshalLogArray calls f(dst, enc).
func (f ArrayMarshalerFunc) MarshalLogArray(dst *buffer.Buffer, enc ArrayEncoder) (*buffer.Buffer, error) {
	return f(dst, enc)
}
