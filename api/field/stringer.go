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

// Stringer constructs a field backed by fmt.Stringer.
//
// The constructor does not call String. Runtime encoders decide when and how
// to materialize the string representation.
// Nil and typed-nil stringers become Null(key).
func Stringer(key string, value fmt.Stringer) Field {
	if isNil(value) {
		return Null(key)
	}
	return Field{Key: key, Type: StringerType, Interface: value}
}
