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

package level

import (
	"testing"

	"arcoris.dev/arclog/api/buffer"
)

// TestEncoderContract verifies the function shape expected by entry encoders.
func TestEncoderContract(t *testing.T) {
	t.Parallel()

	var enc Encoder = func(dst *buffer.Buffer, lvl Level) *buffer.Buffer {
		dst.AppendString(lvl.String())
		return dst
	}

	dst := buffer.New(0)
	got := enc(dst, Error)
	if got != dst {
		t.Fatalf("Encoder returned a different buffer")
	}
	if got.String() != "error" {
		t.Fatalf("encoded level = %q, want %q", got.String(), "error")
	}
}
