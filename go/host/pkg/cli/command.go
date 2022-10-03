package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"github.com/trendyol/smurfs/go/host/pkg/process"
	"github.com/trendyol/smurfs/go/protos"
)

// CommandWrapper is responsible for wrapping the command manifest to cobra command
type CommandWrapper interface {
	Wrap(*protos.Command) func(cmd *cobra.Command, args []string)
}

type wrapper struct {
	exec          process.Exec
	pluginManager plugin.Manager
}

func NewWrapper(exec process.Exec) CommandWrapper {
	return &wrapper{
		exec: exec,
	}
}

func (w *wrapper) Wrap(cmdManifest *protos.Command) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		logger := logrus.
			WithContext(ctx).
			WithFields(logrus.Fields{
				"command": cmdManifest.Name,
				"args":    args,
			})

		logger.Debugf("Running command: %s", cmdManifest.Name)

		commandManifest := GetCommandManifest(ctx)
		if commandManifest == nil {
			logger.Error("command manifest is not found in the context")
			return
		}

		// 1: find related plugin with the command
		// 2: download if not installed or version is not matched
		// 3: run command with the necessary args

		pluginReceipt, err := w.pluginManager.GetPluginReceipt(ctx, commandManifest.PluginRef)
		if err != nil {
			logger.WithError(err).Errorf("could not get plugin receipt %q", commandManifest.PluginRef)
			return
		}

		if err := w.exec.Run(ctx, cmdManifest.Name, args...); err != nil {
			logger.Errorf("Error running command: %s", cmdManifest.Name)
		}

		logger.Debugf("Running command: %s", cmdManifest.Name)
	}
}
