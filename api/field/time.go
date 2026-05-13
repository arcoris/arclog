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

const (
	minUnixNano = -1 << 63
	maxUnixNano = 1<<63 - 1
)

var (
	minTime = time.Unix(0, minUnixNano)
	maxTime = time.Unix(0, maxUnixNano)
)

// Time constructs a time field descriptor.
//
// Values representable as Unix nanoseconds use the compact form and preserve
// the original location separately. Values outside that range retain the full
// time.Time value.
func Time(key string, value time.Time) Field {
	if value.Before(minTime) || value.After(maxTime) {
		return Field{Key: key, Type: TimeFullType, Interface: value}
	}
	return Field{Key: key, Type: TimeType, Integer: value.UnixNano(), Interface: value.Location()}
}
