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

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Parallel()

	zero := time.Time{}
	gotZero := Time("ts", zero)
	if gotZero.Type != TimeFullType {
		t.Fatalf("zero time type = %v, want %v", gotZero.Type, TimeFullType)
	}
	storedZero, ok := gotZero.Interface.(time.Time)
	if !ok || !storedZero.Equal(zero) {
		t.Fatalf("zero time value = %#v", gotZero.Interface)
	}

	utc := time.Date(2026, 5, 13, 10, 11, 12, 13, time.UTC)
	gotUTC := Time("ts", utc)
	if gotUTC.Type != TimeType || gotUTC.Integer != utc.UnixNano() || gotUTC.Interface != time.UTC {
		t.Fatalf("UTC time = %#v", gotUTC)
	}

	loc := time.FixedZone("UTC+3", 3*60*60)
	normal := time.Date(2026, 5, 13, 10, 11, 12, 13, loc)
	gotNormal := Time("ts", normal)
	if gotNormal.Type != TimeType || gotNormal.Integer != normal.UnixNano() || gotNormal.Interface != loc {
		t.Fatalf("localized time = %#v", gotNormal)
	}

	for _, value := range []time.Time{minTime, maxTime} {
		got := Time("ts", value)
		if got.Type != TimeType || got.Integer != value.UnixNano() {
			t.Fatalf("compact boundary %v = %#v", value, got)
		}
	}

	outOfRange := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)
	gotFull := Time("ts", outOfRange)
	if gotFull.Type != TimeFullType {
		t.Fatalf("out-of-range time type = %v, want %v", gotFull.Type, TimeFullType)
	}
	stored, ok := gotFull.Interface.(time.Time)
	if !ok || !stored.Equal(outOfRange) {
		t.Fatalf("out-of-range time = %#v", gotFull.Interface)
	}
}
