package providers

import (
	"context"
	"fmt"
	"github.com/trendyol/smurfs/go/host/pkg/models"
)

type githubProvider struct{}

func NewGithubProvider() DownloadProvider {
	return &githubProvider{}
}

func (p *githubProvider) ResolveArchive(ctx context.Context, distribution models.Distribution) (models.Archive, error) {
	//TODO: It has many way to implement GitlabProvider. I couldn't decide which one is better but I would like to list all of them and I'll implement basic one first after all we can implement other cases and ways
	// 1. Implement fetching archive from public github repository
	// 2. Implement fetching archive from private github repository. We can prompt user for username and password or we can use github token

	archive := models.Archive{
		URL: fmt.Sprintf("%s/releases/download/%s/%s", distribution.Executable.Address, distribution.Version, distribution.Executable.Entrypoint),
	}
	return archive, nil
}
