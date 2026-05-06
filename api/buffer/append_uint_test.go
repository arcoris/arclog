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

import "testing"

func TestBufferAppendUints(t *testing.T) {
	buf := newBuffer()

	buf.AppendUint(1)
	buf.AppendByte(',')
	buf.AppendUint64(2)
	buf.AppendByte(',')
	buf.AppendUint32(3)
	buf.AppendByte(',')
	buf.AppendUint16(4)
	buf.AppendByte(',')
	buf.AppendUint8(5)

	if got, want := string(buf.Bytes()), "1,2,3,4,5"; got != want {
		t.Fatalf("AppendUint* = %q, want %q", got, want)
	}
}

func TestBufferAppendUintptr(t *testing.T) {
	buf := newBuffer()

	buf.AppendUintptr(0x1234)

	if got, want := string(buf.Bytes()), "0x1234"; got != want {
		t.Fatalf("AppendUintptr = %q, want %q", got, want)
	}
}
