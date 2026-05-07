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

// Field is a compact, tagged representation of one structured logging field.
//
// The meaning of Integer, String, and Interface depends on Type. Primitive
// constructors avoid Interface when practical so fields can be created cheaply
// on disabled log paths. More complex values, such as object marshalers,
// arrays, errors, stringers, byte slices, and reflected values, are retained in
// Interface and encoded later.
//
// The zero value is a skip field. It is safe to pass through AddTo and through
// field slices.
type Field struct {
	// Key is the field name used by keyed object encoders.
	Key string
	// Type selects the active storage slot and encoder dispatch path.
	Type Type
	// Integer stores compact numeric, boolean, duration, time, and floating-point
	// bit-pattern payloads.
	Integer int64
	// String stores string payloads.
	String string
	// Interface stores payloads that do not fit in Integer or String.
	Interface any
}

// IsSkip reports whether f is a no-op field.
func (f Field) IsSkip() bool { return f.Type == SkipType }
