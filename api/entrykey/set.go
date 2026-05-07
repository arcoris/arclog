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

package entrykey

// Set groups configurable log-entry keys.
//
// Set is useful for runtime encoders that allow users to rename or omit standard
// entry fields while still starting from the canonical arclog vocabulary.
//
// The zero value is intentionally empty. Use DefaultSet when the canonical
// arclog names are desired.
type Set struct {
	// Time is the key for the entry timestamp.
	Time Key

	// Level is the key for the entry severity level.
	Level Key

	// Logger is the key for the logger name.
	Logger Key

	// Message is the key for the entry message.
	Message Key

	// Caller is the key for source-code call-site information.
	Caller Key

	// Function is the key for the function name associated with the caller.
	Function Key

	// Stacktrace is the key for stack trace information.
	Stacktrace Key

	// Error is the conventional key for attached error values.
	Error Key
}

// DefaultSet returns the canonical arclog key set.
func DefaultSet() Set {
	return Set{
		Time:       Time,
		Level:      Level,
		Logger:     Logger,
		Message:    Message,
		Caller:     Caller,
		Function:   Function,
		Stacktrace: Stacktrace,
		Error:      Error,
	}
}

// Metadata returns the entry metadata keys from s in canonical emission order.
//
// Error is intentionally excluded because it is a conventional user/error field
// rather than structural metadata that every entry carries.
//
// Empty keys are preserved. This lets encoder configuration use an empty key as
// an omission marker without this package making the omission decision.
func (s Set) Metadata() []Key {
	return []Key{
		s.Time,
		s.Level,
		s.Logger,
		s.Message,
		s.Caller,
		s.Function,
		s.Stacktrace,
	}
}

// Known returns every key in s in canonical order, including Error.
//
// Empty keys are preserved.
func (s Set) Known() []Key {
	return []Key{
		s.Time,
		s.Level,
		s.Logger,
		s.Message,
		s.Caller,
		s.Function,
		s.Stacktrace,
		s.Error,
	}
}

// IsDefault reports whether s equals DefaultSet.
func (s Set) IsDefault() bool {
	return s == DefaultSet()
}

// IsZero reports whether all keys in s are empty.
func (s Set) IsZero() bool {
	return s == Set{}
}
