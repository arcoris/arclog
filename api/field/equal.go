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
	"bytes"
	"reflect"
)

// Equal reports whether f and other have the same logical field content.
//
// Byte-string fields compare bytes rather than slice identity. Payloads that
// are commonly non-comparable, such as marshalers, reflected values, errors,
// stringers, and full time values, are compared with reflect.DeepEqual. Equal is
// intended for tests and contract checks, not hot-path logging.
func (f Field) Equal(other Field) bool {
	if f.Key != other.Key || f.Type != other.Type || f.Integer != other.Integer || f.String != other.String {
		return false
	}
	switch f.Type {
	case ByteStringType:
		return bytes.Equal(f.Interface.([]byte), other.Interface.([]byte))
	case ObjectMarshalerType, ArrayMarshalerType, InlineMarshalerType, ReflectType, StringerType, ErrorType, TimeFullType:
		return reflect.DeepEqual(f.Interface, other.Interface)
	default:
		return f.Interface == other.Interface
	}
}
