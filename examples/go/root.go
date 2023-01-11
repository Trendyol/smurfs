package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host"
	"github.com/trendyol/smurfs/go/host/auth"
	"github.com/trendyol/smurfs/go/host/pkg/cli"
	"github.com/trendyol/smurfs/go/host/pkg/models"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"github.com/trendyol/smurfs/go/host/pkg/util/pathutil"
)

var plugins = []*plugin.Plugin{
	{
		Name:             "micro1",
		ShortDescription: "Micro CLI",
		LongDescription:  "Micro CLI",
		Usage:            "micro",
		Source: map[string]interface{}{
			"type":   "binary",
			"binary": "./micro1",
		},
		Distributions: []models.Distribution{
			{
				Version:          "1.0.0",
				Targets:          []string{"darwin_arm64", "darwin_amd64", "linux_arm64", "linux_amd64", "windows_x64", "windows_x86"},
				SkipVerification: true,
				Executable: models.Executable{
					Provider: "local",
					Address:  "./micro1",
				},
			},
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
		Distributions: []models.Distribution{
			{
				Executable: models.Executable{
					Provider: "local",
					Address:  "./micro2",
				},
				Version:          "1.0.0",
				Targets:          []string{"darwin_arm64"},
				SkipVerification: true,
			},
		},
	},
	{
		Name:             "onboarding",
		ShortDescription: "Onboarding CLI",
		LongDescription:  "Onboarding CLI",
		Usage:            "onboarding",
		Source: map[string]interface{}{
			"type":      "gitlab",
			"projectId": "12345",
			"gitlabUrl": "https://gitlab.abc.com",
		},
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
	fmt.Printf("HOST-DEBUG: %s\n", message)
}

func (l *Logger) Info(message string, args ...interface{}) {
	fmt.Printf("HOST-INFO: %s\n", message)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	fmt.Printf("HOST-WARN: %s\n", message)
}

func (l *Logger) Error(message string, args ...interface{}) {
	fmt.Printf("HOST-ERROR: %s\n", message)
}

func (l *Logger) Fatal(message string, args ...interface{}) {
	fmt.Printf("HOST-FATAL: %s\n", message)
}

type Auth struct {
}

func (a Auth) GetToken() (auth.TokenResponse, error) {
	return auth.TokenResponse{
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
		RptToken:     "rpt-token",
	}, nil
}

func (a Auth) GetUserInfo() (auth.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

type MetadataStorage struct {
}

func (m MetadataStorage) Get(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MetadataStorage) Set(key string, value string) error {
	//TODO implement me
	panic("implement me")
}

func main() {
	hostAuth := &Auth{}
	smurfHost, err := host.InitializeHost(&cli.Options{
		Plugins:    plugins,
		RootCmd:    rootCmd,
		Auth:       hostAuth,
		PluginPath: fmt.Sprintf("%s/.smurfs", pathutil.GetHomeDir()),
	})
	if err != nil {
		panic(err)
	}

	if err := smurfHost.Execute(); err != nil {
		panic(err)
	}
}
