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

package buffer_test

import (
	"testing"

	"arcoris.dev/arclog/api/buffer"
)

func newBuffer() *buffer.Buffer {
	return &buffer.Buffer{}
}

func TestBufferWriteMethods(t *testing.T) {
	buf := newBuffer()

	n, err := buf.Write([]byte("foo"))
	if err != nil {
		t.Fatalf("Write returned error: %v", err)
	}
	if n != len("foo") {
		t.Fatalf("Write returned n = %d, want %d", n, len("foo"))
	}

	n, err = buf.WriteString("bar")
	if err != nil {
		t.Fatalf("WriteString returned error: %v", err)
	}
	if n != len("bar") {
		t.Fatalf("WriteString returned n = %d, want %d", n, len("bar"))
	}

	if err := buf.WriteByte('!'); err != nil {
		t.Fatalf("WriteByte returned error: %v", err)
	}

	if got, want := string(buf.Bytes()), "foobar!"; got != want {
		t.Fatalf("buffer contents = %q, want %q", got, want)
	}
}

func TestBufferAppendBytes(t *testing.T) {
	buf := newBuffer()

	buf.AppendByte('[')
	buf.AppendBytes([]byte("foo"))
	buf.AppendByte(',')
	buf.AppendString("bar")
	buf.AppendByte(']')

	if got, want := string(buf.Bytes()), "[foo,bar]"; got != want {
		t.Fatalf("buffer contents = %q, want %q", got, want)
	}
}
