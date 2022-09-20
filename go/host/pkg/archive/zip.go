package archive

import (
	"archive/zip"
	"context"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
)

type zipExtractor struct{}

func (e *zipExtractor) Extract(ctx context.Context, sourceFilePath, destinationFolderPath string) error {
	zipReader, err := zip.OpenReader(sourceFilePath)
	if err != nil {
		return errors.Wrapf(err, "could not open zip reader for source path %q", sourceFilePath)
	}

	for _, f := range zipReader.File {
		if err := suspiciousPath(f.Name); err != nil {
			return err
		}

		path := filepath.Join(destinationFolderPath, filepath.FromSlash(f.Name))
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(path, f.Mode()); err != nil {
				return errors.Wrapf(err, "could not create directory %q", path)
			}
			continue
		}

		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return errors.Wrapf(err, "could not create directory %q", dir)
		}
		src, err := f.Open()
		if err != nil {
			return errors.Wrapf(err, "could not open file %q", f.Name)
		}

		dst, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		if err != nil {
			src.Close()
			return errors.Wrapf(err, "could not open file %q", path)
		}
		closeAll := func() {
			src.Close()
			dst.Close()
		}

		if _, err := io.Copy(dst, src); err != nil {
			closeAll()
			return errors.Wrapf(err, "could not copy file %q to destination", f.Name)
		}
		closeAll()
	}

}
