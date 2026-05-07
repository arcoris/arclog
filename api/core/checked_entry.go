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
	"errors"

	"arcoris.dev/arclog/api/field"
)

var (
	// ErrCheckedEntryWritten is returned when Write is called more than once on
	// the same CheckedEntry.
	//
	// CheckedEntry is a one-shot write handle. Retrying the same value risks
	// duplicating output in cores that already accepted the first write.
	ErrCheckedEntryWritten = errors.New("core: checked entry already written")
)

// CheckedEntry is an Entry together with the Core values that agreed to write it
// during the check phase.
//
// A nil *CheckedEntry represents a disabled entry and is safe to write as a
// no-op. Core.Check implementations usually call AddCore on the incoming
// CheckedEntry when they decide to write the entry.
//
// CheckedEntry is not safe for concurrent mutation. Runtime loggers should build
// and write a checked entry within one log attempt.
type CheckedEntry struct {
	entry   Entry
	cores   []Core
	written bool
}

// AddCore adds c to ce and returns the resulting CheckedEntry.
//
// AddCore is safe to call on a nil receiver. If c is nil, AddCore returns ce
// unchanged. The first successful AddCore call fixes the Entry stored in ce;
// subsequent calls add more cores for the same entry and leave the Entry
// unchanged.
//
// AddCore does not clone entry. Entry is a value type; fields are supplied later
// to Write.
func (ce *CheckedEntry) AddCore(entry Entry, c Core) *CheckedEntry {
	if c == nil {
		return ce
	}

	if ce == nil {
		ce = &CheckedEntry{entry: entry}
	}

	ce.cores = append(ce.cores, c)
	return ce
}

// Entry returns the entry stored in ce.
//
// For a nil CheckedEntry, Entry returns the zero Entry value.
func (ce *CheckedEntry) Entry() Entry {
	if ce == nil {
		return Entry{}
	}

	return ce.entry
}

// Len reports the number of cores registered on ce.
func (ce *CheckedEntry) Len() int {
	if ce == nil {
		return 0
	}

	return len(ce.cores)
}

// IsEmpty reports whether ce has no registered cores.
func (ce *CheckedEntry) IsEmpty() bool {
	return ce.Len() == 0
}

// Write writes fields to every Core registered on ce.
//
// Write on a nil CheckedEntry is a no-op. Write returns all Core.Write errors as
// WriteErrors. Calling Write more than once on the same CheckedEntry returns
// ErrCheckedEntryWritten.
//
// The fields slice is borrowed and passed to each selected Core in order.
// Implementations must treat it as read-only unless they own the complete
// pipeline and document a stronger contract.
func (ce *CheckedEntry) Write(fields ...field.Field) error {
	if ce == nil {
		return nil
	}

	if ce.written {
		return ErrCheckedEntryWritten
	}
	ce.written = true

	var errs WriteErrors
	for _, c := range ce.cores {
		if c == nil {
			continue
		}

		errs = AppendError(errs, c.Write(ce.entry, fields))
	}

	return errs.Err()
}
