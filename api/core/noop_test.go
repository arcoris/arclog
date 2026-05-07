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
	"arcoris.dev/arclog/api/level"
)

func TestNoopCore(t *testing.T) {
	t.Parallel()

	noop := core.Noop()
	var ce *core.CheckedEntry
	ce = ce.AddCore(core.Entry{Level: level.Info}, &recordingCore{})

	if noop.Enabled(level.Info) {
		t.Fatal("Noop core reported enabled")
	}
	if got := noop.With(nil); got == nil {
		t.Fatal("Noop.With returned nil")
	}
	if ce := noop.Check(core.Entry{Level: level.Info}, nil); ce != nil {
		t.Fatal("Noop.Check returned non-nil CheckedEntry")
	}
	if got := noop.Check(core.Entry{Level: level.Info}, ce); got != ce {
		t.Fatal("Noop.Check did not preserve existing CheckedEntry")
	}
	if err := noop.Write(core.Entry{}, nil); err != nil {
		t.Fatalf("Noop.Write error = %v", err)
	}
	if err := noop.Sync(); err != nil {
		t.Fatalf("Noop.Sync error = %v", err)
	}
}
