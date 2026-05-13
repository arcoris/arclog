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

// Error constructs an error field using the default key "error".
//
// Nil and typed-nil errors produce Skip because there is no error value to
// attach. Non-nil errors are stored as ErrorType in Field.Interface. The
// constructor does not call Error; string conversion belongs to encoders.
func Error(err error) Field {
	return NamedError("error", err)
}

// NamedError constructs an error field using key.
//
// Nil and typed-nil errors produce Skip. The key is stored exactly as provided
// for non-nil errors.
func NamedError(key string, err error) Field {
	if isNil(err) {
		return Skip()
	}
	return Field{Key: key, Type: ErrorType, Interface: err}
}
