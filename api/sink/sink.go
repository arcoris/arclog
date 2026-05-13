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

package sink

import "io"

// Sink consumes already encoded log payload bytes.
//
// Sink is intentionally a composition of io.Writer and Syncer so implementations
// can use the familiar Write and Sync method set. Despite using io.Writer,
// Sink is not a general-purpose writer wrapper: Write receives complete arclog
// payload bytes after runtime encoding has already happened. A Sink should not
// inspect field descriptors, choose JSON formatting, map values to OTLP, or
// write through api/buffer directly.
//
// The slice passed to Write is borrowed and may alias a pooled runtime buffer.
// Implementations must consume p before Write returns, must not mutate p, and
// must not retain or publish p after Write returns unless they first make their
// own copy. Asynchronous, buffered, batching, or queueing sinks must copy p
// before returning when the bytes will be consumed later. Callers may reuse or
// release the backing storage immediately after Write returns.
//
// Sync flushes or synchronizes implementation-defined sink state. It does not
// imply a specific durability level unless the concrete implementation
// documents one.
//
// Sink does not imply concurrency safety. A concrete implementation or wrapper
// must document and provide any required synchronization.
//
// Sink does not include Close and does not own resource lifetime.
type Sink interface {
	io.Writer
	Syncer
}
