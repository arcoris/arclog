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

package clock_test

import (
	"testing"
	"time"

	"arcoris.dev/arclog/api/clock"
)

type fixedClock struct {
	value time.Time
}

func (c fixedClock) Now() time.Time {
	return c.value
}

var _ clock.Clock = fixedClock{}

func TestClockContractReturnsImplementationTimestamp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want time.Time
	}{
		{
			name: "utc timestamp",
			want: time.Date(2026, 5, 8, 12, 0, 0, 123, time.UTC),
		},
		{
			name: "zero timestamp",
			want: time.Time{},
		},
		{
			name: "local location",
			want: time.Date(2026, 5, 8, 12, 0, 0, 123, time.FixedZone("clock-test", 3*60*60)),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var c clock.Clock = fixedClock{value: tt.want}
			if got := c.Now(); got != tt.want {
				t.Fatalf("Now() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
