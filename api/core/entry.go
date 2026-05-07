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
	"time"

	"arcoris.dev/arclog/api/caller"
	"arcoris.dev/arclog/api/level"
	"arcoris.dev/arclog/api/stack"
)

// Entry contains the metadata associated with a log entry.
//
// Entry does not include structured fields. Fields are passed separately to Core
// methods so call-site fields, context fields, hook-added fields, and fields
// attached by Core.With can keep their own ownership and lifecycle rules.
//
// The zero value is valid and represents an empty entry. Level is the zero value
// of level.Level; code that needs to reject invalid severities should call
// level.Level.IsValid explicitly.
type Entry struct {
	// Time is the timestamp associated with the entry.
	//
	// A zero Time means that the runtime pipeline has not assigned a timestamp or
	// intentionally omitted it.
	Time time.Time

	// Level is the severity level associated with the entry.
	Level level.Level

	// LoggerName is the logical logger name associated with the entry.
	//
	// The empty string means that no logger name was supplied or exposed to the
	// core layer.
	LoggerName string

	// Message is the human-readable log message.
	Message string

	// Caller describes the selected call site, when caller annotation is enabled.
	//
	// Caller may be undefined because caller capture is optional and may be
	// expensive.
	Caller caller.Caller

	// Stack describes the captured stack trace, when stack capture is enabled.
	//
	// Stack may be empty because stack capture is optional and usually reserved
	// for higher-severity entries.
	Stack stack.Stack
}

// Clone returns an Entry with independent stack-frame storage.
//
// Other Entry fields are value fields and are copied directly. Clone is useful
// when an implementation needs to retain Entry after the current call and cannot
// rely on the caller-owned Stack frame slice remaining unchanged.
func (e Entry) Clone() Entry {
	if !e.Stack.IsEmpty() {
		e.Stack = stack.Clone(e.Stack.Frames())
	}

	return e
}

// IsZero reports whether e is the zero Entry value.
//
// IsZero is exact: a valid entry with only Level set to the zero-value level is
// considered zero. Use explicit field checks for semantic validation.
func (e Entry) IsZero() bool {
	return e.Time.IsZero() &&
		e.Level == 0 &&
		e.LoggerName == "" &&
		e.Message == "" &&
		!e.Caller.Defined &&
		e.Caller.PC == 0 &&
		e.Caller.File == "" &&
		e.Caller.Line == 0 &&
		e.Caller.Function == "" &&
		e.Stack.IsEmpty()
}
