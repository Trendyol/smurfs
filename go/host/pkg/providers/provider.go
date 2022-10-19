package providers

import (
	"context"
	"github.com/trendyol/smurfs/go/host/pkg/models"
)

const (
	GitHubExecutableProvider = models.ExecutableProvider("gitHub")
	GitlabExecutableProvider = models.ExecutableProvider("gitlab")
	DirectExecutableProvider = models.ExecutableProvider("direct")
	LocalExecutableProvider  = models.ExecutableProvider("local")
)

type DownloadProvider interface {
	ResolveArchive(ctx context.Context, distribution models.Distribution) (models.Archive, error)
}

var providerMap = map[models.ExecutableProvider]DownloadProvider{
	GitHubExecutableProvider: NewGithubProvider(),
	DirectExecutableProvider: NewDirectProvider(),
	GitlabExecutableProvider: NewGitlabProvider(),
	LocalExecutableProvider:  NewLocalProvider(),
}

func GetProviders() map[models.ExecutableProvider]DownloadProvider {
	return providerMap
}
