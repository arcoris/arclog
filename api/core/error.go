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

import "strings"

// WriteErrors is an aggregate of errors returned by multiple cores.
//
// It implements Unwrap() []error so errors.Is and errors.As can inspect the
// individual causes. A zero WriteErrors value represents no error.
type WriteErrors []error

// AppendError appends err to errs when err is non-nil.
//
// Keeping nil errors out of the common aggregation path lets CheckedEntry and
// teeCore avoid extra filtering work when all callers use AppendError.
func AppendError(errs WriteErrors, err error) WriteErrors {
	if err == nil {
		return errs
	}

	return append(errs, err)
}

// Err returns nil when errs contains no non-nil errors.
//
// Otherwise Err returns a compact WriteErrors value. If errs already contains
// only non-nil errors, Err returns errs without allocating.
func (errs WriteErrors) Err() error {
	if len(errs) == 0 {
		return nil
	}

	needsCompact := false
	for _, err := range errs {
		if err == nil {
			needsCompact = true
			break
		}
	}

	if !needsCompact {
		return errs
	}

	out := make(WriteErrors, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			out = append(out, err)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

// Error returns a stable diagnostic representation of errs.
//
// Nil errors are ignored. The exact string is intended for diagnostics only; use
// errors.Is or errors.As against Err or Unwrap for programmatic checks.
func (errs WriteErrors) Error() string {
	errs = compactErrors(errs)

	switch len(errs) {
	case 0:
		return ""
	case 1:
		return errs[0].Error()
	}

	var b strings.Builder
	b.WriteString("multiple core errors:")
	for _, err := range errs {
		if err == nil {
			continue
		}
		b.WriteString(" ")
		b.WriteString(err.Error())
		b.WriteByte(';')
	}

	return b.String()
}

// Unwrap returns the non-nil errors contained in errs.
//
// The returned slice is caller-owned and may be modified by the caller.
func (errs WriteErrors) Unwrap() []error {
	compact := compactErrors(errs)
	out := make([]error, len(compact))
	copy(out, compact)
	return out
}

// compactErrors returns a WriteErrors value containing only non-nil errors.
//
// When errs is already compact, compactErrors returns errs without allocating.
func compactErrors(errs WriteErrors) WriteErrors {
	if len(errs) == 0 {
		return nil
	}

	needsCompact := false
	for _, err := range errs {
		if err == nil {
			needsCompact = true
			break
		}
	}
	if !needsCompact {
		return errs
	}

	out := make(WriteErrors, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			out = append(out, err)
		}
	}
	return out
}
