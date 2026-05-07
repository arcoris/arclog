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

package field_test

import (
	"arcoris.dev/arclog/api/field"
	"testing"
)

func TestEqualByteStringComparesContents(t *testing.T) {
	t.Parallel()
	a := field.ByteString("b", []byte("abc"))
	b := field.ByteString("b", []byte("abc"))
	if !a.Equal(b) {
		t.Fatal("not equal")
	}
}
func TestEqualDetectsDifferentType(t *testing.T) {
	t.Parallel()
	if field.String("x", "1").Equal(field.Int("x", 1)) {
		t.Fatal("different types equal")
	}
}
