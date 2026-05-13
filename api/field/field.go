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

// Field is a compact descriptor for one structured log field.
//
// Key stores the field name. Type identifies which storage slot is meaningful.
// Integer stores bools, signed integers, unsigned integer bit patterns, float
// bit patterns, duration nanoseconds, and compact time Unix nanoseconds. String
// stores string values. Bytes stores borrowed byte slices for BytesType.
// Interface stores slow or special values, including errors, stringers,
// reflected values, full time values, and other non-inline payloads.
//
// Unsigned integer constructors store a bit-preserving representation in
// Integer. Encoders must recover the value according to Type, for example
// uint64(f.Integer) for Uint64Type.
//
// Float constructors store IEEE 754 bits in Integer.
type Field struct {
	// Key is the structured field name.
	//
	// Constructors store the key exactly as provided. They do not reject empty
	// keys, normalize dotted names, sanitize UTF-8, or apply redaction policy.
	Key string

	// Type selects the active storage slot and value interpretation.
	//
	// The zero value is SkipType, so the zero Field is a skipped descriptor.
	Type Type

	// Integer stores compact numeric and time-like payloads.
	//
	// BoolType uses 0 or 1. Signed integer types store their signed value.
	// Unsigned integer types store a bit-preserving int64 representation.
	// Float32Type and Float64Type store IEEE 754 bits. DurationType stores
	// nanoseconds. TimeType stores Unix nanoseconds.
	Integer int64

	// String stores the payload for StringType.
	String string

	// Bytes stores the borrowed payload for BytesType.
	//
	// The slice aliases caller-owned memory. Encoders that retain a field
	// beyond the current log call or context-binding operation must copy or
	// encode the bytes before returning.
	Bytes []byte

	// Interface stores slow-path and special payloads.
	//
	// ErrorType stores error values. StringerType stores fmt.Stringer values.
	// ReflectType stores arbitrary values. TimeFullType stores time.Time values
	// that cannot be represented by UnixNano. TimeType stores the original
	// *time.Location here while Integer stores Unix nanoseconds.
	Interface any
}

// IsSkip reports whether f is a no-op field descriptor.
//
// A skipped field carries no structured value and should be ignored by runtime
// encoders. The zero value of Field reports true.
func (f Field) IsSkip() bool { return f.Type == SkipType }

// IsNull reports whether f is an explicit null field descriptor.
//
// Null is different from Skip: Skip means "no field", while Null means "a
// field with this key and a null value".
func (f Field) IsNull() bool { return f.Type == NullType }
