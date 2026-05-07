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

package predicate

import (
	"arcoris.dev/arclog/api/caller"
	"arcoris.dev/arclog/api/level"
)

// Entry carries the entry metadata visible to Predicate implementations.
//
// Entry deliberately contains only API-layer values. It does not reference a
// runtime logger, core, encoder, writer, context, clock, or hook implementation.
// Runtime packages may leave fields at their zero values when the information is
// unavailable or intentionally not captured before predicate evaluation.
//
// The zero value is a valid but mostly unspecified entry: Level is the zero
// value of level.Level, strings are empty, and Caller is undefined. Predicates
// that require a particular metadata value should check it explicitly rather
// than assume the runtime always filled every field.
type Entry struct {
	// Level is the severity associated with the entry.
	//
	// Runtime code should pass a valid level.Level when level information is
	// available. Predicates that care about invalid levels should call
	// level.Level.IsValid explicitly. The zero value is level.Debug in the
	// current level model, but predicates should not use zero-ness as a proxy
	// for "level was provided"; that distinction belongs to the runtime
	// pipeline.
	Level level.Level

	// Logger is the implementation-defined logger name associated with the
	// entry.
	//
	// The empty string means that no logger name was supplied or that the runtime
	// does not expose logger names to predicates. Logger names are ordinary
	// strings; this package does not define hierarchical separators, wildcard
	// matching, or normalization rules.
	Logger string

	// Message is the human-readable entry message before encoding.
	//
	// Predicates may inspect Message for routing or filtering, but should prefer
	// structured fields when the information has a stable key. Message is not
	// encoded or escaped by this package.
	Message string

	// Caller is the source location associated with the entry.
	//
	// Caller may be undefined because caller capture can be expensive. A
	// predicate must check Caller.Defined before relying on File, Line, PC, or
	// Function. Runtime path trimming and symbolization remain outside this
	// package.
	Caller caller.Caller
}
