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
	"errors"
	"testing"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
)

func TestAddCoreOnNilReceiver(t *testing.T) {
	t.Parallel()

	entry := core.Entry{Level: level.Info, Message: "test"}
	fake := &recordingCore{enabled: true}

	var ce *core.CheckedEntry
	ce = ce.AddCore(entry, fake)

	if ce == nil {
		t.Fatal("AddCore returned nil")
	}
	if ce.Len() != 1 {
		t.Fatalf("Len() = %d, want 1", ce.Len())
	}
	if ce.Entry().Message != "test" {
		t.Fatalf("Entry().Message = %q, want test", ce.Entry().Message)
	}
}

func TestAddCoreIgnoresNilCore(t *testing.T) {
	t.Parallel()

	var ce *core.CheckedEntry
	ce = ce.AddCore(core.Entry{}, nil)

	if ce != nil {
		t.Fatal("AddCore with nil core created CheckedEntry")
	}
}

func TestAddCorePreservesFirstEntry(t *testing.T) {
	t.Parallel()

	first := core.Entry{Level: level.Info, Message: "first"}
	second := core.Entry{Level: level.Error, Message: "second"}

	var ce *core.CheckedEntry
	ce = ce.AddCore(first, &recordingCore{})
	ce = ce.AddCore(second, &recordingCore{})

	if got := ce.Entry(); !entriesEqual(got, first) {
		t.Fatalf("Entry() = %#v, want first entry %#v", got, first)
	}
	if ce.Len() != 2 {
		t.Fatalf("Len() = %d, want 2", ce.Len())
	}
}

func TestNilCheckedEntryWriteIsNoop(t *testing.T) {
	t.Parallel()

	var ce *core.CheckedEntry
	if err := ce.Write(field.String("k", "v")); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
}

func TestZeroCheckedEntryWriteIsOneShotNoop(t *testing.T) {
	t.Parallel()

	var ce core.CheckedEntry
	if !ce.IsEmpty() {
		t.Fatal("zero CheckedEntry is not empty")
	}
	if err := ce.Write(); err != nil {
		t.Fatalf("first Write() error = %v", err)
	}
	if err := ce.Write(); !errors.Is(err, core.ErrCheckedEntryWritten) {
		t.Fatalf("second Write() error = %v, want ErrCheckedEntryWritten", err)
	}
}

func TestCheckedEntryWriteCallsCoresInOrder(t *testing.T) {
	t.Parallel()

	var calls []string
	first := &recordingCore{name: "first", enabled: true, calls: &calls}
	second := &recordingCore{name: "second", enabled: true, calls: &calls}

	var ce *core.CheckedEntry
	entry := core.Entry{Level: level.Info}
	ce = ce.AddCore(entry, first)
	ce = ce.AddCore(entry, second)

	if err := ce.Write(field.String("k", "v")); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	want := []string{"first", "second"}
	if len(calls) != len(want) {
		t.Fatalf("calls = %#v, want %#v", calls, want)
	}
	for i := range want {
		if calls[i] != want[i] {
			t.Fatalf("calls = %#v, want %#v", calls, want)
		}
	}
}

func TestCheckedEntryWritePassesEntryAndFields(t *testing.T) {
	t.Parallel()

	entry := core.Entry{Level: level.Warn, Message: "hello"}
	fields := []field.Field{field.String("k", "v")}
	recorder := &recordingCore{}

	var ce *core.CheckedEntry
	ce = ce.AddCore(entry, recorder)

	if err := ce.Write(fields...); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	if !entriesEqual(recorder.writeEntry, entry) {
		t.Fatalf("writeEntry = %#v, want %#v", recorder.writeEntry, entry)
	}
	if len(recorder.writeFields) != 1 || !recorder.writeFields[0].Equal(fields[0]) {
		t.Fatalf("writeFields = %#v, want %#v", recorder.writeFields, fields)
	}
}

func TestCheckedEntryWriteAggregatesErrors(t *testing.T) {
	t.Parallel()

	errA := errors.New("a")
	errB := errors.New("b")

	var ce *core.CheckedEntry
	ce = ce.AddCore(core.Entry{}, &recordingCore{writeErr: errA})
	ce = ce.AddCore(core.Entry{}, &recordingCore{writeErr: errB})

	err := ce.Write()
	if !errors.Is(err, errA) {
		t.Fatalf("Write() error does not contain errA: %v", err)
	}
	if !errors.Is(err, errB) {
		t.Fatalf("Write() error does not contain errB: %v", err)
	}
}

func TestCheckedEntryDoubleWrite(t *testing.T) {
	t.Parallel()

	var ce *core.CheckedEntry
	ce = ce.AddCore(core.Entry{}, &recordingCore{})

	if err := ce.Write(); err != nil {
		t.Fatalf("first Write() error = %v", err)
	}
	if err := ce.Write(); !errors.Is(err, core.ErrCheckedEntryWritten) {
		t.Fatalf("second Write() error = %v, want ErrCheckedEntryWritten", err)
	}
}
