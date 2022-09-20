package util

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"sigs.k8s.io/yaml"
)

// ReadYAMLFromFile loads a file from the FS into the provided object.
func ReadYAMLFromFile(path string, as interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	err = DecodeYAML(f, &as)
	return errors.Wrapf(err, "failed to parse yaml file %q", path)
}

func EncodeToYAML(obj interface{}) ([]byte, error) {
	return yaml.Marshal(obj)
}

// DecodeYAML tries to decode file as YAML format
func DecodeYAML(r io.ReadCloser, as interface{}) error {
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, &as)
}

// RemoveSymLink removes a symlink reference if exists.
func RemoveSymLink(path string) error {
	file, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "failed to stat symlink %q", path)
	}

	if file.Mode()&os.ModeSymlink == 0 {
		return errors.Errorf("file %q is not a symlink (mode=%s)", path, file.Mode())
	}
	if err := os.Remove(path); err != nil {
		return errors.Wrapf(err, "failed to remove the symlink in %q", path)
	}
	return nil
}
