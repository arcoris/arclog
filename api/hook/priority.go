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

// Priority controls ordering among hooks registered for the same phase.
//
// Lower values run earlier. Runtime managers may use Priority when ordering
// registered hooks, but this package does not provide a manager implementation,
// tie-breaking policy, or cross-phase ordering rule.
type Priority int

const (
	// PriorityFirst is a conventional value for hooks that should run before
	// ordinary hooks.
	PriorityFirst Priority = -1000

	// PriorityHigh is a conventional value for hooks that should run before
	// default-priority hooks.
	PriorityHigh Priority = -100

	// PriorityDefault is the default hook ordering value.
	PriorityDefault Priority = 0

	// PriorityLow is a conventional value for hooks that should run after
	// default-priority hooks.
	PriorityLow Priority = 100

	// PriorityLast is a conventional value for hooks that should run after
	// ordinary hooks.
	PriorityLast Priority = 1000
)

// Before reports whether p sorts before other in ascending priority order.
func (p Priority) Before(other Priority) bool {
	return p < other
}
