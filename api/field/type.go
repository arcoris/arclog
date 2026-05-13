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

import "fmt"

// Type identifies the logical value kind stored in a Field.
//
// Type is part of the descriptor contract between API call sites and future
// runtime encoders. It does not name an output format. A JSON encoder, OTLP
// exporter, and binary stream encoder may all interpret the same Type
// differently, but they must read the same Field storage slots documented here.
type Type uint8

const (
	// SkipType marks a descriptor that should be ignored.
	SkipType Type = iota
	// NullType marks an explicit null value under Field.Key.
	NullType

	// BoolType stores a boolean in Field.Integer as 0 or 1.
	BoolType
	// BytesType stores a borrowed byte slice in Field.Bytes.
	BytesType

	// DurationType stores a time.Duration in Field.Integer as nanoseconds.
	DurationType

	// Float64Type stores math.Float64bits(value) in Field.Integer.
	Float64Type
	// Float32Type stores math.Float32bits(value) in Field.Integer.
	Float32Type

	// IntType stores an int converted to int64 in Field.Integer.
	IntType
	// Int8Type stores an int8 converted to int64 in Field.Integer.
	Int8Type
	// Int16Type stores an int16 converted to int64 in Field.Integer.
	Int16Type
	// Int32Type stores an int32 converted to int64 in Field.Integer.
	Int32Type
	// Int64Type stores an int64 in Field.Integer.
	Int64Type

	// StringType stores a string in Field.String.
	StringType

	// TimeType stores Unix nanoseconds in Field.Integer and location in
	// Field.Interface.
	TimeType
	// TimeFullType stores a full time.Time in Field.Interface.
	TimeFullType

	// UintType stores a uint bit pattern in Field.Integer.
	UintType
	// Uint8Type stores a uint8 converted to int64 in Field.Integer.
	Uint8Type
	// Uint16Type stores a uint16 converted to int64 in Field.Integer.
	Uint16Type
	// Uint32Type stores a uint32 converted to int64 in Field.Integer.
	Uint32Type
	// Uint64Type stores a uint64 bit pattern in Field.Integer.
	Uint64Type

	// ReflectType stores an arbitrary non-nil value in Field.Interface.
	ReflectType
	// NamespaceType marks the start or selection of a logical namespace.
	NamespaceType
	// StringerType stores a fmt.Stringer in Field.Interface.
	StringerType
	// ErrorType stores an error in Field.Interface.
	ErrorType
)

// String returns a stable lowercase diagnostic name for t.
//
// String is intended for tests, debugging, and configuration diagnostics. It is
// not a wire-format contract for encoded log records.
func (t Type) String() string {
	switch t {
	case SkipType:
		return "skip"
	case NullType:
		return "null"
	case BoolType:
		return "bool"
	case BytesType:
		return "bytes"
	case DurationType:
		return "duration"
	case Float64Type:
		return "float64"
	case Float32Type:
		return "float32"
	case IntType:
		return "int"
	case Int8Type:
		return "int8"
	case Int16Type:
		return "int16"
	case Int32Type:
		return "int32"
	case Int64Type:
		return "int64"
	case StringType:
		return "string"
	case TimeType:
		return "time"
	case TimeFullType:
		return "time_full"
	case UintType:
		return "uint"
	case Uint8Type:
		return "uint8"
	case Uint16Type:
		return "uint16"
	case Uint32Type:
		return "uint32"
	case Uint64Type:
		return "uint64"
	case ReflectType:
		return "reflect"
	case NamespaceType:
		return "namespace"
	case StringerType:
		return "stringer"
	case ErrorType:
		return "error"
	default:
		return fmt.Sprintf("Type(%d)", t)
	}
}
