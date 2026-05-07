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

package predicate

import "arcoris.dev/arclog/api/core"

// Entry is the core entry metadata visible to Predicate implementations.
//
// Entry is an alias for core.Entry so the core package remains the single source
// of truth for log-entry metadata. The predicate package adds filtering
// semantics around that value; it does not define an alternate entry model.
//
// All value semantics, zero-value behavior, and retention rules are inherited
// from core.Entry. In particular, Stack may retain caller-owned frame storage
// unless the runtime cloned the entry before invoking a predicate. Predicate
// implementations that retain Entry beyond a ShouldLog call must clone it first
// when they cannot prove the runtime already transferred ownership.
//
// Structured fields are passed separately to Predicate.ShouldLog. That keeps
// entry metadata and field ownership independent, and it lets future hook and
// core layers decide when field slices are borrowed, copied, or extended.
type Entry = core.Entry
