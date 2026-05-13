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
	"time"
)

// Equal reports whether f and other have the same logical field content.
//
// Equal is intended for tests and API contract checks, not hot-path logging.
// BytesType compares Field.Bytes by contents. TimeFullType compares time.Time
// values with time.Time.Equal when both payloads have the expected type.
func (f Field) Equal(other Field) bool {
	if f.Key != other.Key || f.Type != other.Type || f.Integer != other.Integer || f.String != other.String {
		return false
	}

	switch f.Type {
	case BytesType:
		return bytes.Equal(f.Bytes, other.Bytes)
	case TimeType:
		return reflect.DeepEqual(f.Interface, other.Interface)
	case TimeFullType:
		left, lok := f.Interface.(time.Time)
		right, rok := other.Interface.(time.Time)
		if lok && rok {
			return left.Equal(right)
		}
		return reflect.DeepEqual(f.Interface, other.Interface)
	case ReflectType, StringerType, ErrorType:
		return reflect.DeepEqual(f.Interface, other.Interface)
	default:
		return f.Interface == other.Interface
	}
}
