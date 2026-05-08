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

package predicate_test

import (
	"fmt"
	"strings"

	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
	"arcoris.dev/arclog/api/predicate"
)

func ExampleAnd() {
	atLeastWarn := predicate.Func(func(entry predicate.Entry, _ []field.Field) bool {
		return entry.Level.Enabled(level.Warn)
	})
	hasMessage := predicate.Func(func(entry predicate.Entry, _ []field.Field) bool {
		return entry.Message != ""
	})

	p := predicate.And(atLeastWarn, hasMessage)

	fmt.Println(p.ShouldLog(predicate.Entry{Level: level.Error, Message: "failed"}, nil))
	fmt.Println(p.ShouldLog(predicate.Entry{Level: level.Debug, Message: "debug"}, nil))

	// Output:
	// true
	// false
}

func ExampleOr() {
	isStartup := predicate.Func(func(entry predicate.Entry, _ []field.Field) bool {
		return strings.Contains(entry.Message, "startup")
	})
	isError := predicate.Func(func(entry predicate.Entry, _ []field.Field) bool {
		return entry.Level.Enabled(level.Error)
	})

	p := predicate.Or(isStartup, isError)

	fmt.Println(p.ShouldLog(predicate.Entry{Level: level.Info, Message: "startup complete"}, nil))
	fmt.Println(p.ShouldLog(predicate.Entry{Level: level.Warn, Message: "retrying"}, nil))

	// Output:
	// true
	// false
}
