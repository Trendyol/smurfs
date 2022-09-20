package download

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

type Downloader interface {
	Download(ctx context.Context, uri, destinationFolder string) error
}

type downloader struct{}

func (d *downloader) Download(ctx context.Context, uri, destinationFolder string) error {
	// todo: override client
	resp, err := http.Get(uri)
	if err != nil {
		return errors.Wrapf(err, "could not download %s", uri)
	}
	defer resp.Body.Close()

	uriSegments := strings.Split(uri, "/")
	filePath := path.Join(destinationFolder, uriSegments[len(uriSegments)-1])

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return errors.Wrapf(err, "could not download %s", uri)
	}

	return nil
}
