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

	"arcoris.dev/arclog/api/encoder"
)

// Any chooses a field constructor for value and falls back to Reflect when no
// specific constructor applies.
//
// Any is a convenience boundary for call sites that already hold dynamic values.
// Hot paths should prefer typed constructors so dispatch remains explicit and
// avoids avoidable interface matching. Slice, marshaler, error, and Stringer
// ownership follows the constructor selected for value.
func Any(key string, value any) Field {
	switch value := value.(type) {
	case encoder.ObjectMarshaler:
		return Object(key, value)
	case encoder.ArrayMarshaler:
		return Array(key, value)
	case Fields:
		return Object(key, value)
	case []Field:
		return Dict(key, value...)
	case bool:
		return Bool(key, value)
	case *bool:
		return BoolPtr(key, value)
	case []byte:
		return ByteString(key, value)
	case complex128:
		return Complex128(key, value)
	case *complex128:
		return Complex128Ptr(key, value)
	case complex64:
		return Complex64(key, value)
	case *complex64:
		return Complex64Ptr(key, value)
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
	case uintptr:
		return Uintptr(key, value)
	case *uintptr:
		return UintptrPtr(key, value)
	case error:
		return Error(key, value)
	case fmt.Stringer:
		return Stringer(key, value)
	default:
		return Reflect(key, value)
	}
}
