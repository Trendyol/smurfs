package install

import (
	"github.com/trendyol/smurfs/host/pkg/download"
	"github.com/trendyol/smurfs/host/pkg/environment"
	"github.com/trendyol/smurfs/host/pkg/install/receipt"
	"github.com/trendyol/smurfs/host/pkg/plugin"
	"github.com/trendyol/smurfs/host/pkg/util/pathutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

type Installer interface {
	InstallPlugin(plugin plugin.Plugin, opts InstallOpts) error
}

type installer struct {
	paths environment.Paths
}

func New(paths environment.Paths) *installer {
	return &installer{
		paths: paths,
	}
}

// InstallOpts specifies options for plugin installation operation.
type InstallOpts struct {
	ArchiveFileOverride string
}

type installOperation struct {
	pluginName string
	platform   plugin.Runnable

	installDir string
	binDir     string
}

// Plugin lifecycle errors
var (
	ErrIsAlreadyInstalled = errors.New("can't install, the newest version is already installed")
	ErrIsNotInstalled     = errors.New("plugin is not installed")
	ErrIsAlreadyUpgraded  = errors.New("can't upgrade, the newest version is already installed")
)

// Install will download and install a plugin. The operation tries
// to not get the plugin dir in a bad state if it fails during the process.
func Install(p environment.Paths, plugin plugin.Plugin, opts InstallOpts) error {
	klog.V(2).Infof("Looking for installed versions")
	_, err := receipt.Load(p.PluginInstallReceiptPath(plugin.Name))

	// todo: compare versions
	if err == nil {
		return ErrIsAlreadyInstalled
	} else if !os.IsNotExist(err) {
		return errors.Wrap(err, "failed to look up plugin receipt")
	}

	// The actual install should be the last action so that a failure during receipt
	// saving does not result in an installed plugin without receipt.
	klog.V(3).Infof("Install plugin %s at version=%s", plugin.Name, plugin.Spec.Version)
	if err := install(installOperation{
		pluginName: plugin.Name,
		platform:   plugin.Spec.Platform,
		binDir:     p.BinPath(),
		installDir: p.PluginVersionInstallPath(plugin.Name, plugin.Spec.Version),
	}, opts); err != nil {
		return errors.Wrap(err, "install failed")
	}

	klog.V(3).Infof("Storing install receipt for plugin %s", plugin.Name)
	err = receipt.Store(receipt.New(plugin, metav1.Now()), p.PluginInstallReceiptPath(plugin.Name))

	return errors.Wrap(err, "installation receipt could not be stored, uninstall may fail")
}

func install(op installOperation, opts InstallOpts) error {
	// Download and extract
	klog.V(3).Infof("Creating download staging directory")

	downloadStagingDir, err := os.MkdirTemp("", "ty-downloads")
	if err != nil {
		return errors.Wrapf(err, "could not create staging dir %q", downloadStagingDir)
	}

	klog.V(3).Infof("Successfully created download staging directory %q", downloadStagingDir)

	defer func() {
		klog.V(3).Infof("Deleting the download staging directory %s", downloadStagingDir)
		if err := os.RemoveAll(downloadStagingDir); err != nil {
			klog.Warningf("failed to clean up download staging directory: %s", err)
		}
	}()

	if err := downloadAndExtract(downloadStagingDir, op.platform.URI, op.platform.Sha256, opts.ArchiveFileOverride); err != nil {
		return errors.Wrap(err, "failed to unpack into staging dir")
	}

	applyDefaults(&op.platform)
	if err := moveToInstallDir(downloadStagingDir, op.installDir, op.platform.Files); err != nil {
		return errors.Wrap(err, "failed while moving files to the installation directory")
	}

	subPathAbs, err := filepath.Abs(op.installDir)
	if err != nil {
		return errors.Wrapf(err, "failed to get the absolute fullPath of %q", op.installDir)
	}
	fullPath := filepath.Join(op.installDir, filepath.FromSlash(op.platform.Bin))
	pathAbs, err := filepath.Abs(fullPath)
	if err != nil {
		return errors.Wrapf(err, "failed to get the absolute fullPath of %q", fullPath)
	}
	if _, ok := pathutil.IsSubPath(subPathAbs, pathAbs); !ok {
		return errors.Wrapf(err, "the fullPath %q does not extend the sub-fullPath %q", fullPath, op.installDir)
	}
	err = createOrUpdateLink(op.binDir, fullPath, op.pluginName)
	return errors.Wrap(err, "failed to link installed plugin")
}

func applyDefaults(platform *plugin.Runnable) {
	if platform.Files == nil {
		platform.Files = []plugin.FileOperation{{From: "*", To: "."}}
		klog.V(4).Infof("file operation not specified, assuming %v", platform.Files)
	}
}

// downloadAndExtract downloads the specified archive uri (or uses the provided overrideFile, if a non-empty value)
// while validating its checksum with the provided sha256sum, and extracts its contents to extractDir that must be.
// created.
func downloadAndExtract(extractDir, uri, sha256sum, overrideFile string) error {
	var fetcher download.Fetcher = download.HTTPFetcher{}
	if overrideFile != "" {
		fetcher = download.NewFileFetcher(overrideFile)
	}

	verifier := download.NewSha256Verifier(sha256sum)
	err := download.NewDownloader(verifier, fetcher).Get(uri, extractDir)
	return errors.Wrap(err, "failed to unpack the plugin archive")
}

func createOrUpdateLink(binDir, binary, plugin string) error {
	dst := filepath.Join(binDir, pluginNameToBin(plugin, IsWindows()))

	if err := removeLink(dst); err != nil {
		return errors.Wrap(err, "failed to remove old symlink")
	}
	if _, err := os.Stat(binary); os.IsNotExist(err) {
		return errors.Wrapf(err, "can't create symbolic link, source binary (%q) cannot be found in extracted archive", binary)
	}

	// Create new
	klog.V(2).Infof("Creating symlink to %q at %q", binary, dst)
	if err := os.Symlink(binary, dst); err != nil {
		return errors.Wrapf(err, "failed to create a symlink from %q to %q", binary, dst)
	}
	klog.V(2).Infof("Created symlink at %q", dst)

	return nil
}

// removeLink removes a symlink reference if exists.
func removeLink(path string) error {
	fi, err := os.Lstat(path)
	if os.IsNotExist(err) {
		klog.V(3).Infof("No file found at %q", path)
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "failed to read the symlink in %q", path)
	}

	if fi.Mode()&os.ModeSymlink == 0 {
		return errors.Errorf("file %q is not a symlink (mode=%s)", path, fi.Mode())
	}
	if err := os.Remove(path); err != nil {
		return errors.Wrapf(err, "failed to remove the symlink in %q", path)
	}
	klog.V(3).Infof("Removed symlink from %q", path)
	return nil
}

// IsWindows sees if KREW_OS or runtime.GOOS to find out if current execution mode is win32.
func IsWindows() bool {
	goos := runtime.GOOS
	if env := os.Getenv("KREW_OS"); env != "" {
		goos = env
	}
	return goos == "windows"
}

// pluginNameToBin creates the name of the symlink file for the plugin name.
// It converts dashes to underscores.
func pluginNameToBin(name string, isWindows bool) string {
	name = strings.ReplaceAll(name, "-", "_")
	name = "ty-" + name
	if isWindows {
		name += ".exe"
	}
	return name
}

// CleanupStaleKrewInstallations removes the versions that aren't the current version.
func CleanupStaleKrewInstallations(dir, currentVersion string) error {
	ls, err := os.ReadDir(dir)
	if err != nil {
		return errors.Wrap(err, "failed to read krew store directory")
	}
	klog.V(2).Infof("Found %d entries in krew store directory", len(ls))
	for _, d := range ls {
		klog.V(2).Infof("Found a krew installation: %s (%s)", d.Name(), d.Type())
		if d.IsDir() && d.Name() != currentVersion {
			klog.V(1).Infof("Deleting stale krew install directory: %s", d.Name())
			p := filepath.Join(dir, d.Name())
			if err := os.RemoveAll(p); err != nil {
				return errors.Wrapf(err, "failed to remove stale krew version at path '%s'", p)
			}
			klog.V(1).Infof("Stale installation directory removed")
		}
	}
	return nil
}