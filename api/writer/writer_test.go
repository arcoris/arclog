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

package writer_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"arcoris.dev/arclog/api/writer"
)

func TestWriteSyncerContract(t *testing.T) {
	t.Parallel()

	sink := &recordingSink{}
	var _ writer.WriteSyncer = sink

	n, err := io.WriteString(sink, "record\n")
	if err != nil {
		t.Fatalf("WriteString() error = %v", err)
	}
	if n != len("record\n") {
		t.Fatalf("WriteString() n = %d, want %d", n, len("record\n"))
	}
	if got := sink.String(); got != "record\n" {
		t.Fatalf("written bytes = %q, want %q", got, "record\n")
	}

	if err := sink.Sync(); err != nil {
		t.Fatalf("Sync() error = %v", err)
	}
	if !sink.synced {
		t.Fatal("Sync() did not mark sink as synced")
	}
}

func TestWriteSyncerSyncError(t *testing.T) {
	t.Parallel()

	want := errors.New("sync failed")
	sink := &recordingSink{syncErr: want}

	if err := sink.Sync(); !errors.Is(err, want) {
		t.Fatalf("Sync() error = %v, want %v", err, want)
	}
}

// recordingSink is a small test double that implements writer.WriteSyncer
// without adding production helper implementations to the writer package.
type recordingSink struct {
	bytes.Buffer
	synced  bool
	syncErr error
}

func (s *recordingSink) Sync() error {
	s.synced = true
	return s.syncErr
}
