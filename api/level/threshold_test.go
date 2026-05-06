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

package level

import (
	"testing"
)

var _ Threshold = (*testThreshold)(nil)

type testThreshold struct {
	lvl Level
}

func (t *testThreshold) Enabled(lvl Level) bool {
	return lvl.Enabled(t.lvl)
}

func (t *testThreshold) Level() Level {
	return t.lvl
}

func (t *testThreshold) SetLevel(lvl Level) {
	t.lvl = lvl
}

// TestThresholdContract verifies the minimal mutable-threshold contract without
// importing any runtime implementation.
func TestThresholdContract(t *testing.T) {
	t.Parallel()

	threshold := &testThreshold{lvl: Info}
	if !threshold.Enabled(Error) {
		t.Fatalf("Error should be enabled at Info threshold")
	}
	if threshold.Enabled(Debug) {
		t.Fatalf("Debug should be disabled at Info threshold")
	}

	threshold.SetLevel(Warn)
	if threshold.Level() != Warn {
		t.Fatalf("Level() = %v, want %v", threshold.Level(), Warn)
	}
	if threshold.Enabled(Info) {
		t.Fatalf("Info should be disabled at Warn threshold")
	}
}
