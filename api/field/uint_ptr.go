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

// UintPtr constructs a field from *uint.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Uint.
func UintPtr(key string, value *uint) Field {
	if value == nil {
		return Null(key)
	}
	return Uint(key, *value)
}

// Uint8Ptr constructs a field from *uint8.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Uint8.
func Uint8Ptr(key string, value *uint8) Field {
	if value == nil {
		return Null(key)
	}
	return Uint8(key, *value)
}

// Uint16Ptr constructs a field from *uint16.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Uint16.
func Uint16Ptr(key string, value *uint16) Field {
	if value == nil {
		return Null(key)
	}
	return Uint16(key, *value)
}

// Uint32Ptr constructs a field from *uint32.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Uint32.
func Uint32Ptr(key string, value *uint32) Field {
	if value == nil {
		return Null(key)
	}
	return Uint32(key, *value)
}

// Uint64Ptr constructs a field from *uint64.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Uint64.
func Uint64Ptr(key string, value *uint64) Field {
	if value == nil {
		return Null(key)
	}
	return Uint64(key, *value)
}
