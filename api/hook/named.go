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

package hook

// Named is an optional diagnostic identity contract for hooks.
//
// Names are useful for debugging, metrics, and configuration reports. They are
// intentionally not required by PreWriteHook, PostWriteHook, ErrorHook, or
// Manager because registration handles are a safer removal mechanism than global
// name uniqueness.
type Named interface {
	// Name returns a stable diagnostic name for the hook.
	Name() string
}
