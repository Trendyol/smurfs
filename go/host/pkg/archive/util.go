package archive

import (
	"github.com/pkg/errors"
	"strings"
)

func suspiciousPath(path string) error {
	if strings.Contains(path, "..") {
		return errors.Errorf("refusing to unpack archive with suspicious entry %q", path)
	}

	if strings.HasPrefix(path, `/`) || strings.HasPrefix(path, `\`) {
		return errors.Errorf("refusing to unpack archive with absolute entry %q", path)
	}

	return nil
}
