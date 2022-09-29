package plugin

import (
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/util"
	"k8s.io/utils/strings/slices"
	"time"
)

var (
	ErrPluginNotCompatible = errors.New("plugin is not compatible with the current platform")
)

type Plugin struct {
	Name             string                 `yaml:"name"`
	ShortDescription string                 `yaml:"shortDescription"`
	LongDescription  string                 `yaml:"longDescription"`
	Usage            string                 `yaml:"usage"`
	Source           map[string]interface{} `yaml:"source"`
	Distributions    []Distribution         `yaml:"distributions"`
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

func (p Plugin) GenerateReceipt(distribution Distribution) Receipt {
	return Receipt{
		Name:        p.Name,
		Description: p.ShortDescription,
		Executable: ExecutableArchive{
			Executable: distribution.Executable,
		},
		InstalledAt: time.Now(),
	}
}

type Distribution struct {
	Targets    []string   `yaml:"targets"`
	Version    string     `yaml:"version"`
	Executable Executable `yaml:"executable"`
}

type Executable struct {
	// Provider specifies how the Address will be used. (Required)
	Provider ExecutableProvider `yaml:"provider"`

	// Address is the location of the executable archive. (Required)
	Address string `yaml:"address"`

	// SHA256 is the SHA256 checksum of the executable archive. (Optional)
	//
	// Required for the URI provider.
	SHA256 string `yaml:"sha256"`

	// Entrypoint specifies which file will be executed in the archive. (Required)
	Entrypoint string `yaml:"entrypoint"`
}

type ExecutableProvider string

const (
	GitHubExecutableProvider = ExecutableProvider("GitHub")
	GitLabExecutableProvider = ExecutableProvider("GitLab")
	URIExecutableProvider    = ExecutableProvider("URI")
	LocalExecutableProvider  = ExecutableProvider("Local")
)
