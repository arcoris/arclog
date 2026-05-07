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
	"testing"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
)

func BenchmarkTeeSingleCoreConstruction(b *testing.B) {
	c := &recordingCore{}

	b.ReportAllocs()
	for b.Loop() {
		_ = core.Tee(nil, c, nil)
	}
}

func BenchmarkTeeMultiCoreConstruction(b *testing.B) {
	first := &recordingCore{}
	second := &recordingCore{}

	b.ReportAllocs()
	for b.Loop() {
		_ = core.Tee(first, nil, second)
	}
}

func BenchmarkCheckedEntryWriteSingleCore(b *testing.B) {
	entry := core.Entry{Level: level.Info, Message: "bench"}
	fields := []field.Field{field.String("k", "v")}

	b.ReportAllocs()
	for b.Loop() {
		var ce *core.CheckedEntry
		ce = ce.AddCore(entry, &recordingCore{})
		_ = ce.Write(fields...)
	}
}
