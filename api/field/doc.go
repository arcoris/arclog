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

// Package field defines compact structured logging fields for arclog API
// packages.
//
// Field is a small tagged value that describes how a single key-value pair
// should be appended to an encoder. It is intentionally closer to zap's
// allocation-aware field union than to zerolog's mutable event builder: fields
// are value objects that can be created at call sites, passed through level and
// core checks, and dispatched later only if the log entry is actually written.
//
// # Responsibility boundary
//
// This package owns field representation and field-to-encoder dispatch. It does
// not own concrete JSON, console, or binary formatting; it does not write to
// sinks; it does not decide whether an entry is enabled; and it does not manage
// logger, core, hook, or writer lifecycle. Concrete formatting belongs to
// runtime encoders that implement encoder.ObjectEncoder and encoder.ArrayEncoder.
//
// # Dependency direction
//
// Field dispatch depends on encoder contracts, but encoder contracts deliberately
// do not depend on fields. The intended API graph is:
//
//	field -> encoder -> buffer
//
// Keeping this direction prevents the API cycle that occurs when an entry-level
// encoder tries to depend on []field.Field while Field.AddTo also needs encoder
// methods.
//
// # Allocation model
//
// Primitive constructors store values directly in Field.Integer or Field.String
// when possible. Constructors for marshalers, errors, stringers, byte slices,
// and reflected values retain references in Field.Interface. Any is a
// convenience constructor; callers on hot paths should prefer typed
// constructors.
//
// Constructors treat nil and typed-nil marshalers, errors, and fmt.Stringer
// values as explicit nil or skip fields according to the constructor contract.
// Non-nil error and fmt.Stringer values are converted during AddTo through the
// encoder/convert package. That conversion is strict: Error and String panics
// propagate to the caller because panic recovery and diagnostic formatting are
// runtime policy, not part of the field contract.
//
// # Ownership
//
// Constructors do not copy byte slices or user-provided objects. If a caller
// passes a []byte, marshaler, error, stringer, or reflected value, that value
// must remain safe to observe until the field is encoded. Use an application-
// level copy before constructing a field when mutation or asynchronous retention
// is possible.
package field
