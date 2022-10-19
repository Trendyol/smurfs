package process

import (
	"context"
	"fmt"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"os"
	"os/exec"
)

type Executor interface {
	// Run executes the command with the given arguments
	Run(ctx context.Context, receipt *plugin.Receipt, args ...string) error
}

type executor struct{}

func NewExec() Executor {
	return &executor{}
}

func (e *executor) Run(ctx context.Context, receipt *plugin.Receipt, args ...string) error {
	executablePath := receipt.Executable.Executable.Entrypoint

	cmd := exec.CommandContext(ctx, executablePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	fmt.Println(err)
	return err
}
