package host

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type SmurfHost struct {
}

type Options struct {
	Plugins []Plugin
	RootCmd *cobra.Command
}

type Plugin struct {
	Name             string
	ShortDescription string
	LongDescription  string
	Usage            string
	Binaries         []struct {
		Name   string
		Target string
		Args   interface{}
	}
}

func InitializeHost(options Options) (*SmurfHost, error) {
	if options.RootCmd == nil {
		options.RootCmd = &cobra.Command{
			Use:   "host",
			Short: "Host CLI",
			Long:  "Host CLI",
		}
	}

	for _, plugin := range options.Plugins {
		cmd := &cobra.Command{
			Use:   plugin.Name,
			Short: plugin.ShortDescription,
			Long:  plugin.LongDescription,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Micro CLI is running")
			},
		}

		options.RootCmd.AddCommand(cmd)
	}
	if err := options.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &SmurfHost{}, nil
}
