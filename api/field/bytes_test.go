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

func TestBytes(t *testing.T) {
	t.Parallel()

	var nilBytes []byte
	gotNil := Bytes("data", nilBytes)
	if gotNil.Type != BytesType || gotNil.Bytes != nil || gotNil.Interface != nil {
		t.Fatalf("Bytes(nil) = %#v", gotNil)
	}

	payload := []byte("abc")
	got := Bytes("data", payload)
	if got.Type != BytesType || got.Interface != nil {
		t.Fatalf("Bytes() = %#v", got)
	}
	if len(got.Bytes) != len(payload) || &got.Bytes[0] != &payload[0] {
		t.Fatal("Bytes() must retain the borrowed slice")
	}

	payload[0] = 'z'
	if got.Bytes[0] != 'z' {
		t.Fatal("Bytes() must expose borrowed slice semantics")
	}
}
