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

// Package caller defines API-side contracts for call-site metadata attached to
// arclog entries.
//
// Caller is a small value type. Its zero value is an undefined caller:
// Defined is false and the remaining fields are unspecified. When Defined is
// true, File, Line, Function, and PC describe a best-effort source location,
// but path trimming, function-name style, and symbolization fidelity are
// implementation-defined.
//
// Resolver is the capture contract. Runtime packages may implement it with
// stack walking, symbol caches, platform-specific lookup, or deterministic test
// doubles. This package does not perform caller lookup itself. Caller capture
// can be relatively expensive, so capture policy belongs to runtime loggers and
// configuration, not to this API package.
//
// Encoder appends a Caller representation to a buffer and returns the
// authoritative buffer for subsequent writes. Implementations must not release
// dst or the returned buffer; lifetime remains with the caller. Formatting
// choices such as full paths, short paths, placeholders for undefined callers,
// or machine-readable layouts belong to concrete encoders outside this package.
package caller
