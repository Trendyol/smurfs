package environment

import (
	"github.com/trendyol/smurfs/go/host/pkg/consts"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Paths contains all important environment paths
type Paths struct {
	base string
	tmp  string
}

func NewPaths(base string) Paths {
	return Paths{base: base, tmp: os.TempDir()}
}

// BasePath returns the base directory.
func (p Paths) BasePath() string { return p.base }

// InstallReceiptsPath returns the base directory where plugin receipts are stored.
//
// e.g. {BasePath}/receipts
func (p Paths) InstallReceiptsPath() string { return filepath.Join(p.base, "receipts") }

// BinPath returns the path where plugin executable symbolic links are found.
// This path should be added to $PATH in client machine.
//
// e.g. {BasePath}/bin
func (p Paths) BinPath() string { return filepath.Join(p.base, "bin") }

// InstallPath returns the base directory for plugin installations.
//
// e.g. {BasePath}/store
func (p Paths) InstallPath() string { return filepath.Join(p.base, "store") }

// PluginInstallPath returns the path to install the plugin.
//
// e.g. {InstallPath}/{version}/{..files..}
func (p Paths) PluginInstallPath(plugin string) string {
	return filepath.Join(p.InstallPath(), plugin)
}

// PluginInstallReceiptPath returns the path to the installation receipt for plugin.
//
// e.g. {InstallReceiptsPath}/{plugin}.yaml
func (p Paths) PluginInstallReceiptPath(plugin string) string {
	return filepath.Join(p.InstallReceiptsPath(), plugin+consts.YAMLExtension)
}

// PluginVersionInstallPath returns the path to the specified version of specified
// plugin.
//
// e.g. {PluginInstallPath}/{plugin}/{version}
func (p Paths) PluginVersionInstallPath(plugin, version string) string {
	return filepath.Join(p.InstallPath(), plugin, version)
}

// Realpath evaluates symbolic links. If the path is not a symbolic link, it
// returns the cleaned path. Symbolic links with relative paths return error.
func Realpath(path string) (string, error) {
	s, err := os.Lstat(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to stat the currently executed path (%q)", path)
	}

	if s.Mode()&os.ModeSymlink != 0 {
		if path, err = os.Readlink(path); err != nil {
			return "", errors.Wrap(err, "failed to resolve the symlink of the currently executed version")
		}
		if !filepath.IsAbs(path) {
			return "", errors.Errorf("symbolic link is relative (%s)", path)
		}
	}
	return filepath.Clean(path), nil
}
