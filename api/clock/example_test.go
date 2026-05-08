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

package clock_test

import (
	"fmt"
	"time"

	"arcoris.dev/arclog/api/clock"
)

func ExampleFunc() {
	fixed := time.Date(2026, 5, 8, 10, 30, 0, 0, time.UTC)
	clk := clock.Func(func() time.Time {
		return fixed
	})

	fmt.Println(clk.Now().Format(time.RFC3339))

	// Output:
	// 2026-05-08T10:30:00Z
}
