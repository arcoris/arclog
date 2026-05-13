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

// Type identifies which storage slot of Field is meaningful.
type Type uint8

const (
	SkipType Type = iota
	NullType

	BoolType
	BytesType

	DurationType

	Float64Type
	Float32Type

	IntType
	Int8Type
	Int16Type
	Int32Type
	Int64Type

	StringType

	TimeType
	TimeFullType

	UintType
	Uint8Type
	Uint16Type
	Uint32Type
	Uint64Type

	ReflectType
	NamespaceType
	StringerType
	ErrorType
	AnyType
)

// String returns a stable diagnostic name for t.
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
	case AnyType:
		return "any"
	default:
		return fmt.Sprintf("Type(%d)", t)
	}
}
