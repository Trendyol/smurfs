package receipt

import (
	"github.com/trendyol/smurfs/host/pkg/plugin"
	"github.com/trendyol/smurfs/host/pkg/util"
	"os"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Store saves the given receipt at the destination.
// The caller has to ensure that the destination directory exists.
func Store(receipt plugin.Receipt, dest string) error {
	yamlBytes, err := yaml.Marshal(receipt)
	if err != nil {
		return errors.Wrapf(err, "convert to yaml")
	}

	err = os.WriteFile(dest, yamlBytes, 0o644)
	return errors.Wrapf(err, "write plugin receipt %q", dest)
}

// Load reads the plugin receipt at the specified destination.
// If not found, it returns os.IsNotExist error.
func Load(path string) (plugin.Receipt, error) {
	var receipt plugin.Receipt
	err := util.ReadFromFile(path, &receipt)
	return receipt, err
}

// New returns a new receipt with the given plugin
func New(p plugin.Plugin, timestamp metav1.Time) plugin.Receipt {
	p.CreationTimestamp = timestamp
	return plugin.Receipt{
		Plugin: p,
	}
}
