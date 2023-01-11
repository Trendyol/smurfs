package cli

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host/auth"
	"github.com/trendyol/smurfs/go/host/logger"
	"github.com/trendyol/smurfs/go/host/pkg/environment"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"github.com/trendyol/smurfs/go/host/storage"
)

type Options struct {
	Plugins         []*plugin.Plugin
	RootCmd         *cobra.Command
	HostAddress     string
	Logger          logger.Logger
	Auth            auth.Auth
	MetadataStorage storage.MetadataStorage
	PluginPath      string
}

// Builder is responsible for building the CLI from command manifests
type Builder interface {
	Build(options *Options) *cobra.Command
}

type cliBuilder struct {
	paths          environment.Paths
	commandWrapper CommandWrapper
}

func NewBuilder(paths environment.Paths, commandWrapper CommandWrapper) Builder {
	return &cliBuilder{
		paths:          paths,
		commandWrapper: commandWrapper,
	}
}

func (c *cliBuilder) Build(options *Options) *cobra.Command {
	for _, p := range options.Plugins {
		ctx := ContextWithCommandManifest(context.Background(), p)

		subCommand := &cobra.Command{
			Use:                p.Name,
			Run:                c.commandWrapper.Wrap(p),
			Short:              p.ShortDescription,
			Example:            p.Usage,
			DisableFlagParsing: true,
		}

		subCommand.SetContext(ctx)

		for _, flag := range p.Flags {
			if flag.Repeated {
				subCommand.Flags().StringSliceP(flag.Name, flag.Short, []string{}, flag.Description)
			} else {
				subCommand.Flags().StringP(flag.Name, flag.Short, "", flag.Description)
			}

			if flag.Required {
				subCommand.MarkFlagRequired(flag.Name)
			}
		}

		options.RootCmd.AddCommand(subCommand)
	}

	return options.RootCmd
}
