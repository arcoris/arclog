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

package hook_test

import (
	"context"
	"testing"

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/hook"
)

type preWriteHookContract struct{}

var _ hook.PreWriteHook = preWriteHookContract{}

func (preWriteHookContract) PreWrite(context.Context, core.Entry, []field.Field) (core.Entry, []field.Field, error) {
	return core.Entry{}, nil, nil
}

func TestPreWriteHookContract(t *testing.T) {
	t.Parallel()

	entry := core.Entry{Message: "before"}
	fields := []field.Field{field.String("k", "v")}

	gotEntry, gotFields, err := (preWriteHookContract{}).PreWrite(context.Background(), entry, fields)
	if err != nil {
		t.Fatalf("PreWrite() error = %v", err)
	}
	if !gotEntry.IsZero() {
		t.Fatalf("entry = %#v, want zero entry from contract test hook", gotEntry)
	}
	if gotFields != nil {
		t.Fatalf("fields = %#v, want nil", gotFields)
	}
}
