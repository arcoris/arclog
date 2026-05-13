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

import "time"

// Duration constructs a duration field using nanoseconds.
//
// The raw nanosecond count is stored in Field.Integer. Formatting policy, such
// as rendering as a number, string, or structured duration, belongs to runtime
// encoders.
func Duration(key string, value time.Duration) Field {
	return Field{Key: key, Type: DurationType, Integer: int64(value)}
}
