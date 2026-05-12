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

package exit_test

import (
	"errors"
	"testing"

	"arcoris.dev/arclog/runtime/internal/exit"
)

// Tests in this file intentionally do not call t.Parallel. The package under
// test owns a process-global exit function, and stubbing that function is an
// exclusive test operation.

func TestWithStubRecordsExit(t *testing.T) {
	recorder := exit.WithStub(func() {
		exit.With(exit.Failure)
	})

	if !recorder.Exited() {
		t.Fatal("Exited() = false, want true")
	}
	if got := recorder.Code(); got != exit.Failure {
		t.Fatalf("Code() = %d, want %d", got, exit.Failure)
	}
	if got := recorder.Calls(); got != 1 {
		t.Fatalf("Calls() = %d, want 1", got)
	}
	if !recorder.Restored() {
		t.Fatal("Restored() = false, want true")
	}
}

func TestWithStubWithoutExit(t *testing.T) {
	recorder := exit.WithStub(func() {})

	if recorder.Exited() {
		t.Fatal("Exited() = true, want false")
	}
	if got := recorder.Code(); got != 0 {
		t.Fatalf("Code() = %d, want 0", got)
	}
	if got := recorder.Calls(); got != 0 {
		t.Fatalf("Calls() = %d, want 0", got)
	}
	if !recorder.Restored() {
		t.Fatal("Restored() = false, want true")
	}
}

func TestWithStubNilFunction(t *testing.T) {
	recorder := exit.WithStub(nil)

	if recorder.Exited() {
		t.Fatal("Exited() = true, want false")
	}
	if !recorder.Restored() {
		t.Fatal("Restored() = false, want true")
	}
}

func TestWithStubRestoresAfterPanic(t *testing.T) {
	want := errors.New("boom")

	func() {
		defer func() {
			if recovered := recover(); recovered != want {
				t.Fatalf("recover() = %v, want %v", recovered, want)
			}
		}()

		_ = exit.WithStub(func() {
			panic(want)
		})
	}()

	recorder := exit.WithStub(func() {
		exit.With(exit.Failure)
	})

	if !recorder.Exited() {
		t.Fatal("recorder was not installed correctly after panic restore")
	}
}

func TestManualStubAndUnstub(t *testing.T) {
	recorder := exit.Stub()
	defer recorder.Unstub()

	exit.With(exit.Failure)

	if !recorder.Exited() {
		t.Fatal("Exited() = false, want true")
	}
	if got := recorder.Code(); got != exit.Failure {
		t.Fatalf("Code() = %d, want %d", got, exit.Failure)
	}

	recorder.Unstub()
	if !recorder.Restored() {
		t.Fatal("Restored() = false, want true")
	}
}

func TestUnstubIsIdempotent(t *testing.T) {
	recorder := exit.Stub()
	recorder.Unstub()
	recorder.Unstub()

	if !recorder.Restored() {
		t.Fatal("Restored() = false, want true")
	}
}

func TestMultipleExitCallsRecordLastCodeAndCount(t *testing.T) {
	recorder := exit.WithStub(func() {
		exit.With(exit.Success)
		exit.With(exit.Failure)
	})

	if !recorder.Exited() {
		t.Fatal("Exited() = false, want true")
	}
	if got := recorder.Code(); got != exit.Failure {
		t.Fatalf("Code() = %d, want %d", got, exit.Failure)
	}
	if got := recorder.Calls(); got != 2 {
		t.Fatalf("Calls() = %d, want 2", got)
	}
}
