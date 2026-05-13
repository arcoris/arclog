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

import "math"

// Float32 constructs a float32 field preserving the IEEE 754 bit pattern.
func Float32(key string, value float32) Field {
	return Field{Key: key, Type: Float32Type, Integer: int64(math.Float32bits(value))}
}

// Float64 constructs a float64 field preserving the IEEE 754 bit pattern.
func Float64(key string, value float64) Field {
	return Field{Key: key, Type: Float64Type, Integer: int64(math.Float64bits(value))}
}
