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

package field

import (
	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
)

// Fields is an ordered collection of fields that can marshal itself as an
// object.
//
// The slice is retained by value only; the Field values and any payloads they
// reference keep their normal ownership semantics.
type Fields []Field

// MarshalLogObject appends all fields in fs to enc in order.
//
// The method stops at the first field error and returns the buffer produced up
// to that point. Callers must continue with the returned buffer even when err is
// non-nil.
func (fs Fields) MarshalLogObject(dst *buffer.Buffer, enc encoder.ObjectEncoder) (*buffer.Buffer, error) {
	var err error
	for i := range fs {
		dst, err = fs[i].AddTo(dst, enc)
		if err != nil {
			return dst, err
		}
	}
	return dst, nil
}

// Dict constructs an object field from an ordered list of fields.
//
// The variadic slice is retained through Fields; callers should not mutate
// caller-owned field payloads until the entry is encoded.
func Dict(key string, fields ...Field) Field { return Object(key, Fields(fields)) }

// DictObject returns an encoder.ObjectMarshaler backed by fields.
//
// It is useful when a caller needs an object marshaler directly rather than a
// keyed object field.
func DictObject(fields ...Field) encoder.ObjectMarshaler { return Fields(fields) }
