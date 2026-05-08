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

// Package nilx contains internal typed-nil inspection for API packages.
//
// The package exists because API constructors and encoder-bound conversion code
// often receive interface values that may hold typed nil pointers, maps, slices,
// functions, channels, or interfaces. Direct comparison with nil only catches a
// nil interface value.
//
// This package is an implementation detail of the api tree. It is not a public
// arclog extension contract and should not grow into a generic utility package.
// It exists to keep typed-nil behavior consistent in constructors and
// conversion functions without exporting that inspection policy to plugin
// authors.
package nilx
