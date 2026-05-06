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
// Packages below api define small contracts and value types that runtime
// loggers, encoders, fields, and third-party plugins can share. They are not
// the runtime implementation layer: concrete loggers, cores, sinks, JSON or
// console encoders, caller resolvers, and writer adapters belong outside api.
//
// The root arcoris.dev/arclog package is expected to be the user-facing facade.
// API packages must stay dependency-light and must not import the root package,
// runtime packages, or internal packages. In particular, encoder deliberately
// does not import field; the intended direction is field -> encoder -> buffer,
// not encoder -> field.
//
// Encoder contracts intentionally pass and return *buffer.Buffer. Callers must
// continue with the returned buffer, which leaves room for low-allocation paths,
// buffer replacement, and explicit ownership boundaries. API documentation must
// avoid global zero-allocation promises; allocation guarantees belong to
// benchmarked concrete paths.
package api
