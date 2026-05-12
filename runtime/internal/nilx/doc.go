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

// Package nilx contains runtime-internal typed-nil inspection.
//
// Runtime adapter and encoder packages often accept interface values that may
// hold nil concrete pointers, slices, maps, functions, channels, or interfaces.
// A direct comparison with nil only catches a nil interface, so those packages
// use IsNil to keep typed-nil behavior consistent at runtime boundaries.
//
// This package is not part of the public arclog API and should not grow into a
// general utilities package. It deliberately contains only the narrow nil
// inspection needed by runtime constructors, adapters, and encoder paths.
package nilx
