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

package importtest

import (
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	rootImport     = "arcoris.dev/arclog"
	apiFieldImport = "arcoris.dev/arclog/api/field"
)

func TestAPIPackagesDoNotImportRuntimeBoundaries(t *testing.T) {
	t.Parallel()

	apiDir := filepath.Join(repoRoot(t), "api")
	err := filepath.WalkDir(apiDir, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		checkFileImports(t, apiDir, path)
		return nil
	})
	if err != nil {
		t.Fatalf("walk api imports: %v", err)
	}
}

func checkFileImports(t *testing.T, apiDir, path string) {
	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
	if err != nil {
		t.Fatalf("parse imports in %s: %v", path, err)
	}

	for _, imported := range file.Imports {
		importPath := strings.Trim(imported.Path.Value, `"`)
		if isForbiddenAPIImport(apiDir, path, importPath) {
			t.Fatalf("%s imports forbidden API boundary %q", path, importPath)
		}
	}
}

func isForbiddenAPIImport(apiDir, path, importPath string) bool {
	if importPath == rootImport {
		return true
	}
	if strings.HasPrefix(importPath, rootImport+"/runtime/") {
		return true
	}
	if strings.HasPrefix(importPath, rootImport+"/internal/") {
		return true
	}
	return isEncoderPackage(apiDir, path) &&
		(importPath == apiFieldImport || strings.HasPrefix(importPath, apiFieldImport+"/"))
}

func isEncoderPackage(apiDir, path string) bool {
	rel, err := filepath.Rel(apiDir, path)
	if err != nil {
		return false
	}
	return rel == "encoder" || strings.HasPrefix(rel, "encoder"+string(filepath.Separator))
}

func repoRoot(t *testing.T) string {
	t.Helper()

	dir, err := filepath.Abs(".")
	if err != nil {
		t.Fatalf("resolve working directory: %v", err)
	}

	for {
		if exists(filepath.Join(dir, "go.mod")) {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find repository root")
		}
		dir = parent
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
