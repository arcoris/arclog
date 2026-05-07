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
	"errors"
	"testing"

	"arcoris.dev/arclog/api/hook"
)

func TestWriteResult(t *testing.T) {
	t.Parallel()

	if hook.Success().Failed() {
		t.Fatal("Success().Failed() = true")
	}

	wantErr := errors.New("write failed")
	result := hook.Failure(wantErr)
	if !result.Failed() {
		t.Fatal("Failure(err).Failed() = false")
	}
	if !errors.Is(result.Err, wantErr) {
		t.Fatalf("Err = %v, want %v", result.Err, wantErr)
	}
}
