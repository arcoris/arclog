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

type namedHook struct {
	name string
}

var _ hook.Named = namedHook{}

func (h namedHook) Name() string {
	return h.name
}

func TestNamedContract(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hook hook.Named
		want string
	}{
		{name: "non-empty", hook: namedHook{name: "audit"}, want: "audit"},
		{name: "empty allowed", hook: namedHook{}, want: ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.hook.Name(); got != tt.want {
				t.Fatalf("Name() = %q, want %q", got, tt.want)
			}
		})
	}
}
