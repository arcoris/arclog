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
	"math"
	"time"
)

// Bool constructs a boolean field.
func Bool(key string, value bool) Field {
	var integer int64
	if value {
		integer = 1
	}
	return Field{Key: key, Type: BoolType, Integer: integer}
}

// String constructs a string field.
func String(key string, value string) Field {
	return Field{Key: key, Type: StringType, String: value}
}

// Int constructs an int field.
func Int(key string, value int) Field {
	return Field{Key: key, Type: IntType, Integer: int64(value)}
}

// Int8 constructs an int8 field.
func Int8(key string, value int8) Field {
	return Field{Key: key, Type: Int8Type, Integer: int64(value)}
}

// Int16 constructs an int16 field.
func Int16(key string, value int16) Field {
	return Field{Key: key, Type: Int16Type, Integer: int64(value)}
}

// Int32 constructs an int32 field.
func Int32(key string, value int32) Field {
	return Field{Key: key, Type: Int32Type, Integer: int64(value)}
}

// Int64 constructs an int64 field.
func Int64(key string, value int64) Field {
	return Field{Key: key, Type: Int64Type, Integer: value}
}

// Uint constructs a uint field.
func Uint(key string, value uint) Field {
	return Field{Key: key, Type: UintType, Integer: int64(value)}
}

// Uint8 constructs a uint8 field.
func Uint8(key string, value uint8) Field {
	return Field{Key: key, Type: Uint8Type, Integer: int64(value)}
}

// Uint16 constructs a uint16 field.
func Uint16(key string, value uint16) Field {
	return Field{Key: key, Type: Uint16Type, Integer: int64(value)}
}

// Uint32 constructs a uint32 field.
func Uint32(key string, value uint32) Field {
	return Field{Key: key, Type: Uint32Type, Integer: int64(value)}
}

// Uint64 constructs a uint64 field.
func Uint64(key string, value uint64) Field {
	return Field{Key: key, Type: Uint64Type, Integer: int64(value)}
}

// Float32 constructs a float32 field preserving the IEEE 754 bit pattern.
func Float32(key string, value float32) Field {
	return Field{Key: key, Type: Float32Type, Integer: int64(math.Float32bits(value))}
}

// Float64 constructs a float64 field preserving the IEEE 754 bit pattern.
func Float64(key string, value float64) Field {
	return Field{Key: key, Type: Float64Type, Integer: int64(math.Float64bits(value))}
}

// Duration constructs a duration field using nanoseconds.
func Duration(key string, value time.Duration) Field {
	return Field{Key: key, Type: DurationType, Integer: int64(value)}
}
