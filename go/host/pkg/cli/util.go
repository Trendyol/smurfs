package cli

import (
	"context"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
)

const (
	CommandManifestContextKey = "commandManifest"
)

func ContextWithCommandManifest(ctx context.Context, commandManifest *plugin.Plugin) context.Context {
	return context.WithValue(ctx, CommandManifestContextKey, commandManifest)
}
