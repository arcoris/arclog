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
	"sync"
	"testing"

	apilevel "arcoris.dev/arclog/api/level"
)

func TestAtomicLevelZeroValueIsInfo(t *testing.T) {
	t.Parallel()

	var atomicLevel AtomicLevel
	if got := atomicLevel.Level(); got != apilevel.Info {
		t.Fatalf("zero AtomicLevel.Level() = %v, want %v", got, apilevel.Info)
	}
}

func TestNewAtomicLevelIsInfo(t *testing.T) {
	t.Parallel()

	if got := NewAtomicLevel().Level(); got != apilevel.Info {
		t.Fatalf("NewAtomicLevel().Level() = %v, want %v", got, apilevel.Info)
	}
}

func TestNewAtomicLevelAtAcceptsSeverityThresholds(t *testing.T) {
	t.Parallel()

	if got := NewAtomicLevelAt(apilevel.Warn).Level(); got != apilevel.Warn {
		t.Fatalf("NewAtomicLevelAt(Warn).Level() = %v, want %v", got, apilevel.Warn)
	}
}

func TestNewAtomicLevelAtAcceptsOff(t *testing.T) {
	t.Parallel()

	if got := NewAtomicLevelAt(apilevel.Off).Level(); got != apilevel.Off {
		t.Fatalf("NewAtomicLevelAt(Off).Level() = %v, want %v", got, apilevel.Off)
	}
}

func TestNewAtomicLevelAtPanicsOnInvalidThreshold(t *testing.T) {
	t.Parallel()

	assertPanic(t, func() {
		_ = NewAtomicLevelAt(apilevel.Level(16))
	})
}

func TestParseAtomicLevel(t *testing.T) {
	t.Parallel()

	atomicLevel, err := ParseAtomicLevel("warning")
	if err != nil {
		t.Fatalf("ParseAtomicLevel returned error: %v", err)
	}
	if got := atomicLevel.Level(); got != apilevel.Warn {
		t.Fatalf("ParseAtomicLevel(warning).Level() = %v, want %v", got, apilevel.Warn)
	}
}

func TestAtomicLevelEnabled(t *testing.T) {
	t.Parallel()

	atomicLevel := NewAtomicLevelAt(apilevel.Warn)
	if !atomicLevel.Enabled(apilevel.Error) {
		t.Fatal("Enabled(Error) = false, want true")
	}
	if atomicLevel.Enabled(apilevel.Info) {
		t.Fatal("Enabled(Info) = true, want false")
	}
}

func TestNilAtomicLevelDisablesAllLevels(t *testing.T) {
	t.Parallel()

	var atomicLevel *AtomicLevel
	if atomicLevel.Enabled(apilevel.Fatal) {
		t.Fatal("nil AtomicLevel.Enabled(Fatal) = true, want false")
	}
	if got := atomicLevel.Level(); got != apilevel.Off {
		t.Fatalf("nil AtomicLevel.Level() = %v, want %v", got, apilevel.Off)
	}
}

func TestAtomicLevelSetLevel(t *testing.T) {
	t.Parallel()

	atomicLevel := NewAtomicLevel()
	atomicLevel.SetLevel(apilevel.Off)

	if got := atomicLevel.Level(); got != apilevel.Off {
		t.Fatalf("Level after SetLevel(Off) = %v, want %v", got, apilevel.Off)
	}
}

func TestAtomicLevelSetLevelPanicsOnInvalid(t *testing.T) {
	t.Parallel()

	atomicLevel := NewAtomicLevel()
	assertPanic(t, func() {
		atomicLevel.SetLevel(apilevel.Level(16))
	})
}

func TestAtomicLevelTrySetLevel(t *testing.T) {
	t.Parallel()

	atomicLevel := NewAtomicLevel()
	if !atomicLevel.TrySetLevel(apilevel.Error) {
		t.Fatal("TrySetLevel(Error) = false, want true")
	}
	if got := atomicLevel.Level(); got != apilevel.Error {
		t.Fatalf("Level after TrySetLevel(Error) = %v, want %v", got, apilevel.Error)
	}
	if !atomicLevel.TrySetLevel(apilevel.Off) {
		t.Fatal("TrySetLevel(Off) = false, want true")
	}
}

func TestAtomicLevelTrySetLevelReturnsFalseOnInvalid(t *testing.T) {
	t.Parallel()

	atomicLevel := NewAtomicLevelAt(apilevel.Warn)
	if atomicLevel.TrySetLevel(apilevel.Level(16)) {
		t.Fatal("TrySetLevel(Level(16)) = true, want false")
	}
	if got := atomicLevel.Level(); got != apilevel.Warn {
		t.Fatalf("Level after failed TrySetLevel = %v, want %v", got, apilevel.Warn)
	}
}

func TestAtomicLevelMarshalText(t *testing.T) {
	t.Parallel()

	got, err := NewAtomicLevelAt(apilevel.Critical).MarshalText()
	if err != nil {
		t.Fatalf("MarshalText returned error: %v", err)
	}
	if string(got) != "critical" {
		t.Fatalf("MarshalText = %q, want %q", string(got), "critical")
	}
}

func TestAtomicLevelUnmarshalText(t *testing.T) {
	t.Parallel()

	atomicLevel := NewAtomicLevel()
	if err := atomicLevel.UnmarshalText([]byte("error3")); err != nil {
		t.Fatalf("UnmarshalText returned error: %v", err)
	}
	if got := atomicLevel.Level(); got != apilevel.Level(10) {
		t.Fatalf("Level after UnmarshalText = %v, want %v", got, apilevel.Level(10))
	}
}

func TestAtomicLevelUnmarshalTextLeavesValueUnchangedOnError(t *testing.T) {
	t.Parallel()

	atomicLevel := NewAtomicLevelAt(apilevel.Warn)
	if err := atomicLevel.UnmarshalText([]byte("panic")); err == nil {
		t.Fatal("UnmarshalText returned nil error, want error")
	}
	if got := atomicLevel.Level(); got != apilevel.Warn {
		t.Fatalf("Level after failed UnmarshalText = %v, want %v", got, apilevel.Warn)
	}
}

func TestAtomicLevelConcurrentSetAndEnabled(t *testing.T) {
	t.Parallel()

	atomicLevel := NewAtomicLevel()
	levels := []apilevel.Level{apilevel.Trace, apilevel.Debug, apilevel.Info, apilevel.Warn, apilevel.Error, apilevel.Off}

	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func(offset int) {
			defer wg.Done()
			for j := 0; j < 1_000; j++ {
				atomicLevel.SetLevel(levels[(offset+j)%len(levels)])
				_ = atomicLevel.Enabled(apilevel.Error)
			}
		}(i)
	}
	wg.Wait()
}

func assertPanic(t *testing.T, fn func()) {
	t.Helper()

	defer func() {
		if recover() == nil {
			t.Fatal("panic = nil, want panic")
		}
	}()

	fn()
}
