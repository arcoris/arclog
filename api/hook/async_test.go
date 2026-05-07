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

type asyncDeclaration bool

func (a asyncDeclaration) Async() bool {
	return bool(a)
}

func TestAllowsAsync(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hook any
		want bool
	}{
		{name: "declares async", hook: asyncDeclaration(true), want: true},
		{name: "declares sync", hook: asyncDeclaration(false), want: false},
		{name: "no async contract", hook: struct{}{}, want: false},
		{name: "nil", hook: nil, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := hook.AllowsAsync(tt.hook); got != tt.want {
				t.Fatalf("AllowsAsync() = %v, want %v", got, tt.want)
			}
		})
	}
}
