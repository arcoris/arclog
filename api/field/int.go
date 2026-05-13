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

// Int constructs an int field.
//
// The value is widened to int64 and stored in Field.Integer.
func Int(key string, value int) Field {
	return Field{Key: key, Type: IntType, Integer: int64(value)}
}

// Int8 constructs an int8 field.
//
// The value is widened to int64 and stored in Field.Integer.
func Int8(key string, value int8) Field {
	return Field{Key: key, Type: Int8Type, Integer: int64(value)}
}

// Int16 constructs an int16 field.
//
// The value is widened to int64 and stored in Field.Integer.
func Int16(key string, value int16) Field {
	return Field{Key: key, Type: Int16Type, Integer: int64(value)}
}

// Int32 constructs an int32 field.
//
// The value is widened to int64 and stored in Field.Integer.
func Int32(key string, value int32) Field {
	return Field{Key: key, Type: Int32Type, Integer: int64(value)}
}

// Int64 constructs an int64 field.
//
// The value is stored directly in Field.Integer.
func Int64(key string, value int64) Field {
	return Field{Key: key, Type: Int64Type, Integer: value}
}
