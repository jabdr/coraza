package io

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// FSReadFile fixes slashes in embedfs on windows
func FSReadFile(fsys fs.FS, name string) ([]byte, error) {
	filename := name
	if filepath.Separator == '\\' {
		filename = strings.ReplaceAll(filename, "\\", "/")
	}

	return fs.ReadFile(fsys, filename)
}
