package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	tcs := []struct {
		path     string
		result   bool
		hasError bool
	}{
		{".", true, false},
		{"file_test.go", true, false},
		{"file_test.gogo", false, false},
	}

	for _, tc := range tcs {
		result, err := Exists(tc.path)
		ast.Equal(tc.result, result)
		ast.Equal(tc.hasError, err != nil)
	}
}

func TestIsDir(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	tcs := []struct {
		path     string
		result   bool
		hasError bool
	}{
		{".", true, false},
		{"file_test.go", false, false},
		{"file_test.gogo", false, true},
	}

	for _, tc := range tcs {
		result, err := IsDir(tc.path)
		ast.Equal(tc.result, result)
		ast.Equal(tc.hasError, err != nil)
	}
}

func TestMakeDir(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	tcs := []struct {
		path     string
		perm     os.FileMode
		result   bool
		hasError bool
	}{
		{"temp", 0644, true, false},
		{"temp", 0644, false, false},
		{"temp/test/1", 0644, true, false},
	}

	for _, tc := range tcs {
		result, err := MakeDir(tc.path, tc.perm)
		ast.Equal(tc.result, result)
		ast.Equal(tc.hasError, err != nil)
	}

	_ = os.RemoveAll("temp")
}

func TestListFiles(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	tcs := []struct {
		dir         string
		filePattern string
		size        int
		hasError    bool
	}{
		{".", "", 2, false},
		{".", ".*test.*", 1, false},
		{".", ".*no.*", 0, false},
		{"no", "", 0, true},
	}

	for _, tc := range tcs {
		fps, err := ListFiles(tc.dir, tc.filePattern)
		ast.Equal(tc.size, len(fps))
		ast.Equal(tc.hasError, err != nil)
	}
}
