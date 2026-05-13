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

func TestReflect(t *testing.T) {
	t.Parallel()

	if got := Reflect("key", nil); !got.Equal(Null("key")) {
		t.Fatalf("Reflect(nil) = %#v", got)
	}

	var ptr *int
	if got := Reflect("key", ptr); !got.Equal(Null("key")) {
		t.Fatalf("Reflect(typed nil ptr) = %#v", got)
	}

	var slice []string
	if got := Reflect("key", slice); !got.Equal(Null("key")) {
		t.Fatalf("Reflect(typed nil slice) = %#v", got)
	}

	value := struct{ Name string }{Name: "arc"}
	got := Reflect("key", value)
	if got.Key != "key" || got.Type != ReflectType || got.Interface != value {
		t.Fatalf("Reflect(value) = %#v", got)
	}
}
