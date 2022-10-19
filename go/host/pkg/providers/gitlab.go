package providers

import (
	"context"
	"github.com/trendyol/smurfs/go/host/pkg/models"
)

type gitlabProvider struct{}

func NewGitlabProvider() DownloadProvider {
	return &gitlabProvider{}
}

func (g *gitlabProvider) ResolveArchive(ctx context.Context, distribution models.Distribution) (models.Archive, error) {
	//TODO implement me
	panic("implement me")
}
