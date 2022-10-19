package providers

import (
	"context"
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/models"
)

type localProvider struct{}

func NewLocalProvider() DownloadProvider {
	return &localProvider{}
}

func (u *localProvider) ResolveArchive(ctx context.Context, distribution models.Distribution) (models.Archive, error) {
	archive := models.Archive{
		URL:             distribution.Executable.Address,
		SHA256:          distribution.Executable.SHA256,
		CanSkipDownload: true,
	}

	if archive.URL == "" {
		return models.Archive{}, errors.Wrapf(models.ErrEmptyArchiveAddress, "archive address cannot be empty for distribution %+v", distribution)
	}
	return archive, nil
}
