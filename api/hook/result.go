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

package hook

// WriteResult describes the write outcome observed by PostWriteHook.
//
// The API-level result intentionally contains only the write error because
// core.Core.Write returns only error. Runtime implementations may track byte
// counts, sink names, retry counts, durations, or routing diagnostics internally,
// but those details are not part of the stable hook contract.
//
// The zero value represents a successful write attempt.
type WriteResult struct {
	// Err is the error returned by the write attempt.
	//
	// Err is borrowed from the runtime pipeline. It may wrap lower-level sink or
	// encoder errors, but this package does not define error classes or retry
	// policy.
	Err error
}

// Success returns the zero successful write result.
func Success() WriteResult {
	return WriteResult{}
}

// Failure returns a write result containing err.
//
// Passing nil records no error, so Failure(nil).Failed() reports false. The
// function does not synthesize sentinel errors because the hook API should
// reflect the concrete error returned by the write path.
func Failure(err error) WriteResult {
	return WriteResult{Err: err}
}

// Failed reports whether r contains a non-nil write error.
func (r WriteResult) Failed() bool {
	return r.Err != nil
}
