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

	"arcoris.dev/arclog/api/core"
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/hook"
)

var (
	_ hook.PreWriteHook  = hook.PreWriteFunc(nil)
	_ hook.PostWriteHook = hook.PostWriteFunc(nil)
	_ hook.ErrorHook     = hook.ErrorFunc(nil)
)

type namedHook struct{}

func (namedHook) Name() string { return "named" }

var _ hook.Named = namedHook{}

type preWriteValueHook struct{}

func (preWriteValueHook) PreWrite(context.Context, core.Entry, []field.Field) (core.Entry, []field.Field, error) {
	return core.Entry{}, nil, nil
}

var _ hook.PreWriteHook = preWriteValueHook{}
