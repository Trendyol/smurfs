package plugin

import "github.com/pkg/errors"

var (
	ErrPluginAlreadyInstalled = errors.New("plugin already installed")
	ErrPluginNotFound         = errors.New("plugin not found")
	ErrEmptyArchiveAddress    = errors.New("empty archive address")
	ErrUnknownArchiveProvider = errors.New("unknown archive provider")
)
