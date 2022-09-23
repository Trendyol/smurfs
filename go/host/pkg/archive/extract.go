package archive

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/pkg/errors"
)

const (
	zipMimeType   = "application/zip"
	tarGzMimeType = "application/x-gzip"
)

type Extractor interface {
	Extract(ctx context.Context, sourceArchivePath, destinationFolderPath string) error
}

type extractorManager struct {
	extractors map[string]Extractor
}

func NewExtractorManager(extractors map[string]Extractor) *extractorManager {
	return &extractorManager{
		extractors: extractors,
	}
}

func (e *extractorManager) Extract(ctx context.Context, sourceArchivePath, destinationFolderPath string) error {
	mimeType, err := mimetype.DetectFile(sourceArchivePath)
	if err != nil {
		return errors.Wrapf(err, "could not find mimetype of the archive file %s", sourceArchivePath)
	}

	extractor, ok := e.extractors[mimeType.String()]
	if !ok {
		return errors.Errorf("unsupported archive type %s", mimeType.String())
	}

	return extractor.Extract(ctx, sourceArchivePath, destinationFolderPath)
}
