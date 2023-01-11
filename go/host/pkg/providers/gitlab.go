package providers

import (
	"context"
	"fmt"
	"github.com/trendyol/smurfs/go/host/pkg/models"
	"net/http"
	"runtime"
)

type gitlabRelease struct {
	TagName string `json:"tag_name"`
	Assets  struct {
		Links []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"links"`
	} `json:"assets"`
}

type gitlabProvider struct {
	client http.Client
}

func NewGitlabProvider(client http.Client) DownloadProvider {
	return &gitlabProvider{
		client: client,
	}
}

func (g *gitlabProvider) ResolveArchive(ctx context.Context, distribution models.Distribution) (models.Archive, error) {
	projectId := distribution.Executable.ProviderData["projectId"]
	gitlabUrl := distribution.Executable.ProviderData["gitlabUrl"]
	token, hasToken := distribution.Executable.ProviderData["token"]

	distributionName := fmt.Sprintf("%s_%s_%s_%s.tar.gz", distribution.Executable.Archive, distribution.Version, runtime.GOOS, runtime.GOARCH)
	url := fmt.Sprintf("%s/api/v4/projects/%s/packages/generic/%s/%s/%s", gitlabUrl, projectId, distribution.Executable.Archive, distribution.Version, distributionName)

	if hasToken {
		url = fmt.Sprintf("%s?private_token=%s", url, token)
	}

	return models.Archive{
		URL:             url,
		SHA256:          "",
		CanSkipDownload: false,
	}, nil
}
