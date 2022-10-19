package archive

import (
	"context"
	"fmt"
	"github.com/trendyol/smurfs/go/host/pkg/util"
	"path"
	"path/filepath"
)

type defaultExtractor struct{}

func NewDefaultExtractor() Extractor {
	return &defaultExtractor{}
}

func (e *defaultExtractor) Extract(ctx context.Context, sourceFilePath, destinationFolderPath string) error {
	destinationFilePath := filepath.Base(sourceFilePath)
	fmt.Printf("Source %s, Destination %s\n", sourceFilePath, destinationFilePath)
	_, err := util.CopyFile(sourceFilePath, path.Join(destinationFolderPath, destinationFilePath))

	return err
}
