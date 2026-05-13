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

func TestZeroValueFieldIsSkip(t *testing.T) {
	t.Parallel()

	var f Field
	if !f.IsSkip() {
		t.Fatal("zero value field must be skip")
	}
	if !Skip().Equal(f) {
		t.Fatal("Skip() must equal the zero field")
	}
}

func TestFieldIsNull(t *testing.T) {
	t.Parallel()

	if !Null("key").IsNull() {
		t.Fatal("Null(key) must report IsNull")
	}
	if Skip().IsNull() {
		t.Fatal("Skip() must not report IsNull")
	}
}
