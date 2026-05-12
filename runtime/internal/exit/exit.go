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

package exit

import (
	"os"
	"sync"
)

var (
	// stateMu protects exitFn installation and lookup.
	//
	// With copies exitFn while holding stateMu and then calls the copied function
	// after releasing the lock. This avoids holding the global state lock while
	// executing either os.Exit or a test recorder.
	stateMu sync.Mutex

	// exitFn is the process termination function currently used by With.
	//
	// In production it is os.Exit. Tests replace it with Recorder.record through
	// Stub or WithStub. All access must go through stateMu.
	exitFn = os.Exit
)

// With terminates the process with code.
//
// In production, With calls os.Exit through the currently installed exit
// function. Tests may replace that function by installing a Recorder with Stub
// or WithStub.
//
// With does not validate code, synchronize with log flushing, or translate
// panic/fatal policy. Exit-code and terminal-action policy belongs to the
// runtime component that calls With.
func With(code int) {
	stateMu.Lock()
	fn := exitFn
	stateMu.Unlock()

	fn(code)
}
