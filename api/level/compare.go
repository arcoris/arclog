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

// Enabled reports whether l passes threshold as a record severity.
//
// The check is inclusive: a record at the threshold is enabled, as are all more
// severe records. Off disables every record. Off and out-of-range values are
// never enabled as record severities.
func (l Level) Enabled(threshold Level) bool {
	return l.IsSeverity() && threshold.IsSeverity() && l >= threshold
}

// Higher reports whether l is a valid severity that is more severe than other.
//
// Off and out-of-range values do not participate in severity ordering and
// return false.
func (l Level) Higher(other Level) bool {
	return l.IsSeverity() && other.IsSeverity() && l > other
}

// Lower reports whether l is a valid severity that is less severe than other.
//
// Off and out-of-range values do not participate in severity ordering and
// return false.
func (l Level) Lower(other Level) bool {
	return l.IsSeverity() && other.IsSeverity() && l < other
}
