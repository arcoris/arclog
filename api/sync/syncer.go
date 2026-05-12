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

package sync

// Syncer flushes or synchronizes implementation-defined sink state.
//
// Durability guarantees depend on the implementation. Sync may flush buffered
// bytes, forward to fsync, flush an exporter client, or do nothing when there is
// no buffered state to flush.
//
// Syncer does not define ownership or closing semantics. Closing a file,
// network connection, or other resource is the responsibility of the component
// that owns that resource, not the logging API contract.
type Syncer interface {
	// Sync flushes or synchronizes implementation-defined sink state.
	Sync() error
}

// SyncFunc adapts a function to Syncer.
//
// A nil SyncFunc is invalid and will panic when Sync is called.
type SyncFunc func() error

// Sync calls f.
func (f SyncFunc) Sync() error {
	return f()
}
