package plugin

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/trendyol/smurfs/host/pkg/archive"
	"github.com/trendyol/smurfs/host/pkg/consts"
	"github.com/trendyol/smurfs/host/pkg/download"
	"github.com/trendyol/smurfs/host/pkg/environment"
	"github.com/trendyol/smurfs/host/pkg/install/receipt"
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
	paths      environment.Paths
	downloader download.Downloader
	extractor  archive.Extractor
	verifier   download.Verifier
}

func NewManager(
	paths environment.Paths,
	downloader download.Downloader,
	extractor archive.Extractor,
	verifier download.Verifier,
) Manager {
	return &manager{
		paths:      paths,
		downloader: downloader,
		extractor:  extractor,
		verifier:   verifier,
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
		pluginReceipt, err := receipt.Load(pluginReceiptPath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load receipt")
		}

		receipts = append(receipts, pluginReceipt)
	}

	return receipts, nil
}

func (m *manager) GetPluginReceipt(ctx context.Context, pluginName string) (Receipt, error) {
	installReceiptsPath := m.paths.InstallReceiptsPath()

	pluginReceiptPath := path.Join(installReceiptsPath, pluginName+consts.YAMLExtension)
	pluginReceipt, err := receipt.Load(pluginReceiptPath)
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
		return errors.Wrap(err, "failed to get plugin receipt")
	}
	if err == nil {
		return errors.Errorf("plugin %q is already installed", plugin.Name)
	}

	tempDir, err := os.MkdirTemp("", "smurfs-temp")
	if err != nil {
		return errors.Wrap(err, "could not create temp dir for plugin installation")
	}

	// clean tempDir after the installation
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			logger.WithError(err).Warningf("could not remove temp dir %q after the installation of the plugin %s", tempDir, plugin.Name)
		}
	}()

	err = m.downloader.Download(ctx, plugin.Spec.Runnable.URI, tempDir)
	if err != nil {
		return errors.Wrap(err, "could not download plugin")
	}

	downloadedArchivePath := path.Join(tempDir, filepath.Base(plugin.Spec.Runnable.URI))

	sha256Verifier := download.NewSha256Verifier(plugin.Spec.Runnable.Sha256)
	if err = sha256Verifier.VerifyFile(ctx, downloadedArchivePath); err != nil {
		return errors.Wrap(err, "could not verify downloaded archive")
	}

	if err = m.extractor.Extract(ctx, downloadedArchivePath, tempDir); err != nil {
		return errors.Wrap(err, "could not extract downloaded archive")
	}

	// todo: move plugin to store path

	return nil
}

func (m *manager) Uninstall(ctx context.Context, plugin Plugin) error {
	//TODO implement me
	panic("implement me")
}
