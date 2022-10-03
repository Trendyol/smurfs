package plugin

import (
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/util"
	"os"
	"time"
)

type Receipt struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Executable  ExecutableArchive `yaml:"executable"`
	InstalledAt time.Time         `yaml:"installedAt"`
}

type ExecutableArchive struct {
	Version    string     `yaml:"version"`
	Executable Executable `yaml:",inline"`
	ArchiveURL string     `yaml:"archiveURL"`
	SHA256Sum  string     `yaml:"sha256Sum"`
}

func (r Receipt) Store(destinationPath string) error {
	yamlBytes, err := util.EncodeToYAML(r)
	if err != nil {
		return errors.Wrapf(err, "could not convert to yaml")
	}

	err = os.WriteFile(destinationPath, yamlBytes, 0o644)
	return errors.Wrapf(err, "could not write plugin receipt %q", destinationPath)
}

func LoadReceiptFrom(path string) (Receipt, error) {
	var receipt Receipt
	err := util.ReadYAMLFromFile(path, &receipt)
	return receipt, err
}
