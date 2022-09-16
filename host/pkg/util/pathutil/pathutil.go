package pathutil

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// IsSubPath checks if the extending path is an extension of the basePath, it will return the extending path
// elements. Both paths have to be absolute or have the same root directory. The remaining path elements
func IsSubPath(basePath, subPath string) (string, bool) {
	extendingPath, err := filepath.Rel(basePath, subPath)
	if err != nil {
		return "", false
	}
	if strings.HasPrefix(extendingPath, "..") {
		return "", false
	}
	return extendingPath, true
}

// ReplaceBase will return a replacement path with replacement as a base of the path instead of the old base. a/b/c, a, d -> d/b/c
func ReplaceBase(path, old, replacement string) (string, error) {
	extendingPath, ok := IsSubPath(old, path)
	if !ok {
		return "", errors.Errorf("can't replace %q in %q, it is not a subpath", old, path)
	}
	return filepath.Join(replacement, extendingPath), nil
}
