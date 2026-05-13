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

// Package field defines typed structured log field descriptors.
//
// A Field describes a key, value kind, and compact value storage. It is a
// descriptor, not an encoder operation. The package does not write JSON, OTLP,
// binary streams, bytes, or sink output. Runtime encoders decide how to apply
// fields to physical output.
//
// The zero value of Field is a skipped field.
//
// Constructors do not validate or normalize keys. Key policy belongs to
// encoders or higher-level policy layers.
//
// Bytes is first-class borrowed storage for BytesType. Callers must not mutate
// the slice until the log call or context binding that consumes the field has
// finished observing it.
//
// Any is a convenience constructor, not a field storage kind. It maps common Go
// values to typed fields and falls back to Reflect.
//
// Object, array, and custom marshaling contracts are intentionally out of scope
// until api/encoder exists. Field.AddTo is intentionally not part of this
// package.
package field
