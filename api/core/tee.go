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

package core

import (
	"arcoris.dev/arclog/api/field"
	"arcoris.dev/arclog/api/level"
)

// Tee returns a Core that duplicates entries to all non-nil cores.
//
// Nil cores are ignored. When no non-nil cores are supplied, Tee returns Noop.
// When exactly one non-nil core is supplied, Tee returns it unchanged.
//
// When multiple cores are supplied, Tee copies the selected Core values into an
// immutable snapshot. Mutating the caller's variadic source after construction
// does not change the resulting fan-out.
func Tee(cores ...Core) Core {
	count := 0
	var single Core
	for _, c := range cores {
		if c == nil {
			continue
		}
		count++
		single = c
	}

	switch count {
	case 0:
		return Noop()
	case 1:
		return single
	}

	out := make(teeCore, 0, count)
	for _, c := range cores {
		if c != nil {
			out = append(out, c)
		}
	}
	return out
}

// teeCore owns an immutable snapshot of downstream cores.
//
// The slice is constructed only by Tee. Methods assume the slice contains no
// nil cores and do not copy it on the write path.
type teeCore []Core

var _ Core = teeCore(nil)

// Enabled reports true when any underlying Core is enabled for lvl.
func (tc teeCore) Enabled(lvl level.Level) bool {
	for _, c := range tc {
		if c.Enabled(lvl) {
			return true
		}
	}

	return false
}

// With returns a Tee of the underlying cores after applying With to each core.
//
// The fields slice is borrowed and forwarded to each Core.With. Individual cores
// are responsible for copying fields if their returned Core retains them.
func (tc teeCore) With(fields []field.Field) Core {
	cores := make([]Core, 0, len(tc))
	for _, c := range tc {
		cores = append(cores, c.With(fields))
	}

	return Tee(cores...)
}

// Check forwards the check phase to every underlying Core in order.
func (tc teeCore) Check(entry Entry, ce *CheckedEntry) *CheckedEntry {
	for _, c := range tc {
		ce = c.Check(entry, ce)
	}

	return ce
}

// Write writes entry and fields to every underlying Core in order.
//
// All cores are attempted even if an earlier core returns an error. Non-nil
// errors are aggregated as WriteErrors.
func (tc teeCore) Write(entry Entry, fields []field.Field) error {
	var errs WriteErrors
	for _, c := range tc {
		errs = AppendError(errs, c.Write(entry, fields))
	}

	return errs.Err()
}

// Sync synchronizes every underlying Core in order.
//
// All cores are attempted even if an earlier core returns an error. Non-nil
// errors are aggregated as WriteErrors.
func (tc teeCore) Sync() error {
	var errs WriteErrors
	for _, c := range tc {
		errs = AppendError(errs, c.Sync())
	}

	return errs.Err()
}
