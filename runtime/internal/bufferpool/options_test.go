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

package bufferpool

import "testing"

func TestNormalizeOptionsUsesDefaults(t *testing.T) {
	t.Parallel()

	got := normalizeOptions(Options{})

	if got.InitialCapacity != DefaultInitialCapacity {
		t.Fatalf("InitialCapacity = %d, want %d", got.InitialCapacity, DefaultInitialCapacity)
	}
	if got.MaxRetainedCapacity != DefaultMaxRetainedCapacity {
		t.Fatalf("MaxRetainedCapacity = %d, want %d", got.MaxRetainedCapacity, DefaultMaxRetainedCapacity)
	}
}

func TestNormalizeOptionsUsesPositiveValues(t *testing.T) {
	t.Parallel()

	got := normalizeOptions(Options{
		InitialCapacity:     2048,
		MaxRetainedCapacity: 4096,
	})

	if got.InitialCapacity != 2048 {
		t.Fatalf("InitialCapacity = %d, want 2048", got.InitialCapacity)
	}
	if got.MaxRetainedCapacity != 4096 {
		t.Fatalf("MaxRetainedCapacity = %d, want 4096", got.MaxRetainedCapacity)
	}
}

func TestNormalizeOptionsRejectsNonPositiveValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		options Options
	}{
		{name: "zero", options: Options{InitialCapacity: 0, MaxRetainedCapacity: 0}},
		{name: "negative", options: Options{InitialCapacity: -1, MaxRetainedCapacity: -1}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := normalizeOptions(tt.options)
			if got.InitialCapacity != DefaultInitialCapacity {
				t.Fatalf("InitialCapacity = %d, want %d", got.InitialCapacity, DefaultInitialCapacity)
			}
			if got.MaxRetainedCapacity != DefaultMaxRetainedCapacity {
				t.Fatalf("MaxRetainedCapacity = %d, want %d", got.MaxRetainedCapacity, DefaultMaxRetainedCapacity)
			}
		})
	}
}

func TestNormalizeOptionsRaisesMaxRetainedCapacityToInitialCapacity(t *testing.T) {
	t.Parallel()

	got := normalizeOptions(Options{
		InitialCapacity:     4096,
		MaxRetainedCapacity: 1024,
	})

	if got.InitialCapacity != 4096 {
		t.Fatalf("InitialCapacity = %d, want 4096", got.InitialCapacity)
	}
	if got.MaxRetainedCapacity != 4096 {
		t.Fatalf("MaxRetainedCapacity = %d, want 4096", got.MaxRetainedCapacity)
	}
}

func TestNormalizePositive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    int
		fallback int
		want     int
	}{
		{name: "positive", value: 3, fallback: 9, want: 3},
		{name: "zero", value: 0, fallback: 9, want: 9},
		{name: "negative", value: -3, fallback: 9, want: 9},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := normalizePositive(tt.value, tt.fallback); got != tt.want {
				t.Fatalf("normalizePositive(%d, %d) = %d, want %d", tt.value, tt.fallback, got, tt.want)
			}
		})
	}
}
