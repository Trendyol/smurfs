package plugin

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/models"
	"github.com/trendyol/smurfs/go/host/pkg/util"
	"os"
	"path/filepath"
	"time"
)

type Receipt struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Executable  ExecutableArchive `yaml:"executable"`
	InstalledAt time.Time         `yaml:"installedAt"`
}

type ExecutableArchive struct {
	Version    string            `yaml:"version"`
	Executable models.Executable `yaml:",inline"`
	ArchiveURL string            `yaml:"archiveURL"`
	SHA256Sum  string            `yaml:"sha256Sum"`
}

func (r Receipt) Store(pluginPath, destinationPath string) (Receipt, error) {

	r.Executable.Executable.Entrypoint = fmt.Sprintf("%s/%s/%s/%s/%s", pluginPath, "store", r.Name, r.Executable.Version, r.Name)
	yamlBytes, err := util.EncodeToYAML(r)
	if err != nil {
		return r, errors.Wrapf(err, "could not convert to yaml")
	}

	err = os.MkdirAll(filepath.Dir(destinationPath), 0755)
	if err != nil {
		return r, err
	}

	err = os.WriteFile(destinationPath, yamlBytes, 0o644)
	return r, errors.Wrapf(err, "could not write plugin receipt %q", destinationPath)
}

func LoadReceiptFrom(path string) (Receipt, error) {
	var receipt Receipt
	err := util.ReadYAMLFromFile(path, &receipt)
	return receipt, err
}
