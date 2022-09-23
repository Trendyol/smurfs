package cli

import (
	"context"
	"github.com/trendyol/smurfs/go/protos"
)

const (
	CommandManifestContextKey = "commandManifest"
)

func ContextWithCommandManifest(ctx context.Context, commandManifest *protos.Command) context.Context {
	return context.WithValue(ctx, CommandManifestContextKey, commandManifest)
}

func GetCommandManifest(ctx context.Context) *protos.Command {
	command, _ := ctx.Value(CommandManifestContextKey).(*protos.Command)
	return command
}
