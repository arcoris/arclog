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

// Float32Ptr constructs a field from *float32.
func Float32Ptr(key string, value *float32) Field {
	if value == nil {
		return Null(key)
	}
	return Float32(key, *value)
}

// Float64Ptr constructs a field from *float64.
func Float64Ptr(key string, value *float64) Field {
	if value == nil {
		return Null(key)
	}
	return Float64(key, *value)
}
