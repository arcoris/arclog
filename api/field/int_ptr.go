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

// IntPtr constructs a field from *int.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Int.
func IntPtr(key string, value *int) Field {
	if value == nil {
		return Null(key)
	}
	return Int(key, *value)
}

// Int8Ptr constructs a field from *int8.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Int8.
func Int8Ptr(key string, value *int8) Field {
	if value == nil {
		return Null(key)
	}
	return Int8(key, *value)
}

// Int16Ptr constructs a field from *int16.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Int16.
func Int16Ptr(key string, value *int16) Field {
	if value == nil {
		return Null(key)
	}
	return Int16(key, *value)
}

// Int32Ptr constructs a field from *int32.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Int32.
func Int32Ptr(key string, value *int32) Field {
	if value == nil {
		return Null(key)
	}
	return Int32(key, *value)
}

// Int64Ptr constructs a field from *int64.
//
// A nil pointer becomes Null(key). A non-nil pointer is dereferenced
// immediately and stored with Int64.
func Int64Ptr(key string, value *int64) Field {
	if value == nil {
		return Null(key)
	}
	return Int64(key, *value)
}
