package plugin

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/trendyol/smurfs/go/host/pkg/archive"
	"github.com/trendyol/smurfs/go/host/pkg/consts"
	"github.com/trendyol/smurfs/go/host/pkg/environment"
	"github.com/trendyol/smurfs/go/host/pkg/verifier"
	"os"
	"path"
	"path/filepath"
)

// TODO: use progressbar

// Manager is the interface for managing plugins.
type Manager interface {
	// List lists all installed plugins.
	List(ctx context.Context) ([]Receipt, error)
	// GetPluginReceipt returns the receipt of the given plugin.
	GetPluginReceipt(ctx context.Context, pluginName string) (Receipt, error)
	// Install installs a plugin.
	Install(ctx context.Context, plugin Plugin) error
	// Uninstall uninstalls a plugin.
	Uninstall(ctx context.Context, plugin Plugin) error
}

type manager struct {
	paths          environment.Paths
	downloader     Downloader
	extractor      archive.Extractor
	sha256Verifier verifier.Verifier
}

func NewManager(
	paths environment.Paths,
	downloader Downloader,
	extractor archive.Extractor,
	sha256Verifier verifier.Verifier,
) Manager {
	return &manager{
		paths:          paths,
		downloader:     downloader,
		extractor:      extractor,
		sha256Verifier: sha256Verifier,
	}
}

func (m *manager) List(ctx context.Context) ([]Receipt, error) {
	logger := logrus.WithContext(ctx)
	installReceiptsPath := m.paths.InstallReceiptsPath()

	dirEntries, err := os.ReadDir(installReceiptsPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read install receipts directory")
	}

	receipts := make([]Receipt, 0, len(dirEntries))
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			logger.Debugf("skipping directory %q in receipts folder %s", dirEntry.Name(), installReceiptsPath)
			continue
		}

		pluginReceiptPath := path.Join(installReceiptsPath, dirEntry.Name())
		pluginReceipt, err := LoadReceiptFrom(pluginReceiptPath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load receipt")
		}

		receipts = append(receipts, pluginReceipt)
	}

	return receipts, nil
}

func (m *manager) GetPluginReceipt(ctx context.Context, pluginName string) (Receipt, error) {
	pluginReceiptPath := path.Join(m.paths.InstallReceiptsPath(), pluginName+consts.YAMLExtension)
	pluginReceipt, err := LoadReceiptFrom(pluginReceiptPath)
	if err != nil {
		return Receipt{}, errors.Wrap(err, "failed to load receipt")
	}

	return pluginReceipt, nil
}

func (m *manager) Install(ctx context.Context, plugin Plugin) error {
	logger := logrus.WithContext(ctx)

	// check plugin is already installed
	_, err := m.GetPluginReceipt(ctx, plugin.Name)
	if err != nil {
		return errors.Wrapf(err, "could not get plugin %q receipt", plugin.Name)
	}
	if err == nil {
		return errors.Wrapf(ErrPluginAlreadyInstalled, "plugin %q is already installed", plugin.Name)
	}

	distribution, err := plugin.GetCompatibleDistribution()
	if err != nil {
		return errors.Wrapf(err, "could not get compatible distribution for plugin %q", plugin.Name)
	}

	// download plugin to a temporary directory
	tempDir, err := os.MkdirTemp("", "smurfs-temp")
	if err != nil {
		return errors.Wrapf(err, "could not create temporary directory for plugin %q", plugin.Name)
	}

	// clean tempDir after the installation
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			logger.WithError(err).Warningf("could not remove temp dir %q after the installation of the plugin %s", tempDir, plugin.Name)
		}
	}()

	err = m.downloader.Download(ctx, distribution, tempDir)
	if err != nil {
		return errors.Wrapf(err, "could not download plugin %q", plugin.Name)
	}

	archivePath := path.Join(tempDir, filepath.Base(distribution.Executable.Address))

	// todo: verify downloaded archive here

	if err = m.extractor.Extract(ctx, archivePath, tempDir); err != nil {
		return errors.Wrapf(err, "could not extract plugin %q", plugin.Name)
	}

	receipt := plugin.GenerateReceipt(distribution)

	// move archive contents to plugin installation directory
	if err = m.moveArchiveContents(archivePath, receipt); err != nil {
		return errors.Wrapf(err, "could not move archive contents of plugin %q", plugin.Name)
	}

	// save receipt
	receiptPath := path.Join(m.paths.InstallReceiptsPath(), plugin.Name+consts.YAMLExtension)
	if err = receipt.Store(receiptPath); err != nil {
		return errors.Wrapf(err, "could not store receipt for plugin %q", plugin.Name)
	}

	return nil
}

func (m *manager) Uninstall(ctx context.Context, plugin Plugin) error {
	receiptPath := path.Join(m.paths.InstallReceiptsPath(), plugin.Name+consts.YAMLExtension)
	if err := os.RemoveAll(receiptPath); err != nil {
		return errors.Wrapf(err, "could not remove receipt for plugin %q", plugin.Name)
	}

	pluginPath := path.Join(m.paths.InstallPath(), plugin.Name)
	if err := os.RemoveAll(pluginPath); err != nil {
		return errors.Wrapf(err, "could not remove plugin %q", plugin.Name)
	}

	return nil
}

func (m *manager) moveArchiveContents(archivePath string, receipt Receipt) error {
	archiveContents, err := os.ReadDir(archivePath)
	if err != nil {
		return errors.Wrapf(err, "could not read archive contents %q", archivePath)
	}

	pluginInstallPath := path.Join(m.paths.InstallPath(), receipt.Name, receipt.Executable.Version)

	for _, archiveContent := range archiveContents {
		name := archiveContent.Name()
		sourcePath := path.Join(archivePath, name)
		destinationPath := path.Join(pluginInstallPath, name)
		if err = os.Rename(sourcePath, destinationPath); err != nil {
			return errors.Wrapf(err, "could not move file %q for plugin %q", name, receipt.Name)
		}
	}

	return nil
}
