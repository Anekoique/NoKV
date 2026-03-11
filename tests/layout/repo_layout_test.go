package layout_test

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func repoRoot(t *testing.T) string {
	t.Helper()

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "../.."))
}

func TestRepoRootContainsOnlyHighSignalDomains(t *testing.T) {
	root := repoRoot(t)

	entries, err := os.ReadDir(root)
	if err != nil {
		t.Fatalf("read repo root: %v", err)
	}

	requiredDirs := map[string]bool{
		"cluster":       false,
		"cmd":           false,
		"docs":          false,
		"engine":        false,
		"observability": false,
		"pb":            false,
		"pd":            false,
		"storage":       false,
		"tests":         false,
		"tools":         false,
		"utils":         false,
	}

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".go") {
			t.Fatalf("repo root should not contain Go source files, found %s", name)
		}
		if entry.IsDir() {
			if _, ok := requiredDirs[name]; ok {
				requiredDirs[name] = true
			}
		}
	}

	for dir, found := range requiredDirs {
		if !found {
			t.Fatalf("expected top-level directory %q to exist", dir)
		}
	}
}
