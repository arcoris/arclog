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
	"fmt"

	"arcoris.dev/arclog/api/field"
)

func ExampleString() {
	f := field.String("service", "api")

	fmt.Println(f.Key)
	fmt.Println(f.Type)
	fmt.Println(f.String)

	// Output:
	// service
	// StringType
	// api
}

func ExampleDict() {
	f := field.Dict("http",
		field.Int("status", 200),
		field.String("method", "GET"),
	)

	fmt.Println(f.Key)
	fmt.Println(f.Type)

	// Output:
	// http
	// ObjectMarshalerType
}
