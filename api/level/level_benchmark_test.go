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

import "testing"

var (
	benchLevelBool   bool
	benchLevelInt    int
	benchLevelString string
	benchLevelValue  Level
)

func BenchmarkLevelEnabled(b *testing.B) {
	for b.Loop() {
		benchLevelBool = Error.Enabled(Info)
	}
}

func BenchmarkSeverityNumber(b *testing.B) {
	for b.Loop() {
		benchLevelInt = Error.SeverityNumber()
	}
}

func BenchmarkSeverityTextNamed(b *testing.B) {
	for b.Loop() {
		benchLevelString = Error.SeverityText()
	}
}

func BenchmarkStringNamed(b *testing.B) {
	for b.Loop() {
		benchLevelString = Error.String()
	}
}

func BenchmarkParseNamed(b *testing.B) {
	for b.Loop() {
		var err error
		benchLevelValue, err = Parse("error")
		if err != nil {
			b.Fatal(err)
		}
	}
}
