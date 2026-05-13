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

func TestDurationPtr(t *testing.T) {
	t.Parallel()

	value := time.Second
	tests := []struct {
		name string
		got  Field
		want Field
	}{
		{name: "nil", got: DurationPtr("dur", nil), want: Null("dur")},
		{name: "value", got: DurationPtr("dur", &value), want: Duration("dur", value)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if !tt.got.Equal(tt.want) {
				t.Fatalf("got %#v, want %#v", tt.got, tt.want)
			}
		})
	}
}
