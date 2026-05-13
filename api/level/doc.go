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

// Package level defines severity values and threshold checks for arclog.
//
// Level is aligned with OpenTelemetry Logs severity numbers using:
//
//	Level = SeverityNumber - severityNumberShift
//
// severityNumberShift is the OpenTelemetry SeverityNumber for INFO. Using it as
// the offset makes Info the zero value while preserving a direct conversion to
// the OpenTelemetry severity range. The public constants intentionally expose a
// small arclog domain set: Trace, Debug, Info, Notice, Warn, Error, Critical,
// Fatal, and Off.
//
// arclog still supports the full OpenTelemetry SeverityNumber range internally.
// The valid record severity range is Trace through Level(15), which maps to
// OpenTelemetry SeverityNumber 1 through 24. Use FromSeverityNumber to bridge
// OpenTelemetry sublevels into Level values.
//
// Notice is arclog's source severity name for SeverityNumber 10, numerically
// equivalent to OpenTelemetry INFO2. Critical is arclog's source severity name
// for SeverityNumber 18, numerically equivalent to OpenTelemetry ERROR2.
//
// Off is a threshold-only sentinel. It disables all record levels when used as
// a threshold and is never a record severity.
//
// Fatal is only a severity value in this package. Panic is not a level. Panic
// behavior, process exit, retry policy, encoder behavior, sink behavior, and
// runtime side effects are outside this package. Level values do not exit or
// panic as a side effect.
//
// Encoders and exporters may use SeverityNumber and SeverityText for JSON,
// OTLP, or other display formats, but this API package intentionally does not
// import encoder, sink, OpenTelemetry, protobuf, or runtime packages.
package level
