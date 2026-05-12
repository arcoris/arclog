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

package level

import (
	"errors"
	"fmt"
)

// MarshalText implements encoding.TextMarshaler.
//
// Valid severities and Off are serialized with String. Out-of-range values are
// rejected because they are not stable configuration or wire values.
func (l Level) MarshalText() ([]byte, error) {
	if !l.IsThreshold() {
		return nil, fmt.Errorf("level.MarshalText: invalid level %d", int(l))
	}
	return []byte(l.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
//
// UnmarshalText accepts the same names and aliases as Parse. On parse failure
// it leaves the receiver unchanged. A nil receiver returns an explicit error.
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errors.New("level.UnmarshalText: nil receiver")
	}

	parsed, err := Parse(string(text))
	if err != nil {
		return err
	}
	*l = parsed
	return nil
}
