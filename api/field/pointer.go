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

import "time"

// BoolPtr constructs a field from *bool.
func BoolPtr(key string, value *bool) Field {
	if value == nil {
		return Null(key)
	}
	return Bool(key, *value)
}

// DurationPtr constructs a field from *time.Duration.
func DurationPtr(key string, value *time.Duration) Field {
	if value == nil {
		return Null(key)
	}
	return Duration(key, *value)
}

// Float64Ptr constructs a field from *float64.
func Float64Ptr(key string, value *float64) Field {
	if value == nil {
		return Null(key)
	}
	return Float64(key, *value)
}

// Float32Ptr constructs a field from *float32.
func Float32Ptr(key string, value *float32) Field {
	if value == nil {
		return Null(key)
	}
	return Float32(key, *value)
}

// IntPtr constructs a field from *int.
func IntPtr(key string, value *int) Field {
	if value == nil {
		return Null(key)
	}
	return Int(key, *value)
}

// Int8Ptr constructs a field from *int8.
func Int8Ptr(key string, value *int8) Field {
	if value == nil {
		return Null(key)
	}
	return Int8(key, *value)
}

// Int16Ptr constructs a field from *int16.
func Int16Ptr(key string, value *int16) Field {
	if value == nil {
		return Null(key)
	}
	return Int16(key, *value)
}

// Int32Ptr constructs a field from *int32.
func Int32Ptr(key string, value *int32) Field {
	if value == nil {
		return Null(key)
	}
	return Int32(key, *value)
}

// Int64Ptr constructs a field from *int64.
func Int64Ptr(key string, value *int64) Field {
	if value == nil {
		return Null(key)
	}
	return Int64(key, *value)
}

// StringPtr constructs a field from *string.
func StringPtr(key string, value *string) Field {
	if value == nil {
		return Null(key)
	}
	return String(key, *value)
}

// TimePtr constructs a field from *time.Time.
func TimePtr(key string, value *time.Time) Field {
	if value == nil {
		return Null(key)
	}
	return Time(key, *value)
}

// UintPtr constructs a field from *uint.
func UintPtr(key string, value *uint) Field {
	if value == nil {
		return Null(key)
	}
	return Uint(key, *value)
}

// Uint8Ptr constructs a field from *uint8.
func Uint8Ptr(key string, value *uint8) Field {
	if value == nil {
		return Null(key)
	}
	return Uint8(key, *value)
}

// Uint16Ptr constructs a field from *uint16.
func Uint16Ptr(key string, value *uint16) Field {
	if value == nil {
		return Null(key)
	}
	return Uint16(key, *value)
}

// Uint32Ptr constructs a field from *uint32.
func Uint32Ptr(key string, value *uint32) Field {
	if value == nil {
		return Null(key)
	}
	return Uint32(key, *value)
}

// Uint64Ptr constructs a field from *uint64.
func Uint64Ptr(key string, value *uint64) Field {
	if value == nil {
		return Null(key)
	}
	return Uint64(key, *value)
}
