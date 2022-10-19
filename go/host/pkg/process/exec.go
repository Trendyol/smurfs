package process

import (
	"context"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
)

type Exec interface {
	// Run executes the command with the given arguments
	Run(ctx context.Context, receipt *plugin.Receipt, args ...string) error
}

type exec struct{}

func NewExec() Exec {
	return &exec{}
}

func (e *exec) Run(ctx context.Context, receipt *plugin.Receipt, args ...string) error {
	/*
		if plugin.Source["type"] == "binary" {
			cmd := proc.Command(plugin.Source["binary"].(string), args...)
			cmd.Stdout = os.Stdout

			err := cmd.Start()

			err = cmd.Wait()

			return err
		}
	*/
	panic("implement me")
}
