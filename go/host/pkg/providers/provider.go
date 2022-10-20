package providers

import (
	"context"
	"github.com/trendyol/smurfs/go/host/pkg/models"
	"net/http"
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

var providerMap = map[models.ExecutableProvider]DownloadProvider{}

func InitProviders(client http.Client) {
	providerMap[GitHubExecutableProvider] = NewGithubProvider()
	providerMap[DirectExecutableProvider] = NewDirectProvider()
	providerMap[GitlabExecutableProvider] = NewGitlabProvider(client)
	providerMap[LocalExecutableProvider] = NewLocalProvider()
}

func GetProviders() map[models.ExecutableProvider]DownloadProvider {
	return providerMap
}
