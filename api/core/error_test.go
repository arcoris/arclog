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
)

func TestWriteErrorsErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		errs core.WriteErrors
		want error
	}{
		{name: "empty", errs: nil, want: nil},
		{name: "nil only", errs: core.WriteErrors{nil, nil}, want: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := tt.errs.Err(); err != tt.want {
				t.Fatalf("Err() = %v, want %v", err, tt.want)
			}
		})
	}

	want := errors.New("write failed")
	err := core.AppendError(nil, want).Err()
	if !errors.Is(err, want) {
		t.Fatalf("Err() = %v, want to contain %v", err, want)
	}
}

func TestAppendErrorIgnoresNil(t *testing.T) {
	t.Parallel()

	errs := core.AppendError(nil, nil)
	if len(errs) != 0 {
		t.Fatalf("len(errs) = %d, want 0", len(errs))
	}
}

func TestWriteErrorsErrorString(t *testing.T) {
	t.Parallel()

	errA := errors.New("a")
	errB := errors.New("b")

	tests := []struct {
		name string
		errs core.WriteErrors
		want string
	}{
		{name: "empty", errs: nil, want: ""},
		{name: "nil only", errs: core.WriteErrors{nil}, want: ""},
		{name: "single", errs: core.WriteErrors{errA}, want: "a"},
		{name: "multiple", errs: core.WriteErrors{errA, nil, errB}, want: "multiple core errors: a; b;"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.errs.Error(); got != tt.want {
				t.Fatalf("Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestWriteErrorsUnwrap(t *testing.T) {
	t.Parallel()

	errA := errors.New("a")
	errB := errors.New("b")
	err := core.WriteErrors{errA, nil, errB}

	unwrapped := err.Unwrap()
	if len(unwrapped) != 2 {
		t.Fatalf("len(Unwrap()) = %d, want 2", len(unwrapped))
	}
	unwrapped[0] = nil
	if !errors.Is(err, errA) {
		t.Fatal("mutating Unwrap result changed original WriteErrors")
	}
}
