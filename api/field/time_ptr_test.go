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

func TestTimePtr(t *testing.T) {
	t.Parallel()

	value := time.Date(2026, 5, 13, 10, 11, 12, 13, time.UTC)
	tests := []struct {
		name string
		got  Field
		want Field
	}{
		{name: "nil", got: TimePtr("ts", nil), want: Null("ts")},
		{name: "value", got: TimePtr("ts", &value), want: Time("ts", value)},
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
