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

// Package level defines the stable severity contract used by arclog API types.
//
// This package belongs to the public API module. It is intentionally small: it
// defines the severity value type, parsing and text-marshaling behavior,
// comparison helpers, the Enabler and Threshold contracts, and the minimal
// encoder function shape used by higher-level entry encoders.
//
// Runtime implementations do not belong here. Atomic thresholds, configuration
// loaders, encoder registries, default JSON or console encoders, sampling, and
// sink-specific routing should live in runtime packages that depend on this API
// package, not the other way around.
//
// # Severity order
//
// Level values are ordered from least severe to most severe:
//
//	Trace < Debug < Info < Notice < Warn < Error < Critical < Fatal < Panic
//
// The standard threshold rule is inclusive: a record is enabled when its level
// is valid and greater than or equal to the configured threshold.
//
// The zero value of Level is Info. That makes uninitialized level variables
// represent ordinary operational records, but callers should still validate
// external configuration with Parse or IsValid before using it as a threshold.
//
// Invalid is a sentinel for failed parsing or out-of-range values. It is not a
// valid log-entry severity and must not be accepted as a normal threshold.
package level
