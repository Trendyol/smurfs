package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/trendyol/smurfs/go/host/pkg/models"
	"net/http"
	"net/url"
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

	releaseUrl, err := url.Parse(fmt.Sprintf("%s/api/v4/projects/%s/releases", gitlabUrl, projectId))
	if err != nil {
		return models.Archive{}, errors.Wrap(err, "failed to parse gitlab url")
	}
	request := http.Request{
		Method: "GET",
		URL:    releaseUrl,
	}

	if hasToken {
		request.Header = http.Header{
			"PRIVATE-TOKEN": []string{token},
		}
	}

	resp, err := g.client.Do(&request)
	if err != nil {
		return models.Archive{}, errors.Wrapf(err, "failed to get releases from gitlab")
	}

	if resp.StatusCode != http.StatusOK {
		return models.Archive{}, errors.Errorf("failed to get releases from gitlab, status code: %d", resp.StatusCode)
	}

	var releases []gitlabRelease

	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return models.Archive{}, errors.Wrapf(err, "failed to decode gitlab releases")
	}

	distributionName := fmt.Sprintf("%s_%s_%s.tar.gz", distribution.Executable.Archive, distribution.Version, distribution.Targets[0])

	for _, release := range releases {
		if release.TagName == fmt.Sprintf("v%s", distribution.Version) {
			for _, asset := range release.Assets.Links {
				if asset.Name == distributionName {
					return models.Archive{
						URL:             asset.URL,
						SHA256:          "",
						CanSkipDownload: false,
					}, nil
				}
			}
		}
	}

	return models.Archive{}, errors.Errorf("failed to find archive for version: %s", distribution.Version)
}
