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

package writer_test

import (
	"errors"
	"testing"

	"arcoris.dev/arclog/api/writer"
)

func TestSyncFuncCallsFunction(t *testing.T) {
	t.Parallel()

	called := false
	syncer := writer.SyncFunc(func() error {
		called = true
		return nil
	})

	if err := syncer.Sync(); err != nil {
		t.Fatalf("Sync() error = %v", err)
	}
	if !called {
		t.Fatal("Sync() did not call the adapter function")
	}
}

func TestSyncFuncReturnsError(t *testing.T) {
	t.Parallel()

	want := errors.New("flush failed")
	syncer := writer.SyncFunc(func() error {
		return want
	})

	if err := syncer.Sync(); !errors.Is(err, want) {
		t.Fatalf("Sync() error = %v, want %v", err, want)
	}
}

func TestNilSyncFuncPanics(t *testing.T) {
	t.Parallel()

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("Sync() did not panic for nil SyncFunc")
		}
	}()

	var syncer writer.SyncFunc
	_ = syncer.Sync()
}
