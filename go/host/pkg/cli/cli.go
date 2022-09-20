package cli

import (
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host/pkg/environment"
	"github.com/trendyol/smurfs/go/host/protos"
)

const (
	cliName      = "ty"
	version      = "0.0.1"
	shortMessage = "Trendyol CLI"
)

// Builder is responsible for building the CLI from command manifests
type Builder interface {
	Build(commandManifests []*protos.Command) *cobra.Command
}

type cliBuilder struct {
	paths          environment.Paths
	commandWrapper CommandWrapper
}

func New(paths environment.Paths, commandWrapper CommandWrapper) Builder {
	return &cliBuilder{
		paths:          paths,
		commandWrapper: commandWrapper,
	}
}

func (c *cliBuilder) Build(commandManifests []*protos.Command) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     cliName,
		Version: version,
		Short:   shortMessage,
	}

	for _, cmdManifest := range commandManifests {
		subCommand := &cobra.Command{
			Use:     cmdManifest.Name,
			Run:     c.commandWrapper.Wrap(cmdManifest),
			Short:   cmdManifest.Description,
			Example: cmdManifest.Usage,
		}

		for _, flag := range cmdManifest.Flags {
			if flag.Repeated {
				subCommand.Flags().StringSliceP(flag.Name, flag.Short, []string{}, flag.Description)
			} else {
				subCommand.Flags().StringP(flag.Name, flag.Short, "", flag.Description)
			}

			if flag.Required {
				subCommand.MarkFlagRequired(flag.Name)
			}
		}

		rootCmd.AddCommand(subCommand)
	}

	return rootCmd
}

var commandManifests = []*protos.Command{
	{
		Name:        "login",
		Description: "to login with LDAP credentials",
		Flags: []*protos.CommandFlag{
			{
				Name:        "email",
				Required:    true,
				Repeated:    false,
				Description: "User email",
			},
			{
				Name:        "password",
				Required:    true,
				Repeated:    false,
				Description: "User password",
			},
		},
	},
	{
		Name:        "logout",
		Description: "to logout from CLI",
		Flags: []*protos.CommandFlag{
			{
				Name:        "email",
				Required:    false,
				Repeated:    false,
				Description: "User email",
			},
		},
	},
}
