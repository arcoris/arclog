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

	apilevel "arcoris.dev/arclog/api/level"
)

func TestStaticThresholdBehavior(t *testing.T) {
	t.Parallel()

	threshold := NewStaticThreshold(apilevel.Warn)
	if got := threshold.Level(); got != apilevel.Warn {
		t.Fatalf("Level() = %v, want %v", got, apilevel.Warn)
	}
	if !threshold.Enabled(apilevel.Error) {
		t.Fatal("Enabled(Error) = false, want true")
	}
	if threshold.Enabled(apilevel.Info) {
		t.Fatal("Enabled(Info) = true, want false")
	}
}

func TestStaticThresholdAcceptsOff(t *testing.T) {
	t.Parallel()

	threshold := NewStaticThreshold(apilevel.Off)
	if threshold.Enabled(apilevel.Fatal) {
		t.Fatal("Off static threshold enabled Fatal, want false")
	}
}

func TestStaticThresholdPanicsOnInvalid(t *testing.T) {
	t.Parallel()

	assertPanic(t, func() {
		_ = NewStaticThreshold(apilevel.Level(16))
	})
}

func TestStaticThresholdSetLevelIsImmutable(t *testing.T) {
	t.Parallel()

	threshold := NewStaticThreshold(apilevel.Warn)
	threshold.SetLevel(apilevel.Error)

	if got := threshold.Level(); got != apilevel.Warn {
		t.Fatalf("Level after SetLevel(Error) = %v, want %v", got, apilevel.Warn)
	}
}

func TestStaticThresholdSetLevelPanicsOnInvalid(t *testing.T) {
	t.Parallel()

	threshold := NewStaticThreshold(apilevel.Warn)
	assertPanic(t, func() {
		threshold.SetLevel(apilevel.Level(16))
	})
}
