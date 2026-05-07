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
	"math"
	"time"

	"arcoris.dev/arclog/api/encoder"
	"arcoris.dev/arclog/api/internal/nilx"
)

const (
	minUnixNano = -1 << 63
	maxUnixNano = 1<<63 - 1
)

var (
	// minTimeInt64 and maxTimeInt64 bound the compact Unix-nanosecond time
	// representation. Values outside this range keep the full time.Time value.
	minTimeInt64 = time.Unix(0, minUnixNano)
	maxTimeInt64 = time.Unix(0, maxUnixNano)
)

// Skip returns a no-op field that AddTo ignores.
//
// The zero value of Field is equivalent to Skip.
func Skip() Field { return Field{} }

// Nil constructs a field that explicitly carries a nil value through the
// encoder's reflection path.
//
// Concrete encoders decide how reflected nil values are represented.
func Nil(key string) Field { return Reflect(key, nil) }

// Bool constructs a boolean field without retaining caller-owned state.
func Bool(key string, value bool) Field {
	var i int64
	if value {
		i = 1
	}
	return Field{Key: key, Type: BoolType, Integer: i}
}

// ByteString constructs a byte-string field.
//
// The slice is retained, not copied. Callers must not mutate value until the log
// entry has been encoded. The constructor performs no escaping, validation, or
// UTF-8 checks; those decisions belong to the concrete encoder.
func ByteString(key string, value []byte) Field {
	return Field{Key: key, Type: ByteStringType, Interface: value}
}

// Complex128 constructs a complex128 field without retaining caller-owned
// mutable state.
func Complex128(key string, value complex128) Field {
	return Field{Key: key, Type: Complex128Type, Interface: value}
}

// Complex64 constructs a complex64 field without retaining caller-owned mutable
// state.
func Complex64(key string, value complex64) Field {
	return Field{Key: key, Type: Complex64Type, Interface: value}
}

// Duration constructs a time.Duration field using the duration's nanosecond
// count.
func Duration(key string, value time.Duration) Field {
	return Field{Key: key, Type: DurationType, Integer: int64(value)}
}

// Float64 constructs a float64 field preserving the IEEE 754 bit pattern,
// including NaN payloads.
func Float64(key string, value float64) Field {
	return Field{Key: key, Type: Float64Type, Integer: int64(math.Float64bits(value))}
}

// Float32 constructs a float32 field preserving the IEEE 754 bit pattern,
// including NaN payloads.
func Float32(key string, value float32) Field {
	return Field{Key: key, Type: Float32Type, Integer: int64(math.Float32bits(value))}
}

// Int constructs an int field.
func Int(key string, value int) Field { return Field{Key: key, Type: IntType, Integer: int64(value)} }

// Int64 constructs an int64 field.
func Int64(key string, value int64) Field { return Field{Key: key, Type: Int64Type, Integer: value} }

// Int32 constructs an int32 field.
func Int32(key string, value int32) Field {
	return Field{Key: key, Type: Int32Type, Integer: int64(value)}
}

// Int16 constructs an int16 field.
func Int16(key string, value int16) Field {
	return Field{Key: key, Type: Int16Type, Integer: int64(value)}
}

// Int8 constructs an int8 field.
func Int8(key string, value int8) Field {
	return Field{Key: key, Type: Int8Type, Integer: int64(value)}
}

// String constructs a string field.
func String(key string, value string) Field { return Field{Key: key, Type: StringType, String: value} }

// Time constructs a time.Time field.
//
// Values representable as Unix nanoseconds use compact storage plus the
// original location. Values outside that range retain the full time.Time value.
func Time(key string, value time.Time) Field {
	if value.Before(minTimeInt64) || value.After(maxTimeInt64) {
		return Field{Key: key, Type: TimeFullType, Interface: value}
	}
	return Field{Key: key, Type: TimeType, Integer: value.UnixNano(), Interface: value.Location()}
}

// Uint constructs a uint field.
//
// The value is stored in the field's signed integer slot and reconstructed as
// uint during AddTo.
func Uint(key string, value uint) Field {
	return Field{Key: key, Type: UintType, Integer: int64(value)}
}

// Uint64 constructs a uint64 field.
//
// All bits are preserved by storing the value in the field's signed integer
// slot and reconstructing it as uint64 during AddTo.
func Uint64(key string, value uint64) Field {
	return Field{Key: key, Type: Uint64Type, Integer: int64(value)}
}

// Uint32 constructs a uint32 field.
func Uint32(key string, value uint32) Field {
	return Field{Key: key, Type: Uint32Type, Integer: int64(value)}
}

// Uint16 constructs a uint16 field.
func Uint16(key string, value uint16) Field {
	return Field{Key: key, Type: Uint16Type, Integer: int64(value)}
}

// Uint8 constructs a uint8 field.
func Uint8(key string, value uint8) Field {
	return Field{Key: key, Type: Uint8Type, Integer: int64(value)}
}

// Uintptr constructs a uintptr field.
func Uintptr(key string, value uintptr) Field {
	return Field{Key: key, Type: UintptrType, Integer: int64(value)}
}

// Object constructs a field backed by an encoder.ObjectMarshaler.
//
// A nil or typed-nil marshaler is represented as Nil(key). Non-nil marshalers
// are retained by reference and are invoked later by AddTo.
func Object(key string, value encoder.ObjectMarshaler) Field {
	if nilx.IsNil(value) {
		return Nil(key)
	}
	return Field{Key: key, Type: ObjectMarshalerType, Interface: value}
}

// Array constructs a field backed by an encoder.ArrayMarshaler.
//
// A nil or typed-nil marshaler is represented as Nil(key). Non-nil marshalers
// are retained by reference and are invoked later by AddTo.
func Array(key string, value encoder.ArrayMarshaler) Field {
	if nilx.IsNil(value) {
		return Nil(key)
	}
	return Field{Key: key, Type: ArrayMarshalerType, Interface: value}
}

// Inline constructs a field that appends an object marshaler into the current
// encoder namespace.
//
// A nil or typed-nil marshaler becomes Skip because there is no key at which to
// encode an explicit nil value.
func Inline(value encoder.ObjectMarshaler) Field {
	if nilx.IsNil(value) {
		return Skip()
	}
	return Field{Type: InlineMarshalerType, Interface: value}
}

// Reflect constructs a field for the encoder's reflection path.
//
// Reflect is the compatibility fallback for values without a dedicated field
// constructor. It may allocate or fail depending on the concrete encoder.
func Reflect(key string, value any) Field {
	return Field{Key: key, Type: ReflectType, Interface: value}
}

// Namespace opens or selects a namespace named by key for subsequent fields.
//
// The exact representation of namespaces is encoder-defined; this field only
// requests the transition.
func Namespace(key string) Field { return Field{Key: key, Type: NamespaceType} }

// Stringer constructs a field backed by fmt.Stringer.
//
// A nil or typed-nil Stringer is represented as Nil(key). String conversion is
// deferred until AddTo and may allocate or panic according to the String method.
func Stringer(key string, value fmt.Stringer) Field {
	if nilx.IsNil(value) {
		return Nil(key)
	}
	return Field{Key: key, Type: StringerType, Interface: value}
}

// Error constructs a field backed by an error.
//
// A nil or typed-nil error is represented as Nil(key). Error string conversion
// is deferred until AddTo and may allocate or panic according to the error
// value.
func Error(key string, value error) Field {
	if nilx.IsNil(value) {
		return Nil(key)
	}
	return Field{Key: key, Type: ErrorType, Interface: value}
}
