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

package entrykey

// Key identifies a canonical field name used for arclog log-entry metadata.
//
// Key is a small string type that prevents raw string literals from spreading
// across API and runtime packages. It does not imply that user-defined fields
// must use this package.
type Key string

// String returns k as a string.
func (k Key) String() string {
	return string(k)
}

// IsZero reports whether k is empty.
//
// The empty key is not canonical. Runtime encoder configuration may use an
// empty key as an omission marker, but omission policy belongs to that
// configuration layer rather than to this package.
func (k Key) IsZero() bool {
	return k == ""
}

// IsKnown reports whether k is one of the canonical keys defined by this
// package.
func (k Key) IsKnown() bool {
	switch k {
	case Time,
		Level,
		Logger,
		Message,
		Caller,
		Function,
		Stacktrace,
		Error:
		return true
	default:
		return false
	}
}
