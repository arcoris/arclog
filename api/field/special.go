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
	"reflect"
	"time"
)

const (
	minUnixNano = -1 << 63
	maxUnixNano = 1<<63 - 1
)

var (
	minTime = time.Unix(0, minUnixNano)
	maxTime = time.Unix(0, maxUnixNano)
)

// Skip returns the zero field descriptor.
func Skip() Field {
	return Field{}
}

// Null constructs an explicit null field descriptor.
func Null(key string) Field {
	return Field{Key: key, Type: NullType}
}

// Bytes constructs a bytes field descriptor.
//
// Bytes stores a borrowed slice. Callers must not mutate value until the log
// call or context binding that consumes the field has finished observing it.
func Bytes(key string, value []byte) Field {
	return Field{Key: key, Type: BytesType, Interface: value}
}

// Time constructs a time field descriptor.
//
// Values representable as Unix nanoseconds use the compact form and preserve
// the original location separately. Values outside that range retain the full
// time.Time value.
func Time(key string, value time.Time) Field {
	if value.Before(minTime) || value.After(maxTime) {
		return Field{Key: key, Type: TimeFullType, Interface: value}
	}
	return Field{Key: key, Type: TimeType, Integer: value.UnixNano(), Interface: value.Location()}
}

// Namespace constructs a namespace descriptor under key.
//
// Namespace is descriptor-only. Runtime encoders decide how namespace scopes
// are opened and closed in their own output formats.
func Namespace(key string) Field {
	return Field{Key: key, Type: NamespaceType}
}

// Error constructs an error field using the default key "error".
func Error(err error) Field {
	return NamedError("error", err)
}

// NamedError constructs an error field using key.
func NamedError(key string, err error) Field {
	if isNil(err) {
		return Skip()
	}
	return Field{Key: key, Type: ErrorType, Interface: err}
}

// Stringer constructs a field backed by fmt.Stringer.
//
// The constructor does not call String. Runtime encoders decide when and how
// to materialize the string representation.
func Stringer(key string, value fmt.Stringer) Field {
	if isNil(value) {
		return Null(key)
	}
	return Field{Key: key, Type: StringerType, Interface: value}
}

// Reflect constructs a reflection-backed field descriptor.
//
// Reflect stores the provided value without encoding it. Nil and typed-nil
// values become explicit null fields.
func Reflect(key string, value any) Field {
	if isNil(value) {
		return Null(key)
	}
	return Field{Key: key, Type: ReflectType, Interface: value}
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}
