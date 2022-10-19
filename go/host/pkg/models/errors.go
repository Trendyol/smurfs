package models

import "errors"

var (
	ErrPluginNotInstalled     = errors.New("plugin not installed")
	ErrPluginAlreadyInstalled = errors.New("plugin already installed")
	ErrPluginNotFound         = errors.New("plugin not found")
	ErrEmptyArchiveAddress    = errors.New("empty archive address")
	ErrUnknownArchiveProvider = errors.New("unknown archive provider")
)
