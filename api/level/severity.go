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

// FromSeverityNumber converts an OpenTelemetry SeverityNumber to Level.
//
// The valid OpenTelemetry severity range is minSeverityNumber through
// maxSeverityNumber. Values outside that range return false. Notice is returned
// for SeverityNumber 10, and Critical is returned for SeverityNumber 18.
func FromSeverityNumber(n int) (Level, bool) {
	if n < minSeverityNumber || n > maxSeverityNumber {
		return 0, false
	}
	return Level(n - severityNumberShift), true
}

// SeverityNumber returns the OpenTelemetry-compatible severity number for l.
//
// Valid record severities map by adding severityNumberShift to the Level value.
// Off and out-of-range values return 0, the OpenTelemetry unspecified severity
// number.
func (l Level) SeverityNumber() int {
	if !l.IsSeverity() {
		return 0
	}
	return int(l) + severityNumberShift
}

// SeverityText returns uppercase severity text suitable for JSON and OTLP.
//
// Named levels use arclog display names such as "NOTICE" and "CRITICAL".
// Custom OpenTelemetry sublevels use OTel short names such as "TRACE2",
// "INFO3", or "FATAL4". Off returns "OFF". Out-of-range values return a
// diagnostic string of the form "LEVEL(<number>)".
func (l Level) SeverityText() string {
	if l == Off {
		return "OFF"
	}

	n := l.SeverityNumber()
	if n >= minSeverityNumber && n <= maxSeverityNumber {
		// Indexes are OpenTelemetry SeverityNumber values. Element zero is left
		// empty because OTel SeverityNumber 0 is UNSPECIFIED and not a record
		// severity in arclog.
		return [...]string{
			1:  "TRACE",
			2:  "TRACE2",
			3:  "TRACE3",
			4:  "TRACE4",
			5:  "DEBUG",
			6:  "DEBUG2",
			7:  "DEBUG3",
			8:  "DEBUG4",
			9:  "INFO",
			10: "NOTICE",
			11: "INFO3",
			12: "INFO4",
			13: "WARN",
			14: "WARN2",
			15: "WARN3",
			16: "WARN4",
			17: "ERROR",
			18: "CRITICAL",
			19: "ERROR3",
			20: "ERROR4",
			21: "FATAL",
			22: "FATAL2",
			23: "FATAL3",
			24: "FATAL4",
		}[n]
	}

	return fmt.Sprintf("LEVEL(%d)", int(l))
}
