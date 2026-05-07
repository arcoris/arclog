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

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/field"
)

func TestFieldsMarshalLogObject(t *testing.T) {
	t.Parallel()

	dst := buffer.New(0)
	enc := recordingEncoder{}
	fs := field.Fields{field.String("name", "arcoris"), field.Int("n", 2)}
	got, err := fs.MarshalLogObject(dst, enc)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "name=arcoris;n=2;" {
		t.Fatalf("buffer=%q", got.String())
	}
}

func TestFieldsMarshalLogObjectStopsAtFirstError(t *testing.T) {
	t.Parallel()

	fs := field.Fields{
		field.String("name", "arcoris"),
		{Type: field.Type(255)},
		field.String("after", "ignored"),
	}
	got, err := fs.MarshalLogObject(buffer.New(0), recordingEncoder{})
	if !errors.Is(err, field.ErrUnsupportedType) {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "name=arcoris;" {
		t.Fatalf("buffer=%q", got.String())
	}
}

func TestDict(t *testing.T) {
	t.Parallel()

	f := field.Dict("obj", field.String("name", "arcoris"))
	if f.Type != field.ObjectMarshalerType {
		t.Fatalf("Type=%s", f.Type)
	}
}

func TestDictObject(t *testing.T) {
	t.Parallel()

	got, err := field.DictObject(field.String("name", "arcoris")).MarshalLogObject(buffer.New(0), recordingEncoder{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got.String() != "name=arcoris;" {
		t.Fatalf("buffer=%q", got.String())
	}
}
