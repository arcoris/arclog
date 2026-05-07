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
	"testing"

	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
	"arcoris.dev/arclog/api/predicate"
)

func TestAlwaysAndNever(t *testing.T) {
	t.Parallel()

	entry := predicate.Entry{Level: level.Error, LoggerName: "api"}
	fields := []field.Field{field.String("service", "auth")}

	tests := []struct {
		name string
		p    predicate.Predicate
		want bool
	}{
		{name: "always", p: predicate.Always(), want: true},
		{name: "never", p: predicate.Never(), want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.p.ShouldLog(entry, fields); got != tt.want {
				t.Fatalf("ShouldLog() = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestConstantEvaluationDoesNotAllocate(t *testing.T) {
	entry := predicate.Entry{Level: level.Info, LoggerName: "api"}
	fields := []field.Field{field.String("service", "auth")}

	tests := []struct {
		name string
		p    predicate.Predicate
	}{
		{name: "always", p: predicate.Always()},
		{name: "never", p: predicate.Never()},
	}

	for _, tt := range tests {
		allocs := testing.AllocsPerRun(1000, func() {
			_ = tt.p.ShouldLog(entry, fields)
		})
		if allocs != 0 {
			t.Fatalf("%s allocs per evaluation = %g, want 0", tt.name, allocs)
		}
	}
}
