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
	"fmt"

	apilevel "arcoris.dev/arclog/api/level"
)

// StaticThreshold is an immutable severity threshold.
//
// StaticThreshold is useful for fixed core graphs where no runtime mutation is
// needed. It implements api/level.Threshold for compatibility, but SetLevel does
// not change the stored value.
type StaticThreshold struct {
	level apilevel.Level
}

var _ apilevel.Threshold = StaticThreshold{}

// NewStaticThreshold returns an immutable threshold at lvl.
//
// NewStaticThreshold panics if lvl is not a valid severity threshold or Off.
func NewStaticThreshold(lvl apilevel.Level) StaticThreshold {
	if !lvl.IsThreshold() {
		panic(fmt.Sprintf("runtime/level: invalid threshold %d", int(lvl)))
	}
	return StaticThreshold{level: lvl}
}

// Enabled reports whether lvl passes the static threshold.
func (s StaticThreshold) Enabled(lvl apilevel.Level) bool {
	return lvl.Enabled(s.level)
}

// Level returns the static threshold level.
func (s StaticThreshold) Level() apilevel.Level {
	return s.level
}

// SetLevel validates lvl but does not mutate the static threshold.
//
// This method exists so StaticThreshold can satisfy api/level.Threshold in code
// that accepts either mutable or immutable thresholds. Use AtomicLevel when
// runtime mutation is required.
func (s StaticThreshold) SetLevel(lvl apilevel.Level) {
	if !lvl.IsThreshold() {
		panic(fmt.Sprintf("runtime/level: invalid threshold %d", int(lvl)))
	}
}
