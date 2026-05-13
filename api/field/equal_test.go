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
	"errors"
	"testing"
	"time"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	bytesA := []byte("same")
	bytesB := []byte("same")
	timeA := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)
	timeB := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)
	errA := errors.New("same")
	errB := errors.New("same")

	tests := []struct {
		name string
		a    Field
		b    Field
		want bool
	}{
		{name: "same primitive", a: String("k", "v"), b: String("k", "v"), want: true},
		{name: "different key", a: String("a", "v"), b: String("b", "v"), want: false},
		{name: "different type", a: String("k", "v"), b: Bytes("k", []byte("v")), want: false},
		{name: "different value", a: String("k", "a"), b: String("k", "b"), want: false},
		{name: "bytes content", a: Bytes("k", bytesA), b: Bytes("k", bytesB), want: true},
		{name: "bytes different content", a: Bytes("k", []byte("same")), b: Bytes("k", []byte("diff")), want: false},
		{name: "full time equal", a: Time("k", timeA), b: Time("k", timeB), want: true},
		{name: "time full different", a: Time("k", timeA), b: Time("k", timeA.Add(time.Second)), want: false},
		{name: "reflect equal", a: Reflect("k", struct{ A int }{A: 1}), b: Reflect("k", struct{ A int }{A: 1}), want: true},
		{name: "reflect different", a: Reflect("k", struct{ A int }{A: 1}), b: Reflect("k", struct{ A int }{A: 2}), want: false},
		{name: "error equal", a: NamedError("k", errA), b: NamedError("k", errB), want: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.a.Equal(tt.b); got != tt.want {
				t.Fatalf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
