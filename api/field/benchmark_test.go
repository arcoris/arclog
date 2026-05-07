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
	"testing"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/field"
)

func BenchmarkStringConstructor(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		_ = field.String("name", "arcoris")
	}
}

func BenchmarkAddToString(b *testing.B) {
	f := field.String("name", "arcoris")
	enc := recordingEncoder{}
	dst := buffer.New(0)

	b.ReportAllocs()

	for b.Loop() {
		dst.Reset()
		_, _ = f.AddTo(dst, enc)
	}
}
