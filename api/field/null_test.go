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

import "testing"

func TestNull(t *testing.T) {
	t.Parallel()

	got := Null("key")
	want := Field{Key: "key", Type: NullType}
	if !got.Equal(want) {
		t.Fatalf("Null() = %#v, want %#v", got, want)
	}
	if !got.IsNull() {
		t.Fatal("Null() must report IsNull")
	}
	if got.IsSkip() {
		t.Fatal("Null() must not report IsSkip")
	}
}
