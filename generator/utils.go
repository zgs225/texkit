package generator

import (
	"errors"
	"path/filepath"
	"strings"
)

// buildPath build a filesystem path from pkger path and o.OutputDir. It removes
// prefix "templates" of path
func buildPath(path string, o *Option) (string, error) {
	pieces := strings.Split(path, ":")

	if len(pieces) != 2 {
		return "", errors.New("invalid pkger path")
	}

	path = strings.TrimPrefix(pieces[1], "/templates")
	path = filepath.Join(o.OutputDir, path)

	if filepath.IsAbs(path) {
		return path, nil
	}

	return filepath.Abs(path)
}
