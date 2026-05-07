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

package core_test

import (
	"reflect"
	"sync"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
)

type recordingCore struct {
	mu          sync.Mutex
	name        string
	enabled     bool
	calls       *[]string
	withFields  []field.Field
	writeEntry  core.Entry
	writeFields []field.Field
	writeErr    error
	syncErr     error
	syncCalls   int
}

var _ core.Core = (*recordingCore)(nil)

func entriesEqual(a, b core.Entry) bool {
	return reflect.DeepEqual(a, b)
}

func (c *recordingCore) Enabled(level.Level) bool {
	return c.enabled
}

func (c *recordingCore) With(fields []field.Field) core.Core {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.withFields = append([]field.Field(nil), fields...)
	return c
}

func (c *recordingCore) Check(entry core.Entry, ce *core.CheckedEntry) *core.CheckedEntry {
	if !c.Enabled(entry.Level) {
		return ce
	}

	return ce.AddCore(entry, c)
}

func (c *recordingCore) Write(entry core.Entry, fields []field.Field) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.calls != nil {
		*c.calls = append(*c.calls, c.name)
	}
	c.writeEntry = entry
	c.writeFields = append([]field.Field(nil), fields...)

	return c.writeErr
}

func (c *recordingCore) Sync() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.syncCalls++
	return c.syncErr
}
