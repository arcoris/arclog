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

package level_test

import (
	"fmt"

	"arcoris.dev/arclog/api/level"
)

func ExampleParse() {
	lvl, err := level.Parse("warning")

	fmt.Println(lvl)
	fmt.Println(err == nil)

	// Output:
	// warn
	// true
}

func ExampleLevel_Enabled() {
	threshold := level.Warn

	fmt.Println(level.Error.Enabled(threshold))
	fmt.Println(level.Debug.Enabled(threshold))
	fmt.Println(level.Invalid.Enabled(threshold))

	// Output:
	// true
	// false
	// false
}
