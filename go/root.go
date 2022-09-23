package main

import (
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host"
)

var plugins = []host.Plugin{
	{
		Name:             "micro1",
		ShortDescription: "Micro CLI",
		LongDescription:  "Micro CLI",
		Usage:            "micro",
	},
	{
		Name:             "micro2",
		ShortDescription: "Micro CLI",
		LongDescription:  "Micro CLI",
		Usage:            "micro",
	},
}

var rootCmd = &cobra.Command{
	Use:   "host",
	Short: "Host CLI",
	Long:  "Host CLI",
}

func main() {
	smurfHost, err := host.InitializeHost(host.Options{
		Plugins: plugins,
		RootCmd: rootCmd,
	})
	if err != nil {
		panic(err)
	}

	if err := smurfHost.Execute(); err != nil {
		panic(err)
	}
}
