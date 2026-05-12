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

package exit_test

import (
	"fmt"

	"arcoris.dev/arclog/runtime/internal/exit"
)

func ExampleWithStub() {
	recorder := exit.WithStub(func() {
		exit.With(exit.Failure)
	})

	fmt.Println(recorder.Exited())
	fmt.Println(recorder.Code())
	fmt.Println(recorder.Calls())

	// Output:
	// true
	// 1
	// 1
}

func ExampleStub() {
	recorder := exit.Stub()
	defer recorder.Unstub()

	exit.With(exit.Failure)

	fmt.Println(recorder.Exited())
	fmt.Println(recorder.Code())

	// Output:
	// true
	// 1
}

func ExampleWith() {
	recorder := exit.WithStub(func() {
		exit.With(exit.Failure)
	})

	fmt.Println(recorder.Exited())

	// Output:
	// true
}
