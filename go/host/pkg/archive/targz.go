package archive

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
)

type tarGzExtractor struct{}

func NewTarGzExtractor() Extractor {
	return &tarGzExtractor{}
}

func (e *tarGzExtractor) Extract(ctx context.Context, sourceFilePath, destinationFolderPath string) error {
	file, err := os.Open(sourceFilePath)
	if err != nil {
		return errors.Wrapf(err, "could not open targz source path %q", sourceFilePath)
	}

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return errors.Wrapf(err, "could not create gzip reader for targz source path %q", sourceFilePath)
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrapf(err, "could not read next tar entry for targz source path %q", sourceFilePath)
		}

		// see https://golang.org/cl/78355 for handling pax_global_header
		if hdr.Name == "pax_global_header" {
			continue
		}

		if err := suspiciousPath(hdr.Name); err != nil {
			return err
		}

		path := filepath.Join(destinationFolderPath, filepath.FromSlash(hdr.Name))
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, os.FileMode(hdr.Mode)); err != nil {
				return errors.Wrap(err, "could not create directory from targz")
			}
		case tar.TypeReg:
			dir := filepath.Dir(path)
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return errors.Wrap(err, "could not create directory for targz")
			}
			f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
			if err != nil {
				return errors.Wrapf(err, "could not create file %q", path)
			}

			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return errors.Wrapf(err, "could not copy %q from targz into file", hdr.Name)
			}
			f.Close()
		default:
			return errors.Errorf("could not handle tar header type %d in %q", hdr.Typeflag, hdr.Name)
		}
	}
	return nil
}
