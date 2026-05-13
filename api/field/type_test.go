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

import "testing"

func TestTypeString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		typ  Type
		want string
	}{
		{name: "skip", typ: SkipType, want: "skip"},
		{name: "null", typ: NullType, want: "null"},
		{name: "bool", typ: BoolType, want: "bool"},
		{name: "bytes", typ: BytesType, want: "bytes"},
		{name: "duration", typ: DurationType, want: "duration"},
		{name: "float64", typ: Float64Type, want: "float64"},
		{name: "float32", typ: Float32Type, want: "float32"},
		{name: "int", typ: IntType, want: "int"},
		{name: "int8", typ: Int8Type, want: "int8"},
		{name: "int16", typ: Int16Type, want: "int16"},
		{name: "int32", typ: Int32Type, want: "int32"},
		{name: "int64", typ: Int64Type, want: "int64"},
		{name: "string", typ: StringType, want: "string"},
		{name: "time", typ: TimeType, want: "time"},
		{name: "time full", typ: TimeFullType, want: "time_full"},
		{name: "uint", typ: UintType, want: "uint"},
		{name: "uint8", typ: Uint8Type, want: "uint8"},
		{name: "uint16", typ: Uint16Type, want: "uint16"},
		{name: "uint32", typ: Uint32Type, want: "uint32"},
		{name: "uint64", typ: Uint64Type, want: "uint64"},
		{name: "reflect", typ: ReflectType, want: "reflect"},
		{name: "namespace", typ: NamespaceType, want: "namespace"},
		{name: "stringer", typ: StringerType, want: "stringer"},
		{name: "error", typ: ErrorType, want: "error"},
		{name: "unknown", typ: Type(255), want: "Type(255)"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.typ.String(); got != tt.want {
				t.Fatalf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
