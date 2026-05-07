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

// Package entrykey defines canonical field names for arclog log-entry metadata.
//
// The package replaces a broader "schema" package name with a narrower and more
// explicit responsibility: stable keys used to encode the standard parts of a
// log entry. It is a vocabulary package, not a JSON Schema package, not a
// validation registry, and not a semantic-conventions package.
//
// # Responsibility boundary
//
// This package names the fields that arclog itself may produce for entry
// metadata: timestamp, severity level, logger name, message, caller, function,
// stack trace, and the conventional error field. It does not reserve
// application-domain names and does not define HTTP, gRPC, database, messaging,
// cloud, user, payment, or tracing semantic conventions.
//
// Concrete encoders may use these keys as defaults, but encoder configuration
// owns omission and renaming policy. For example, an encoder may choose to omit
// Stacktrace by using an empty key in its own configuration. That behavior is
// intentionally outside this package.
//
// # Naming policy
//
// Keys are lower-case ASCII identifiers intended to be stable across encoders
// and integrations. The package uses "message" instead of shortened forms such
// as "msg", and "stacktrace" as one word to match common structured logging
// practice.
//
// # Stability
//
// Renaming or repurposing a key is a breaking API change. Adding a new key is
// usually safe, but it should still be reviewed as part of the public log-entry
// contract because field names are often consumed by downstream log pipelines.
package entrykey
