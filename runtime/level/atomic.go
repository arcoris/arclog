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
	"errors"
	"fmt"
	"sync/atomic"

	apilevel "arcoris.dev/arclog/api/level"
)

// AtomicLevel is a concurrency-safe mutable severity threshold.
//
// The zero value is ready to use and represents api/level.Info because Info is
// the zero value of api/level.Level. Off is accepted as a threshold and disables
// all record severities.
type AtomicLevel struct {
	level atomic.Int32
}

var _ apilevel.Threshold = (*AtomicLevel)(nil)

// NewAtomicLevel returns an AtomicLevel initialized to api/level.Info.
func NewAtomicLevel() *AtomicLevel {
	return &AtomicLevel{}
}

// NewAtomicLevelAt returns an AtomicLevel initialized to lvl.
//
// NewAtomicLevelAt panics if lvl is not a valid severity threshold or Off.
func NewAtomicLevelAt(lvl apilevel.Level) *AtomicLevel {
	a := &AtomicLevel{}
	a.SetLevel(lvl)
	return a
}

// ParseAtomicLevel parses s and returns an AtomicLevel initialized to it.
//
// ParseAtomicLevel accepts the same names and aliases as api/level.Parse.
func ParseAtomicLevel(s string) (*AtomicLevel, error) {
	lvl, err := apilevel.Parse(s)
	if err != nil {
		return nil, err
	}
	return NewAtomicLevelAt(lvl), nil
}

// Enabled reports whether lvl passes the current threshold.
//
// A nil *AtomicLevel disables all record levels. This makes optional runtime
// thresholds safe to call when no threshold has been configured.
func (a *AtomicLevel) Enabled(lvl apilevel.Level) bool {
	if a == nil {
		return false
	}
	return lvl.Enabled(a.Level())
}

// Level returns the current threshold.
//
// A nil *AtomicLevel reports Off, which matches its disabled Enabled behavior.
func (a *AtomicLevel) Level() apilevel.Level {
	if a == nil {
		return apilevel.Off
	}
	return apilevel.Level(a.level.Load())
}

// SetLevel updates the current threshold.
//
// SetLevel accepts valid severity thresholds and Off. It panics if lvl is out
// of range so invalid configuration does not silently change logging behavior.
func (a *AtomicLevel) SetLevel(lvl apilevel.Level) {
	if !a.TrySetLevel(lvl) {
		panic(fmt.Sprintf("runtime/level: invalid threshold %d", int(lvl)))
	}
}

// TrySetLevel updates the current threshold if lvl is valid.
//
// It returns true for valid severity thresholds and Off. It returns false and
// leaves the threshold unchanged for out-of-range values and nil receivers.
func (a *AtomicLevel) TrySetLevel(lvl apilevel.Level) bool {
	if a == nil || !lvl.IsThreshold() {
		return false
	}
	a.level.Store(int32(lvl))
	return true
}

// MarshalText implements encoding.TextMarshaler for the current threshold.
func (a *AtomicLevel) MarshalText() ([]byte, error) {
	return a.Level().MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
//
// On parse error the current level is left unchanged. A nil receiver returns an
// explicit error instead of panicking.
func (a *AtomicLevel) UnmarshalText(text []byte) error {
	if a == nil {
		return errors.New("runtime/level.AtomicLevel.UnmarshalText: nil receiver")
	}

	lvl, err := apilevel.Parse(string(text))
	if err != nil {
		return err
	}
	a.SetLevel(lvl)
	return nil
}
