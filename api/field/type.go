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

// Type identifies which storage slot of Field is meaningful and how the field
// should be dispatched to an encoder.
//
// The zero value is SkipType, making the zero Field a no-op. This is a deliberate
// API choice: invalid or optional constructors can return Field{} without
// causing a panic in the logging path.
type Type uint8

const (
	// SkipType marks a field as a no-op.
	SkipType Type = iota
	// BoolType stores a bool in Field.Integer as 0 or 1.
	BoolType
	// ByteStringType stores byte-oriented text in Field.Interface as a []byte.
	ByteStringType
	// Complex128Type stores a complex128 in Field.Interface.
	Complex128Type
	// Complex64Type stores a complex64 in Field.Interface.
	Complex64Type
	// DurationType stores a time.Duration in Field.Integer as nanoseconds.
	DurationType
	// Float64Type stores a float64 in Field.Integer using math.Float64bits.
	Float64Type
	// Float32Type stores a float32 in Field.Integer using math.Float32bits.
	Float32Type
	// Int64Type stores an int64 in Field.Integer.
	Int64Type
	// Int32Type stores an int32 in Field.Integer.
	Int32Type
	// Int16Type stores an int16 in Field.Integer.
	Int16Type
	// Int8Type stores an int8 in Field.Integer.
	Int8Type
	// IntType stores an int in Field.Integer.
	IntType
	// StringType stores a string in Field.String.
	StringType
	// TimeType stores a time.Time as Unix nanoseconds in Field.Integer and a
	// *time.Location in Field.Interface.
	TimeType
	// TimeFullType stores a time.Time directly in Field.Interface for values that
	// cannot be represented as Unix nanoseconds.
	TimeFullType
	// Uint64Type stores a uint64 bit pattern in Field.Integer.
	Uint64Type
	// Uint32Type stores a uint32 in Field.Integer.
	Uint32Type
	// Uint16Type stores a uint16 in Field.Integer.
	Uint16Type
	// Uint8Type stores a uint8 in Field.Integer.
	Uint8Type
	// UintType stores a uint bit pattern in Field.Integer.
	UintType
	// UintptrType stores a uintptr bit pattern in Field.Integer.
	UintptrType
	// ObjectMarshalerType stores an encoder.ObjectMarshaler in Field.Interface.
	ObjectMarshalerType
	// ArrayMarshalerType stores an encoder.ArrayMarshaler in Field.Interface.
	ArrayMarshalerType
	// InlineMarshalerType stores an encoder.ObjectMarshaler whose fields should be
	// appended into the current namespace instead of under Field.Key.
	InlineMarshalerType
	// ReflectType stores an arbitrary value in Field.Interface for the encoder's
	// reflection path.
	ReflectType
	// NamespaceType opens or selects a namespace named by Field.Key.
	NamespaceType
	// StringerType stores a fmt.Stringer in Field.Interface.
	StringerType
	// ErrorType stores an error in Field.Interface.
	ErrorType
)

// String returns a stable diagnostic name for t.
func (t Type) String() string {
	switch t {
	case SkipType:
		return "SkipType"
	case BoolType:
		return "BoolType"
	case ByteStringType:
		return "ByteStringType"
	case Complex128Type:
		return "Complex128Type"
	case Complex64Type:
		return "Complex64Type"
	case DurationType:
		return "DurationType"
	case Float64Type:
		return "Float64Type"
	case Float32Type:
		return "Float32Type"
	case Int64Type:
		return "Int64Type"
	case Int32Type:
		return "Int32Type"
	case Int16Type:
		return "Int16Type"
	case Int8Type:
		return "Int8Type"
	case IntType:
		return "IntType"
	case StringType:
		return "StringType"
	case TimeType:
		return "TimeType"
	case TimeFullType:
		return "TimeFullType"
	case Uint64Type:
		return "Uint64Type"
	case Uint32Type:
		return "Uint32Type"
	case Uint16Type:
		return "Uint16Type"
	case Uint8Type:
		return "Uint8Type"
	case UintType:
		return "UintType"
	case UintptrType:
		return "UintptrType"
	case ObjectMarshalerType:
		return "ObjectMarshalerType"
	case ArrayMarshalerType:
		return "ArrayMarshalerType"
	case InlineMarshalerType:
		return "InlineMarshalerType"
	case ReflectType:
		return "ReflectType"
	case NamespaceType:
		return "NamespaceType"
	case StringerType:
		return "StringerType"
	case ErrorType:
		return "ErrorType"
	default:
		return fmt.Sprintf("Type(%d)", t)
	}
}
