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

func TestDuration(t *testing.T) {
	t.Parallel()

	tests := []time.Duration{0, -time.Second, 3 * time.Second}
	for _, value := range tests {
		got := Duration("dur", value)
		want := Field{Key: "dur", Type: DurationType, Integer: int64(value)}
		if !got.Equal(want) {
			t.Fatalf("Duration(%v) = %#v, want %#v", value, got, want)
		}
	}
}
