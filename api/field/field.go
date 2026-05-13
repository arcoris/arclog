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
	Key       string
	Type      Type
	Integer   int64
	String    string
	Bytes     []byte
	Interface any
}

// IsSkip reports whether f is a no-op field.
func (f Field) IsSkip() bool { return f.Type == SkipType }

// IsNull reports whether f is an explicit null field.
func (f Field) IsNull() bool { return f.Type == NullType }
