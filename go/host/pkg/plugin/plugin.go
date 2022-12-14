package plugin

import (
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/models"
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
	Distributions    []models.Distribution  `yaml:"distributions"`
	Flags            []struct {
		Name        string `yaml:"name"`
		Short       string `yaml:"short"`
		Description string `yaml:"description"`
		Repeated    bool   `yaml:"required"`
		Required    bool   `yaml:"required"`
	}
}

func (p Plugin) GetCompatibleDistribution() (models.Distribution, error) {
	platform := util.OSArch().String()
	for _, distribution := range p.Distributions {
		if slices.Contains(distribution.Targets, platform) {
			return distribution, nil
		}
	}
	return models.Distribution{}, errors.Wrapf(ErrPluginNotCompatible, "plugin %s is not compatible with the current platform %s", p.Name, platform)
}

func (p Plugin) GenerateReceipt(distribution models.Distribution) Receipt {
	return Receipt{
		Name:        p.Name,
		Description: p.ShortDescription,
		Executable: ExecutableArchive{
			Version:    distribution.Version,
			Executable: distribution.Executable,
		},
		InstalledAt: time.Now(),
	}
}
