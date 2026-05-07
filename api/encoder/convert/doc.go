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

// Package convert contains small encoder-bound conversion functions for values
// commonly represented as strings.
//
// The package is not a concrete encoder package. It does not define JSON,
// console, reflection, field dispatch, panic recovery, escaping, quoting, or
// formatting policy beyond converting already-provided error and fmt.Stringer
// values to strings and forwarding those strings to encoder.ObjectEncoder or
// encoder.ArrayEncoder.
//
// Conversion functions preserve the encoder returned-buffer contract. They pass
// dst to the encoder method and return the authoritative buffer returned by that
// method, so callers must continue with the returned value.
//
// Nil and typed-nil error or fmt.Stringer values are encoded as empty strings.
// Non-nil values are strict: Error and String are called directly, and any panic
// from user code propagates to the caller. Panic recovery and diagnostic string
// formatting are runtime policy, not API conversion behavior.
//
// Conversion functions assume the supplied encoder is non-nil. They do not own
// field dispatch validation, so callers that accept optional encoders should
// enforce their nil-encoder policy before calling this package.
//
// The functions in this package should stay small and allocation-conscious. They
// do not allocate by themselves on the normal path; user-provided Error or
// String implementations may allocate.
package convert
