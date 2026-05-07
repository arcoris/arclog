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

package hook_test

import (
	"testing"

	"arcoris.dev/arclog/api/hook"
)

func TestPriorityBefore(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		left  hook.Priority
		right hook.Priority
		want  bool
	}{
		{name: "high before default", left: hook.PriorityHigh, right: hook.PriorityDefault, want: true},
		{name: "default before high", left: hook.PriorityDefault, right: hook.PriorityHigh, want: false},
		{name: "same priority", left: hook.PriorityDefault, right: hook.PriorityDefault, want: false},
		{name: "last after low", left: hook.PriorityLast, right: hook.PriorityLow, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.left.Before(tt.right); got != tt.want {
				t.Fatalf("Before() = %v, want %v", got, tt.want)
			}
		})
	}
}
