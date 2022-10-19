package cli

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host/pkg/models"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"github.com/trendyol/smurfs/go/host/pkg/process"
)

// CommandWrapper is responsible for wrapping the command manifest to cobra command
type CommandWrapper interface {
	Wrap(plugin2 *plugin.Plugin) func(cmd *cobra.Command, args []string)
}

type wrapper struct {
	exec          process.Executor
	pluginManager plugin.Manager
}

func NewCmdWrapper(exec process.Executor, pluginManager plugin.Manager) CommandWrapper {
	return &wrapper{
		exec:          exec,
		pluginManager: pluginManager,
	}
}

func (w *wrapper) Wrap(plugin *plugin.Plugin) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		logger := logrus.
			WithContext(ctx).
			WithFields(logrus.Fields{
				"command": plugin.Name,
				"args":    args,
			})

		logger.Debugf("Running command: %s", plugin.Name)

		// 1: find related plugin with the command
		// 2: download if not installed or version is not matched
		// 3: run command with the necessary args

		pluginReceipt, err := w.pluginManager.GetPluginReceipt(ctx, plugin.Name)
		if err != nil && errors.Is(err, models.ErrPluginNotInstalled) {
			pluginReceipt, err = w.pluginManager.Install(ctx, *plugin)
			if err != nil {
				logger.WithError(err).Errorf("Failed to install plugin: %s", plugin.Name)
				return
			}
		}

		if err != nil {
			logger.WithError(err).Errorf("could not get plugin receipt %q", plugin.Name)
			return
		}

		if err := w.exec.Run(ctx, &pluginReceipt, args...); err != nil {
			logger.Errorf("Error running command: %s", plugin.Name)
		}

		logger.Debugf("Running command: %s", plugin.Name)
	}
}
