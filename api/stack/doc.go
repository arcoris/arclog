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

// Package stack defines API-side contracts for stack trace metadata attached to
// arclog entries.
//
// The package is a value and contract layer. It does not walk the Go runtime
// stack, symbolize program counters, trim file paths, format stack traces, own
// stack pools, or decide when stack traces should be captured. Those concerns
// belong to runtime packages that depend on this API package.
//
// # Model
//
// A Stack is represented as an ordered sequence of caller.Caller frames. Reusing
// caller.Caller avoids a second source-location model for file, line, function,
// and program-counter data. Frames are expected to be ordered from the most
// specific selected application frame outward toward older callers. The exact
// skip policy, wrapper-frame hiding, runtime-frame filtering, and maximum depth
// are capture policy and must be documented by the Capturer implementation.
//
// # Ownership
//
// Stack is a small value that references a frame slice. New retains the supplied
// slice without copying it; this keeps stack-capture paths allocation-aware but
// makes ownership explicit. Clone creates an independent snapshot for paths that
// may outlive the caller-owned frame slice, such as asynchronous hooks or sinks.
//
// # API boundary
//
// This package intentionally does not expose runtime.Frame, runtime.Callers, a
// pooled mutable stack iterator, or zap-style First/Full capture depth. Those
// mechanisms are runtime implementation details. Runtime packages may use them
// internally and still return this package's stable Stack value.
//
// # Encoding
//
// Encoder is a value-specific append contract. It preserves arclog's
// returned-buffer style, but it does not define JSON, console, or binary stack
// formatting. Concrete runtime encoders decide whether a stack is rendered as a
// multiline string, an array of frames, a compact frame list, or another format.
package stack
