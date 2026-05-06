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

package level_test

import (
	"testing"

	"arcoris.dev/arclog/api/level"
)

var _ level.Enabler = level.EnablerFunc(nil)

// TestEnablerFunc verifies the function adapter without relying on package
// internals.
func TestEnablerFunc(t *testing.T) {
	t.Parallel()

	enabler := level.EnablerFunc(func(lvl level.Level) bool {
		return lvl.Enabled(level.Warn)
	})

	if enabler.Enabled(level.Info) {
		t.Fatalf("Info should be disabled at Warn threshold")
	}
	if !enabler.Enabled(level.Error) {
		t.Fatalf("Error should be enabled at Warn threshold")
	}
}

// TestEnablerFuncNilPanics documents that a nil function adapter is invalid.
func TestEnablerFuncNilPanics(t *testing.T) {
	t.Parallel()

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatalf("nil EnablerFunc did not panic")
		}
	}()

	var enabler level.EnablerFunc
	_ = enabler.Enabled(level.Info)
}
