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
	"strings"
)

// Parse converts a textual level name into a Level.
//
// Parse trims leading and trailing whitespace and performs a case-insensitive
// match. It accepts canonical names and a small set of common aliases:
//
//	"info", "information", "informational" -> Info
//	"warn", "warning" -> Warn
//	"error", "err" -> Error
//	"critical", "crit" -> Critical
//
// Unknown names return Invalid and a non-nil error. Callers should treat that
// error as invalid configuration or invalid user input rather than silently
// falling back to a default level.
func Parse(s string) (Level, error) {
	name := strings.TrimSpace(strings.ToLower(s))
	switch name {
	case "trace":
		return Trace, nil
	case "debug":
		return Debug, nil
	case "info", "information", "informational":
		return Info, nil
	case "notice":
		return Notice, nil
	case "warn", "warning":
		return Warn, nil
	case "error", "err":
		return Error, nil
	case "critical", "crit":
		return Critical, nil
	case "fatal":
		return Fatal, nil
	case "panic":
		return Panic, nil
	default:
		return Invalid, fmt.Errorf("level.Parse: unknown level %q", s)
	}
}

// MustParse returns the parsed level or panics if s is not recognized.
//
// MustParse is intended for hard-coded package-level defaults and tests. It
// should not be used for user input, configuration files, environment values,
// command-line flags, or network input because those paths need ordinary error
// handling.
func MustParse(s string) Level {
	lvl, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return lvl
}
