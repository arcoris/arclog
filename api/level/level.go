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

// Level identifies a log record severity or a threshold sentinel.
//
// Record severities use an OpenTelemetry-shifted numeric model. The shift is
// the OpenTelemetry SeverityNumber for INFO, so arclog's Info level can remain
// the zero value:
//
//	Level = OpenTelemetry SeverityNumber - severityNumberShift
//
// The zero value is Info, which makes an uninitialized Level suitable for
// normal operational logging. Custom OpenTelemetry sublevels are valid when
// they fall within Trace through Level(15), inclusive. Off is not a record
// severity; it is only valid as a threshold sentinel.
type Level int8

const (
	// Trace is the most verbose named severity.
	//
	// Trace maps to OpenTelemetry SeverityNumber 1. It is intended for detailed
	// diagnostics that are normally disabled outside targeted investigation.
	Trace Level = -8

	// Debug is a verbose diagnostic severity.
	//
	// Debug maps to OpenTelemetry SeverityNumber 5. It is intended for
	// troubleshooting details that are too noisy for normal operations.
	Debug Level = -4

	// Info is the default operational severity and the zero value of Level.
	//
	// Info maps to OpenTelemetry's INFO SeverityNumber. Use it for normal
	// service lifecycle events and steady-state operational messages.
	Info Level = 0

	// Notice is a named INFO sublevel for noteworthy but expected conditions.
	//
	// Notice maps to OpenTelemetry SeverityNumber 10, equivalent to OTel INFO2.
	Notice Level = 1

	// Warn marks degraded or undesirable conditions that do not necessarily make
	// the current operation fail.
	//
	// Warn maps to OpenTelemetry SeverityNumber 13.
	Warn Level = 4

	// Error marks a failed operation after which the process can continue.
	//
	// Error maps to OpenTelemetry SeverityNumber 17.
	Error Level = 8

	// Critical marks severe degradation or integrity risk while the process is
	// still running.
	//
	// Critical maps to OpenTelemetry SeverityNumber 18, equivalent to OTel
	// ERROR2.
	Critical Level = 18 - severityNumberShift

	// Fatal marks an unrecoverable operational severity.
	//
	// Fatal maps to OpenTelemetry SeverityNumber 21. It does not terminate the
	// process; exit and panic behavior belong to higher runtime or logger layers.
	Fatal Level = 12

	// Off is a threshold-only sentinel that disables all record severities.
	//
	// Off is not a record severity and maps to OpenTelemetry SeverityNumber 0.
	Off Level = 127
)

const (
	// severityNumberShift is the OpenTelemetry SeverityNumber for INFO.
	//
	// arclog stores Info as Level(0), so converting between Level and
	// OpenTelemetry SeverityNumber adds or subtracts this offset.
	severityNumberShift = 9

	minSeverityNumber = 1
	maxSeverityNumber = 24
)

const (
	trace2 Level = -7
	trace3 Level = -6
	trace4 Level = -5

	debug2 Level = -3
	debug3 Level = -2
	debug4 Level = -1

	info2 Level = Notice
	info3 Level = 2
	info4 Level = 3

	warn2 Level = 5
	warn3 Level = 6
	warn4 Level = 7

	error2 Level = Critical
	error3 Level = 10
	error4 Level = 11

	fatal2 Level = 13
	fatal3 Level = 14
	fatal4 Level = 15

	// maxSeverity is the highest valid record severity in the shifted OTel
	// range. Off intentionally sits outside this range as a threshold-only
	// sentinel.
	maxSeverity = fatal4
)

// IsSeverity reports whether l is a valid record severity.
//
// Valid record severities cover the full shifted OpenTelemetry range:
// Trace through Level(15), inclusive. Off and out-of-range values are not
// record severities.
func (l Level) IsSeverity() bool {
	return l >= Trace && l <= maxSeverity
}

// IsNamed reports whether l is one of arclog's named severity levels.
//
// Custom OpenTelemetry sublevels, Off, and out-of-range values return false.
func (l Level) IsNamed() bool {
	switch l {
	case Trace, Debug, Info, Notice, Warn, Error, Critical, Fatal:
		return true
	default:
		return false
	}
}

// IsThreshold reports whether l can be used as a threshold.
//
// Valid record severities and Off are valid thresholds. Out-of-range values are
// rejected so configuration mistakes do not silently enable or disable logs.
func (l Level) IsThreshold() bool {
	return l.IsSeverity() || l == Off
}
