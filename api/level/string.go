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

import "fmt"

// String returns the canonical lowercase name for l.
//
// For valid levels, String returns one of "trace", "debug", "info",
// "notice", "warn", "error", "critical", "fatal", or "panic". For Invalid,
// it returns "invalid". For any other out-of-range value, it returns a
// diagnostic fallback of the form "Level(<numeric>)".
//
// The fallback is for diagnostics only and should not be used as a stable wire
// format.
func (l Level) String() string {
	switch l {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Notice:
		return "notice"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Critical:
		return "critical"
	case Fatal:
		return "fatal"
	case Panic:
		return "panic"
	case Invalid:
		return "invalid"
	default:
		return fmt.Sprintf("Level(%d)", int8(l))
	}
}
