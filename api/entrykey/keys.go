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

const (
	// Time is the canonical key for the timestamp associated with a log entry.
	Time Key = "time"

	// Level is the canonical key for the severity level associated with a log
	// entry.
	Level Key = "level"

	// Logger is the canonical key for the logger name associated with a log
	// entry.
	Logger Key = "logger"

	// Message is the canonical key for the human-readable log message.
	Message Key = "message"

	// Caller is the canonical key for source-code call-site information.
	Caller Key = "caller"

	// Function is the canonical key for the function name associated with a
	// caller frame.
	Function Key = "function"

	// Stacktrace is the canonical key for stack trace information.
	Stacktrace Key = "stacktrace"

	// Error is the conventional key for an error value attached to a log entry.
	//
	// Error is included because error fields are common enough to benefit from a
	// shared spelling. It does not imply special runtime treatment unless a
	// concrete field constructor, encoder, or core explicitly documents such
	// behavior.
	Error Key = "error"
)

var metadataKeys = [...]Key{
	Time,
	Level,
	Logger,
	Message,
	Caller,
	Function,
	Stacktrace,
}

var knownKeys = [...]Key{
	Time,
	Level,
	Logger,
	Message,
	Caller,
	Function,
	Stacktrace,
	Error,
}

// Metadata returns the canonical keys that describe log-entry metadata.
//
// The returned slice is a copy and may be modified by the caller.
func Metadata() []Key {
	return cloneKeys(metadataKeys[:])
}

// Known returns every canonical key defined by this package.
//
// The returned slice is a copy and may be modified by the caller.
func Known() []Key {
	return cloneKeys(knownKeys[:])
}

// cloneKeys returns a caller-owned copy of keys so exported vocabulary lists
// cannot be mutated through a returned slice.
func cloneKeys(keys []Key) []Key {
	out := make([]Key, len(keys))
	copy(out, keys)
	return out
}
