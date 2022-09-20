package plugin

import (
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/util"
	"k8s.io/utils/strings/slices"
	"os"
	"time"
)

var (
	ErrPluginNotCompatible = errors.New("plugin is not compatible with the current platform")
)

type Plugin struct {
	Name          string         `yaml:"name"`
	Description   string         `yaml:"description"`
	Distributions []Distribution `yaml:"distributions"`
}

func (p Plugin) GetCompatibleDistribution() (Distribution, error) {
	platform := util.OSArch().String()
	for _, distribution := range p.Distributions {
		if slices.Contains(distribution.Targets, platform) {
			return distribution, nil
		}
	}
	return Distribution{}, errors.Wrapf(ErrPluginNotCompatible, "plugin %s is not compatible with the current platform %s", p.Name, platform)
}

func (p Plugin) GenerateReceipt() Receipt {
	return Receipt{
		Plugin:      p,
		InstalledAt: time.Now(),
	}
}

type Distribution struct {
	Targets    []string   `yaml:"targets"`
	Version    string     `yaml:"version"`
	Executable Executable `yaml:"executable"`
}

type Executable struct {
	// Provider specifies how the Address will be used
	Provider ExecutableProvider `yaml:"provider"`

	// Address is the location of the executable archive
	Address string `yaml:"uri"`

	// Entrypoint specifies which file will be executed in the archive
	Entrypoint string `yaml:"entrypoint"`
}

type ExecutableProvider string

const (
	GitHubExecutableProvider = ExecutableProvider("github")
	GitLabExecutableProvider = ExecutableProvider("gitlab")
	URIExecutableProvider    = ExecutableProvider("uri")
)

type Receipt struct {
	Plugin      `yaml:",inline" json:",inline"`
	InstalledAt time.Time `json:"installedAt" yaml:"installedAt"`
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

//// Plugin describes a plugin manifest file.
//type Plugin struct {
//	metav1.TypeMeta   `json:",inline" yaml:",inline"`
//	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata"`
//
//	Spec Spec `json:"spec"`
//}
//
//type Spec struct {
//	Name             string   `json:"name"`
//	Version          string   `json:"version"`
//	Description      string   `json:"description"`
//	ShortDescription string   `json:"shortDescription"`
//	Runnable         Runnable `json:"runnable"`
//}
//
//type Runnable struct {
//	// Archive address of the plugin
//	URI string `json:"uri" yaml:"uri"`
//
//	// Sha256 of the plugin to check integrity
//	Sha256 string          `json:"sha256" yaml:"sha256"`
//	Files  []FileOperation `json:"files"  yaml:"files"`
//
//	// Entrypoint is the path to the binary in the archive
//	Entrypoint string `json:"name" yaml:"name"`
//
//	Bin string `json:"bin" yaml:"bin"`
//}
//
//// FileOperation todo: do not use it
//type FileOperation struct {
//	From string `json:"from"`
//	To   string `json:"to"`
//}
//
//// Receipt is a record of a plugin installation.
//type Receipt struct {
//	Plugin      `yaml:",inline" json:",inline"`
//	InstalledAt time.Time `json:"installedAt" yaml:"installedAt"`
//}
