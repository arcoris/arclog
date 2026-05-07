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

package core

import (
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
)

// NoopCore is a Core that never enables entries and drops all writes.
//
// The zero value is ready to use. NoopCore has no mutable state and is safe for
// concurrent use.
type NoopCore struct{}

// Noop returns a no-op Core.
//
// It is useful as an explicit disabled Core when nil would make ownership or
// configuration logic ambiguous.
func Noop() Core {
	return NoopCore{}
}

// Enabled reports false for every level.
func (NoopCore) Enabled(level.Level) bool {
	return false
}

// With returns the receiver unchanged.
func (c NoopCore) With([]field.Field) Core {
	return c
}

// Check returns ce unchanged.
func (NoopCore) Check(_ Entry, ce *CheckedEntry) *CheckedEntry {
	return ce
}

// Write drops entry and fields.
func (NoopCore) Write(Entry, []field.Field) error {
	return nil
}

// Sync is a no-op.
func (NoopCore) Sync() error {
	return nil
}
