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

import (
	"go/parser"
	"go/token"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestBufferpoolImportBoundary(t *testing.T) {
	t.Parallel()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	dir := filepath.Dir(filename)
	allowedImports := map[string]struct{}{
		"arcoris.dev/arclog/api/buffer": {},
		"arcoris.dev/pool":              {},
	}

	matches, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		t.Fatalf("glob source files: %v", err)
	}

	for _, name := range matches {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(token.NewFileSet(), name, nil, parser.ImportsOnly)
		if err != nil {
			t.Fatalf("parse %s imports: %v", filepath.Base(name), err)
		}
		for _, spec := range file.Imports {
			path := strings.Trim(spec.Path.Value, "\"")
			if !strings.HasPrefix(path, "arcoris.dev/") {
				continue
			}
			if _, ok := allowedImports[path]; !ok {
				t.Fatalf("runtime/internal/bufferpool must stay independent of core, encoders, and writers; %s imports %q", filepath.Base(name), path)
			}
		}
	}
}
