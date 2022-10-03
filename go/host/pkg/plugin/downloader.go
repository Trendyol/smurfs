package plugin

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/download"
	"github.com/trendyol/smurfs/go/host/pkg/environment"
)

type Downloader interface {
	Download(ctx context.Context, distribution Distribution, destinationFolder string) error
}

type downloaderImpl struct {
	paths          environment.Paths
	providers      map[ExecutableProvider]DownloadProvider
	fileDownloader download.FileDownloader
}

func NewDownloader(
	paths environment.Paths,
	providers map[ExecutableProvider]DownloadProvider,
	fileDownloader download.FileDownloader,
) Downloader {
	return &downloaderImpl{
		paths:          paths,
		providers:      providers,
		fileDownloader: fileDownloader,
	}
}

func (d *downloaderImpl) Download(ctx context.Context, distribution Distribution, destinationFolder string) error {
	provider, ok := d.providers[distribution.Executable.Provider]
	if !ok {
		return errors.Wrapf(ErrUnknownArchiveProvider, "provider %s is not registered", distribution.Executable.Provider)
	}

	archive, err := provider.ResolveArchive(ctx, distribution)
	if err != nil {
		return errors.Wrap(err, "failed to resolve archive")
	}

	return d.fileDownloader.Download(ctx, archive.URL, destinationFolder)
}

type DownloadProvider interface {
	ResolveArchive(ctx context.Context, distribution Distribution) (Archive, error)
}

type uriProvider struct{}

func NewURIProvider() *uriProvider {
	return &uriProvider{}
}

func (u *uriProvider) ResolveArchive(ctx context.Context, distribution Distribution) (Archive, error) {
	archive := Archive{
		URL:    distribution.Executable.Address,
		SHA256: distribution.Executable.SHA256,
	}

	if archive.URL == "" {
		return Archive{}, errors.Wrapf(ErrEmptyArchiveAddress, "archive address cannot be empty for distribution %+v", distribution)
	}
	return archive, nil
}

type gitlabProvider struct{}

func (g *gitlabProvider) ResolveArchive(ctx context.Context, distribution Distribution) (Archive, error) {
	//TODO implement me
	panic("implement me")
}

type githubProvider struct{}

func (p *githubProvider) ResolveArchive(ctx context.Context, distribution Distribution) (Archive, error) {
	//TODO implement me
	archive := Archive{
		URL: fmt.Sprintf("%s/releases/download/%s/%s", distribution.Executable.Address, distribution.Version, distribution.Executable.GetArchiveName()),
	}
	return archive, nil
}
