package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
)

var plugins = []plugin.Plugin{
	{
		Name:             "micro1",
		ShortDescription: "Micro CLI",
		LongDescription:  "Micro CLI",
		Usage:            "micro",
		Source: map[string]interface{}{
			"type":   "binary",
			"binary": "./micro1",
		},
	},
	{
		Name:             "micro2",
		ShortDescription: "Micro CLI",
		LongDescription:  "Micro CLI",
		Usage:            "micro1",
		Source: map[string]interface{}{
			"type":   "binary",
			"binary": "./micro2",
		},
	},
	{
		Name:             "onboarding",
		ShortDescription: "Onboarding CLI",
		LongDescription:  "Onboarding CLI",
		Usage:            "onboarding",
	},
}

var rootCmd = &cobra.Command{
	Use:   "host",
	Short: "Host CLI",
	Long:  "Host CLI",
}

type Logger struct {
}

func (l *Logger) Debug(message string, args ...interface{}) {

}

func (l *Logger) Info(message string, args ...interface{}) {
	fmt.Println("Fuck")
}

func (l *Logger) Warn(message string, args ...interface{}) {

}

func (l *Logger) Error(message string, args ...interface{}) {

}

func (l *Logger) Fatal(message string, args ...interface{}) {

}

func main() {
	logger := &Logger{}
	smurfHost, err := host.InitializeHost(host.Options{
		Plugins: plugins,
		RootCmd: rootCmd,
		Logger:  logger,
	})
	if err != nil {
		panic(err)
	}

	if err := smurfHost.Execute(); err != nil {
		panic(err)
	}
}
