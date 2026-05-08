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
	"fmt"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/level"
)

func ExampleNoop() {
	c := core.Noop()
	entry := core.Entry{Level: level.Info, Message: "ready"}

	checked := c.Check(entry, nil)
	err := c.Write(entry, nil)

	fmt.Println(checked == nil)
	fmt.Println(err)

	// Output:
	// true
	// <nil>
}
