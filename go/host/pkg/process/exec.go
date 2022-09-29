package process

import (
	"context"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"os"
	proc "os/exec"
)

type Exec interface {
	// Run executes the command with the given arguments
	Run(ctx context.Context, plugin plugin.Plugin, args ...string) error
}

type exec struct{}

func NewExec() Exec {
	return &exec{}
}

func (e *exec) Run(ctx context.Context, plugin plugin.Plugin, args ...string) error {
	if plugin.Source["type"] == "binary" {
		cmd := proc.Command(plugin.Source["binary"].(string), args...)
		cmd.Stdout = os.Stdout

		err := cmd.Start()

		err = cmd.Wait()

		return err
	}

	panic("implement me")
}
