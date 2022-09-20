package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host/pkg/process"
	"github.com/trendyol/smurfs/go/host/protos"
)

// CommandWrapper is responsible for wrapping the command manifest to cobra command
type CommandWrapper interface {
	Wrap(*protos.Command) func(cmd *cobra.Command, args []string)
}

type wrapper struct {
	exec process.Exec
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

		if err := w.exec.Run(ctx, cmdManifest.Name, args...); err != nil {
			logger.Errorf("Error running command: %s", cmdManifest.Name)
		}

		logger.Debugf("Running command: %s", cmdManifest.Name)
	}
}
