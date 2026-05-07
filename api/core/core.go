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

// Core is the minimal contract for a component that accepts checked log entries.
//
// Core implementations decide whether they are enabled for a level, can attach
// structured context through With, participate in the check phase, write entries
// selected during the check phase, and synchronize buffered state.
//
// Core is a contract, not a concrete encoder-backed implementation. A Core may
// write to a file, network sink, memory sink, test recorder, fan-out composite,
// or another downstream component.
//
// Unless a concrete implementation documents otherwise, Core methods are
// expected to be safe for concurrent use by multiple goroutines. Field slices
// passed to Core methods are borrowed and must be treated as read-only.
type Core interface {
	level.Enabler

	// With returns a Core that includes fields in every subsequent write.
	//
	// Implementations should treat the input slice as caller-owned and copy or
	// encode it if the returned Core retains the fields after With returns.
	With(fields []field.Field) Core

	// Check adds this Core to ce when it will write entry.
	//
	// If the Core will not write entry, Check must return ce unchanged. Check must
	// not perform expensive encoding or sink I/O. Check may be called with a nil
	// CheckedEntry; implementations should pass that value to AddCore when they
	// decide to write.
	Check(entry Entry, ce *CheckedEntry) *CheckedEntry

	// Write writes entry and fields.
	//
	// Write is called only after the check phase selected the Core. It must not
	// repeat the level or predicate decision made by Check. Implementations that
	// retain fields beyond the call must copy the slice.
	Write(entry Entry, fields []field.Field) error

	// Sync flushes or synchronizes buffered state held by the Core.
	Sync() error
}
