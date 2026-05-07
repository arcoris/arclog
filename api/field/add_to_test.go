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

package field_test

import (
	"errors"
	"testing"
	"time"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/encoder"
	"arcoris.dev/arclog/api/field"
)

type testStringer string

func (s testStringer) String() string { return string(s) }

type panicStringer struct{}

func (panicStringer) String() string { panic("stringer panic") }

type panicError struct{}

func (panicError) Error() string { panic("error panic") }

func TestAddToPrimitive(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got, err := field.String("name", "arcoris").AddTo(dst, recordingEncoder{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "name=arcoris;" {
		t.Fatalf("buffer=%q", got.String())
	}
}
func TestAddToSkip(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got, err := field.Skip().AddTo(dst, recordingEncoder{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got != dst {
		t.Fatal("different buffer")
	}
	if got.Len() != 0 {
		t.Fatalf("len=%d", got.Len())
	}
}

func TestAddToSkipAllowsNilEncoder(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got, err := field.Skip().AddTo(dst, nil)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got != dst {
		t.Fatal("different buffer")
	}
}

func TestAddToNilEncoder(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got, err := field.String("name", "arcoris").AddTo(dst, nil)
	if !errors.Is(err, field.ErrNilEncoder) {
		t.Fatalf("err=%v", err)
	}
	if got != dst {
		t.Fatal("different buffer")
	}
}

func TestAddToObject(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	f := field.Object("obj", encoder.ObjectMarshalerFunc(func(dst *buffer.Buffer, enc encoder.ObjectEncoder) (*buffer.Buffer, error) {
		return enc.AddString(dst, "inner", "value"), nil
	}))
	got, err := f.AddTo(dst, recordingEncoder{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "obj={inner=value;};" {
		t.Fatalf("buffer=%q", got.String())
	}
}

func TestAddToArray(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	f := field.Array("items", encoder.ArrayMarshalerFunc(func(dst *buffer.Buffer, enc encoder.ArrayEncoder) (*buffer.Buffer, error) {
		return enc.AppendInt(dst, 7), nil
	}))
	got, err := f.AddTo(dst, recordingEncoder{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "items=[7;];" {
		t.Fatalf("buffer=%q", got.String())
	}
}

func TestAddToInline(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	f := field.Inline(encoder.ObjectMarshalerFunc(func(dst *buffer.Buffer, enc encoder.ObjectEncoder) (*buffer.Buffer, error) {
		return enc.AddString(dst, "inline", "value"), nil
	}))
	got, err := f.AddTo(dst, recordingEncoder{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "inline=value;" {
		t.Fatalf("buffer=%q", got.String())
	}
}

func TestAddToNamespace(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	got, err := field.Namespace("http").AddTo(dst, recordingEncoder{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "http." {
		t.Fatalf("buffer=%q", got.String())
	}
}

func TestAddToReflectStringerErrorAndTime(t *testing.T) {
	t.Parallel()

	loc := time.FixedZone("test", 3*60*60)
	when := time.Date(2026, 5, 7, 10, 30, 0, 123, loc)
	tests := []struct {
		name string
		f    field.Field
		want string
	}{
		{name: "reflect", f: field.Reflect("any", 42), want: "any=42;"},
		{name: "stringer", f: field.Stringer("stringer", testStringer("value")), want: "stringer=value;"},
		{name: "error", f: field.Error("error", errors.New("failed")), want: "error=failed;"},
		{name: "time", f: field.Time("time", when), want: "time=2026-05-07T07:30:00.000000123Z;"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.f.AddTo(buffer.New(0), recordingEncoder{})
			if err != nil {
				t.Fatalf("err=%v", err)
			}
			if got.String() != tt.want {
				t.Fatalf("buffer=%q, want %q", got.String(), tt.want)
			}
		})
	}
}

func TestAddToPropagatesObjectError(t *testing.T) {
	t.Parallel()

	want := errors.New("object failed")
	f := field.Object("obj", encoder.ObjectMarshalerFunc(func(dst *buffer.Buffer, enc encoder.ObjectEncoder) (*buffer.Buffer, error) { return dst, want }))
	if _, err := f.AddTo(buffer.New(0), recordingEncoder{}); !errors.Is(err, want) {
		t.Fatalf("err=%v", err)
	}
}

func TestAddToPropagatesArrayError(t *testing.T) {
	t.Parallel()

	want := errors.New("array failed")
	f := field.Array("arr", encoder.ArrayMarshalerFunc(func(dst *buffer.Buffer, enc encoder.ArrayEncoder) (*buffer.Buffer, error) {
		dst.AppendString("partial")
		return dst, want
	}))
	got, err := f.AddTo(buffer.New(0), recordingEncoder{})
	if !errors.Is(err, want) {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "arr=[partial" {
		t.Fatalf("buffer=%q", got.String())
	}
}

func TestAddToUnsupportedType(t *testing.T) {
	t.Parallel()

	_, err := field.Field{Type: field.Type(255)}.AddTo(buffer.New(0), recordingEncoder{})
	if !errors.Is(err, field.ErrUnsupportedType) {
		t.Fatalf("err=%v", err)
	}
}

func TestAddToStringerPropagatesPanic(t *testing.T) {
	t.Parallel()

	requirePanic(t, func() {
		_, _ = field.Stringer("stringer", panicStringer{}).AddTo(buffer.New(0), recordingEncoder{})
	})
}

func TestAddToErrorPropagatesPanic(t *testing.T) {
	t.Parallel()

	requirePanic(t, func() {
		_, _ = field.Error("error", panicError{}).AddTo(buffer.New(0), recordingEncoder{})
	})
}

func requirePanic(t *testing.T, fn func()) {
	t.Helper()

	defer func() {
		if recover() == nil {
			t.Fatal("function did not panic")
		}
	}()

	fn()
}
