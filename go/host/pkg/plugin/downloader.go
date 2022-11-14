package plugin

import (
	"context"
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/download"
	"github.com/trendyol/smurfs/go/host/pkg/environment"
	"github.com/trendyol/smurfs/go/host/pkg/models"
	"github.com/trendyol/smurfs/go/host/pkg/providers"
	"github.com/trendyol/smurfs/go/host/pkg/util"
	"os"
	"path"
	"path/filepath"
)

type Downloader interface {
	Download(ctx context.Context, distribution models.Distribution, destinationFolder string) (string, error)
}

type downloaderImpl struct {
	paths          environment.Paths
	providers      map[models.ExecutableProvider]providers.DownloadProvider
	fileDownloader download.FileDownloader
}

func NewDownloader(
	paths environment.Paths,
	providers map[models.ExecutableProvider]providers.DownloadProvider,
	fileDownloader download.FileDownloader,
) Downloader {
	return &downloaderImpl{
		paths:          paths,
		providers:      providers,
		fileDownloader: fileDownloader,
	}
}

func (d *downloaderImpl) Download(ctx context.Context, distribution models.Distribution, destinationFolder string) (string, error) {
	provider, ok := d.providers[distribution.Executable.Provider]
	if !ok {
		return "", errors.Wrapf(models.ErrUnknownArchiveProvider, "provider %s is not registered", distribution.Executable.Provider)
	}

	archive, err := provider.ResolveArchive(ctx, distribution)
	if err != nil {
		return "", errors.Wrap(err, "failed to resolve archive")
	}

	if !archive.CanSkipDownload {
		return d.fileDownloader.Download(ctx, archive.URL, destinationFolder)
	} else {
		currentPath, _ := os.Getwd()
		sourceFullPath := path.Join(currentPath, archive.URL)
		destinationFullPath := path.Join(destinationFolder, filepath.Base(distribution.Executable.Address))
		_, err := util.CopyFile(sourceFullPath, destinationFullPath)
		return path.Join(destinationFolder, filepath.Base(distribution.Executable.Address)), err
	}
}
