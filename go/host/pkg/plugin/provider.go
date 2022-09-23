package plugin

import (
	"context"
)

type Provider interface {
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
	return archive, nil
}

type gitlabProvider struct{}

func (g *gitlabProvider) ResolveArchive(ctx context.Context, distribution Distribution) (Archive, error) {
	//TODO implement me
	panic("implement me")
}

type githubProvider struct{}
