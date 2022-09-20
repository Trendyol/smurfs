package receipt

import (
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"github.com/trendyol/smurfs/go/host/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"

	"github.com/pkg/errors"
)

// Store saves the given receipt at the destination.
// The caller has to ensure that the destination directory exists.
func Store(receipt plugin.Receipt, dest string) error {
	yamlBytes, err := util.EncodeToYAML(receipt)
	if err != nil {
		return errors.Wrapf(err, "could not convert to yaml")
	}

	err = os.WriteFile(dest, yamlBytes, 0o644)
	return errors.Wrapf(err, "could not write plugin receipt %q", dest)
}

// Load reads the plugin receipt at the specified destination.
// If not found, it returns os.IsNotExist error.
func Load(path string) (plugin.Receipt, error) {
	var receipt plugin.Receipt
	err := util.ReadYAMLFromFile(path, &receipt)
	return receipt, err
}

// New returns a new receipt with the given plugin
func New(p plugin.Plugin, time metav1.Time) plugin.Receipt {
	return plugin.Receipt{
		Plugin:      p,
		InstalledAt: time.Time,
	}
}
