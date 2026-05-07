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

type managerContract struct{}

var _ hook.Manager = managerContract{}

func (managerContract) AddPreWrite(hook.PreWriteHook, hook.Priority) hook.Registration {
	return hook.RegistrationFunc(func() bool { return true })
}

func (managerContract) AddPostWrite(hook.PostWriteHook, hook.Priority) hook.Registration {
	return hook.RegistrationFunc(func() bool { return true })
}

func (managerContract) AddError(hook.ErrorHook, hook.Priority) hook.Registration {
	return hook.RegistrationFunc(func() bool { return true })
}

func (managerContract) FirePreWrite(ctx context.Context, entry core.Entry, fields []field.Field) (core.Entry, []field.Field, error) {
	entry.Message = "checked"
	fields = append(fields, field.String("phase", "pre"))

	return entry, fields, nil
}

func (managerContract) FirePostWrite(context.Context, core.Entry, []field.Field, hook.WriteResult) error {
	return nil
}

func (managerContract) FireError(context.Context, core.Entry, []field.Field, error) {}

func (managerContract) Stop(context.Context) error {
	return nil
}

func TestManagerContract(t *testing.T) {
	t.Parallel()

	manager := managerContract{}

	registrations := []struct {
		name         string
		registration hook.Registration
	}{
		{name: "pre-write", registration: manager.AddPreWrite(nil, hook.PriorityDefault)},
		{name: "post-write", registration: manager.AddPostWrite(nil, hook.PriorityDefault)},
		{name: "error", registration: manager.AddError(nil, hook.PriorityDefault)},
	}

	for _, tt := range registrations {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if !tt.registration.Remove() {
				t.Fatal("registration should remove successfully")
			}
		})
	}
}

func TestManagerFirePreWriteReturnsAuthoritativeValues(t *testing.T) {
	t.Parallel()

	manager := managerContract{}
	entry := core.Entry{Message: "input"}
	fields := []field.Field{field.String("k", "v")}

	gotEntry, gotFields, err := manager.FirePreWrite(context.Background(), entry, fields)
	if err != nil {
		t.Fatalf("FirePreWrite() error = %v", err)
	}
	if gotEntry.Message != "checked" {
		t.Fatalf("Message = %q, want checked", gotEntry.Message)
	}
	if len(gotFields) != 2 {
		t.Fatalf("len(fields) = %d, want 2", len(gotFields))
	}
}

func TestManagerFirePostWriteAndErrorAndStop(t *testing.T) {
	t.Parallel()

	manager := managerContract{}

	if err := manager.FirePostWrite(context.Background(), core.Entry{}, nil, hook.Success()); err != nil {
		t.Fatalf("FirePostWrite() error = %v", err)
	}

	manager.FireError(context.Background(), core.Entry{}, nil, nil)

	if err := manager.Stop(context.Background()); err != nil {
		t.Fatalf("Stop() error = %v", err)
	}
}
