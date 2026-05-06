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

// Package encoder defines stable encoder-side contracts used by arclog API
// packages and third-party extensions.
//
// The package is intentionally limited to contracts that must be shared between
// field constructors, custom field marshalers, and concrete encoder
// implementations. It does not define a complete log-entry encoder because a
// complete entry encoder would have to depend on the field package, while field
// values need to call these interfaces from Field.AddTo. Keeping this package
// independent from field prevents an import cycle:
//
//	field -> encoder -> buffer
//
// Concrete JSON, console, or binary encoders belong in runtime packages. This
// package only specifies the object and array write contracts that those
// implementations expose to API-level marshalers.
//
// The package depends on buffer and deliberately does not depend on field. That
// direction keeps the API graph acyclic for future field dispatch code:
// field-like packages may call encoder contracts, while encoder contracts stay
// independent from concrete field representations.
//
// # Buffer ownership
//
// Encoder methods append to the supplied buffer and return the buffer that
// should be used by the caller after the operation. Implementations MAY return
// the same buffer or a replacement buffer. Callers MUST continue with the
// returned value.
//
// # Concurrency
//
// ObjectEncoder and ArrayEncoder values are builder-style contracts. Unless a
// concrete implementation documents stronger guarantees, callers MUST treat
// them as not safe for concurrent mutation.
//
// # Error model
//
// Primitive append operations do not return errors. Structured operations that
// execute user-provided marshalers or reflection paths return an error so
// implementations can propagate marshaling failures to the caller.
package encoder
