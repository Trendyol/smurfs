package main

import (
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host"
)

type MyLogger struct {
}

func (l *MyLogger) Debug(message string, args ...interface{}) {
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "host",
		Short: "Host CLI",
		Long:  "Host CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	plugins := []host.Plugin{
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
	_, err := host.InitializeHost(host.Options{
		Plugins: plugins,
		RootCmd: rootCmd,
	})
	if err != nil {
		panic(err)
	}
}
