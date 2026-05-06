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

package buffer_test

import (
	"strconv"
	"testing"
	"time"
)

func TestBufferAppendTime(t *testing.T) {
	buf := newBuffer()
	tm := time.Date(2026, 5, 6, 7, 8, 9, 123_000_000, time.FixedZone("TEST", 3*60*60))
	layout := "2006-01-02T15:04:05.000Z07:00"

	buf.AppendTime(tm, layout)

	if got, want := string(buf.Bytes()), tm.Format(layout); got != want {
		t.Fatalf("AppendTime = %q, want %q", got, want)
	}
}

func TestBufferAppendDuration(t *testing.T) {
	buf := newBuffer()
	duration := 2*time.Second + 500*time.Millisecond

	buf.AppendDuration(duration)

	if got, want := string(buf.Bytes()), strconv.FormatInt(int64(duration), 10); got != want {
		t.Fatalf("AppendDuration = %q, want %q", got, want)
	}
}
