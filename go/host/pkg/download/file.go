package download

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

type FileDownloader interface {
	Download(ctx context.Context, uri, destinationFolder string) (string, error)
}

type fileDownloader struct {
	httpClient *http.Client
}

func NewFileDownloader(httpClient *http.Client) FileDownloader {
	return &fileDownloader{
		httpClient: httpClient,
	}
}

func (d *fileDownloader) Download(ctx context.Context, uri, destinationFolder string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return "", errors.Wrapf(err, "could not create download request for %s", uri)
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return "", errors.Wrapf(err, "could not download %s", uri)
	}
	defer resp.Body.Close()

	parsedUrl, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}

	pathSegments := strings.Split(parsedUrl.Path, "/")
	filePath := path.Join(destinationFolder, pathSegments[len(pathSegments)-1])

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "could not download %s", uri)
	}

	return filePath, nil
}
