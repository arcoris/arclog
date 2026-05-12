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

package exit_test

import (
	"testing"

	"arcoris.dev/arclog/runtime/internal/exit"
)

func TestNilRecorderMethods(t *testing.T) {
	var recorder *exit.Recorder

	recorder.Unstub()

	if recorder.Exited() {
		t.Fatal("nil Recorder Exited() = true, want false")
	}
	if got := recorder.Code(); got != 0 {
		t.Fatalf("nil Recorder Code() = %d, want 0", got)
	}
	if got := recorder.Calls(); got != 0 {
		t.Fatalf("nil Recorder Calls() = %d, want 0", got)
	}
	if recorder.Restored() {
		t.Fatal("nil Recorder Restored() = true, want false")
	}
}
