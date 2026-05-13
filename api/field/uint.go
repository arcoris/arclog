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

// Uint constructs a uint field.
func Uint(key string, value uint) Field {
	return Field{Key: key, Type: UintType, Integer: int64(value)}
}

// Uint8 constructs a uint8 field.
func Uint8(key string, value uint8) Field {
	return Field{Key: key, Type: Uint8Type, Integer: int64(value)}
}

// Uint16 constructs a uint16 field.
func Uint16(key string, value uint16) Field {
	return Field{Key: key, Type: Uint16Type, Integer: int64(value)}
}

// Uint32 constructs a uint32 field.
func Uint32(key string, value uint32) Field {
	return Field{Key: key, Type: Uint32Type, Integer: int64(value)}
}

// Uint64 constructs a uint64 field.
func Uint64(key string, value uint64) Field {
	return Field{Key: key, Type: Uint64Type, Integer: int64(value)}
}
