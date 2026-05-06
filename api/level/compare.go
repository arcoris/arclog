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

package level

// Enabled reports whether l passes the inclusive threshold check.
//
// Enabled returns true when both l and threshold are valid levels and l is
// greater than or equal to threshold. Invalid and out-of-range values are never
// considered enabled, even if their raw numeric values would otherwise satisfy
// the comparison.
func (l Level) Enabled(threshold Level) bool {
	return l.IsValid() && threshold.IsValid() && l >= threshold
}

// Higher reports whether l is valid and more severe than other.
//
// If either operand is Invalid or out of range, Higher returns false. Invalid
// values do not participate in the severity ordering.
func (l Level) Higher(other Level) bool {
	return l.IsValid() && other.IsValid() && l > other
}

// Lower reports whether l is valid and less severe than other.
//
// If either operand is Invalid or out of range, Lower returns false. Invalid
// values do not participate in the severity ordering.
func (l Level) Lower(other Level) bool {
	return l.IsValid() && other.IsValid() && l < other
}
