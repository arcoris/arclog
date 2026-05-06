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

// Level identifies the severity of a log entry.
//
// The numeric order is part of the API contract for comparison purposes: lower
// values are more verbose and less severe, while higher values are less verbose
// and more severe. Code may compare valid Level values directly, but should not
// persist or expose the raw numeric representation as a stable wire format.
type Level int8

const (
	// Trace is the most verbose diagnostic level.
	//
	// Use Trace for very fine-grained internal diagnostics such as detailed state
	// transitions, tight-loop observations, and temporary investigation output.
	// Trace is normally disabled outside targeted debugging sessions.
	Trace Level = iota - 1

	// Debug is a verbose diagnostic level for development and troubleshooting.
	//
	// Use Debug for control-flow decisions, configuration details, and state
	// changes that are useful during active investigation but too noisy for the
	// normal operational stream.
	Debug

	// Info is the baseline level for normal operational events.
	//
	// Use Info for service lifecycle events, important state transitions, and
	// steady-state observability that operators can keep enabled by default.
	Info

	// Notice marks noteworthy but expected conditions.
	//
	// Use Notice for events that are unusual enough to be visible above Info, but
	// still remain within the expected operating envelope and do not yet indicate
	// degraded behavior.
	Notice

	// Warn marks undesirable or degraded conditions that do not necessarily make
	// the current operation fail.
	//
	// Use Warn for transient failures with retry, soft-limit pressure, degraded
	// dependencies, fallback paths, or conditions that may become errors if they
	// persist.
	Warn

	// Error marks a failed operation after which the process can continue.
	//
	// Use Error when a logical operation failed as intended to be observed by
	// callers or operators. Error is for operation-level failures, not for normal
	// informational messages and not for process-terminating failures.
	Error

	// Critical marks severe system-level degradation or integrity risk.
	//
	// Use Critical when the system is still running but correctness, availability,
	// or data integrity is materially compromised. Critical is more severe than an
	// isolated operation failure and is often suitable for incident workflows.
	Critical

	// Fatal marks an unrecoverable operational failure that should terminate the
	// process or component.
	//
	// Fatal is for failures such as invalid mandatory configuration, startup
	// initialization failure, unavailable required dependencies, or unsafe runtime
	// state where continuing execution would be misleading.
	Fatal

	// Panic marks violated invariants and programmer errors.
	//
	// Panic is intended for impossible states, corrupted internal structures, and
	// bugs that should fail loudly through panic semantics rather than through an
	// orderly operational shutdown.
	Panic

	// Invalid is a sentinel returned for failed parsing and out-of-range values.
	//
	// Invalid is not a valid log-entry severity and must not be used as a normal
	// threshold. Public APIs should reject Invalid instead of silently coercing it
	// to another level.
	Invalid = Panic + 1
)

const (
	// minLevel is the lowest real severity and anchors IsValid.
	minLevel = Trace

	// maxLevel is the highest real severity and keeps Invalid outside the valid
	// range even though it is numerically adjacent to Panic.
	maxLevel = Panic
)

// IsValid reports whether l is one of the defined log-entry severities.
//
// IsValid returns true for Trace through Panic, inclusive. It returns false for
// Invalid and for any out-of-range numeric value.
func (l Level) IsValid() bool {
	return l >= minLevel && l <= maxLevel
}
