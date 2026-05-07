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

package writer

import "io"

// WriteSyncer is the minimal sink contract for encoded log records.
//
// Write receives already-encoded bytes. Sync flushes or synchronizes any
// buffered state held by the sink. The contract intentionally embeds io.Writer
// instead of defining a custom write method so existing Go sinks can be adapted
// by runtime packages without changing their write semantics.
//
// WriteSyncer does not require implementations to be safe for concurrent use.
// A concrete implementation or wrapper must document and provide any required
// synchronization.
//
// WriteSyncer does not own resource lifetime. It does not include Close because
// loggers often write to resources owned by the application, process supervisor,
// test harness, or runtime integration.
type WriteSyncer interface {
	io.Writer
	Syncer
}
