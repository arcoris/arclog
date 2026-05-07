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

// Package api documents the stable extension-contract layer for arclog.
//
// Packages below api define the small contracts and value types that runtime
// loggers, cores, encoders, fields, predicates, writers, clocks, and third-party
// plugins can share. The API layer is the vocabulary for extension points, not
// the place where runtime behavior is implemented. Concrete logger facades,
// concrete core implementations, sinks, JSON or console encoders, caller
// capture, predicate policy wiring, hook managers, writer adapters, and clock
// selection belong outside api.
//
// The root arcoris.dev/arclog package is expected to be the user-facing facade.
// API packages must stay dependency-light and must not import the root package,
// runtime packages, or internal implementation packages. Dependency direction is
// part of the public design: field dispatch may depend on encoder contracts,
// encoder contracts may depend on buffer, and encoder deliberately does not
// import field. The intended direction is:
//
//	field -> encoder -> buffer
//
// Packages that sit above entry metadata, such as predicate, should depend on
// core.Entry rather than inventing their own entry shape. That keeps core as the
// source of truth for entry metadata while still allowing fields to stay
// caller-owned and passed separately.
//
// Encoder contracts intentionally pass and return *buffer.Buffer. Callers must
// continue with the returned buffer, which leaves room for low-allocation paths,
// buffer replacement, and explicit ownership boundaries. Comments in api must
// avoid global zero-allocation promises; allocation guarantees belong to
// benchmarked concrete runtime paths and should remain path-specific.
package api
