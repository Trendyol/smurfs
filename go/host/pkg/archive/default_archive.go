package archive

import (
	"context"
)

type defaultExtractor struct{}

func NewDefaultExtractor() Extractor {
	return &defaultExtractor{}
}

func (e *defaultExtractor) Extract(ctx context.Context, sourceFilePath, destinationFolderPath string) error {
	return nil
}
