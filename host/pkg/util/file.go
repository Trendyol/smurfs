package util

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"sigs.k8s.io/yaml"
)

// ReadFromFile loads a file from the FS into the provided object.
func ReadFromFile(path string, as interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	err = DecodeYAML(f, &as)
	return errors.Wrapf(err, "failed to parse yaml file %q", path)
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
