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

// String returns the canonical lowercase text for l.
//
// Named levels use arclog names such as "notice" and "critical". Custom
// OpenTelemetry sublevels use lowercase OTel short names such as "trace2",
// "info3", or "fatal4". Off returns "off". Out-of-range values return a
// diagnostic string of the form "level(<number>)".
func (l Level) String() string {
	if l == Off {
		return "off"
	}

	n := l.SeverityNumber()
	if n >= minSeverityNumber && n <= maxSeverityNumber {
		// Indexes are OpenTelemetry SeverityNumber values. Element zero is left
		// empty because OTel SeverityNumber 0 is UNSPECIFIED and not a record
		// severity in arclog.
		return [...]string{
			1:  "trace",
			2:  "trace2",
			3:  "trace3",
			4:  "trace4",
			5:  "debug",
			6:  "debug2",
			7:  "debug3",
			8:  "debug4",
			9:  "info",
			10: "notice",
			11: "info3",
			12: "info4",
			13: "warn",
			14: "warn2",
			15: "warn3",
			16: "warn4",
			17: "error",
			18: "critical",
			19: "error3",
			20: "error4",
			21: "fatal",
			22: "fatal2",
			23: "fatal3",
			24: "fatal4",
		}[n]
	}

	return fmt.Sprintf("level(%d)", int(l))
}
