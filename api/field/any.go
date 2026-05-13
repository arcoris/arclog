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
	"fmt"
	"time"
)

// Any chooses a typed field constructor for value and falls back to Reflect.
//
// Any is a convenience boundary for dynamic values. It does not introduce an
// AnyType storage kind: supported concrete Go types become the same descriptors
// produced by the typed constructors, and unsupported values become Reflect.
// Hot paths with statically known value types should call typed constructors
// directly, especially for []byte where interface boxing can allocate before
// Any receives the value.
func Any(key string, value any) Field {
	if value == nil {
		return Null(key)
	}

	switch value := value.(type) {
	case bool:
		return Bool(key, value)
	case *bool:
		return BoolPtr(key, value)
	case []byte:
		return Bytes(key, value)
	case time.Duration:
		return Duration(key, value)
	case *time.Duration:
		return DurationPtr(key, value)
	case float64:
		return Float64(key, value)
	case *float64:
		return Float64Ptr(key, value)
	case float32:
		return Float32(key, value)
	case *float32:
		return Float32Ptr(key, value)
	case int:
		return Int(key, value)
	case *int:
		return IntPtr(key, value)
	case int64:
		return Int64(key, value)
	case *int64:
		return Int64Ptr(key, value)
	case int32:
		return Int32(key, value)
	case *int32:
		return Int32Ptr(key, value)
	case int16:
		return Int16(key, value)
	case *int16:
		return Int16Ptr(key, value)
	case int8:
		return Int8(key, value)
	case *int8:
		return Int8Ptr(key, value)
	case string:
		return String(key, value)
	case *string:
		return StringPtr(key, value)
	case time.Time:
		return Time(key, value)
	case *time.Time:
		return TimePtr(key, value)
	case uint:
		return Uint(key, value)
	case *uint:
		return UintPtr(key, value)
	case uint64:
		return Uint64(key, value)
	case *uint64:
		return Uint64Ptr(key, value)
	case uint32:
		return Uint32(key, value)
	case *uint32:
		return Uint32Ptr(key, value)
	case uint16:
		return Uint16(key, value)
	case *uint16:
		return Uint16Ptr(key, value)
	case uint8:
		return Uint8(key, value)
	case *uint8:
		return Uint8Ptr(key, value)
	case error:
		return NamedError(key, value)
	case fmt.Stringer:
		return Stringer(key, value)
	default:
		return Reflect(key, value)
	}
}
