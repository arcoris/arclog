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

// Parse converts text into a Level.
//
// Parse trims surrounding whitespace and is case-insensitive. It accepts named
// arclog levels, the aliases "warning", "crit", and "disabled", and the OTel
// short names TRACE2 through FATAL4. "panic", "unspecified", and "invalid" are
// intentionally rejected because they are not arclog severity configuration
// values.
func Parse(s string) (Level, error) {
	name := strings.ToLower(strings.TrimSpace(s))
	switch name {
	case "trace":
		return Trace, nil
	case "trace2":
		return trace2, nil
	case "trace3":
		return trace3, nil
	case "trace4":
		return trace4, nil
	case "debug":
		return Debug, nil
	case "debug2":
		return debug2, nil
	case "debug3":
		return debug3, nil
	case "debug4":
		return debug4, nil
	case "info":
		return Info, nil
	case "notice", "info2":
		return info2, nil
	case "info3":
		return info3, nil
	case "info4":
		return info4, nil
	case "warn", "warning":
		return Warn, nil
	case "warn2":
		return warn2, nil
	case "warn3":
		return warn3, nil
	case "warn4":
		return warn4, nil
	case "error":
		return Error, nil
	case "critical", "crit", "error2":
		return error2, nil
	case "error3":
		return error3, nil
	case "error4":
		return error4, nil
	case "fatal":
		return Fatal, nil
	case "fatal2":
		return fatal2, nil
	case "fatal3":
		return fatal3, nil
	case "fatal4":
		return fatal4, nil
	case "off", "disabled":
		return Off, nil
	default:
		return Info, fmt.Errorf("level.Parse: unknown level %q", s)
	}
}

// MustParse returns the parsed Level or panics if s is not recognized.
//
// MustParse is intended for package-level constants, test fixtures, and other
// hard-coded values. User input and configuration paths should call Parse so
// errors can be reported normally.
func MustParse(s string) Level {
	lvl, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return lvl
}
