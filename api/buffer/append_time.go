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

package buffer

import (
	"strconv"
	"time"
)

// AppendTime appends t formatted with layout.
//
// The layout follows Go's time formatting rules. The method delegates to
// time.Time.AppendFormat so formatting appends directly into the buffer's byte
// slice and avoids an intermediate string allocation.
func (b *Buffer) AppendTime(t time.Time, layout string) {
	b.data = t.AppendFormat(b.data, layout)
}

// AppendDuration appends d as a base-10 number of nanoseconds.
//
// Higher-level encoders that need seconds, milliseconds, strings, or floating
// point durations SHOULD convert d to the target representation before using a
// numeric append helper.
func (b *Buffer) AppendDuration(d time.Duration) {
	b.data = strconv.AppendInt(b.data, int64(d), 10)
}
