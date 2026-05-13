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

import (
	"errors"
	"testing"
)

func TestSinkContract(t *testing.T) {
	t.Parallel()

	sink := &fakeSink{writeN: len("record\n")}
	var _ Sink = sink

	n, err := sink.Write([]byte("record\n"))
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	if n != len("record\n") {
		t.Fatalf("Write() n = %d, want %d", n, len("record\n"))
	}
	if got := string(sink.written); got != "record\n" {
		t.Fatalf("written bytes = %q, want %q", got, "record\n")
	}

	if err := sink.Sync(); err != nil {
		t.Fatalf("Sync() error = %v", err)
	}
	if !sink.synced {
		t.Fatal("Sync() did not mark sink as synced")
	}
}

func TestSinkWriteError(t *testing.T) {
	t.Parallel()

	want := errors.New("write failed")
	sink := &fakeSink{writeErr: want}

	if _, err := sink.Write([]byte("record\n")); !errors.Is(err, want) {
		t.Fatalf("Write() error = %v, want %v", err, want)
	}
}

func TestSinkSyncError(t *testing.T) {
	t.Parallel()

	want := errors.New("sync failed")
	sink := &fakeSink{syncErr: want}

	if err := sink.Sync(); !errors.Is(err, want) {
		t.Fatalf("Sync() error = %v, want %v", err, want)
	}
}

// fakeSink is a small test double that implements Sink without adding
// production helper implementations to the sink package.
type fakeSink struct {
	written  []byte
	writeN   int
	writeErr error
	synced   bool
	syncErr  error
}

func (s *fakeSink) Write(p []byte) (int, error) {
	s.written = append(s.written, p...)
	return s.writeN, s.writeErr
}

func (s *fakeSink) Sync() error {
	s.synced = true
	return s.syncErr
}
