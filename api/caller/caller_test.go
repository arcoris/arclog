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

package caller_test

import (
	"testing"

	"arcoris.dev/arclog/api/buffer"
	"arcoris.dev/arclog/api/caller"
)

func TestEncoderContract(t *testing.T) {
	t.Parallel()

	enc := caller.Encoder(func(dst *buffer.Buffer, c caller.Caller) *buffer.Buffer {
		if !c.Defined {
			return dst
		}
		dst.AppendString(c.File)
		dst.AppendByte(':')
		dst.AppendInt(c.Line)
		return dst
	})

	dst := buffer.New(0)
	got := enc(dst, caller.Caller{Defined: true, File: "logger.go", Line: 42})
	if got != dst {
		t.Fatalf("Encoder returned a different buffer")
	}
	if got.String() != "logger.go:42" {
		t.Fatalf("encoded caller = %q, want %q", got.String(), "logger.go:42")
	}
}

func TestEncoderContractUndefinedCaller(t *testing.T) {
	t.Parallel()

	enc := caller.Encoder(func(dst *buffer.Buffer, c caller.Caller) *buffer.Buffer {
		if !c.Defined {
			return dst
		}
		dst.AppendString(c.File)
		return dst
	})

	dst := buffer.New(0)
	got := enc(dst, caller.Caller{})
	if got != dst {
		t.Fatalf("Encoder returned a different buffer")
	}
	if got.Len() != 0 {
		t.Fatalf("undefined caller encoded %q, want empty buffer", got.String())
	}
}

func TestResolverContract(t *testing.T) {
	t.Parallel()

	resolver := testResolver{
		caller: caller.Caller{Defined: true, File: "entry.go", Line: 7, Function: "log"},
	}

	got := resolver.Caller(1)
	if !got.Defined {
		t.Fatal("Caller() returned undefined caller")
	}
	if got.File != "entry.go" || got.Line != 7 || got.Function != "log" {
		t.Fatalf("Caller() = %#v", got)
	}
	if resolver.skip != 1 {
		t.Fatalf("skip = %d, want 1", resolver.skip)
	}
}

// testResolver records the requested skip value and returns a fixed Caller.
type testResolver struct {
	caller caller.Caller
	skip   int
}

func (r *testResolver) Caller(skip int) caller.Caller {
	r.skip = skip
	return r.caller
}
