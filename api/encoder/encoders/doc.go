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

// Package encoders contains small reusable helpers for common API-level
// encoding cases.
//
// The helpers in this package are not complete log encoders. They handle small
// cross-cutting operations such as nil-safe error and fmt.Stringer conversion.
// Runtime encoders may use these helpers from field dispatch code.
//
// Strict helpers call user-provided Error or String methods directly and allow
// panics to propagate. Safe helpers recover those panics and encode a diagnostic
// string in the fixed form "PANIC=<value>". The conversions may allocate
// depending on user implementations and formatting.
package encoders
